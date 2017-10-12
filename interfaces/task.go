package interfaces

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Assignee    string `json:"assignee"`
	Assignor    string `json:"assignor"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Description string `json:"description"`
}

type TaskQuery struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	Assignee    string    `json:"assignee"`
	Assignor    string    `json:"assignor"`
	StartTime   string    `json:"startTime"`
	EndTime     string    `json:"endTime"`
	Description string    `json:"description"`
	Comments    []Comment `json:"comments"`
}
