package main

import "fmt"

type Car struct {
	brand string
	model string
	year int
	km int
	fuel float32
	Engine
}

type Engine struct {
	name string
	isElectrified bool
	horsePower float32
	torque float32
	volume float32
	averageFuelCons float32
	isRunning bool
}

func (c *Car) start() bool {
	if c.fuel > 0.1 && !c.isRunning {
		c.fuel -= 0.1
		c.isRunning = true
		return true
	}
	return false
}

func (c *Car) stop() bool {
	if c.isRunning {
		c.isRunning = false
		return true
	}
	return false
}

func calculateFuel(distance int, c Car) float32 {
	if !c.isElectrified {
		fuelNeeded := c.averageFuelCons * float32(distance)/100
		return fuelNeeded
	} else {
		return 0
	}
}

func calculateDistanceFromFuel(fuel float32, avgCons float32) int {
	return int(fuel*100/avgCons)
}

func (c *Car) drive(driveDistance int)  {
	if c.isRunning {
		totalFuelNeeded := calculateFuel(driveDistance, *c)
		if totalFuelNeeded < c.fuel {
			c.km += driveDistance
			c.fuel -= totalFuelNeeded
		} else {
			c.km += calculateDistanceFromFuel(c.fuel, c.averageFuelCons)
			c.fuel = 0
			c.isRunning = false
		}
	}
}



func main() {
	tsi := Engine{"TSI", false, 130, 150, 1.5, 6.8, false}
	polo := Car{"VW", "Polo", 2019, 10000, 44.0, tsi}

	if polo.start() {
		polo.drive(300)
	}
	polo.stop()
	polo.drive(20)

	fmt.Println(polo)
	fmt.Println(tsi)
}
