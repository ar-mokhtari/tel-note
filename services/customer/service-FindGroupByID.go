package customer

import "tel-note/protocol"

func FindGroupByID(ID uint) protocol.CustomerGroup {
	return storageGroup.FindGroupByID(ID)
}
