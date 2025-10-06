package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type Person struct {
	Name     string
	Age      int
	Email    string
	Location string
	Job      string
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		return
	}
	template.Execute(w, nil)
}

func idcardHandler(w http.ResponseWriter, r *http.Request) {
	var person Person

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			http.Error(w, "Invalid age value", http.StatusBadRequest)
			return
		}

		person = Person{
			Name:     r.FormValue("name"),
			Age:      age,
			Email:    r.FormValue("email"),
			Location: r.FormValue("location"),
			Job:      r.FormValue("job"),
		}
	} else {
		person = Person{
			Name:     "Kik's la Légende",
			Age:      20,
			Email:    "kik's.legend@gmail.com",
			Location: "Lyon, France",
			Job:      "Student",
		}
	}

	renderTemplate(w, "pages/infos.html", person)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homepageHandler)
	http.HandleFunc("/idcard", idcardHandler)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
