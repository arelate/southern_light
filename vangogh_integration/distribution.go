package vangogh_integration

import (
	"errors"
	"maps"
	"slices"
	"strings"
)

type Store int

const (
	StoreUnknown Store = iota
	StoreVangogh
	StoreSteam
	StoreEpicGames
	StoreGog
)

var storeStrings = map[Store]string{
	StoreUnknown:   "unknown",
	StoreVangogh:   "vangogh",
	StoreSteam:     "Steam",
	StoreEpicGames: "EGS",
	StoreGog:       "GOG",
}

func (store Store) String() string {
	if stStr, ok := storeStrings[store]; ok {
		return stStr
	}
	return storeStrings[StoreUnknown]
}

func ParseStore(stStr string) Store {
	stStr = strings.ToLower(stStr)
	for store, str := range storeStrings {
		if strings.ToLower(str) == stStr {
			return store
		}
	}
	return StoreUnknown
}

func (store Store) ErrUnsupportedStore() error {
	return errors.New("unsupported store: " + store.String())
}

func AllStores() []string {
	return slices.Collect(maps.Values(storeStrings))
}
