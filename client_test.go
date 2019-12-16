package apivideosdk

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

var (
	mux *http.ServeMux

	client *Client

	server *httptest.Server
)

var paginationJSON = `{
    "currentPage": 1,
    "pageSize": 25,
    "pagesTotal": 1,
    "itemsTotal": 11,
    "currentPageItems": 11,
    "links": [
      {
        "rel": "self",
        "uri": "https://ws.api.video"
      },
      {
        "rel": "first",
        "uri": "https://ws.api.video"
      },
      {
        "rel": "last",
        "uri": "https://ws.api.video"
      }
    ]
  }`

var paginationStruct = Pagination{
	CurrentPage:      1,
	PageSize:         25,
	PagesTotal:       1,
	ItemsTotal:       11,
	CurrentPageItems: 11,
	Links: []Link{
		{
			Rel: "self",
			URI: "https://ws.api.video",
		},
		{
			Rel: "first",
			URI: "https://ws.api.video",
		},
		{
			Rel: "last",
			URI: "https://ws.api.video",
		},
	},
}

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	mux.HandleFunc("/auth/api-key", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
			"token_type": "Bearer",
			"expires_in": 3600,
			"access_token": "fakeToken",
			"refresh_token": "fakeToken"
		  }`)
	})

	client = NewClient("apiKey")
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}

func createTempFile(filename string, size int64) string {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := f.Truncate(size); err != nil {
		log.Fatal(err)
	}
	return filename
}
