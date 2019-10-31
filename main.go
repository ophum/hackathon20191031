package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ophum/hackathon20191031/room"
)

var rooms []room.Room

func sampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("rooms", rooms)
		c.Next()
	}
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"rooms": rooms})
}

func CreateRoom(c *gin.Context) {
	name := c.PostForm("name")
	id := len(rooms) + 1

	room := room.NewRoom(id, name)
	rooms = append(rooms, *room)

	c.Redirect(301, "/")
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*.html")
	r.Use(sampleMiddleware())

	r.GET("/", Index)

	r.POST("/room", CreateRoom)

	r.Run()

}
