package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Parse error", http.StatusBadRequest)
		return
	}

	name := r.FormValue("your-name")
	email := r.FormValue("your-email")
	subject := r.FormValue("your-subject")
	message := r.FormValue("your-message")

	from := "mamdouhhazemm@gmail.com"
	password := "nodh nviw kmln aeet"
	to := "mamdouhhazemm@gmail.com"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	body := fmt.Sprintf("Name: %s\nEmail: %s\nSubject: %s\n\nMessage:\n%s", name, email, subject, message)
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Contact Form Submission\n\n" + body

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		http.Error(w, "Failed to send email: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/index.html#contactForm", http.StatusSeeOther)
}

func main() {
	fs := http.FileServer(http.Dir("./DrWaelAlsaidy"))
	http.Handle("/", fs)

	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
