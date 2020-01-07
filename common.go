package apivideosdk

import (
	"fmt"
	"regexp"
	"strings"
)

//Pagination represents a pagination object
type Pagination struct {
	CurrentPage      int    `json:"currentPage,omitempty"`
	PageSize         int    `json:"pageSize,omitempty"`
	PagesTotal       int    `json:"pagesTotal,omitempty"`
	ItemsTotal       int    `json:"itemsTotal,omitempty"`
	CurrentPageItems int    `json:"currentPageItems,omitempty"`
	Links            []Link `json:"links,omitempty"`
}

//Link represents a link
type Link struct {
	Rel string `json:"rel,omitempty"`
	URI string `json:"uri,omitempty"`
}

func checkVideoID(videoID string) error {
	if !strings.HasPrefix(videoID, "vi") {
		return fmt.Errorf("Video id %s is invalid, it must start with 'vi'", videoID)
	}
	return nil
}

func checkTimecode(timecode string) error {

	var rxPat = regexp.MustCompile(`^[0-9]{2}(:[0-9]{2}){3}$`)

	if !rxPat.MatchString(timecode) {
		return fmt.Errorf("Timecode format is invalid, it must of type '00:00:00:00'")
	}
	return nil
}

func checkOpts(opts *VideoOpts) error {

	var rxPat = regexp.MustCompile(`^(publishedAt|updatedAt|title)$`)

	if opts.SortBy != "" && !rxPat.MatchString(opts.SortBy) {
		return fmt.Errorf("SortBy value is invalid, it must be 'publishedAt', 'updatedAt' or 'title'")
	}

	rxPat = regexp.MustCompile(`^(asc|desc)$`)
	if opts.SortOrder != "" && !rxPat.MatchString(opts.SortOrder) {
		return fmt.Errorf("SortOrder value is invalid, it must be 'asc' or 'desc'")
	}

	return nil
}

func checkPlayerID(PlayerID string) error {
	if !strings.HasPrefix(PlayerID, "pl") && !strings.HasPrefix(PlayerID, "pt") {
		return fmt.Errorf("Player id %s is invalid, it must start with 'pl' or 'pt'", PlayerID)
	}
	return nil
}

func checkLivestreamID(livestreamID string) error {
	if !strings.HasPrefix(livestreamID, "li") {
		return fmt.Errorf("Livestream id %s is invalid, it must start with 'li'", livestreamID)
	}
	return nil
}

func checkSessionID(sessionID string) error {
	if !strings.HasPrefix(sessionID, "ps") {
		return fmt.Errorf("Session id %s is invalid, it must start with 'ps'", sessionID)
	}
	return nil
}
