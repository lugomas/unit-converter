package converter

func ConvertWeight(value float64, from, to string) float64 {
	// Conversion logic for weight (mg, g, kg, etc.)
	// Example: Convert grams to kilograms
	if from == "gram" && to == "kilogram" {
		return value / 1000
	}
	return value
}
