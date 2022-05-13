package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID         string    `json:"encoded_video_folder" valid:"uuid" gorm:"type:uuid;primary_key"` // internal identification and video folder location
	ResourceID string    `json:"resource_id" valid:"notnull" gorm:"type:varchar(255)"`           // id from client
	FilePath   string    `json:"file_path" valid:"notnull" gorm:"type:varchar(255)`              // bucket`s location
	CreatedAt  time.Time `json:"-" valid:"-"`                                                    // no validate
	Jobs       []*Job    `json:"-" valid:"-" gorm:"ForeignKey:VideoID"`
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
