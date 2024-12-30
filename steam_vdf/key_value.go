package steam_vdf

type KeyValue struct {
	Key      string
	Val      *string
	Children []*KeyValue
}
