package entity

import (
	"errors"
	"time"
)

// User usuario
type User struct {
	ID           int
	Name         string
	Age          int
	Birthday     time.Time
	BirthCountry string
}

var userIndex map[int]*User
var users []User

// Create create the user
func (u *User) Create() {
	if _, ok := userIndex[u.ID]; !ok {
		users = append(users, *u)
		userIndex[u.ID] = u
	}
}

// Delete delte element
func (u *User) Delete() {
	for i, user := range users {
		if u.ID == user.ID {
			users = append(users[:i], users[i+1:]...)
		}
	}
	delete(userIndex, u.ID)
}

// FindUserByID find user by id
func FindUserByID(ID int) (u *User, e error) {
	var ok bool
	u, ok = userIndex[ID]
	if !ok {
		e = errors.New("Didnt find the index")
	}
	return
}

// GetUserList get user list
func GetUserList() (user []User) {
	return users
}

// DeleteAllUsers  delete all elements
func DeleteAllUsers() {
	userIndex = make(map[int]*User)
	users = nil
}
