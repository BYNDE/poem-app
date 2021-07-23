package poem

type Poems struct {
	Id       int    `json:"id"`
	Title    string `json:"title" binding:"required"`
	Text     string `json:"text" binding:"required"`
	AuthorId int
}

type Authors struct {
	Id      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	SurName string `json:"surname" binding:"required"`
	PoemsId int
}
