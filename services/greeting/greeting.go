package greeting

import (
	"encoding/json"
	"os"
	"tel-note/protocol"
)

func ReadGreetingNote() (string, string) {
	data := new(protocol.Greeting)
	greetingConfig, _ := os.ReadFile("internal/env/greeting.json")
	json.Unmarshal([]byte(greetingConfig), &data)
	return data.General, data.Description

}

func ChangeGreeting(newNote, newDesc string) {
	data := new(protocol.Greeting)
	data.General, data.Description = newNote, newDesc
	result, _ := json.Marshal(data)
	_ = os.WriteFile("internal/env/greeting.json", result, 0644)
}
