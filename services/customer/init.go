package customer

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init(mux *http.ServeMux) {
	mux.Handle(env.CustomerListR, &GetCustomer)
	mux.Handle(env.CustomerGroupListR, &GetCustomerGroup)
	mux.Handle(env.CustGRelationListR, &GetCustomerGroupRelation)
	mux.Handle(env.NewCustomerR, &AddCustomer)
	mux.Handle(env.DeleteCustomerR, &DeleteCustomer)
	mux.Handle(env.EditCustomerR, &EditCustomer)
	mux.Handle(env.FindCustGrpRelationByGrpIDR, &FindCustGrpRelationByGrpID)
	mux.Handle(env.FindCustIDR, &FindCustomerByID)
	mux.Handle(env.FindCustGrpIDR, &FindGroupByID)
	mux.Handle(env.NewCustomerGroupR, &NewGroup)
	mux.Handle(env.NewCustGRelationR, &NewGrpRelation)
}
