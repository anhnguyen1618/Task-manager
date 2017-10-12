package interfaces

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
	Date    string `json:"date"`
}
