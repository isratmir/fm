package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const url  = "https://api.forismatic.com/api/1.0/?method=getQuote&format=json&lang=en"

type Quote struct {
	Text string `json:"quoteText"`
	Author string `json:"quoteAuthor"`
	Link string `json:"quoteLink"`
}

func getQuote() []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func main() {
	args := os.Args
	fmt.Println(args)
	data := getQuote()
	quote := Quote{}
	err := json.Unmarshal(data, &quote)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(quote.Text)
	fmt.Println(quote.Author)
}