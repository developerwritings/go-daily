package handlers

import (
	"go_gin_crud/db"
	"go_gin_crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type User struct {
// 	Name string
// }

func Handlers() {

}
func GetUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create User
	dB := db.SetupDB()
	dB.Create((&input))
	//task := User{Name: input.Name}

	c.JSON(http.StatusOK, gin.H{"data": input})
}
