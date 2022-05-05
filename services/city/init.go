package city

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init() {
	http.Handle(env.CallDistanceTimeTwoCitiesR, &DistanceTimeService)
	http.Handle(env.FindCityByCharR, &FindByCharService)
	http.Handle(env.DeleteCityByIdR, &DeleteCity)
	http.Handle(env.EditCityByIdR, &EditCity)
	http.Handle(env.FindCityByIdR, &FindCityID)
	http.Handle(env.ListOfCitiesR, &GetCity)
	http.Handle(env.InsertNewCityR, &NewCity)
}
