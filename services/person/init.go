package person

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.FindOnePersonByIdR, &FindPersonID)
	http.Handle(env.ListOfPersonR, &GetPerson)
}
