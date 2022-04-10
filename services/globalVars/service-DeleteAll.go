package globalVars

import "tel-note/protocol"

func DeleteAllGlobalVars() {
	CustomerMapStore = make(map[uint]*protocol.Customer)
}
