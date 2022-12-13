package gog_integration

import (
	"net/url"
	"path"
)

const supportedVideoProvider = "youtube"

func youtubeVideoId(videoRawUrl string) string {
	if videoRawUrl == "" {
		return videoRawUrl
	}

	videoUrl, err := url.Parse(videoRawUrl)
	if err != nil {
		return ""
	}

	_, videoId := path.Split(videoUrl.Path)

	return videoId
}
