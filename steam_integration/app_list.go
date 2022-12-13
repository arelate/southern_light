package steam_integration

type GetAppListResponse struct {
	AppList AppList `json:"applist"`
}

type AppList struct {
	Apps []App `json:"apps"`
}

type App struct {
	AppId uint32 `json:"appid"`
	Name  string `json:"name"`
}
