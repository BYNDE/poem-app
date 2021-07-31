package poem

type Poems struct {
	Id    int    `json:"id"`
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}

type Authors struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

type AuthorsList struct {
	Id       int
	AuthorId int
	PoemId   int
}
