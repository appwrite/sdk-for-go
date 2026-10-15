package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDedicatedDatabaseBranchModel(t *testing.T) {
	model := DedicatedDatabaseBranch{BranchId: "a1b2c3d4", BranchName: "branch-a1b2c3d4", Namespace: "db-myproject-mydb-branch-a1b2c3d4", ExpiresAt: 1711411200, Host: "db-myproject-mydb-a1b2c3d4.fra.appwrite.center", Database: "db-myproject-mydb-a1b2c3d4", Username: "appwrite", Password: "********", Ssl: true, Engine: "postgresql", ConnectionString: "postgresql://appwrite:****@db-myproject-mydb-a1b2c3d4.fra.appwrite.center:5432/db-myproject-mydb-a1b2c3d4?sslmode=require"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseBranch
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.BranchId != model.BranchId {
		t.Errorf("Expected BranchId %v, got %v", model.BranchId, result.BranchId)
	}
	if result.BranchName != model.BranchName {
		t.Errorf("Expected BranchName %v, got %v", model.BranchName, result.BranchName)
	}
	if result.Namespace != model.Namespace {
		t.Errorf("Expected Namespace %v, got %v", model.Namespace, result.Namespace)
	}
	if result.ExpiresAt != model.ExpiresAt {
		t.Errorf("Expected ExpiresAt %v, got %v", model.ExpiresAt, result.ExpiresAt)
	}
	if result.Host != model.Host {
		t.Errorf("Expected Host %v, got %v", model.Host, result.Host)
	}
	if result.Database != model.Database {
		t.Errorf("Expected Database %v, got %v", model.Database, result.Database)
	}
	if result.Username != model.Username {
		t.Errorf("Expected Username %v, got %v", model.Username, result.Username)
	}
	if result.Password != model.Password {
		t.Errorf("Expected Password %v, got %v", model.Password, result.Password)
	}
	if result.Ssl != model.Ssl {
		t.Errorf("Expected Ssl %v, got %v", model.Ssl, result.Ssl)
	}
	if result.Engine != model.Engine {
		t.Errorf("Expected Engine %v, got %v", model.Engine, result.Engine)
	}
	if result.ConnectionString != model.ConnectionString {
		t.Errorf("Expected ConnectionString %v, got %v", model.ConnectionString, result.ConnectionString)
	}
}

func TestDedicatedDatabaseBranchModelDefaults(t *testing.T) {
	var fresh DedicatedDatabaseBranch
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.BranchId)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing branchId to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.BranchName)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing branchName to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Namespace)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing namespace to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ExpiresAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing expiresAt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Host)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing host to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Database)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing database to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Username)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing username to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Password)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing password to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Ssl)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ssl to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Engine)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing engine to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ConnectionString)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing connectionString to default to %v, got %v", want, got)
		}
	}

	reused := DedicatedDatabaseBranch{BranchId: "a1b2c3d4", BranchName: "branch-a1b2c3d4", Namespace: "db-myproject-mydb-branch-a1b2c3d4", ExpiresAt: 1711411200, Host: "db-myproject-mydb-a1b2c3d4.fra.appwrite.center", Database: "db-myproject-mydb-a1b2c3d4", Username: "appwrite", Password: "********", Ssl: true, Engine: "postgresql", ConnectionString: "postgresql://appwrite:****@db-myproject-mydb-a1b2c3d4.fra.appwrite.center:5432/db-myproject-mydb-a1b2c3d4?sslmode=require"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
