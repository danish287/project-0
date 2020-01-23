package main

import (
	"fmt"

	"github.com/danish287/project-0/shaker"
	"github.com/danish287/project-0/config"
)

func main() {

	

	if Zodiac == "EMPTY" {
		fmt.Println("\nPlease enter a zodiac sign using the flag Zodiac. \nExample: -Zodiac=Aquarius")
	}
	
	if zodiacURL == "Please try again using valid arguments." {
		fmt.Println("\nPlease try again using valid arguments.\n")
	} else {
		//answer := shaker.GetMyResponse(zodiacURL, ReadingFor)
		zodiacURL := config.MyURL(shaker.GetZodiacURL(Zodiac, ReadingFor, ReadingType))

		fmt.Println("\n\n ", zodiacURL)
	}

}
