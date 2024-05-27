package faas

import (
	"github.com/keirwhitlock/hello-api/handlers/rest"
	"net/http"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	rest.TranslateHandler(w, r)
}
