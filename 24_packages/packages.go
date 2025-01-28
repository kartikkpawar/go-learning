package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/kartikkpawar/go-learnings/auth"
	"github.com/kartikkpawar/go-learnings/user"
)

func main() {
	auth.LoginWithCredentials("kartikkpawar", "password_kartikkpawar")

	session := auth.GetSession()

	fmt.Println(session)

	user := user.User{
		Name:  "Kartik Pawar",
		Email: "user@email,com",
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)

	color.HiRed("TESTING 3rd party package")
}
