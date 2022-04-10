package controllers

import (
	"net/http"
	"shorten/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

var SavePath string

func Ecoding_url(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func UploadImage(c *gin.Context) {

	file, _ := c.FormFile("file")
	//name := c.PostForm("user_id")

	//filename := file.Filename
	//header := file.Header

	SavePath = "./image/" + time.Now().Format("20060102150405") + file.Filename
	setPath(&SavePath)
	err := c.SaveUploadedFile(file, SavePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "上傳檔案出錯：" + err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "上傳成功\n"})
	}
	//fmt.Println(SavePath)

	//	c.SaveUploadedFile(file, "tmp/"+filename) // save file to tmp folder in current directory

	//filename := name + ".png"

	//fmt.Println(header)

	Pass_data(c, SavePath)
	//c.Writer.WriteString(string(data))
}

func Pass_data(c *gin.Context, SavePath string) {

	f := middlewares.Dump_to_python(SavePath)
	c.String(http.StatusOK, "您輸入的文字為: \n%s", f)
}

func setPath(s *string) string {
	Path := *s
	return Path
}
