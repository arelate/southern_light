// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import (
	"net/http"
	"net/url"
)

const (
	acceptTextHTML = "text/html"
	acceptImages   = "image/webp,image/png,image/svg+xml,image/*"
)

func setDefaultHeaders(req *http.Request, host, accept string) {
	const (
		acceptLanguageHeader = "en-us"
		connectionHeader     = "keep-alive"
		userAgentHeader      = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) " +
			"AppleWebKit/537.36 (KHTML, like Gecko) " +
			"Chrome/97.0.4692.99 " +
			"Safari/537.36"
	)
	if host != "" {
		req.Host = host
	}
	if accept != "" {
		req.Header.Set("Accept", accept)
	}
	req.Header.Set("Accept-Language", acceptLanguageHeader)
	req.Header.Set("Connection", connectionHeader)
	req.Header.Set("User-Agent", userAgentHeader)
	req.Header.Set("Origin", WwwGogHost)
}

func addAuthHostDefaultHeaders(req *http.Request) {
	setDefaultHeaders(req, authHost, acceptTextHTML)
}

func addLoginHostDefaultHeaders(req *http.Request) {
	setDefaultHeaders(req, loginHost, acceptTextHTML)
}

func setLoginFormHeaders(req *http.Request, referer *url.URL) {
	const (
		urlEncoded = "application/x-www-form-urlencoded"
	)

	req.Header.Set("Origin", loginHost)
	req.Header.Set("Referer", referer.String())
	req.Header.Set("Content-Type", urlEncoded)
}
