package apivideosdk

import (
	"net/http"
)

const accountBasePath = "account"

// AccountServiceI is an interface representing the Account
// endpoints of the api.video API
// See: https://docs.api.video/5.1/captions
type AccountServiceI interface {
	Get() (*Account, error)
}

// AccountService communicating with the Account
// endpoints of the api.video API
type AccountService struct {
	client *Client
}

// Account represents an api.video Account
type Account struct {
	Quota *Quota `json:"quota,omitempty"`
	Term  *Term  `json:"term,omitempty"`
}

// Quota represents a Quota
type Quota struct {
	QuotaUsed      int `json:"quotaUsed,omitempty"`
	QuotaRemaining int `json:"quotaRemaining,omitempty"`
	QuotaTotal     int `json:"quotaTotal,omitempty"`
}

// Term represents a Term
type Term struct {
	StartAt string `json:"startAt,omitempty"`
	EndAt   string `json:"endAt,omitempty"`
}

//Get returns an Account
func (s *AccountService) Get() (*Account, error) {

	req, err := s.client.prepareRequest(http.MethodGet, accountBasePath, nil)
	if err != nil {
		return nil, err
	}

	a := new(Account)
	_, err = s.client.do(req, a)

	if err != nil {
		return nil, err
	}

	return a, nil
}
