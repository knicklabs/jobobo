package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/knicklabs/jobobo/models"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// Middleware to find a job by ID
func findJobMW(h buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		id, err := c.ParamInt("job_id")
		if err == nil {
			j := &models.Job{}
			tx := c.Get("tx").(*pop.Connection)
			err = tx.Find(j, id)
			if err != nil {
				return c.Error(404, errors.WithStack(err))
			}
			c.Set("job", j)
		}
		return h(c)
	}
}

// JobsList renders an html page that shows jobs
func JobsList(c buffalo.Context) error {
	jobs := &models.Jobs{}
	tx := c.Get("tx").(*pop.Connection)
	err := tx.All(jobs)
	if err != nil {
		return c.Error(404, errors.WithStack(err))
	}
	c.Set("jobs", jobs)
	return c.Render(200, r.HTML("jobs/index.html"))
}

// JobsRedirect redirects to the jobs list on the homepage
func JobsRedirect(c buffalo.Context) error {
	return c.Redirect(301, "/")
}

// JobsShow renders an html page for viewing a job
func JobsShow(c buffalo.Context) error {
	return c.Render(200, r.HTML("jobs/show.html"))
}

// JobsNew renders a form for adding a new job
func JobsNew(c buffalo.Context) error {
	c.Set("job", models.Job{})
	return c.Render(200, r.HTML("jobs/new.html"))
}

// JobsCreate creates a job
func JobsCreate(c buffalo.Context) error {
	j := &models.Job{}
	err := c.Bind(j)
	if err != nil {
		return errors.WithStack(err)
	}
	tx := c.Get("tx").(*pop.Connection)
	validationErrors, err := j.ValidateSave(tx)
	if err != nil {
		return errors.WithStack(err)
	}
	if validationErrors.HasAny() {
		c.Set("validationErrors", validationErrors.Errors)
		c.Set("job", j)
		return c.Render(422, r.HTML("jobs/new.html"))
	}
	err = tx.Create(j)
	if err != nil {
		return errors.WithStack(err)
	}
	return c.Redirect(301, "/jobs/%d", j.ID)
}

// JobsEdit renders a form for editing a job
func JobsEdit(c buffalo.Context) error {
	return c.Render(200, r.HTML("jobs/edit.html"))
}

// JobsUpdate updates a job
func JobsUpdate(c buffalo.Context) error {
	tx := c.Get("tx").(*pop.Connection)
	j := c.Get("job").(*models.Job)

	err := c.Bind(j)
	if err != nil {
		return errors.WithStack(err)
	}

	validationErrors, err := j.ValidateUpdate(tx)
	if err != nil {
		return errors.WithStack(err)
	}
	if validationErrors.HasAny() {
		c.Set("validationErrors", validationErrors.Errors)
		c.Set("job", j)
		return c.Render(422, r.HTML("jobs/edit.html"))
	}
	err = tx.Update(j)
	if err != nil {
		return errors.WithStack(err)
	}

	err = tx.Reload(j)
	if err != nil {
		return errors.WithStack(err)
	}
	return c.Redirect(301, "/users/%d", j.ID)
}

// JobsDelete removes a job
func JobsDelete(c buffalo.Context) error {
	tx := c.Get("tx").(*pop.Connection)
	j := c.Get("job").(*models.Job)

	err := tx.Destroy(j)
	if err != nil {
		return errors.WithStack(err)
	}
	return c.Redirect(301, "/")
}
