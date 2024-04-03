package main
import "bytes"
import "fmt"

func main() {
	buf := bytes.NewBuffer([]byte("Go"))
	b := make([]byte, 2)
	n, _ := buf.Read(b)
	fmt.Println(string(b[:n])) // "Go"


}

func processLogData(logData []byte) {
    lines := bytes.Split(logData, []byte("\n"))
    for _, line := range lines {
        if bytes.Contains(line, []byte("ERROR")) {
            fmt.Println("Error found:", string(line))
        }
    }
}
