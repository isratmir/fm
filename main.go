package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const BaseURL = "https://api.forismatic.com/api/1.0/?method=getQuote&format=json"

type Quote struct {
	Text   string `json:"quoteText"`
	Author string `json:"quoteAuthor"`
	Link   string `json:"quoteLink"`
}

func getQuote(lang string) (string, Quote, error) {
	q := Quote{}
	url := BaseURL + "&lang=" + lang
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return url, q, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return url, q, err
	}

	err = json.Unmarshal(body, &q)
	if err != nil {
		log.Print(err)
		return url, q, err
	}

	return url, q, nil
}

func main() {
	lang := flag.String("l", "en", "Quote language")
	flag.Parse()

	_, q, err := getQuote(*lang)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n- %s\n", q.Text, q.Author)
}
