package sex

import (
	"tel-note/protocol"
)

func FindSexByID(ID uint8) (protocol.Sex, protocol.ResponseStatus) {
	if starte, data := storage.FindSexByID(ID); starte {
		return data, protocol.ResponseStatus{State: true}
	}
	return protocol.Sex{}, protocol.ResponseStatus{State: false}
}
