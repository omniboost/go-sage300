package netsuite

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"text/template"
	"time"

	"golang.org/x/oauth2"
)

const (
	scope                = "*"
	oauthStateString     = ""
	authorizationTimeout = 60 * time.Second
	tokenTimeout         = 5 * time.Second
)

type Oauth2Config struct {
	CompanyID string
	oauth2.Config
}

func NewOauthRoundTripper(rtp http.RoundTripper, tokenURL, companyID string, params url.Values) *OauthRoundTripper {
	return &OauthRoundTripper{rtp: rtp, tokenURL: tokenURL, params: params}
}

type OauthRoundTripper struct {
	rtp       http.RoundTripper
	tokenURL  string
	companyID string
	params    url.Values
}

func (t *OauthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.String() == t.tokenURL {
		tmpl, err := template.New("host").Parse(req.URL.Host)
		if err != nil {
			return nil, err
		}
		buf := new(bytes.Buffer)
		err = tmpl.Execute(buf, map[string]interface{}{"account_id": t.companyID})
		if err != nil {
			return nil, err
		}
		req.URL.Host = buf.String()

		q := req.URL.Query()
		for k, vv := range t.params {
			for _, v := range vv {
				q.Add(k, v)
			}
		}
		req.URL.RawQuery = q.Encode()
	}

	return http.DefaultTransport.RoundTrip(req)
}

func NewOauth2Config(companyID string) *Oauth2Config {
	config := &Oauth2Config{
		CompanyID: companyID,
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{scope},
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://{{.account_id}}.suitetalk.api.netsuite.com/services/rest/auth/oauth2/v1/token",
			},
		},
	}

	return config
}

func (c *Oauth2Config) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	params := url.Values{"company": []string{c.CompanyID}}
	rtp := NewOauthRoundTripper(http.DefaultTransport, c.Config.Endpoint.TokenURL, c.CompanyID, params)
	ctx = context.WithValue(ctx, oauth2.HTTPClient, &http.Client{Transport: rtp})
	return c.Config.Client(ctx, t)
}
