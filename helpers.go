package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func ContainsInt(xs []int, x int) bool {
	for _, i := range xs {
		if i == x {
			return true
		}
	}
	return false
}

func ContainsString(xs []string, x string) bool {
	for _, i := range xs {
		if i == x {
			return true
		}
	}
	return false
}

func DownloadFile(httpClient *http.Client, url string, filepath string) error {
	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer closeDef(resp.Body)

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer closeDef(file)

	_, err = io.Copy(file, resp.Body)

	return err
}

func closeDef(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}
