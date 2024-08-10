package db

import (
	"fmt"

	"github.com/jpdhaliwal22/Todo-Golang/model"
	"gorm.io/gorm/clause"
)

// AddTask implements Database.
func (m *mysqlDb) AddTask(task model.Task) (model.Task, error) {
	result := m.db.Create(&task)

	if result.Error != nil {
		fmt.Println("Add task failed")
		return model.Task{}, result.Error
	}

	return task, nil
}

// GetTaskList implements Database.
func (m *mysqlDb) GetTaskList(filters map[string]interface{}) ([]model.Task, error) {
	var tasks []model.Task

	result := m.db.Debug().Where(filters).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("tasks", tasks)
	return tasks, nil
}

func (m *mysqlDb) DeleteTask(id string) error {
	result := m.db.Debug().Where("id=?", id).Delete(&model.Task{})
	if result.Error != nil {

		return result.Error
	}

	return nil
}

func (m *mysqlDb) UpdateTask(task model.Task) error {
	err := m.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},                          // key colume
		DoUpdates: clause.AssignmentColumns([]string{"status", "detail"}), // column needed to be updated
	}).Create(&task).Error

	return err
}
