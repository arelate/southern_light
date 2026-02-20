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
	Values     []*KeyValues
}

type ValveDataFile []*KeyValues

func (kv *KeyValues) Val(key string) string {
	for _, value := range kv.Values {
		if value.Key == key {
			switch value.Value {
			case nil:
				return ""
			default:
				return *value.Value
			}
		}
	}
	return ""
}

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
