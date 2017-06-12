package meetup

import (
	"net/http"
	"time"
)

/*ClientConfig contains all the configurations available for the Meetup Client */
type ClientConfig struct {
	ClientTimeout time.Duration
	Key           string
}

type Credentials struct {
	key string
}

/*Init initializes all the configurations specified */
func Init(timeout time.Duration, apiKey string) *ClientConfig {
	return &ClientConfig{
		ClientTimeout: timeout,
		Key:           apiKey,
	}
}

/*Client constructs MeetupClinet*/
func (config *ClientConfig) Client() *ClientParams {
	return &ClientParams{
		client: &http.Client{Timeout: config.ClientTimeout},
		key:    config.Key,
	}
}
