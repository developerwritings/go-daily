package main

import (
	"go_gin_crud/db"
	"go_gin_crud/routes"
)

func main() {
	// err := godotenv.Load("local.env")
	// if err != nil {
	// 	//log.Fatalf("Some error occured. Err: %s", err)
	// 	fmt.Println("ERROR: ", err)
	// }

	// user := os.Getenv("USER")
	// fmt.Println(user)
	db.SetupDB()
	router := routes.Routes()

	router.Run(":8814")
}
