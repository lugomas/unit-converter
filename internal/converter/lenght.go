package converter

func ConvertLength(value float64, from, to string) float64 {
	// Conversion factors to meters
	conversionToMeter := map[string]float64{
		"millimeter": 0.001,
		"centimeter": 0.01,
		"meter":      1,
		"kilometer":  1000,
		"inch":       0.0254,
		"foot":       0.3048,
		"yard":       0.9144,
		"mile":       1609.34,
	}

	// Check if the input units exist in the map
	fromFactor, okFrom := conversionToMeter[from]
	toFactor, okTo := conversionToMeter[to]

	if !okFrom || !okTo {
		// If either the "from" or "to" unit is invalid, return the original value
		return value
	}

	// Convert the value to meters first, then to the target unit
	valueInMeters := value * fromFactor
	convertedValue := valueInMeters / toFactor

	return convertedValue
}
