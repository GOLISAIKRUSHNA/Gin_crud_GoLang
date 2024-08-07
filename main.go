package main

import (
	"io"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
																																																																																																																								
	router.GET("/getdat																																																																																																																																																																																																																																																																																																																																																						a", getdata)
	router.GET("/getquerystring", getquerystring)
	router.GET("/geturl/:name/:age", geturl)
	router.POST("/postdata", postdata)
	router.PUT("/putdata/:id", putdata)
	router.DELETE("/deletedata/:id", deletedata)

	router.Run() //:5000

}

func deletedata(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"message": "Data deleted successfully",
		"id":      id,
	})
}

func putdata(c *gin.Context) {

	id := c.Param("id")

	body := c.Request.Body
	value, _ := io.ReadAll(body)

	c.JSON(200, gin.H{
		"data": "update data from putdata function",
		"id":   id,
		"body": string(value),
	})

}

// http://localhost:5000/geturl/Saikrushna/22
func geturl(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"data": "This geturl data with saikrushna and 22",
		"name": name,
		"age":  age,
	})
}

// http://localhost:5000/getquerystring?name=saibhai&age=22
func getquerystring(c *gin.Context) {
	namee := c.Query("name")
	agee := c.Query("age")
	c.JSON(200, gin.H{
		"data": "Hi i am gin framework for query string",
		"name": namee,
		"age":  agee,
	})

}

func getdata(c *gin.Context) {
	c.JSON(200, gin.H{
		"data":    "apna college",
		"Id":      1,
		"Subject": "youtuber and getall",
	})
}

func postdata(c *gin.Context) {
	body := c.Request.Body
	value, _ := io.ReadAll(body)

	c.JSON(200, gin.H{
		"data":    "coding wallah sir for java of post",
		"Id":      2,
		"Subject": "This is post method",
		"body":    string(value),
	})
}
