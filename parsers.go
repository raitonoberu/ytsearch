package ytsearch

import (
	"strconv"
)

type path []interface{}

func getValue(source interface{}, path path) interface{} {
	value := source
	for _, element := range path {
		mustBreak := false
		switch element.(type) {
		case string:
			if val, ok := value.(map[string]interface{})[element.(string)]; ok {
				value = val
			} else {
				value = nil
				mustBreak = true
			}
		case int:
			if len(value.([]interface{})) != 0 {
				value = value.([]interface{})[element.(int)]
			} else {
				value = nil
				mustBreak = true
			}
		}
		if mustBreak {
			break
		}
	}
	return value
}

// parseSource extracts content from response.
func parseSource(response interface{}, newPage bool) ([]map[string]interface{}, string, int) {
	var responseContent []interface{}
	if !newPage {
		content := getValue(response, path{"contents", "twoColumnSearchResultsRenderer", "primaryContents", "sectionListRenderer", "contents"})
		if content != nil {
			responseContent = content.([]interface{})
		} else {
			return nil, "", 0
		}
	} else {
		content := getValue(response, path{"onResponseReceivedCommands", 0, "appendContinuationItemsAction", "continuationItems"})
		if content != nil {
			responseContent = content.([]interface{})
		} else {
			return nil, "", 0
		}
	}

	// converting []interface{} to []map[string]interface{}
	responseContentMaps := make([]map[string]interface{}, len(responseContent))
	for index, value := range responseContent {
		responseContentMaps[index] = value.(map[string]interface{})
	}

	var responseSource []map[string]interface{}
	var continuationKey string

	if responseContent != nil {
		for _, element := range responseContentMaps {
			if _, ok := element["itemSectionRenderer"]; ok {
				newSource := getValue(element, path{"itemSectionRenderer", "contents"}).([]interface{})
				// converting []interface{} to []map[string]interface{}
				responseSource = responseSource[:0]
				for _, value := range newSource {
					responseSource = append(responseSource, value.(map[string]interface{}))
				}
			}
			if _, ok := element["continuationItemRenderer"]; ok {
				continuationKey = getValue(element, path{"continuationItemRenderer", "continuationEndpoint", "continuationCommand", "token"}).(string)
			}
		}
	} else {
		responseSource = getValue(responseContent, path{"contents", "twoColumnSearchResultsRenderer", "primaryContents", "richGridRenderer", "contents"}).([]map[string]interface{})
		continuationKey = getValue(responseSource, path{"continuationItemRenderer", "continuationEndpoint", "continuationCommand", "token"}).(string)
	}

	estimatedResults, _ := strconv.Atoi(
		getValue(response, path{"estimatedResults"}).(string),
	)

	return responseSource, continuationKey, estimatedResults
}

// parseComponents splits source into various components.
func parseComponents(responseSource []map[string]interface{}) *SearchResult {
	result := &SearchResult{}

	for _, element := range responseSource {
		if videoElement, ok := element["videoRenderer"]; ok {
			videoComponent := parseVideoComponent(videoElement.(map[string]interface{}))
			result.Videos = append(result.Videos, videoComponent)
			continue
		}
		if channelElement, ok := element["channelRenderer"]; ok {
			channelComponent := parseChannelComponent(channelElement.(map[string]interface{}))
			result.Channels = append(result.Channels, channelComponent)
			continue
		}
		if playlistElement, ok := element["playlistRenderer"]; ok {
			playlistComponent := parsePlaylistComponent(playlistElement.(map[string]interface{}))
			result.Playlists = append(result.Playlists, playlistComponent)
			continue
		}
		if shelfElement, ok := element["shelfRenderer"]; ok {
			shelfComponent := parseShelfComponent(shelfElement.(map[string]interface{}))
			result.Shelves = append(result.Shelves, shelfComponent)
			continue
		}
		if richItemElement, ok := element["richItemRenderer"]; ok {
			// initial fallback handling for FindVideos
			if richItemElementContent, ok := richItemElement.(map[string]interface{})["content"]; ok {
				if videoElement, ok := richItemElementContent.(map[string]interface{})["videoRenderer"]; ok {
					videoComponent := parseVideoComponent(videoElement.(map[string]interface{}))
					result.Videos = append(result.Videos, videoComponent)
				}
			}
			continue
		}
	}
	return result
}

func parseVideoComponent(video map[string]interface{}) *VideoItem {
	item := &VideoItem{}
	if id := getValue(video, path{"videoId"}); id != nil {
		item.ID = id.(string)
		item.URL = "https://www.youtube.com/watch?v=" + item.ID
	}
	if title := getValue(video, path{"title", "runs", 0, "text"}); title != nil {
		item.Title = title.(string)
	}
	if publishedTime := getValue(video, path{"publishedTimeText", "simpleText"}); publishedTime != nil {
		item.PublishedTime = publishedTime.(string)
	}
	if duration := getValue(video, path{"lengthText", "simpleText"}); duration != nil {
		item.Duration = durationToInt(duration)
	}
	if viewCount := getValue(video, path{"viewCountText", "simpleText"}); viewCount != nil {
		item.ViewCount = viewCountToInt(viewCount)
	}
	if thumbnails := getValue(video, path{"thumbnail", "thumbnails"}); thumbnails != nil {
		for _, thumbnail := range thumbnails.([]interface{}) {
			item.Thumbnails = append(item.Thumbnails, Thumbnail{
				URL:    thumbnail.(map[string]interface{})["url"].(string),
				Height: int(thumbnail.(map[string]interface{})["height"].(float64)),
				Width:  int(thumbnail.(map[string]interface{})["width"].(float64)),
			})
		}
	}
	if richThumbnail := getValue(video, path{"richThumbnail", "movingThumbnailRenderer", "movingThumbnailDetails", "thumbnails", 0}); richThumbnail != nil {
		item.RichThumbnail = Thumbnail{
			URL:    richThumbnail.(map[string]interface{})["url"].(string),
			Height: int(richThumbnail.(map[string]interface{})["height"].(float64)),
			Width:  int(richThumbnail.(map[string]interface{})["width"].(float64)),
		}
	}
	if description := getValue(video, path{"detailedMetadataSnippets", 0, "snippetText", "runs"}); description != nil {
		item.Description = descriptionToStr(description.([]interface{}))
	}
	item.Channel = Channel{}
	if channelTitle := getValue(video, path{"ownerText", "runs", 0, "text"}); channelTitle != nil {
		item.Channel.Title = channelTitle.(string)
	}
	if channelId := getValue(video, path{"ownerText", "runs", 0, "navigationEndpoint", "browseEndpoint", "browseId"}); channelId != nil {
		item.Channel.ID = channelId.(string)
		item.Channel.URL = "https://www.youtube.com/channel/" + item.Channel.ID
	}
	if channelThumbnails := getValue(video, path{"channelThumbnailSupportedRenderers", "channelThumbnailWithLinkRenderer", "thumbnail", "thumbnails"}); channelThumbnails != nil {
		for _, thumbnail := range channelThumbnails.([]interface{}) {
			item.Channel.Thumbnails = append(item.Channel.Thumbnails, Thumbnail{
				URL:    thumbnail.(map[string]interface{})["url"].(string),
				Height: int(thumbnail.(map[string]interface{})["height"].(float64)),
				Width:  int(thumbnail.(map[string]interface{})["width"].(float64)),
			})
		}
	}
	return item
}

func parseChannelComponent(channel map[string]interface{}) *ChannelItem {
	item := &ChannelItem{}
	if id := getValue(channel, path{"channelId"}); id != nil {
		item.ID = id.(string)
		item.URL = "https://www.youtube.com/channel/" + item.ID
	}
	if title := getValue(channel, path{"title", "simpleText"}); title != nil {
		item.Title = title.(string)
	}
	if thumbnails := getValue(channel, path{"thumbnail", "thumbnails"}); thumbnails != nil {
		for _, thumbnail := range thumbnails.([]interface{}) {
			item.Thumbnails = append(item.Thumbnails, Thumbnail{
				URL:    "http:" + thumbnail.(map[string]interface{})["url"].(string),
				Height: int(thumbnail.(map[string]interface{})["height"].(float64)),
				Width:  int(thumbnail.(map[string]interface{})["width"].(float64)),
			})
		}
	}
	if videoCount := getValue(channel, path{"videoCountText", "runs", 0, "text"}); videoCount != nil {
		item.VideoCount, _ = strconv.Atoi(videoCount.(string))
	}
	if description := getValue(channel, path{"descriptionSnippet", "runs"}); description != nil {
		item.Description = descriptionToStr(description.([]interface{}))
	}
	if subscribers := getValue(channel, path{"subscriberCountText", "simpleText"}); subscribers != nil {
		item.Subscribers = subscribers.(string)
	}
	return item
}

func parsePlaylistComponent(playlist map[string]interface{}) *PlaylistItem {
	item := &PlaylistItem{}
	if id := getValue(playlist, path{"playlistId"}); id != nil {
		item.ID = id.(string)
		item.URL = "https://www.youtube.com/playlist?list=" + item.ID
	}
	if title := getValue(playlist, path{"title", "simpleText"}); title != nil {
		item.Title = title.(string)
	}
	if videoCount := getValue(playlist, path{"videoCount"}); videoCount != nil {
		item.VideoCount, _ = strconv.Atoi(videoCount.(string))
	}
	item.Channel = Channel{}
	if channelTitle := getValue(playlist, path{"shortBylineText", "runs", 0, "text"}); channelTitle != nil {
		item.Channel.Title = channelTitle.(string)
	}
	if channelId := getValue(playlist, path{"shortBylineText", "runs", 0, "navigationEndpoint", "browseEndpoint", "browseId"}); channelId != nil {
		item.Channel.ID = channelId.(string)
		item.Channel.URL = "https://www.youtube.com/channel/" + item.Channel.ID
	}
	if thumbnails := getValue(playlist, path{"thumbnailRenderer", "playlistVideoThumbnailRenderer", "thumbnail", "thumbnails"}); thumbnails != nil {
		for _, thumbnail := range thumbnails.([]interface{}) {
			item.Thumbnails = append(item.Thumbnails, Thumbnail{
				URL:    "http:" + thumbnail.(map[string]interface{})["url"].(string),
				Height: int(thumbnail.(map[string]interface{})["height"].(float64)),
				Width:  int(thumbnail.(map[string]interface{})["width"].(float64)),
			})
		}
	}
	return item
}

func parseShelfComponent(shelf map[string]interface{}) *ShelfItem {
	item := &ShelfItem{}
	if title := getValue(shelf, path{"title", "simpleText"}); title != nil {
		item.Title = title.(string)
	}
	items := getValue(shelf, path{"content", "verticalListRenderer", "items"})
	for _, shelfItem := range items.([]interface{}) {
		if videoElement, ok := shelfItem.(map[string]interface{})["videoRenderer"]; ok {
			videoComponent := parseVideoComponent(videoElement.(map[string]interface{}))
			item.Items = append(item.Items, videoComponent)
		}
	}
	return item
}

func parseSuggestions(response map[string]interface{}) []string {
	suggestions := response["refinements"]
	if suggestions == nil {
		return []string{}
	}
	result := make([]string, len(suggestions.([]interface{})))
	for index, item := range suggestions.([]interface{}) {
		result[index] = item.(string)
	}
	return result
}
