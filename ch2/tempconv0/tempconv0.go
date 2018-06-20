package tempconv

import (
	"fmt"
)

// Celsius the temperature scale
type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g degreeC", c)
}

// Fahrenheit the temperature scale
type Fahrenheit float64

const (
	// AbsoluteZeroC absolute zero in Celsius
	AbsoluteZeroC Celsius = -273.15

	// FreezingC freezing temperature in Celsius
	FreezingC Celsius = 0

	// BoilingC boiling point in Celsius
	BoilingC Celsius = 100
)

// CToF convert Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC convert Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
