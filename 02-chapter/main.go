package main

func main() {
	weatherData := NewWeatherData()
	
	NewCurrentConditionsDisplay(weatherData)
	NewStatisticsDisplay(weatherData)
	// forecastDisplay := ForecaseDisplay{}

	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)
}