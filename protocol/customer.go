package protocol

import "time"

type (
	Customer struct {
		PersonID    uint
		Description string
		CreateAt    time.Time
		UpdatedAt   time.Time
	}
	CustomerGroup struct {
		GroupID   uint
		GroupName string
	}
	CustomerGroupStorage  []*CustomerGroup
	CustomerGroupRelation struct {
		ID         uint
		CustomerID uint
		GroupID    uint
	}
	CustomerGRelationStorage []*CustomerGroupRelation
	CustomerStorage          struct {
		Data map[uint]*Customer
	}
	CustomerServices interface {
		NewCustomer(newCustomer Customer)
		EditCustomer(id uint, EditedCustomer Customer)
		DeleteCustomerById(id uint)
		FindCustomerByID(ID uint) Customer
	}
	CustomerGroupServices interface {
		GetCustomerGroup() CustomerGroupStorage
		NewGroup(groupName string)
		FindGroupByID(ID uint) CustomerGroup
	}
	CustomerGroupRelationServices interface {
		GetCustomerGroupRelation() CustomerGRelationStorage
		NewRelation(customerID uint, groupID uint)
		FindCustomerByGroupID(ID uint) CustomerStorage
	}
)
