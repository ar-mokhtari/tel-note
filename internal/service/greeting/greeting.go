package greeting

import (
	"encoding/json"
	"fmt"
	"os"
	"tel-note/internal/config"
	"tel-note/internal/service/identity"
)

func ReadGreetingNote() (string, string) {
	data := new(config.Greeting)
	plan, _ := os.ReadFile("internal/config/greeting.json")
	json.Unmarshal([]byte(plan), &data)
	return data.General, data.Description

}

func ChangeGreeting() {
	fmt.Println("Do you want to change greeting? (yes or no)")
	var changeStatus string
	fmt.Scanln(&changeStatus)
	if changeStatus == config.OkStatus {
		var note, desc string
		fmt.Println("Enter new greeting note:")
		fmt.Scanln(&note)
		fmt.Println("Enter new greeting description:")
		fmt.Scanln(&desc)
		data := new(config.Greeting)
		data.General, data.Description = note, desc
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
