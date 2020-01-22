package main

import (
	"fmt"
	"net/http"
)

// var tpl *template.Template

// func init() {
// 	tpl = template.Must(template.ParseGlob("../../web/*.html"))
// }

func main() {
	println("Server is running on port 8080")

	http.Handle("/", http.FileServer(http.Dir("web")))
	//http.Handle("/", http.FileServer("../../web/*.html"))
	//"../../web/*.html"
	http.HandleFunc("/horoscope", func(w http.ResponseWriter, r *http.Request) {
		var name = r.FormValue("name")
		if name == "" {
			name = "Anonymous NAME"
		}
		fmt.Fprint(w, name)
	})
	http.ListenAndServe(":8080", nil)

	//http.ListenAndServe(":8080", nil)
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	//io.WriteString(w, "DANIA")
// 	tpl.ExecuteTemplate(w, "index.html", nil)

// }
