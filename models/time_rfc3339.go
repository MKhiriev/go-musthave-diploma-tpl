package models

import (
	"encoding/json"
	"time"
)

type RFC3339Time time.Time

func (t RFC3339Time) MarshalJSON() ([]byte, error) {
	stamp := time.Time(t).Format(time.RFC3339)
	return json.Marshal(stamp)
}

//goland:noinspection GoMixedReceiverTypes
func (t *RFC3339Time) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	parsed, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}

	*t = RFC3339Time(parsed)
	return nil
}

//goland:noinspection GoMixedReceiverTypes
func (t RFC3339Time) ToTime() time.Time {
	return time.Time(t)
}
