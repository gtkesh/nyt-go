# nyt-go
nyt-go is a Golang client implementation of the [New York Times API](https://developer.nytimes.com/).

**Documentation:** [![GoDoc](https://godoc.org/github.com/gtkesh/nyt-go?status.svg)](https://godoc.org/github.com/gtkesh/nyt-go/nyt)


Install via `go get github.com/gtkesh/nyt-go`.


Currently supported NYT APIs:
- [Article Search](https://developer.nytimes.com/article_search_v2.json)

TODO:
- Complete README
- Add tests
- Support more APIs


## Usage:

```go
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
articles, err = nytClient.GetArticles(
  "Barack Obama",
  // Note how you can chain options.
  nyt.WithBeginDate("20090000"),
  nyt.WithEndDate("20110000"),
  nyt.SortedByNewest(),
)
if err != nil {
  log.Printf("error getting articles %v\n", err)
}
// Print out the url of the first article from the results.
fmt.Println(articles[0].WebURL)
```
