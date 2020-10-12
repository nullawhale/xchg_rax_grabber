package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tkanos/gonfig"
	"log"
	"net/http"
)

type Conf struct {
	Url string
}

func getXchg(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#snip_content_cell").Each(func(i int, s *goquery.Selection) {
		band := s.
			Find("table").
			Find("tbody").
			Find("tr").
			Find("td").
			Find("pre").Text()
		fmt.Printf("%s\n", band)
	})
}

func main() {
	conf := Conf{}
	err := gonfig.GetConf("config/top.secret.json", &conf)
	if err != nil {
		log.Fatal(err)
	}

	url := conf.Url

	for i, b := range make([]uint16, 64) {
		fmt.Printf("Page: %d\n", i+1)
		s := fmt.Sprintf("%s%02x.html", url, b)
		getXchg(s)
	}
}
