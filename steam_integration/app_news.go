package steam_integration

type GetNewsForAppResponse struct {
	AppNews AppNews `json:"appnews"`
}

type AppNews struct {
	AppId     uint32     `json:"appid"`
	NewsItems []NewsItem `json:"newsitems"`
	Count     uint32     `json:"count"`
}

type NewsItem struct {
	GId           string   `json:"gid"`
	Title         string   `json:"title"`
	Url           string   `json:"url"`
	IsExternalUrl bool     `json:"is_external_url"`
	Author        string   `json:"author"`
	Contents      string   `json:"contents"`
	FeedLabel     string   `json:"feedlabel"`
	Date          int64    `json:"date"`
	FeedName      string   `json:"feedname"`
	FeedType      int      `json:"feed_type"`
	AppId         uint32   `json:"appid"`
	Tags          []string `json:"tags"`
}
