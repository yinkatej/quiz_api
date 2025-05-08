package main

import (
	"fmt"
	"net/http"
	"test/quiz"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	// Load the quiz from the JSON file
	loadedQuiz, err := quiz.LoadQuiz()
	if err != nil {
		fmt.Println("Error loading quiz:", err)
		return
	}

	// Print the loaded quiz
	// fmt.Printf("Loaded Quiz: %+v\n", loadedQuiz.Quiz)

	// Example of accessing a question
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	
	r.Use(cors.New(config))
	r.GET("/quiz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"quiz": loadedQuiz.Quiz})
	})
	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Quiz API!"})
	})
	r.StaticFile("/image", "./image/ict1.jpg")
	r.StaticFile("/image2", "./image/ict2.jpg")
	r.StaticFile("/image3", "./image/ict3.jpg")
	
	r.Run(":8080")
	fmt.Println("Server running on port 8080")
}