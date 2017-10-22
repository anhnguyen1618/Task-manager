package models

import (
	"database/sql"
	"time"

	"github.com/anhnguyen300795/Task-manager/interfaces"
)

type Tasks struct {
	DB *sql.DB
}

func (model *Tasks) GetAll() []interfaces.TaskQuery {
	db := model.DB
	rows, err := db.Query(
		`SELECT ID, title, status, assignee, assignor, start_time, end_time,description
		 	FROM tasks`)
	if err != nil {
		panic(err.Error())
	}

	tasks := []interfaces.TaskQuery{}

	commentModel := &Comments{model.DB}

	for rows.Next() {
		var id int
		var title string
		var status string
		var assignee string
		var assignor string
		var start_time string
		var end_time string
		var description string
		err = rows.Scan(&id, &title, &status, &assignee, &assignor, &start_time, &end_time, &description)

		if err != nil {
			panic(err.Error())
		}

		comments := commentModel.GetByTaskID(id)

		task := interfaces.TaskQuery{id, title, status, assignee, assignor, start_time, end_time, description, comments}
		tasks = append(tasks, task)
	}

	return tasks
}

func (model *Tasks) GetOne(id int) *interfaces.TaskQuery {
	db := model.DB
	row := db.QueryRow(
		`SELECT title, status, assignee, assignor, start_time, end_time,description
		 	FROM tasks
		 	WHERE ID = ?`, id)

	var title string
	var status string
	var assignee string
	var assignor string
	var start_time string
	var end_time string
	var description string
	err := row.Scan(&title, &status, &assignee, &assignor, &start_time, &end_time, &description)
	if err != nil {
		return nil
	}

	commentModel := &Comments{model.DB}
	comments := commentModel.GetByTaskID(id)
	task := &interfaces.TaskQuery{id, title, status, assignee, assignor, start_time, end_time, description, comments}
	return task
}

func (model *Tasks) Add(task *interfaces.Task) (int64, error) {
	db := model.DB

	result, err := db.Exec(
		`INSERT INTO tasks(id, title, status, assignee, assignor, start_time, end_time, description)
		 VALUES(?, ?, ?, ?, ?, ?, ?, ?)`,
		nil, task.Title, task.Status, task.Assignee, task.Assignor, time.Now().String(), task.EndTime, task.Description)

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}

func (model *Tasks) Update(task *interfaces.Task) error {
	db := model.DB

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

func (model *Tasks) Delete(id int) error {
	db := model.DB

	_, err := db.Exec(`DELETE FROM tasks WHERE id = ?`, id)

	return err
}
