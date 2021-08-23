package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	dialogflow "github.com/lucasansei/pythia/dialogflow/detect_intent"
)

func main() {
	router := gin.Default()

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	router.GET("/text", func(c *gin.Context) {
		text := c.PostForm("text")

		res, err := dialogflow.DetectIntentText("experimentation-lab-60c88", "123456", text, "pt-BR")
		fmt.Println(res)
		fmt.Println(err)

		c.JSON(200, gin.H{
			"dialogflow_result": res,
		})
	})

	router.POST("/audio", func(c *gin.Context) {
		file, _ := c.FormFile("audio")
		fmt.Println(file.Filename)

		filename := "./tmp-audio-folder/" + file.Filename

		c.SaveUploadedFile(file, filename)
		res, err := dialogflow.DetectIntentAudio("experimentation-lab-60c88", "123456", filename, "pt-BR")
		fmt.Println(err)

		e := os.Remove(filename)
		if e != nil {
			fmt.Println(e)
		}

		c.JSON(200, gin.H{
			"dialogflow_result": res,
		})
	})

	router.Run()
}
