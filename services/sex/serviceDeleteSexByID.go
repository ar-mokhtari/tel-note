package sex

import "tel-note/protocol"

func DeleteSexByID(ID uint8) protocol.ResponseStatus {
	if storage.DeleteSex(ID) {
		return protocol.ResponseStatus{State: true}
	}
	return protocol.ResponseStatus{State: false}
}
