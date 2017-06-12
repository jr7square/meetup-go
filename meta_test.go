package meetup

import (
	"testing"
	"time"

	"github.com/dghubble/sling"
)

//const key = "1c79452f27445bf6c365104a2b6132"
func TestnewMetaService(t *testing.T) {
	meta := newMetaService(sling.New().Base("https://api.meetup.com/"), key)

	if meta == nil {
		t.Errorf("meta is nil")
	}

	if meta.key != key {
		t.Errorf("meta.key is not assigned properly")
	}
}

func TestMetaService_Status(t *testing.T) {
	config := Init(time.Second, key)
	client := config.Client()
	meta := newMetaService(sling.New().Client(client.client).Base("https://api.meetup.com/"), key)

	status, resp, err := meta.Status()
	//
	if status.Message != "" {
		t.Errorf("message is not empty")
	}

	if status.Status != "ok" {
		t.Errorf("status was not ok, it was")
	}

	if status.Title != "" {
		t.Errorf("title is not an empty string")
	}

	if resp == nil {
		t.Errorf("http response is nil")
	}

	if err != nil {
		t.Errorf("Error when making request")
	}

	if status != nil && err != nil {
		t.Errorf("status and err are both assigned")
	}
}
