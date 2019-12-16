package apivideosdk

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const statisticsBasePath = "analytics"

// StatisticsServiceI is an interface representing the Players
// endpoints of the api.video API
// See: https://docs.api.video/5.1/players
type StatisticsServiceI interface {
	GetVideoSessions(videoID string, opts *SessionVideoOpts) (*StatisticList, error)
	GetLivestreamSessions(LivestreamID string, opts *SessionLivestreamOpts) (*StatisticList, error)
	GetSessionEvents(SessionID string, opts *SessionEventOpts) (*SessionEventList, error)
}

// StatisticsService communicating with the Statistics
// endpoints of the api.video API
type StatisticsService struct {
	client *Client
}

//StatisticList represents a list of staticics
type StatisticList struct {
	Data       []Statistic `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

//Statistic represents a Statistic
type Statistic struct {
	Session    *Session    `json:"session,omitempty"`
	Location   *Location   `json:"location,omitempty"`
	Referrer   *Referrer   `json:"referrer,omitempty"`
	Device     *Device     `json:"device,omitempty"`
	Os         *Os         `json:"os,omitempty"`
	SessClient *SessClient `json:"client,omitempty"`
}

//Session represents a Session
type Session struct {
	SessionID string            `json:"sessionId,omitempty"`
	LoadedAt  string            `json:"loadedAt,omitempty"`
	EndedAt   string            `json:"endedAt,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

//Location represents a Location
type Location struct {
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
}

//Referrer represents a Referrer
type Referrer struct {
	URL        string `json:"url,omitempty"`
	Medium     string `json:"medium,omitempty"`
	Source     string `json:"source,omitempty"`
	SearchTerm string `json:"searchTerm,omitempty"`
}

//Device represents a Device
type Device struct {
	Type   string `json:"type,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	Model  string `json:"model,omitempty"`
}

//Os represents a Os
type Os struct {
	Name      string `json:"name,omitempty"`
	Shortname string `json:"shortname,omitempty"`
	Version   string `json:"version,omitempty"`
}

//SessClient represents a SessClient
type SessClient struct {
	Type    string `json:"type,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

//SessionEvent represents a SessionEvent
type SessionEvent struct {
	Type      string `json:"type,omitempty"`
	EmittedAt string `json:"emittedAt,omitempty"`
	At        int    `json:"at,omitempty"`
	From      int    `json:"from,omitempty"`
	To        int    `json:"to,omitempty"`
}

//SessionEventList represents a SessionEventList
type SessionEventList struct {
	Data       []SessionEvent `json:"data,omitempty"`
	Pagination *Pagination    `json:"pagination,omitempty"`
}

//SessionVideoOpts represents a query string to search on videos statistics
type SessionVideoOpts struct {
	CurrentPage int               `url:"currentPage,omitempty"`
	PageSize    int               `url:"pageSize,omitempty"`
	Period      string            `url:"period,omitempty"`
	Metadata    map[string]string `url:"-"`
}

//SessionLivestreamOpts represents a query string to search on livestream statistics
type SessionLivestreamOpts struct {
	CurrentPage int    `url:"currentPage,omitempty"`
	PageSize    int    `url:"pageSize,omitempty"`
	Period      string `url:"period,omitempty"`
}

//SessionEventOpts represents a query string to search on sessions
type SessionEventOpts struct {
	CurrentPage int `url:"currentPage,omitempty"`
	PageSize    int `url:"pageSize,omitempty"`
}

//GetVideoSessions returns a StatisticList containing all sessions for a video
func (s *StatisticsService) GetVideoSessions(videoID string, opts *SessionVideoOpts) (*StatisticList, error) {

	err := checkVideoID(videoID)
	if err != nil {
		return nil, err
	}

	v, err := query.Values(opts)

	if err != nil {
		return nil, err
	}
	qs := v.Encode()

	if len(opts.Metadata) > 0 {
		u := url.Values{}
		for k, v := range opts.Metadata {
			u.Add(fmt.Sprintf("metadata[%s]", k), v)
		}
		qs = fmt.Sprintf("%s&%s", qs, u.Encode())
	}

	path := fmt.Sprintf("%s/videos/%s?%s", statisticsBasePath, videoID, qs)

	req, err := s.client.prepareRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	sl := new(StatisticList)
	_, err = s.client.do(req, sl)

	if err != nil {
		return nil, err
	}

	return sl, nil
}

//GetLivestreamSessions returns a StatisticList containing all sessions for a video
func (s *StatisticsService) GetLivestreamSessions(livestreamID string, opts *SessionLivestreamOpts) (*StatisticList, error) {

	err := checkLivestreamID(livestreamID)
	if err != nil {
		return nil, err
	}

	v, err := query.Values(opts)

	if err != nil {
		return nil, err
	}
	qs := v.Encode()

	path := fmt.Sprintf("%s/live-streams/%s?%s", statisticsBasePath, livestreamID, qs)

	req, err := s.client.prepareRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	sl := new(StatisticList)
	_, err = s.client.do(req, sl)

	if err != nil {
		return nil, err
	}

	return sl, nil
}

//GetSessionEvents returns a StatisticList containing all stats for one session
func (s *StatisticsService) GetSessionEvents(sessionID string, opts *SessionEventOpts) (*SessionEventList, error) {

	err := checkSessionID(sessionID)
	if err != nil {
		return nil, err
	}

	v, err := query.Values(opts)

	if err != nil {
		return nil, err
	}
	qs := v.Encode()

	path := fmt.Sprintf("%s/sessions/%s/events?%s", statisticsBasePath, sessionID, qs)

	req, err := s.client.prepareRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	sel := new(SessionEventList)
	_, err = s.client.do(req, sel)

	if err != nil {
		return nil, err
	}

	return sel, nil
}
