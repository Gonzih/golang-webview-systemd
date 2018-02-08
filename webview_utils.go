package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func slurpFile(fname string) string {
	js, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	return string(js)
}

func slurpURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func slurp(resource string) string {
	if _, err := os.Stat(resource); os.IsNotExist(err) {
		return slurpURL(resource)
	}

	return slurpFile(resource)
}

func setBodyHTML(content string) string {
	return fmt.Sprintf("(function(content){ document.body.innerHTML = content; }(`%s`))", content)
}
