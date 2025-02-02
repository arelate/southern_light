package vangogh_integration

import (
	"fmt"
	"github.com/boggydigital/kevlar"
	"iter"
)

var interestingNewProductTypes = map[ProductType]bool{
	CatalogProducts:      true,
	AccountProducts:      true,
	UserWishlistProducts: true,
	Details:              true,
}

var interestingUpdatedProductTypes = map[ProductType]bool{
	UserWishlistProducts: true,
	Details:              true,
	SteamAppNews:         true,
}

func Updates(since int64) (map[string]map[string]bool, error) {
	updates := make(map[string]map[string]bool)

	for _, pt := range LocalProducts() {

		vr, err := NewProductReader(pt)
		if err != nil {
			return nil, err
		}

		if interestingNewProductTypes[pt] {
			categorize(vr.Since(since, kevlar.Create),
				fmt.Sprintf("new in %s", pt.HumanReadableString()),
				updates)
		}

		if interestingUpdatedProductTypes[pt] {
			categorize(vr.Since(since, kevlar.Create, kevlar.Update),
				fmt.Sprintf("updates in %s", pt.HumanReadableString()),
				updates)
		}
	}

	return updates, nil
}

func categorize(ids iter.Seq2[string, kevlar.MutationType], cat string, updates map[string]map[string]bool) {
	for id := range ids {
		if updates[cat] == nil {
			updates[cat] = make(map[string]bool, 0)
		}
		updates[cat][id] = true
	}
}
