package meetup

import (
	"io/ioutil"
	"net/http"
)

/*Client is a go client for the MeetUp Api */
type Client struct {
	client *http.Client
	key    string
}

/*Status executes a status request */
func (mc *Client) Status() ([]byte, error) {
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

	return ioutil.ReadAll(resp.Body)
}
