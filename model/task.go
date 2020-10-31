package model

import "time"

type (
	Task struct {
		ID            uint16    `json:"id,omitempty" sql:"id,omitempty" validate:"required_without=UserID"`
		UserID        uint16    `json:"user_id,omitempty" sql:"user_id,omitempty" validate:"required_without=ID"`
		State         string    `json:"state,omitempty" sql:"state,omitempty"`
		Title         string    `json:"title,omitempty" sql:"title,omitempty"`
		Description   string    `json:"description,omitempty" sql:"description,omitempty"`
		IsActive      *bool     `json:"-" sql:"is_active,omitempty"`
		CreatedAt     time.Time `json:"created_at,omitempty" sql:"created_at,omitempty"`
		InactivatedAt time.Time `json:"inactivated_at,omitempty" sql:"inactivated_at,omitempty"`
	}
)
