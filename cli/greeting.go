package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tel-note/internal/config"
	"tel-note/internal/service/greeting"
	"tel-note/internal/service/identity"
)

func ShowGreeting() {
	if identity.IsAdmin {
		fmt.Println("Hi Admin ... ")
	}
	fmt.Println(greeting.ReadGreetingNote())
}

func ChangeGreeting() {
	fmt.Println("Do you want to change greeting? (yes or no)")
	var changeStatus string
	fmt.Scanln(&changeStatus)
	if strings.ToLower(changeStatus) == config.OkStatus {
		note := bufio.NewScanner(os.Stdin)
		desc := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter new greeting note:")
		note.Scan()
		fmt.Println("Enter new greeting description:")
		desc.Scan()
		greeting.ChangeGreeting(note.Text(), desc.Text())
	}
}
