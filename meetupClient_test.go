package meetup

import (
  "testing"
  "time"
)

const key = "1c79452f27445bf6c365104a2b6132"

func TestNewClient(t *testing.T) {
  config := Init(time.Second, key)
  client := config.Client()

  meetup := NewClient(client)

  if meetup == nil {
    t.Errorf("meetup is nil")
  }
  if meetup.sling == nil {
    t.Errorf("meetup sling is nil")
  }
  if meetup.Meta == nil {
    t.Errorf("meetup Meta is nil")
  }

}

// func TestStatus(t *testing.T) {
//   config := Init(time.Second, key)
//   client := config.Client()
//   meetup := NewClient(client)
//
//
//   status, err := meetup.Meta.Status()
//   if err != nil {
//     t.Errorf("Something went wrong")
//   }
//   if status.Status != "ok" {
//     t.Errorf("api is not ok")
//   }
//   if status.Message != "" {
//     t.Errorf("Message is not empty string")
//   }
//   if status.Title != "" {
//     t.Errorf("Title is not empty ")
//   }
// }
