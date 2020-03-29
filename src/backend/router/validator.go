package router

import "gopkg.in/go-playground/validator.v9"

type (
	UploadData struct {
		FileType  string `json:"fileType" validate:"required"`
		Data string `json:"data" validate:"required"`
		FileName string `json:"fileName" validate:"required"`
	}

	DeletePayload struct {
		Key  string `json:"key" validate:"required"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
