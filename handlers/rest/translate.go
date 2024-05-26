package rest

import (
	"encoding/json"
	"github.com/keirwhitlock/hello-api/translation"
	"net/http"
	"strings"
)

const defaultLanguage = "english"

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = defaultLanguage
	}

	word := strings.ReplaceAll(r.URL.Path, "/", "")
	translated_word := translation.Translate(word, language)

	if translated_word == "" {
		language = ""
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp := Resp{
		Language:    language,
		Translation: translated_word,
	}
	err := enc.Encode(resp)
	if err != nil {
		panic("unable to encode response")
	}
}
