package tides

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Tides []Tide

type Tide struct {
	DateTime     time.Time
	State        string
	HeightMetres float64
}

func GetTides() Tides {
	res, err := http.Get("https://www.tidetime.org/europe/united-kingdom/leigh-on-sea.htm")
	if err != nil {
		log.Fatal("Could not retrieve tides data")
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

	tides := []Tide{}
	doc.Find("#tideTable tbody tr td").Each(func(i int, s *goquery.Selection) {
		s.Find("ul li").Each(func(j int, s2 *goquery.Selection) {
			state := s2.Find(".tidal-state").Text()
			time := strings.ReplaceAll(s2.Find("strong").Text(), state, "")
			dateTime := createDate(dates[i], time)
			height := createHeight(strings.Split(s2.Text(), "(")[1])
			tides = append(tides, Tide{DateTime: dateTime, State: state, HeightMetres: height})
		})
	})

	return tides
}

func createHeight(heightString string) float64 {
	reg, err := regexp.Compile(`[^\.0-9]+`)
	if err != nil {
		log.Fatal(err)
	}
	height, err := strconv.ParseFloat(reg.ReplaceAllString(heightString, ""), 16)
	if err != nil {
		log.Fatal(err)
	}
	return height
}

func createDate(dateStr string, timeStr string) time.Time {
	now := time.Now()
	timeStr = strings.Trim(timeStr, " ")
	dateStr = strings.Trim(dateStr, " ")

	// Remove alpha from date
	reg, err := regexp.Compile(`[^0-9]+`)
	if err != nil {
		log.Fatal(err)
	}
	dayNo, err := strconv.Atoi(reg.ReplaceAllString(dateStr, ""))
	if err != nil {
		log.Fatal(err)
	}

	hours, err := strconv.Atoi(timeStr[:2])
	if err != nil {
		log.Fatal(err)
	}
	minutes, err := strconv.Atoi(timeStr[3:5])
	if err != nil {
		log.Fatal(err)
	}
	amPm := timeStr[5:]

	if amPm == "pm" && hours < 12 {
		hours += 12
	}
	if amPm == "am" && hours == 12 {
		hours = 0
	}
	london, err := time.LoadLocation("Europe/London")
	if err != nil {
		log.Fatal(err)
	}
	month := now.Month()
	if dayNo < now.Day() {
		month += 1
	}
	return time.Date(now.Year(), month, dayNo, hours, minutes, 0, 0, london)
}
