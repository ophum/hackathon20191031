package main

import (
	"github.com/ophum/hackathon20191031/message"
	"github.com/gin-gonic/gin"
	"github.com/ophum/hackathon20191031/room"
	"strconv"
	"fmt"
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

func RoomIndex(c *gin.Context)  {
	id := c.Param("id")
	i, _ := strconv.ParseInt(id,10,32)
	room := rooms[i-1]
	c.HTML(200, "room.html", gin.H{"room": room})
}

func SendMessage(c *gin.Context) {
	i := c.PostForm("id")
	id, _ := strconv.ParseInt(i,10,32)
	name := c.PostForm("name")
	mes := c.PostForm("message")
	date := ""
	m := message.NewMessage(int(id), name, mes, date)
	rooms[id-1].Messages = append(rooms[id-1].Messages, *m)
	c.Redirect(301, fmt.Sprintf("/room/%d", id))
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*.html")
	r.Use(sampleMiddleware())

	r.GET("/", Index)
	r.GET("/room/:id", RoomIndex)
	r.POST("/room", CreateRoom)
	r.POST("/send", SendMessage)

	r.Run(":8000")

}
