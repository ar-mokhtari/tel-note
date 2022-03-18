package general

import (
	"strconv"
)

func StrToUint(input string) (output uint) {
	if result, err := strconv.Atoi(input); err == nil {
		output = uint(result)
	}
	return output
}
