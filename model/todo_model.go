package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
