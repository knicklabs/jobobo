package actions

import (
	"html/template"
	"strings"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),

		// Add template helpers here:
		Helpers: render.Helpers{
			"errClass": func(hasError bool) string {
				if hasError {
					return "has-error"
				}
				return ""
			},
			"errMsg": func(errors map[string][]string, key string) template.HTML {
				if value, ok := errors[key]; ok {
					element  := strings.Join(value, "<br>")
					elements := []string{"<span class='help-block'>", element, "</span>"}
					html := strings.Join(elements, "")
					return template.HTML(html)
				}

				return template.HTML("")
			},
			"helpMsg": func(message string) template.HTML {
				elements := []string{"<span class='help-block'>", message, "</span>"}
				html := strings.Join(elements, "")
				return template.HTML(html)
			},
		},
	})
}