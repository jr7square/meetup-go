package meetup

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
)

/*Client is a go client for the MeetUp Api */
type Client struct {
	client *http.Client
	key    string
}

/*Status executes a status request */
func (mc *Client) Status() (*Status, error) {
	url := "https://api.meetup.com/status?key" + mc.key + "&sign=true"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := mc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err:= ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	status := new(Status)
	json.Unmarshal(body, status)
	if err != nil {
		return nil, err
	}
	return status, err
}
