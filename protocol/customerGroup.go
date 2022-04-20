package protocol

type (
	CustomerGroup struct {
		GroupID   uint   `json:"group_id"`
		GroupName string `json:"group_name"`
	}
	CustomerGroupStorage  []*CustomerGroup
	CustomerGroupServices interface {
		GetCustomerGroup() CustomerGroupStorage
		NewGroup(groupName string)
		FindGroupByID(ID uint) CustomerGroup
	}
)
