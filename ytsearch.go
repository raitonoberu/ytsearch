// Package ytsearch: search for YouTube videos, channels & playlists. Without YouTube Data API.
package ytsearch

import "net/http"

// Search creates a new SearchClient with default parameters.
func Search(query string) *SearchClient {
	return &SearchClient{
		Query:           query,
		Limit:           0,
		Language:        "en",
		Region:          "US",
		SearchFilter:    NoFilter,
		SortOrder:       RelevanceOrder,
		FindVideos:      true,
		FindChannels:    true,
		FindPlaylists:   true,
		FindShelves:     true,
		FindSuggestions: true,
		HTTPClient:      &http.Client{},
	}
}

// VideoSearch creates a new SearchClient for video search.
func VideoSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Limit:        0,
		Language:     "en",
		Region:       "US",
		SearchFilter: VideoFilter,
		SortOrder:    RelevanceOrder,
		FindVideos:   true,
		HTTPClient:   &http.Client{},
	}
}

// ChannelSearch creates a new SearchClient for channel search.
func ChannelSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Limit:        0,
		Language:     "en",
		Region:       "US",
		SearchFilter: ChannelFilter,
		SortOrder:    RelevanceOrder,
		FindChannels: true,
		HTTPClient:   &http.Client{},
	}
}

// PlaylistSearch creates a new SearchClient for playlist search.
func PlaylistSearch(query string) *SearchClient {
	return &SearchClient{
		Query:         query,
		Limit:         0,
		Language:      "en",
		Region:        "US",
		SearchFilter:  PlaylistFilter,
		SortOrder:     RelevanceOrder,
		FindPlaylists: true,
		HTTPClient:    &http.Client{},
	}
}
