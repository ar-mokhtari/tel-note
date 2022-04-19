package customer

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.CustomerListR, &GetCustomer)
	http.Handle(env.CustomerGroupListR, &GetCustomerGroup)
	http.Handle(env.CustGRelationListR, &GetCustomerGroupRelation)
	http.Handle(env.NewCustomerR, &AddCustomer)
	http.Handle(env.DeleteCustomerR, &DeleteCustomer)
	http.Handle(env.EditCustomerR, &EditCustomer)
	http.Handle(env.FindCustGroupIDR, &FindCustomerGroupID)
}
