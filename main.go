package main
import "bytes"
import "fmt"

func main() {
	a := []byte("go")
    b := []byte("golang")
    fmt.Println(bytes.Compare(a, b)) // -1
    fmt.Println(bytes.Equal(a, b))   // false

}