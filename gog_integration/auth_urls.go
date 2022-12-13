// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import (
	"net/url"
)

const (
	gogClientId = "46755278331571209"
)

func authUrl(host string) *url.URL {
	authURL := &url.URL{
		Scheme: HttpsScheme,
		Host:   host,
		Path:   authPath,
	}

	setAuthQuery(authURL)

	return authURL
}

func setAuthQuery(url *url.URL) {
	q := url.Query()
	q.Set("client_id", gogClientId)
	q.Set("redirect_uri", OnLoginSuccessUrl().String())
	q.Set("response_type", "code")
	q.Set("layout", "default")
	q.Set("brand", "gog")
	q.Set("gog_lc", "en-US")
	url.RawQuery = q.Encode()
}

func AuthHostUrl() *url.URL {
	return authUrl(authHost)
}

func LoginHostUrl() *url.URL {
	return authUrl(loginHost)
}
