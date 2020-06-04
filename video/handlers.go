package video

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
