package actions

import (
	"github.com/gobuffalo/buffalo"
)

// Renders an HTML page
func Home(c buffalo.Context) error {
	return c.Render(200, r.HTML("home/index.html"))
}
