package main

import (
	"fmt"
	"net/http"

	"github.com/danish287/project-0/config"
	"github.com/danish287/project-0/shaker"
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
		var myZodiac = r.FormValue("zod")
		var myChooseType = r.FormValue("type")
		var myChooseReading = r.FormValue("reading")
		if myZodiac == "" || myChooseType == "" || myChooseReading == "" {
			fmt.Fprint(w, "Please try again using valid arguments.")
		} else {
			//config.CONFIGFILE = "conf.json"
			zodiacURL := shaker.GetZodiacURL(myZodiac, myChooseType, myChooseReading)
			config.ReadingFor = myChooseReading
			//fmt.Println("\n\n ", shaker.GetMyResponse(zodiacURL, config.ReadingFor))
			fmt.Fprint(w, "\n\n ", shaker.GetMyResponse(zodiacURL, config.ReadingFor))
		}

	})
	http.ListenAndServe(":8080", nil)

	//http.ListenAndServe(":8080", nil)
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	//io.WriteString(w, "DANIA")
// 	tpl.ExecuteTemplate(w, "index.html", nil)

// }
