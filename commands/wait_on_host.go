package commands

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/tylergets/workon/logger"
	"net/http"
	"time"
)

type WaitOnHost struct{}

func (c *WaitOnHost) Execute(log *logger.Logger, data map[string]string) error {
	url := data["url"]
	if url == "" {
		log.Error("No URL specified.")
		return errors.New("no URL specified")
	}
	log.Info("Waiting on URL: %s", url)

	return waitForHTTP(log, url)
}

func waitForHTTP(log *logger.Logger, url string) error {
	// Creating a custom HTTP client to ignore SSL errors
	httpClient := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	for {
		resp, err := httpClient.Get(url)
		if err != nil {
			log.Info("HTTP service at URL %s not ready, retrying...", url)
			// print the error
			fmt.Println(err.Error())
			time.Sleep(time.Second)
			continue
		}
		if resp.StatusCode == 200 {
			log.Info("HTTP service at URL %s is now ready with 200 OK", url)
			resp.Body.Close()
			break
		}
		resp.Body.Close()
		log.Info("HTTP service at URL %s not ready with status code: %d, retrying...", url, resp.StatusCode)
		time.Sleep(time.Second)
	}
	return nil
}

func init() {
	Register("wait-on-host", &WaitOnHost{})
}
