package platform

import "errors"

type Platforms struct {
	Id     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title" binding:"required"`
	Text   string `json:"text" db:"text" binding:"required"`
	Author string `json:"author"`
}

type Authors struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

type AuthorsList struct {
	Id         int
	AuthorId   int
	PlatformId int
}

type UpdatePlatformInput struct {
	Title *string `json:"title"`
	Text  *string `json:"text"`
}

func (i UpdatePlatformInput) Validate() error {
	if i.Title == nil && i.Text == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateAuthorInput struct {
	Name *string `json:"name"`
}

func (i UpdateAuthorInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
