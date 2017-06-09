package meetup

import (
  "testing"
  "time"
)

const key = "1c79452f27445bf6c365104a2b6132"

func TestStatus(t *testing.T) {
  config := Init(time.Second, key)
  client := config.NewMeetupClient()

  status, err := client.Status()
  if err != nil {
    t.Errorf("Something went wrong")
  }
  if status.Status != "ok" {
    t.Errorf("api is not ok")
  }
  if status.Message != "" {
    t.Errorf("Message is not empty string")
  }
  if status.Title != "" {
    t.Errorf("Title is not empty ")
  }
}
