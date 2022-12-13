// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import (
	"net/url"
	"strings"
)

func AddToWishlistUrl(id string) *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   strings.Replace(wishlistAddPathTemplate, "{id}", id, 1),
	}
}

func RemoveFromWishlistUrl(id string) *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   strings.Replace(wishlistRemovePathTemplate, "{id}", id, 1),
	}
}
