package steam_vdf

type KeyValues struct {
	Key        string
	Value      *string
	Type       BinaryType
	TypedValue any
	Values     []*KeyValues
}

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
