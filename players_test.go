package apivideosdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"
)

var playerJSONResponses = []string{
	`{
		"playerId": "pt3Lony8J6NozV71Yxn8KVFn",
		"assets": {
			"logo": "https://cdn.api.video/player/pl3Lony8J6NozV71Yxn8KVFn/logo.png",
			"link": "https://api.video"
		},
		"shapeMargin": 3,
		"shapeRadius": 10,
		"shapeAspect": "flat",
		"shapeBackgroundTop": "rgba(50, 50, 50, 0.7)",
		"shapeBackgroundBottom": "rgba(50, 50, 50, 0.8)",
		"text": "rgba(255, 255, 255, 0.95)",
		"link": "rgba(255, 0, 0, 0.95)",
		"linkHover": "rgba(255, 255, 255, 0.75)",
		"linkActive": "rgba(255, 0, 0, 0.75)",
		"trackPlayed": "rgba(255, 255, 255, 0.95)",
		"trackUnplayed": "rgba(255, 255, 255, 0.1)",
		"trackBackground": "rgba(0, 0, 0, 0)",
		"backgroundTop": "rgba(72, 4, 45, 1)",
		"backgroundBottom": "rgba(94, 95, 89, 1)",
		"backgroundText": "rgba(255, 255, 255, 0.95)",
		"language": "en",
		"enableApi": false,
		"enableControls": true,
		"forceAutoplay": false,
		"hideTitle": false,
		"forceLoop": false
	}`,
	`{
		"playerId": "pl3zY7qtojdW2EvMIU37707q",
		"assets": {
			"logo": "https://cdn.api.video/player/pl3zY7qtojdW2EvMIU37707q/logo.png",
			"link": "https://api.video"
		},
		"shapeMargin": 3,
		"shapeRadius": 10,
		"shapeAspect": "flat",
		"shapeBackgroundTop": "rgba(50, 50, 50, 0.7)",
		"shapeBackgroundBottom": "rgba(50, 50, 50, 0.8)",
		"text": "rgba(255, 255, 255, 0.95)",
		"link": "rgba(255, 0, 0, 0.95)",
		"linkHover": "rgba(255, 255, 255, 0.75)",
		"linkActive": "rgba(255, 0, 0, 0.75)",
		"trackPlayed": "rgba(255, 255, 255, 0.95)",
		"trackUnplayed": "rgba(255, 255, 255, 0.1)",
		"trackBackground": "rgba(0, 0, 0, 0)",
		"backgroundTop": "rgba(72, 4, 45, 1)",
		"backgroundBottom": "rgba(94, 95, 89, 1)",
		"backgroundText": "rgba(255, 255, 255, 0.95)",
		"language": "en",
		"enableApi": false,
		"enableControls": true,
		"forceAutoplay": false,
		"hideTitle": false,
		"forceLoop": false
	}`,
}

var playerStructs = []Player{
	Player{
		PlayerID:              "pt3Lony8J6NozV71Yxn8KVFn",
		ShapeMargin:           3,
		ShapeRadius:           10,
		ShapeAspect:           "flat",
		ShapeBackgroundTop:    "rgba(50, 50, 50, 0.7)",
		ShapeBackgroundBottom: "rgba(50, 50, 50, 0.8)",
		Text:                  "rgba(255, 255, 255, 0.95)",
		Link:                  "rgba(255, 0, 0, 0.95)",
		LinkHover:             "rgba(255, 255, 255, 0.75)",
		LinkActive:            "rgba(255, 0, 0, 0.75)",
		TrackPlayed:           "rgba(255, 255, 255, 0.95)",
		TrackUnplayed:         "rgba(255, 255, 255, 0.1)",
		TrackBackground:       "rgba(0, 0, 0, 0)",
		BackgroundTop:         "rgba(72, 4, 45, 1)",
		BackgroundBottom:      "rgba(94, 95, 89, 1)",
		BackgroundText:        "rgba(255, 255, 255, 0.95)",
		Language:              "en",
		EnableAPI:             false,
		EnableControls:        true,
		ForceAutoplay:         false,
		HideTitle:             false,
		ForceLoop:             false,
		Assets: &PlayerAssets{
			Logo: "https://cdn.api.video/player/pl3Lony8J6NozV71Yxn8KVFn/logo.png",
			Link: "https://api.video",
		},
	},
	Player{
		PlayerID:              "pl3zY7qtojdW2EvMIU37707q",
		ShapeMargin:           3,
		ShapeRadius:           10,
		ShapeAspect:           "flat",
		ShapeBackgroundTop:    "rgba(50, 50, 50, 0.7)",
		ShapeBackgroundBottom: "rgba(50, 50, 50, 0.8)",
		Text:                  "rgba(255, 255, 255, 0.95)",
		Link:                  "rgba(255, 0, 0, 0.95)",
		LinkHover:             "rgba(255, 255, 255, 0.75)",
		LinkActive:            "rgba(255, 0, 0, 0.75)",
		TrackPlayed:           "rgba(255, 255, 255, 0.95)",
		TrackUnplayed:         "rgba(255, 255, 255, 0.1)",
		TrackBackground:       "rgba(0, 0, 0, 0)",
		BackgroundTop:         "rgba(72, 4, 45, 1)",
		BackgroundBottom:      "rgba(94, 95, 89, 1)",
		BackgroundText:        "rgba(255, 255, 255, 0.95)",
		Language:              "en",
		EnableAPI:             false,
		EnableControls:        true,
		ForceAutoplay:         false,
		HideTitle:             false,
		ForceLoop:             false,
		Assets: &PlayerAssets{
			Logo: "https://cdn.api.video/player/pl3zY7qtojdW2EvMIU37707q/logo.png",
			Link: "https://api.video",
		},
	},
}

var playerRequestJSON = `{
	"shapeMargin": 3,
	"shapeRadius": 10,
	"shapeAspect": "flat",
	"shapeBackgroundTop": "rgba(50, 50, 50, 0.7)",
	"shapeBackgroundBottom": "rgba(50, 50, 50, 0.8)",
	"text": "rgba(255, 255, 255, 0.95)",
	"link": "rgba(255, 0, 0, 0.95)",
	"linkHover": "rgba(255, 255, 255, 0.75)",
	"linkActive": "rgba(255, 0, 0, 0.75)",
	"trackPlayed": "rgba(255, 255, 255, 0.95)",
	"trackUnplayed": "rgba(255, 255, 255, 0.1)",
	"trackBackground": "rgba(0, 0, 0, 0)",
	"backgroundTop": "rgba(72, 4, 45, 1)",
	"backgroundBottom": "rgba(94, 95, 89, 1)",
	"backgroundText": "rgba(255, 255, 255, 0.95)",
	"language": "en",
	"enableApi": false,
	"enableControls": true,
	"forceAutoplay": false,
	"hideTitle": false,
	"forceLoop": false
  }`

var playerRequestStruct = PlayerRequest{
	ShapeMargin:           3,
	ShapeRadius:           10,
	ShapeAspect:           "flat",
	ShapeBackgroundTop:    "rgba(50, 50, 50, 0.7)",
	ShapeBackgroundBottom: "rgba(50, 50, 50, 0.8)",
	Text:                  "rgba(255, 255, 255, 0.95)",
	Link:                  "rgba(255, 0, 0, 0.95)",
	LinkHover:             "rgba(255, 255, 255, 0.75)",
	LinkActive:            "rgba(255, 0, 0, 0.75)",
	TrackPlayed:           "rgba(255, 255, 255, 0.95)",
	TrackUnplayed:         "rgba(255, 255, 255, 0.1)",
	TrackBackground:       "rgba(0, 0, 0, 0)",
	BackgroundTop:         "rgba(72, 4, 45, 1)",
	BackgroundBottom:      "rgba(94, 95, 89, 1)",
	BackgroundText:        "rgba(255, 255, 255, 0.95)",
	Language:              "en",
	EnableAPI:             false,
	EnableControls:        true,
	ForceAutoplay:         false,
	HideTitle:             false,
	ForceLoop:             false,
}

func TestPlayers_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/pt3Lony8J6NozV71Yxn8KVFn", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, playerJSONResponses[0])
	})

	player, err := client.Players.Get("pt3Lony8J6NozV71Yxn8KVFn")
	if err != nil {
		t.Errorf("Players.Get error: %v", err)
	}

	expected := &playerStructs[0]
	if !reflect.DeepEqual(player, expected) {
		t.Errorf("Players.Get\n got=%#v\nwant=%#v", player, expected)
	}
}

func TestPlayers_List(t *testing.T) {
	setup()
	defer teardown()
	JSONResp := fmt.Sprintf(
		`{"data":[%s,%s], "pagination":%s}`,
		playerJSONResponses[0],
		playerJSONResponses[1],
		paginationJSON)

	mux.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		expectedQuery := url.Values{
			"currentPage": []string{"1"},
			"pageSize":    []string{"25"},
		}
		if !reflect.DeepEqual(r.URL.Query(), expectedQuery) {
			t.Errorf("Request querystring\n got=%#v\nwant=%#v", r.URL.Query(), expectedQuery)
		}
		fmt.Fprint(w, JSONResp)
	})

	opts := &PlayerOpts{
		CurrentPage: 1,
		PageSize:    25,
	}
	players, err := client.Players.List(opts)
	if err != nil {
		t.Errorf("Players.List error: %v", err)
	}

	expected := &PlayerList{
		Data:       playerStructs,
		Pagination: &paginationStruct,
	}
	if !reflect.DeepEqual(players, expected) {
		t.Errorf("Players.List\n got=%#v\nwant=%#v", players, expected)
	}
}

func TestPlayers_Create(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		expectedBody := map[string]interface{}{
			"shapeMargin":           float64(3),
			"shapeRadius":           float64(10),
			"shapeAspect":           "flat",
			"shapeBackgroundTop":    "rgba(50, 50, 50, 0.7)",
			"shapeBackgroundBottom": "rgba(50, 50, 50, 0.8)",
			"text":                  "rgba(255, 255, 255, 0.95)",
			"link":                  "rgba(255, 0, 0, 0.95)",
			"linkHover":             "rgba(255, 255, 255, 0.75)",
			"linkActive":            "rgba(255, 0, 0, 0.75)",
			"trackPlayed":           "rgba(255, 255, 255, 0.95)",
			"trackUnplayed":         "rgba(255, 255, 255, 0.1)",
			"trackBackground":       "rgba(0, 0, 0, 0)",
			"backgroundTop":         "rgba(72, 4, 45, 1)",
			"backgroundBottom":      "rgba(94, 95, 89, 1)",
			"backgroundText":        "rgba(255, 255, 255, 0.95)",
			"language":              "en",
			"enableApi":             false,
			"enableControls":        true,
			"forceAutoplay":         false,
			"hideTitle":             false,
			"forceLoop":             false,
		}
		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}
		if !reflect.DeepEqual(v, expectedBody) {
			t.Errorf("Request body\n got=%#v\n want=%#v", v, expectedBody)
		}
		fmt.Fprint(w, playerJSONResponses[0])
	})

	player, err := client.Players.Create(&playerRequestStruct)
	if err != nil {
		t.Errorf("Players.Create error: %v", err)
	}

	expected := &playerStructs[0]
	if !reflect.DeepEqual(player, expected) {
		t.Errorf("Players.Create\n got=%#v\nwant=%#v", player, expected)
	}
}

func TestPlayers_Update(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/pt3Lony8J6NozV71Yxn8KVFn", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		expectedBody := map[string]interface{}{
			"shapeMargin":           float64(3),
			"shapeRadius":           float64(10),
			"shapeAspect":           "flat",
			"shapeBackgroundTop":    "rgba(50, 50, 50, 0.7)",
			"shapeBackgroundBottom": "rgba(50, 50, 50, 0.8)",
			"text":                  "rgba(255, 255, 255, 0.95)",
			"link":                  "rgba(255, 0, 0, 0.95)",
			"linkHover":             "rgba(255, 255, 255, 0.75)",
			"linkActive":            "rgba(255, 0, 0, 0.75)",
			"trackPlayed":           "rgba(255, 255, 255, 0.95)",
			"trackUnplayed":         "rgba(255, 255, 255, 0.1)",
			"trackBackground":       "rgba(0, 0, 0, 0)",
			"backgroundTop":         "rgba(72, 4, 45, 1)",
			"backgroundBottom":      "rgba(94, 95, 89, 1)",
			"backgroundText":        "rgba(255, 255, 255, 0.95)",
			"language":              "en",
			"enableApi":             false,
			"enableControls":        true,
			"forceAutoplay":         false,
			"hideTitle":             false,
			"forceLoop":             false,
		}
		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}
		if !reflect.DeepEqual(v, expectedBody) {
			t.Errorf("Request body\n got=%#v\n want=%#v", v, expectedBody)
		}
		fmt.Fprint(w, playerJSONResponses[0])
	})

	player, err := client.Players.Update("pt3Lony8J6NozV71Yxn8KVFn", &playerRequestStruct)
	if err != nil {
		t.Errorf("Players.Update error: %v", err)
	}

	expected := &playerStructs[0]
	if !reflect.DeepEqual(player, expected) {
		t.Errorf("Players.Update\n got=%#v\nwant=%#v", player, expected)
	}
}

func TestPlayers_Delete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/pt3Lony8J6NozV71Yxn8KVFn", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	err := client.Players.Delete("pt3Lony8J6NozV71Yxn8KVFn")
	if err != nil {
		t.Errorf("Players.Delete error: %v", err)
	}
}

func TestPlayers_UploadLogo(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/players/pt3Lony8J6NozV71Yxn8KVFn/logo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, playerJSONResponses[0])
	})

	file := createTempFile("test.logo", 1024*1024)
	defer os.Remove(file)

	player, err := client.Players.UploadLogo("pt3Lony8J6NozV71Yxn8KVFn", "https://api.video", file)
	if err != nil {
		t.Errorf("Captions.Upload error: %v", err)
	}

	expected := &playerStructs[0]
	if !reflect.DeepEqual(player, expected) {
		t.Errorf("Captions.Upload\n got=%#v\nwant=%#v", player, expected)
	}
}
