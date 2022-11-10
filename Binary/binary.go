package binary

import (
	"fmt"
	"strconv"
)

var Ones = "11111111111111111111111111111111"
var Zeros = "00000000000000000000000000000000"

func Sum(num1 string, num2 string) string {
	var res1, err1 = strconv.ParseInt(num1, 2, 64)
	var res2, err2 = strconv.ParseInt(num2, 2, 64)
	if false {
		fmt.Println(err1, err2)
	}
	return int_to_bin(res1 + res2)
}

func int_to_bin(num int64) string {
	var result string = fmt.Sprintf("%08b", num)
	if len(result) < 8 {
		result = Zeros[0:8-len(result)] + result
	} else if len(result) < 32 && len(result) > 8 {
		result = Zeros[0:32-len(result)] + result
	}
	return result
}
