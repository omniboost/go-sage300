package netsuite

import "github.com/omniboost/go-netsuite/utils"

func (c *Client) NewAcctRecGetRequest() AcctRecGetRequest {
	r := AcctRecGetRequest{
		NetsuiteDataGetRequest: c.NewNetsuiteDataGetRequest(),
	}
	r.NetsuiteDataGetRequest.QueryParams().BusinessObject = "AR1"
	return r
}

type AcctRecGetRequest struct {
	NetsuiteDataGetRequest
}

func (r *AcctRecGetRequest) NewResponseBody() *AcctRecGetResponseBody {
	return &AcctRecGetResponseBody{}
}

type AcctRecGetResponseBody struct {
	AcctRec []AcctRec     `json:"CONTACT"`
	UsrLink []interface{} `json:"USR_LINK"`
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
