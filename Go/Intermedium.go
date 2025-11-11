package main
import "fmt"

type Intermedium struct{
	Tyres
}

func NewIntermedium (A, B, C, D string, lapsMax int, isWet bool, tyreColor string) Intermedium{
	tyres := NewTyres(A, B, C, D, lapsMax, isWet, tyreColor)
	return Intermedium{Tyres: tyres}
}

func (i Intermedium) Durability(laps int) string{
	if (laps>i.durabilityMax){
		i.Pits()
		return fmt.Sprintf("The tyres have been on track %d laps\nThe Intermediums are being changed", laps)
	}else{
		return fmt.Sprintf("The tyre have been on track %d laps\nCan continue", laps)
	}
}

func (i Intermedium) PSI(){
	fmt.Println("The air pressure required is 37")
}

func (i Intermedium) AreWet(){
	fmt.Println("The wheather is medium")
}