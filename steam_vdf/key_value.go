package steam_vdf

import (
	"fmt"
	"io"
	"strings"
)

type KeyValues struct {
	Key    string
	Value  *string
	Values []*KeyValues
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

func (kv *KeyValues) WriteString(w io.Writer, depth int) error {

	// (depth times \t)"key"
	if _, err := io.WriteString(w, strings.Repeat("\t", depth)); err != nil {
		return err
	}
	if _, err := io.WriteString(w, fmt.Sprintf("\"%s\"", kv.Key)); err != nil {
		return err
	}

	// \t\t"value"
	if kv.Value != nil {
		if _, err := io.WriteString(w, strings.Repeat("\t", 2)); err != nil {
			return err
		}
		if _, err := io.WriteString(w, fmt.Sprintf("\"%s\"", *kv.Value)); err != nil {
			return err
		}
	}

	// \n
	if _, err := io.WriteString(w, "\n"); err != nil {
		return err
	}

	// (depth times \t){\n
	if kv.Values != nil || (kv.Values == nil && kv.Value == nil) {
		if _, err := io.WriteString(w, strings.Repeat("\t", depth)); err != nil {
			return err
		}
		if _, err := io.WriteString(w, "{"); err != nil {
			return err
		}
		if _, err := io.WriteString(w, "\n"); err != nil {
			return err
		}
	}

	// write Values at depth + 1
	for _, val := range kv.Values {
		if err := val.WriteString(w, depth+1); err != nil {
			return err
		}
	}

	// (depth times \t)}\n
	if kv.Values != nil || (kv.Values == nil && kv.Value == nil) {
		if _, err := io.WriteString(w, strings.Repeat("\t", depth)); err != nil {
			return err
		}
		if _, err := io.WriteString(w, "}"); err != nil {
			return err
		}
		if _, err := io.WriteString(w, "\n"); err != nil {
			return err
		}
	}

	return nil
}
