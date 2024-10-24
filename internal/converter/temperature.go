package converter

func ConvertTemperature(value float64, from, to string) float64 {
	if from == "celsius" && to == "fahrenheit" {
		return value*9/5 + 32
	} else if from == "fahrenheit" && to == "celsius" {
		return (value - 32) * 5 / 9
	}
	return value
}
