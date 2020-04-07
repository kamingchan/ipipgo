package ipipgo

import "net/http"

var client = http.DefaultClient

func SetClient(c *http.Client) {
	client = c
}

var header = func() http.Header {
	h := http.Header{}
	h.Set("Accept", "*/*")
	h.Set("Accept-Language", "zh-cn")
	h.Set("User-Agent", "BestTrace/Mac V1.30")
	return h
}()

func SetHeader(h http.Header) {
	header = h
}

func httpGet(url string) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header = header
	return client.Do(req)
}
