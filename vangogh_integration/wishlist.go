package vangogh_integration

import (
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
)

func AddToLocalWishlist(
	ids []string,
	tpw nod.TotalProgressWriter) ([]string, error) {

	processedIds := make([]string, 0, len(ids))

	rxa, err := NewReduxWriter(UserWishlistProperty)
	if err != nil {
		return processedIds, err
	}

	if tpw != nil {
		tpw.TotalInt(len(ids))
	}

	for _, id := range ids {
		// remove "false" reduction
		//if rxa.HasValue(UserWishlistProperty, UserWishlistProperty, id) {
		//	if err := rxa.CutValues(UserWishlistProperty, UserWishlistProperty, id); err != nil {
		//		if tpw != nil {
		//			tpw.Increment()
		//		}
		//		return processedIds, err
		//	}
		//}

		if !rxa.HasValue(UserWishlistProperty, UserWishlistProperty, id) {
			if err := rxa.AddValues(UserWishlistProperty, UserWishlistProperty, id); err != nil {
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

	rxa, err := NewReduxWriter(UserWishlistProperty)
	if err != nil {
		return processedIds, err
	}

	if tpw != nil {
		tpw.TotalInt(len(ids))
	}

	for _, id := range ids {
		if rxa.HasValue(UserWishlistProperty, UserWishlistProperty, id) {
			if err := rxa.CutValues(UserWishlistProperty, UserWishlistProperty, id); err != nil {
				if tpw != nil {
					tpw.Increment()
				}
				return processedIds, err
			}
		}

		//if !rxa.HasValue(WishlistedProperty, id, FalseValue) {
		//	if err := rxa.AddValues(WishlistedProperty, id, FalseValue); err != nil {
		//		if tpw != nil {
		//			tpw.Increment()
		//		}
		//		return processedIds, err
		//	}
		//}

		processedIds = append(processedIds, id)
		if tpw != nil {
			tpw.Increment()
		}
	}

	ptDir, err := AbsProductTypeDir(UserWishlistProducts)
	if err != nil {
		return processedIds, err
	}
	kvPt, err := kevlar.New(ptDir, kevlar.JsonExt)
	if err != nil {
		return processedIds, err
	}

	for _, id := range processedIds {
		if err = kvPt.Cut(id); err != nil {
			return processedIds, err
		}
	}

	// don't check err because we're immediately returning it
	return processedIds, err
}
