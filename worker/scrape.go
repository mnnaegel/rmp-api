package worker

import (
	"fmt"
	"rmpParser/models"
)

// Create a worker that will scrape the data of each professor on this page `https://www.ratemyprofessors.com/school?sid=1491`
func Scrape () []model.Professor {
	// create a new page scraper with url set to "https://www.ratemyprofessors.com/school?sid=1491"
	
	scraper := PageScraper{
		URL: "https://www.ratemyprofessors.com/search/teachers?query=*&sid=1491",
	}
	
	// fetch the document from the page scraper
	doc, err := scraper.FetchDocument(); if err != nil {
		fmt.Println("Error fetching document")
	}

	if (scraper.Status != "200") {
		fmt.Println("Error. Received status code:", scraper.Status)
	}

	s, err := doc.Html(); if err != nil {
		fmt.Println("Error getting html")
	}

	fmt.Println(scraper.Status, s)

	// Wait for the page to load 

	// professors on the page
	fmt.Println("Scraping for professors")
	professors := scraper.scrapeProfessors(doc)

	fmt.Println(professors)
	return professors
}