package model

import "time"

type (
	UserID = uint16
	User   struct {
		ID        UserID    `json:"id" sql:"id"`
		Name      string    `json:"name,omitempty" sql:"name"`
		NickName  string    `json:"nickname,omitempty" sql:"nickname"`
		ImgURL    string    `json:"img_url,omitempty" sql:"img_url"`
		CreatedAt time.Time `json:"created_at,omitempty" sql:"created_at"`
	}
)
