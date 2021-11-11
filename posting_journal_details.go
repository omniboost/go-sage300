package sage300

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-sage300/odata"
	"github.com/omniboost/go-sage300/utils"
)

func (c *Client) NewPostingJournalDetailsGetRequest() PostingJournalDetailsGetRequest {
	r := PostingJournalDetailsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type PostingJournalDetailsGetRequest struct {
	client      *Client
	queryParams *PostingJournalDetailsGetRequestQueryParams
	pathParams  *PostingJournalDetailsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody PostingJournalDetailsGetRequestBody
}

func (r PostingJournalDetailsGetRequest) NewQueryParams() *PostingJournalDetailsGetRequestQueryParams {
	selectFields, _ := utils.Fields(&TaxGroup{})
	return &PostingJournalDetailsGetRequestQueryParams{
		Select: odata.NewSelect(selectFields),
		Filter: odata.NewFilter(),
		Top:    odata.NewTop(),
		Skip:   odata.NewSkip(),
	}
}

type PostingJournalDetailsGetRequestQueryParams struct {
	// @TODO: check if this an OData struct or something
	Select *odata.Select `schema:"$select,omitempty"`
	Filter *odata.Filter `schema:"$filter,omitempty"`
	Top    *odata.Top    `schema:"$top,omitempty"`
	Skip   *odata.Skip   `schema:"$skip,omitempty"`
}

func (p PostingJournalDetailsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostingJournalDetailsGetRequest) QueryParams() *PostingJournalDetailsGetRequestQueryParams {
	return r.queryParams
}

func (r PostingJournalDetailsGetRequest) NewPathParams() *PostingJournalDetailsGetRequestPathParams {
	return &PostingJournalDetailsGetRequestPathParams{}
}

type PostingJournalDetailsGetRequestPathParams struct {
}

func (p *PostingJournalDetailsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostingJournalDetailsGetRequest) PathParams() *PostingJournalDetailsGetRequestPathParams {
	return r.pathParams
}

func (r *PostingJournalDetailsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostingJournalDetailsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostingJournalDetailsGetRequest) Method() string {
	return r.method
}

func (r PostingJournalDetailsGetRequest) NewRequestBody() PostingJournalDetailsGetRequestBody {
	return PostingJournalDetailsGetRequestBody{}
}

type PostingJournalDetailsGetRequestBody struct {
}

func (r *PostingJournalDetailsGetRequest) RequestBody() *PostingJournalDetailsGetRequestBody {
	return nil
}

func (r *PostingJournalDetailsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *PostingJournalDetailsGetRequest) SetRequestBody(body PostingJournalDetailsGetRequestBody) {
	r.requestBody = body
}

func (r *PostingJournalDetailsGetRequest) NewResponseBody() *PostingJournalDetailsGetResponseBody {
	return &PostingJournalDetailsGetResponseBody{}
}

type PostingJournalDetailsGetResponseBody struct {
	OdataContext  string                `json:"@odata.context"`
	Value         PostingJournalDetails `json:"value"`
	OdataNextLink string                `json:"@odata.nextLink"`
}

func (r *PostingJournalDetailsGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/v{{.api_version}}/{{.tenant_id}}/{{.company_id}}/GL/GLPostingJournalDetails", r.PathParams())
	return &u, err
}

func (r *PostingJournalDetailsGetRequest) Do() (PostingJournalDetailsGetResponseBody, error) {
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

func (r *PostingJournalDetailsGetRequest) All() (PostingJournalDetailsGetResponseBody, error) {
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
