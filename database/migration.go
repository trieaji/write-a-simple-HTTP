package database

import (
	"fmt"
	"tchtest/models"
	"tchtest/pkg/mysql"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.Data{}, &models.User{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Berhasil Luur")
}
