package fetcher

import (
	"io"
	"net/http"
)

//Fetch downloads the service tags manifest
func Fetch() ([]io.Reader, error) {
	urls := []string{
		"https://www.gstatic.com/ipranges/goog.json",
		"https://www.gstatic.com/ipranges/cloud.json",
	}

	var client http.Client
	var docs []io.Reader

	for _, u := range urls {
		doc, err := client.Get(u)
		if err != nil {
			docs = append(docs, doc.Body)
		}
	}

	return docs, nil
}
