package models

// Idea is a struct that represents an idea entity.
type Idea struct {
	// Id is the primary key of the idea, and it is marked with the "primary_key" tag.
	Id int `json:"id" gorm:"primary_key"`

	// IdeaName represents the name/title of the idea.
	IdeaName string `json:"idea"`

	// IdeaOwner represents the owner/creator of the idea.
	IdeaOwner string `json:"owner"`
}
