package epic_games

import (
	"io"
	"net/http"
)

func GetCatalogItem(namespace, itemId string, accessToken string, client *http.Client) (string, error) {

	catalogItemUrl := CatalogItemUrl(namespace, itemId, true, true, "US", "en")

	req, err := http.NewRequest(http.MethodGet, catalogItemUrl.String(), http.NoBody)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bts), nil
}
