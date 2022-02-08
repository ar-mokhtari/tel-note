package greeting

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"tel-note/internal/config"
	"tel-note/internal/service/identity"
)

func ReadGreetingNote() (string, string) {
	data := new(config.Greeting)
	greetingConfig, _ := os.ReadFile("internal/config/greeting.json")
	json.Unmarshal([]byte(greetingConfig), &data)
	return data.General, data.Description

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
		data := new(config.Greeting)
		data.General, data.Description = note.Text(), desc.Text()
		result, _ := json.Marshal(data)
		_ = os.WriteFile("internal/config/greeting.json", result, 0644)
	}
}

func Greeting() {
	if identity.IsAdmin {
		fmt.Println("Hi Admin ... ")
	}
	fmt.Println(ReadGreetingNote())
}
