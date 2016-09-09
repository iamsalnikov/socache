package vk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/iamsalnikov/socache/helpers/validator"
)

// VK is a cache cleaner for vk.com
type VK struct {
	BaseURL string
}

// New create VK object
func New() *VK {
	return &VK{
		BaseURL: "https://api.vk.com/method/pages.clearCache",
	}
}

// Clear function drop cache
func (v VK) Clear(url string) (bool, error) {
	if !validator.IsURL(url) {
		return false, errors.New(url + " - is not valid url")
	}

	return v.sendRequest(url)
}

func (v VK) sendRequest(url string) (bool, error) {
	response, err := http.Get(v.BaseURL + "?url=" + url)
	defer response.Body.Close()

	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return false, err
	}

	var answer Answer
	err = json.Unmarshal(body, &answer)

	if err != nil {
		return false, err
	}

	if answer.Response == 1 {
		return true, nil
	}

	return false, errors.New("Response from VK is 0")
}
