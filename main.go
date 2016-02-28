package app

import (
	"strings"
	"net/http"
	. "flotilla"
)

func init() {
  Handle("/").Default(handle)
}

func handle(r *http.Request) {
	url := r.URL
	parts := strings.Split(url.Host, ".")
	Ensure(len(parts) > 1, StatusForbidden)
	url.Host = strings.Join(parts[1:], ".")
  Redirect(StatusTemporaryRedirect, url.String())
}

