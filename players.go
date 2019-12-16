package apivideosdk

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const playersBasePath = "players"

// PlayersServiceI is an interface representing the Players
// endpoints of the api.video API
// See: https://docs.api.video/5.1/players
type PlayersServiceI interface {
	Get(playerID string) (*Player, error)
	List(opts *PlayerOpts) (*PlayerList, error)
	Create(createRequest *PlayerRequest) (*Player, error)
	Update(playerID string, updateRequest *PlayerRequest) (*Player, error)
	Delete(playerID string) error
	UploadLogo(playerID string, link string, filepath string) (*Player, error)
}

// PlayersService communicating with the Players
// endpoints of the api.video API
type PlayersService struct {
	client *Client
}

// Player represents an api.video Player
type Player struct {
	PlayerID              string        `json:"playerId,omitempty"`
	ShapeMargin           int           `json:"shapeMargin,omitempty"`
	ShapeRadius           int           `json:"shapeRadius,omitempty"`
	ShapeAspect           string        `json:"shapeAspect,omitempty"`
	ShapeBackgroundTop    string        `json:"shapeBackgroundTop,omitempty"`
	ShapeBackgroundBottom string        `json:"shapeBackgroundBottom,omitempty"`
	Text                  string        `json:"text,omitempty"`
	Link                  string        `json:"link,omitempty"`
	LinkHover             string        `json:"linkHover,omitempty"`
	LinkActive            string        `json:"linkActive,omitempty"`
	TrackPlayed           string        `json:"trackPlayed,omitempty"`
	TrackUnplayed         string        `json:"trackUnplayed,omitempty"`
	TrackBackground       string        `json:"trackBackground,omitempty"`
	BackgroundTop         string        `json:"backgroundTop,omitempty"`
	BackgroundBottom      string        `json:"backgroundBottom,omitempty"`
	BackgroundText        string        `json:"backgroundText,omitempty"`
	Language              string        `json:"language,omitempty"`
	EnableAPI             bool          `json:"enableApi,omitempty"`
	EnableControls        bool          `json:"enableControls,omitempty"`
	ForceAutoplay         bool          `json:"forceAutoplay,omitempty"`
	HideTitle             bool          `json:"hideTitle,omitempty"`
	ForceLoop             bool          `json:"forceLoop,omitempty"`
	Assets                *PlayerAssets `json:"assets,omitempty"`
}

// PlayerAssets represents ths assets of a Player
type PlayerAssets struct {
	Logo string `json:"logo,omitempty"`
	Link string `json:"link,omitempty"`
}

// PlayerRequest represents a request to create / update a Player
type PlayerRequest struct {
	ShapeMargin           int    `json:"shapeMargin,omitempty"`
	ShapeRadius           int    `json:"shapeRadius,omitempty"`
	ShapeAspect           string `json:"shapeAspect,omitempty"`
	ShapeBackgroundTop    string `json:"shapeBackgroundTop,omitempty"`
	ShapeBackgroundBottom string `json:"shapeBackgroundBottom,omitempty"`
	Text                  string `json:"text,omitempty"`
	Link                  string `json:"link,omitempty"`
	LinkHover             string `json:"linkHover,omitempty"`
	LinkActive            string `json:"linkActive,omitempty"`
	TrackPlayed           string `json:"trackPlayed,omitempty"`
	TrackUnplayed         string `json:"trackUnplayed,omitempty"`
	TrackBackground       string `json:"trackBackground,omitempty"`
	BackgroundTop         string `json:"backgroundTop,omitempty"`
	BackgroundBottom      string `json:"backgroundBottom,omitempty"`
	BackgroundText        string `json:"backgroundText,omitempty"`
	Language              string `json:"language,omitempty"`
	EnableAPI             bool   `json:"enableApi"`
	EnableControls        bool   `json:"enableControls"`
	ForceAutoplay         bool   `json:"forceAutoplay"`
	HideTitle             bool   `json:"hideTitle"`
	ForceLoop             bool   `json:"forceLoop"`
}

//PlayerOpts represents a query string to search on players
type PlayerOpts struct {
	CurrentPage int `url:"currentPage,omitempty"`
	PageSize    int `url:"pageSize,omitempty"`
}

// PlayerList represents a list of player
type PlayerList struct {
	Data       []Player    `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

//Get returns a Player by id
func (s *PlayersService) Get(PlayerID string) (*Player, error) {

	err := checkPlayerID(PlayerID)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s", playersBasePath, PlayerID)

	req, err := s.client.prepareRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	p := new(Player)
	_, err = s.client.do(req, p)

	if err != nil {
		return nil, err
	}

	return p, nil
}

//List returns a PlayerList containing all players
func (s *PlayersService) List(opts *PlayerOpts) (*PlayerList, error) {

	v, err := query.Values(opts)

	if err != nil {
		return nil, err
	}
	qs := v.Encode()

	path := fmt.Sprintf("%s?%s", playersBasePath, qs)

	req, err := s.client.prepareRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	pl := new(PlayerList)
	_, err = s.client.do(req, pl)

	if err != nil {
		return nil, err
	}

	return pl, nil
}

//Create a player and returns it
func (s *PlayersService) Create(createRequest *PlayerRequest) (*Player, error) {

	req, err := s.client.prepareRequest(http.MethodPost, playersBasePath, createRequest)
	if err != nil {
		return nil, err
	}

	p := new(Player)
	_, err = s.client.do(req, p)

	if err != nil {
		return nil, err
	}

	return p, nil
}

//Update a player and returns it
func (s *PlayersService) Update(playerID string, updateRequest *PlayerRequest) (*Player, error) {

	err := checkPlayerID(playerID)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s", playersBasePath, playerID)

	req, err := s.client.prepareRequest(http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, err
	}

	p := new(Player)
	_, err = s.client.do(req, p)

	if err != nil {
		return nil, err
	}

	return p, nil
}

//Delete a video container
func (s *PlayersService) Delete(playerID string) error {

	err := checkPlayerID(playerID)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s", playersBasePath, playerID)

	req, err := s.client.prepareRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.do(req, nil)

	if err != nil {
		return err
	}

	return nil
}

//UploadLogo upload the logo of a player
func (s *PlayersService) UploadLogo(playerID string, link string, filePath string) (*Player, error) {

	err := checkPlayerID(playerID)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s/logo", playersBasePath, playerID)

	fields := map[string]string{
		"link": link,
	}

	req, err := s.client.prepareUploadRequest(path, filePath, fields)

	if err != nil {
		return nil, err
	}

	p := new(Player)

	_, err = s.client.do(req, p)

	if err != nil {
		return nil, err
	}

	return p, nil
}
