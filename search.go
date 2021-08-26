package ytsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SearchClient struct {
	Query string // required

	Limit    int    // default: 0
	Language string // default: en
	Region   string // default: US

	SearchFilter SearchFilter
	SortOrder    SortOrder

	// CustomParams allows you to copy and paste params from YouTube.
	CustomParams string

	FindVideos      bool
	FindChannels    bool
	FindPlaylists   bool
	FindShelves     bool
	FindSuggestions bool

	HTTPClient *http.Client

	continuationKey string
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

// parseSource extracts content from response.
func (search *SearchClient) parseSource(
	response map[string]interface{},
) ([]map[string]interface{}, int, error) {
	var responseContent []interface{}
	if search.continuationKey == "" {
		content := getValue(response, contentPath)
		if content != nil {
			responseContent = content.([]interface{})
		} else {
			return nil, 0, &PageDoesntExistError{}
		}
	} else {
		content := getValue(response, continuationContentPath)
		if content != nil {
			responseContent = content.([]interface{})
		} else {
			return nil, 0, &PageDoesntExistError{}
		}
	}

	// converting []interface{} to []map[string]interface{}
	responseContentMaps := make([]map[string]interface{}, len(responseContent))
	for index, value := range responseContent {
		responseContentMaps[index] = value.(map[string]interface{})
	}

	var responseSource []map[string]interface{}
	if responseContent != nil {
		for _, element := range responseContentMaps {
			if _, ok := element[itemSectionKey]; ok {
				newSource := getValue(element, path{itemSectionKey, "contents"}).([]interface{})
				// converting []interface{} to []map[string]interface{}
				responseSource = responseSource[:0]
				for _, value := range newSource {
					responseSource = append(responseSource, value.(map[string]interface{}))
				}
			}
			if _, ok := element[continuationItemKey]; ok {
				search.continuationKey = getValue(element, continuationKeyPath).(string)
			}
		}
	} else {
		responseSource = getValue(responseContent, fallbackContentPath).([]map[string]interface{})
		search.continuationKey = getValue(responseSource, continuationKeyPath).(string)
	}

	estimatedResults, _ := strconv.Atoi(
		getValue(response, path{"estimatedResults"}).(string),
	)

	return responseSource, estimatedResults, nil
}

// getComponents splits source into various components.
func (search *SearchClient) getComponents(responseSource []map[string]interface{}) *SearchResult {
	result := &SearchResult{}
	if search.FindVideos {
		result.Videos = []*VideoItem{}
	}
	if search.FindChannels {
		result.Channels = []*ChannelItem{}
	}
	if search.FindPlaylists {
		result.Playlists = []*PlaylistItem{}
	}
	if search.FindShelves {
		result.Shelves = []*ShelfItem{}
	}

	for _, element := range responseSource {
		if videoElement, ok := element[videoElementKey]; ok {
			if search.FindVideos {
				videoComponent := getVideoComponent(videoElement.(map[string]interface{}))
				result.Videos = append(result.Videos, videoComponent)
				if search.Limit != 0 && len(result.Videos) >= search.Limit {
					break
				}
			}
			continue
		}
		if channelElement, ok := element[channelElementKey]; ok {
			if search.FindChannels {
				channelComponent := getChannelComponent(channelElement.(map[string]interface{}))
				result.Channels = append(result.Channels, channelComponent)
				if search.Limit != 0 && len(result.Channels) >= search.Limit {
					break
				}
			}
			continue
		}
		if playlistElement, ok := element[playlistElementKey]; ok {
			if search.FindPlaylists {
				playlistComponent := getPlaylistComponent(playlistElement.(map[string]interface{}))
				result.Playlists = append(result.Playlists, playlistComponent)
				if search.Limit != 0 && len(result.Playlists) >= search.Limit {
					break
				}
			}
			continue
		}
		if shelfElement, ok := element[shelfElementKey]; ok {
			if search.FindShelves {
				shelfComponent := getShelfComponent(shelfElement.(map[string]interface{}))
				result.Shelves = append(result.Shelves, shelfComponent)
				if search.Limit != 0 && len(result.Shelves) >= search.Limit {
					break
				}
			}
			continue
		}
		if richItemElement, ok := element[richItemKey]; ok {
			if search.FindVideos {
				// initial fallback handling for FindVideos
				if richItemElementContent, ok := richItemElement.(map[string]interface{})["content"]; ok {
					if videoElement, ok := richItemElementContent.(map[string]interface{})[videoElementKey]; ok {
						videoComponent := getVideoComponent(videoElement.(map[string]interface{}))
						result.Videos = append(result.Videos, videoComponent)
					}
				}
				if search.Limit != 0 && len(result.Videos) >= search.Limit {
					break
				}
			}
			continue
		}
	}
	return result
}

// Next returns content from the next page.
func (search *SearchClient) Next() (*SearchResult, error) {
	response, err := search.makeRequest()
	if err != nil {
		return nil, err
	}
	responseSource, estimatedResults, err := search.parseSource(response)
	if err != nil {
		return nil, err
	}

	result := search.getComponents(responseSource)
	result.EstimatedResults = estimatedResults
	if search.FindSuggestions {
		result.Suggestions = getSuggestions(response)
	}
	return result, nil
}
