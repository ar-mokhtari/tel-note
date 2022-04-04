package contact

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.NewContactRecordR, &NewContactPool)
	http.Handle(env.ListOfContactR, &GetPool)
	http.Handle(env.FindAndEditContactByContactIdR, &EditPool)
	http.Handle(env.FindOneContactByIdR, &FindByIDPool)
	http.Handle(env.FindContactContainingSomeCharacterR, &FindByCharPool)
	http.Handle(env.DeleteContactByIdR, &DelContactIDPool)
	http.Handle(env.DeleteAllContactsR, &DelAllContactPool)
}
