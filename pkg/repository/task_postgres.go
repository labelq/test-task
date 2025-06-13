package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"test"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (t *TaskPostgres) GetTaskList() ([]test.TaskList, error) {
	var list []test.TaskList

	query := fmt.Sprintf("SELECT * FROM %s", tasksTable)
	err := t.db.Select(&list, query)

	return list, err
}

func (t *TaskPostgres) AddTask(task test.AddTask) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, instruction, point) values ($1, $2, $3) RETURNING id", tasksTable)

	row := t.db.QueryRow(query, task.Name, task.Instruction, task.Point)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TaskPostgres) AddTaskNotPoint(task test.AddTaskNotPoint) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, instruction) values ($1, $2) RETURNING id", tasksTable)

	row := t.db.QueryRow(query, task.Name, task.Instruction)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
