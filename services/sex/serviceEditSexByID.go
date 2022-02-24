package sex

import (
	"tel-note/protocol"
)

func EditSexByID(ID uint8, NewSexName string) protocol.ResponseStatus {
	if storage.EditSex(ID, NewSexName) {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}
