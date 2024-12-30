package steam_vdf

import (
	"fmt"
	"path"
	"strings"
)

type KeyValue struct {
	Key string
	Val *string
}

func (kv *KeyValue) String() string {
	if kv.Val != nil {
		return fmt.Sprintf("%s=%s", kv.Key, *kv.Val)
	} else {
		return fmt.Sprintf("%s={}", kv.Key)
	}
}

func GetValByAbsKey(keyValues []*KeyValue, absKey string) *string {
	for _, kv := range keyValues {
		if kv.Key == absKey {
			return kv.Val
		}
	}
	return nil
}

func GetValByRelKey(keyValues []*KeyValue, relKey string) *string {
	for _, kv := range keyValues {
		if strings.HasSuffix(kv.Key, relKey) {
			return kv.Val
		}
	}
	return nil
}

func GetDirectChildren(keyValues []*KeyValue, absKey string) []string {
	children := make([]string, 0)
	for _, kv := range keyValues {
		pfx, sfx := path.Split(kv.Key)
		pfx = strings.TrimSuffix(pfx, "/")
		if strings.HasPrefix(pfx, absKey) && len(pfx) > len(absKey) {
			break
		}
		if pfx == absKey {
			children = append(children, sfx)
		}
	}
	return children
}
