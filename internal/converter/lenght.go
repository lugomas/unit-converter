package converter

func ConvertLength(value float64, from, to string) float64 {
	// Conversion logic for length (mm, cm, m, etc.)
	// Example: Convert meters to kilometers
	switch from {
	case "meter":
		if to == "kilometer" {
			return value / 1000
		}
	case "kilometer":
		if to == "meter" {
			return value * 1000
		}
	}

	return value
}
