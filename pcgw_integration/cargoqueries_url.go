package pcgw_integration

import (
	"fmt"
	"github.com/arelate/southern_light"
	"net/url"
	"strings"
)

const (
	actionParam = "action"
	tablesParam = "tables"
	fieldsParam = "fields"
	whereParam  = "where"
	formatParam = "format"
)

func PageIdCargoQueryUrl(gogId string) *url.URL {

	u := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   pcgwHost,
		Path:   apiPath,
	}

	q := u.Query()
	q.Set(actionParam, "cargoquery")
	q.Set(tablesParam, "Infobox_game")
	q.Set(fieldsParam, "_pageID=PageID")
	q.Set(whereParam, fmt.Sprintf("Infobox_game.GOGcom_ID HOLDS \"%s\"", gogId))
	q.Set(formatParam, "json")
	u.RawQuery = q.Encode()

	return u
}

func EngineCargoQueryUrl(pageId string) *url.URL {

	u := &url.URL{
		Scheme: southern_light.HttpsScheme,
		Host:   pcgwHost,
		Path:   apiPath,
	}

	q := u.Query()
	q.Set(actionParam, "cargoquery")
	q.Set(tablesParam, "Infobox_game_engine")
	q.Set(fieldsParam, strings.Join([]string{"Engine", "Build"}, ","))
	q.Set(whereParam, fmt.Sprintf("_pageId=%s", pageId))
	q.Set(formatParam, "json")
	u.RawQuery = q.Encode()

	return u
}
