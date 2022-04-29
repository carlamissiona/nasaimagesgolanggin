package main

import ( 
			 
	"net/http"
	 
	"github.com/gin-gonic/gin"


	"fmt"
  	 
    _"github.com/PuerkitoBio/goquery"
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
		
	r.GET("/api/available", func(c *gin.Context) {
		
		resp, err := http.Get("https://epic.gsfc.nasa.gov/api/natural/available")
		if err != nil {
			// handle err
	    }
		defer resp.Body.Close()


		fmt.Println("resp" , resp)
		fmt.Println(resp)

		daily := resp 

		c.JSON(
		http.StatusOK,
		daily,
	)
		 
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
