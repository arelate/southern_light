package epic_games

import (
	"encoding/json/v2"
	"net/http"
	"net/url"

	"github.com/boggydigital/dolo"
)

const (
	// https://github.com/MixV2/EpicResearch/blob/master/docs/auth/auth_clients.md
	clientId     = "34a02cf8f4414e29b15921876da36f9a" // launcherAppClient2
	clientSecret = "daafbccc737745039dffe53d94fc76cf" // launcherAppClient2
)

func ApiRedirectUrl() *url.URL {
	u := new(url.URL{
		Scheme: httpsScheme,
		Host:   epicGamesHost,
		Path:   apiRedirectPath,
	})

	q := u.Query()

	q.Set("clientId", clientId)
	q.Set("responseType", "code")

	u.RawQuery = q.Encode()

	return u
}

func GetAuthorizationCode(client *http.Client) (string, error) {

	aruUrl := ApiRedirectUrl().String()

	req, err := http.NewRequest(http.MethodGet, aruUrl, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", dolo.DefaultUserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var apiRedirectResponse ApiRedirectResponse
	if err = json.UnmarshalRead(resp.Body, &apiRedirectResponse); err != nil {
		return "", err
	}

	return apiRedirectResponse.AuthorizationCode, nil
}

type ApiRedirectResponse struct {
	AuthorizationCode string `json:"authorizationCode"`
	//Warning           string      `json:"warning"`
	//RedirectUrl       string      `json:"redirectUrl"`
	//ExchangeCode      any `json:"exchangeCode"`
	//Sid               any `json:"sid"`
}
