package apivideosdk

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var tokenJSONResponse string = `{
	"token": "toXXX"
  }`

var tokenStruct = UploadToken{
	Token: "toXXX",
}

func TestUploadTokens_Generate(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/upload-tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, tokenJSONResponse)
	})

	token, err := client.UploadTokens.Generate()
	if err != nil {
		t.Errorf("UploadTokens.Generate error: %v", err)
	}

	expected := &tokenStruct
	if !reflect.DeepEqual(token, expected) {
		t.Errorf("UploadTokens.Generate\n got=%#v\nwant=%#v", token, expected)
	}
}
