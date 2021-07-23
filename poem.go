package poem

type Poems struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Text     string `json:"text"`
	AuthorId int
}

type Authors struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	SurName string `json:"surname"`
	PoemsId int
}
