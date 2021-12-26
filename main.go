package main

import (
	"memberclub/store"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var inMemoryStore = store.InMemoryMemberClubStore{}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
	r.GET("/members", getMembers)
	r.POST("/member", addMember)
	r.Run("localhost:8080")
}

func getMembers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, inMemoryStore.GetMembers())
}

func addMember(c *gin.Context) {
	var newMember store.Member

	if err := c.BindJSON(&newMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong when creating a new member.",
		})
		return
	}

	if err := inMemoryStore.AddMember(newMember); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, inMemoryStore.GetMembers())
}
