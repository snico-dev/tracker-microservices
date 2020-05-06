package common

import "net/http"

// Route describe a HTTP route
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}
