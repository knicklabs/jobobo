package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/knicklabs/jobobo/models"
	"github.com/gobuffalo/buffalo/render"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// Middleware to find a job by ID
func FindJobMW(h buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		id := c.Param("job_id")

		if id == "" {
			return h(c)
		}

		j := &models.Job{}
		tx := c.Get("tx").(*pop.Connection)
		err := tx.Find(j, id)
		if err != nil {
			return c.Error(404, errors.WithStack(err))
		}
		c.Set("job", j)
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
	return c.Render(200, render.JSON(jobs))
}

// JobsShow renders an html page for viewing a job
func JobsShow(c buffalo.Context) error {
	return c.Render(200, render.JSON(c.Get("job")))
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
		return c.Render(422, render.JSON(validationErrors))
	}
	err = tx.Create(j)
	if err != nil {
		return errors.WithStack(err)
	}
	return c.Render(201, render.JSON(j))
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
		return c.Render(422, render.JSON(validationErrors))
	}
	err = tx.Update(j)
	if err != nil {
		return errors.WithStack(err)
	}

	err = tx.Reload(j)
	if err != nil {
		return errors.WithStack(err)
	}
	return c.Render(200, render.JSON(j))
}

// JobsDelete removes a job
func JobsDelete(c buffalo.Context) error {
	tx := c.Get("tx").(*pop.Connection)
	j := c.Get("job").(*models.Job)

	err := tx.Destroy(j)
	if err != nil {
		return errors.WithStack(err)
	}
	return c.Render(200, render.JSON(j))
}
