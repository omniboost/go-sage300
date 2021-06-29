package accountviewnet

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-accountview.new/utils"
)

func (c *Client) NewProductGroupsRequest() ProductGroupsRequest {
	r := ProductGroupsRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ProductGroupsRequest struct {
	client      *Client
	queryParams *ProductGroupsRequestQueryParams
	pathParams  *ProductGroupsRequestPathParams
	method      string
	headers     http.Header
	requestBody ProductGroupsRequestBody
}

func (r ProductGroupsRequest) NewQueryParams() *ProductGroupsRequestQueryParams {
	return &ProductGroupsRequestQueryParams{}
}

type ProductGroupsRequestQueryParams struct {
}

func (p ProductGroupsRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ProductGroupsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r ProductGroupsRequest) NewPathParams() *ProductGroupsRequestPathParams {
	return &ProductGroupsRequestPathParams{}
}

type ProductGroupsRequestPathParams struct{}

func (p *ProductGroupsRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ProductGroupsRequest) PathParams() *ProductGroupsRequestPathParams {
	return r.pathParams
}

func (r *ProductGroupsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ProductGroupsRequest) SetMethod(method string) {
	r.method = method
}

func (r *ProductGroupsRequest) Method() string {
	return r.method
}

func (r ProductGroupsRequest) NewRequestBody() ProductGroupsRequestBody {
	return ProductGroupsRequestBody{}
}

type ProductGroupsRequestBody struct {
}

func (r *ProductGroupsRequest) RequestBody() *ProductGroupsRequestBody {
	return nil
}

func (r *ProductGroupsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *ProductGroupsRequest) SetRequestBody(body ProductGroupsRequestBody) {
	r.requestBody = body
}

func (r *ProductGroupsRequest) NewResponseBody() *ProductGroupsRequestResponseBody {
	return &ProductGroupsRequestResponseBody{}
}

type ProductGroupsRequestResponseBody ProductGroups

func (r *ProductGroupsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("export/productgroups", r.PathParams())
	return &u
}

func (r *ProductGroupsRequest) Do() (ProductGroupsRequestResponseBody, error) {
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
