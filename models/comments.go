package models

import (
	"database/sql"
	"time"

	"../interfaces"
)

type Comments struct {
	Db *sql.DB
}

func (model *Comments) Get(taskID int) []interfaces.Comment {
	db := model.Db

	commentRows, err := db.Query(
		`SELECT ID, content, author, date
		 FROM comments
		 WHERE task_id=?`, taskID)

	if err != nil {
		panic(err.Error())
	}

	comments := []interfaces.Comment{}

	for commentRows.Next() {
		var conmentId int
		var content string
		var author string
		var date string
		err = commentRows.Scan(&conmentId, &content, &author, &date)

		if err != nil {
			panic(err.Error())
		}

		comment := interfaces.Comment{conmentId, content, author, date}
		comments = append(comments, comment)
	}

	return comments
}

func (model *Comments) GetById(commentID int) *interfaces.Comment {
	db := model.Db

	commentRows := db.QueryRow(
		`SELECT ID, content, author, date
		 FROM comments
		 WHERE ID=?`, commentID)

	var conmentId int
	var content string
	var author string
	var date string
	err := commentRows.Scan(&conmentId, &content, &author, &date)

	if err != nil {
		return nil
		// panic(err.Error())
	}

	comment := &interfaces.Comment{conmentId, content, author, date}

	return comment
}

func (model *Comments) Add(newComment *interfaces.Comment, taskID int) (int64, error) {
	db := model.Db
	result, err := db.Exec(
		`INSERT INTO comments(ID, content, author, task_id, date)
		 VALUES(?, ?, ?, ?, ?)`,
		nil, newComment.Content, newComment.Author, taskID, time.Now().String())
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}

func (model *Comments) Update(updatedComment *interfaces.Comment) error {
	db := model.Db
	_, err := db.Exec(
		`UPDATE comments
		 SET content = ?
		 WHERE id = ?`,
		updatedComment.Content, updatedComment.Id)

	if err != nil {
		return err
	}

	return nil
}

func (model *Comments) Delete(id int) error {
	db := model.Db
	_, err := db.Exec(`DELETE FROM comments WHERE id = ?`, id)

	return err
}
