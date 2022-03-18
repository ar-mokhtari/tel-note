package person

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tel-note/protocol"
)

type servPersonRoute struct{}

var ServPersonRoute *servPersonRoute

func FindPersonByID(ID uint) (state protocol.ResponseStatus, result protocol.Person) {
	if status, res := storage.FindPersonByID(ID); status {
		return protocol.ResponseStatus{State: true}, res
	}
	return protocol.ResponseStatus{State: false}, protocol.Person{}
}

func (servPerson *servPersonRoute) ServeFindPersonByIDHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	personID := r.FormValue("pid")
	personUID, _ := strconv.ParseUint(personID, 10, 8)
	if status, data := FindPersonByID(uint(personUID)); status.State {
		json.NewEncoder(w).Encode(data)
	} else {
		fmt.Fprintf(w, "not found")
	}
}
