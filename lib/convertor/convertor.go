package convertor

import (
	"errors"
	"strconv"
	"strings"
	"tel-note/lib/validator"
)

func StrToUint(input string) (output uint) {
	if result, err := strconv.Atoi(input); err == nil {
		output = uint(result)
	}
	return output
}

func StrToSlice(input string) (output []string) {
	inputPack := strings.Split(input, ",")
	return inputPack
}

func StrSliceToUintSlice(input []string) (err error, output []uint) {
	for _, data := range input {
		if validator.IsNumber(data) {
			output = append(output, StrToUint(data))
		} else {
			return errors.New("invalid format"), []uint{}
		}
	}
	return nil, output
}
