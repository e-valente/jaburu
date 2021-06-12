package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	c.Set("page-title", "Jaburu Home")
	c.Set("page-description", "Jaburu Home")
	return c.Render(http.StatusOK, r.HTML("index.html"))
}
