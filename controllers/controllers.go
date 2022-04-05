package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ecoding_url(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func UploadImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	//name := c.PostForm("user_id")

	//filename := file.Filename
	header := file.Header
	//	c.SaveUploadedFile(file, "tmp/"+filename) // save file to tmp folder in current directory

	//filename := name + ".png"

	fmt.Println(header)
	c.Redirect(http.StatusMovedPermanently, "localhost:8080/index")
}
