package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/PuerkitoBio/goquery"
    "github.com/andlabs/ui"
)

func scrapeWebsite(url string) {
    res, err := http.Get(url)
    if err != nil {
        log.Fatal("Failed to scrape website:", err)
    }
    defer res.Body.Close()

    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatal("Failed to parse HTML:", err)
    }

    // Do something with the doc object (extract data, etc.)
    // For now, let's just print the title
    title := doc.Find("title").First().Text()
    ui.MsgBox("Scraped Content", fmt.Sprintf("Title: %s", title))
}

func onSubmit(entry *ui.Entry) {
    url := entry.Text()
    scrapeWebsite(url)
}

func main() {
    err := ui.Main(func() {
        box := ui.NewVerticalBox()
        window := ui.NewWindow("Web Scraper", 400, 120, false)
        window.OnClosing(func(*ui.Window) bool {
            ui.Quit()
            return true
        })

        entry := ui.NewEntry()
        entry.SetReadOnly(false)

        button := ui.NewButton("Scrape")
        button.OnClicked(func(*ui.Button) {
            onSubmit(entry)
        })

        box.Append(ui.NewLabel("Enter website URL:"), false)
        box.Append(entry, false)
        box.Append(button, false)

        window.SetChild(box)
        window.Show()
    })

    if err != nil {
        log.Fatal(err)
    }
}
