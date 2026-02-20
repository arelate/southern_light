package steam_vdf

type KeyValues struct {
	Key        string
	Value      *string
	Type       BinaryType
	TypedValue any
	Values     []*KeyValues
}

type ValveDataFile []*KeyValues

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

func (vdf ValveDataFile) GetKv(pathParts ...string) (*KeyValues, error) {
	return nil, nil
}

func (vdf ValveDataFile) GetVdf(pathParts ...string) (ValveDataFile, error) {
	return nil, nil
}
