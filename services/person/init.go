package person

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.FindOnePersonByIdR, &FindPersonID)
	http.Handle(env.ListOfPersonR, &GetPerson)
	http.Handle(env.DeletePersonR, &DeletePerson)
	http.Handle(env.EditPersonR, &EditPerson)
}
