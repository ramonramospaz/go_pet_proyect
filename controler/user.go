package controler

import (
	"errors"
	"time"
	"twelve_webapp/entity"
)

// User usuario
type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	BirthCountry string `json:"birthCountry"`
}

// CreateUser ..
func CreateUser(u User) (err error) {
	if (u.ID != 0) &&
		(u.Name != "") &&
		(u.Age != 0) &&
		(u.BirthCountry != "") {
		newUser := entity.User{
			ID:           u.ID,
			Name:         u.Name,
			Age:          u.Age,
			Birthday:     time.Now(),
			BirthCountry: u.BirthCountry,
		}
		newUser.Create()
	} else {
		err = errors.New("The fields; id, name, age, birthCountry are obligatory")
	}
	return
}

// ListUsers ...
func ListUsers() (listUsers []User) {
	users := entity.GetUserList()
	if len(users) > 0 {
		for _, u := range users {
			newUser := User{
				ID:           u.ID,
				Name:         u.Name,
				Age:          u.Age,
				BirthCountry: u.BirthCountry,
			}
			listUsers = append(listUsers, newUser)
		}
	} else {
		listUsers = make([]User, 0)
	}
	return
}

// FindUserByID ...
func FindUserByID(ID int) (user User, err error) {
	u, err := entity.FindUserByID(ID)
	if err == nil {
		user = User{
			ID:           u.ID,
			Name:         u.Name,
			Age:          u.Age,
			BirthCountry: u.BirthCountry,
		}
	}
	return
}

// FindAndDeleteUserByID yea hell name
func FindAndDeleteUserByID(ID int) (err error) {
	u, err := entity.FindUserByID(ID)
	if err == nil {
		u.Delete()
	}
	return

}

// DeleteAllUsers delete all
func DeleteAllUsers() {
	entity.DeleteAllUsers()
}
