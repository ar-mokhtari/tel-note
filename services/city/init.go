package city

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.CallDistanceTimeTwoCitiesR, &DistanceTimeService)
	http.Handle(env.FindCityByCharR, &FindByCharService)
	http.Handle(env.DeleteCityByIdR, &DeleteCityPool)
	http.Handle(env.EditCityByIdR, &EditCity)
	http.Handle(env.FindCityByIdR, &FindCityIDPool)
	http.Handle(env.ListOfCitiesR, &GetCityPool)
	http.Handle(env.InsertNewCityR, &NewCityPool)
}
