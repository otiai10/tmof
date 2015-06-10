package tmof

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func HexToIPv4(hexadecimal string) string {
	a, _ := hex.DecodeString(hexadecimal)
	s := []string{}
	for _, b := range a {
		s = append(s, fmt.Sprintf("%d", b))
	}
	return strings.Join(s, ".")
}
