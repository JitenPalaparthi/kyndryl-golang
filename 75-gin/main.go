package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"demo/models"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
)

func main() {

	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8085"
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/person/add", Add)

	r.Run(":" + PORT) // listen and serve on 0.0.0.0:8080
}

func Add(c *gin.Context) {

	p1 := new(models.Person)

	if err := c.Bind(p1); err != nil {
		log.Println("Error:", err)
		c.String(http.StatusBadRequest, "Invalid person data")
		c.Abort()
		return
	}

	if !p1.Validate() {
		c.String(http.StatusBadRequest, "invalid data")
		c.Abort()
		return
	}
	p1.LastUpdated = time.Now().Unix()

	str, err := p1.ToString()
	if err != nil {
		log.Println("Error:", err)
		c.String(http.StatusBadRequest, "Invalid person data")
		c.Abort()
		return
	}
	err = SaveToFile("persons.txt", str)
	if err != nil {
		log.Println("Error:", err)
		c.String(http.StatusBadRequest, "Invalid person data")
		c.Abort()
		return
	}

	c.String(201, "Person successfuly added")

}

func SaveToFile(filename string, data string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(data + "\n"))
	if err != nil {
		return err
	}
	return nil
}
