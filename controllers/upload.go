package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadResume(c *gin.Context) {
	file, err := c.FormFile("resume")
	if err != nil {
		fmt.Println("Error retrieving the file:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".pdf" && ext != ".docx" {
		fmt.Println("Invalid file extension:", ext)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only PDF or DOCX files are allowed"})
		return
	}

	err = c.SaveUploadedFile(file, "./uploads/"+file.Filename)
    if err != nil {
        fmt.Println("Error saving the file:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
        return
    }

    fmt.Println("File uploaded successfully:", file.Filename)
    c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})

}
	
