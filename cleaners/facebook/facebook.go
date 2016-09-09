package facebook

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/asaskevich/govalidator"
)

// Facebook is a cache cleaner for facebook.com
type Facebook struct {
	BaseURL string
}

// NewFacebook function create Facebook cleaner object
func NewFacebook() *Facebook {
	return &Facebook{
		BaseURL: "https://graph.facebook.com",
	}
}

// Clear function drop cache
func (f Facebook) Clear(url string) (bool, error) {
	if !govalidator.IsURL(url) {
		return false, errors.New(url + " - is not valid url")
	}

	return f.sendRequest(url)
}

func (f Facebook) sendRequest(address string) (bool, error) {
	data := url.Values{}
	data.Set("id", address)
	data.Add("scrape", "true")

	response, err := http.PostForm(f.BaseURL, data)
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

	var nilError AnswerError
	return answer.Error == nilError, errors.New(answer.Error.Message)
}
