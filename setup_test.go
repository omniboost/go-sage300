package accountviewnet_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	accountviewnet "github.com/omniboost/go-accountview.new"
)

var (
	client *accountviewnet.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	subscriptionKey := os.Getenv("SUBSCRIPTION_KEY")
	serviceKey := os.Getenv("SERVICE_KEY")
	environment := os.Getenv("ENVIRONMENT")
	debug := os.Getenv("DEBUG")

	client = accountviewnet.NewClient(nil, subscriptionKey, serviceKey)
	if debug != "" {
		client.SetDebug(true)
	}
	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}
	if environment != "" {
		client.SetEnvironment(environment)
	}
	m.Run()
}
