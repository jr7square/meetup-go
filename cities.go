package meetup

import (
  "net/http"
  "github.com/dghubble/sling"
)

type City struct {
  Zip string `json:"zip"`
  Country string `json:"country"`
  LocalizedCountryName string `json:"loalized_contry_name"`
  Distance float64 `json:"distance"`
  Name string `json:"city"`
  Longitude float64 `json:"lon"`
  Ranking int `json:"ranking"`
  Id int `json:"id"`
  State string `json:"state"`
  MemberCount string `json:"41033"`
  Latitude float64 `json:"lat"`
}

type CitiesResponse struct {
  Results []City `json:"results"`
  Meta ResponseMeta `json:"meta"`
}

type ReqCityParams struct {
  Country string `url:"country,omitempty"`
  Query string `url:"query,omitempty"`
  Longitude float64 `url:"lon,omitempty"`
  State string `url:"state,omitempty"`
  Radius float64 `url:"radius,omitempty"`
  Latitude float64 `url:"lat,omitempty"`
  Order string `url:"order,omitempty"`
  Page int `url:"page,omitempty"`
  Offset int `url:"offset,omitempty"`
  Desc string `url:"desc,omitempty"`
  Only string `url:"only,omitempty"`
  Omit string `url:"omit,omitempty"`
}

type CityService struct {
  sling *sling.Sling
}

func newCityService(sling *sling.Sling) *CityService {
  return &CityService{
    sling: sling.Path("2/cities"),
  }
}

func (cs *CityService) GetCities(params *ReqCityParams) (CitiesResponse, *http.Response, error) {
  cities := new(CitiesResponse)
  apiError := new(ApiError)
  resp, err := cs.sling.New().QueryStruct(params).Receive(cities, apiError)
  return *cities, resp, relevantError(err, *apiError)
}
