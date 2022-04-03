package person

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	var personObject storageMemory
	storage = &personObject

	http.Handle(env.FindOnePersonByIdR, &ServPersonRoute)
}
