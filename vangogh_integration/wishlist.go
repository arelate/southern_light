package vangogh_integration

import (
	"github.com/boggydigital/nod"
	"github.com/boggydigital/redux"
)

func AddToLocalWishlist(
	ids []string,
	tpw nod.TotalProgressWriter) ([]string, error) {

	processedIds := make([]string, 0, len(ids))

	rdx, err := redux.NewWriter(AbsReduxDir(), UserWishlistProperty)
	if err != nil {
		return processedIds, err
	}

	if tpw != nil {
		tpw.TotalInt(len(ids))
	}

	for _, id := range ids {

		if err = rdx.ReplaceValues(UserWishlistProperty, id, TrueValue); err != nil {
			if tpw != nil {
				tpw.Increment()
			}
			return processedIds, err
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

	rdx, err := redux.NewWriter(AbsReduxDir(), UserWishlistProperty)
	if err != nil {
		return processedIds, err
	}

	if tpw != nil {
		tpw.TotalInt(len(ids))
	}

	for _, id := range ids {
		if err = rdx.ReplaceValues(UserWishlistProperty, id, FalseValue); err != nil {
			if tpw != nil {
				tpw.Increment()
			}
			return processedIds, err
		}

		processedIds = append(processedIds, id)
		if tpw != nil {
			tpw.Increment()
		}
	}

	return processedIds, err
}
