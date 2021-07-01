package accountviewnet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-accountview.net/utils"
)

func (c *Client) NewAccountviewDataGetRequest() AccountviewDataGetRequest {
	r := AccountviewDataGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountviewDataGetRequest struct {
	client      *Client
	queryParams *AccountviewDataGetRequestQueryParams
	pathParams  *AccountviewDataGetRequestPathParams
	method      string
	headers     http.Header
	requestBody AccountviewDataGetRequestBody
}

func (r AccountviewDataGetRequest) NewQueryParams() *AccountviewDataGetRequestQueryParams {
	return &AccountviewDataGetRequestQueryParams{}
}

type AccountviewDataGetRequestQueryParams struct {
	BusinessObject string `schema:"BusinessObject"`
	PageSize       int    `schema:"PageSize"`

	FilterControlSource1  string `schema:"FilterControlSource1,omitempty"`
	FilterOperator1       string `schema:"FilterOperator1,omitempty"`
	FilterValueType1      string `schema:"FilterValueType1,omitempty"`
	FilterValue1          string `schema:"FilterValue1,omitempty"`
	FilterIsListOfValues1 bool   `schema:"FilterIsListOfValues1,omitempty"`

	FilterControlSource2  string `schema:"FilterControlSource2,omitempty"`
	FilterOperator2       string `schema:"FilterOperator2,omitempty"`
	FilterValueType2      string `schema:"FilterValueType2,omitempty"`
	FilterValue2          string `schema:"FilterValue2,omitempty"`
	FilterIsListOfValues2 bool   `schema:"FilterIsListOfValues2,omitempty"`
}

func (p AccountviewDataGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AccountviewDataGetRequest) QueryParams() *AccountviewDataGetRequestQueryParams {
	return r.queryParams
}

func (r AccountviewDataGetRequest) NewPathParams() *AccountviewDataGetRequestPathParams {
	return &AccountviewDataGetRequestPathParams{}
}

type AccountviewDataGetRequestPathParams struct {
}

func (p *AccountviewDataGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccountviewDataGetRequest) PathParams() *AccountviewDataGetRequestPathParams {
	return r.pathParams
}

func (r *AccountviewDataGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountviewDataGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountviewDataGetRequest) Method() string {
	return r.method
}

func (r AccountviewDataGetRequest) NewRequestBody() AccountviewDataGetRequestBody {
	return AccountviewDataGetRequestBody{}
}

type AccountviewDataGetRequestBody struct {
}

func (r *AccountviewDataGetRequest) RequestBody() *AccountviewDataGetRequestBody {
	return nil
}

func (r *AccountviewDataGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *AccountviewDataGetRequest) SetRequestBody(body AccountviewDataGetRequestBody) {
	r.requestBody = body
}

func (r *AccountviewDataGetRequest) NewResponseBody() *AccountviewDataGetResponseBody {
	return &AccountviewDataGetResponseBody{}
}

type AccountviewDataGetResponseBody struct{}

func (r *AccountviewDataGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("accountviewdata", r.PathParams())
	return &u
}

func (r *AccountviewDataGetRequest) Do() (AccountviewDataGetResponseBody, error) {
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
