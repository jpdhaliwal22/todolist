package service

import (
	"fmt"

	"github.com/jpdhaliwal22/Todo-Golang/model"
)

// AddTask implements Service.
func (s *svc) AddTask(task model.Task) (model.Task, error) {

	task, err := s.db.AddTask(task)
	if err != nil {
		return model.Task{}, err

	}

	fmt.Println("newlycreated Task ", task)

	return task, nil
}

// GetTaskList implements Service.
func (s *svc) GetTaskList(filter map[string]interface{}) ([]model.Task, error) {
	taskList, err := s.db.GetTaskList(filter)
	if err != nil {
		return nil, err

	}

	return taskList, nil
}

func (s *svc) UpdateTask(task model.Task) error {
	return s.db.UpdateTask(task)
}

func (s *svc) DeleteTask(taskId string) error {
	return s.db.DeleteTask(taskId)
}
