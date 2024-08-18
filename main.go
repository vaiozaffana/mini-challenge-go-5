package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Email   string
	Address string
	Job     string
	Reason  string
}

var users = map[string]User{
	"vaiopasha@example.com": {
		Email:   "vaiopasha@example.com",
		Address: "Sidoarjo",
		Job:     "Student",
		Reason:  "Alasan Vaio",
	},
	"vaiozaffana@example.com": {
		Email:   "vaiozaffana@example.com",
		Address: "Surabaya",
		Job:     "Pekerjaan Zaffana",
		Reason:  "Alasan Zaffana",
	},
	"pashazaffana@example.com": {
		Email:   "pashazaffana@example.com",
		Address: "Jawa Timur",
		Job:     "Pekerjaan Pasha",
		Reason:  "Alasan Pasha",
	},
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		AvailableEmails []string
		ErrorMessage    string
	}

	availableEmails := make([]string, 0, len(users))
	for Email := range users {
		availableEmails = append(availableEmails, Email)
	}

	data := PageData{AvailableEmails: availableEmails}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		if user, ok := users[email]; ok {
			tmpl := template.Must(template.ParseFiles("templates/profile.html"))
			tmpl.Execute(w, user)
			return
		} else {
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			return
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, data)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	message := "Email is not registered!"

	tmpl := template.Must(template.ParseFiles("templates/error.html"))
	tmpl.Execute(w, template.HTML(message))
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/error", errorHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
