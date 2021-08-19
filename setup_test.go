package netsuite_test

import (
	"context"
	"os"
	"testing"

	netsuite "github.com/omniboost/go-netsuite"
	"golang.org/x/oauth2"
)

var (
	client *netsuite.Client
)

func TestMain(m *testing.M) {
	baseURL := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	refreshToken := os.Getenv("REFRESH_TOKEN")
	tokenURL := os.Getenv("TOKEN_URL")
	companyID := os.Getenv("COMPANY_ID")
	debug := os.Getenv("DEBUG")

	oauthConfig := netsuite.NewOauth2Config(companyID)
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	// b, _ := json.MarshalIndent(oauthConfig, "", "  ")
	// log.Fatal(string(b))

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background(), token)

	client = netsuite.NewClient(httpClient, companyID)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != "" {
		client.SetBaseURL(baseURL)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
