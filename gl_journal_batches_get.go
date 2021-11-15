package sage300

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-sage300/odata"
	"github.com/omniboost/go-sage300/utils"
)

func (c *Client) NewGLJournalBatchGetRequest() GLJournalBatchGetRequest {
	r := GLJournalBatchGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GLJournalBatchGetRequest struct {
	client      *Client
	queryParams *GLJournalBatchGetRequestQueryParams
	pathParams  *GLJournalBatchGetRequestPathParams
	method      string
	headers     http.Header
	requestBody GLJournalBatchGetRequestBody
}

func (r GLJournalBatchGetRequest) NewQueryParams() *GLJournalBatchGetRequestQueryParams {
	selectFields, _ := utils.Fields(&GLAccount{})
	return &GLJournalBatchGetRequestQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type GLJournalBatchGetRequestQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p GLJournalBatchGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GLJournalBatchGetRequest) QueryParams() *GLJournalBatchGetRequestQueryParams {
	return r.queryParams
}

func (r GLJournalBatchGetRequest) NewPathParams() *GLJournalBatchGetRequestPathParams {
	return &GLJournalBatchGetRequestPathParams{}
}

type GLJournalBatchGetRequestPathParams struct {
}

func (p *GLJournalBatchGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GLJournalBatchGetRequest) PathParams() *GLJournalBatchGetRequestPathParams {
	return r.pathParams
}

func (r *GLJournalBatchGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GLJournalBatchGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *GLJournalBatchGetRequest) Method() string {
	return r.method
}

func (r GLJournalBatchGetRequest) NewRequestBody() GLJournalBatchGetRequestBody {
	return GLJournalBatchGetRequestBody{}
}

type GLJournalBatchGetRequestBody struct {
}

func (r *GLJournalBatchGetRequest) RequestBody() *GLJournalBatchGetRequestBody {
	return nil
}

func (r *GLJournalBatchGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GLJournalBatchGetRequest) SetRequestBody(body GLJournalBatchGetRequestBody) {
	r.requestBody = body
}

func (r *GLJournalBatchGetRequest) NewResponseBody() *GLJournalBatchGetResponseBody {
	return &GLJournalBatchGetResponseBody{}
}

type GLJournalBatchGetResponseBody struct {
	OdataContext  string     `json:"@odata.context"`
	Value         GLAccounts `json:"value"`
	OdataNextLink string     `json:"@odata.nextLink"`
}

func (r *GLJournalBatchGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/v{{.api_version}}/{{.tenant_id}}/{{.company_id}}/GL/GLJournalBatches", r.PathParams())
	return &u, err
}

func (r *GLJournalBatchGetRequest) Do() (GLJournalBatchGetResponseBody, error) {
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

func (r *GLJournalBatchGetRequest) All() (GLJournalBatchGetResponseBody, error) {
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
