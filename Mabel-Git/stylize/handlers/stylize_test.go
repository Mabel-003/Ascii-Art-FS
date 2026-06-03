package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	asciiartgenerator "stylize/asciiArtGenerator"
	"testing"
)

func TestAsciiPage(t *testing.T) {
	_ = os.Chdir("../")
	form := url.Values{}
	form.Set("inputText", "Hello")
	form.Set("banner", "shadow")

	req, err := http.NewRequest("POST", "/ascii-art", strings.NewReader(form.Encode()))
	if err != nil {
		log.Print(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	AsciiPage(rr, req)

	status := rr.Code
	expected := http.StatusOK

	if status != expected {
		t.Errorf("Got: %v\nExpected: %v", status, expected)
	}
	result := asciiartgenerator.PrintAsciiArt("Hello", "shadow")
	if !strings.Contains(rr.Body.String(), result) {
		t.Errorf("Got :\n%v\nExpected:\n%v", result, rr.Body.String())
	}
}
