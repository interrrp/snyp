// Package main implements a Snyp server.
package main

import (
	"github.com/labstack/echo/v4"

	"github.com/interrrp/snyp/route"
	"github.com/interrrp/snyp/util"
)

func main() {
	// Clear the console so the output looks nicer
	// in development environments.
	util.ClearConsole()

	e := echo.New()

	for _, r := range route.Routes {
		r.Register(e.Router())
	}

	e.Start(util.EnvOr("ADDR", ":3000"))
}