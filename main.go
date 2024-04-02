package main
import (
	"fmt"
)


func main() {
	a  := [5]int{1,2,3,4,5}
	for i, _ :=  range a{
		fmt.Println("i", i)
	}
}