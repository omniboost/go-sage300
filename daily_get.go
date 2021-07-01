package accountviewnet

import "github.com/omniboost/go-accountview.net/utils"

func (c *Client) NewDailyGetRequest() DailyGetRequest {
	r := DailyGetRequest{
		AccountviewDataGetRequest: c.NewAccountviewDataGetRequest(),
	}
	r.AccountviewDataGetRequest.QueryParams().BusinessObject = "DJ1"
	return r
}

type DailyGetRequest struct {
	AccountviewDataGetRequest
}

func (r *DailyGetRequest) NewResponseBody() *DailyGetResponseBody {
	return &DailyGetResponseBody{}
}

type DailyGetResponseBody struct {
	Daily []Daily `json:"DAILY"`
}

func (r *DailyGetRequest) Do() (DailyGetResponseBody, error) {
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
