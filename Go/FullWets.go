package main
import "fmt"

type FullWet struct{
	Tyres
}

func NewFullWet (A, B, C, D string, lapsMax int, isWet bool, tyreColor string) FullWet{
	tyres := NewTyres(A, B, C, D, lapsMax, isWet, tyreColor)
	return FullWet{Tyres: tyres}
}

func (f FullWet) Durability(laps int) string{
	if (laps>f.durabilityMax){
		f.Pits()
		return fmt.Sprintf("The tyres have been on track %d laps\nThe FullWets are being changed", laps)
	}else{
		return fmt.Sprintf("The tyre have been on track %d laps\nCan continue", laps)
	}
}

func (f FullWet) PSI(){
	fmt.Println("The air pressure required is 33")
}

func (f FullWet) AreWet(){
	fmt.Println("The wheather is heavy")
}