package interfaces

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Assignee    int    `json:"assignee"`
	Assignor    int    `json:"assignor"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Description string `json:"description"`
}

type TaskQuery struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Status       string `json:"status"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
	Description  string `json:"description"`
	AssigneeName string `json:"assigneeName"`
	AssignorName string `json:"assignorName"`
}
