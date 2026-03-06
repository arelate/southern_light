package egs_integration

import (
	"encoding/json/v2"
	"errors"
	"net/http"
	"time"
)

type CatalogItem struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	KeyImages   []struct {
		Type         string    `json:"type"`
		Url          string    `json:"url"`
		Md5          string    `json:"md5"`
		Width        int       `json:"width"`
		Height       int       `json:"height"`
		Size         int       `json:"size"`
		UploadedDate time.Time `json:"uploadedDate"`
	} `json:"keyImages"`
	Categories []struct {
		Path string `json:"path"`
	} `json:"categories"`
	Namespace        string    `json:"namespace"`
	Status           string    `json:"status"`
	CreationDate     time.Time `json:"creationDate"`
	LastModifiedDate time.Time `json:"lastModifiedDate"`
	CustomAttributes struct {
		SupportedPlatforms struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"SupportedPlatforms"`
	} `json:"customAttributes"`
	EntitlementName     string        `json:"entitlementName"`
	EntitlementType     string        `json:"entitlementType"`
	ItemType            string        `json:"itemType"`
	Developer           string        `json:"developer"`
	DeveloperId         string        `json:"developerId"`
	EulaIds             []string      `json:"eulaIds"`
	EndOfSupport        bool          `json:"endOfSupport"`
	MainGameItemList    []interface{} `json:"mainGameItemList"`
	EsrbGameRatingValue string        `json:"esrbGameRatingValue"`
	AgeGatings          struct {
	} `json:"ageGatings"`
	Unsearchable bool `json:"unsearchable"`
}

func GetCatalogItem(namespace, itemId string, token string, client *http.Client) (*CatalogItem, error) {

	ciUrl := CatalogItemUrl(namespace, itemId, true, true, "US", "en")

	resp, err := getResponse(ciUrl, token, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var catalogItemResponse map[string]CatalogItem
	if err = json.UnmarshalRead(resp.Body, &catalogItemResponse); err != nil {
		return nil, err
	}

	if catalogItem, ok := catalogItemResponse[itemId]; ok {
		return &catalogItem, nil
	}

	return nil, errors.New("catalog item not found")
}
