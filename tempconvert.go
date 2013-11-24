// tempconvert.go
package goconvtemps

import (
	"fmt"
	"strings"
)

// ConvertTemps takes a temperature of type float64 and a temperature unit of measure
// of type string and returns the converted value as type string.  Valid units of measure
// are "C" and "F". If unit is "C" (Celsius), the temperature is converted to Fahrenheit and vice versa.
// Example:
//     temp, unit := 0, C => "32F"
func ConvertTemps(temp float64, unit string) string {

	var converted float64
	var newTemp string

	if strings.ToLower(unit) == "c" {
		converted = ((temp * 9) / 5) + 32
		unit = "F"
	} else if strings.ToLower(unit) == "f" {
		converted = ((temp - 32) * 5) / 9
		unit = "C"
	} else {
		return fmt.Sprintf("Conversion Limitation: Cannot convert temperature (unit %v)", unit)
	}

	newTemp = fmt.Sprintf("%v%v", converted, strings.ToUpper(unit))
	return newTemp
}
