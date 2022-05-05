package globalVars

import "github.com/ar-mokhtari/tel-note/protocol"

func DeleteAllGlobalVars() {
	CustomerMapStore = make(map[uint]*protocol.Customer)
}
