package main

import "fmt"

type Observer interface {
	Update(float64, float64, float64)
}

type DisplayElement interface {
	Display()
}

type CurrentConditionsDisplay struct {
	temperature float64
	humidity float64
}

func NewCurrentConditionsDisplay(weatherData *WeatherData) *CurrentConditionsDisplay{
	c := &CurrentConditionsDisplay{}
	weatherData.RegisterObserver(c)
	return c
}

func (c *CurrentConditionsDisplay) Update(temp float64, humidity float64, pressure float64) {
	c.temperature = temp
	c.humidity = humidity
	c.Display()
}

func (c *CurrentConditionsDisplay) Display() {
	fmt.Println("💧 현재 상태: 온도 ", c.temperature, "F, 습도 ", c.humidity, "%")
}

// StatisticsDisplay struct
type StatisticsDisplay struct {
	minTemp, maxTemp, avgTemp float64
	temperatureCount          int
}

// NewStatisticsDisplay creates a new StatisticsDisplay instance
func NewStatisticsDisplay(weatherData *WeatherData) *StatisticsDisplay {
	s := &StatisticsDisplay{}
	weatherData.RegisterObserver(s)
	return s
}

// Update calculates avg, min and max temp and displays the information
func (sd *StatisticsDisplay) Update(temp, humidity, pressure float64) {
	if sd.temperatureCount == 0 {
		sd.minTemp = temp
		sd.maxTemp = temp
		sd.avgTemp = temp
	} else {
		if temp < sd.minTemp {
			sd.minTemp = temp
		}
		if temp > sd.maxTemp {
			sd.maxTemp = temp
		}
		sd.avgTemp = (sd.avgTemp*float64(sd.temperatureCount) + temp) / float64(sd.temperatureCount+1)
	}

	sd.temperatureCount++
	sd.Display()
}

// Display displays the statistics
func (sd *StatisticsDisplay) Display() {
	fmt.Printf("🌡️ Temperature statistics: Min %.2f°F, Max %.2f°F, Average %.2f°F\n", sd.minTemp, sd.maxTemp, sd.avgTemp)
}

type ForecaseDisplay struct{}
func (f ForecaseDisplay) Update(temp, humidity, pressure float64) {
	fmt.Println(temp, humidity, pressure)
}