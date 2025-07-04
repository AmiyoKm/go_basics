package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Message string
		}{
			"Hi , This is coming from the backend",
		}

		tmpl.Execute(w, data)

	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		time := fmt.Sprintf("Current time : %s", time.Now().Format(time.RFC1123))
		w.Write([]byte(time))
	})

	http.ListenAndServe(":8080", nil)
}
