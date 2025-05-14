package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Configure CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:7070"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Serve static files
	r.Static("/", "./frontend")

	// Contact form route
	r.POST("/contact", func(c *gin.Context) {
		// Validate request method (already handled by Gin router)

		// Parse form data
		name := c.PostForm("your-name")
		email := c.PostForm("your-email")
		subject := c.PostForm("your-subject")
		message := c.PostForm("your-message")

		// Email configuration (MOVE TO ENV VARIABLES IN PRODUCTION!)
		from := "mamdouhhazemm@gmail.com"
		password := "nodh nviw kmln aeet"
		to := "mamdouhhazemm@gmail.com"
		smtpHost := "smtp.gmail.com"
		smtpPort := "587"

		// Create email body
		body := fmt.Sprintf("Name: %s\nEmail: %s\nSubject: %s\n\nMessage:\n%s",
			name, email, subject, message)
		msg := "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: Contact Form Submission\n\n" + body

		// SMTP authentication
		auth := smtp.PlainAuth("", from, password, smtpHost)

		// Send email
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to send email: " + err.Error(),
			})
			return
		}

		// Redirect to frontend contact form
		c.Redirect(http.StatusSeeOther, "/index.html#contactForm")
	})

	// Start server
	fmt.Println("Listening on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server startup error: ", err)
	}
}
