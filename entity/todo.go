package entity

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Task        string `json:"task"`
	Description string `json:"description"`
}
