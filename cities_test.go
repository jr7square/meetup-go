package meetup

import (
  "testing"
  "time"
  "github.com/dghubble/sling"
)

func TestCityService_GetCities(t *testing.T) {
  config := Init(time.Second, key)
  client := config.Client()
  cs := newCityService(sling.New().Client(client.client).Base("https://api.meetup.com/"))
  params := &ReqCityParams{}

  cities, resp, err := cs.GetCities(params)

  if cities.Results == nil {
    t.Errorf("cities.Results is nil")
  }

  if cities.Results[0].Name != "Raleigh" {
    t.Errorf("The first element of cities result is not Raleigh")
  }

  if resp == nil {
    t.Errorf("response is empty")
  }

  if err != nil {
    t.Errorf("there was an error with the request")
  }

  params2 := &ReqCityParams{
    Only: "city",
  }

  cities2, resp2, err2 := cs.GetCities(params2)

  if resp2 == nil {
    t.Errorf("response is empty")
  }

  if err2 != nil {
    t.Errorf("there was an error with the request")
  }

  if cities2.Results[0].Zip != "" {
    t.Errorf("zip is present is result, %s", cities2.Results[0].Zip)
  }

  if cities2.Results[0].Name == "" {
    t.Errorf("City name is not found")
  }

}
