package main
import (
	"fmt"
)


func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
}