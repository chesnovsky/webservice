package main

import (
	"fmt"

	"github.com/chesnovsky.ru/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Hui",
		LastName:  "Pizda",
	}
	fmt.Println(u)
}
