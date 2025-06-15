package main

import (
	"04-associations/data"
	"04-associations/domain"
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

	db, err := gorm.Open(sqlite.Open("associates.db"), gormConfig)
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	// AutoMigrate will only create tables, missing columns and missing indexes, and won't change existing column's type or delete unused columns to protect your data.
	err = db.AutoMigrate(&domain.Company{}, &domain.CreditCard{}, &domain.Photo{}, &domain.Language{}, &domain.User{})
	if err != nil {
		panic("failed to migrate schema")
	}

	// populate data
	data.PopulateData(db)

	// Query (Read)
	var user domain.User
	db.Preload("Company").Preload("CreditCard").Preload("Photos").First(&user, 1)
	// SELECT * FROM `company` WHERE `company`.`id` = 1 AND `company`.`deleted_at` IS NULL
	// SELECT * FROM `credit_card` WHERE `credit_card`.`user_id` = 1 AND `credit_card`.`deleted_at` IS NULL
	// SELECT * FROM `photo` WHERE `photo`.`user_id` = 1 AND `photo`.`deleted_at` IS NULL
	// SELECT * FROM `user` WHERE `user`.`id` = 1 AND `user`.`deleted_at` IS NULL ORDER BY `user`.`id` LIMIT 1

	var joinsUser domain.User
	db.Joins("Company").Joins("CreditCard").First(&joinsUser, 1)

	fmt.Println("exit")
}
