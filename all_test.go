package tmof

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestHexToIPv4(t *testing.T) {
	Expect(t, HexToIPv4("A402000A")).ToBe("164.2.0.10")
	Expect(t, HexToIPv4("3801000A")).ToBe("56.1.0.10")
	Expect(t, HexToIPv4("1701000A")).ToBe("23.1.0.10")
}
