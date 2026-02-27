package epic_games

import (
	"encoding/base64"
	"encoding/json/v2"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// https://github.com/MixV2/EpicResearch/blob/master/docs/auth/auth_clients.md
	clientId     = "34a02cf8f4414e29b15921876da36f9a" // launcherAppClient2
	clientSecret = "daafbccc737745039dffe53d94fc76cf" // launcherAppClient2
)

const (
	UserAgent = "EpicGamesLauncher/12.2.9-16657426 UnrealEngine/4.23.0-14907503+++Portal+Release-Live"
)

type GetApiRedirectResponse struct {
	Warning           string `json:"warning"`
	RedirectUrl       string `json:"redirectUrl"`
	AuthorizationCode string `json:"authorizationCode"`
	ExchangeCode      any    `json:"exchangeCode"`
	Sid               any    `json:"sid"`
}

type PostTokenResponse struct {
	AccessToken      string    `json:"access_token"`
	ExpiresIn        int       `json:"expires_in"`
	ExpiresAt        time.Time `json:"expires_at"`
	TokenType        string    `json:"token_type"`
	RefreshToken     string    `json:"refresh_token"`
	RefreshExpires   int       `json:"refresh_expires"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at"`
	AccountId        string    `json:"account_id"`
	ClientId         string    `json:"client_id"`
	InternalClient   bool      `json:"internal_client"`
	ClientService    string    `json:"client_service"`
	Scope            any       `json:"scope"`
	Display          string    `json:"displayName"`
	App              string    `json:"app"`
	InAppId          string    `json:"in_app_id"`
	Acr              string    `json:"acr"`
	AuthTime         time.Time `json:"auth_time"`
}

func GetApiRedirect(client *http.Client) (*GetApiRedirectResponse, error) {

	aruUrl := ApiRedirectUrl().String()

	req, err := http.NewRequest(http.MethodGet, aruUrl, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var getApiRedirectResponse GetApiRedirectResponse
	if err = json.UnmarshalRead(resp.Body, &getApiRedirectResponse); err != nil {
		return nil, err
	}

	return &getApiRedirectResponse, nil
}

func PostToken(authorizationCode string, client *http.Client) (*PostTokenResponse, error) {

	aotUrl := AccountApiOauthToken().String()

	payload := make(url.Values)
	payload.Add("grant_type", "authorization_code")
	payload.Add("code", authorizationCode)

	req, err := http.NewRequest(http.MethodPost, aotUrl, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}

	clientCredentials := clientId + ":" + clientSecret
	b64cc := base64.StdEncoding.EncodeToString([]byte(clientCredentials))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "basic "+b64cc)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var postTokenReponse PostTokenResponse
	if err = json.UnmarshalRead(resp.Body, &postTokenReponse); err != nil {
		return nil, err
	}

	return &postTokenReponse, nil
}
