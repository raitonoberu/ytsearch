package ytsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		return nil, err
	}
	url := "https://www.youtube.com/youtubei/v1/search?key=" + searchKey
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	request.Header = requestHeader
	response, err := search.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	text, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var objmap map[string]interface{}
	err = json.Unmarshal(text, &objmap)
	if err != nil {
		return nil, err
	}
	return objmap, nil
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
	response, err := search.makeRequest()
	if err != nil {
		return nil, err
	}
	responseSource, continuationKey, estimatedResults, err := parseSource(response, search.newPage)
	if err != nil {
		return nil, err
	}

	result := parseComponents(responseSource)
	result.EstimatedResults = estimatedResults
	result.Suggestions = parseSuggestions(response)

	search.continuationKey = continuationKey
	search.newPage = true

	return result, nil
}
