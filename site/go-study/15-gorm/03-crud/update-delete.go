package main

import (
	"03-crud/domain"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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

	newUser := domain.User{Name: "Jinzhu", Email: "jinzhu.chang@email.com"}
	err = saveUser(db, &newUser)
	if err != nil {
		panic("failed to create user")
	}

	err = DeleteUser(db, newUser.ID)

	newUser = domain.User{Name: "James B.", Email: "james007@secret.com"}
	err = saveUser(db, &newUser)
	if err != nil {
		panic("failed to create user")
	}
	err = DeleteUserByName(db, "James B.")
	if err != nil {
		panic("failed to delete user")
	}

	err = DeleteAllUsers(db)
	if err != nil {
		panic("failed to delete all users")
	}
	fmt.Println("exit!")
}

func saveUser(db *gorm.DB, user *domain.User) error {
	// Update user
	return db.Save(user).Error
}

// DeleteUser deletes user by id
func DeleteUser(db *gorm.DB, id uint) error {
	return db.Delete(&domain.User{}, id).Error
}

// DeleteAllUsers deletes all users
func DeleteAllUsers(db *gorm.DB) error {
	return db.Where("1 = 1").Delete(&domain.User{}).Error
}

// DeleteUserByName deletes user by name
func DeleteUserByName(db *gorm.DB, name string) error {
	return db.Where("name = ?", name).Delete(&domain.User{}).Error
}

// DeleteUsers deletes users by conditions
func DeleteUsers(db *gorm.DB, conditions map[string]interface{}) error {
	return db.Where(conditions).Delete(&domain.User{}).Error
}
