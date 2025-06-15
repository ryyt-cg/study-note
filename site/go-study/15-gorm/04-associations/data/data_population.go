package data

import (
	"04-associations/domain"
	"gorm.io/gorm"
	"time"
)

func PopulateData(db *gorm.DB) {
	// Create Companies
	companies := []domain.Company{
		{Model: gorm.Model{ID: 1}, Name: "Apple"},
		{Model: gorm.Model{ID: 2}, Name: "Google"},
	}
	db.Create(companies)

	// Create CreditCards
	creditCards := []domain.CreditCard{
		{Model: gorm.Model{ID: 1}, Number: "1111-1111-1111-1111", UserID: 1},
		{Model: gorm.Model{ID: 2}, Number: "2222-2222-2222-2222", UserID: 2},
		{Model: gorm.Model{ID: 3}, Number: "3333-3333-3333-3333", UserID: 3},
	}
	db.Create(creditCards)

	// Create Photos
	photos := []domain.Photo{
		{Model: gorm.Model{ID: 1}, Description: "Photo 1", UserID: 1},
		{Model: gorm.Model{ID: 2}, Description: "Photo 2", UserID: 1},
		{Model: gorm.Model{ID: 3}, Description: "Photo 3", UserID: 2},
		{Model: gorm.Model{ID: 4}, Description: "Photo 4", UserID: 2},
		{Model: gorm.Model{ID: 5}, Description: "Photo 5", UserID: 3},
		{Model: gorm.Model{ID: 6}, Description: "Photo 6", UserID: 3},
	}
	db.Create(photos)

	languages := []domain.Language{
		{Model: gorm.Model{ID: 1}, Name: "English"},
		{Model: gorm.Model{ID: 2}, Name: "Japan"},
		{Model: gorm.Model{ID: 3}, Name: "German"},
		{Model: gorm.Model{ID: 4}, Name: "Spanish"},
	}
	db.Create(languages)

	// Create Users
	users := []domain.User{
		{Model: gorm.Model{ID: 1}, Name: "User 1", Email: "user1@gmail.com", BirthDay: time.Now(), CompanyId: 1, CreditCard: creditCards[0],
			Photos: []domain.Photo{photos[0], photos[1]}},
		{Model: gorm.Model{ID: 2}, Name: "User 2", Email: "user2@outlook.com", BirthDay: time.Now(), CompanyId: 1, CreditCard: creditCards[1],
			Photos: []domain.Photo{photos[2], photos[3]}},
		{Model: gorm.Model{ID: 3}, Name: "User 3", Email: "user3@example.com", BirthDay: time.Now(), CompanyId: 2, CreditCard: creditCards[2],
			Photos: []domain.Photo{photos[4], photos[5]}},
	}
	db.Create(users)

}
