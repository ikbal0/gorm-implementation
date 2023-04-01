package main

import (
	"errors"
	"fmt"
	"gorm-implementation/database"
	"gorm-implementation/entity"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	// createUser("john.doe2@mail.com")
	getUserById(1)
}

func getUserById(id uint) {
	db := database.GetDB()

	user := entity.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("error finding user:", err)
	}

	fmt.Printf("User Data: %+v \n", user)
}

// func createUser(email string) {
// 	db := database.GetDB()

// 	User := entity.User{
// 		Email: email,
// 	}

// 	tx := db.Begin()

// 	err := db.Create(&User).Error

// 	if err != nil {
// 		tx.Rollback()
// 		fmt.Println("Error creating user data:", err)
// 		return
// 	}

// 	fmt.Println("New User Data:", User)
// }
