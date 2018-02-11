package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func md5sum(in string) string {
	hasher := md5.New()
	hasher.Write([]byte(in))
	return hex.EncodeToString(hasher.Sum(nil))
}

func slurpFile(fname string) string {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("Error in slurpFile: %s", err)
	}

	return string(content)
}

func slurpURL(url string) string {
	res, err := http.Head(url)
	if err != nil {
		log.Fatalf("Error executing http head: %s", err)
	}

	etag := res.Header.Get("Etag")
	tmpFname := fmt.Sprintf("/tmp/webview-cache-%s", md5sum(etag))

	if _, err := os.Stat(tmpFname); !os.IsNotExist(err) {
		log.Println("Reading response for %s from cache at %s", url, tmpFname)
		return slurpFile(tmpFname)
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error executing http get: %s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading resp body: %s", err)
	}

	_, err = os.Create(tmpFname)
	if err != nil {
		log.Fatalf("Error creating cache file: %s", err)
	}

	err = ioutil.WriteFile(tmpFname, body, 0644)
	if err != nil {
		log.Fatalf("Error writing to cache file: %s", err)
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
