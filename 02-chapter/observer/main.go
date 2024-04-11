package main

import "fmt"

type Subject interface {
	RegisterObserver(Observer)
	RemoveObserver(Observer)
	NotifyObservers()
}

type Observer interface {
	Update(float64, float64, float64)
}

type DisplayElement interface {
	Display()
}

type WeatherData struct {
	observers []Observer
	temp float64
	humidity float64 // 습도
	pressure float64 // 기압
}

func New() *WeatherData{
	return &WeatherData{
		observers: []Observer{},
	}
}

func (w WeatherData) RegisterObserver(o Observer) {
	w.observers = append(w.observers, o)
	fmt.Println("옵저버 등록", w.observers)
}

func (w WeatherData) RemoveObserver(o Observer) {
	for i, o := range w.observers {
		fmt.Println(i, o)
	}
}

func (w WeatherData) NotifyObservers() {
	fmt.Println("Noti!", w.observers)
	for i := range w.observers {
		w.observers[i].Update(w.temp, w.humidity, w.pressure)
	}
}



func (w WeatherData) MeasurementsChanged() {
	w.NotifyObservers()
}

func (w WeatherData) SetMeasurements(temp, humidity, pressure float64) {
	w.temp = temp
	w.humidity = humidity
	w.pressure = pressure
	w.MeasurementsChanged()	
}

type CurrentConditionsDisplay struct {
	temperature float64
	humidity float64
	weatherData WeatherData
}

func NewCurrentConditionsDisplay(weatherData WeatherData) CurrentConditionsDisplay{
	c := CurrentConditionsDisplay{
		weatherData: weatherData,
	}
	weatherData.RegisterObserver(c)
	return c
}
func (c CurrentConditionsDisplay) Update(temp, humidity, pressure float64) {
	fmt.Println("update!")
	c.temperature = temp
	c.humidity = humidity
	c.Display()
}

func (c CurrentConditionsDisplay) Display() {
	fmt.Println("현재 상태: 온도 ", c.temperature, "F, 습도 ", c.humidity, "%")
}

type StatisticsDisplay struct{}
func (s StatisticsDisplay) Update(temp, humidity, pressure float64) {
	fmt.Println(temp, humidity, pressure)
}

type ForecaseDisplay struct{}
func (f ForecaseDisplay) Update(temp, humidity, pressure float64) {
	fmt.Println(temp, humidity, pressure)
}
func main() {
	weatherData := New()
	
	NewCurrentConditionsDisplay(*weatherData)
	// statisticDisplay := StatisticsDisplay{}
	// forecastDisplay := ForecaseDisplay{}

	weatherData.SetMeasurements(80, 65, 30.4)
}