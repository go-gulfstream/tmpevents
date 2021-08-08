package tmpevents

import (
	"encoding/json"

	gulfstreamevent "github.com/go-gulfstream/gulfstream/pkg/event"
)

func init() {
	gulfstreamevent.RegisterCodec(SessionRegistered, &SessionRegisteredPayload{})
	gulfstreamevent.RegisterCodec(SessionRegistered, &SessionEnabledPayload{})
}

func (c *SessionRegisteredPayload) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *SessionRegisteredPayload) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *SessionEnabledPayload) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *SessionEnabledPayload) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}
