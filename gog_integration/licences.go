package gog_integration

import "strconv"

type Licence string

func (lc Licence) GetId() int {
	lcStr, err := strconv.Atoi(string(lc))
	if err != nil {
		return 0
	}
	return lcStr
}

type Licences []Licence

func (lcs Licences) GetProducts() []IdGetter {
	idg := make([]IdGetter, 0, len(lcs))
	for _, lc := range lcs {
		idg = append(idg, lc)
	}
	return idg
}
