package rss

import (
	"log"
	"os"

	"github.com/mmcdole/gofeed"
)

func Rss() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("") // add your RSS feed url here

	f, err := os.Create("./newsletter.html")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.WriteString("<h2>Check Out This Week's Articles: </h2><br> ") // replace this line with your preferred newsletter intro

	for i := 0; i < 5; i++ {
		f.WriteString("<p><a href='" + feed.Items[i].Link + "'>" + feed.Items[i].Title + "</a> - " + feed.Items[i].Description + "</p> \n")

	}

	f.WriteString("Thanks for being part of the Opensource Geeks community💻🐧") // replace this line with your preffered newsletter outro or conclusion
}
