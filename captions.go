package apivideosdk

import (
	"fmt"
	"net/http"
)

// CaptionsServiceI is an interface representing the Captions
// endpoints of the api.video API
// See: https://docs.api.video/5.1/captions
type CaptionsServiceI interface {
	Get(videoID string, language string) (*Caption, error)
	List(videoID string) (*CaptionList, error)
	Upload(videoID string, language string, filepath string) (*Caption, error)
	Update(videoID string, language string, updateRequest *CaptionRequest) (*Caption, error)
	Delete(videoID string, language string) error
}

// CaptionsService communicating with the Captions
// endpoints of the api.video API
type CaptionsService struct {
	client *Client
}

// Caption represents an api.video Caption
type Caption struct {
	URI     string `json:"uri,omitempty"`
	Src     string `json:"src,omitempty"`
	Srclang string `json:"srclang,omitempty"`
	Default bool   `json:"default,omitempty"`
}

// CaptionList represents a list of captions
type CaptionList struct {
	Data       []Caption   `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// CaptionRequest represents a request to update a Caption
type CaptionRequest struct {
	Default bool `json:"default"`
}

//Get returns a Caption by video id and language
func (s *CaptionsService) Get(videoID string, language string) (*Caption, error) {

	err := checkVideoID(videoID)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s/captions/%s", videosBasePath, videoID, language)

	req, err := s.client.prepareRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	c := new(Caption)
	_, err = s.client.do(req, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

//List returns a slice of Caption containing all captions for a videoId
func (s *CaptionsService) List(videoID string) (*CaptionList, error) {

	err := checkVideoID(videoID)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s/captions", videosBasePath, videoID)

	req, err := s.client.prepareRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	c := new(CaptionList)
	_, err = s.client.do(req, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

//Upload a vtt for a video and language
func (s *CaptionsService) Upload(videoID string, language string, filePath string) (*Caption, error) {

	err := checkVideoID(videoID)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s/captions/%s", videosBasePath, videoID, language)

	req, err := s.client.prepareUploadRequest(path, filePath, nil)

	if err != nil {
		return nil, err
	}

	c := new(Caption)

	_, err = s.client.do(req, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

//Update a video container and returns it
func (s *CaptionsService) Update(videoID string, language string, updateRequest *CaptionRequest) (*Caption, error) {

	err := checkVideoID(videoID)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s/captions/%s", videosBasePath, videoID, language)

	req, err := s.client.prepareRequest(http.MethodPatch, path, updateRequest)
	if err != nil {
		return nil, err
	}

	c := new(Caption)
	_, err = s.client.do(req, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

//Delete a caption
func (s *CaptionsService) Delete(videoID string, language string) error {

	err := checkVideoID(videoID)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s/captions/%s", videosBasePath, videoID, language)

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
