package meetup

import (
  "net/http"
	"github.com/dghubble/sling"
)

type Status struct {
  Message string `json:"message"`
  Status string `json:"status"`
  Title string `json:"title"`
}

//refactor this to another file called client.go
type ResponseMeta struct {
  Next string `json:"next"`
  Method string `json:"categories"`
  TotalCount int `json:"total_count"`
  Link string `json:"link"`
  Count int `json:"count"`
  Description string `json:"description"`
  Lon string `json"lon"`
  Title string `json:"title"`
  Url string `json:"url"`
  SignedUrl string `json:"signed_url"`
  Id string `json:"id"`
  Updated string `json:"updated"`
  Lat string `json:"lat"`
}

type MetaService struct {
  sling *sling.Sling
  key string
}

func newMetaService(sling *sling.Sling, apiKey string) *MetaService {
  return &MetaService{
    sling: sling.Path("status?sign=true&key=" + apiKey),
    key: apiKey,
  }
}

func (ms *MetaService) Status() (*Status, *http.Response, error) {
  status := new(Status)
  apiError := new(ApiError)
  resp, err := ms.sling.New().Receive(status, apiError)
  return status, resp, relevantError(err, *apiError)
}

// func (ms *MetaService) Status() (*Status, error) {
// 	url := "https://api.meetup.com/status?key" + mc.key + "&sign=true"
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	resp, err := mc.client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
//
// 	body, err:= ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	status := new(Status)
// 	json.Unmarshal(body, status)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return status, err
// }
