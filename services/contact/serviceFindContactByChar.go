package contact

import (
	"tel-note/protocol"
	"unsafe"
)

func FindContactByChar(insertChar string) (protocol.ContactStorage, uint) {
	if resultData, data := storage.FindContactByChar(insertChar); resultData {
		return data, uint(unsafe.Sizeof(data))
	}
	return protocol.ContactStorage{}, 0
}
