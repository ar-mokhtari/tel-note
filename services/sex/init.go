package sex

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init(mux *http.ServeMux) {
	mux.Handle(env.ListOfSexR, &GetSex)
	mux.Handle(env.DeleteSexR, &DeleteSexByID)
	mux.Handle(env.EditSexR, &EditSex)
	mux.Handle(env.FindSexR, &FindSexID)
	mux.Handle(env.InsertNewSexR, &AddSex)
}
