package shaker

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

//CONFIGFILE stores the links where we get horoscopes
const CONFIGFILE string = "conf.json"

//MyHoroscope stores the response for the requested horoscope
var MyHoroscope string

//MyURL stores the URL we are trying to get
var MyURL string

//myZodiac stores the zodiac user input
var myZodiac string

//LinkCluster derives Links from json file
type LinkCluster struct {
	General []string
	Love    []string
}

//ZodiacCluster derives links from json file
type ZodiacCluster struct {
	Zodiac   string
	Readings []LinkCluster
}

//Configuration derives links from json file
type Configuration struct {
	Links []ZodiacCluster
}

var linkList []string

//GetZodiacURL is a function that takes a string of integers and returns the first thing in tthe listt
func GetZodiacURL(zodiac string, when string) string {
	linkNum := GetType(when)
	config := Configuration{}
	myLink, err := os.Open(CONFIGFILE)
	myZodiac = zodiac
	json.NewDecoder(myLink).Decode(&config)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(linkNum)
	link2 := config.Links[1].Readings[0].General[linkNum]
	link2 = strings.ReplaceAll(link2, "https", "http")

	return link2
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

//GetMyResponse gets requested horoscope
func GetMyResponse(url string, ReadingFor string) string {
	//println("\nMY RUSL \n", url)
	resp, err := http.Get(url)
	if err != nil {
		//fmt.Println("ERROR IS HERE")
		//fmt.Println("NOT NIL")
		panic(err)
	}

	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)

	if resp.Status == "200 OK" {

		if ReadingFor == "yearly" {
			//<h2>Aquarius Horoscope</h2>
			heading := "<h2>" + strings.Title(myZodiac) + " Horoscope</h2>"
			//fmt.Println(heading)
			count := 300
			for i := 0; scanner.Scan() && i < count; i++ {
				myLine := scanner.Text()
				//fmt.Println(myLine)
				if strings.Contains(myLine, heading) {
					//MyHoroscope = myLine
					count = i + 2

				}
				MyHoroscope = scanner.Text()
			}

		} else {
			for scanner.Scan() {
				myLine := scanner.Text()
				//fmt.Println(myLine)
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

	return CleanResponse(MyHoroscope)
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
	//txt = strings.ReplaceAll(txt, "&nbsp;-&nbsp;", "")
	return txt
}
