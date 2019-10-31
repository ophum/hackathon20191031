package main

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*.html")

	r.GET("/", Index)

	r.Run()

}
