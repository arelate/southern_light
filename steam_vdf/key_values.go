package steam_vdf

import (
	"errors"
)

var ErrVdfKeyNotFound = errors.New("vdf key not found")

type KeyValues struct {
	Key        string
	Value      *string
	Type       BinaryType
	TypedValue any
	Values     ValveDataFile
}

type ValveDataFile []*KeyValues

func (vdf ValveDataFile) At(pathParts ...string) (*KeyValues, error) {

	current := new(KeyValues{Values: vdf})

	for _, part := range pathParts {

		found := false
		for _, kv := range current.Values {
			if kv.Key == part {
				current = kv
				found = true
			}
		}

		if !found {
			return nil, ErrVdfKeyNotFound
		}
	}

	return current, nil
}

func (vdf ValveDataFile) Val(pathParts ...string) (string, bool) {
	at, err := vdf.At(pathParts...)
	if errors.Is(err, ErrVdfKeyNotFound) {
		return "", false
	} else if err != nil {
		return "", false
	}

	switch at.Value {
	case nil:
		return "", false
	default:
		return *at.Value, true
	}
}
