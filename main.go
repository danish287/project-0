package main

import (
	"fmt"

	"github.com/danish287/project-0/config"
	"github.com/danish287/project-0/shaker"
)

func main() {
	userIn := config.ZodiacSign

	if userIn == "EMPTY" {
		fmt.Println("\nPlease enter a zodiac sign using the flag Zodiac. \nExample: -Zodiac=Aquarius")
	}

	zodiacURL := shaker.GetZodiacURL(userIn, config.ReadingFor, config.ReadingType)

	if zodiacURL == "Please try again using valid arguments." {
		fmt.Println("Please try again using valid arguments.")
	} else {
		fmt.Println("\n\n ", shaker.GetMyResponse(zodiacURL, ReadingFor))
	}

}
