package meetup

import (
	"net/http"
	"github.com/dghubble/sling"
)

const meetupAPI = "https://api.meetup.com/"

/*Client is a go client for the MeetUp Api */
type ClientParams struct {
	client *http.Client
	key    string
}

type MeetupClient struct {
	sling *sling.Sling
	Meta *MetaService
}

func NewClient(client *ClientParams) *MeetupClient {
	base := sling.New().Client(client.client).Base(meetupAPI)
	return &MeetupClient {
					sling: base,
					Meta:  newMetaService(base.New(), client.key),
				}
}
