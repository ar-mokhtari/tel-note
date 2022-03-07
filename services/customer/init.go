package customer

func Init() {
	var customerObject storageMemory
	var customerGroupObject group
	var customerGRelationObject relation
	storageCustomer = &customerObject
	storageGroup = &customerGroupObject
	storageRelation = &customerGRelationObject
}
