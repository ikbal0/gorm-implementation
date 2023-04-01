package main

import (
	"fmt"
	"gorm-implementation/database"
	"gorm-implementation/entity"
)

func main() {
	database.StartDB()

	// createUser("john.doe6@mail.com")
	// getUserById(1)

	// updateUserById(1, "amanda@email.com")

	// =============================================================

	// createProduct(1, "YLO", "YYY")
	// createProduct(1, "YLO", "DDDD")

	// getUserWithProducts()

	deleteProductById(1)
}

func deleteProductById(id uint) {
	db := database.GetDB()
	product := entity.Product{}

	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted", id)
}

// func getUserWithProducts() {
// 	db := database.GetDB()

// 	users := entity.User{}
// 	err := db.Preload("Products").Find(&users).Error

// 	if err != nil {
// 		fmt.Println("Error getting user data with product:", err.Error())
// 		return
// 	}

// 	fmt.Println("User Data With Products")
// 	fmt.Printf("%+v", users)
// }

// func createProduct(userId uint, brand string, name string) {
// 	db := database.GetDB()

// 	Product := entity.Product{
// 		UserID: userId,
// 		Brand:  brand,
// 		Name:   name,
// 	}

// 	err := db.Create(&Product).Error

// 	if err != nil {
// 		fmt.Println("Error creating product data", err.Error())
// 		return
// 	}

// 	fmt.Println("New Product Data", Product)
// }

// func updateUserById(id uint, email string) {
// 	db := database.GetDB()

// 	user := entity.User{}

// 	err := db.Model(&user).Where("id = ?", id).Updates(entity.User{Email: email}).Error

// 	if err != nil {
// 		fmt.Println("Error updating user data:", err)
// 		return
// 	}

// 	fmt.Printf("Update user's email: %+v \n", user.Email)
// }

// func getUserById(id uint) {
// 	db := database.GetDB()

// 	user := entity.User{}

// 	err := db.First(&user, "id = ?", id).Error

// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			fmt.Println("User data not found")
// 			return
// 		}
// 		print("error finding user:", err)
// 	}

// 	fmt.Printf("User Data: %+v \n", user)
// }

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

// 	tx.Commit()

// 	fmt.Println("New User Data:", User)
// }
