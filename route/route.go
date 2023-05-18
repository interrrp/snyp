// Package route contains all routes for the server.
package route

import (
	"github.com/labstack/echo/v4"
)

// Routes contains all routes.
var Routes = []Route{CreateSnippet, GetSnippet}

// A Route is a single route on a server.
type Route struct {
	Path    string
	Method  Method
	Handler echo.HandlerFunc
}

// Register registers the route to a router.
//
// For Echo apps, you generally just need to pass in
// the Router()'s result which is a pointer to the app's router.
func (r *Route) Register(e *echo.Router) {
	e.Add(string(r.Method), r.Path, r.Handler)
}
