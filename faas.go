// Package faas : function as a service handler
package faas

import (
	"github.com/keirwhitlock/hello-api/handlers/rest"
	"github.com/keirwhitlock/hello-api/translation"
	"net/http"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	service := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(service)
	translateHandler.TranslateHandler(w, r)
}
