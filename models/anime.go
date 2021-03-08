package models

// Anime is an entity that represents a new anime search result.
type Anime struct {
	Type               string
	Name               string
	TorrentLink        string
	ThreadLink         string
	Magnet             string
	Date               string
	Size               string
	Seeders            string
	Leechers           string
	CompletedDownloads string
}
