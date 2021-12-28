package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Filter string

const (
	All                Filter = "all"
	Downloading        Filter = "downloading"
	Seeding            Filter = "seeding"
	Completed          Filter = "completed"
	Paused             Filter = "paused"
	Active             Filter = "active"
	Inactive           Filter = "inactive"
	Resumed            Filter = "resumed"
	Stalled            Filter = "stalled"
	StalledUploading   Filter = "stalled_uploading"
	StalledDownloading Filter = "stalled_downloading"
	Errored            Filter = "errored"
)

type SearchCriteria struct {
	Filter   *Filter
	Category *string
	Tag      *string
	Sort     *string
	Reverse  bool
	Limit    *int
	Offset   *int
	Hashes   *[]string // Seperated by '|'
}

type TorrentList struct {
	Torrents []TorrentData
}

type TorrentData struct {
	Dlspeed       int     `json:"dlspeed"`
	Eta           int     `json:"eta"`
	FLPiecePrio   bool    `json:"f_l_piece_prio"`
	ForceStart    bool    `json:"force_start"`
	Hash          string  `json:"hash"`
	Category      string  `json:"category"`
	Tags          string  `json:"tags"`
	Name          string  `json:"name"`
	NumComplete   int     `json:"num_complete"`
	NumIncomplete int     `json:"num_incomplete"`
	NumLeechs     int     `json:"num_leechs"`
	NumSeeds      int     `json:"num_seeds"`
	Priority      int     `json:"priority"`
	Progress      float64 `json:"progress"`
	Ratio         int     `json:"ratio"`
	SeqDl         bool    `json:"seq_dl"`
	Size          int     `json:"size"`
	State         string  `json:"state"`
	SuperSeeding  bool    `json:"super_seeding"`
	Upspeed       int     `json:"upspeed"`
}

func (c *Client) SearchTorrents(search SearchCriteria) (*TorrentList, error) {
	url := fmt.Sprintf("%s%s%s%s", c.hostname, c.apiVer, "/torrents/info", search.Marshal())
	c.logger.Println(url)
	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		c.logger.Printf("Failed to create request: %s", err)
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: "SID", Value: c.token})

	resp, err := c.httpClient.Do(req)
	// TODO: Missing case for not 200 status
	if err != nil || resp.StatusCode != 200 {
		c.logger.Printf("Failed to send request: %s", err)
		return nil, err
	}

	c.logger.Println(resp.Body)

	var (
		torrentList TorrentList
		bodyBytes   []byte
	)
	_, err = resp.Body.Read(bodyBytes)
	if err != nil {
		c.logger.Printf("Failed to read resp body")
		return nil, err
	}

	c.logger.Println(bodyBytes)

	err = json.Unmarshal(bodyBytes, &torrentList)
	if err != nil {
		c.logger.Printf("Failed to unmarshal API response")
		return nil, err
	}

	return &torrentList, nil
}

// TODO: This really should return an error
func (s *SearchCriteria) Marshal() string {
	reqString := "?"
	if s.Filter != nil {
		reqString += fmt.Sprintf("filter=%s", *s.Filter)
	}
	if s.Category != nil {
		reqString += fmt.Sprintf("category=%s", *s.Category)
	}
	if s.Tag != nil {
		reqString += fmt.Sprintf("tag=%s", *s.Tag)
	}
	if s.Sort != nil {
		reqString += fmt.Sprintf("sort=%s", *s.Sort)
	}
	if s.Limit != nil {
		limit := strconv.Itoa(*s.Limit)
		reqString += fmt.Sprintf("limit=%s", limit)
	}
	if s.Offset != nil {
		offset := strconv.Itoa(*s.Offset)
		reqString += fmt.Sprintf("offset=%s", offset)
	}
	if s.Hashes != nil {
		var hashes string
		for _, hash := range *s.Hashes {
			hashes += hash + "|"
		}
		reqString += fmt.Sprintf("hashes=%s", hashes)
	}
	return reqString
}
