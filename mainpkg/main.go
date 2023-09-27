package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type bodyVitals struct {
	Height int
	Weight int
	Eyes   string
	Hair   string
	BMI    float32
}

type person struct {
	Name string
	Age  int

	Vitals bodyVitals
}

var tpl *template.Template
var person1 person

func main() {
	person1 = person{
		Name: "Robert",
		Age:  33,

		Vitals: bodyVitals{
			Height: 65,
			Weight: 190,
			Eyes:   "Green",
			Hair:   "Black",
			BMI:    13.2,
		},
	}

	var err error
	tpl, err = template.ParseGlob("../template/*.html")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	http.HandleFunc("/person", personInfoHandler)
	http.ListenAndServe(":8888", nil)

}

func personInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("personInfoHandler running")
	tpl.ExecuteTemplate(w, "person.html", person1)
}
