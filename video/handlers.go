package video

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func hello(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		"hello world",
	)
}

func saveVideo(c *gin.Context) {
	_, err := c.FormFile("video")
	if err != nil {
		log.Fatal("err:", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"err": err,
			},
		)
	}
	id := c.PostForm("id")
	log.Println("id:", id)
	c.JSON(
		http.StatusOK,
		gin.H{
			"id": id,
		},
	)
}

func endPoint(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("client connnected")
	for {
		msgType, data, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("type:", msgType)
		log.Println("data:", len(data))
		ws.WriteMessage(msgType, data)
	}
}
