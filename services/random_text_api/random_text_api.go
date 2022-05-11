package random_text_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiUrl = "https://baconipsum.com/api/?type=meat-and-filler"

func GetText() (string, error) {
	resp, err := http.DefaultClient.Get(apiUrl)
	if err != nil {
		return "", fmt.Errorf("unable to reach [%v]: %v", apiUrl, err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read response body: %v", err)
	}
	return string(content), nil
}
