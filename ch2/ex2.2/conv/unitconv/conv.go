package unitconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func FToM(f Feet) Meters { return Meters(f * 0.3048) }

func MToF(m Meters) Feet { return Feet(m * 3.28084)}

func KToP(k Kilograms) Pounds { return Pounds(k * 2.20462)}

func PToK(p Pounds) Kilograms {return Kilograms(p * 0.453592)}