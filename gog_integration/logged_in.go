// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func LoggedIn(client *http.Client) (bool, error) {

	resp, err := client.Get(UserDataUrl().String())

	if err != nil {
		return false, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var ud UserData

	err = json.Unmarshal(respBody, &ud)
	if err != nil {
		return false, err
	}

	err = resp.Body.Close()
	if err != nil {
		return ud.IsLoggedIn, err
	}

	return ud.IsLoggedIn, nil
}
