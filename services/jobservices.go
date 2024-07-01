package services

import (
	"Go_authentication/config"
	"Go_authentication/models"
	"fmt"
)

// CreateJob creates a new job record in the database.
func CreateJob(job *models.Job) error {
	// Log the incoming job creation request
	fmt.Printf("Creating job: %+v\n", job)

	// Perform the database operation to create a new job record
	result := config.DB.Create(job)

	// Check for errors during the database operation
	if result.Error != nil {
		// Log the error and return it
		fmt.Printf("Error creating job: %v\n", result.Error)
		return result.Error
	}

	// Log the successful creation of the job
	fmt.Printf("Job created successfully: %+v\n", job)
	return nil
}

func ListJobs() ([]models.Job, error) {
	var jobs []models.Job
	result := config.DB.Find(&jobs)
	if result.Error != nil {
        fmt.Printf("Error finding jobs: %v\n", result.Error)
        return nil, result.Error
    }
    fmt.Printf("Successfully found %d jobs\n", len(jobs))
    return jobs, nil
}

// ApplyToJob creates a new application for a job by a user
func ApplyToJob(jobID, userID uint) error {
    fmt.Println("Starting ApplyToJob function")
    fmt.Printf("Received jobID: %d, userID: %d\n", jobID, userID)

    application := models.Application{
        JobID:  jobID,
        UserID: userID,
    }

    fmt.Println("Creating application:", application)

    result := config.DB.Create(&application)
    if result.Error != nil {
        fmt.Println("Error occurred while creating application:", result.Error)
        return result.Error
    }

    fmt.Println("Application created successfully")
    return nil
}
