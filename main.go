package main
import "bytes"
import "fmt"

func main() {
	buf := bytes.NewBuffer([]byte("Go"))
	b := make([]byte, 2)
	n, _ := buf.Read(b)
	fmt.Println(string(b[:n])) // "Go"


}