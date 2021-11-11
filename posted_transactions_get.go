package sage300

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-sage300/odata"
	"github.com/omniboost/go-sage300/utils"
)

func (c *Client) NewPostedTransactionsGetRequest() PostedTransactionsGetRequest {
	r := PostedTransactionsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type PostedTransactionsGetRequest struct {
	client      *Client
	queryParams *PostedTransactionsGetRequestQueryParams
	pathParams  *PostedTransactionsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody PostedTransactionsGetRequestBody
}

func (r PostedTransactionsGetRequest) NewQueryParams() *PostedTransactionsGetRequestQueryParams {
	selectFields, _ := utils.Fields(&TaxGroup{})
	return &PostedTransactionsGetRequestQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type PostedTransactionsGetRequestQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p PostedTransactionsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostedTransactionsGetRequest) QueryParams() *PostedTransactionsGetRequestQueryParams {
	return r.queryParams
}

func (r PostedTransactionsGetRequest) NewPathParams() *PostedTransactionsGetRequestPathParams {
	return &PostedTransactionsGetRequestPathParams{}
}

type PostedTransactionsGetRequestPathParams struct {
}

func (p *PostedTransactionsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostedTransactionsGetRequest) PathParams() *PostedTransactionsGetRequestPathParams {
	return r.pathParams
}

func (r *PostedTransactionsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostedTransactionsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostedTransactionsGetRequest) Method() string {
	return r.method
}

func (r PostedTransactionsGetRequest) NewRequestBody() PostedTransactionsGetRequestBody {
	return PostedTransactionsGetRequestBody{}
}

type PostedTransactionsGetRequestBody struct {
}

func (r *PostedTransactionsGetRequest) RequestBody() *PostedTransactionsGetRequestBody {
	return nil
}

func (r *PostedTransactionsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *PostedTransactionsGetRequest) SetRequestBody(body PostedTransactionsGetRequestBody) {
	r.requestBody = body
}

func (r *PostedTransactionsGetRequest) NewResponseBody() *PostedTransactionsGetResponseBody {
	return &PostedTransactionsGetResponseBody{}
}

type PostedTransactionsGetResponseBody struct {
	OdataContext  string             `json:"@odata.context"`
	Value         PostedTransactions `json:"value"`
	OdataNextLink string             `json:"@odata.nextLink"`
}

func (r *PostedTransactionsGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/v{{.api_version}}/{{.tenant_id}}/{{.company_id}}/GL/GLPostedTransactions", r.PathParams())
	return &u, err
}

func (r *PostedTransactionsGetRequest) Do() (PostedTransactionsGetResponseBody, error) {
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

func (r *PostedTransactionsGetRequest) All() (PostedTransactionsGetResponseBody, error) {
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
