package person

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init() {
	http.Handle(env.FindOnePersonByIdR, &FindPersonID)
	http.Handle(env.ListOfPersonR, &GetPerson)
	http.Handle(env.DeletePersonR, &DeletePerson)
	http.Handle(env.EditPersonR, &EditPerson)
	http.Handle(env.FindPersonCharR, &FindCharPerson)
	http.Handle(env.NewPersonR, &AddPerson)
}
