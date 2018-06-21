// Package tempconv performs Celsius and Fahrenheit conversions
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

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g degreeF", f)
}

const (
	// AbsoluteZeroC absolute zero in Celsius
	AbsoluteZeroC Celsius = -273.15

	// FreezingC freezing temperature in Celsius
	FreezingC Celsius = 0

	// BoilingC boiling point in Celsius
	BoilingC Celsius = 100
)
