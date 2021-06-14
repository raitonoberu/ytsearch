package ytsearch

type SearchFilter string

const (
	// NoFilter will search for videos, channels, and playlists.
	// Default for Search.
	NoFilter SearchFilter = "%253D"

	// VideoFilter will search only for videos.
	// Default for VideoSearch.
	VideoFilter SearchFilter = "SAhAB"

	// ChannelFilter will search only for channels
	// Default for ChannelSearch.
	ChannelFilter SearchFilter = "SAhAC"

	// PlaylistFilter will search only for playlist.
	// Default for PlaylistSearch.
	PlaylistFilter SearchFilter = "SAhAD"
)

type SortOrder string

const (
	// RelevanceOrder will sort results by relevancy. Default.
	RelevanceOrder SortOrder = "CAA"

	// UploadDateOrder will sort results by upload date.
	UploadDateOrder SortOrder = "CIA"

	// ViewCountOrder will sort results by view count.
	ViewCountOrder SortOrder = "CMA"

	// RatingOrder will sort results by rating.
	RatingOrder SortOrder = "CEA"
)
