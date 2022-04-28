package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

var Router * gin.Engine



func main() {
	

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.LoadHTMLGlob("templates/assets/*")
	r.LoadHTMLGlob("templates/assets/css/*")
	r.LoadHTMLGlob("templates/assets/css/js/*")
	r.LoadHTMLGlob("templates/assets/sass/*")
	r.LoadHTMLGlob("templates/assets/webfonts/*")
	r.LoadHTMLGlob("templates/images/*")
//         r.Static("/assets", "./templates/assets")
// 	r.Static("/assets/css", "./templates/assets/css")
// 	r.Static("/images", "./templates/images")
	r.GET("/updatesemail", func(c *gin.Context) {
		sendemail("codetuna@protonmail.com")
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	
	r.GET("/app", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
			"payload": "hi",
		})
	})
	 
	r.Run()

 


	// Handle the index route
	
	//router.GET("/images/view/:id", showEachItem)
	

	// Start serving the application
}
