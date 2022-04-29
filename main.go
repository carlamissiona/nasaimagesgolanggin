package main

import ( 
			 
	"net/http"
	 
	"github.com/gin-gonic/gin"
)

var Router * gin.Engine

func main() {
 	
	r := gin.Default()
      
    r.StaticFS("/searches", http.Dir("searches")) 

 	r.Static("/assets", "./assets")   
 	r.Static("/images", "./images")   
	

	r.LoadHTMLGlob("templates/*.html")
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
		
	r.GET("/search", func(c *gin.Context) {
		c.HTML(http.StatusOK, "search.html", gin.H{
			"title": "Search",
			"payload": "hi",
		})
	})
	  

	r.Run()

 


	// Handle the index route
	
	//router.GET("/images/view/:id", showEachItem)
	

	// Start serving the application
}
