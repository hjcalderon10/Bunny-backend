package model

import "time"

type (
	UserID = uint16
	User   struct {
		ID            UserID    `json:"id" sql:"id"`
		Name          string    `json:"name,omitempty" sql:"name,omitempty"`
		NickName      string    `json:"nickname,omitempty" sql:"nickname,omitempty"`
		ImgURL        string    `json:"img_url,omitempty" sql:"img_url,omitempty"`
		IsActive      *bool     `json:"-" sql:"is_active,omitempty"`
		CreatedAt     time.Time `json:"created_at,omitempty" sql:"created_at,omitempty"`
		InactivatedAt time.Time `json:"inactivated_at,omitempty" sql:"inactivated_at,omitempty"`
	}
)
