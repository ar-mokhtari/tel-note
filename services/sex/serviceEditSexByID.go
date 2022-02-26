package sex

import (
	"tel-note/protocol"
)

func EditSexByID(NewSex protocol.Sex) protocol.ResponseStatus {
	if storage.EditSex(NewSex) {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}
