package contact

import (
	"encoding/json"
	"net/http"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type contactPool struct{}

var PoolContact contactPool

func (allData *contactPool) NewContact(inputContact protocol.Contact) (status protocol.ResponseStatus, resData []*protocol.Contact) {
	if status.State = storage.AddContact(inputContact); status.State {
		return protocol.ResponseStatus{State: true}, resData
	}
	return protocol.ResponseStatus{State: false}, []*protocol.Contact{}

}

func (allData *contactPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	if status, _ := allData.NewContact(result); status.State {
		json.NewEncoder(w).Encode(struct {
			State   int
			Message string
		}{200, "New contact added ..."})
	} else {
		json.NewEncoder(w).Encode(struct {
			State   int
			Message string
		}{400, "some thing wrong"})
	}
}
