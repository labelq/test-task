package test

type User struct {
	Id       int     `json:"-" db:"id"`
	Name     string  `json:"name" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Points   int     `json:"points"`
	Referrer *string `json:"referrer"`
}

type UserInfo struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Password string  `json:"-" db:"password"`
	Points   int     `json:"points"`
	Referrer *string `json:"referrer"`
}

type LeaderBoard struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

type UserTaskComplete struct {
	TaskId   int  `json:"task_id" db:"id"`
	Complete bool `json:"complete"`
}

type UserPoint struct {
	Name  string `json:"name"`
	Point int    `json:"point"`
}

type UserReferrer struct {
	Name     string `json:"name"`
	Referrer int    `json:"referrer"`
}

type UserTaskResponse struct {
	Name    string `json:"user_name"`
	TaskIDs []int  `json:"task_ids"`
	TaskId  int    `json:"-" db:"task_id"`
}
