package egs_integration

import (
	"encoding/base64"
	"encoding/json/v2"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/arelate/southern_light/vangogh_integration"
	"github.com/boggydigital/camino"
	"github.com/boggydigital/coost"
	"github.com/boggydigital/kevlar"
)

const (
	// https://github.com/MixV2/EpicResearch/blob/master/docs/auth/auth_clients.md
	clientId     = "34a02cf8f4414e29b15921876da36f9a" // launcherAppClient2
	clientSecret = "daafbccc737745039dffe53d94fc76cf" // launcherAppClient2
)

const (
	UserAgent = "UELauncher/15.18.2-29993784+++Portal+Release-Live Windows/10.0.19041.1.256.64bit"
)

const (
	egsCookiesFilename = "egs-cookies.json"
	egsTokenKey        = "egs-token"
	jsonCatalogItemPfx = "{\"id\""
)

type GrantType string

const (
	GrantTypeRefreshToken      GrantType = "refresh_token"
	GrantTypeExchangeToken               = "exchange_code"
	GrantTypeAuthorizationCode           = "authorization_code"
	GrantTypeClientCredentials           = "client_credentials"
)

var egsClient *http.Client
var egsTokenVerifiedRecently bool

var (
	eosOverlayGameAsset = GameAsset{
		AppName:       "98bc04bc842e4906993fd6d6644ffb8d",
		LabelName:     "Epic Online Services Overlay",
		CatalogItemId: "cc15684f44d849e89e9bf4cec0508b68",
		Namespace:     "302e5ede476149b1bc3e4fe6ae45e50e",
	}

	eosHelperGameAsset = GameAsset{
		AppName:       "c9e2eb9993a1496c99dc529b49a07339",
		LabelName:     "Epic Online Services Helper",
		Namespace:     "302e5ede476149b1bc3e4fe6ae45e50e",
		CatalogItemId: "1108a9c0af47438da91331753b22ea21",
	}
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

func GetApiRedirect(client *http.Client) (io.ReadCloser, error) {

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

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	return resp.Body, nil
}

func PostToken(token string, grantType GrantType, client *http.Client) (io.ReadCloser, error) {

	aotUrl := AccountApiOauthTokenUrl()

	payload := make(url.Values)
	payload.Add("grant_type", string(grantType))
	payload.Add("token_type", "eg1")

	switch grantType {
	case GrantTypeRefreshToken:
		payload.Add("refresh_token", token)
	case GrantTypeExchangeToken:
		payload.Add("exchange_code", token)
	case GrantTypeAuthorizationCode:
		payload.Add("code", token)
	case GrantTypeClientCredentials:
		// do nothing
	}

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

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	return resp.Body, nil
}

func doResponse(req *http.Request, token string, client *http.Client) (io.ReadCloser, error) {

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

	return resp.Body, nil
}

func getResponse(u *url.URL, token string, client *http.Client) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	return doResponse(req, token, client)
}

func GetVerifyToken(token string, client *http.Client) (io.ReadCloser, error) {

	aovUrl := AccountApiOauthVerifyUrl()

	return getResponse(aovUrl, token, client)
}

func DeleteToken(token string, client *http.Client) error {

	aokUrl := AccountApiOauthKillUrl(token)

	req, err := http.NewRequest(http.MethodDelete, aokUrl.String(), http.NoBody)
	if err != nil {
		return err
	}

	readCloser, err := doResponse(req, token, client)
	if err != nil {
		return err
	}

	return readCloser.Close()
}

func GetClient() (*http.Client, error) {

	if egsClient == nil {
		cookiesDir := camino.GetRel(vangogh_integration.Cookies, vangogh_integration.Metadata)
		egsCookiePath := filepath.Join(cookiesDir, egsCookiesFilename)

		jar, err := coost.Read(HostUrl(), egsCookiePath)
		if err != nil {
			return nil, err
		}

		egsClient = http.DefaultClient
		egsClient.Jar = jar
	}

	return egsClient, nil
}

func PostStoreToken(token string, grantType GrantType) error {

	var err error

	var client *http.Client
	client, err = GetClient()
	if err != nil {
		return err
	}

	tokensDir := camino.GetRel(vangogh_integration.Tokens, vangogh_integration.Metadata)
	kvTokens, err := kevlar.New(tokensDir, kevlar.JsonExt)
	if err != nil {
		return err
	}

	var rcPostTokenResponse io.ReadCloser

	rcPostTokenResponse, err = PostToken(token, grantType, client)
	if err != nil {
		return err
	}

	defer rcPostTokenResponse.Close()

	return kvTokens.Set(egsTokenKey, rcPostTokenResponse)
}

func GetAccessToken(cookieStr string) error {

	cookiesDir := camino.GetRel(vangogh_integration.Cookies, vangogh_integration.Metadata)
	egsCookiePath := filepath.Join(cookiesDir, egsCookiesFilename)

	tokensDir := camino.GetRel(vangogh_integration.Tokens, vangogh_integration.Metadata)
	kvTokens, err := kevlar.New(tokensDir, kevlar.JsonExt)
	if err != nil {
		return err
	}

	if err = coost.Import(cookieStr, HostUrl(), egsCookiePath); err != nil {
		return err
	}

	if kvTokens.Has(egsTokenKey) {
		if err = kvTokens.Cut(egsTokenKey); err != nil {
			return err
		}
	}

	var client *http.Client
	client, err = GetClient()
	if err != nil {
		return err
	}

	var apiRedirectResponse GetApiRedirectResponse
	var rcApiRedirectResponse io.ReadCloser

	rcApiRedirectResponse, err = GetApiRedirect(client)
	if err != nil {
		return err
	}

	defer rcApiRedirectResponse.Close()

	if err = json.UnmarshalRead(rcApiRedirectResponse, &apiRedirectResponse); err != nil {
		return err
	}

	return PostStoreToken(apiRedirectResponse.AuthorizationCode, GrantTypeAuthorizationCode)
}

func RefreshToken(refreshToken string) error {
	if refreshToken == "" {
		return errors.New("refresh token not present")
	}
	return PostStoreToken(refreshToken, GrantTypeRefreshToken)
}

func GetStoredPostTokenResponse() (*PostTokenResponse, error) {
	tokensDir := camino.GetRel(vangogh_integration.Tokens, vangogh_integration.Metadata)
	kvTokens, err := kevlar.New(tokensDir, kevlar.JsonExt)
	if err != nil {
		return nil, err
	}

	var rcEgsToken io.ReadCloser
	rcEgsToken, err = kvTokens.Get(egsTokenKey)
	if err != nil {
		return nil, err
	}
	defer rcEgsToken.Close()

	var ptr PostTokenResponse
	if err = json.UnmarshalRead(rcEgsToken, &ptr); err != nil {
		return nil, err
	}

	return &ptr, nil
}

func VerifyToken(client *http.Client) (*PostTokenResponse, error) {

	ptr, err := GetStoredPostTokenResponse()
	if err != nil {
		return nil, err
	}

	if ptr.AccessToken == "" {
		return nil, errors.New("empty access token, re-connect EGS")
	}

	if egsTokenVerifiedRecently {
		return ptr, nil
	}

	if ptr.ExpiresAt.Sub(time.Now()) < time.Hour {
		if err = RefreshToken(ptr.RefreshToken); err != nil {
			return nil, err
		}

		if ptr, err = GetStoredPostTokenResponse(); err != nil {
			return nil, err
		}
	}

	var rcVerifyTokenResponse io.ReadCloser

	rcVerifyTokenResponse, err = GetVerifyToken(ptr.AccessToken, client)
	if err != nil {
		return nil, err
	}

	defer rcVerifyTokenResponse.Close()

	var verifyTokenResponse GetVerifyTokenResponse

	if err = json.UnmarshalRead(rcVerifyTokenResponse, &verifyTokenResponse); err != nil {
		return nil, err
	}

	if verifyTokenResponse.Token == "" {
		return nil, errors.New("empty access token, re-connect EGS")
	}

	if ptr.ExpiresAt.Sub(time.Now()) > time.Hour*3 {
		egsTokenVerifiedRecently = true
	}

	return ptr, nil
}

func SetupConnection(cookieStr string, reset bool) error {

	var err error

	if reset {
		if err = ResetConnection(); err != nil {
			return err
		}
	}

	if cookieStr != "" {
		if err = GetAccessToken(cookieStr); err != nil {
			return err
		}
	}

	client, err := GetClient()
	if err != nil {
		return err
	}

	_, err = VerifyToken(client)
	return err
}

func ResetConnection() error {

	cookiesDir := camino.GetRel(vangogh_integration.Cookies, vangogh_integration.Metadata)
	egsCookiePath := filepath.Join(cookiesDir, egsCookiesFilename)

	tokensDir := camino.GetRel(vangogh_integration.Tokens, vangogh_integration.Metadata)
	kvTokens, err := kevlar.New(tokensDir, kevlar.JsonExt)
	if err != nil {
		return err
	}

	if _, err = os.Stat(egsCookiePath); err == nil {
		if err = os.Remove(egsCookiePath); err != nil {
			return nil
		}
	}

	if kvTokens.Has(egsTokenKey) {
		if err = kvTokens.Cut(egsTokenKey); err != nil {
			return err
		}
	}

	return nil
}
