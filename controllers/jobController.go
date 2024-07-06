package controllers

import (
	"Go_authentication/models"
	"Go_authentication/services"
	"Go_authentication/utils"
	"fmt"
	"net/http"

	//"Go_authentication/config"
	//"Go_authentication/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type JobRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CompanyName string `json:"company_name"`
}

func CreateJob(c *gin.Context) {
	var jobRequest JobRequest
	if err := c.ShouldBindJSON(&jobRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response("Invalid request", err.Error()))
		return
	}

	//userID, _ := strconv.Atoi(c.GetString("user_id"))
	userIDStr, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, utils.Response("User ID not found in context", nil))
		return
	}
	userIDs, ok := userIDStr.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Response("interface error", nil))
		return
	}
	userIdString := fmt.Sprintf("%v", userIDs)
	fmt.Println("Value", userIdString)
	userID, err := strconv.Atoi(userIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response("Invalid user ID", err.Error()))
		return
	}
	job := models.Job{
		Title:       jobRequest.Title,
		Description: jobRequest.Description,
		CompanyName: jobRequest.CompanyName,
		PostedByID:  uint(userID),
		PostedOn:    time.Now(),
	}

	if err := services.CreateJob(&job); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response("Database error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Response("Job created successfully", nil))
}

func ListJobs(c *gin.Context) {
	jobs, err := services.ListJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response("Database error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Response("Jobs fetched successfully", jobs))
}

func ApplyToJob(c *gin.Context) {
	fmt.Println("Entering ApplyToJob controller")
	// Get job_id from query parameters
	jobIDStr := c.Query("job_id")
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		fmt.Println("Invalid job_id:", err)
		c.JSON(http.StatusBadRequest, utils.Response("Invalid job_id", err.Error()))
		return
	}
	// Get user_id from context
	userIDStr, exists := c.Get("user_id")
	if !exists {
		fmt.Println("User ID not found in context")
		c.JSON(http.StatusUnauthorized, utils.Response("User ID not found in context", nil))
		return
	}

	userID, err := strconv.Atoi(userIDStr.(string))
	if err != nil {
		fmt.Println("Invalid user_id:", err)
		c.JSON(http.StatusInternalServerError, utils.Response("Invalid user_id", err.Error()))
		return
	}

	// Call the service to apply for the job
	if err := services.ApplyToJob(uint(jobID), uint(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response("Database error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Response("Applied to job successfully", nil))
}
