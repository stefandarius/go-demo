package tasks

import "go-server/dtos"

type TaskService struct{}

func New() TaskService {
	return TaskService{}
}

func (ts TaskService) GetTasks() ([]dtos.Task, error) {
	return []dtos.Task{
		{Id: 1, Description: "Buy Milk", Completed: false},
		{Id: 2, Description: "Buy Eggs", Completed: false},
		{Id: 3, Description: "Buy Bread", Completed: false},
	}, nil
}
