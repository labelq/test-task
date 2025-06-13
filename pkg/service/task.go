package service

import (
	"test"
	"test/pkg/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (t *TaskService) GetTaskList() ([]test.TaskList, error) {
	list, err := t.repo.GetTaskList()
	if err != nil {
		return []test.TaskList{}, err
	}

	return list, nil
}

func (t *TaskService) AddTask(task test.AddTask) (int, error) {
	if task.Point == 0 {
		var taskNotPoint test.AddTaskNotPoint
		taskNotPoint = test.AddTaskNotPoint{
			Name:        task.Name,
			Instruction: task.Instruction,
		}

		id, err := t.repo.AddTaskNotPoint(taskNotPoint)
		if err != nil {
			return 0, err
		}

		return id, nil
	}

	id, err := t.repo.AddTask(task)
	if err != nil {
		return 0, err
	}

	return id, nil
}
