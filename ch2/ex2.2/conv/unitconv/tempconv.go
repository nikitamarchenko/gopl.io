package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64

type Meters float64
type Feet float64

type Kilograms float64
type Pounds float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (f Feet) String() string { return fmt.Sprintf("%gft", f)}
func (m Meters) String() string { return fmt.Sprintf("%gm", m)}

func (k Kilograms) String() string {return fmt.Sprintf("%gkg", k)}
func (p Pounds) String() string {return fmt.Sprintf("%glbs", p)}