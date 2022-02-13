package storage

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type (
	AllDataTool interface {
		GetANewCodeID()
	}
	AllData struct {
		ContactData []*Contact
		CityData    []*City
		jobData     []*jobInfo
	}
)

func (MainData *AllData) GetANewCodeID() uint {
	baseData := reflect.ValueOf(&MainData).Elem()
	for i := 0; i < baseData.NumField(); i++ {
		val := baseData.Field(i).Interface()
		subData := make(map[int]interface{})
		subData[0] = val
		for _, data := range subData {
			fmt.Println(json.Marshal(data))
		}
	}
	return 1
}
