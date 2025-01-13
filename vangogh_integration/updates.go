package vangogh_integration

import (
	"fmt"
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
			createdAfter, err := vr.CreatedAfter(since)
			if err != nil {
				return nil, err
			}
			categorize(createdAfter,
				fmt.Sprintf("new in %s", pt.HumanReadableString()),
				updates)
		}

		if interestingUpdatedProductTypes[pt] {
			updatedAfter, err := vr.UpdatedAfter(since)
			if err != nil {
				return nil, err
			}
			categorize(updatedAfter,
				fmt.Sprintf("updates in %s", pt.HumanReadableString()),
				updates)
		}
	}

	return updates, nil
}

func categorize(ids []string, cat string, updates map[string]map[string]bool) {
	for _, id := range ids {
		if updates[cat] == nil {
			updates[cat] = make(map[string]bool, 0)
		}
		updates[cat][id] = true
	}
}
