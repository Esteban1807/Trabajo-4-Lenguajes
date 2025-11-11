package main
import "fmt"

type Tyres struct{
	Car
	durabilityMax int
	wet bool
	color string
}

func NewTyres(A, B, C, D string, lapsMax int, isWet bool, tyreColor string) Tyres {
	car := Car{
		rightFront: A,
		leftFront:  B,
		rightRear:  C,
		leftRear:   D,
	}
	return Tyres{
		Car:           car,
		durabilityMax: lapsMax,
		wet:           isWet,
		color:         tyreColor,
	}
}

func (t Tyres) Durability(laps int) string {
	return fmt.Sprintf("The tyres have been on track %d laps", laps)
}

func (t Tyres) ShowAll(){
	fmt.Printf("The tyres are %s %s %s %s\nTheir durability is %d laps\nThe tyres are wet? %t\nTyres color: %s\n",
		t.rightFront, t.leftFront, t.rightRear, t.leftRear, t.durabilityMax, t.wet, t.color)
}

func (t Tyres) PSI(){}