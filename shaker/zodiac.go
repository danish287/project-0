package shaker

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//CONFIGFILE stores the links where we get horoscopes
const CONFIGFILE string = "conf.json"

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
	config.Links = ZodiacCluster
	config.Links.Readings = LinkCluster
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config.Links.Zodiac)
	// if config.Zodiac == zodiac {
	// 	linkList = config.Readings.General
	// }
	//fmt.Println(linkNum)
	//fmt.Println(linkList)
	return 0 //linkList[linkNum]
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
