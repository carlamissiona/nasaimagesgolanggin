package main

import ( 
	
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var Router * gin.Engine

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func StaticHandler(c *gin.Context) {
	filepath := c.Param("filepath")
	data := Assets.Files["templates/assets"+filepath].Data
	c.Writer.Write(data)
}

func main() {
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}

 	r := gin.Default()
        r.SetHTMLTemplate(t)
	r.GET("assets/*filepath", StaticHandler)
// 	r.LoadHTMLGlob("templates/*")
// 	r.LoadHTMLGlob("templates/assets/*")
// 	r.LoadHTMLGlob("templates/assets/css/*")
// 	r.LoadHTMLGlob("templates/assets/css/js/*")
// 	r.LoadHTMLGlob("templates/assets/sass/*")
// 	r.LoadHTMLGlob("templates/assets/webfonts/*")
// 	r.LoadHTMLGlob("templates/images/*")
// //         r.Static("/assets", "./templates/assets")
// // 	r.Static("/assets/css", "./templates/assets/css")
// // 	r.Static("/images", "./templates/images")
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
