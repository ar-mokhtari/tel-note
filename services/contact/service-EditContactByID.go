package contact

import (
	"encoding/json"
	"net/http"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type contactEditPool struct{}

var EditPool contactEditPool

func (allData *contactEditPool) EditContactByID(newData protocol.Contact, ID uint) *protocol.ResponseStatus {
	if storage.EditContactByID(newData, ID) {
		return &protocol.ResponseStatus{State: true}
	}
	return &protocol.ResponseStatus{State: true, String: "not found"}
}

func (allData *contactEditPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := r.Header.Get("ID")
	PersonID := r.Header.Get("PersonID")
	JobID := r.Header.Get("JobID")
	Tel := r.Header.Get("Tel")
	Cellphone := r.Header.Get("CellphoneCollection")
	Address := r.Header.Get("Address")
	Description := r.Header.Get("Description")
	var cellPack []protocol.CellPhone
	cellPhoneCollection := convertor.StrToSlice(Cellphone)
	if (len(cellPhoneCollection))%2 == 0 {
		var tempCell protocol.CellPhone
		for index, data := range cellPhoneCollection {
			if (index+1)%2 != 0 {
				tempCell.CellPhone = data
			} else {
				tempCell.Description = data
				cellPack = append(cellPack, tempCell)
				tempCell = protocol.CellPhone{}
			}
		}
	}
	result := protocol.Contact{
		PersonID:            convertor.StrToUint(PersonID),
		JobID:               convertor.StrToUint(JobID),
		Tel:                 Tel,
		CellphoneCollection: cellPack,
		Address:             Address,
		Description:         Description,
	}
	if status := allData.EditContactByID(result, convertor.StrToUint(ID)); status.State {
		json.NewEncoder(w).Encode(struct {
			State   int
			Message string
		}{200, "contact edited ..."})
	} else {
		json.NewEncoder(w).Encode(struct {
			State   int
			Message string
		}{400, "some thing wrong"})
	}
}
