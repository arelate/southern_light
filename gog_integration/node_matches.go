// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import (
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"strings"
)

type inputLoginTokenMatcher struct{}

func (iltm *inputLoginTokenMatcher) Match(n *html.Node) bool {
	return n != nil &&
		n.Type == html.ElementNode &&
		n.Data == "input" &&
		match_node.AttrVal(n, "name") == "login[_token]"
}

type inputSecondStepAuthTokenMatcher struct{}

func (issatm *inputSecondStepAuthTokenMatcher) Match(n *html.Node) bool {
	return n != nil &&
		n.Type == html.ElementNode &&
		n.Data == "input" &&
		match_node.AttrVal(n, "name") == "second_step_authentication[_token]"
}

type scriptReCaptchaMatcher struct{}

func (srcm *scriptReCaptchaMatcher) Match(n *html.Node) bool {
	return n != nil &&
		n.Type == html.ElementNode &&
		n.Data == "script" &&
		strings.HasPrefix(match_node.AttrVal(n, "src"), ReCaptchaUrl().String())
}
