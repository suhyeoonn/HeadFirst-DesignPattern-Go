package main

type Subject interface {
	RegisterObserver(Observer)
	RemoveObserver(Observer)
	NotifyObservers()
}

type WeatherData struct {
	observers []Observer
	temp float64
	humidity float64 // 습도
	pressure float64 // 기압
}

func NewWeatherData() *WeatherData{
	return &WeatherData{}
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.temp = 10
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
	// TODO
}

func (w *WeatherData) NotifyObservers() {
	for _, observer := range w.observers {		
		observer.Update(w.temp, w.humidity, w.pressure)
	}
}

func (w *WeatherData) MeasurementsChanged() {
	w.NotifyObservers()
}

func (w *WeatherData) SetMeasurements(temp, humidity, pressure float64) {
	w.temp = temp
	w.humidity = humidity
	w.pressure = pressure
	w.MeasurementsChanged()	
}