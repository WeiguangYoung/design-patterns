package main

import (
	"log"
)

type Subject interface {
	registerObserver(Observer)
	removeObserver(Observer)
	notifyObervers()
}
type Observer interface {
	getType() string
	update(temp float64, humidity float64, pressure float64)
}

type DisplayElement interface {
	display()
}

type WeatherData struct {
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func NewWeatherData() WeatherData {
	return WeatherData{
		observers: []Observer{},
	}
}

func (w *WeatherData) registerObserver(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) removeObserver(o Observer) {
	index := -1
	for i, thisO := range w.observers {
		if thisO.getType() == o.getType() {
			index = i
			break
		}
	}

	if index >= 0 {
		w.observers = append(w.observers[:index], w.observers[index:]...)
	}
}

func (w *WeatherData) notifyObervers() {
	for _, o := range w.observers {
		o.update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *WeatherData) measurementsChanged() {
	w.notifyObervers()
}

func (w *WeatherData) setMeasurements(temp float64, humidity float64, pressure float64) {
	w.temperature = temp
	w.humidity = humidity
	w.pressure = pressure

	w.measurementsChanged()
}

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
	weatherData Subject
}

func NewCurrentConditionsDisplay(w Subject) CurrentConditionsDisplay {
	c := &CurrentConditionsDisplay{
		weatherData: w,
	}
	w.registerObserver(c)
	return *c
}
func (c *CurrentConditionsDisplay) getType() string {
	return "CurrentConditionsDisplay"
}
func (c *CurrentConditionsDisplay) update(temp float64, humidity float64, pressure float64) {
	c.temperature = temp
	c.humidity = humidity

	c.display()
}

func (c *CurrentConditionsDisplay) display() {
	log.Printf("Current conditions: %.2f degrees and %.2f%% humidity", c.temperature, c.humidity)
}

func main() {
	weatherData := NewWeatherData()
	NewCurrentConditionsDisplay(&weatherData)

	weatherData.setMeasurements(80, 65, 30.4)
	weatherData.setMeasurements(82, 70, 29.2)
	weatherData.setMeasurements(78, 90, 29.2)
}
