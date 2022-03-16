package contact

import (
	"tel-note/protocol"
)

func FindContactByChar(insertChar string) (protocol.ContactStorage, uint) {
	if resultData, data := storage.FindContactByChar(insertChar); resultData {
		return data, uint(len(data.ContactData))
	}
	return protocol.ContactStorage{}, 0
}
