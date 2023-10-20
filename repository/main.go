package repository

import (
	"aldinoh8/example-gin/entity"

	"gorm.io/gorm"
)

type Main struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Main {
	return Main{DB: db}
}

func (repository Main) Create(task, description string) (entity.Todo, error) {
	todo := entity.Todo{
		Task:        task,
		Description: description,
	}

	db := repository.DB

	if err := db.Create(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (repository Main) FindAll() ([]entity.Todo, error) {
	result := []entity.Todo{}

	db := repository.DB

	if err := db.Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}
