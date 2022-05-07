package city

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init(mux *http.ServeMux) {
	mux.Handle(env.CallDistanceTimeTwoCitiesR, &DistanceTimeService)
	mux.Handle(env.FindCityByCharR, &FindByCharService)
	mux.Handle(env.DeleteCityByIdR, &DeleteCity)
	mux.Handle(env.EditCityByIdR, &EditCity)
	mux.Handle(env.FindCityByIdR, &FindCityID)
	mux.Handle(env.ListOfCitiesR, &GetCity)
	mux.Handle(env.InsertNewCityR, &NewCity)
}
