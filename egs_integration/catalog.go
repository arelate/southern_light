package egs_integration

import (
	"io"
	"net/http"
	"time"
)

type TypeValue struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

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
	Namespace        string               `json:"namespace"`
	Status           string               `json:"status"`
	CreationDate     time.Time            `json:"creationDate"`
	LastModifiedDate time.Time            `json:"lastModifiedDate"`
	CustomAttributes map[string]TypeValue `json:"customAttributes"`
	EntitlementName  string               `json:"entitlementName"`
	EntitlementType  string               `json:"entitlementType"`
	ItemType         string               `json:"itemType"`
	ReleaseInfo      []struct {
		Id       string   `json:"id"`
		AppId    string   `json:"appId"`
		Platform []string `json:"platform"`
	} `json:"releaseInfo"`
	Developer    string   `json:"developer"`
	DeveloperId  string   `json:"developerId"`
	EulaIds      []string `json:"eulaIds"`
	EndOfSupport bool     `json:"endOfSupport"`
	//MainGameItem        CatalogItem   `json:"mainGameItem"`
	MainGameItemList    []CatalogItem `json:"mainGameItemList"`
	DlcItemList         []CatalogItem `json:"dlcItemList"`
	EsrbGameRatingValue string        `json:"esrbGameRatingValue"`
	AgeGatings          struct {
	} `json:"ageGatings"`
	Unsearchable bool `json:"unsearchable"`
}

func GetCatalogItem(namespace, itemId string, token string, client *http.Client) (io.ReadCloser, error) {

	ciUrl := CatalogItemUrl(namespace, itemId, true, true, "US", "en")

	return getResponse(ciUrl, token, client)
}
