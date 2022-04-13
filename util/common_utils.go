package util

import (
	"io/ioutil"
	"net/http"
)

// Reading the byte stream from a html page given the url
func ReadBytes(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// this checks whether a url is reachable
func IsReachable(url string) bool {

	if len(url) < 4 {
		return false
	}

	response, err := http.Get(url)

	if err != nil {
		return false
	}

	if response.StatusCode == 200 {
		return true
	}

	return false
}
