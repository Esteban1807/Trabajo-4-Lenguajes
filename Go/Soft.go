package main
import "fmt"

type Soft struct {
	Tyres
}

func NewSoft(A, B, C, D string, lapsMax int, isWet bool, tyreColor string) Soft {
	tyres := NewTyres(A, B, C, D, lapsMax, isWet, tyreColor)
	return Soft{Tyres: tyres}
}
 
func (s Soft) Durability(laps int) string {
	if laps > s.durabilityMax {
		s.Pits()
		return fmt.Sprintf("The tyres have been on track %d laps\nThe Softs are being changed", laps)
	}
	return fmt.Sprintf("The tyres have been on track %d laps\nCan continue", laps)
}

func (s Soft) PSI() {
	fmt.Println("The air pressure required is 34")
}