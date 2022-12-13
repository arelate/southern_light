// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import "net/url"

func loginData(username, password, token string) string {
	data := url.Values{
		"login[username]":   {username},
		"login[password]":   {password},
		"login[login_flow]": {"default"},
		"login[_token]":     {token},
	}
	return data.Encode()
}
