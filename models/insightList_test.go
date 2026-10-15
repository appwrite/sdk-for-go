package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestInsightListModel(t *testing.T) {
	model := InsightList{Total: 5, Insights: []Insight{Insight{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ReportId: "5e5ea5c16897e", Type: "tablesDBIndex", Severity: "warning", Status: "active", ResourceType: "databases", ResourceId: "main", ParentResourceType: "tables", ParentResourceId: "orders", Title: "Missing index on collection orders", Summary: "Queries against `orders.status` are scanning the full collection.", Ctas: []InsightCTA{InsightCTA{Label: "Create missing index", Service: "tablesDB", Method: "createIndex", Params: map[string]interface{}{}}}}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result InsightList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestInsightListModelDefaults(t *testing.T) {
	var fresh InsightList
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

	reused := InsightList{Total: 5, Insights: []Insight{Insight{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ReportId: "5e5ea5c16897e", Type: "tablesDBIndex", Severity: "warning", Status: "active", ResourceType: "databases", ResourceId: "main", ParentResourceType: "tables", ParentResourceId: "orders", Title: "Missing index on collection orders", Summary: "Queries against `orders.status` are scanning the full collection.", Ctas: []InsightCTA{InsightCTA{Label: "Create missing index", Service: "tablesDB", Method: "createIndex", Params: map[string]interface{}{}}}}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
