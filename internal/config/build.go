package config

import (
	"fmt"
	"tel-note/internal/storage/memory"
)

var (
	AllContact []*memory.Contact
	AllCity    []*memory.City
	MainData   memory.AllData
)

func Init() *Config {
	cnf := Config{DB: Database{Memory: memory.AllData{}}}
	fmt.Println(cnf)
	//switch cnf.DB {
	//case Database{Memory}:
	//	fmt.Println("")
	////case Database{Memory: memory.AllData{}}:
	////	AllContact = make([]*memory.Contact, 0)
	////	AllCity = make([]*memory.City, 0)
	////	MainData = []*memory.AllData{{ContactData: AllContact, CityData: AllCity}}
	//}
	AllContact = make([]*memory.Contact, 0)
	AllCity = make([]*memory.City, 0)
	MainData = memory.AllData{ContactData: AllContact, CityData: AllCity}
	return &cnf
}
