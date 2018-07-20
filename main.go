package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"os/exec"
)

type rq struct {
	Address string `json:"address"`
}


func main() {
	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {
		var json rq
		if err := c.BindJSON(&json);
		err != nil {
			fmt.Println("Error: ", err)
		}
		// b := exec.Command("cmd", "/C", "start", "ping","www.baidu.com")
		b := exec.Command("/bin/bash", "-c", "youtube-dl "+json.Address)
		if err := b.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
		c.JSON(http.StatusOK, json.Address)
	})

	router.Run(":8080")
}
