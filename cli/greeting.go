package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tel-note/env"
	"tel-note/services/greeting"
	"tel-note/services/identity"
)

func ShowGreeting() {
	if identity.IsRegulator {
		fmt.Println("Hi Regulator ... ")
	}
	fmt.Println(greeting.ReadGreetingNote())
}

func ChangeGreeting() {
	fmt.Println("Do you want to change greeting? (yes or no)")
	var changeStatus string
	fmt.Scanln(&changeStatus)
	if strings.ToLower(changeStatus) == env.YES {
		note := bufio.NewScanner(os.Stdin)
		desc := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter new greeting note:")
		note.Scan()
		fmt.Println("Enter new greeting description:")
		desc.Scan()
		greeting.ChangeGreeting(note.Text(), desc.Text())
	}
}
