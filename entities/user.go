package entities

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id           int32  `json:"id"`
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Phone_number int64  `json:"phone_no"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	City         string `json:"city"`
	State        string `json:"state"`
}

var users = [] user{
	{Id: 1, First_name: "John", Last_name: "Smith", Phone_number: 904568123, Email: "john@gmial.com", Address: "UTD, APT #A", City: "Charlotte", State: "North Carolina"},
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func AddUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
