package steam_vdf

type KeyValues struct {
	Key    string
	Value  *string
	Values []*KeyValues
}
