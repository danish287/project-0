package shaker

import (
	"bufio"
	"encoding/json"
	"fmt"
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

//GetZodiac is a function that takes a string of integers and returns the first thing in tthe listt
func GetZodiac(zodiac string, when string) int {
	linkNum := GetType(when)
	config := Configuration{}
	myLink, err := os.Open(CONFIGFILE)
	json.NewDecoder(myLink).Decode(&config)
	fmt.Println("YUUUUUUh\n", config.Links[0].Readings[0].General[linkNum])

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(linkNum)

	return 0
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

//GetResponse gets requested horoscope
func GetResponse(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)

	if resp.Status == "200 OK" {
		for i := 0; scanner.Scan() && i < 268; i++ {
			if i == 267 {
				MyHoroscope = scanner.Text()
				print("\n\n\n\nANSWER", MyHoroscope)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

//CleanResponse cleans response given by the HTTP request
func CleanResponse(txt string) string {
	txt = strings.ReplaceAll(txt, "strong", "")
	txt = strings.ReplaceAll(txt, "<p>", "")
	txt = strings.ReplaceAll(txt, "</p>", "")
	txt = strings.ReplaceAll(txt, "<>", "")
	txt = strings.ReplaceAll(txt, "</>", "")
	return txt
}
