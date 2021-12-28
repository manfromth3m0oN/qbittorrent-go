package main

import (
	"log"

	"github.com/manfromth3m0oN/qbittorrent-go/pkg"
)

func main() {
	client := pkg.NewClient("http://localhost:8080")
	client.Login("admin", "adminadmin")
	defer client.Logout()

	sort := "downloaded"
	searchCriteria := pkg.SearchCriteria{
		Filter:   nil,
		Category: nil,
		Tag:      nil,
		Sort:     &sort,
		Reverse:  false,
		Limit:    nil,
		Offset:   nil,
		Hashes:   nil,
	}

	torrents, err := client.SearchTorrents(searchCriteria)
	if err != nil {
		log.Fatalf("Failed to get torrents: %s", err)
	}

	for _, torrent := range torrents.Torrents {
		log.Println(torrent.Name)
	}
}
