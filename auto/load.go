package auto

import (
	"api/database"
	"api/models"
	"log"
)

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal()
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.User).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&models.User).Error
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		err =
	}
}
