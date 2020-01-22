package shaker

import (
	"bufio"
	"net/http"
	"strings"

	"github.com/danish287/project-0/config"
)

//MyHoroscope stores the response for the requested horoscope
var MyHoroscope string

//myZodiac stores the zodiac user input
var myZodiac string

var linkList []string
var output string

//GetZodiacURL is a function that takes a string of integers and returns the first thing in the listt
func GetZodiacURL(zodiac string, when string, myGenre string) string {
	linkNum := GetType(when)

	myZodiac = zodiac
	mySign := GetZodiac(zodiac)
	wantedGenre := GetGenre(myGenre)

	if linkNum == -1 || mySign == -1 || wantedGenre == -1 {
		return "Please try again using valid arguments."
	}

	if myGenre == "general" {

		output = config.ConfigMe.Links[mySign].Readings[0].General[linkNum]
	} else if myGenre == "love" {
		output = config.ConfigMe.Links[mySign].Readings[0].Love[linkNum]
	} else if myGenre == "career" {
		output = config.ConfigMe.Links[mySign].Readings[0].Career[linkNum]
	} else {
		output = config.ConfigMe.Links[mySign].Readings[0].Money[0]
	}

	if when == "yearly" {
		output = config.ConfigMe.Links[mySign].Readings[0].General[3]
	}

	output = strings.ReplaceAll(output, "https", "http")

	return output
}

//GetGenre returns an integer referring to the genre of reading the user wants
func GetGenre(genre string) int {
	readingGenre := []string{"general", "love", "career", "money"}
	const (
		Gen = iota
		Luv
		Car
		Yr
	)
	switch {
	case readingGenre[Gen] == genre:
		return Gen
	case readingGenre[Luv] == genre:
		return Luv
	case readingGenre[Car] == genre:
		return Car
	case readingGenre[Yr] == genre:
		return Yr
	}
	return -1
}

//GetType returns an integer referring to the type of reading the user wants
func GetType(when string) int {
	readingLink := []string{"daily", "weekly", "monthly", "yearly"}
	const (
		Daily = iota
		Weekly
		Monthly
		Yearly
	)
	switch {
	case readingLink[Daily] == when:
		return Daily
	case readingLink[Weekly] == when:
		return Weekly
	case readingLink[Monthly] == when:
		return Monthly
	case readingLink[Yearly] == when:
		return Yearly
	}
	return -1

}

//GetZodiac returns an integer referencing the user's zodiac sign
func GetZodiac(userInput string) int {
	readingLink := []string{"aries", "taurus", "gemini", "cancer", "leo", "virgo", "libra", "scorpio", "sagittarius", "capricorn", "aquarius", "pisces"}
	const (
		Aries = iota
		Taurus
		Gemini
		Cancer
		Leo
		Virgo
		Libra
		Scorpio
		Sagittarius
		Capricorn
		Aquarius
		Pisces
	)
	switch {
	case readingLink[Aries] == userInput:
		return Aries
	case readingLink[Taurus] == userInput:
		return Taurus
	case readingLink[Gemini] == userInput:
		return Gemini
	case readingLink[Cancer] == userInput:
		return Cancer
	case readingLink[Leo] == userInput:
		return Leo
	case readingLink[Virgo] == userInput:
		return Virgo
	case readingLink[Libra] == userInput:
		return Libra
	case readingLink[Scorpio] == userInput:
		return Scorpio
	case readingLink[Sagittarius] == userInput:
		return Sagittarius
	case readingLink[Capricorn] == userInput:
		return Capricorn
	case readingLink[Aquarius] == userInput:
		return Aquarius
	case readingLink[Pisces] == userInput:
		return Pisces
	}
	return -1

}

//GetMyResponse gets requested horoscope
func GetMyResponse(url string, ReadingFor string) string {

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)

	if resp.Status == "200 OK" {
		if config.ReadingFor == "yearly" {
			heading := "<h2>" + strings.Title(myZodiac) + " Horoscope</h2>"
			count := 300
			for i := 0; scanner.Scan() && i < count; i++ {
				myLine := scanner.Text()
				if strings.Contains(myLine, heading) {
					count = i + 2
				}
				MyHoroscope = scanner.Text()
			}

		} else {

			for scanner.Scan() {
				myLine := scanner.Text()
				if strings.Contains(myLine, "<p><strong") {
					MyHoroscope = myLine
					break

				}
			}
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	MyHoroscope = CleanResponse(MyHoroscope)

	return MyHoroscope
}

//CleanResponse cleans response given by the HTTP request
func CleanResponse(txt string) string {
	txt = strings.ReplaceAll(txt, "strong", "")
	txt = strings.ReplaceAll(txt, "<p>", "")
	txt = strings.ReplaceAll(txt, "</p>", "")
	txt = strings.ReplaceAll(txt, "<>", "")
	txt = strings.ReplaceAll(txt, "</>", "")
	txt = strings.ReplaceAll(txt, "%!(EXTRA string=", "")
	txt = strings.ReplaceAll(txt, ")", "")
	txt = strings.ReplaceAll(txt, "&nbsp;", "")
	txt = strings.ReplaceAll(txt, "<br>", " ")
	return txt
}
