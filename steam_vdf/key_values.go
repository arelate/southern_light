package steam_vdf

import (
	"errors"
)

type KeyValues struct {
	Key        string
	Value      *string
	Type       BinaryType
	TypedValue any
	Values     ValveDataFile
}

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

type ValveDataFile []*KeyValues

var ErrVdfKeyNotFound = errors.New("vdf key not found")

func GetKevValuesByKey(keyValues []*KeyValues, key string) *KeyValues {

	queue := make(map[*KeyValues]bool)

	for _, kv := range keyValues {
		queue[kv] = true
	}

	for {
		var next *KeyValues
		for kv, flag := range queue {
			if flag == true {
				next = kv
				break
			}
		}

		if next == nil {
			break
		}

		if next.Key == key {
			return next
		}

		for _, kv := range next.Values {
			queue[kv] = true
		}

		queue[next] = false
	}

	return nil
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
