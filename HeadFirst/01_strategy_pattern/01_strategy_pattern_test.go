package strategy

import "testing"


func TestStrategy(t *testing.T) {
		// 1
		duck := MallardDuck{}
		duck.flyBehavior = FlyWithWings{}
		duck.quackBehavior = Quack{}
	
		duck.display()
		duck.performQuack()
		duck.performFly()
	
		// 2
		duck2 := ModelDuck{}
		duck2.flyBehavior = FlyNoWay{}
		duck2.quackBehavior = MuteQuack{}
	
		duck2.display()
		duck2.performFly()
		duck2.setFlyBehavior(FlyRocketPowered{})
		duck2.performFly()
}
