package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAnalyticsAnnotationListModel(t *testing.T) {
	model := AnalyticsAnnotationList{Total: 5, Annotations: []AnalyticsAnnotation{AnalyticsAnnotation{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", PropertyId: "5e5ea5c16897e", Date: "2020-10-15T06:38:00.000+00:00", Content: "Launched the new pricing page", UserId: "5e5ea5c16897e"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AnalyticsAnnotationList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestAnalyticsAnnotationListModelDefaults(t *testing.T) {
	var fresh AnalyticsAnnotationList
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

	reused := AnalyticsAnnotationList{Total: 5, Annotations: []AnalyticsAnnotation{AnalyticsAnnotation{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", PropertyId: "5e5ea5c16897e", Date: "2020-10-15T06:38:00.000+00:00", Content: "Launched the new pricing page", UserId: "5e5ea5c16897e"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
