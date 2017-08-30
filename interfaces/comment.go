package interfaces

type Comment struct {
	Id         int    `json:"id"`
	Content    string `json:"content"`
	AuthorName string `json:"authorName"`
	AuthorID   int    `json:"authorId"`
	Date       string `json:"date"`
}
