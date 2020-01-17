package apivideosdk

import (
	"fmt"
	"net/http"
)

// ChaptersServiceI is an interface representing the Chapters
// endpoints of the api.video API
// See: https://docs.api.video/5.1/chapters
type ChaptersServiceI interface {
	Get(videoID string, language string) (*Chapter, error)
	List(videoID string) (*ChapterList, error)
	Upload(videoID string, language string, filepath string) (*Chapter, error)
	Delete(videoID string, language string) error
}

// ChaptersService communicating with the Chapters
// endpoints of the api.video API
type ChaptersService struct {
	client *Client
}

// Chapter represents an api.video Chapter
type Chapter struct {
	URI      string `json:"uri,omitempty"`
	Src      string `json:"src,omitempty"`
	Language string `json:"language,omitempty"`
}

// ChapterList represents a list of chapters
type ChapterList struct {
	Data       []Chapter   `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

//Get returns a Chapter by video id and language
func (s *ChaptersService) Get(videoID string, language string) (*Chapter, error) {

	err := checkVideoID(videoID)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s/chapters/%s", videosBasePath, videoID, language)

	req, err := s.client.prepareRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	c := new(Chapter)
	_, err = s.client.do(req, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

//List returns a slice of Chapter containing all chapters for a videoId
func (s *ChaptersService) List(videoID string) (*ChapterList, error) {

	err := checkVideoID(videoID)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s/chapters", videosBasePath, videoID)

	req, err := s.client.prepareRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	c := new(ChapterList)
	_, err = s.client.do(req, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

//Upload a vtt for a video and language
func (s *ChaptersService) Upload(videoID string, language string, filePath string) (*Chapter, error) {

	err := checkVideoID(videoID)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s/chapters/%s", videosBasePath, videoID, language)

	req, err := s.client.prepareUploadRequest(path, filePath, nil)

	if err != nil {
		return nil, err
	}

	c := new(Chapter)

	_, err = s.client.do(req, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

//Delete a chapter
func (s *ChaptersService) Delete(videoID string, language string) error {

	err := checkVideoID(videoID)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s/chapters/%s", videosBasePath, videoID, language)

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
