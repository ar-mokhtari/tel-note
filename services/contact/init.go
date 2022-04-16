package contact

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.NewContactRecordR, &NewContact)
	http.Handle(env.ListOfContactR, &GetContact)
	http.Handle(env.FindAndEditContactByContactIdR, &EditContact)
	http.Handle(env.FindOneContactByIdR, &FindContactID)
	http.Handle(env.FindContactContainingSomeCharacterR, &FindContactChar)
	http.Handle(env.DeleteContactByIdR, &DelContactID)
	http.Handle(env.DeleteAllContactsR, &DelAllContact)
}
