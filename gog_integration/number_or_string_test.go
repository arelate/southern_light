package gog_integration

import (
	"encoding/json/v2"
	"testing"
)

func TestAccountPageTagProductCount(t *testing.T) {
	// GOG returns tag productCount as a JSON number; older responses used strings.
	for _, tc := range []struct {
		name, body, want string
	}{
		{"number", `{"tags":[{"id":"158240544","name":"COMPLETED","productCount":0}],"appliedFilters":{"tags":[{"id":"1","name":"X","productCount":12}]}}`, "0"},
		{"string", `{"tags":[{"id":"158240544","name":"COMPLETED","productCount":"3"}],"appliedFilters":{"tags":[]}}`, "3"},
		{"null", `{"tags":[{"id":"1","name":"X","productCount":null}]}`, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var ap AccountPage
			if err := json.Unmarshal([]byte(tc.body), &ap); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if got := string(ap.Tags[0].ProductCount); got != tc.want {
				t.Fatalf("productCount = %q, want %q", got, tc.want)
			}
		})
	}

	var ap AccountPage
	if err := json.Unmarshal([]byte(`{"tags":[{"productCount":true}]}`), &ap); err == nil {
		t.Fatal("expected an error for a boolean productCount")
	}
}
