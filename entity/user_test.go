package entity

import (
	"testing"
)

func TestUser(t *testing.T) {
	DeleteAllUsers()
	user1 := User{ID: 15281565, Name: "Ramon", Age: 38, BirthCountry: "Venezuela"}
	t.Logf("User created %v\n", user1)
	user1.Create()
	listUser := GetUserList()
	t.Logf("List created %v\n", listUser)
	if len(listUser) != 0 {
		if user1.ID != listUser[0].ID {
			t.Error("The user was not created")
		}
	} else {
		t.Error("The user can be saved")
	}

}

func TestDuplicatedUser(t *testing.T) {
	DeleteAllUsers()
	t.Log("Creating Users")
	user1 := User{ID: 15281565, Name: "Ramon", Age: 38, BirthCountry: "Venezuela"}
	user1.Create()
	t.Logf("User created %v\n", user1)
	user2 := User{ID: 15281565, Name: "Ramon", Age: 38, BirthCountry: "Venezuela"}
	user2.Create()
	t.Logf("User duplicated created %v\n", user2)
	listUser := GetUserList()
	if len(listUser) != 1 {
		t.Error("User duplicated created")
	}

}

func TestCreationListUser(t *testing.T) {
	DeleteAllUsers()
	user1 := User{ID: 15281565, Name: "Ramon", Age: 38, BirthCountry: "Venezuela"}
	user1.Create()
	user2 := User{ID: 15281567, Name: "Fabiola", Age: 37, BirthCountry: "Venezuela"}
	user2.Create()
	user3 := User{ID: 23446172, Name: "Franlis", Age: 29, BirthCountry: "Venezuela"}
	user3.Create()
	user4 := User{ID: 10456789, Name: "Papa", Age: 68, BirthCountry: "Venezuela"}
	user4.Create()
	listUser := GetUserList()

	if len(listUser) != 4 {
		t.Error("Error creating users")
	}

	if listUser[0].ID != 15281565 {
		t.Error("The first user is not right")
	}

	if listUser[1].ID != 15281567 {
		t.Error("The Second user is not right")
	}

	if listUser[2].ID != 23446172 {
		t.Error("The third user is not right")
	}

	if listUser[3].ID != 10456789 {
		t.Error("The four user is not right")
	}
}

func TestFindUser(t *testing.T) {
	DeleteAllUsers()
	user1 := User{ID: 15281565, Name: "Ramon", Age: 38, BirthCountry: "Venezuela"}
	user1.Create()
	user2 := User{ID: 15281567, Name: "Fabiola", Age: 37, BirthCountry: "Venezuela"}
	user2.Create()
	user3 := User{ID: 23446172, Name: "Franlis", Age: 29, BirthCountry: "Venezuela"}
	user3.Create()
	user4 := User{ID: 10456789, Name: "Papa", Age: 68, BirthCountry: "Venezuela"}
	user4.Create()
	user, e := FindUserByID(15281567)
	if e != nil {
		t.Error("The user was not found")
	}
	if user.ID != 15281567 {
		t.Error("The user is not right")
	}
}

func TestDeleteUser(t *testing.T) {
	DeleteAllUsers()
	user1 := User{ID: 15281565, Name: "Ramon", Age: 38, BirthCountry: "Venezuela"}
	user1.Create()
	user2 := User{ID: 15281567, Name: "Fabiola", Age: 37, BirthCountry: "Venezuela"}
	user2.Create()
	user3 := User{ID: 23446172, Name: "Franlis", Age: 29, BirthCountry: "Venezuela"}
	user3.Create()
	user4 := User{ID: 10456789, Name: "Papa", Age: 68, BirthCountry: "Venezuela"}
	user4.Create()
	beforeErase := len(users)
	user3.Delete()
	afterErase := len(users)
	if beforeErase == afterErase {
		t.Error("The user was not erased")
	}
	_, e := FindUserByID(user3.ID)
	if e == nil {
		t.Error("The user was not erased")
	}

	listUser := GetUserList()

	if len(listUser) == 3 {
		if listUser[0].ID != 15281565 {
			t.Error("The first user is wrong")
		}

		if listUser[1].ID != 15281567 {
			t.Error("The second user is wrong")
		}

		if listUser[2].ID != 10456789 {
			t.Error("The third user is wrong")
		}
	} else {
		t.Error("The amount of users are wrong")
	}

	user1.Delete()

	listUser = GetUserList()

	if len(listUser) == 2 {
		if listUser[0].ID == 15281565 {
			t.Error("The first user is wrong")
		}
	} else {
		t.Error("The amount of users are wrong")
	}

}

func TestStoreStruct(t *testing.T) {
	t.Log("This test is only to create the file")
	SaveFiles()
}
