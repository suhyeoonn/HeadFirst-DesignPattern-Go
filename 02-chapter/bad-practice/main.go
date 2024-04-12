package badpractice

import "fmt"

type WeatherData struct {
	temp float64
	humidity float64 // 습도
	pressure float64 // 기압
}

func (w WeatherData) MeasurementsChanged() {
	temp := w.temp
	humidity := w.humidity
	pressure := w.pressure

	CurrentConditionsDisplay{}.Update(temp, humidity, pressure)
	StatisticsDisplay{}.Update(temp, humidity, pressure)
	ForecaseDisplay{}.Update(temp, humidity, pressure)
}

type CurrentConditionsDisplay struct {}
func (c CurrentConditionsDisplay) Update(temp, humidity, pressure float64) {
	fmt.Println(temp, humidity, pressure)
}

type StatisticsDisplay struct{}
func (s StatisticsDisplay) Update(temp, humidity, pressure float64) {
	fmt.Println(temp, humidity, pressure)
}

type ForecaseDisplay struct{}
func (f ForecaseDisplay) Update(temp, humidity, pressure float64) {
	fmt.Println(temp, humidity, pressure)
}
