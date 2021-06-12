package actions

import (
	"github.com/gobuffalo/buffalo"
	"net/http"
)

// FalcoEventsHandler is a default handler to serve up the falco events
func FalcoRulesHandler(c buffalo.Context) error {
	c.Set("page-title", "Falco rules")
	c.Set("page-description", "Falco Rules")
	return c.Render(http.StatusOK, r.HTML("falco_rules.html"))
}
