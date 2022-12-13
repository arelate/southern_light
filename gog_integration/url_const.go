// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

// hosts
const (
	GogHost       = "gog.com"
	WwwGogHost    = "www." + GogHost
	apiHost       = "api." + GogHost
	imagesHost    = "images.gog-statics.com"
	menuHost      = "menu." + GogHost
	cdnHost       = "cdn." + GogHost
	loginHost     = "login." + GogHost
	authHost      = "auth." + GogHost
	ItemsHost     = "items." + GogHost
	CatalogHost   = "catalog." + GogHost
	reCaptchaHost = "www.recaptcha.net"
)

// paths
const (
	storePagePath              = "/games/ajax/filtered"
	accountPath                = "/account"
	detailsPathTemplate        = accountPath + "/" + GameMedia + "Details/{id}.json"
	accountPagePath            = accountPath + "/getFilteredProducts"
	accountWishlistPath        = accountPath + "/wishlist"
	wishlistSearchPath         = accountWishlistPath + "/search"
	userPath                   = "/user"
	wishlistPath               = userPath + "/wishlist"
	wishlistAddPathTemplate    = wishlistPath + "/add/{id}"
	wishlistRemovePathTemplate = wishlistPath + "/remove/{id}"
	ordersPath                 = accountPath + "/settings/orders/data"
	apiV1PathTemplate          = "/products/{id}"
	apiV2GamesPath             = "/v2/games/"
	imagesPathTemplate         = "/{image_id}"
	licencesPath               = "/v1/account/licences"
	catalogPath                = "/v1/catalog"
	accountTagsPath            = "/account/tags"
	tagsCreatePath             = accountTagsPath + "/add"
	tagsDeletePath             = accountTagsPath + "/delete"
	tagsUpdatePath             = accountTagsPath + "/update"
	tagsAddPath                = accountTagsPath + "/attach"
	tagsRemovePath             = accountTagsPath + "/detach"
	securePath                 = "/secure"
	authPath                   = "/auth"
	loginCheckPath             = "/login_check"
	loginTwoStepPath           = "/login/two_step"
	reCaptchaPath              = "/recaptcha/api.js"
	userDataPath               = "/userData.json"
	userWishlistPath           = "/user/wishlist.json"
	onLoginSuccessPath         = "/on_login_success"
)
