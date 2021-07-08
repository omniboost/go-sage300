package accountviewnet

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/omniboost/go-accountview.net/utils"
)

func (c *Client) NewAccountviewDataPostRequest() AccountviewDataPostRequest {
	r := AccountviewDataPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountviewDataPostRequest struct {
	client      *Client
	queryParams *AccountviewDataPostRequestQueryParams
	pathParams  *AccountviewDataPostRequestPathParams
	method      string
	headers     http.Header
	requestBody AccountviewDataPostRequestBody
}

func (r AccountviewDataPostRequest) NewQueryParams() *AccountviewDataPostRequestQueryParams {
	return &AccountviewDataPostRequestQueryParams{}
}

type AccountviewDataPostRequestQueryParams struct{}

func (p AccountviewDataPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AccountviewDataPostRequest) QueryParams() *AccountviewDataPostRequestQueryParams {
	return r.queryParams
}

func (r AccountviewDataPostRequest) NewPathParams() *AccountviewDataPostRequestPathParams {
	return &AccountviewDataPostRequestPathParams{}
}

type AccountviewDataPostRequestPathParams struct {
}

func (p *AccountviewDataPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccountviewDataPostRequest) PathParams() *AccountviewDataPostRequestPathParams {
	return r.pathParams
}

func (r *AccountviewDataPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountviewDataPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountviewDataPostRequest) Method() string {
	return r.method
}

func (r AccountviewDataPostRequest) NewRequestBody() AccountviewDataPostRequestBody {
	return AccountviewDataPostRequestBody{}
}

type AccountviewDataPostRequestBody struct {
	BookDate       string    `json:"BookDate"`
	BusinessObject string    `json:"BusinessObject"`
	Table          Table     `json:"Table"`
	TableData      TableData `json:"TableData"`
}

func (r *AccountviewDataPostRequest) RequestBody() *AccountviewDataPostRequestBody {
	return &r.requestBody
}

func (r *AccountviewDataPostRequest) RequestBodyInterface() interface{} {
	return r.RequestBody()
}

func (r *AccountviewDataPostRequest) SetRequestBody(body AccountviewDataPostRequestBody) {
	r.requestBody = body
}

func (r *AccountviewDataPostRequest) NewResponseBody() *AccountviewDataPostResponseBody {
	return &AccountviewDataPostResponseBody{}
}

type AccountviewDataPostResponseBody struct {
	json.RawMessage
}

func (r *AccountviewDataPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("accountviewdata", r.PathParams())
	return &u
}

func (r *AccountviewDataPostRequest) Do() (AccountviewDataPostResponseBody, error) {
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
