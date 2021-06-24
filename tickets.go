package trivec

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-trivec/utils"
)

func (c *Client) NewTicketsRequest() TicketsRequest {
	r := TicketsRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TicketsRequest struct {
	client      *Client
	queryParams *TicketsRequestQueryParams
	pathParams  *TicketsRequestPathParams
	method      string
	headers     http.Header
	requestBody TicketsRequestBody
}

func (r TicketsRequest) NewQueryParams() *TicketsRequestQueryParams {
	return &TicketsRequestQueryParams{}
}

type TicketsRequestQueryParams struct {
}

func (p TicketsRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TicketsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r TicketsRequest) NewPathParams() *TicketsRequestPathParams {
	return &TicketsRequestPathParams{}
}

type TicketsRequestPathParams struct {
	Date Date
}

func (p *TicketsRequestPathParams) Params() map[string]string {
	return map[string]string{
		"date": p.Date.Format("20060102"),
	}
}

func (r *TicketsRequest) PathParams() *TicketsRequestPathParams {
	return r.pathParams
}

func (r *TicketsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TicketsRequest) SetMethod(method string) {
	r.method = method
}

func (r *TicketsRequest) Method() string {
	return r.method
}

func (r TicketsRequest) NewRequestBody() TicketsRequestBody {
	return TicketsRequestBody{}
}

type TicketsRequestBody struct {
}

func (r *TicketsRequest) RequestBody() *TicketsRequestBody {
	return nil
}

func (r *TicketsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *TicketsRequest) SetRequestBody(body TicketsRequestBody) {
	r.requestBody = body
}

func (r *TicketsRequest) NewResponseBody() *TicketsRequestResponseBody {
	return &TicketsRequestResponseBody{}
}

type TicketsRequestResponseBody Tickets

func (r *TicketsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("export/tickets/day/{{.date}}", r.PathParams())
	return &u
}

func (r *TicketsRequest) Do() (TicketsRequestResponseBody, error) {
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
