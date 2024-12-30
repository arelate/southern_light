package steam_vdf

import (
	"path"
	"strings"
)

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
