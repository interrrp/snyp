// Provides the route for getting snippets.

package route

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/interrrp/snyp/db"
	"github.com/interrrp/snyp/util"
)

// GetSnippet is the route for getting snippets.
var GetSnippet = Route{
	Path:    "/api/snippets/:id",
	Method:  MethodGET,
	Handler: handleGetSnippet,
}

// handleGetSnippet handles the route for getting snippets.
func handleGetSnippet(c echo.Context) error {
	id := c.Param("id")

	var s db.Snippet
	db.DB.First(&s, "id = ?", id)

	if s == (db.Snippet{}) {
		return echo.NewHTTPError(
			http.StatusNotFound,
			"item not found",
		)
	}

	c.JSON(http.StatusOK, util.H{
		"id":        s.ID,
		"name":      s.Name,
		"content":   s.Content,
		"createdAt": s.CreatedAt,
		"updatedAt": s.UpdatedAt,
	})

	return nil
}
