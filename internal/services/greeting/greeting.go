package greeting

import (
	"encoding/json"
	"os"
	"tel-note/internal/config"
)

func ReadGreetingNote() (string, string) {
	data := new(config.Greeting)
	greetingConfig, _ := os.ReadFile("internal/config/greeting.json")
	json.Unmarshal([]byte(greetingConfig), &data)
	return data.General, data.Description

}

func ChangeGreeting(newNote, newDesc string) {
	data := new(config.Greeting)
	data.General, data.Description = newNote, newDesc
	result, _ := json.Marshal(data)
	_ = os.WriteFile("internal/config/greeting.json", result, 0644)
}
