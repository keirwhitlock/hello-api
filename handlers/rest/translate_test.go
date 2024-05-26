package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTranslateAPI(t *testing.T) {

	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/hello",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Endpoint:            "/hello?language=german",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
		{
			Endpoint:            "/hello?language=dutch",
			StatusCode:          http.StatusNotFound,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
	}

	handler := http.HandlerFunc(TranslateHandler)

	for _, test := range tt {
		// Arrage

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.Endpoint, nil)
		handler.ServeHTTP(rr, req)

		// Assert
		if rr.Code != test.StatusCode {
			t.Errorf("exected status %d but received %d", test.StatusCode, rr.Code)
		}

		var resp Resp
		json.Unmarshal(rr.Body.Bytes(), &resp)

		if resp.Language != test.ExpectedLanguage {
			t.Errorf(`expected language to be "%s" but received "%s"`, test.ExpectedLanguage, resp.Language)
		}

		if resp.Translation != test.ExpectedTranslation {
			t.Errorf(`expected translation to be "%s" but received "%s"`, test.ExpectedTranslation, resp.Translation)
		}
	}
}
