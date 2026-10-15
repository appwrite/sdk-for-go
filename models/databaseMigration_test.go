package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDatabaseMigrationModel(t *testing.T) {
	model := DatabaseMigration{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ProjectId: "5e5ea5c16897e", DatabaseId: "5e5ea5c16897e", Specification: "s-2vcpu-4gb", Phase: "pending", Attempt: 0, LastError: "string", LagDocuments: 0, ChangelogWatermark: 0, VerifiedAt: "2020-10-15T06:38:00.000+00:00", CutoverAt: "2020-10-15T06:38:00.000+00:00", SoakUntil: "2020-10-15T06:38:00.000+00:00", AutoCutover: true, CutoverRequested: true, Paused: true}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DatabaseMigration
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.CreatedAt != model.CreatedAt {
		t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
	}
	if result.UpdatedAt != model.UpdatedAt {
		t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
	}
	if result.ProjectId != model.ProjectId {
		t.Errorf("Expected ProjectId %v, got %v", model.ProjectId, result.ProjectId)
	}
	if result.DatabaseId != model.DatabaseId {
		t.Errorf("Expected DatabaseId %v, got %v", model.DatabaseId, result.DatabaseId)
	}
	if result.Specification != model.Specification {
		t.Errorf("Expected Specification %v, got %v", model.Specification, result.Specification)
	}
	if result.Phase != model.Phase {
		t.Errorf("Expected Phase %v, got %v", model.Phase, result.Phase)
	}
	if result.Attempt != model.Attempt {
		t.Errorf("Expected Attempt %v, got %v", model.Attempt, result.Attempt)
	}
	if result.LastError != model.LastError {
		t.Errorf("Expected LastError %v, got %v", model.LastError, result.LastError)
	}
	if result.LagDocuments != model.LagDocuments {
		t.Errorf("Expected LagDocuments %v, got %v", model.LagDocuments, result.LagDocuments)
	}
	if result.ChangelogWatermark != model.ChangelogWatermark {
		t.Errorf("Expected ChangelogWatermark %v, got %v", model.ChangelogWatermark, result.ChangelogWatermark)
	}
	if result.VerifiedAt != model.VerifiedAt {
		t.Errorf("Expected VerifiedAt %v, got %v", model.VerifiedAt, result.VerifiedAt)
	}
	if result.CutoverAt != model.CutoverAt {
		t.Errorf("Expected CutoverAt %v, got %v", model.CutoverAt, result.CutoverAt)
	}
	if result.SoakUntil != model.SoakUntil {
		t.Errorf("Expected SoakUntil %v, got %v", model.SoakUntil, result.SoakUntil)
	}
	if result.AutoCutover != model.AutoCutover {
		t.Errorf("Expected AutoCutover %v, got %v", model.AutoCutover, result.AutoCutover)
	}
	if result.CutoverRequested != model.CutoverRequested {
		t.Errorf("Expected CutoverRequested %v, got %v", model.CutoverRequested, result.CutoverRequested)
	}
	if result.Paused != model.Paused {
		t.Errorf("Expected Paused %v, got %v", model.Paused, result.Paused)
	}
}

func TestDatabaseMigrationModelDefaults(t *testing.T) {
	var fresh DatabaseMigration
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Id)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing $id to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.CreatedAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing $createdAt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.UpdatedAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing $updatedAt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ProjectId)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing projectId to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DatabaseId)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing databaseId to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Specification)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing specification to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"pending\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Phase)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing phase to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Attempt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing attempt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LastError)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing lastError to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LagDocuments)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing lagDocuments to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ChangelogWatermark)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing changelogWatermark to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.VerifiedAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing verifiedAt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.CutoverAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing cutoverAt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SoakUntil)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing soakUntil to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.AutoCutover)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing autoCutover to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.CutoverRequested)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing cutoverRequested to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Paused)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing paused to default to %v, got %v", want, got)
		}
	}

	reused := DatabaseMigration{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ProjectId: "5e5ea5c16897e", DatabaseId: "5e5ea5c16897e", Specification: "s-2vcpu-4gb", Phase: "pending", Attempt: 0, LastError: "string", LagDocuments: 0, ChangelogWatermark: 0, VerifiedAt: "2020-10-15T06:38:00.000+00:00", CutoverAt: "2020-10-15T06:38:00.000+00:00", SoakUntil: "2020-10-15T06:38:00.000+00:00", AutoCutover: true, CutoverRequested: true, Paused: true}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
