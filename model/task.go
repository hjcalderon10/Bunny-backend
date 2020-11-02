package model

import "time"

type (
	TaskID = uint16
	Task   struct {
		ID          TaskID    `json:"id,omitempty" sql:"id" validate:"required_without=UserID"`
		UserID      uint16    `json:"user_id,omitempty" sql:"user_id" validate:"required_without=ID"`
		State       string    `json:"state,omitempty" sql:"state"`
		Title       string    `json:"title,omitempty" sql:"title"`
		Description string    `json:"description,omitempty" sql:"description"`
		CreatedAt   time.Time `json:"created_at,omitempty" sql:"created_at"`
	}

	TaskState struct {
		ID    uint16 `json:"id,omitempty" sql:"id"`
		State string `json:"state,omitempty" sql:"state"`
	}
)
