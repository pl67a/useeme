package main

import (
	"./database"

	"github.com/gin-gonic/gin"
)

// const (
// 	TEST_NAME = "Nick"
// )

func sayHello(c *gin.Context) {
	// var user User
	// GetUser(db, TEST_NAME, user)
	// c.String(200, "Hello %s, Age %d", user.Name, user.Age)
	c.String(200, "Hello world")
}

func setupRouters() *gin.Engine {
	r := gin.Default()
	r.GET("/greeting", sayHello)
	return r
}

func main() {
	db := database.InitDB()

	// db := InitDB()
	// db.AutoMigrate(&User{})
	// CreateUser(db, &User{Name: TEST_NAME, Age: 12})
	r := setupRouters()
	r.Run(":9000")
}
