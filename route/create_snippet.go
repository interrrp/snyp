// Provides the route for creating snippets.

package route

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/interrrp/snyp/db"
	"github.com/interrrp/snyp/util"
)

// CreateSnippet is the route for creating snippets.
var CreateSnippet = Route{
	Path:    "/api/snippets",
	Method:  MethodPOST,
	Handler: handleCreateSnippet,
}

// createSnippetParams are the JSON body parameters
// to the snippet creation route.
type createSnippetParams struct {
	Name    string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// handleCreateSnippet handles the snippet creation route.
func handleCreateSnippet(c echo.Context) error {
	var params createSnippetParams
	if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	if err := util.Validate(params); err != "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err,
		)
	}

	s := &db.Snippet{
		Name:    params.Name,
		Content: params.Content,
	}

	tx := db.DB.Create(s)
	if tx.Error != nil {
		return tx.Error
	}

	c.JSON(http.StatusCreated, util.H{
		"id":        s.ID,
		"name":      s.Name,
		"content":   s.Content,
		"createdAt": s.CreatedAt,
	})

	return nil
}
