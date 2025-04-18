package vangogh_integration

import (
	"encoding/json"
	"fmt"
	"github.com/arelate/southern_light/gog_integration"
	"github.com/boggydigital/nod"
	"github.com/boggydigital/pathways"
	"github.com/boggydigital/redux"
	"net/http"
	"net/url"
	"slices"
)

func postTagResp(httpClient *http.Client, url *url.URL, respVal interface{}) error {
	resp, err := httpClient.Post(url.String(), "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(&respVal)
}

func TagIdByName(tagName string) (string, error) {
	reduxDir, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return "", err
	}

	rdx, err := redux.NewReader(reduxDir, TagNameProperty)
	if err != nil {
		return "", err
	}

	tagIds := rdx.Match(map[string][]string{TagNameProperty: {tagName}})
	if tagIds == nil {
		return "", fmt.Errorf("unknown tag-name %s", tagName)
	}
	tagId := ""
	for ti := range tagIds {
		if tagId != "" {
			return "", fmt.Errorf("ambiguous tag-name %s, matching tag-ids: %v",
				tagName,
				tagIds)
		}
		tagId = ti
	}
	return tagId, nil
}

func CreateTag(httpClient *http.Client, tagName string) error {

	reduxDir, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return err
	}

	rdx, err := redux.NewWriter(reduxDir, TagNameProperty)
	if err != nil {
		return err
	}

	createTagUrl := gog_integration.CreateTagUrl(tagName)
	var ctResp gog_integration.CreateTagResp
	if err := postTagResp(httpClient, createTagUrl, &ctResp); err != nil {
		return err
	}
	if ctResp.Id == "" {
		return fmt.Errorf("invalid create tag response")
	}

	if !rdx.HasValue(TagNameProperty, ctResp.Id, tagName) {
		if err := rdx.AddValues(TagNameProperty, ctResp.Id, tagName); err != nil {
			return err
		}
	}

	return nil
}

func DeleteTag(httpClient *http.Client, tagName, tagId string) error {

	reduxDir, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return err
	}

	rdx, err := redux.NewWriter(reduxDir, TagNameProperty)
	if err != nil {
		return err
	}

	deleteTagUrl := gog_integration.DeleteTagUrl(tagId)
	var dtResp gog_integration.DeleteTagResp
	if err := postTagResp(httpClient, deleteTagUrl, &dtResp); err != nil {
		return err
	}
	if dtResp.Status != "deleted" {
		return fmt.Errorf("invalid delete tag response")
	}

	if rdx.HasValue(TagNameProperty, tagId, tagName) {
		if err := rdx.CutValues(TagNameProperty, tagId, tagName); err != nil {
			return err
		}
	}

	return nil
}

func AddTags(
	httpClient *http.Client,
	ids, tags []string,
	tpw nod.TotalProgressWriter) error {

	reduxDir, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return err
	}

	rdx, err := redux.NewWriter(reduxDir, TagIdProperty)
	if err != nil {
		return err
	}

	nod.TotalInt(tpw, len(ids)*len(tags))

	for _, id := range ids {
		for _, tag := range tags {

			if rdx.HasValue(TagIdProperty, id, tag) {
				nod.Increment(tpw)
				continue
			}

			addTagUrl := gog_integration.AddTagUrl(id, tag)
			var artResp gog_integration.AddRemoveTagResp
			if err := postTagResp(httpClient, addTagUrl, &artResp); err != nil {
				if tpw != nil {
					tpw.Increment()
				}
				return err
			}
			if !artResp.Success {
				nod.Increment(tpw)
				return fmt.Errorf("failed to add tag %s", tag)
			}

			if err := rdx.AddValues(TagIdProperty, id, tag); err != nil {
				nod.Increment(tpw)
				return err
			}

			nod.Increment(tpw)
		}
	}

	return nil
}

func RemoveTags(
	httpClient *http.Client,
	ids, tags []string,
	tpw nod.TotalProgressWriter) error {

	reduxDir, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return err
	}

	rdx, err := redux.NewWriter(reduxDir, TagIdProperty)
	if err != nil {
		return err
	}

	nod.TotalInt(tpw, len(ids)*len(tags))

	for _, id := range ids {
		for _, tag := range tags {

			if !rdx.HasValue(TagIdProperty, id, tag) {
				nod.Increment(tpw)
				continue
			}

			removeTagUrl := gog_integration.RemoveTagUrl(id, tag)
			var artResp gog_integration.AddRemoveTagResp
			if err := postTagResp(httpClient, removeTagUrl, &artResp); err != nil {
				nod.Increment(tpw)
				return err
			}
			if !artResp.Success {
				nod.Increment(tpw)
				return fmt.Errorf("failed to remove tag %s", tag)
			}

			if err := rdx.CutValues(TagIdProperty, id, tag); err != nil {
				nod.Increment(tpw)
				return err
			}

			nod.Increment(tpw)
		}
	}

	return nil
}

func diffTagProperty(
	tagProperty string,
	id string,
	newTags []string) (add []string, rem []string, err error) {

	add = make([]string, 0)
	rem = make([]string, 0)

	reduxDir, err := pathways.GetAbsRelDir(Redux)
	if err != nil {
		return nil, nil, err
	}

	rdx, err := redux.NewReader(reduxDir, tagProperty)
	if err != nil {
		return add, rem, err
	}

	//we need empty slice to detect new values
	currentVals, _ := rdx.GetAllValues(tagProperty, id)

	for _, tag := range newTags {
		if !slices.Contains(currentVals, tag) {
			add = append(add, tag)
		}
	}

	for _, tag := range currentVals {
		if !slices.Contains(newTags, tag) {
			rem = append(rem, tag)
		}
	}

	return add, rem, nil
}

func DiffTags(id string, newTags []string) (add []string, rem []string, err error) {
	return diffTagProperty(TagIdProperty, id, newTags)
}
