package ytsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SearchClient struct {
	Query string // required

	Language string // default: en
	Region   string // default: US

	SearchFilter SearchFilter
	SortOrder    SortOrder

	// CustomParams allows you to copy and paste params from YouTube.
	CustomParams string

	HTTPClient *http.Client

	continuationKey string
	newPage         bool
}

// makeRequest makes HTTP POST request and returns result as a map.
func (search *SearchClient) makeRequest() (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"query": search.Query,
		"context": map[string]map[string]interface{}{
			"client": {
				"clientName":       "WEB",
				"clientVersion":    "2.20210224.06.00",
				"newVisitorCookie": true,
			},
			"user": {
				"lockedSafetyMode": false,
			},
		},
	}
	if search.CustomParams == "" {
		payload["params"] = fmt.Sprintf("%s%s", search.SortOrder, search.SearchFilter)
	} else {
		payload["params"] = search.CustomParams
	}
	clientData := payload["context"].(map[string]map[string]interface{})["client"]
	clientData["hl"] = search.Language
	clientData["gl"] = search.Region
	if search.continuationKey != "" {
		payload["continuation"] = search.continuationKey
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal payload: %w", err)
	}
	url := "https://www.youtube.com/youtubei/v1/search?key=" + searchKey
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("couldn't create request: %w", err)
	}
	request.Header = requestHeader
	response, err := search.HTTPClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer response.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("couldn't decode JSON: %w", err)
	}
	return result, nil
}

// NextExists returns whether the Next call will return new content.
func (search *SearchClient) NextExists() bool {
	if !search.newPage {
		return true
	}
	if search.continuationKey != "" {
		return true
	}
	return false
}

// Next returns content from the next page.
func (search *SearchClient) Next() (*SearchResult, error) {
	if !search.NextExists() {
		return nil, PageDoesntExistError
	}
	response, err := search.makeRequest()
	if err != nil {
		return nil, fmt.Errorf("makeRequest failed: %w", err)
	}
	responseSource, continuationKey, estimatedResults := parseSource(response, search.newPage)

	result := parseComponents(responseSource)
	result.EstimatedResults = estimatedResults
	result.Suggestions = parseSuggestions(response)

	search.continuationKey = continuationKey
	search.newPage = true

	return result, nil
}
