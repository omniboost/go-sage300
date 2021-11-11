package sage300

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-sage300/odata"
	"github.com/omniboost/go-sage300/utils"
)

func (c *Client) NewTaxGroupsGetRequest() TaxGroupsGetRequest {
	r := TaxGroupsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaxGroupsGetRequest struct {
	client      *Client
	queryParams *TaxGroupsGetRequestQueryParams
	pathParams  *TaxGroupsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody TaxGroupsGetRequestBody
}

func (r TaxGroupsGetRequest) NewQueryParams() *TaxGroupsGetRequestQueryParams {
	selectFields, _ := utils.Fields(&TaxGroup{})
	return &TaxGroupsGetRequestQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type TaxGroupsGetRequestQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p TaxGroupsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TaxGroupsGetRequest) QueryParams() *TaxGroupsGetRequestQueryParams {
	return r.queryParams
}

func (r TaxGroupsGetRequest) NewPathParams() *TaxGroupsGetRequestPathParams {
	return &TaxGroupsGetRequestPathParams{}
}

type TaxGroupsGetRequestPathParams struct {
}

func (p *TaxGroupsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TaxGroupsGetRequest) PathParams() *TaxGroupsGetRequestPathParams {
	return r.pathParams
}

func (r *TaxGroupsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TaxGroupsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *TaxGroupsGetRequest) Method() string {
	return r.method
}

func (r TaxGroupsGetRequest) NewRequestBody() TaxGroupsGetRequestBody {
	return TaxGroupsGetRequestBody{}
}

type TaxGroupsGetRequestBody struct {
}

func (r *TaxGroupsGetRequest) RequestBody() *TaxGroupsGetRequestBody {
	return nil
}

func (r *TaxGroupsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *TaxGroupsGetRequest) SetRequestBody(body TaxGroupsGetRequestBody) {
	r.requestBody = body
}

func (r *TaxGroupsGetRequest) NewResponseBody() *TaxGroupsGetResponseBody {
	return &TaxGroupsGetResponseBody{}
}

type TaxGroupsGetResponseBody struct {
	OdataContext  string    `json:"@odata.context"`
	Value         TaxGroups `json:"value"`
	OdataNextLink string    `json:"@odata.nextLink"`
}

func (r *TaxGroupsGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/v{{.api_version}}/{{.tenant_id}}/{{.company_id}}/TX/TXTaxGroups", r.PathParams())
	return &u, err
}

func (r *TaxGroupsGetRequest) Do() (TaxGroupsGetResponseBody, error) {
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

func (r *TaxGroupsGetRequest) All() (TaxGroupsGetResponseBody, error) {
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
