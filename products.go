package trivec

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-trivec/utils"
)

func (c *Client) NewProductsRequest() ProductsRequest {
	r := ProductsRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ProductsRequest struct {
	client      *Client
	queryParams *ProductsRequestQueryParams
	pathParams  *ProductsRequestPathParams
	method      string
	headers     http.Header
	requestBody ProductsRequestBody
}

func (r ProductsRequest) NewQueryParams() *ProductsRequestQueryParams {
	return &ProductsRequestQueryParams{}
}

type ProductsRequestQueryParams struct {
}

func (p ProductsRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ProductsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r ProductsRequest) NewPathParams() *ProductsRequestPathParams {
	return &ProductsRequestPathParams{}
}

type ProductsRequestPathParams struct {
}

func (p *ProductsRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ProductsRequest) PathParams() *ProductsRequestPathParams {
	return r.pathParams
}

func (r *ProductsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ProductsRequest) SetMethod(method string) {
	r.method = method
}

func (r *ProductsRequest) Method() string {
	return r.method
}

func (r ProductsRequest) NewRequestBody() ProductsRequestBody {
	return ProductsRequestBody{}
}

type ProductsRequestBody struct {
}

func (r *ProductsRequest) RequestBody() *ProductsRequestBody {
	return nil
}

func (r *ProductsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *ProductsRequest) SetRequestBody(body ProductsRequestBody) {
	r.requestBody = body
}

func (r *ProductsRequest) NewResponseBody() *ProductsRequestResponseBody {
	return &ProductsRequestResponseBody{}
}

type ProductsRequestResponseBody Products

func (r *ProductsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("export/products", r.PathParams())
	return &u
}

func (r *ProductsRequest) Do() (ProductsRequestResponseBody, error) {
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
