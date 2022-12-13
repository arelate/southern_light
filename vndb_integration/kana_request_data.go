package vndb_integration

type KanaRequestData struct {
	Filters []string `json:"filters"`
	Fields  string   `json:"fields"`
}

//curl https://api.vndb.org/kana/vn --header 'Content-Type: application/json' --data '{
//"filters": ["id", "=", "v14018"],
//"fields": "tags.name,tags.rating,tags.spoiler,rating"
//}'
