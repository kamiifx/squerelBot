package clients

import (
	"io"
	"net/http"
	"squerelBot/utils"
	"strconv"
)

func GetImage(page int, imageQuery string) ([]byte, error) {
	url := "https://pixabay.com/api/?key=" +
		utils.PixBayKey.GetValue() +
		"&q=" + imageQuery + "&image_type=photo&page=" +
		strconv.Itoa(page)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
