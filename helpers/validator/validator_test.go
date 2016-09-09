package validator

import "testing"

func TestIsURL(t *testing.T) {
	if !IsURL("https://facebook.com") {
		t.Error("https://facebook.com - is url, but got false")
	}

	if !IsURL("http://facebook.com") {
		t.Error("http://facebook.com - is url, but got false")
	}

	if IsURL("ftp://facebook.com") {
		t.Error("ftp://facebook.com - is not url, but got true")
	}

	if !IsURL("facebook.com") {
		t.Error("facebook.com - is url, but got false")
	}
}
