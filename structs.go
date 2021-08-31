package ytsearch

type SearchResult struct {
	EstimatedResults int             `json:"estimated_results"`
	Videos           []*VideoItem    `json:"videos"`
	Channels         []*ChannelItem  `json:"channels"`
	Playlists        []*PlaylistItem `json:"playlists"`
	Shelves          []*ShelfItem    `json:"shelves"`
	Suggestions      []string        `json:"suggestions"`
}

type VideoItem struct {
	Id            string      `json:"id"`
	Title         string      `json:"title"`
	PublishedTime string      `json:"published_time"`
	Duration      int         `json:"duration"`
	ViewCount     int         `json:"view_count"`
	Thumbnails    []Thumbnail `json:"thumbnails"`
	RichThumbnail Thumbnail   `json:"rich_thumbnail"`
	Description   string      `json:"description"`
	Channel       Channel     `json:"channel"`
	Link          string      `json:"link"`
}

type ChannelItem struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	Thumbnails  []Thumbnail `json:"thumbnails"`
	VideoCount  int         `json:"videoCount"`
	Description string      `json:"description"`
	Subscribers string      `json:"subscribers"`
	Link        string      `json:"link"`
}

type PlaylistItem struct {
	Id         string      `json:"id"`
	Title      string      `json:"title"`
	VideoCount int         `json:"videoCount"`
	Channel    Channel     `json:"channel"`
	Thumbnails []Thumbnail `json:"thumbnails"`
	Link       string      `json:"link"`
}

type ShelfItem struct {
	Title string       `json:"title"`
	Items []*VideoItem `json:"items"`
}

type Thumbnail struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Channel struct {
	Id         string       `json:"id"`
	Title      string       `json:"title"`
	Thumbnails []Thumbnail `json:"thumbnails,omitempty"`
	Link       string       `json:"link"`
}
