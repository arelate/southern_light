// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type ManualDownload struct {
	ManualUrl string `json:"manualUrl"`
	Name      string `json:"name"`
	Version   string `json:"version"`
	Date      string `json:"date"`
	Type      string `json:"type"`
	Info      int    `json:"info"`
	Size      string `json:"size"`
}

type LangOSManualDownloads struct {
	Language string
	Windows  []ManualDownload `json:"windows"`
	Mac      []ManualDownload `json:"mac"`
	Linux    []ManualDownload `json:"linux"`
}

type Details struct {
	Title           string `json:"title"`
	BackgroundImage string `json:"backgroundImage"`
	CdKey           string `json:"cdKey"`
	TextInformation string `json:"textInformation"`
	//GetDownloads comment explains why this is [][]interface{}
	Downloads              []interface{}    `json:"downloads"`
	GalaxyDownloads        []interface{}    `json:"galaxyDownloads"`
	Extras                 []ManualDownload `json:"extras"`
	DLCs                   []Details        `json:"dlcs"`
	Tags                   []accountTag     `json:"tags"`
	IsPreOrder             bool             `json:"isPreOrder"`
	ReleaseTimestamp       int64            `json:"releaseTimestamp"`
	Messages               []string         `json:"messages"`
	Changelog              string           `json:"changelog"`
	ForumLink              string           `json:"forumLink"`
	IsBaseProductMissing   bool             `json:"isBaseProductMissing"`
	MissingBaseProduct     interface{}      `json:"missingBaseProduct"`
	Features               []string         `json:"features"`
	SimpleGalaxyInstallers []struct {
		Path string `json:"path"`
		Os   string `json:"os"`
	} `json:"simpleGalaxyInstallers"`
}

func (det *Details) GetTitle() string {
	return det.Title
}

func (det *Details) GetFeatures() []string {
	return det.Features
}

func (det *Details) getAccountTags() []accountTag {
	return det.Tags
}

func (det *Details) GetTagIds() []string {
	return getTagIds(det)
}

func (det *Details) GetTagNames(tagIds []string) map[string]string {
	return getTagNames(tagIds, det)
}

func encodeDecode(rawVal interface{}, output interface{}) error {
	b := bytes.NewBuffer(nil)
	if err := json.NewEncoder(b).Encode(rawVal); err != nil {
		return err
	}
	if err := json.NewDecoder(b).Decode(output); err != nil {
		return err
	}
	return nil
}

// GetGameDownloads extracts downloads for languages and operating systems
// from GOG.com details. Currently, this data is stored as array of arrays,
// where the inner array holds the language as the first element and the object
// containing Windows, Mac, Linux downloads as the second element.
// To process that we go through the array, save the language of the current
// element, then encode/decode the next elements using a LangOSDownloads struct.
func (det *Details) GetGameDownloads() ([]LangOSManualDownloads, error) {

	langOsManDls := make([]LangOSManualDownloads, 0, len(det.Downloads))

	for _, dl := range det.Downloads {

		lang := ""
		//GOG.com has the following formats in details:
		//- games:
		//	[["Native_Language_Name",{"windows":[ManualDownload,...],"mac":[ManualDownload,...],"linux":[ManualDownload,...]}],...]
		gameDls, ok := dl.([]interface{})
		if !ok {
			return langOsManDls, fmt.Errorf("gog_types: cannot unmarshal game details download as []interface{}")
		}

		for i, dlPart := range gameDls {
			if i == 0 {
				lang = dlPart.(string)
			} else {
				var losDl *LangOSManualDownloads
				if err := encodeDecode(dlPart, &losDl); err != nil {
					return langOsManDls, err
				}
				losDl.Language = lang
				langOsManDls = append(langOsManDls, *losDl)
			}
		}
	}

	return langOsManDls, nil
}

func (det *Details) GetGOGRelease() string {
	if det.ReleaseTimestamp == 0 {
		return ""
	}
	//see comment for ApiProduct.GetGOGRelease that explains those special dates
	if det.ReleaseTimestamp == 693612000 ||
		det.ReleaseTimestamp == 978299999 {
		return ""
	}
	if l, err := time.LoadLocation("Europe/Nicosia"); err == nil {
		t := time.Unix(det.ReleaseTimestamp, 0)
		return t.In(l).Format("2006.01.02")
	}
	return ""
}

func (det *Details) GetForumUrl() string {
	return urlPathFromLink(det.ForumLink)
}

func (det *Details) GetChangelog() string {
	return det.Changelog
}

func (det *Details) GetOperatingSystems() ([]string, error) {

	downloads, err := det.GetGameDownloads()
	if err != nil {
		return nil, err
	}

	osm := make(map[string]any)

	for _, langDownload := range downloads {
		if len(langDownload.Windows) > 0 {
			osm["Windows"] = nil
		}
		if len(langDownload.Mac) > 0 {
			osm["macOS"] = nil
		}
		if len(langDownload.Linux) > 0 {
			osm["Linux"] = nil
		}
	}

	operatingSystems := make([]string, 0, len(osm))

	for ios := range osm {
		operatingSystems = append(operatingSystems, ios)
	}

	return operatingSystems, nil
}
