package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url  = "https://api.forismatic.com/api/1.0/?method=getQuote&format=json"

type Quote struct {
	Text string `json:"quoteText"`
	Author string `json:"quoteAuthor"`
	Link string `json:"quoteLink"`
}

func (q *Quote) getQuote(lang string) {
	URL := url + "&lang="+lang
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err2 := json.Unmarshal(body, q)
	if err2 != nil{
		log.Fatal(err)
	}
}

func main() {
	lang := flag.String("l", "en", "Quote language")
	flag.Parse()
	quote := Quote{}
	quote.getQuote(*lang)

	fmt.Printf("%s\n- %s\n", quote.Text, quote.Author)
}