package meetup

import (
  "testing"
  "time"
	"github.com/dghubble/sling"
)

func TestCategoryService_GetCategories(t *testing.T) {
  config := Init(time.Second, key)
  client := config.Client()
  cs := newCategoryService(sling.New().Client(client.client).Base("https://api.meetup.com/"), key)
  params := &ReqCategoryParams{
    Page: "20",
  }

  categories, resp, err := cs.GetCategories(params)

  if categories.Results == nil {
    t.Errorf("categories.Result is empty")
  }

  if len(categories.Results) == 0 {
    t.Errorf("categories. Result is of length zero")
  }

  if categories.Meta.TotalCount != 20 {
    t.Errorf("categories.Meta.TotalCount is not correct, it is %d", categories.Meta.TotalCount)
  }

  if resp == nil {
    t.Errorf("rest api response is nil")
  }

  if err != nil {
    t.Errorf("there was an error with the request, it error is %s", err.Error())
  }

}
