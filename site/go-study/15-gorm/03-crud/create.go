package main

import (
	"03-crud/domain"
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func main() {
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	db, err := gorm.Open(sqlite.Open("crud.db"), gormConfig)
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	//  CREATE TABLE `user` (`id` integer PRIMARY KEY AUTOINCREMENT,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`name` text,`email` text,`birth_day` datetime)
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		panic("failed to migrate schema")
	}

	newUser := domain.User{Name: "Jinzhu", Email: "jinzhu.chang@email.com", BirthDay: time.Now()}
	// Create new Product & assign ID number
	// CREATE INDEX `idx_user_deleted_at` ON `user`(`deleted_at`)
	err = CreateUser(db, &newUser)
	if err != nil {
		panic("failed to create user")
	}

	// Query (Read)
	user, err := GetUserByID(db, newUser.ID)
	if err != nil {
		panic("failed to get user")
	}

	fmt.Println(user)

	var userByName domain.User
	db.First(&userByName, "name = ?", "Jinzhu") // find product with code D42

	// Conditions
	var userByEmail domain.User
	db.Where("email = ?", "jinzhu.chang@email.com").First(&userByEmail)

	// LIKE
	var users []domain.User
	db.Where("name LIKE ?", "%jin%").Find(&users)

	// Get Soft Delete Users
	softDeletedUsers, err := GetSoftDeletedUsers(db)
	if err != nil {
		panic("failed to get not active users")
	}

	for _, u := range softDeletedUsers {
		jsonU, _ := json.Marshal(u)
		fmt.Println(string(jsonU))
	}

	fmt.Println("exit!")
}

func CreateUser(db *gorm.DB, user *domain.User) error {
	return db.Create(user).Error
}

// GetUserByID find user with integer primary key
func GetUserByID(db *gorm.DB, id uint) (*domain.User, error) {
	var user domain.User
	err := db.First(&user, id).Error
	return &user, err
}

// GetSoftDeletedUsers gets not active users
func GetSoftDeletedUsers(db *gorm.DB) ([]domain.User, error) {
	var users []domain.User
	err := db.Unscoped().Where("deleted_at IS NOT NULL").Find(&users).Error
	return users, err
}
