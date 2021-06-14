package ytsearch

import (
	"reflect"
	"strconv"
)

func getValue(source interface{}, path []interface{}) interface{} {
	value := source
	for _, element := range path {
		mustBreak := false
		switch reflect.TypeOf(element).Kind() {
		case reflect.String:
			if val, ok := value.(map[string]interface{})[element.(string)]; ok {
				value = val
			} else {
				value = nil
				mustBreak = true
			}
		case reflect.Int:
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

func getVideoComponent(video map[string]interface{}) *VideoItem {
	item := &VideoItem{}
	if id := getValue(video, []interface{}{"videoId"}); id != nil {
		item.Id = id.(string)
		item.Link = "https://www.youtube.com/watch?v=" + item.Id
	}
	if title := getValue(video, []interface{}{"title", "runs", 0, "text"}); title != nil {
		item.Title = title.(string)
	}
	if publishedTime := getValue(video, []interface{}{"publishedTimeText", "simpleText"}); publishedTime != nil {
		item.PublishedTime = publishedTime.(string)
	}
	if duration := getValue(video, []interface{}{"lengthText", "simpleText"}); duration != nil {
		item.Duration = durationToInt(duration)
	}
	if viewCount := getValue(video, []interface{}{"viewCountText", "simpleText"}); viewCount != nil {
		item.ViewCount = viewCountToInt(viewCount)
	}
	if thumbnails := getValue(video, []interface{}{"thumbnail", "thumbnails"}); thumbnails != nil {
		for _, thumbnail := range thumbnails.([]interface{}) {
			item.Thumbnails = append(item.Thumbnails, &Thumbnail{
				Url:    thumbnail.(map[string]interface{})["url"].(string),
				Height: int(thumbnail.(map[string]interface{})["height"].(float64)),
				Width:  int(thumbnail.(map[string]interface{})["width"].(float64)),
			})
		}
	}
	if richThumbnail := getValue(video, []interface{}{"richThumbnail", "movingThumbnailRenderer", "movingThumbnailDetails", "thumbnails", 0}); richThumbnail != nil {
		item.RichThumbnail = &Thumbnail{
			Url:    richThumbnail.(map[string]interface{})["url"].(string),
			Height: int(richThumbnail.(map[string]interface{})["height"].(float64)),
			Width:  int(richThumbnail.(map[string]interface{})["width"].(float64)),
		}
	}
	if description := getValue(video, []interface{}{"detailedMetadataSnippets", 0, "snippetText", "runs"}); description != nil {
		item.Description = descriptionToStr(description.([]interface{}))
	}
	item.Channel = &Channel{}
	if channelTitle := getValue(video, []interface{}{"ownerText", "runs", 0, "text"}); channelTitle != nil {
		item.Channel.Title = channelTitle.(string)
	}
	if channelId := getValue(video, []interface{}{"ownerText", "runs", 0, "navigationEndpoint", "browseEndpoint", "browseId"}); channelId != nil {
		item.Channel.Id = channelId.(string)
		item.Channel.Link = "https://www.youtube.com/channel/" + item.Channel.Id
	}
	if channelThumbnails := getValue(video, []interface{}{"channelThumbnailSupportedRenderers", "channelThumbnailWithLinkRenderer", "thumbnail", "thumbnails"}); channelThumbnails != nil {
		for _, thumbnail := range channelThumbnails.([]interface{}) {
			item.Channel.Thumbnails = append(item.Channel.Thumbnails, &Thumbnail{
				Url:    thumbnail.(map[string]interface{})["url"].(string),
				Height: int(thumbnail.(map[string]interface{})["height"].(float64)),
				Width:  int(thumbnail.(map[string]interface{})["width"].(float64)),
			})
		}
	}
	return item
}

func getChannelComponent(channel map[string]interface{}) *ChannelItem {
	item := &ChannelItem{}
	if id := getValue(channel, []interface{}{"channelId"}); id != nil {
		item.Id = id.(string)
		item.Link = "https://www.youtube.com/channel/" + item.Id
	}
	if title := getValue(channel, []interface{}{"title", "simpleText"}); title != nil {
		item.Title = title.(string)
	}
	if thumbnails := getValue(channel, []interface{}{"thumbnail", "thumbnails"}); thumbnails != nil {
		for _, thumbnail := range thumbnails.([]interface{}) {
			item.Thumbnails = append(item.Thumbnails, &Thumbnail{
				Url:    "http:" + thumbnail.(map[string]interface{})["url"].(string),
				Height: int(thumbnail.(map[string]interface{})["height"].(float64)),
				Width:  int(thumbnail.(map[string]interface{})["width"].(float64)),
			})
		}
	}
	if videoCount := getValue(channel, []interface{}{"videoCountText", "runs", 0, "text"}); videoCount != nil {
		item.VideoCount, _ = strconv.Atoi(videoCount.(string))
	}
	if description := getValue(channel, []interface{}{"descriptionSnippet", "runs"}); description != nil {
		item.Description = descriptionToStr(description.([]interface{}))
	}
	if subscribers := getValue(channel, []interface{}{"subscriberCountText", "simpleText"}); subscribers != nil {
		item.Subscribers = subscribers.(string)
	}
	return item
}

func getPlaylistComponent(playlist map[string]interface{}) *PlaylistItem {
	item := &PlaylistItem{}
	if id := getValue(playlist, []interface{}{"playlistId"}); id != nil {
		item.Id = id.(string)
		item.Link = "https://www.youtube.com/playlist?list=" + item.Id
	}
	if title := getValue(playlist, []interface{}{"title", "simpleText"}); title != nil {
		item.Title = title.(string)
	}
	if videoCount := getValue(playlist, []interface{}{"videoCount"}); videoCount != nil {
		item.VideoCount, _ = strconv.Atoi(videoCount.(string))
	}
	item.Channel = &Channel{}
	if channelTitle := getValue(playlist, []interface{}{"shortBylineText", "runs", 0, "text"}); channelTitle != nil {
		item.Channel.Title = channelTitle.(string)
	}
	if channelId := getValue(playlist, []interface{}{"shortBylineText", "runs", 0, "navigationEndpoint", "browseEndpoint", "browseId"}); channelId != nil {
		item.Channel.Id = channelId.(string)
		item.Channel.Link = "https://www.youtube.com/channel/" + item.Channel.Id
	}
	if thumbnails := getValue(playlist, []interface{}{"thumbnailRenderer", "playlistVideoThumbnailRenderer", "thumbnail", "thumbnails"}); thumbnails != nil {
		for _, thumbnail := range thumbnails.([]interface{}) {
			item.Thumbnails = append(item.Thumbnails, &Thumbnail{
				Url:    "http:" + thumbnail.(map[string]interface{})["url"].(string),
				Height: int(thumbnail.(map[string]interface{})["height"].(float64)),
				Width:  int(thumbnail.(map[string]interface{})["width"].(float64)),
			})
		}
	}
	return item
}

func getShelfComponent(shelf map[string]interface{}) *ShelfItem {
	item := &ShelfItem{}
	if title := getValue(shelf, []interface{}{"title", "simpleText"}); title != nil {
		item.Title = title.(string)
	}
	items := getValue(shelf, []interface{}{"content", "verticalListRenderer", "items"})
	for _, shelfItem := range items.([]interface{}) {
		if videoElement, ok := shelfItem.(map[string]interface{})[videoElementKey]; ok {
			videoComponent := getVideoComponent(videoElement.(map[string]interface{}))
			item.Items = append(item.Items, videoComponent)
		}
	}
	return item
}

func getSuggestions(response map[string]interface{}) []string {
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
