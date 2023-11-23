package entity

type StudentRepository interface {
	FindAll() ([]Student, error)
	FindById(id string) (Student, error)
	Create(student Student) (Student, error)
	Update(student Student) (Student, error)
	Delete(id string) error
	SearchStudentByCpf(cpf string) (Student, error)
}
