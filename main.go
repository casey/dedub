package app

import (
	. "flotilla"
	"net/http"
	"strings"
)

func init() {
	Handle("/").Default(handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	parts := strings.Split(url.Host, ".")
	Ensure(len(parts) > 1, StatusForbidden)
	url.Host = strings.Join(parts[1:], ".")
	Redirect(StatusMovedPermanently, url.String())
}
