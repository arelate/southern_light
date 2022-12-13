// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_integration

import (
	"strconv"
	"time"
)

type AccountProduct struct {
	IsGalaxyCompatible bool     `json:"isGalaxyCompatible"`
	Tags               []string `json:"tags"`
	Id                 int      `json:"id"`
	Availability       struct {
		IsAvailable          bool `json:"isAvailable"`
		IsAvailableInAccount bool `json:"isAvailableInAccount"`
	} `json:"availability"`
	Title   string `json:"title"`
	Image   string `json:"image"`
	Url     string `json:"url"`
	WorksOn struct {
		Windows bool `json:"Windows"`
		Mac     bool `json:"Mac"`
		Linux   bool `json:"Linux"`
	} `json:"worksOn"`
	Category     string `json:"category"`
	Rating       int    `json:"rating"`
	IsComingSoon bool   `json:"isComingSoon"`
	IsMovie      bool   `json:"isMovie"`
	IsGame       bool   `json:"isGame"`
	Slug         string `json:"slug"`
	Updates      int    `json:"updates"`
	IsNew        bool   `json:"isNew"`
	DlcCount     int    `json:"dlcCount"`
	ReleaseDate  struct {
		Date         string `json:"date"`
		TimezoneType int    `json:"timezone_type"`
		Timezone     string `json:"timezone"`
	} `json:"releaseDate"`
	IsBaseProductMissing bool `json:"isBaseProductMissing"`
	IsHidingDisabled     bool `json:"isHidingDisabled"`
	IsInDevelopment      bool `json:"isInDevelopment"`
	IsHidden             bool `json:"isHidden"`
}

func (ap *AccountProduct) GetId() int {
	return ap.Id
}

func (ap *AccountProduct) GetTitle() string {
	return ap.Title
}

func (ap *AccountProduct) GetImage() string {
	return ap.Image
}

func (ap *AccountProduct) GetRating() string {
	return strconv.Itoa(ap.Rating)
}

func (ap *AccountProduct) GetTagIds() []string {
	return ap.Tags
}

func (ap *AccountProduct) GetOperatingSystems() []string {
	os := make([]string, 0, 3)
	if ap.WorksOn.Windows {
		os = append(os, windows)
	}
	if ap.WorksOn.Mac {
		os = append(os, macos)
	}
	if ap.WorksOn.Linux {
		os = append(os, linux)
	}
	return os
}

func (ap *AccountProduct) GetSlug() string {
	return normalizeSlug(ap.Slug)
}

func (ap *AccountProduct) GetGlobalRelease() string {
	if l, err := time.LoadLocation(ap.ReleaseDate.Timezone); err == nil {
		if t, err := time.ParseInLocation("2006-01-02 15:04:05.0", ap.ReleaseDate.Date, l); err == nil {
			return t.Format("2006.01.02")
		}
	}
	return ""
}

func (ap *AccountProduct) GetStoreUrl() string {
	return ap.Url
}
