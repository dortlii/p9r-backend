package models

import "gorm.io/gorm"

type Namespace struct {
	gorm.Model
	Name string `json:"name"`
}
