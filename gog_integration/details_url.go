// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
	"strings"
)

func DetailsUrl(id string) *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   GogHost,
		Path:   strings.Replace(detailsPathTemplate, "{id}", id, 1),
	}
}
