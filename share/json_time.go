package share

import (
	"encoding/json"
	"time"
)

type JSONTime struct {
	Value time.Time
	Valid bool
	Set   bool
}

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	t.Set = true

	if string(data) == "null" {
		// The key was set to null
		t.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp time.Time
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	t.Value = temp
	t.Valid = true
	return nil
}
