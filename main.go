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
type Hero struct {
    Id     int `json:"id"`
    Name   string `json:"name"`
}

func main() {
	router := gin.Default()
	// router.LoadHTMLGlob("/")
	router.LoadHTMLFiles("index.html")
	router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl",gin.H{
            "title": "index",
        })
    })

	router.POST("/api/post", func(c *gin.Context) {
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

	router.GET("/api/heroes", func(c *gin.Context) {
		hero := []Hero{
			// Id:     1,
			// Name:  "name",
			{1,"t1"},
			{2,"t2"},
			{3,"t3"},
			{4,"t4"},
		}
		// name := c.Param("name")
		//c.String(http.StatusOK, "Hello %s", name)
		c.JSON(http.StatusOK, hero)
	})


	router.Run(":8080")
}
