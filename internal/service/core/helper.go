package core

import "fmt"

func FindLastIDPlusOne(MainData interface{}) uint {
	result := make(map[uint]interface{})
	result[121] = MainData
	for _, data := range result {
		fmt.Println(data)
	}
	return 1
}
