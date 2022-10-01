package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connection Database
func DatabaseInit() {
	var err error
	dsn := "root:@tcp(localhost:3306)/data?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	/* if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} */

	fmt.Println("Connected to Database")
}
