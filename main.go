package main
 
import "github.com/gin-gonic/gin"  //go get github.com/gin-gonic/gin
import "go_l/mypackage"


func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	println("trans.Pi", trans.Pi)

	r := gin.Default()
	r.GET("/ping", ping)
	r.Run()	// listen and serve on 0.0.0.0:8080
}
