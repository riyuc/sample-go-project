package services

import (
	"fmt"

	"github.com/riyuc/sample-go-project/models"
	"github.com/riyuc/sample-go-project/utils"
)

func Run() {
	user := models.User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	fmt.Println(user.GetFullName())
	fmt.Println(utils.HelperFunction())
}
