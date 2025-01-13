package vangogh_integration

import "sort"

type siKV struct {
	key string
	val int
}

type strIntKeyValues []siKV

func (kv strIntKeyValues) Len() int {
	return len(kv)
}

func (kv strIntKeyValues) Swap(i, j int) {
	kv[i], kv[j] = kv[j], kv[i]
}

func (kv strIntKeyValues) Less(i, j int) bool {
	return kv[i].val < kv[j].val
}

func (kv strIntKeyValues) GetKey(i int) string {
	return kv[i].key
}

func SortStrIntMap(m map[string]int, desc bool) []string {
	kvs := make(strIntKeyValues, 0, len(m))

	for k, v := range m {
		kvs = append(kvs, siKV{k, v})
	}

	var sortInterface sort.Interface = kvs
	if desc {
		sortInterface = sort.Reverse(sortInterface)
	}

	sort.Sort(sortInterface)

	sorted := make([]string, 0, len(kvs))

	for i := 0; i < len(kvs); i++ {
		sorted = append(sorted, kvs[i].key)
	}

	return sorted
}
