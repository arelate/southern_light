package epic_games

import (
	"encoding/json/v2"
	"net/http"
	"time"
)

type LibraryItems struct {
	ResponseMetadata struct {
		NextCursor string `json:"nextCursor"`
		StateToken string `json:"stateToken"`
	} `json:"responseMetadata"`
	Records []struct {
		Namespace       string        `json:"namespace"`
		CatalogItemId   string        `json:"catalogItemId"`
		AppName         string        `json:"appName"`
		Country         string        `json:"country"`
		Platform        []string      `json:"platform"`
		ProductId       string        `json:"productId"`
		SandboxName     string        `json:"sandboxName"`
		SandboxType     string        `json:"sandboxType"`
		RecordType      string        `json:"recordType"`
		AcquisitionDate time.Time     `json:"acquisitionDate"`
		Dependencies    []interface{} `json:"dependencies"`
		AvailableDate   time.Time     `json:"availableDate,omitempty"`
	} `json:"records"`
}

func GetLibraryItems(cursor string, token string, client *http.Client) (*LibraryItems, error) {

	liUrl := LibraryItemsUrl(true, cursor)

	resp, err := getResponse(liUrl, token, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var libraryItems LibraryItems
	if err = json.UnmarshalRead(resp.Body, &libraryItems); err != nil {
		return nil, err
	}

	return &libraryItems, nil
}
