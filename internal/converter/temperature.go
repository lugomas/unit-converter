package converter

func ConvertTemperature(value float64, from, to string) float64 {
	if from == "Celsius" && to == "Fahrenheit" {
		return value*9/5 + 32
	} else if from == "Fahrenheit" && to == "Celsius" {
		return (value - 32) * 5 / 9
	}
	return value
}
