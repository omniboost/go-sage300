package sage300

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-sage300/odata"
	"github.com/omniboost/go-sage300/utils"
)

func (c *Client) NewGLAccountsGetRequest() GLAccountsGetRequest {
	r := GLAccountsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GLAccountsGetRequest struct {
	client      *Client
	queryParams *GLAccountsGetRequestQueryParams
	pathParams  *GLAccountsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody GLAccountsGetRequestBody
}

func (r GLAccountsGetRequest) NewQueryParams() *GLAccountsGetRequestQueryParams {
	selectFields, _ := utils.Fields(&GLAccount{})
	return &GLAccountsGetRequestQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type GLAccountsGetRequestQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p GLAccountsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GLAccountsGetRequest) QueryParams() *GLAccountsGetRequestQueryParams {
	return r.queryParams
}

func (r GLAccountsGetRequest) NewPathParams() *GLAccountsGetRequestPathParams {
	return &GLAccountsGetRequestPathParams{}
}

type GLAccountsGetRequestPathParams struct {
}

func (p *GLAccountsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GLAccountsGetRequest) PathParams() *GLAccountsGetRequestPathParams {
	return r.pathParams
}

func (r *GLAccountsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GLAccountsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *GLAccountsGetRequest) Method() string {
	return r.method
}

func (r GLAccountsGetRequest) NewRequestBody() GLAccountsGetRequestBody {
	return GLAccountsGetRequestBody{}
}

type GLAccountsGetRequestBody struct {
}

func (r *GLAccountsGetRequest) RequestBody() *GLAccountsGetRequestBody {
	return nil
}

func (r *GLAccountsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GLAccountsGetRequest) SetRequestBody(body GLAccountsGetRequestBody) {
	r.requestBody = body
}

func (r *GLAccountsGetRequest) NewResponseBody() *GLAccountsGetResponseBody {
	return &GLAccountsGetResponseBody{}
}

type GLAccountsGetResponseBody struct {
	OdataContext  string     `json:"@odata.context"`
	Value         GLAccounts `json:"value"`
	OdataNextLink string     `json:"@odata.nextLink"`
}

func (r *GLAccountsGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/v{{.api_version}}/{{.tenant_id}}/{{.company_id}}/GL/GLAccounts", r.PathParams())
	return &u, err
}

func (r *GLAccountsGetRequest) Do() (GLAccountsGetResponseBody, error) {
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

func (r *GLAccountsGetRequest) All() (GLAccountsGetResponseBody, error) {
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
