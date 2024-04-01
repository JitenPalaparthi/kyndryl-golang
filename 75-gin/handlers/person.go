package handlers

import (
	"demo/interfaces"
	"demo/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	FileName string
	IPerson  interfaces.IPerson
}

func (p *Person) Add(c *gin.Context) {
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

	result, err := p.IPerson.Add(p1)
	if err != nil {
		log.Println("Error:", err)
		c.String(http.StatusBadRequest, "Invalid person data")
		c.Abort()
		return
	}

	c.String(201, fmt.Sprint(result.(int)))

}

func (p *Person) Delete(c *gin.Context) {
	fmt.Println("-----<")
	id := c.Param("id")
	fmt.Println("-----<", id)
	if id == "" {
		log.Println("Error:", "invalid id")
		c.String(http.StatusBadRequest, "Invalid id")
		c.Abort()
	}
	result, err := p.IPerson.Delete(id)
	if err != nil {
		log.Println("Error:", err)
		c.String(http.StatusBadRequest, "Invalid data")
		c.Abort()
		return
	}
	c.String(http.StatusAccepted, fmt.Sprint(result.(int64)))
}
