package main

import (
	"fmt"
	"log"

	"github.com/gtkesh/nyt-go/nyt"
)

func main() {
	nytClient := nyt.NewClient("api_secret")

	articles, err := nytClient.GetArticles("Deez Nutz")
	if err != nil {
		log.Printf("error getting articles %v\n", err)
	}
	fmt.Println(articles)

	articles, err = nytClient.GetArticles("Deez Nutz", nyt.WithBeginDate("20001109"))
	if err != nil {
		log.Printf("error getting articles %v\n", err)
	}
	fmt.Println(articles)
}
