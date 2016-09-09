package vk

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/asaskevich/govalidator"
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
	if !govalidator.IsURL(url) {
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

	answer := new(Answer)
	err = answer.UnmarshalJSON(body)

	if err != nil {
		return false, err
	}

	return answer.Response == 1, nil
}
