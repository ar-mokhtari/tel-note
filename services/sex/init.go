package sex

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.ListOfSexR, &GetSex)
	http.Handle(env.DeleteSexR, &DeleteSexByID)
	http.Handle(env.EditSexR, &EditSex)
	http.Handle(env.FindSexR, &FindSexID)
	http.Handle(env.InsertNewSexR, &AddSex)
}
