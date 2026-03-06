package egs_integration

import (
	"encoding/base64"
	"encoding/json/v2"
	"errors"
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
	UserAgent = "UELauncher/11.0.1-14907503+++Portal+Release-Live Windows/10.0.19041.1.256.64bit"
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
	DisplayName      string    `json:"displayName"`
	App              string    `json:"app"`
	InAppId          string    `json:"in_app_id"`
	Acr              string    `json:"acr"`
	AuthTime         time.Time `json:"auth_time"`
}

type GetVerifyTokenResponse struct {
	Token          string    `json:"token"`
	SessionId      string    `json:"session_id"`
	TokenType      string    `json:"token_type"`
	ClientId       string    `json:"client_id"`
	InternalClient bool      `json:"internal_client"`
	ClientService  string    `json:"client_service"`
	AccountId      string    `json:"account_id"`
	ExpiresIn      int       `json:"expires_in"`
	ExpiresAt      time.Time `json:"expires_at"`
	AuthMethod     string    `json:"auth_method"`
	DisplayName    string    `json:"display_name"`
	App            string    `json:"app"`
	InAppId        string    `json:"in_app_id"`
	Scope          any       `json:"scope"`
	Acr            string    `json:"acr"`
	AuthTime       time.Time `json:"auth_time"`
}

func GetApiRedirect(client *http.Client) (*GetApiRedirectResponse, error) {

	aruUrl := ApiRedirectUrl()

	req, err := http.NewRequest(http.MethodGet, aruUrl.String(), http.NoBody)
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

	aotUrl := AccountApiOauthTokenUrl()

	payload := make(url.Values)
	payload.Add("grant_type", "authorization_code")
	payload.Add("code", authorizationCode)
	payload.Add("token_type", "eg1")

	req, err := http.NewRequest(http.MethodPost, aotUrl.String(), strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}

	clientIdSecret := clientId + ":" + clientSecret
	base64cis := base64.StdEncoding.EncodeToString([]byte(clientIdSecret))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+base64cis)

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

func doResponse(req *http.Request, token string, client *http.Client) (*http.Response, error) {

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	return resp, nil
}

func getResponse(u *url.URL, token string, client *http.Client) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	return doResponse(req, token, client)
}

func GetVerifyToken(token string, client *http.Client) (*GetVerifyTokenResponse, error) {

	aovUrl := AccountApiOauthVerifyUrl()

	resp, err := getResponse(aovUrl, token, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getVerifyTokenResponse GetVerifyTokenResponse

	if err = json.UnmarshalRead(resp.Body, &getVerifyTokenResponse); err != nil {
		return nil, err
	}

	return &getVerifyTokenResponse, nil
}

func DeleteToken(token string, client *http.Client) error {

	aokUrl := AccountApiOauthKillUrl(token)

	req, err := http.NewRequest(http.MethodDelete, aokUrl.String(), http.NoBody)
	if err != nil {
		return err
	}

	resp, err := doResponse(req, token, client)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
