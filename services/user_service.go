package services

import (
	"fmt"

	"github.com/riyuc/sample-go-project.git/models"
	"github.com/riyuc/sample-go-project.git/utils"
)

func Run() {
	user := models.User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	fmt.Println(user.GetFullName())
	fmt.Println(utils.HelperFunction())
}
