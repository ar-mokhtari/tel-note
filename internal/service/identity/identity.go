package identity

import "fmt"

var IsAdmin bool

func Identity() {
	var userStatus string
	fmt.Println("user or admin?")
	fmt.Scanln(&userStatus)
	if userStatus == "admin" {
		IsAdmin = true
	}
}
