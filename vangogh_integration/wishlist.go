package vangogh_integration

import (
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
)

func AddToLocalWishlist(
	ids []string,
	tpw nod.TotalProgressWriter) ([]string, error) {

	processedIds := make([]string, 0, len(ids))

	rxa, err := NewReduxWriter(WishlistedProperty)
	if err != nil {
		return processedIds, err
	}

	if tpw != nil {
		tpw.TotalInt(len(ids))
	}

	for _, id := range ids {
		// remove "false" reduction
		if rxa.HasValue(WishlistedProperty, id, FalseValue) {
			if err := rxa.CutValues(WishlistedProperty, id, FalseValue); err != nil {
				if tpw != nil {
					tpw.Increment()
				}
				return processedIds, err
			}
		}

		if !rxa.HasValue(WishlistedProperty, id, TrueValue) {
			if err := rxa.AddValues(WishlistedProperty, id, TrueValue); err != nil {
				if tpw != nil {
					tpw.Increment()
				}
				return processedIds, err
			}
		}

		processedIds = append(processedIds, id)
		if tpw != nil {
			tpw.Increment()
		}
	}

	return processedIds, nil
}

func RemoveFromLocalWishlist(
	ids []string,
	tpw nod.TotalProgressWriter) ([]string, error) {

	processedIds := make([]string, 0, len(ids))

	rxa, err := NewReduxWriter(WishlistedProperty)
	if err != nil {
		return processedIds, err
	}

	if tpw != nil {
		tpw.TotalInt(len(ids))
	}

	for _, id := range ids {
		if rxa.HasValue(WishlistedProperty, id, TrueValue) {
			if err := rxa.CutValues(WishlistedProperty, id, TrueValue); err != nil {
				if tpw != nil {
					tpw.Increment()
				}
				return processedIds, err
			}
		}

		if !rxa.HasValue(WishlistedProperty, id, FalseValue) {
			if err := rxa.AddValues(WishlistedProperty, id, FalseValue); err != nil {
				if tpw != nil {
					tpw.Increment()
				}
				return processedIds, err
			}
		}

		processedIds = append(processedIds, id)
		if tpw != nil {
			tpw.Increment()
		}
	}

	ptDir, err := AbsLocalProductTypeDir(UserWishlistProducts)
	if err != nil {
		return processedIds, err
	}
	kvPt, err := kevlar.NewKeyValues(ptDir, kevlar.JsonExt)
	if err != nil {
		return processedIds, err
	}

	for _, id := range processedIds {
		if _, err := kvPt.Cut(id); err != nil {
			return processedIds, err
		}
	}

	// don't check err because we're immediately returning it
	return processedIds, err
}
