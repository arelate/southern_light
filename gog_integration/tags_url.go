package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func CreateTagUrl(name string) *url.URL {
	createTag := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   GogHost,
		Path:   accountTagsCreatePath,
	}

	q := createTag.Query()
	q.Add("name", name)
	createTag.RawQuery = q.Encode()

	return createTag
}

func DeleteTagUrl(tagId string) *url.URL {
	deleteTag := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   GogHost,
		Path:   accountTagsDeletePath,
	}

	q := deleteTag.Query()
	q.Add("tag_id", tagId)
	deleteTag.RawQuery = q.Encode()

	return deleteTag
}

func UpdateTagsUrl(tagsJson string) *url.URL {
	updateTags := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   GogHost,
		Path:   accountTagsUpdatePath,
	}

	q := updateTags.Query()
	q.Add("tags", tagsJson)
	updateTags.RawQuery = q.Encode()

	return updateTags
}

func AddTagUrl(productId, tagId string) *url.URL {
	addTag := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   GogHost,
		Path:   accountTagsAttachPath,
	}

	q := addTag.Query()
	q.Add("product_id", productId)
	q.Add("tag_id", tagId)
	addTag.RawQuery = q.Encode()

	return addTag
}

func RemoveTagUrl(productId, tagId string) *url.URL {
	removeTag := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   GogHost,
		Path:   accountTagsDetachPath,
	}

	q := removeTag.Query()
	q.Add("product_id", productId)
	q.Add("tag_id", tagId)
	removeTag.RawQuery = q.Encode()

	return removeTag
}
