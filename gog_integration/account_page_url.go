// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

const (
	gameMediaType = "1"
)

func AccountPageUrl(page string) *url.URL {
	return accountPageUrl(
		page,
		AccountSortByPurchaseDate,
		false,
		false)
}

func accountPageUrl(
	page string,
	sortOrder AccountSortOrder,
	updated bool, /* get only updated products */
	hidden bool /* get only hidden products */) *url.URL {

	accountPage := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   GogHost,
		Path:   accountPagePath,
	}

	q := accountPage.Query()
	q.Add("mediaType", gameMediaType)
	if sortOrder != "" {
		q.Add("sortBy", string(sortOrder))
	}
	q.Add("page", page)
	if hidden {
		q.Add("hiddenFlag", "1")
	}
	if updated {
		q.Add("isUpdated", "1")
	}
	accountPage.RawQuery = q.Encode()

	return accountPage
}
