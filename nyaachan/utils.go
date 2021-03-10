package nyaachan

import (
	"fmt"
	"nyaachan/models"
	"strings"
)

// What?
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// A utility function to Change the Anime Entities to MD for Telegram.
func markDownify(d []models.Anime) string {

	result := []string{}
	for _, a := range d {
		if a.Name != "" {
			a.Name = strings.ReplaceAll(a.Name, "[", "{")
			a.Name = strings.ReplaceAll(a.Name, "]", "}")
			result = append(result, fmt.Sprintf(
				"*%s*\n|[Thread Link](%s) | [Torrent Link](%s)|\n|*Seeders -*%v|*Leechers -* %v|*Completed -* %v| \n", a.Name, baseURL+a.ThreadLink, baseURL+a.TorrentLink, a.Seeders, a.Leechers, a.CompletedDownloads,
			))
		}
	}
	resultStr := strings.Join(result, "\n")
	return resultStr
}
