package apivideosdk

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var accountJSONResponse string = `{
	"quota": {
	  "quotaUsed": 6,
	  "quotaRemaining": 54,
	  "quotaTotal": 60
	},
	"term": {
	  "startAt": "2019-02-20",
	  "endAt": "2019-03-20"
	}
  }`

var accountStruct = Account{
	Quota: &Quota{
		QuotaUsed:      6,
		QuotaRemaining: 54,
		QuotaTotal:     60,
	},
	Term: &Term{
		StartAt: "2019-02-20",
		EndAt:   "2019-03-20",
	},
}

func TestAccount_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, accountJSONResponse)
	})

	account, err := client.Account.Get()
	if err != nil {
		t.Errorf("Account.Get error: %v", err)
	}

	expected := &accountStruct
	if !reflect.DeepEqual(account, expected) {
		t.Errorf("Account.Get\n got=%#v\nwant=%#v", account, expected)
	}
}
