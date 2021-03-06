package models

type Person struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Race    string `json:"race"`
	Age     uint   `json:"age"`
	Alive   bool   `json:"alive"`
	Jobs    []Job  `json:"jobs" `
}

type CreatePersonInput struct {
	Name    string `json:"name" bindings:"required"`
	Surname string `json:"surname" bindings:"required"`
	Race    string `json:"race" bindings:"required"`
	Age     uint   `json:"age" bindings:"required"`
	Alive   bool   `json:"alive"`
	Jobs    []Job  `json:"jobs" `
}

type UpdatePersonInput struct {
	Name    string `json:"name" `
	Surname string `json:"surname" `
	Race    string `json:"race"`
	Age     uint   `json:"age" `
	Alive   bool   `json:"alive"`
	Jobs    []Job  `json:"jobs" `
}
