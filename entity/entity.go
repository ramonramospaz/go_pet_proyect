package entity

import (
	"log"
	"twelve_webapp/util"
)

func init() {
	log.Printf("Loading structures")
	err := util.Load(&users, "/tmp/users.gob")
	if err != nil {
		log.Println("The file users.gob cant be load or was not created yet.")
	}
	err = util.Load(&userIndex, "/tmp/userIndex.gob")
	if err != nil {
		log.Println("The file userIndex.gob cant be load or was not created yet.")
		userIndex = make(map[int]*User)
	}
}

// SaveFiles save all files
func SaveFiles() {
	log.Printf("Storing structures")
	err := util.Store(users, "/tmp/users.gob")
	if err != nil {
		log.Fatalln("The file users.gob cant be created")
	}
	err = util.Store(userIndex, "/tmp/userIndex.gob")
	if err != nil {
		log.Println("The file userIndex.gob cant be created")
	}
}
