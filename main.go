package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

var Router * gin.Engine

func showHome(c *gin.Context) {
	 

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": "hi",
		},
	)

}

func showEachItem(c *gin.Context) {
}






func main() {

	r := gin.Default()
	Router.LoadHTMLGlob("templates/*")

	Router.GET("/updatesemail", func(c *gin.Context) {
		sendemail("codetuna@protonmail.com")
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	
	Router.GET("/", showHome)
	r.Run()

 


	// Handle the index route
	
	//router.GET("/images/view/:id", showEachItem)
	

	// Start serving the application
}
