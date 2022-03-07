package customer

import "tel-note/protocol"

func GetCustomerGroupRelation() protocol.CustomerGRelationStorage {
	return storageRelation.GetCustomerGroupRelation()
}
