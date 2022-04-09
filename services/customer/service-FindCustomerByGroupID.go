package customer

import "tel-note/protocol"

func FindCustomerByGroupID(ID uint) protocol.CustomerStorage {
	return storageRelation.FindCustomerByGroupID(ID)
}
