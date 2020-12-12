package core

import (
	"html/template"
	"log"
	"net/http"
)

/* A web server will serve an html template to output the data fetched with the Velib API */
func HandleWeb() {
	http.HandleFunc("/", serveTemplate)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")

	data := GetVelibRecords()

	err := tmpl.ExecuteTemplate(w,"index.html" ,data)
	if err != nil {
		log.Fatal(err)
	}
}
