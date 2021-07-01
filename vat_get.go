package accountviewnet

import "github.com/omniboost/go-accountview.net/utils"

func (c *Client) NewVATGetRequest() VATGetRequest {
	r := VATGetRequest{
		AccountviewDataGetRequest: c.NewAccountviewDataGetRequest(),
	}
	r.AccountviewDataGetRequest.QueryParams().BusinessObject = "VA1"
	return r
}

type VATGetRequest struct {
	AccountviewDataGetRequest
}

func (r *VATGetRequest) NewResponseBody() *VATGetResponseBody {
	return &VATGetResponseBody{}
}

type VATGetResponseBody struct {
	VAT []VAT `json:"VAT"`
}

func (r *VATGetRequest) Do() (VATGetResponseBody, error) {
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
