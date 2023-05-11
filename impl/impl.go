package impl

import "strings"

type PostGrid struct {
	apiKey  string
	baseURL string
	live    bool
}

func New(apiKey string) *PostGrid {
	live := false

	if strings.HasPrefix(apiKey, "live_") {
		live = true
	}

	return &PostGrid{
		apiKey: apiKey,
		live:   live,
	}
}

func (pg *PostGrid) IsLive() bool {
	return pg.live
}
