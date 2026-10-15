package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestResourceTokenListModel(t *testing.T) {
	model := ResourceTokenList{Total: 5, Tokens: []ResourceToken{ResourceToken{Id: "bb8ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", ResourceId: "5e5ea5c168bb8:5e5ea5c168bb8", ResourceType: "files", Expire: "2020-10-15T06:38:00.000+00:00", Secret: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", AccessedAt: "2020-10-15T06:38:00.000+00:00"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result ResourceTokenList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestResourceTokenListModelDefaults(t *testing.T) {
	var fresh ResourceTokenList
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Total)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing total to default to %v, got %v", want, got)
		}
	}

	reused := ResourceTokenList{Total: 5, Tokens: []ResourceToken{ResourceToken{Id: "bb8ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", ResourceId: "5e5ea5c168bb8:5e5ea5c168bb8", ResourceType: "files", Expire: "2020-10-15T06:38:00.000+00:00", Secret: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", AccessedAt: "2020-10-15T06:38:00.000+00:00"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
