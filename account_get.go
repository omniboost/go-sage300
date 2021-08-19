package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite/utils"
)

func (c *Client) NewAccountGetRequest() AccountGetRequest {
	r := AccountGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountGetRequest struct {
	client      *Client
	queryParams *AccountGetRequestQueryParams
	pathParams  *AccountGetRequestPathParams
	method      string
	headers     http.Header
	requestBody AccountGetRequestBody
}

func (r AccountGetRequest) NewQueryParams() *AccountGetRequestQueryParams {
	return &AccountGetRequestQueryParams{}
}

type AccountGetRequestQueryParams struct {
	// Account string `schema:"Account"`
	// PageSize       int    `schema:"PageSize"`
}

func (p AccountGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	// encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	// encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountGetRequest) QueryParams() *AccountGetRequestQueryParams {
	return r.queryParams
}

func (r AccountGetRequest) NewPathParams() *AccountGetRequestPathParams {
	return &AccountGetRequestPathParams{}
}

type AccountGetRequestPathParams struct {
}

func (p *AccountGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccountGetRequest) PathParams() *AccountGetRequestPathParams {
	return r.pathParams
}

func (r *AccountGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountGetRequest) Method() string {
	return r.method
}

func (r AccountGetRequest) NewRequestBody() AccountGetRequestBody {
	return AccountGetRequestBody{}
}

type AccountGetRequestBody struct {
}

func (r *AccountGetRequest) RequestBody() *AccountGetRequestBody {
	return nil
}

func (r *AccountGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *AccountGetRequest) SetRequestBody(body AccountGetRequestBody) {
	r.requestBody = body
}

func (r *AccountGetRequest) NewResponseBody() *AccountGetResponseBody {
	return &AccountGetResponseBody{}
}

type AccountGetResponseBody struct{}

func (r *AccountGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/account", r.PathParams())
	return &u, err
}

func (r *AccountGetRequest) Do() (AccountGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
