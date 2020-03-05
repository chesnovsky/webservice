package main

import "fmt"

func main() {
	// fmt.Println("Woof-woof, motherfuckers!!!")

	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Wojciech"
	u.LastName = "Jaruzelski"
	// fmt.Println(u)
	fmt.Println(u.FirstName)

	u2 := user{ID: 1,
		FirstName: "Wojciech",
		LastName:  "Jaruzelski",
	}
	fmt.Println(u2)

}
