package poem

type Poems struct {
	Id       int    `json:"id"`
	Title    string `json:"title" binding:"required"`
	Text     string `json:"text" binding:"required"`
	AuthorId int    `json:"author_id"`
}

type Authors struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}
