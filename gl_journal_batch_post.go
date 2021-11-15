package sage300

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-sage300/odata"
	"github.com/omniboost/go-sage300/utils"
)

func (c *Client) NewGLJournalBatchPostRequest() GLJournalBatchPostRequest {
	r := GLJournalBatchPostRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GLJournalBatchPostRequest struct {
	client      *Client
	queryParams *GLJournalBatchPostRequestQueryParams
	pathParams  *GLJournalBatchPostRequestPathParams
	method      string
	headers     http.Header
	requestBody GLJournalBatchPostRequestBody
}

func (r GLJournalBatchPostRequest) NewQueryParams() *GLJournalBatchPostRequestQueryParams {
	selectFields, _ := utils.Fields(&GLAccount{})
	return &GLJournalBatchPostRequestQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type GLJournalBatchPostRequestQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p GLJournalBatchPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GLJournalBatchPostRequest) QueryParams() *GLJournalBatchPostRequestQueryParams {
	return r.queryParams
}

func (r GLJournalBatchPostRequest) NewPathParams() *GLJournalBatchPostRequestPathParams {
	return &GLJournalBatchPostRequestPathParams{}
}

type GLJournalBatchPostRequestPathParams struct {
}

func (p *GLJournalBatchPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GLJournalBatchPostRequest) PathParams() *GLJournalBatchPostRequestPathParams {
	return r.pathParams
}

func (r *GLJournalBatchPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GLJournalBatchPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *GLJournalBatchPostRequest) Method() string {
	return r.method
}

func (r GLJournalBatchPostRequest) NewRequestBody() GLJournalBatchPostRequestBody {
	return GLJournalBatchPostRequestBody{}
}

type GLJournalBatchPostRequestBody struct {
}

func (r *GLJournalBatchPostRequest) RequestBody() *GLJournalBatchPostRequestBody {
	return nil
}

func (r *GLJournalBatchPostRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GLJournalBatchPostRequest) SetRequestBody(body GLJournalBatchPostRequestBody) {
	r.requestBody = body
}

func (r *GLJournalBatchPostRequest) NewResponseBody() *GLJournalBatchPostResponseBody {
	return &GLJournalBatchPostResponseBody{}
}

type GLJournalBatchPostResponseBody struct {
	OdataContext  string     `json:"@odata.context"`
	Value         GLAccounts `json:"value"`
	OdataNextLink string     `json:"@odata.nextLink"`
}

func (r *GLJournalBatchPostRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/v{{.api_version}}/{{.tenant_id}}/{{.company_id}}/GL/GLAccounts", r.PathParams())
	return &u, err
}

func (r *GLJournalBatchPostRequest) Do() (GLJournalBatchPostResponseBody, error) {
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

func (r *GLJournalBatchPostRequest) All() (GLJournalBatchPostResponseBody, error) {
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
