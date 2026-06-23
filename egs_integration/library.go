package egs_integration

import (
	"io"
	"net/http"
	"time"
)

type LibraryItems struct {
	ResponseMetadata struct {
		NextCursor string `json:"nextCursor"`
		StateToken string `json:"stateToken"`
	} `json:"responseMetadata"`
	Records []struct {
		Namespace       string    `json:"namespace"`
		CatalogItemId   string    `json:"catalogItemId"`
		AppName         string    `json:"appName"`
		Country         string    `json:"country"`
		Platform        []string  `json:"platform"`
		ProductId       string    `json:"productId"`
		SandboxName     string    `json:"sandboxName"`
		SandboxType     string    `json:"sandboxType"`
		RecordType      string    `json:"recordType"`
		AcquisitionDate time.Time `json:"acquisitionDate"`
		Dependencies    []any     `json:"dependencies"`
		AvailableDate   time.Time `json:"availableDate"`
	} `json:"records"`
}

func GetLibraryItems(cursor string, token string, client *http.Client) (io.ReadCloser, error) {

	liUrl := LibraryItemsUrl(true, cursor)

	return getResponse(liUrl, token, client)
}
