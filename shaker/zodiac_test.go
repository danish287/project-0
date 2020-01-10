package shaker

import "testing"

func TestPrintZodiac(t *testing.T) {
	zodiacSigns := []string{"aquarius", "pices", "aries"}
	n := PrintZodiac(zodiacSigns)

	if len(n) == 0 {
		t.Error("No sign")
	}

}
