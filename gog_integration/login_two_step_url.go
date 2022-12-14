// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import (
	"github.com/arelate/southern_light"
	"net/url"
)

func LoginTwoStepUrl() *url.URL {
	return &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   loginHost,
		Path:   loginTwoStepPath,
	}
}
