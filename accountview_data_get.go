package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite/utils"
)

func (c *Client) NewNetsuiteDataGetRequest() NetsuiteDataGetRequest {
	r := NetsuiteDataGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type NetsuiteDataGetRequest struct {
	client      *Client
	queryParams *NetsuiteDataGetRequestQueryParams
	pathParams  *NetsuiteDataGetRequestPathParams
	method      string
	headers     http.Header
	requestBody NetsuiteDataGetRequestBody
}

func (r NetsuiteDataGetRequest) NewQueryParams() *NetsuiteDataGetRequestQueryParams {
	return &NetsuiteDataGetRequestQueryParams{}
}

type NetsuiteDataGetRequestQueryParams struct {
	BusinessObject string `schema:"BusinessObject"`
	PageSize       int    `schema:"PageSize"`
	Fields         string `schema:"Fields,omitempty"`

	FilterControlSource1  string `schema:"FilterControlSource1,omitempty"`
	FilterOperator1       string `schema:"FilterOperator1,omitempty"`
	FilterValueType1      string `schema:"FilterValueType1,omitempty"`
	FilterValue1          string `schema:"FilterValue1,omitempty"`
	FilterIsListOfValues1 bool   `schema:"FilterIsListOfValues1"`

	FilterControlSource2  string `schema:"FilterControlSource2,omitempty"`
	FilterOperator2       string `schema:"FilterOperator2,omitempty"`
	FilterValueType2      string `schema:"FilterValueType2,omitempty"`
	FilterValue2          string `schema:"FilterValue2,omitempty"`
	FilterIsListOfValues2 bool   `schema:"FilterIsListOfValues2,omitempty"`

	SortFields string `schema:"SortFields,omitempty"`
	SortOrder  string `schema:"SortOrder,omitempty"`
}

func (p NetsuiteDataGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *NetsuiteDataGetRequest) QueryParams() *NetsuiteDataGetRequestQueryParams {
	return r.queryParams
}

func (r NetsuiteDataGetRequest) NewPathParams() *NetsuiteDataGetRequestPathParams {
	return &NetsuiteDataGetRequestPathParams{}
}

type NetsuiteDataGetRequestPathParams struct {
}

func (p *NetsuiteDataGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *NetsuiteDataGetRequest) PathParams() *NetsuiteDataGetRequestPathParams {
	return r.pathParams
}

func (r *NetsuiteDataGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *NetsuiteDataGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *NetsuiteDataGetRequest) Method() string {
	return r.method
}

func (r NetsuiteDataGetRequest) NewRequestBody() NetsuiteDataGetRequestBody {
	return NetsuiteDataGetRequestBody{}
}

type NetsuiteDataGetRequestBody struct {
}

func (r *NetsuiteDataGetRequest) RequestBody() *NetsuiteDataGetRequestBody {
	return nil
}

func (r *NetsuiteDataGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *NetsuiteDataGetRequest) SetRequestBody(body NetsuiteDataGetRequestBody) {
	r.requestBody = body
}

func (r *NetsuiteDataGetRequest) NewResponseBody() *NetsuiteDataGetResponseBody {
	return &NetsuiteDataGetResponseBody{}
}

type NetsuiteDataGetResponseBody struct{}

func (r *NetsuiteDataGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("netsuitedata", r.PathParams())
	return &u
}

func (r *NetsuiteDataGetRequest) Do() (NetsuiteDataGetResponseBody, error) {
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
