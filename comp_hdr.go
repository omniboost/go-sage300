package accountviewnet

import "github.com/omniboost/go-accountview.net/utils"

func (c *Client) NewCompHdrGetRequest() CompHdrGetRequest {
	r := CompHdrGetRequest{
		AccountviewDataGetRequest: c.NewAccountviewDataGetRequest(),
	}
	r.AccountviewDataGetRequest.QueryParams().BusinessObject = "CMP1"
	return r
}

type CompHdrGetRequest struct {
	AccountviewDataGetRequest
}

func (r *CompHdrGetRequest) NewResponseBody() *CompHdrGetResponseBody {
	return &CompHdrGetResponseBody{}
}

type CompHdrGetResponseBody struct {
	CompHdr []CompHdr `json:"COMP_HDR"`
}

func (r *CompHdrGetRequest) Do() (CompHdrGetResponseBody, error) {
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
