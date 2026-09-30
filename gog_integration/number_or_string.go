package gog_integration

import (
	"encoding/json/v2"
	"errors"
)

// numberOrString is a string value that GOG may encode either as a JSON string or as a
// JSON number (e.g. account page tags "productCount": 0). Numbers are kept verbatim.
type numberOrString string

func (ns *numberOrString) UnmarshalJSON(data []byte) error {
	switch {
	case len(data) == 0:
		return errors.New("numberOrString: empty value")
	case string(data) == "null":
		*ns = ""
	case data[0] == '"':
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		*ns = numberOrString(s)
	case data[0] == '-' || (data[0] >= '0' && data[0] <= '9'):
		*ns = numberOrString(data)
	default:
		return errors.New("numberOrString: expected a JSON string or number, got " + string(data))
	}
	return nil
}
