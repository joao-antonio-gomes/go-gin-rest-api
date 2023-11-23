package entity

import "gopkg.in/validator.v2"

type Student struct {
	ID        int    `json:"id"`
	Name      string `json:"name" validate:"nonzero,min=3,max=100,regexp=^[a-zA-Z ]*$"`
	CPF       string `json:"cpf" validate:"regexp=^[0-9]{11}$,nonzero"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func ValidateStudent(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}
