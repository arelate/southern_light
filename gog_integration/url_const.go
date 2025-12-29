// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import (
	"net/url"

	"github.com/arelate/southern_light"
)

// hosts
const (
	GogHost     = "gog.com"
	WwwGogHost  = "www." + GogHost
	apiHost     = "api." + GogHost
	imagesHost  = "images.gog-statics.com"
	menuHost    = "menu." + GogHost
	loginHost   = "login." + GogHost
	authHost    = "auth." + GogHost
	itemsHost   = "items." + GogHost
	catalogHost = "catalog." + GogHost
	gamesDbHost = "gamesdb." + GogHost
)

const (
	userDataPath = "/userData.json"
)

const (
	accountLicencesPath = "/v1/account/licences"
)

const (
	userWishlistJsonPath = "/user/wishlist.json"
)

const (
	userWishlistAddPathTemplate    = "/user/wishlist/add/{id}"
	userWishlistRemovePathTemplate = "/user/wishlist/remove/{id}"
)

const (
	catalogPath = "/v1/catalog"
)

const (
	accountGameDetailsPathTemplate = "/account/gameDetails/{id}.json"
	accountPagePath                = "/account/getFilteredProducts"
	accountOrdersPath              = "/account/settings/orders/data"
)

const (
	apiV2GamesPath = "/v2/games/"
)

const (
	userAccessTokenPath = "/user/accessToken.json"
)

const (
	gamesDbGogExternalReleasePathTemplate     = "/platforms/gog/external_releases/{id}"
	gamesDbGenericExternalReleasePathTemplate = "/platforms/generic/external_releases/{id}"
	gamesDbWishlistedGamesPath                = "/wishlist/user/wishlisted_games"
)

const (
	accountTagsCreatePath = "/account/tags/add"
	accountTagsDeletePath = "/account/tags/delete"
	accountTagsUpdatePath = "/account/tags/update"
	accountTagsAttachPath = "/account/tags/attach"
	accountTagsDetachPath = "/account/tags/detach"
)

const (
	imagesPathTemplate = "/{image_id}"
)

func DefaultUrl() *url.URL {
	return &url.URL{Scheme: southern_light.HttpsScheme, Host: GogHost}
}
