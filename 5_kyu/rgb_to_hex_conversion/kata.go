// https://www.codewars.com/kata/513e08acc600c94f01000001/train/go
package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func decimalToHex(n int) string {
	if n > 255 {
		n = 255
	}

	if n < 0 {
		n = 0
	}

	return strconv.FormatInt(int64(n), 16)
}

func RGB(r, g, b int) string {
	output := fmt.Sprintf("%02s%02s%02s", decimalToHex(r), decimalToHex(g), decimalToHex(b))

	return strings.ToUpper(output)
}
