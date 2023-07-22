package models

import "gorm.io/gorm"

type Idea struct {
	gorm.Model
	Id        int    `json:"id" gorm:"primary_key"`
	IdeaName  string `json:"idea"`
	IdeaOwner string `json:"owner"`
}
