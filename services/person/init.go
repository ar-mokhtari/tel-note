package person

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.FindOnePersonByIdR, &ServPersonRoute)
	http.Handle(env.ListOfPersonR, &GetPerson)
}
