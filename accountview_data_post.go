package netsuite

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite/utils"
)

func (c *Client) NewNetsuiteDataPostRequest() NetsuiteDataPostRequest {
	r := NetsuiteDataPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type NetsuiteDataPostRequest struct {
	client      *Client
	queryParams *NetsuiteDataPostRequestQueryParams
	pathParams  *NetsuiteDataPostRequestPathParams
	method      string
	headers     http.Header
	requestBody NetsuiteDataPostRequestBody
}

func (r NetsuiteDataPostRequest) NewQueryParams() *NetsuiteDataPostRequestQueryParams {
	return &NetsuiteDataPostRequestQueryParams{}
}

type NetsuiteDataPostRequestQueryParams struct{}

func (p NetsuiteDataPostRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *NetsuiteDataPostRequest) QueryParams() *NetsuiteDataPostRequestQueryParams {
	return r.queryParams
}

func (r NetsuiteDataPostRequest) NewPathParams() *NetsuiteDataPostRequestPathParams {
	return &NetsuiteDataPostRequestPathParams{}
}

type NetsuiteDataPostRequestPathParams struct {
}

func (p *NetsuiteDataPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *NetsuiteDataPostRequest) PathParams() *NetsuiteDataPostRequestPathParams {
	return r.pathParams
}

func (r *NetsuiteDataPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *NetsuiteDataPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *NetsuiteDataPostRequest) Method() string {
	return r.method
}

func (r NetsuiteDataPostRequest) NewRequestBody() NetsuiteDataPostRequestBody {
	return NetsuiteDataPostRequestBody{}
}

type NetsuiteDataPostRequestBody struct {
	BookDate       string    `json:"BookDate"`
	BusinessObject string    `json:"BusinessObject"`
	Table          Table     `json:"Table"`
	TableData      TableData `json:"TableData"`
}

func (r *NetsuiteDataPostRequest) RequestBody() *NetsuiteDataPostRequestBody {
	return &r.requestBody
}

func (r *NetsuiteDataPostRequest) RequestBodyInterface() interface{} {
	return r.RequestBody()
}

func (r *NetsuiteDataPostRequest) SetRequestBody(body NetsuiteDataPostRequestBody) {
	r.requestBody = body
}

func (r *NetsuiteDataPostRequest) NewResponseBody() *NetsuiteDataPostResponseBody {
	return &NetsuiteDataPostResponseBody{}
}

type NetsuiteDataPostResponseBody struct {
	json.RawMessage
}

func (r *NetsuiteDataPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("netsuitedata", r.PathParams())
	return &u
}

func (r *NetsuiteDataPostRequest) Do() (NetsuiteDataPostResponseBody, error) {
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
