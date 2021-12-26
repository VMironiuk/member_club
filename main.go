package main

import (
	"memberclub/store"
	"time"

	"github.com/gin-gonic/gin"
)

var inMemoryStore = store.InMemoryMemberClubStore{}

func main() {
	r := gin.Default()
	inMemoryStore.AddMember(store.Member{
		Name:      "John Smith",
		Email:     "john.smith@gmail.com",
		DateAdded: time.Now(),
	})
	r.GET("/", rootRequest)
	r.Run()
}

func rootRequest(c *gin.Context) {
	member := inMemoryStore.GetMembers()[0]
	c.JSON(200, gin.H{
		"name":      member.Name,
		"email":     member.Email,
		"dateAdded": member.DateAdded,
	})
}
