package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"os/exec"
	//"io/ioutil"
	"path/filepath"
	"os"
)

type rq struct {
	Address string `json:"address"`
}
type Hero struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Path struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

func main() {
	router := gin.Default()
	// router.LoadHTMLGlob("/")
// 	router.LoadHTMLFiles("index.html")
// 	router.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.tmpl", gin.H{
// 			"title": "index",
// 		})
// 	})

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

	router.GET("/api/get", func(c *gin.Context) {
		var result []Path

		//获取路径
		path := `/home`
		err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				return nil
			}

			//println(filepath.Base(path))
			re := Path{
				Path: path,
				Name: filepath.Base(path),
			}
			result = append(result, re)
			return nil
		})
		if err != nil {
			fmt.Println("Error: ", err)
		}

		//获取文件名
		//files, _ := ioutil.ReadDir("./")
		//for _, f := range files {
		//	fmt.Println(f)
		//	re := Path{
		//		Path: f.Name(),
		//	}
		//	result = append(result, re)
		//}

		//获取文件名
		//files2, _ := filepath.Glob("*")
		//fmt.Println(files2)

		//obj :=	"filepath.Walk() returned %v\n"
		c.JSON(http.StatusOK, result)

	})

	router.GET("/api/heroes", func(c *gin.Context) {
		hero := []Hero{
			// Id:     1,
			// Name:  "name",
			{1, "t1"},
			{2, "t2"},
			{3, "t3"},
			{4, "t4"},
		}
		// name := c.Param("name")
		//c.String(http.StatusOK, "Hello %s", name)
		c.JSON(http.StatusOK, hero)
	})

	router.Run(":8080")
}
