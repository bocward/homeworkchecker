package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bocward/homeworkchecker/hwprocess"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting app....")
	p := "images/loading.png"
	hwprocess.Process(&p)

	// r := gin.Default()
	// r.POST("/homework", uploadImage)
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func uploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("uploadedFile") //  curl -X POST -F uploadedFile=@<image path> http://localhost:8080/homework
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "File not informed"})
		return
	}
	filename := header.Filename

	dirname := fmt.Sprint("images/")
	newfilename := fmt.Sprint(dirname, filename)

	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		os.Mkdir(dirname, 0700)
	}

	out, err := os.Create(newfilename)
	if err != nil {
		log.Println("Error creating file", err)
		c.JSON(http.StatusNotFound, gin.H{"status": "Error creating file " + filename})
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println("Error saving file", err)
		c.JSON(http.StatusNotFound, gin.H{"status": "Error saving file " + filename})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "File uploaded!"})
}
