package trivec

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-trivec/utils"
)

func (c *Client) NewPaymodesRequest() PaymodesRequest {
	r := PaymodesRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type PaymodesRequest struct {
	client      *Client
	queryParams *PaymodesRequestQueryParams
	pathParams  *PaymodesRequestPathParams
	method      string
	headers     http.Header
	requestBody PaymodesRequestBody
}

func (r PaymodesRequest) NewQueryParams() *PaymodesRequestQueryParams {
	return &PaymodesRequestQueryParams{}
}

type PaymodesRequestQueryParams struct {
}

func (p PaymodesRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PaymodesRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r PaymodesRequest) NewPathParams() *PaymodesRequestPathParams {
	return &PaymodesRequestPathParams{}
}

type PaymodesRequestPathParams struct{}

func (p *PaymodesRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PaymodesRequest) PathParams() *PaymodesRequestPathParams {
	return r.pathParams
}

func (r *PaymodesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PaymodesRequest) SetMethod(method string) {
	r.method = method
}

func (r *PaymodesRequest) Method() string {
	return r.method
}

func (r PaymodesRequest) NewRequestBody() PaymodesRequestBody {
	return PaymodesRequestBody{}
}

type PaymodesRequestBody struct {
}

func (r *PaymodesRequest) RequestBody() *PaymodesRequestBody {
	return nil
}

func (r *PaymodesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *PaymodesRequest) SetRequestBody(body PaymodesRequestBody) {
	r.requestBody = body
}

func (r *PaymodesRequest) NewResponseBody() *PaymodesRequestResponseBody {
	return &PaymodesRequestResponseBody{}
}

type PaymodesRequestResponseBody Paymodes

func (r *PaymodesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("export/paymodes", r.PathParams())
	return &u
}

func (r *PaymodesRequest) Do() (PaymodesRequestResponseBody, error) {
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
