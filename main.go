package main

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

type websiteHealth struct {
	ID                  int
	URL                 string
	Name                string
	Status              string
	LastCheckedDatetime string
}

var google = &websiteHealth{
	ID:                  1,
	URL:                 "https://google.com",
	Name:                "Google",
	Status:              "0",
	LastCheckedDatetime: getCurrentDatetime(),
}

var aparat = &websiteHealth{
	ID:                  2,
	URL:                 "https://aparat.com",
	Name:                "Aparat",
	Status:              "0",
	LastCheckedDatetime: getCurrentDatetime(),
}

var digikala = &websiteHealth{
	ID:                  3,
	URL:                 "https://www.digikala.com/",
	Name:                "Digikala",
	Status:              "0",
	LastCheckedDatetime: getCurrentDatetime(),
}

var wikipedia = &websiteHealth{
	ID:                  4,
	URL:                 "https://www.wikipedia.org/",
	Name:                "Wikipedia",
	Status:              "0",
	LastCheckedDatetime: getCurrentDatetime(),
}

var wailsDoc = &websiteHealth{
	ID:                  5,
	URL:                 "https://wails.app",
	Name:                "Wails Doc",
	Status:              "0",
	LastCheckedDatetime: getCurrentDatetime(),
}

var sampleError = &websiteHealth{
	ID:                  6,
	URL:                 "https://httpstat.us/500",
	Name:                "httpstat 500",
	Status:              "0",
	LastCheckedDatetime: getCurrentDatetime(),
}

var sampleNotFound = &websiteHealth{
	ID:                  7,
	URL:                 "https://httpstat.us/404",
	Name:                "httpstat 404",
	Status:              "0",
	LastCheckedDatetime: getCurrentDatetime(),
}

var websiteHealthList = make([]*websiteHealth, 0)

var client = http.Client{Timeout: 5 * time.Second}

var wg sync.WaitGroup

func getCurrentDatetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func urlHealthCheck(w *websiteHealth, wg *sync.WaitGroup) {
	resp, err := client.Head(w.URL)
	w.LastCheckedDatetime = getCurrentDatetime()
	if err != nil {
		w.Status = "N/A"
		log.Println(err)
	} else {
		w.Status = strconv.Itoa(resp.StatusCode)
	}
	defer wg.Done()
}

func updateListHealthCheck() []*websiteHealth {
	wg.Add(len(websiteHealthList))
	for _, w := range websiteHealthList {
		go urlHealthCheck(w, &wg)
	}
	wg.Wait()
	return websiteHealthList
}

func main() {

	websiteHealthList = append(websiteHealthList, google)
	websiteHealthList = append(websiteHealthList, aparat)
	websiteHealthList = append(websiteHealthList, digikala)
	websiteHealthList = append(websiteHealthList, wikipedia)
	websiteHealthList = append(websiteHealthList, wailsDoc)
	websiteHealthList = append(websiteHealthList, sampleError)
	websiteHealthList = append(websiteHealthList, sampleNotFound)

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  800,
		Height: 500,
		Title:  "Sample Wails Healtcheck App",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(updateListHealthCheck)
	app.Run()
}
