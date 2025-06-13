package service

import (
	"test"
	"test/pkg/repository"
)

type Authorization interface {
	CreateUser(user test.User) (int, error)
	GenerateToken(name, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetUserInfo(id int) (string, int, *string, error)
	GetLeaderBoard() ([]test.LeaderBoard, error)
	CompleteTask(userId int, userTask test.UserTaskComplete) ([]test.UserPoint, error)
	InsertReferrer(userId int, referrer test.UserReferrer) (test.UserReferrer, error)
	GetAllTasksUser(userId int) (*test.UserTaskResponse, error)
}

type Task interface {
	GetTaskList() ([]test.TaskList, error)
	AddTask(task test.AddTask) (int, error)
}

type Service struct {
	Authorization
	User
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		Task:          NewTaskService(repos.Task),
	}
}
