package customer

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.CustomerListR, &GetCustomer)
	http.Handle(env.CustomerGroupListR, &GetCustomerGroup)
	http.Handle(env.CustGRelationListR, &GetCustomerGroupRelation)
}
