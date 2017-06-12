package meetup

import (
)

type Category struct {
  Name string `json:"name"`
  SortName string `json:"sort_name"`
  Id int `json:"id"`
  ShortName string `json:"shortname"`
}

type CategoryError struct {
  Details string `json:"details"`
  Code string `json:"code"`
  Problem string `json:"problem"`
}


type ReqParams struct {
  Fields string
  Order string
  Page string
  Offset string
  Desc string
  Only string
  Omit string
}
