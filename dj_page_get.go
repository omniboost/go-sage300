package accountviewnet

import "github.com/omniboost/go-accountview.net/utils"

func (c *Client) NewDjPageGetRequest() DjPageGetRequest {
	r := DjPageGetRequest{
		AccountviewDataGetRequest: c.NewAccountviewDataGetRequest(),
	}
	r.AccountviewDataGetRequest.QueryParams().BusinessObject = "DJ2P"
	return r
}

type DjPageGetRequest struct {
	AccountviewDataGetRequest
}

func (r *DjPageGetRequest) NewResponseBody() *DjPageGetResponseBody {
	return &DjPageGetResponseBody{}
}

type DjPageGetResponseBody struct {
	DjPage  []DjPage      `json:"DJ_PAGE"`
	DjLine  []DjLine      `json:"DJ_LINE"`
	UsrLink []interface{} `json:"USR_LINK"`
	PageWf  []interface{} `json:"PAGE_WF"`
}

func (r *DjPageGetRequest) Do() (DjPageGetResponseBody, error) {
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
