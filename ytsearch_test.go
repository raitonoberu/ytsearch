package ytsearch

import (
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	search := Search("ncs")
	result, err := search.Next()
	if err != nil {
		t.Error(err)
	}

	if result.EstimatedResults == 0 {
		t.Error("result.EstimatedResults == 0")
	}
	if len(result.Videos) == 0 {
		t.Error("len(result.Videos) == 0")
	}
	if len(result.Channels) == 0 {
		t.Error("len(result.Channels) == 0")
	}
	if len(result.Playlists) == 0 {
		t.Error("len(result.Playlists) == 0")
	}
	if len(result.Shelves) == 0 {
		t.Error("len(result.Shelves) == 0")
	}
	if len(result.Suggestions) == 0 {
		t.Error("len(result.Suggestions) == 0")
	}

	_, err = search.Next()
	if err != nil {
		t.Error(err)
	}
}

func TestVideoSearch(t *testing.T) {
	search := VideoSearch("GopherCon 2020")
	result, err := search.Next()
	if err != nil {
		t.Error(err)
	}

	if len(result.Videos) != 20 {
		t.Error("len(result.Videos) != 20")
	}
	if result.Videos[0].Id == "" {
		t.Error("result.Videos[0].Id == \"\"")
	}
}

func TestChannelSearch(t *testing.T) {
	search := ChannelSearch("Programming")
	result, err := search.Next()
	if err != nil {
		t.Error(err)
	}

	if len(result.Channels) != 20 {
		t.Error("len(result.Channels) != 20")
	}
	if result.Channels[0].Id == "" {
		t.Error("result.Channels[0].Id == \"\"")
	}
}

func TestPlaylistSearch(t *testing.T) {
	search := PlaylistSearch("Music playlist")
	result, err := search.Next()
	if err != nil {
		t.Error(err)
	}

	if len(result.Playlists) != 20 {
		t.Error("len(result.Playlists) != 20")
	}
	if result.Playlists[0].Id == "" {
		t.Error("result.Playlists[0].Id == \"\"")
	}
}

func TestRegionLanguage(t *testing.T) {
	search := VideoSearch("no copyright sound")
	search.Region = "RU"
	search.Language = "ru"
	result, err := search.Next()
	if err != nil {
		t.Error(err)
	}

	if len(result.Videos) == 0 {
		t.Error("len(result.Videos) == 0")
	}
	if !strings.Contains(result.Videos[0].PublishedTime, "назад") {
		t.Error("Region & Language params don't work")
	}
}

func TestSortOrder(t *testing.T) {
	search1 := VideoSearch("ncs")
	search1.SortOrder = UploadDateOrder
	search2 := VideoSearch("ncs")
	search2.SortOrder = ViewCountOrder

	result1, err := search1.Next()
	if err != nil {
		t.Error(err)
	}
	if len(result1.Videos) == 0 {
		t.Error("len(result1.Videos) == 0")
	}
	result2, err := search2.Next()
	if err != nil {
		t.Error(err)
	}
	if len(result2.Videos) == 0 {
		t.Error("len(result2.Videos) == 0")
	}

	if result1.Videos[0].Id == result2.Videos[0].Id {
		t.Error("SortOrder param doesn't work")
	}
}
