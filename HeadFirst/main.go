package main

import "fmt"

type FlyBehavior interface {
	fly()
}

type FlyWithWings struct{}

func (f FlyWithWings) fly() {
	fmt.Println("I'm flying!")
}

type FlyNoWay struct{}

func (f FlyNoWay) fly() {
	fmt.Println("I can't fly!")
}

type FlyRocketPowered struct{}

func (f FlyRocketPowered) fly() {
	fmt.Println("I'm flying with rocket!")
}

type QuackBehavior interface {
	quack()
}

type Quack struct{}

func (q Quack) quack() {
	fmt.Println("Quack")
}

type MuteQuack struct{}

func (q MuteQuack) quack() {
	fmt.Println("<< Silence >>")
}

type Squeak struct{}

func (q Squeak) quack() {
	fmt.Println("Squeak")
}

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d Duck) display() {}

func (d Duck) performFly() {
	d.flyBehavior.fly()
}

func (d Duck) performQuack() {
	d.quackBehavior.quack()
}

func (d Duck) swim() {
	fmt.Println("All ducks float, even decoys!")
}

func (d *Duck) setFlyBehavior(fb FlyBehavior) {
	d.flyBehavior = fb
}
func (d *Duck) setQuackBehavior(qb QuackBehavior) {
	d.quackBehavior = qb
}

type MallardDuck struct {
	Duck
}

func (d MallardDuck) display() {
	fmt.Println("I'm a real Mallard duck")
}

type ModelDuck struct {
	Duck
}

func (d ModelDuck) display() {
	fmt.Println("I'm a model duck")
}

func main() {
	// 1
	duck := MallardDuck{}
	duck.flyBehavior = FlyWithWings{}
	duck.quackBehavior = Quack{}

	duck.performQuack()
	duck.performFly()

	// 2
	duck2 := ModelDuck{}
	duck2.flyBehavior = FlyNoWay{}
	duck2.quackBehavior = MuteQuack{}
	duck2.performFly()
	duck2.setFlyBehavior(FlyRocketPowered{})
	duck2.performFly()
}
