package youtube_search

var requestHeader = map[string][]string{
	"Content-Type": {"application/json; charset=utf-8"},
	"User-Agent": {
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
	},
}

const searchKey = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"
const itemSectionKey = "itemSectionRenderer"
const videoElementKey = "videoRenderer"
const channelElementKey = "channelRenderer"
const playlistElementKey = "playlistRenderer"
const shelfElementKey = "shelfRenderer"
const richItemKey = "richItemRenderer"
const continuationItemKey = "continuationItemRenderer"

var contentPath = []interface{}{"contents", "twoColumnSearchResultsRenderer", "primaryContents", "sectionListRenderer", "contents"}

var fallbackContentPath = []interface{}{"contents", "twoColumnSearchResultsRenderer", "primaryContents", "richGridRenderer", "contents"}

var continuationContentPath = []interface{}{"onResponseReceivedCommands", 0, "appendContinuationItemsAction", "continuationItems"}

var continuationKeyPath = []interface{}{"continuationItemRenderer", "continuationEndpoint", "continuationCommand", "token"}
