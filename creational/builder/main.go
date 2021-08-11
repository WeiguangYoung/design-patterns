package main

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelsWheels        = "steel"
)

type Builder interface {
	Paint(Color) Builder
	Install(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	Driver() error
	Stop() error
}

func NewBuilder() *Car {
	builder := new(Car)
	return builder
}

type Car struct {
	Color  Color
	Wheels Wheels
	Speed  Speed
}

func (car *Car) Paint(color Color) *Car {
	car.Color = color
	return car
}
func (car *Car) Install(wheels Wheels) *Car {
	car.Wheels = wheels
	return car
}
func (car *Car) TopSpeed(speed Speed) *Car {
	car.Speed = speed
	return car
}

func (car *Car) Build() *Car {
	return car
}

func (car *Car) Driver() {
	fmt.Println("driver")
}

func main() {
	assembly := NewBuilder().Paint(RedColor)
	familyCar := assembly.Install(SportsWheels).TopSpeed(50 * MPH).Build()
	familyCar.Driver()
}
