// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ErrNotLoggedIn = errors.New("user is not logged in, please update cookies.txt")

func IsLoggedIn(client *http.Client) error {

	resp, err := client.Get(UserDataUrl().String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var ud UserData
	if err = json.NewDecoder(resp.Body).Decode(&ud); err != nil {
		return err
	}

	switch ud.IsLoggedIn {
	case true:
		return nil
	default:
		return ErrNotLoggedIn
	}
}
