package greeting

import (
	"TelNote/internal/config"
	"TelNote/internal/service/identity"
	"encoding/json"
	"fmt"
	"os"
)

func ReadGreetingNote() (string, string) {
	data := new(config.Greeting)
	plan, _ := os.ReadFile("internal/config/config.json")
	json.Unmarshal([]byte(plan), &data)
	return data.General, data.Description

}

func ChangeGreeting() {
	println("Do you want to change greeting? (yes or no)")
	var changeStatus string
	fmt.Scanln(&changeStatus)
	if changeStatus == "yes" {
		var note, desc string
		println("Enter new greeting note:")
		fmt.Scanln(&note)
		println("Enter new greeting description:")
		fmt.Scanln(&desc)
		data := new(config.Greeting)
		data.General, data.Description = note, desc
		result, _ := json.Marshal(data)
		_ = os.WriteFile("internal/config/config.json", result, 0644)
	}
}

func Greeting() {
	if identity.IsAdmin {
		println("Hi Admin ... ")
	}
	fmt.Println(ReadGreetingNote())
}
