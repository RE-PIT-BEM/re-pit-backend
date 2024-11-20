package main

import (
	"fmt"
	"log"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/constant"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/database"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/domain"
)

func main() {
	db, err := database.NewDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Seeding...")

	err = db.Create(&domain.User{
		NIM:        "245150700111020",
		Name:       "Jevon Mozart",
		Major:      "Teknik Komputer",
		Department: "PSDM",
		Password:   "$2a$12$Pt3tuFzHoXPaJuOkZWNN6.78LRpH4srmpc2AdDWra9W0JTOLTHPhS", // pw: password
		Role:       constant.ROLE_ADMIN,
	}).Error

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Seeding Done")
}
