package person

import "net/http"

func Init() {
	var personObject storageMemory
	storage = &personObject

	http.Handle("/find-person-id", &ServPersonRoute)
}
