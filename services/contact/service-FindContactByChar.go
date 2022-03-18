package contact

import (
	"tel-note/protocol"
)

func FindContactByChar(insertChar string) ([]*protocol.Contact, uint) {
	if resultData, data := storage.FindContactByChar(insertChar); resultData {
		return data, uint(len(data))
	}
	return []*protocol.Contact{}, 0
}
