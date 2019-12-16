package apivideosdk

import (
	"net/http"
)

const uploadTokensBasePath = "upload-tokens"

// UploadTokensServiceI is an interface representing the Upload Tokens
// endpoints of the api.video API
// See: https://docs.api.video/5.1/videos-delegated-upload
type UploadTokensServiceI interface {
	Generate() (*UploadToken, error)
}

// UploadTokensService communicating with the Upload Tokens
// endpoints of the api.video API
type UploadTokensService struct {
	client *Client
}

// UploadToken represents an api.video UploadToken
type UploadToken struct {
	Token string `json:"token,omitempty"`
}

//Generate returns a new generated UploadToken
func (s *UploadTokensService) Generate() (*UploadToken, error) {

	req, err := s.client.prepareRequest(http.MethodPost, uploadTokensBasePath, nil)
	if err != nil {
		return nil, err
	}

	ut := new(UploadToken)
	_, err = s.client.do(req, ut)

	if err != nil {
		return nil, err
	}

	return ut, nil
}
