package rest

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Shreesha5102/ChallongeTool/mybackend/internal/utils/logger"
)

var (
	log = logger.GetLogger()
)

type NewRestInterface interface {
	SetBaseURL(string, string)
	SendRequest(*http.Request) (*http.Response, error)
}

type Client struct {
	client  *http.Client
	baseURL string
}

func NewRestClient(hostName string, timeout time.Duration) (NewRestInterface, error) {
	tr := &http.Transport{
		// #nosec G402
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: timeout}

	restClient := &Client{
		client:  client,
		baseURL: hostName,
	}
	return restClient, nil
}

func (client *Client) SetBaseURL(protocol, hostName string) {
	client.baseURL = fmt.Sprintf("%s://%s", protocol, hostName)
}

func (client *Client) SendRequest(request *http.Request) (*http.Response, error) {
	path := fmt.Sprintf("%s%s", client.baseURL, request.URL.String())
	u, err := url.Parse(path)
	if err != nil {
		log.Error("Error parsing url: ", err.Error())
		return nil, err
	}
	request.URL = u

	log.Debug("Request: ", request.Method, " ", request.URL.String())
	request.Header.Set("Content-Type", "application/json")

	response, err := client.client.Do(request)
	if err != nil {
		log.Error("Error in request : ", err.Error())
		return nil, err
	}

	// defer response.Body.Close() - DONT close here, close at callbacks once we are done with building msg
	log.Info("Request processed successfully with response code : ", response.Status)
	return response, nil
}
