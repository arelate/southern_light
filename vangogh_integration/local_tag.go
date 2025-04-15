package vangogh_integration

import (
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"github.com/boggydigital/redux"
)

func addLocalTag(id, tag string, rdx redux.Writeable, tpw nod.TotalProgressWriter) error {
	if err := rdx.MustHave(LocalTagsProperty); err != nil {
		return err
	}

	if !rdx.HasValue(LocalTagsProperty, id, tag) {
		if err := rdx.AddValues(LocalTagsProperty, id, tag); err != nil {
			nod.Increment(tpw)
			return err
		}
	}
	nod.Increment(tpw)
	return nil
}

func removeLocalTag(id, tag string, rdx redux.Writeable, tpw nod.TotalProgressWriter) error {
	if err := rdx.MustHave(LocalTagsProperty); err != nil {
		return err
	}

	if rdx.HasValue(LocalTagsProperty, id, tag) {
		if err := rdx.CutValues(LocalTagsProperty, id, tag); err != nil {
			nod.Increment(tpw)
			return err
		}
	}

	nod.Increment(tpw)
	return nil
}

func AddLocalTags(ids, tags []string, tpw nod.TotalProgressWriter) error {
	reduxDir, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return err
	}

	rdx, err := redux.NewWriter(reduxDir, LocalTagsProperty)
	if err != nil {
		return err
	}

	nod.TotalInt(tpw, len(ids)*len(tags))

	for _, id := range ids {
		for _, tag := range tags {
			if err = addLocalTag(id, tag, rdx, tpw); err != nil {
				return err
			}
		}
	}

	return nil
}

func RemoveLocalTags(ids, tags []string, tpw nod.TotalProgressWriter) error {
	reduxDir, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return err
	}

	rdx, err := redux.NewWriter(reduxDir, LocalTagsProperty)
	if err != nil {
		return err
	}

	nod.TotalInt(tpw, len(ids)*len(tags))

	for _, id := range ids {
		for _, tag := range tags {
			if err = removeLocalTag(id, tag, rdx, tpw); err != nil {
				return err
			}
		}
	}

	return nil
}

func DiffLocalTags(id string, newTags []string) (add []string, rem []string, err error) {
	return diffTagProperty(LocalTagsProperty, id, newTags)
}
