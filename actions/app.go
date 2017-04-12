package actions

import (
	"github.com/gobuffalo/buffalo"

	"github.com/gobuffalo/buffalo/middleware"
	"github.com/knicklabs/jobobo/models"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env:         ENV,
			SessionName: "_jobobo_session",
		})

		app.Use(middleware.PopTransaction(models.DB))
		app.GET("/", Home)

		jg := app.Group("/api/jobs")
		jg.Use(middleware.SetContentType("application/json"))
		jg.Use(FindJobMW)
		jg.GET("/", JobsList)
		jg.GET("/{job_id}", JobsShow)
		jg.POST("/", JobsCreate)
		jg.PUT("/{job_id}", JobsUpdate)
		jg.DELETE("/{job_id}", JobsDelete)

		app.ServeFiles("/assets", packr.NewBox("../public/assets"))
	}

	return app
}
