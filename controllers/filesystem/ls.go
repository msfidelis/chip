package filesystem

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type RequestLs struct {
	Path string `json:"path"`
}

type ResponseLs struct {
	Path  string   `json:"path"`
	Files []string `json:"files"`
}

func Ls(c *gin.Context) {

	var request RequestLs
	var response ResponseLs

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files, err := os.ReadDir(request.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	fmt.Println(files)

	response.Path = request.Path
	response.Files = fileNames

	c.JSON(200, response)
}
