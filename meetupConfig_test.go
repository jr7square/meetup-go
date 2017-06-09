package meetup

import (
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	config := Init(time.Second, "12345")
	if config.ClientTimeout != time.Second {
		t.Errorf("ClientTimeout is incorrect")
	}
	if config.Key != "12345" {
		t.Errorf("key is incorrect")
	}
}

func TestNewMeetupClient(t *testing.T) {
	config := Init(time.Second, "12345")
	client := config.NewMeetupClient()
	if client.client == nil {
		t.Errorf("client is nil")
	}
	if client.key != config.Key {
		t.Errorf("key is not setup properly")
	}
}
