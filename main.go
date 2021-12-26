package main

import (
	"memberclub/store"
	"time"

	"github.com/gin-gonic/gin"
)

var inMemoryStore = store.InMemoryMemberClubStore{}

func main() {
	fillTestMembers()

	r := gin.Default()
	r.GET("/", rootRequest)
	r.Run()
}

func rootRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"members": inMemoryStore.GetMembers(),
	})
}

func fillTestMembers() {
	inMemoryStore.AddMember(store.Member{
		Name:      "John Show",
		Email:     "j.show@company.com",
		DateAdded: time.Now(),
	})
	inMemoryStore.AddMember(store.Member{
		Name:      "Dana A. Watkins",
		Email:     "Dana2000@mail.com",
		DateAdded: time.Now(),
	})
	inMemoryStore.AddMember(store.Member{
		Name:      "Henri Rousseau",
		Email:     "rousseau@impressionis.fr",
		DateAdded: time.Now(),
	})
	inMemoryStore.AddMember(store.Member{
		Name:      "Mr. Brown",
		Email:     "brown@gmail.com",
		DateAdded: time.Now(),
	})
}
