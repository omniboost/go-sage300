package sage300_test

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	sage300 "github.com/omniboost/go-sage300"
	"golang.org/x/net/proxy"
)

var (
	client *sage300.Client
)

func TestMain(m *testing.M) {
	baseURL := os.Getenv("SAGE_BASE_URL")
	username := os.Getenv("SAGE_USERNAME")
	password := os.Getenv("SAGE_PASSWORD")
	companyID := os.Getenv("SAGE_COMPANY_ID")
	debug := os.Getenv("SAGE_DEBUG")

	client = sage300.NewClient(nil, username, password, companyID)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != "" {
		client.SetBaseURL(baseURL)
	}

	client.SetDisallowUnknownFields(true)

	allProxy := os.Getenv("ALL_PROXY")
	if allProxy != "" {
		proxyUrl, err := url.Parse(allProxy)
		if err != nil {
			log.Fatal(err)
		}

		addr := fmt.Sprintf("%s:%s", proxyUrl.Hostname(), proxyUrl.Port())

		dialSocksProxy, err := proxy.SOCKS5("tcp", addr, nil, proxy.Direct)
		if err != nil {
			fmt.Println("Error connecting to proxy:", err)
		}
		tr := &http.Transport{Dial: dialSocksProxy.Dial}

		httpClient := &http.Client{
			Transport: tr,
			// Timeout: 30 * time.Second,
		}

		client.SetHTTPClient(httpClient)
	}

	m.Run()
}
