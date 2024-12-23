package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/*
В контексте приглашений на мероприятия, RSVP — это запрос подтверждения от приглашённого человека или людей.
RSVP — это акроним французской фразы Répondez s’il vous plaît,
означающей буквально «Будьте добры ответить» или «Пожалуйста, ответьте».
*/

type Rsvp struct {
	Name, Email, Phone string
	WillAttend 		   bool
}

var responces = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 5)

func LoadTemplates() {
	// TODO - load template here
	// There are 5 templates: welcome.html, form.html, list.html, thanks.html, sorry.html + base template(layout.html)
	templateNames := [5]string{"welcome", "form", "list", "thanks", "sorry"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name + ".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}


// welcomeHandler handles root URL
func welcomeHandler(w http.ResponseWriter, r *http.Request){
	templates["welcome"].Execute(w, nil)
}

// listHandler handles /list URL
func listHandler(w http.ResponseWriter, r *http.Request){
	templates["list"].Execute(w, responces)
}

// type formData struct {
// 	*Rsvp
// 	Errors []string
// }

// formHandler handles /form URL
func formHandler(w http.ResponseWriter, r *http.Request){
	// TODO
}

func main() {
	LoadTemplates()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}