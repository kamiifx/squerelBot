package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"squerelBot/clients"
)

type PixBayResponse struct {
	TotalHits int   `json:"totalHits"`
	Hits      []HIT `json:"hits"`
}

type HIT struct {
	ID            int    `json:"id"`
	PreviewUrl    string `json:"previewURL"`
	ImageUrl      string `json:"imageURL"`
	LargeImageUrl string `json:"LargeImageUrl"`
	FullHdUrl     string `json:"fullHDURL"`
}

func QueryImage(query string) (string, error) {
	var page int
	if query == "sphynx" {
		page = 1
	} else {
		page = rand.Intn(20) + 1
	}
	response, err := clients.GetImage(page, query)
	if err != nil {
		return "", fmt.Errorf("Error calling GetImage: %v", err)
	}

	if len(response) == 0 {
		return "", fmt.Errorf("Empty response received from PixBay")
	}

	var pixBayResponse PixBayResponse
	err = json.Unmarshal(response, &pixBayResponse)
	if err != nil {
		return "", fmt.Errorf("Error decoding PixBay Response: %v", err)
	}

	if len(pixBayResponse.Hits) > 0 {
		randomIndex := rand.Intn(len(pixBayResponse.Hits))
		return pixBayResponse.Hits[randomIndex].LargeImageUrl, nil
	}

	return "", nil
}
