package healthchecksio

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type HealthcheckAction string

const (
	// Indicates that the job has started. Requires [HealthcheckSuccess] afterwards to not be considered failed
	HealthcheckStart   HealthcheckAction = "/start"
	HealthcheckSuccess HealthcheckAction = ""
	HealthcheckFail    HealthcheckAction = "/fail"
)

// Send a healthcheck request to healthcheck.io
//
// `pid` is an optional parameter useful to match Start and Finish / Fail healthchecks on the website by
func Healthcheck(url string, action HealthcheckAction, pid string) error {
	url = url + string(action)
	if pid != "" {
		url += "?rid=" + pid
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while sending a healthcheck request (action: %s): %w", action, err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error while reading a healthcheck start response (action: %s): %w", action, err)
	}

	err = handleHealthchecksResponse(res.StatusCode, string(body))
	if err != nil {
		return fmt.Errorf("error while handling a healthcheck response (action: %s): %w. Status: %d, Body: %s",
			action, err, res.StatusCode, string(body))
	}

	return nil
}

func handleHealthchecksResponse(status int, body string) error {
	switch {
	case status == 200 && body == "OK":
		return nil
	case status == 200 && body == "OK (not found)":
		return errors.New("unknown check ID")
	case status == 200 && body == "OK (rate limited)":
		return errors.New("rate limited")
	case status == 400 && body == "invalid url format":
		return errors.New("invalid URL format")
	default:
		return errors.New("unknown response")
	}
}
