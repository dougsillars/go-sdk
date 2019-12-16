package apivideosdk

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

var statsJSONResponses = []string{
	`{
		"session": {
		  "sessionId": "psEmFwGQUAXR2lFHj5nDOpy",
		  "loadedAt": "2019-06-24T11:45:01.109+00",
		  "endedAt": "2019-06-24T11:49:19.243+00"
		},
		"location": {
		  "country": "France",
		  "city": "Paris"
		},
		"referrer": {
		  "url": "https://api.video",
		  "medium": "organic",
		  "source": "https://google.com",
		  "searchTerm": "video encoding hosting and delivery"
		},
		"device": {
		  "type": "desktop",
		  "vendor": "Dell",
		  "model": "unknown"
		},
		"os": {
		  "name": "Microsoft Windows",
		  "shortname": "W10",
		  "version": "Windows10"
		},
		"client": {
		  "type": "browser",
		  "name": "Firefox",
		  "version": "67.0"
		}
	  }`,
	`{
		"session": {
		  "sessionId": "psLQOSDqsdoqsjdoLQSD65o",
		  "loadedAt": "2019-06-25T11:45:01.109+00",
		  "endedAt": "2019-06-25T11:49:19.243+00",
		  "metadata": null
		},
		"location": {
		  "country": "France",
		  "city": "Paris"
		},
		"referrer": {
		  "url": "https://api.video",
		  "medium": "organic",
		  "source": "https://google.com",
		  "searchTerm": "video encoding hosting and delivery"
		},
		"device": {
		  "type": "desktop",
		  "vendor": "Dell",
		  "model": "unknown"
		},
		"os": {
		  "name": "Microsoft Windows",
		  "shortname": "W10",
		  "version": "Windows10"
		},
		"client": {
		  "type": "browser",
		  "name": "Firefox",
		  "version": "67.0"
		}
	  }`,
}

var statsStructs = []Statistic{
	Statistic{
		Session: &Session{
			SessionID: "psEmFwGQUAXR2lFHj5nDOpy",
			LoadedAt:  "2019-06-24T11:45:01.109+00",
			EndedAt:   "2019-06-24T11:49:19.243+00",
		},
		Location: &Location{
			Country: "France",
			City:    "Paris",
		},
		Referrer: &Referrer{
			URL:        "https://api.video",
			Medium:     "organic",
			Source:     "https://google.com",
			SearchTerm: "video encoding hosting and delivery",
		},
		Device: &Device{
			Type:   "desktop",
			Vendor: "Dell",
			Model:  "unknown",
		},
		Os: &Os{
			Name:      "Microsoft Windows",
			Shortname: "W10",
			Version:   "Windows10",
		},
		SessClient: &SessClient{
			Type:    "browser",
			Name:    "Firefox",
			Version: "67.0",
		},
	},
	Statistic{
		Session: &Session{
			SessionID: "psLQOSDqsdoqsjdoLQSD65o",
			LoadedAt:  "2019-06-25T11:45:01.109+00",
			EndedAt:   "2019-06-25T11:49:19.243+00",
		},
		Location: &Location{
			Country: "France",
			City:    "Paris",
		},
		Referrer: &Referrer{
			URL:        "https://api.video",
			Medium:     "organic",
			Source:     "https://google.com",
			SearchTerm: "video encoding hosting and delivery",
		},
		Device: &Device{
			Type:   "desktop",
			Vendor: "Dell",
			Model:  "unknown",
		},
		Os: &Os{
			Name:      "Microsoft Windows",
			Shortname: "W10",
			Version:   "Windows10",
		},
		SessClient: &SessClient{
			Type:    "browser",
			Name:    "Firefox",
			Version: "67.0",
		},
	},
}

var sessEventJSONResponses = []string{
	`{
		"type": "player_session_vod.loaded",
		"emittedAt": "2019-01-01 03:11:35.973+01",
		"at": 0,
		"from":10,
		"to":15
	  }`,
	`{
		"type": "player_session_vod.played",
		"emittedAt": "2019-01-01 03:11:36.232+01",
		"at": 0
	  }`,
	`{
		"type": "player_session_vod.paused",
		"emittedAt": "2019-01-01 03:11:38.837+01",
		"at": 2
	  }`,
}

var sessEventStructs = []SessionEvent{
	SessionEvent{
		Type:      "player_session_vod.loaded",
		EmittedAt: "2019-01-01 03:11:35.973+01",
		At:        0,
		From:      10,
		To:        15,
	},
	SessionEvent{
		Type:      "player_session_vod.played",
		EmittedAt: "2019-01-01 03:11:36.232+01",
		At:        0,
	},
	SessionEvent{
		Type:      "player_session_vod.paused",
		EmittedAt: "2019-01-01 03:11:38.837+01",
		At:        2,
	},
}

func TestStatistics_GetVideoSessions(t *testing.T) {
	setup()
	defer teardown()
	JSONResp := fmt.Sprintf(
		`{"data":[%s,%s], "pagination":%s}`,
		statsJSONResponses[0],
		statsJSONResponses[1],
		paginationJSON)

	mux.HandleFunc("/analytics/videos/vi4k0jvEUuaTdRAEjQ4Jfagz", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		expectedQuery := url.Values{
			"currentPage":    []string{"1"},
			"pageSize":       []string{"25"},
			"period":         []string{"2019-12-05"},
			"metadata[key]":  []string{"value"},
			"metadata[key2]": []string{"value2"},
		}
		if !reflect.DeepEqual(r.URL.Query(), expectedQuery) {
			t.Errorf("Request querystring\n got=%#v\nwant=%#v", r.URL.Query(), expectedQuery)
		}
		fmt.Fprint(w, JSONResp)
	})

	opts := &SessionVideoOpts{
		CurrentPage: 1,
		PageSize:    25,
		Period:      "2019-12-05",
		Metadata:    map[string]string{"key": "value", "key2": "value2"},
	}
	stats, err := client.Statistics.GetVideoSessions("vi4k0jvEUuaTdRAEjQ4Jfagz", opts)
	if err != nil {
		t.Errorf("Statistics.GetVideoSessions error: %v", err)
	}

	expected := &StatisticList{
		Data:       statsStructs,
		Pagination: &paginationStruct,
	}
	if !reflect.DeepEqual(stats, expected) {
		t.Errorf("Statistics.GetVideoSessions\n got=%#v\nwant=%#v", stats, expected)
	}
}

func TestStatistics_GetLivestreamSessions(t *testing.T) {
	setup()
	defer teardown()
	JSONResp := fmt.Sprintf(
		`{"data":[%s,%s], "pagination":%s}`,
		statsJSONResponses[0],
		statsJSONResponses[1],
		paginationJSON)

	mux.HandleFunc("/analytics/live-streams/li2FgWk8CyBKFIGyDK1SimnL", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		expectedQuery := url.Values{
			"currentPage": []string{"1"},
			"pageSize":    []string{"25"},
			"period":      []string{"2019-12-05"},
		}
		if !reflect.DeepEqual(r.URL.Query(), expectedQuery) {
			t.Errorf("Request querystring\n got=%#v\nwant=%#v", r.URL.Query(), expectedQuery)
		}
		fmt.Fprint(w, JSONResp)
	})

	opts := &SessionLivestreamOpts{
		CurrentPage: 1,
		PageSize:    25,
		Period:      "2019-12-05",
	}
	stats, err := client.Statistics.GetLivestreamSessions("li2FgWk8CyBKFIGyDK1SimnL", opts)
	if err != nil {
		t.Errorf("Statistics.GetLivestreamSession error: %v", err)
	}

	expected := &StatisticList{
		Data:       statsStructs,
		Pagination: &paginationStruct,
	}
	if !reflect.DeepEqual(stats, expected) {
		t.Errorf("Statistics.GetLivestreamSession\n got=%#v\nwant=%#v", stats, expected)
	}
}

func TestStatistics_GetSessionEvents(t *testing.T) {
	setup()
	defer teardown()
	JSONResp := fmt.Sprintf(
		`{"data":[%s,%s,%s], "pagination":%s}`,
		sessEventJSONResponses[0],
		sessEventJSONResponses[1],
		sessEventJSONResponses[2],
		paginationJSON)

	mux.HandleFunc("/analytics/sessions/psEmFwGQUAXR2lFHj5nDOpy/events", func(w http.ResponseWriter, r *http.Request) {
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

	opts := &SessionEventOpts{
		CurrentPage: 1,
		PageSize:    25,
	}
	stats, err := client.Statistics.GetSessionEvents("psEmFwGQUAXR2lFHj5nDOpy", opts)
	if err != nil {
		t.Errorf("Statistics.GetSessionEvents error: %v", err)
	}

	expected := &SessionEventList{
		Data:       sessEventStructs,
		Pagination: &paginationStruct,
	}
	if !reflect.DeepEqual(stats, expected) {
		t.Errorf("Statistics.GetSessionEvents\n got=%#v\nwant=%#v", stats, expected)
	}
}
