package config

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"flag"
)

//Zodiac stores name of zodiac
var Zodiac string

//ReadingType stores type of reading user is looking get
var ReadingType string

//ReadingFor specifies when the reader wants the reading
var ReadingFor string

//CONFIGFILE stores the links where we get horoscopes
const CONFIGFILE string = "conf.json"

//MyURL stores the URL we are trying to get
var MyURL string

//LinkCluster derives Links from json file
type LinkCluster struct {
	General []string
	Love    []string
	Career  []string
	Money   []string
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


func init() {

	config := Configuration{}
	myLink, err := os.Open(CONFIGFILE)
	json.NewDecoder(myLink).Decode(&config)

	if err != nil {
		log.Fatal(err)
	}
	
	flag.StringVar(&Zodiac, "Zodiac", "EMPTY", "name of the zodiac sign")
	flag.StringVar(&ReadingFor, "ReadingFor", "daily", "time period of reading")
	flag.StringVar(&ReadingType, "Type", "general", "type of reading")
	flag.Parse()


	
}
