package main

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"gopkg.in/gomail.v2"
)

const (
	BindIP = "localhost"
	Port   = ":8555"
)

func main() {
	u, _ := url.Parse("http://" + BindIP + Port)
	fmt.Printf("Server Started: %v\n", u)

	http.HandleFunc("/submit/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			from := r.FormValue("email")
			message := r.FormValue("message")

			if !isValidEmail(from) {
				http.Error(w, "Invalid email address", http.StatusBadRequest)
				return
			}

			if err := sendEmail(from, name, message); err != nil {
				http.Error(w, "Failed to send email", http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "Email sent successfully!")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	if err := http.ListenAndServe(Port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func sendEmail(from, name, message string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", "integranet.developers@gmail.com")
	m.SetHeader("Subject", "Message from: "+name)
	m.SetBody("text/html", "<p>"+message+"</p>")
	m.SetHeader("Reply-To", from)

	d := gomail.NewDialer("smtp.gmail.com", 587, "integranet.developers@gmail.com", "jrsb tdpc stua uhid")
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func isValidEmail(email string) bool {
	// Basic email validation using regular expression
	// This is a simple validation, you might want to use a more sophisticated method
	// depending on your requirements
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
