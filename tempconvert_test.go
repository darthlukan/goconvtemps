package goconvtemps

import (
	"strings"
	"testing"
)

// TestConvertTemps tests the ConvertTemps function to make sure that it actually
// converts Fahrenheit to Celsius and vice versa.  It also checks that ConvertTemps
// supplies an appropriate error if passed in a unit such as Kelvin as ConvertTemps
// is limited to Celsius and Fahrenheit conversions.
func TestConvertTemps(t *testing.T) {
	const f = 32
	const c = 0

	tempc := ConvertTemps(f, "F")
	if strings.Contains(tempc, "F") {
		t.Errorf("ConvertTemps(%v, %v) = %v, want %v", f, "F", tempc, "C")
	}

	if tempc != "0C" {
		t.Errorf("ConvertTemps(%v, %v) = %v, want %v", f, "F", tempc, "0C")
	}

	tempf := ConvertTemps(c, "C")
	if strings.Contains(tempf, "C") {
		t.Errorf("ConvertTemps(%v, %v) = %v, want %v", c, "C", tempf, "F")
	}

	if tempf != "32F" {
		t.Errorf("ConvertTemps(%v, %v) = %v, want %v", c, "C", tempf, "32F")
	}

	tempk := ConvertTemps(f, "K")
	if !strings.Contains(tempk, "Conversion") {
		t.Errorf("ConvertTemps(%v, %v) = %v, want %v", f, "K", tempk, "Conversion Limitation Error")
	}
}
