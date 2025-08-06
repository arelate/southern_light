package steam_integration

type GetAppListV2Response struct {
	AppList struct {
		Apps []struct {
			AppId uint32 `json:"appid"`
			Name  string `json:"name"`
		} `json:"apps"`
	} `json:"applist"`
}
