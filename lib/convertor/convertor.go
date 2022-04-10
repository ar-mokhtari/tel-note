package convertor

import (
	"errors"
	"strconv"
	"strings"
	"tel-note/lib/validator"
)

func StrToUint(input string) (err error, output uint) {
	if result, err := strconv.Atoi(input); err == nil {
		output = uint(result)
		return nil, output
	} else {
		return errors.New("can't convert to uint"), output
	}
}

func StrToFloat64(input string) (err error, output float64) {
	if result, err := strconv.ParseFloat(input, 64); err == nil {
		return nil, result
	}
	return errors.New("can't convert to float"), output
}

func StrToSlice(input string) (output []string) {
	inputPack := strings.Split(input, ",")
	return inputPack
}

func StrSliceToUintSlice(input []string) (error, []uint) {
	var output []uint
	for _, data := range input {
		if validator.IsNumber(data) {
			if err, uintRes := StrToUint(data); err == nil {
				output = append(output, uintRes)
			} else {
				return err, []uint{}
			}
		} else {
			return errors.New(" invalid format "), []uint{}
		}
	}
	return nil, output
}
