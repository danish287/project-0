package main

import (
	"fmt"
	"net/http"

	"github.com/danish287/project-0/config"
	"github.com/danish287/project-0/shaker"
)

func main() {
	// userIn := config.ZodiacSign

	// if userIn == "EMPTY" {
	// 	fmt.Println("\nPlease enter a zodiac sign using the flag Zodiac. \nExample: -Zodiac=Aquarius")
	// }

	// if zodiacURL == "Please try again using valid arguments." {
	// 	fmt.Println("Please try again using valid arguments.")
	// } else {

	println("Server is running on port 8080")

	http.Handle("/", http.FileServer(http.Dir("web")))
	http.HandleFunc("/horoscope", func(w http.ResponseWriter, r *http.Request) {
			var myZodiac = r.FormValue("ZodiacSign")
			zodiacURL := shaker.GetZodiacURL(userIn, config.ReadingFor, config.ReadingType)
			fmt.Fprint(w, shaker.GetMyResponse(zodiacURL, config.ReadingFor), " ", name)
		})
		http.ListenAndServe(":8080", nil)
		fmt.Println("\n\n ", shaker.GetMyResponse(zodiacURL, config.ReadingFor))
	}

}
