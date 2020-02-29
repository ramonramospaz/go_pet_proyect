package controler

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	DeleteAllUsers()
	user1 := User{
		ID:           15281565,
		Name:         "Ramon",
		Age:          38,
		BirthCountry: "Venezuela",
	}
	err := CreateUser(user1)
	if err != nil {
		t.Error("The user cant be created")
	}

	user2 := User{
		ID:           0,
		Name:         "",
		Age:          0,
		BirthCountry: "",
	}
	err = CreateUser(user2)
	if err == nil {
		t.Error("The user cant be created with empty values")
	}

}

func TestListUsers(t *testing.T) {
	DeleteAllUsers()
	listUsers := ListUsers()
	if len(listUsers) != 0 {
		t.Error("The list of user need to be empty")
	}
	user := User{
		ID:           15281565,
		Name:         "Ramon",
		Age:          38,
		BirthCountry: "Venezuela",
	}
	CreateUser(user)
	user.ID = 15281566
	user.Name = "Fabiola"
	CreateUser(user)
	user.ID = 23446172
	user.Name = "Franlis"
	CreateUser(user)
	listUsers = ListUsers()
	if len(listUsers) != 3 {
		t.Error("The list of user need to be only 3")
	}
	if (listUsers[0].ID != 15281565) &&
		(listUsers[1].ID != 15281566) &&
		(listUsers[2].ID != 23446172) {
		t.Error("The list of user dont have the right ID")
	}
}

func TestFindUserById(t *testing.T) {
	DeleteAllUsers()
	user1 := User{
		ID:           15281565,
		Name:         "Ramon",
		Age:          38,
		BirthCountry: "Venezuela",
	}
	CreateUser(user1)
	user2 := User{
		ID:           15281567,
		Name:         "Fabiola",
		Age:          37,
		BirthCountry: "Venezuela",
	}
	CreateUser(user2)
	userFind, _ := FindUserByID(15281565)
	if userFind.ID != user1.ID {
		t.Error("The user was not found")
	}

	userFind, err := FindUserByID(15281568)
	if err == nil {
		t.Error("There were not user with that ID")
	}
}

func TestFindAndDeleteUserByID(t *testing.T) {

}
