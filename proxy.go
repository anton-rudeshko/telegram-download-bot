package main

import (
	"log"
	"net/http"
	"net/url"
)

func MakeHttpClient(proxyStr string) *http.Client {
	if proxyStr == "" {
		return &http.Client{}
	}

	proxyUrl, err := url.Parse(proxyStr)
	if err != nil {
		log.Fatalf(`Could not parse proxy url "%s": %s`, proxyStr, err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	log.Printf(`Using proxy "%s"`, proxyStr)

	return &http.Client{
		Transport: transport,
	}
}
