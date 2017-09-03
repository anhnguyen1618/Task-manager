package tasks

import (
	"../../database"
	"../../interfaces"
	CommentModel "../comments"
)

// var db = database.DBCon

func GetAll() []interfaces.TaskQuery {
	db := database.DBCon
	rows, err := db.Query(
		`SELECT A.ID, title, status, start_time, end_time, description, assigneeName, users.username AS assignorName
		 FROM
		 (SELECT tasks.ID, title, status, start_time, end_time, assignor, description, username AS assigneeName 
		 	FROM tasks INNER JOIN users ON tasks.assignee = users.id)
		 AS A INNER JOIN users ON A.assignor = users.id`)
	if err != nil {
		panic(err.Error())
	}

	tasks := []interfaces.TaskQuery{}

	for rows.Next() {
		var id int
		var title string
		var status string
		var start_time string
		var end_time string
		var description string
		var assigneeName string
		var assignorName string
		err = rows.Scan(&id, &title, &status, &start_time, &end_time, &description, &assigneeName, &assignorName)

		if err != nil {
			panic(err.Error())
		}

		comments := CommentModel.Get(id)

		task := interfaces.TaskQuery{id, title, status, start_time, end_time, description, assigneeName, assignorName, comments}
		tasks = append(tasks, task)
	}

	return tasks
}

func GetOne(id int) interfaces.TaskQuery {
	db := database.DBCon
	row := db.QueryRow(
		`SELECT A.ID, title, status, start_time, end_time, description, assigneeName, users.username AS assignorName
		 FROM
		 (SELECT tasks.ID, title, status, start_time, end_time, assignor, description, username AS assigneeName 
		 	FROM tasks INNER JOIN users ON tasks.assignee = users.id)
		 AS A INNER JOIN users ON A.assignor = users.id
		 WHERE A.ID=?`, id)

	var title string
	var status string
	var start_time string
	var end_time string
	var description string
	var assigneeName string
	var assignorName string
	err := row.Scan(&id, &title, &status, &start_time, &end_time, &description, &assigneeName, &assignorName)
	if err != nil {
		panic(err.Error())
	}
	comments := CommentModel.Get(id)
	task := interfaces.TaskQuery{id, title, status, start_time, end_time, description, assigneeName, assignorName, comments}
	return task
}

func Add(task *interfaces.Task) (int64, error) {
	db := database.DBCon

	result, err := db.Exec(
		`INSERT INTO tasks(id, title, status, assignee, assignor, start_time, end_time, description)
		 VALUES(?, ?, ?, ?, ?, ?, ?, ?)`,
		nil, task.Title, task.Status, task.Assignee, task.Assignor, task.StartTime, task.EndTime, task.Description)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}

func Update(task *interfaces.Task) error {
	db := database.DBCon

	_, err := db.Exec(
		`UPDATE tasks 
		 SET title = ?, status = ?, assignee = ?, assignor = ?, start_time = ?, end_time = ?, description = ?
		 WHERE id = ?`,
		task.Title, task.Status, task.Assignee, task.Assignor, task.StartTime, task.EndTime, task.Description, task.Id)

	if err != nil {
		return err
	}

	return err
}

func Delete(id int) error {
	db := database.DBCon

	_, err := db.Exec(`DELETE FROM tasks WHERE id = ?`, id)

	return err
}
