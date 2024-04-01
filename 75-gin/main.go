package main

import (
	"demo/database"
	"demo/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
	DSN  string
)

func main() {

	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8085"
	}
	DSN = os.Getenv("DB_CONNECTION")
	if DSN == "" {
		DSN = "host=localhost user=app password=app@123 dbname=demodb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	db, err := database.GetConnection(DSN)
	if err != nil {
		log.Fatal(err)
	}

	personDb := database.Person{DB: db}
	phandler := new(handlers.Person)
	phandler.IPerson = &personDb
	phandler.FileName = "persons.txt"
	r.POST("person/add", phandler.Add)
	r.DELETE("person/:id", phandler.Delete)
	r.Run(":" + PORT) // listen and serve on 0.0.0.0:8080
}

// write Add to a handler package
// write SaveToFile in helper pacakge
