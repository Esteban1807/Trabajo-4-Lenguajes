package main
import "fmt"

type Car struct {
	rightFront string
	leftFront string
	rightRear string
	leftRear string
}

func (c *Car) NewCar (A string, B string, C string, D string){
	c.rightFront=A
	c.leftFront=B
	c.rightRear=C
	c.leftRear=D
}

func (c Car) Pits(){
	fmt.Println("The car is entering to pits and the tyres will be change")
}