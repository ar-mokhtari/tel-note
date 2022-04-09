package sex

import "tel-note/protocol"

func GetSex() []*protocol.Sex {
	return storage.GetSex()
}
