package youtube_search

// TODO:
// • Better exceptions handling
// • Extra features

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
	}
}
