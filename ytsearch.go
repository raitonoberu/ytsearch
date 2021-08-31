// Package ytsearch: search for YouTube videos, channels & playlists. Without YouTube Data API.
package ytsearch

import "net/http"

// Search creates a new SearchClient with default parameters.
func Search(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: NoFilter,
		SortOrder:    RelevanceOrder,
		HTTPClient:   &http.Client{},
	}
}

// VideoSearch creates a new SearchClient for video search.
func VideoSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: VideoFilter,
		SortOrder:    RelevanceOrder,
		HTTPClient:   &http.Client{},
	}
}

// ChannelSearch creates a new SearchClient for channel search.
func ChannelSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: ChannelFilter,
		SortOrder:    RelevanceOrder,
		HTTPClient:   &http.Client{},
	}
}

// PlaylistSearch creates a new SearchClient for playlist search.
func PlaylistSearch(query string) *SearchClient {
	return &SearchClient{
		Query:        query,
		Language:     "en",
		Region:       "US",
		SearchFilter: PlaylistFilter,
		SortOrder:    RelevanceOrder,
		HTTPClient:   &http.Client{},
	}
}
