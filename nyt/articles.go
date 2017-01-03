package nyt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const baseURL = "https://api.nytimes.com/svc/search/v2/articlesearch.json"

type Client struct {
	baseURl    string
	secret     string
	httpClient *http.Client
}

func NewClient(secret string) *Client {
	return &Client{baseURL, secret, &http.Client{}}
}

type Article struct {
	WebURL        string `json:"web_url"`
	Snippet       string `json:"snippet"`
	LeadParagraph string `json:"lead_paragraph"`
	Abstract      string `json:"abstract"`
	PrintPage     string `json:"print_page"`

	Source   string `json:"source"`
	Headline struct {
		Main   string `json:"main"`
		Kicker string `json:"kicker"`
	} `json:"headline"`

	Keywords []keyword `json:"keywords"`

	PubDate        string       `json:"pub_date"`
	DocumentType   string       `json:"document_type"`
	NewsDesk       string       `json:"news_desk"`
	SectionName    string       `json:"section_name"`
	SubsectionName string       `json:"subsection_name"`
	TypeOfMaterial string       `json:"type_of_material"`
	Id             string       `json:"_id"`
	Multimetdia    []multimedia `json:"multimedia"`

	// TODO(giorgi)
	// WordCount      string       `json:"word_count"`
	// ByLine         byline       `json:"byline"`
	// Blog          []struct {
	// } `json:"blog"`
}

type byline struct {
	Contributor string   `json:"contributor"`
	Person      []person `json:"person"`
	Original    string   `json:"original"`
}

type personItem struct {
	Organization string `json:"organization"`
	Role         string `json:"role"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Rank         int    `json:"rank"`
}

type person struct {
}

type keyword struct {
	Rank  string `json:"rank"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type multimedia struct {
	URL       string `json:"url"`
	Format    string `json:"format"`
	Height    int    `json:"height"`
	Width     int    `json:"width"`
	Type      string `json:"type"`
	SubType   string `json:"subtype"`
	Caption   string `json:"caption"`
	Copyright string `json:"copyright"`
}

type getArticlesResponse struct {
	Response struct {
		Articles []Article `json:"docs"`
	} `json:"response"`
	Meta struct {
		Hits   int `json:"hits"`
		Time   int `json:"time"`
		Offset int `json:"offset"`
	} `json:"meta"`
}

type Option func(*ArticleRequestOptions)

type ArticleRequestOptions struct {
	opts map[string]string
}

// WithBeginDate restricts responses to results with publication dates
// of the date specified or later.
// Date format: "YYYYMMDD"
func WithBeginDate(date string) Option {
	return func(opts *ArticleRequestOptions) {
		opts.opts["begin_date"] = date
	}
}

// WithBeginDate restricts responses to results with publication dates
// of the date specified or earlier.
// Date format: "YYYYMMDD"
func WithEndDate(date string) Option {
	return func(opts *ArticleRequestOptions) {
		opts.opts["end_date"] = date
	}
}

func SortedByNewest() Option {
	return func(opts *ArticleRequestOptions) {
		opts.opts["sort"] = "newest"
	}
}

func SortedByOldest() Option {
	return func(opts *ArticleRequestOptions) {
		opts.opts["sort"] = "oldest"
	}
}

func (c *Client) GetArticles(query string, opts ...Option) ([]Article, error) {
	defaultOpts := map[string]string{
		"api-key": c.secret,
		"q":       query,
	}

	options := ArticleRequestOptions{opts: defaultOpts}

	for _, o := range opts {
		o(&options)
	}

	url := endpointFromOpts(c.baseURl, options)

	var resp getArticlesResponse
	err := c.getAndUnmarshal(url, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Response.Articles, nil
}

func (c *Client) getAndUnmarshal(url *url.URL, v *getArticlesResponse) error {
	resp, err := c.httpClient.Get(url.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, v)
	if err != nil {
		return err
	}

	return nil
}

func endpointFromOpts(baseURL string, articleRequestOptions ArticleRequestOptions) *url.URL {
	u, _ := url.ParseRequestURI(baseURL)
	data := url.Values{}
	for k, v := range articleRequestOptions.opts {
		data.Add(k, v)
	}
	u.RawQuery = data.Encode()

	return u
}
