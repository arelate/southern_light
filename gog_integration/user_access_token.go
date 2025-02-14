package gog_integration

type UserAccessToken struct {
	AccessToken string `json:"accessToken"`
	Expiration  int    `json:"accessTokenExpiration"`
	ClientId    string `json:"accessTokenClientId"`
	UserId      string `json:"accessTokenUserId"`
}
