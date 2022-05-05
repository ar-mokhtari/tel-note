package customer

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init() {
	http.Handle(env.CustomerListR, &GetCustomer)
	http.Handle(env.CustomerGroupListR, &GetCustomerGroup)
	http.Handle(env.CustGRelationListR, &GetCustomerGroupRelation)
	http.Handle(env.NewCustomerR, &AddCustomer)
	http.Handle(env.DeleteCustomerR, &DeleteCustomer)
	http.Handle(env.EditCustomerR, &EditCustomer)
	http.Handle(env.FindCustGrpRelationByGrpIDR, &FindCustGrpRelationByGrpID)
	http.Handle(env.FindCustIDR, &FindCustomerByID)
	http.Handle(env.FindCustGrpIDR, &FindGroupByID)
	http.Handle(env.NewCustomerGroupR, &NewGroup)
	http.Handle(env.NewCustGRelationR, &NewGrpRelation)
}
