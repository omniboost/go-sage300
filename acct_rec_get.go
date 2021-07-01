package accountviewnet

import "github.com/omniboost/go-accountview.net/utils"

func (c *Client) NewAcctRecGetRequest() AcctRecGetRequest {
	r := AcctRecGetRequest{
		AccountviewDataGetRequest: c.NewAccountviewDataGetRequest(),
	}
	r.AccountviewDataGetRequest.QueryParams().BusinessObject = "AR1"
	return r
}

type AcctRecGetRequest struct {
	AccountviewDataGetRequest
}

func (r *AcctRecGetRequest) NewResponseBody() *AcctRecGetResponseBody {
	return &AcctRecGetResponseBody{}
}

type AcctRecGetResponseBody struct {
	AcctRec []AcctRec `json:"ACCTREC"`
}

func (r *AcctRecGetRequest) Do() (AcctRecGetResponseBody, error) {
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
