package trivec

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-trivec/utils"
)

func (c *Client) NewPaymodeGroupsRequest() PaymodeGroupsRequest {
	r := PaymodeGroupsRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type PaymodeGroupsRequest struct {
	client      *Client
	queryParams *PaymodeGroupsRequestQueryParams
	pathParams  *PaymodeGroupsRequestPathParams
	method      string
	headers     http.Header
	requestBody PaymodeGroupsRequestBody
}

func (r PaymodeGroupsRequest) NewQueryParams() *PaymodeGroupsRequestQueryParams {
	return &PaymodeGroupsRequestQueryParams{}
}

type PaymodeGroupsRequestQueryParams struct {
}

func (p PaymodeGroupsRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PaymodeGroupsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r PaymodeGroupsRequest) NewPathParams() *PaymodeGroupsRequestPathParams {
	return &PaymodeGroupsRequestPathParams{}
}

type PaymodeGroupsRequestPathParams struct{}

func (p *PaymodeGroupsRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PaymodeGroupsRequest) PathParams() *PaymodeGroupsRequestPathParams {
	return r.pathParams
}

func (r *PaymodeGroupsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PaymodeGroupsRequest) SetMethod(method string) {
	r.method = method
}

func (r *PaymodeGroupsRequest) Method() string {
	return r.method
}

func (r PaymodeGroupsRequest) NewRequestBody() PaymodeGroupsRequestBody {
	return PaymodeGroupsRequestBody{}
}

type PaymodeGroupsRequestBody struct {
}

func (r *PaymodeGroupsRequest) RequestBody() *PaymodeGroupsRequestBody {
	return nil
}

func (r *PaymodeGroupsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *PaymodeGroupsRequest) SetRequestBody(body PaymodeGroupsRequestBody) {
	r.requestBody = body
}

func (r *PaymodeGroupsRequest) NewResponseBody() *PaymodeGroupsRequestResponseBody {
	return &PaymodeGroupsRequestResponseBody{}
}

type PaymodeGroupsRequestResponseBody PaymodeGroups

func (r *PaymodeGroupsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("export/paymodegroups", r.PathParams())
	return &u
}

func (r *PaymodeGroupsRequest) Do() (PaymodeGroupsRequestResponseBody, error) {
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
