package meetup

import (
  "net/http"
  "github.com/dghubble/sling"
)

type Category struct {
  Name string `json:"name"`
  SortName string `json:"sort_name"`
  Id int `json:"id"`
  ShortName string `json:"shortname"`
}

type CategoryResponse struct {
  Results []Category `json:"results"`
  Meta ResponseMeta `json:"meta"`
}

type CategoryError struct {
  Details string `json:"details"`
  Code string `json:"code"`
  Problem string `json:"problem"`
}

type ReqCategoryParams struct {
  Sign string `url:"sign,omitempty"`
  Key string  `url:"key,omitempty"`
  Fields string `url:"fields,omitempty"`
  Order string `url:"order,omitempty"`
  Page string `url:"page,omitempty"`
  Offset string `url:"offset,omitempty"`
  Desc string `url:"desc,omitempty"`
  Only string `url:"only,omitempty"`
  Omit string `url:"omit,omitempty"`
}

type CategoryService struct {
  sling *sling.Sling
  key string
}

func newCategoryService(sling *sling.Sling, apiKey string) *CategoryService {
  return &CategoryService{
    sling: sling.Path("2/categories?key=" + apiKey + "&sign=true"),
    key: apiKey,
  }
}

func (cs *CategoryService) GetCategories(params *ReqCategoryParams) (CategoryResponse, *http.Response, error) {
  categories := new(CategoryResponse)
  apiError := new(ApiError)
  resp, err := cs.sling.New().QueryStruct(params).Receive(categories, apiError)
  return *categories, resp, relevantError(err, *apiError)
}
