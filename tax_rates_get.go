package sage300

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-sage300/odata"
	"github.com/omniboost/go-sage300/utils"
)

func (c *Client) NewTaxRatesGetRequest() TaxRatesGetRequest {
	r := TaxRatesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaxRatesGetRequest struct {
	client      *Client
	queryParams *TaxRatesGetRequestQueryParams
	pathParams  *TaxRatesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody TaxRatesGetRequestBody
}

func (r TaxRatesGetRequest) NewQueryParams() *TaxRatesGetRequestQueryParams {
	selectFields, _ := utils.Fields(&TaxRate{})
	return &TaxRatesGetRequestQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type TaxRatesGetRequestQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p TaxRatesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TaxRatesGetRequest) QueryParams() *TaxRatesGetRequestQueryParams {
	return r.queryParams
}

func (r TaxRatesGetRequest) NewPathParams() *TaxRatesGetRequestPathParams {
	return &TaxRatesGetRequestPathParams{}
}

type TaxRatesGetRequestPathParams struct {
}

func (p *TaxRatesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TaxRatesGetRequest) PathParams() *TaxRatesGetRequestPathParams {
	return r.pathParams
}

func (r *TaxRatesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TaxRatesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *TaxRatesGetRequest) Method() string {
	return r.method
}

func (r TaxRatesGetRequest) NewRequestBody() TaxRatesGetRequestBody {
	return TaxRatesGetRequestBody{}
}

type TaxRatesGetRequestBody struct {
}

func (r *TaxRatesGetRequest) RequestBody() *TaxRatesGetRequestBody {
	return nil
}

func (r *TaxRatesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *TaxRatesGetRequest) SetRequestBody(body TaxRatesGetRequestBody) {
	r.requestBody = body
}

func (r *TaxRatesGetRequest) NewResponseBody() *TaxRatesGetResponseBody {
	return &TaxRatesGetResponseBody{}
}

type TaxRatesGetResponseBody struct {
	OdataContext  string   `json:"@odata.context"`
	Value         TaxRates `json:"value"`
	OdataNextLink string   `json:"@odata.nextLink"`
}

func (r *TaxRatesGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/v{{.api_version}}/{{.tenant_id}}/{{.company_id}}/TX/TXTaxRates", r.PathParams())
	return &u, err
}

func (r *TaxRatesGetRequest) Do() (TaxRatesGetResponseBody, error) {
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

func (r *TaxRatesGetRequest) All() (TaxRatesGetResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := resp.Value

	for resp.OdataNextLink != "" {
		u, err := url.Parse(resp.OdataNextLink)
		if err != nil {
			return resp, err
		}

		skip, err := strconv.Atoi(u.Query().Get("$skip"))
		if err != nil {
			return resp, err
		}

		r.QueryParams().Skip.Set(skip)
		resp, err = r.Do()
		if err != nil {
			return resp, err
		}

		concat = append(concat, resp.Value...)
	}

	resp.Value = concat
	return resp, nil
}
