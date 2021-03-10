package scraper

import (
	"log"
	"net/http"
	"nyaachan/models"

	"github.com/PuerkitoBio/goquery"
)

// FindAnime takes a search query as parameter and returns a
// models.Anime result from nyaa.si . The query param does not
// need to be a urlsafe string. We do it before using the parameter
func FindAnime(uri string) []models.Anime {
	var results []models.Anime
	res, err := http.Get(uri)
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
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		var res models.Anime
		child := s.Find("td")
		child.Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				res.Type, _ = s.Children().Attr("title")
			case 1:
				if len(s.Children().Nodes) == 2 {
					res.ThreadLink, _ = s.Children().Next().Attr("href")
					res.Name, _ = s.Children().Next().Attr("title")
				} else {
					res.ThreadLink, _ = s.Attr("href")
					res.Name, _ = s.Children().Attr("title")
				}
			case 2:
				res.TorrentLink, _ = s.Children().Attr("href")
				res.Magnet, _ = s.Children().Next().Attr("href")
			case 3:
				res.Size = s.Text()
			case 4:
				res.Date = s.Text()
			case 5:
				res.Seeders = s.Text()
			case 6:
				res.Leechers = s.Text()
			case 7:
				res.CompletedDownloads = s.Text()
			}
		})
		results = append(results, res)
	})
	return results
}
