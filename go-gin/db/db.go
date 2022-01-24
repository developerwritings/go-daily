package db

import (
	"fmt"
	"go_gin_crud/models"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func SetupDB() *gorm.DB {

	err := godotenv.Load("local.env")
	if err != nil {
		//log.Fatalf("Some error occured. Err: %s", err)
		fmt.Println("ERROR: ", err)
	}

	//user := os.Getenv("USER")
	//fmt.Println(user)
	VAL := os.Getenv("VAL")
	PASS := os.Getenv("PASS")
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	DBNAME := os.Getenv("DBNAME")
	fmt.Println("Pas", VAL, PASS)
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", VAL, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.User{})
	return db
}
