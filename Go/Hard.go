package main
import "fmt"

type Hard struct{
	Tyres
}

func NewHard (A, B, C, D string, lapsMax int, isWet bool, tyreColor string) Hard{
	tyres := NewTyres(A, B, C, D, lapsMax, isWet, tyreColor)
	return Hard{Tyres:tyres}
}

func (h Hard) Durability(laps int) string{
	if (laps>h.durabilityMax){
		h.Pits()
		return fmt.Sprintf("The tyres have been on track %d laps\nThe Hards are being changed", laps)
	}else{
		return fmt.Sprintf("The tyre have been on track %d laps\nCan continue", laps)
	}
}

func (h Hard) PSI(){
	fmt.Println("The air pressure required is 51")
}