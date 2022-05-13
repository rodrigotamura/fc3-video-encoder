package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID         string    `valid:"uuid"`    // internal identification
	ResourceID string    `valid:"notnull"` // id from client
	FilePath   string    `valid:"notnull"` // bucket`s location
	CreatedAt  time.Time `valid:"-"`       // no validate
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true) // every filed must be required by default
}

func NewVideo() *Video {
	return &Video{}
}

func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video) // validating struct

	if err != nil {
		return err
	}

	return nil
}
