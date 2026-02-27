package epic_games

import (
	"encoding/json/v2"
	"net/http"
	"time"
)

type Entitlement struct {
	Id               string    `json:"id"`
	EntitlementName  string    `json:"entitlementName"`
	Namespace        string    `json:"namespace"`
	CatalogItemId    string    `json:"catalogItemId"`
	Store            string    `json:"store"`
	AccountId        string    `json:"accountId"`
	IdentityId       string    `json:"identityId"`
	EntitlementType  string    `json:"entitlementType"`
	GrantDate        time.Time `json:"grantDate"`
	Consumable       bool      `json:"consumable"`
	Status           string    `json:"status"`
	Active           bool      `json:"active"`
	UseCount         int       `json:"useCount"`
	OriginalUseCount int       `json:"originalUseCount"`
	PlatformType     string    `json:"platformType"`
	Created          time.Time `json:"created"`
	Updated          time.Time `json:"updated"`
	GroupEntitlement bool      `json:"groupEntitlement"`
	Country          string    `json:"country"`
	ReadFromCache    bool      `json:"readFromCache"`
}

func GetEntitlements(accountId string, accessToken string, client *http.Client) ([]Entitlement, error) {

	entUrl := EntitlementsUrl(accountId)

	req, err := http.NewRequest(http.MethodGet, entUrl.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := doResponse(req, accessToken, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var entitlements []Entitlement
	if err = json.UnmarshalRead(resp.Body, &entitlements); err != nil {
		return nil, err
	}

	return entitlements, nil
}
