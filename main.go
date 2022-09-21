package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", WelcomeHandler)
	http.HandleFunc("/foo", FooHandler)
	http.HandleFunc("/bar", BarHandler)
	http.ListenAndServe(":8080", nil)
}

type User struct {
	Name        string
	nationality string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("welcomeform.html")
		check(err)
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		myUser := User{}
		myUser.Name = r.Form.Get("entered_name")
		myUser.nationality = r.Form.Get("entered_nationality")
		t, err := template.ParseFiles("welcomerespone.html")
		check(err)
		t.Execute(w, myUser)
	}

}

func FooHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello from foo!")
}

func BarHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello from bar!")
}
