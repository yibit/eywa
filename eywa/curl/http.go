package curl

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

type proxyHandler func(*http.Request) (*url.URL, error)

// invokeService invoke specified service
func invokeService(method, url string, headers []string, body string) (*http.Response, error) {
	log.Debugf("%s %s %v %s", method, url, headers, body)
	request, err := http.NewRequest(method, url, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}

	for _, v := range headers {
		pairs := strings.Split(v, ":")
		if len(pairs) != 2 {
			continue
		}
		request.Header.Add(strings.TrimSpace(pairs[0]), strings.TrimSpace(pairs[1]))
	}
	setHeaderIfNotExist(request, "Accept-Charset", "UTF-8")
	setHeaderIfNotExist(request, "Accept", "application/json;charset=UTF-8")
	setHeaderIfNotExist(request, "Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{
		Timeout: time.Duration(60) * time.Second,
	}

	return client.Do(request)
}

func setHeaderIfNotExist(request *http.Request, header, value string) {
	if request.Header.Get(header) == "" {
		request.Header.Add(header, value)
	}
}

// invokeYService invoke Y Service
func InvokeYService(method, url string, headers []string, body string) bool {
	response, err := invokeService(method, url, headers, body)
	if err != nil {
		log.Infof("invokeService failed: %s", err.Error())
		return false
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Errorf("invokeService failed: %s", response.Status)
		io.Copy(io.Discard, response.Body)
		return false
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorf("invokeService failed: %v", err)
		return false
	}

	log.Warnf("%s", string(data))

	return true
}
