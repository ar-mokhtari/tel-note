package protocol

type (
	CustomerGroup struct {
		GroupID   uint
		GroupName string
	}
	CustomerGroupStorage  []*CustomerGroup
	CustomerGroupServices interface {
		GetCustomerGroup() CustomerGroupStorage
		NewGroup(groupName string)
		FindGroupByID(ID uint) CustomerGroup
	}
)
