package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Tide struct {
	State  string
	Time   string
	Height string
}

type DayTide struct {
	Date  string
	Tides []Tide
}

func GetTides() []DayTide {
	res, err := http.Get("https://www.tidetime.org/europe/united-kingdom/leigh-on-sea.htm")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	dates := []string{}

	doc.Find("#tideTable th").Each(func(i int, s *goquery.Selection) {
		date := s.Text()
		dates = append(dates, date)
	})

	dayTides := []DayTide{}
	doc.Find("#tideTable tbody tr td").Each(func(i int, s *goquery.Selection) {
		dt := DayTide{Date: dates[i], Tides: []Tide{}}
		s.Find("ul li").Each(func(j int, s2 *goquery.Selection) {
			state := s2.Find(".tidal-state").Text()
			time := strings.ReplaceAll(s2.Find("strong").Text(), state, "")
			height := strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(strings.ReplaceAll(s2.Text(), state, ""), time, ""),
					"(",
					"",
				),
				")", "",
			)
			dt.Tides = append(dt.Tides, Tide{State: state, Time: time, Height: height})
		})
		dayTides = append(dayTides, dt)
	})

	return dayTides
}

func main() {
	dt := GetTides()
	fmt.Printf("%+v\n", dt)
}
