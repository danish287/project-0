package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

var response1 string

func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}

func main() {

	resp, err := http.Get("https://horoscope.com/us/horoscopes/general/horoscope-general-daily-today.aspx?sign=2")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("RESPONSE STATUS", resp.Status)
	scanner := bufio.NewScanner(resp.Body)

	if resp.Status == "200 OK" {
		for i := 0; scanner.Scan() && i < 268; i++ {
			//fmt.Println("SACNNER TEXT", scanner.Text())
			if i == 267 {
				response1 = scanner.Text()
				//print("\nANSWER\n\n", scanner.Text())
			}
		}
	}

	// for scanner.Scan() {
	// 	fmt.Println("SACNNER TEXT", scanner.Text())
	// }

	// bodyBytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// bodyString := string(bodyBytes)
	// fmt.Println(bodyString)
	//fmt.Println("\n\n\n\n")

	fmt.Println("ANSWER\n\n", CleanResponse(response1))

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func CleanResponse(txt string) string {
	txt = strings.ReplaceAll(txt, "strong", "")
	txt = strings.ReplaceAll(txt, "<p>", "")
	txt = strings.ReplaceAll(txt, "</p>", "")
	txt = strings.ReplaceAll(txt, "<>", "")
	txt = strings.ReplaceAll(txt, "</>", "")
	return txt
}
