package ipipgo

import "net/http"

var client = http.DefaultClient

func SetClient(c *http.Client) {
	client = c
}

var header = func() http.Header {
	h := http.Header{}
	h.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	h.Add("Accept-Language", "en-US,en;q=0.9,zh-HK;q=0.8,zh;q=0.7,zh-CN;q=0.6,zh-TW;q=0.5,ja;q=0.4")
	h.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.36")
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
