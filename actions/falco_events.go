package actions

import (
	"github.com/gobuffalo/buffalo"
	"net/http"
)

// FalcoEventsHandler is a default handler to serve up the falco events
func FalcoEventsHandler(c buffalo.Context) error {
	c.Set("page-title", "Falco Events")
	c.Set("page-description", "Falco Events")
	return c.Render(http.StatusOK, r.HTML("falco_events.html"))
}
