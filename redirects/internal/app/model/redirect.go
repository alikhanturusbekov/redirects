package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Redirect struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	ActiveLink  string    `gorm:"varchar(255);not null" json:"active_link,omitempty"`
	HistoryLink string    `gorm:"varchar(255);not null" json:"history_link,omitempty"`
	CreatedAt   time.Time `gorm:"not null" json:"createdAt,omitempty"`
	UpdatedAt   time.Time `gorm:"not null" json:"updatedAt,omitempty"`
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type CreateRedirect struct {
	ActiveLink  string `json:"title" validate:"required"`
	HistoryLink string `json:"content" validate:"required"`
}

type UpdateRedirect struct {
	ActiveLink  string `json:"title,omitempty"`
	HistoryLink string `json:"content,omitempty"`
}
