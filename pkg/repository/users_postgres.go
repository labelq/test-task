package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"test"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) GetUserInfo(id int) (test.UserInfo, error) {
	var user test.UserInfo

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", usersTable)
	err := u.db.Get(&user, query, id)

	return user, err
}

func (u *UserPostgres) GetAllUsers() ([]test.UserInfo, error) {
	var user []test.UserInfo

	query := fmt.Sprintf("SELECT * FROM %s", usersTable)
	err := u.db.Select(&user, query)

	return user, err
}

func (u *UserPostgres) GetPointsTask(taskId int) (int, error) {
	var point int

	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", pointsTableTasks, tasksTable)
	err := u.db.Get(&point, query, taskId)

	return point, err
}

func (u *UserPostgres) UpdatePointsInUsersTable(point int, userId int) (int, string, error) {
	var (
		updatedName   string
		updatedPoints int
	)

	query := fmt.Sprintf("UPDATE %s SET points = points + $1 WHERE id = $2 RETURNING points, name", usersTable)
	row := u.db.QueryRow(query, point, userId)
	if err := row.Scan(&updatedPoints, &updatedName); err != nil {
		return 0, "", err
	}

	return updatedPoints, updatedName, nil
}

func (u *UserPostgres) MarkTaskUser(userId int, taskId int) (sql.Result, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, task_id) values ($1, $2)", users_tasksTable)
	result, err := u.db.Exec(query, userId, taskId)

	return result, err
}

func (u *UserPostgres) InsertReferrer(userId int, referrer test.UserReferrer) (test.UserReferrer, error) {
	var result test.UserReferrer

	query := fmt.Sprintf("UPDATE %s SET referrer = $1 WHERE id = $2 RETURNING name, referrer", usersTable)
	row := u.db.QueryRow(query, referrer.Referrer, userId)
	if err := row.Scan(&result.Name, &result.Referrer); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, errors.New("user does not exist")
		}
		return result, errors.New("failed to update user referrer")
	}

	return result, nil
}

func (u *UserPostgres) GetAllTasksUser(userId int) ([]test.UserTaskResponse, error) {
	var result []test.UserTaskResponse

	query := fmt.Sprintf("SELECT task_id FROM %s WHERE user_id = $1", users_tasksTable)
	err := u.db.Select(&result, query, userId)

	return result, err
}
