package ytsearch

type SearchResult struct {
	EstimatedResults int             `json:"estimatedResults"`
	Videos           []*VideoItem    `json:"videos"`
	Channels         []*ChannelItem  `json:"channels"`
	Playlists        []*PlaylistItem `json:"playlists"`
	Shelves          []*ShelfItem    `json:"shelves"`
	Suggestions      []string        `json:"suggestions"`
}

type VideoItem struct {
	ID            string      `json:"id"`
	Title         string      `json:"title"`
	PublishedTime string      `json:"publishedTime"`
	Duration      int         `json:"duration"`
	ViewCount     int         `json:"viewCount"`
	Thumbnails    []Thumbnail `json:"thumbnails"`
	RichThumbnail Thumbnail   `json:"richThumbnail"`
	Description   string      `json:"description"`
	Channel       Channel     `json:"channel"`
	URL           string      `json:"url"`
}

type ChannelItem struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Thumbnails  []Thumbnail `json:"thumbnails"`
	VideoCount  int         `json:"videoCount"`
	Description string      `json:"description"`
	Subscribers string      `json:"subscribers"`
	URL         string      `json:"url"`
}

type PlaylistItem struct {
	ID         string      `json:"id"`
	Title      string      `json:"title"`
	VideoCount int         `json:"videoCount"`
	Channel    Channel     `json:"channel"`
	Thumbnails []Thumbnail `json:"thumbnails"`
	URL        string      `json:"url"`
}

type ShelfItem struct {
	Title string       `json:"title"`
	Items []*VideoItem `json:"items"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Channel struct {
	ID         string      `json:"id"`
	Title      string      `json:"title"`
	Thumbnails []Thumbnail `json:"thumbnails"`
	URL        string      `json:"url"`
}
