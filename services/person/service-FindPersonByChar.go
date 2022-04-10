package person

import (
	"tel-note/protocol"
)

func FindPersonByChar(inputChar string) (state protocol.ResponseStatus, result []*protocol.Person) {
	if status, res := storage.FindPersonByChar(inputChar); status {
		return protocol.ResponseStatus{State: true}, res
	}
	return protocol.ResponseStatus{State: false}, []*protocol.Person{}
}
