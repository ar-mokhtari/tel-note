package protocol

type (
	CustomerGroupRelation struct {
		ID         uint
		CustomerID uint
		GroupID    uint
	}
	CustomerGRelationStorage      []*CustomerGroupRelation
	CustomerGroupRelationServices interface {
		GetCustomerGroupRelation() CustomerGRelationStorage
		NewRelation(customerID uint, groupID uint)
		FindCustomerByGroupID(ID uint) CustomerStorage
	}
)
