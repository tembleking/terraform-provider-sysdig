package common

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

type SysdigCommonClient interface {
	CreateUser(User) (User, error)
	GetUserById(int) (User, error)
	DeleteUser(int) error
	UpdateUser(User) (User, error)
}

func NewSysdigCommonClient(sysdigAPIToken string, url string, insecure bool) SysdigCommonClient {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
		},
	}

	return &sysdigCommonClient{
		SysdigAPIToken: sysdigAPIToken,
		URL:            url,
		httpClient:     httpClient,
	}
}

type sysdigCommonClient struct {
	SysdigAPIToken string
	URL            string
	httpClient     *http.Client
}

func (client *sysdigCommonClient) doSysdigCommonRequest(method string, url string, payload io.Reader) (*http.Response, error) {
	request, _ := http.NewRequest(method, url, payload)
	request.Header.Set("Authorization", "Bearer "+client.SysdigAPIToken)
	request.Header.Set("Content-Type", "application/json")

	out, _ := httputil.DumpRequestOut(request, true)
	log.Printf("[DEBUG] %s", string(out))
	response, error := client.httpClient.Do(request)

	out, _ = httputil.DumpResponse(response, true)
	log.Printf("[DEBUG] %s", string(out))
	return response, error
}
