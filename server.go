package main

import (
	"server/controllers"
	"server/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(JSONMiddleware())

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// r.GET("/", HomePage)
	// r.POST("/", PostHomePage)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.GET("/query", QueryString)
	// r.GET("/path/:name/:age", PathParameter)

	models.ConnectDatabase()
	r.Run(":4444")
}

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

// func HomePage(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "메인 페이지 입니다.",
// 	})
// }

// func PostHomePage(c *gin.Context) {
// 	body := c.Request.Body
// 	value, err := ioutil.ReadAll(body)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	// fmt.Println(string(value))

// 	var data map[string]interface{}
// 	json.Unmarshal([]byte(value), &data)
// 	c.JSON(http.StatusOK, gin.H{
// 		"name": data["name"],
// 		"age":  data["age"],
// 	})

// 	doc, _ := json.Marshal(data)
// 	// fmt.Println(string(doc))
// 	c.String(http.StatusOK, string(doc))
// }

// func QueryString(c *gin.Context) {
// 	name := c.Query("name")
// 	age := c.Query("age")

// 	c.JSON(http.StatusOK, gin.H{
// 		"name": name,
// 		"age":  age,
// 	})
// }

// func PathParameter(c *gin.Context) {
// 	name := c.Param("name")
// 	age := c.Param("age")

// 	c.JSON(http.StatusOK, gin.H{
// 		"name": name,
// 		"age":  age,
// 	})
// }
