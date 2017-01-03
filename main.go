package main

import (
	"fmt"
	"log"

	"github.com/gtkesh/nyt-go/nyt"
)

func main() {
	nytClient := nyt.NewClient("api_secret")

	// Search articles with "Barack Obama" mentioned in them.
	// https://api.nytimes.com/svc/search/v2/articlesearch.json?&q=Barack+Obama
	articles, err := nytClient.GetArticles("Barack Obama")
	if err != nil {
		log.Printf("error getting articles %v\n", err)
	}
	// Print out the url of the first article from the results.
	fmt.Println(articles[0].WebURL)

	// Search articles with "Barack Obama" mentioned in them in a specified time range and sorted by newest articles first.
	// https://api.nytimes.com/svc/search/v2/articlesearch.json?&q=Barack+Obama&begin_date=20090000&end_date=20110000
	articles, err = nytClient.GetArticles("Barack Obama",
		nyt.WithBeginDate("20090000"),
		nyt.WithEndDate("20110000"),
		nyt.SortedByNewest(),
	)
	if err != nil {
		log.Printf("error getting articles %v\n", err)
	}
	// Print out the url of the first article from the results.
	fmt.Println(articles[0].WebURL)

}
