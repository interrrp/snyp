// Contains the root route.

package route

import (
	"github.com/labstack/echo/v4"

	"github.com/interrrp/snyp/util"
)

// Root is the root route.
var Root = Route{
	Path:    "/",
	Method:  MethodGET,
	Handler: handleRoot,
}

// handleRoot handles the root route.
func handleRoot(c echo.Context) error {
	c.JSON(200, util.H{
		"life":  42,
		"hotel": "Trivago",
	})

	return nil
}
