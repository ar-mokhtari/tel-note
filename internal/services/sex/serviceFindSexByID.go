package sex

import (
	"tel-note/internal/config"
	"tel-note/internal/protocol"
)

func FindSexByID(ID uint8) (protocol.Sex, config.ResponseStatus) {
	if starte, data := storage.FindSexByID(ID); starte {
		return data, config.ResponseStatus{State: true}
	}
	return protocol.Sex{}, config.ResponseStatus{State: false}
}
