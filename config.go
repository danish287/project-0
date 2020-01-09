package main

import (
	"bufio"
	"log"
	"os"
)

var zodiacSigns []string

func init() {
	file, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(file)

	for reader.Scan() {

		zodiacSigns = append(zodiacSigns, reader.Text())
	}
}
