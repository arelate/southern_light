package egs_integration

import (
	"io"
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

func GetUserEntitlements(accountId string, token string, start, count int, client *http.Client) (io.ReadCloser, error) {

	entUrl := EntitlementsUrl(accountId, start, count)

	return getResponse(entUrl, token, client)
}
