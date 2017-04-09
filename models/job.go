package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Job struct {
	ID        uuid.UUID   `json:"id" db:"id"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
	JobCategory    string `json:"job_category" db:"job_category"`
	JobType        string `json:"job_type" db:"job_type"`
	Location       string `json:"location" db:"location"`
	Title          string `json:"title" db:"title"`
	Description    string  `json:"description" db:"description"`
	Instructions   string `json:"instructions" db:"instructions"`
	CompanyName    string `json:"company_name" db:"company_name"`
	CompanyLogo    string `json:"company_logo" db:"company_logo"`
	CompanyEmail   string `json:"company_email" db:"company_email"`
	CompanyWebsite string `json:"company_website" db:"company_website"`
	CompanyTwitter string `json:"company_twitter" db:"company_twitter"`
}

// String is not required by pop and may be deleted
func (j Job) String() string {
	jj, _ := json.Marshal(j)
	return string(jj)
}

// Jobs is not required by pop and may be deleted
type Jobs []Job

// String is not required by pop and may be deleted
func (j Jobs) String() string {
	jj, _ := json.Marshal(j)
	return string(jj)
}

// Validate gets run everytime you call a "pop.Validate" method.
func (j *Job) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return j.validateCommon(tx)
}

// ValidateSave gets run everytime you call "pop.ValidateSave" method.
func (j *Job) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return j.validateCommon(tx)
}

// ValidateUpdate gets run everytime you call "pop.ValidateUpdate" method.
func (j *Job) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return j.validateCommon(tx)
}

// validateCommon provides shared validation rules used by other validation methods,
// such as "pop.Validate", "pop.ValidateSave", and "pop.ValidateUpdate".
func (j *Job) validateCommon(tx *pop.Connection) (*validate.Errors, error) {
	jobCategories := []string{"Copywriting", "Customer Service", "Design", "DevOps", "Management", "Marketing", "Programming"}
	jobTypes := []string{"Full-Time", "Part-Time", "Contract"}
	locations := []string{"Remote", "Fort Erie", "Grimsby", "Lincoln", "Niagara Falls", "Niagara-on-the-Lake", "Pelham", "Port Colborne", "St. Catharines", "Thorold", "Wainfleet", "Welland", "West Lincoln", "Out of Region"}

	validationErrors := validate.Validate(
		&validators.StringInclusion{Name: "Job Category", Field: j.JobCategory, List: jobCategories},
		&validators.StringInclusion{Name: "Job Type", Field: j.JobType, List: jobTypes},
		&validators.StringInclusion{Name: "Location", Field: j.Location, List: locations},
		&validators.StringIsPresent{Name: "Title", Field: j.Title},
		&validators.StringIsPresent{Name: "Description", Field: j.Description},
		&validators.StringIsPresent{Name: "Application instructions", Field: j.Instructions},
		&validators.StringIsPresent{Name: "Company name", Field: j.CompanyName},
		&validators.StringIsPresent{Name: "Email", Field: j.CompanyEmail},
		&validators.StringIsPresent{Name: "Website", Field: j.CompanyWebsite},
		&validators.StringIsPresent{Name: "Twitter handle", Field: j.CompanyTwitter},
	)
	return validationErrors, nil
}
