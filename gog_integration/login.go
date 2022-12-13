// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import (
	"errors"
	"fmt"
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

const (
	reCaptchaError       = "reCAPTCHA present on the page"
	secondStepCodePrompt = "Two-step authentication security code:"
)

func authToken(client *http.Client) (token string, error error) {

	req, err := http.NewRequest(http.MethodGet, AuthHostUrl().String(), nil)
	addAuthHostDefaultHeaders(req)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	// check for captcha presence
	if match_node.Match(doc, &scriptReCaptchaMatcher{}) != nil {
		// TODO: Write how to add cookie from the browser to allow user to unblock themselves
		return "", errors.New(reCaptchaError)
	}

	input := match_node.Match(doc, &inputLoginTokenMatcher{})
	token = match_node.AttrVal(input, "value")

	if err := resp.Body.Close(); err != nil {
		return token, err
	}

	return token, nil
}

func secondStepAuth(client *http.Client, body io.ReadCloser, requestText func(string) string) error {

	doc, err := html.Parse(body)
	if err != nil {
		return err
	}

	input := match_node.Match(doc, &inputSecondStepAuthTokenMatcher{})
	token := match_node.AttrVal(input, "value")

	for token != "" {

		code := ""
		if requestText != nil {
			for len(code) != 4 {
				code = requestText(secondStepCodePrompt)
			}
		} else {
			return fmt.Errorf("2FA token is requied, cannot get it with nil callback")
		}

		data := secondStepData(code, token)

		req, _ := http.NewRequest(http.MethodPost, LoginTwoStepUrl().String(), strings.NewReader(data))
		addLoginHostDefaultHeaders(req)
		setLoginFormHeaders(req, LoginTwoStepUrl())

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		doc, err = html.Parse(resp.Body)
		if err != nil {
			return err
		}

		input = match_node.Match(doc, &inputSecondStepAuthTokenMatcher{})
		token = match_node.AttrVal(input, "value")

		if err := resp.Body.Close(); err != nil {
			return err
		}
	}

	return nil
}

/*
Login to GOG.com for account formData queries using username and password

Overall flow is:
- Get auth token from the page (this would check for reCaptcha as well)
- Post username, password and token for check
- Check for presence of second step auth token
- (Optional) Post 4 digit second step auth code
*/
func Login(client *http.Client, username, password string, requestText func(string) string) error {

	if username == "" {
		resp := requestText("GOG.com username:")
		if resp == "" {
			return fmt.Errorf("username cannot be empty")
		}
		username = resp
	}

	if password == "" {
		resp := requestText(fmt.Sprintf("Enter password for %s:", username))
		if resp == "" {
			return fmt.Errorf("password cannot be empty")
		}
		password = resp
	}

	token, err := authToken(client)
	if err != nil {
		return err
	}

	data := loginData(username, password, token)

	req, err := http.NewRequest(http.MethodPost, LoginCheckUrl().String(), strings.NewReader(data))
	if err != nil {
		return err
	}
	addLoginHostDefaultHeaders(req)
	// GOG.com redirects initial auth request from authHost to loginHost.
	setLoginFormHeaders(req, LoginHostUrl())

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if err := secondStepAuth(client, resp.Body, requestText); err != nil {
		return err
	}

	return resp.Body.Close()
}
