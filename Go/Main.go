package main
import "fmt"

func main(){
	soft := NewSoft("SoftA", "SoftB", "SoftC", "SoftD", 15, false, "Red")
	soft.ShowAll()
	soft.PSI()
	fmt.Println("____________________________")
	fmt.Println(soft.Durability(16))
	fmt.Printf("__________________\n")

	medium := NewMedium("MediumA", "MediumB", "MediumC", "MediumD", 30, false, "Yellow")
	medium.ShowAll()
	medium.PSI()
	fmt.Println("____________________________")
	fmt.Println(medium.Durability(15))
	fmt.Printf("__________________\n")

	hard := NewHard("HardA", "HardB", "HardC", "HardD", 40, false, "White")
	hard.ShowAll()
	hard.PSI()
	fmt.Println("____________________________")
	fmt.Println(hard.Durability(46))
	fmt.Printf("__________________\n")

	full := NewFullWet("FullA", "FullB", "FullC", "FullD", 25, true, "Blue")
	full.ShowAll()
	full.PSI()
	fmt.Println("____________________________")
	fmt.Println(full.Durability(23))
	fmt.Printf("__________________\n")

	inter := NewIntermedium("InterA", "InterB", "InterC", "InterD", 40, true, "Green")
	inter.ShowAll()
	inter.PSI()
	fmt.Println("____________________________")
	fmt.Println(inter.Durability(41))
	fmt.Printf("__________________\n")
}