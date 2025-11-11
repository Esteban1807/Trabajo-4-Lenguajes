package main
import "fmt"

type Medium struct{
	Tyres
}

func NewMedium (A, B, C, D string, lapsMax int, isWet bool, tyreColor string) Medium{
	tyres := NewTyres(A, B, C, D, lapsMax, isWet, tyreColor)
	return Medium{Tyres: tyres}
}

func (m Medium) Durability(laps int) string {
	if (laps > m.durabilityMax){
		m.Pits()
		return fmt.Sprintf("The tyres have been on track %d laps\nThe Mediums are being changed", laps)
	} else{
		return fmt.Sprintf("The tyres have been on track %d laps\nCan continue", laps)
	}
}

func (m Medium) PSI(){
	fmt.Println("The air pressure required is 40")
}