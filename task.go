package test

type TaskList struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Instruction string `json:"instruction"`
	Point       int    `json:"point"`
}

type AddTask struct {
	Name        string `json:"name"`
	Instruction string `json:"instruction"`
	Point       int    `json:"point"`
}

type AddTaskNotPoint struct {
	Name        string `json:"name"`
	Instruction string `json:"instruction"`
}
