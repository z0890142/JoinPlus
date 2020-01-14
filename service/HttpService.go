package service

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func HttpRequestWithQuery(action string, url string, header map[string]string, query map[string]string, Log *logrus.Entry) (resp *http.Response, err error) {

	var req *http.Request

	req, err = http.NewRequest(action, url, nil)
	queryParams := req.URL.Query()
	for key, val := range query {
		queryParams.Add(key, val)
	}
	req.URL.RawQuery = queryParams.Encode()

	for key, value := range header {
		req.Header.Set(key, value)
	}
	httpClient := http.Client{
		Timeout: 15 * time.Second,
		// Transport: &http.Transport{
		// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		// },
	}
	resp, err = httpClient.Do(req)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"action": action,
			"url":    url,
			"header": header,
			"query":  query,
			"error":  err,
		}).Error("http client do error")
		return
	}
	return
}
