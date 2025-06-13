package service

import (
	"errors"
	"sort"
	"test"
	"test/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserInfo(id int) (string, int, *string, error) {
	user, err := s.repo.GetUserInfo(id)
	if err != nil {
		return "", 0, nil, err
	}

	name := user.Name
	points := user.Points
	referrer := user.Referrer

	return name, points, referrer, nil
}

func (s *UserService) GetLeaderBoard() ([]test.LeaderBoard, error) {
	user, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	leaders := top3Leaders(user)
	if len(leaders) == 0 {
		return nil, errors.New("no leaders found")
	}

	return leaders, nil
}

func (s *UserService) CompleteTask(userId int, userTask test.UserTaskComplete) ([]test.UserPoint, error) {
	if userTask.Complete == false {
		return []test.UserPoint{}, errors.New("user not complete")
	}

	point, err := s.repo.GetPointsTask(userTask.TaskId)
	if err != nil {
		return []test.UserPoint{}, err
	}

	points, name, err := s.repo.UpdatePointsInUsersTable(point, userId)
	if err != nil {
		return []test.UserPoint{}, err
	}

	result, err := s.repo.MarkTaskUser(userId, userTask.TaskId)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return []test.UserPoint{}, err
	}

	if rowsAffected == 0 {
		return []test.UserPoint{}, errors.New("insertion failed")
	}

	var user []test.UserPoint

	user = append(user, test.UserPoint{
		Name:  name,
		Point: points,
	})

	return user, err
}

func (s *UserService) InsertReferrer(userId int, referrer test.UserReferrer) (test.UserReferrer, error) {
	result, err := s.repo.InsertReferrer(userId, referrer)
	if err != nil {
		return result, errors.New("referrer insertion failed")
	}

	return result, err
}

func (s *UserService) GetAllTasksUser(userId int) (*test.UserTaskResponse, error) {
	tasks, err := s.repo.GetAllTasksUser(userId)
	if err != nil {
		return nil, err
	}

	name, err := s.repo.GetUserInfo(userId)
	if err != nil {
		return nil, err
	}

	userIDs := make([]int, len(tasks))

	for i, task := range tasks {
		userIDs[i] = task.TaskId
	}

	response := test.UserTaskResponse{
		Name:    name.Name,
		TaskIDs: userIDs,
	}

	return &response, nil
}

func top3Leaders(user []test.UserInfo) []test.LeaderBoard {
	sort.Slice(user, func(i, j int) bool {
		return user[i].Points > user[j].Points
	})

	var leaders []test.LeaderBoard

	if len(user) < 3 {
		for i := 0; i < len(user); i++ {
			leaders = append(leaders, test.LeaderBoard{
				Name:   user[i].Name,
				Points: user[i].Points,
			})
		}

		return leaders
	} else {
		for i := 0; i <= 2; i++ {
			leaders = append(leaders, test.LeaderBoard{
				Name:   user[i].Name,
				Points: user[i].Points,
			})
		}
	}

	return leaders
}
