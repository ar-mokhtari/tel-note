package protocol

type (
	CustomerGroupRelation struct {
		ID         uint `json:"id"`
		CustomerID uint `json:"customer_id"`
		GroupID    uint `json:"group_id"`
	}
	CustomerGRelationStorage      []*CustomerGroupRelation
	CustomerGroupRelationServices interface {
		GetCustomerGroupRelation() CustomerGRelationStorage
		NewRelation(customerID uint, groupID uint)
		FindCustomerGroupRelationByGroupID(ID uint) CustomerStorage
	}
)
