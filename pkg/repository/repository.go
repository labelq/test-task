package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"test"
)

type Authorization interface {
	CreateUser(user test.User) (int, error)
	GetUser(name, password string) (test.User, error)
}

type User interface {
	GetUserInfo(id int) (test.UserInfo, error)
	GetAllUsers() ([]test.UserInfo, error)
	GetPointsTask(taskId int) (int, error)
	UpdatePointsInUsersTable(point int, userId int) (int, string, error)
	MarkTaskUser(userId int, taskId int) (sql.Result, error)
	InsertReferrer(userId int, referrer test.UserReferrer) (test.UserReferrer, error)
	GetAllTasksUser(userId int) ([]test.UserTaskResponse, error)
}

type Task interface {
	GetTaskList() ([]test.TaskList, error)
	AddTask(task test.AddTask) (int, error)
	AddTaskNotPoint(task test.AddTaskNotPoint) (int, error)
}

type Repository struct {
	Authorization
	User
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Task:          NewTaskPostgres(db),
	}
}
