package modelos

import (
	"time"
)

type User struct {
	Id       int
	Name     string
	CreateAt time.Time
	Status   bool
}

func (user *User) AddUser(id int, name string, createAt time.Time, status bool) {
	user.Id = id
	user.Name = name
	user.CreateAt = createAt
	user.Status = status
}
