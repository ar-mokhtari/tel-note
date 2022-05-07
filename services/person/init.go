package person

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init(mux *http.ServeMux) {
	mux.Handle(env.FindOnePersonByIdR, &FindPersonID)
	mux.Handle(env.ListOfPersonR, &GetPerson)
	mux.Handle(env.DeletePersonR, &DeletePerson)
	mux.Handle(env.EditPersonR, &EditPerson)
	mux.Handle(env.FindPersonCharR, &FindCharPerson)
	mux.Handle(env.NewPersonR, &AddPerson)
}
