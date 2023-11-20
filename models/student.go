package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model

	ID   int    `json:"id"`
	Name string `json:"name"`
	CPF  string `json:"cpf"`
}
