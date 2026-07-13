package models

import (
    "encoding/json"
    "testing"
)

func TestActivityEventModel(t *testing.T) {
    model := ActivityEvent{        Id: "5e5ea5c16897e",        ActorType: "user",        ActorId: "610fc2f985ee0",        ActorEmail: "john@appwrite.io",        ActorName: "John Doe",        ResourceParent: "database/ID",        ResourceType: "collection",        ResourceId: "610fc2f985ee0",        Resource: "collections/610fc2f985ee0",        Event: "account.sessions.create",        UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36",        Ip: "127.0.0.1",        Mode: "admin",        Country: "US",        Time: "2020-10-15T06:38:00.000+00:00",        ProjectId: "610fc2f985ee0",        TeamId: "610fc2f985ee0",        Hostname: "appwrite.io",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result ActivityEvent
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.ActorType != model.ActorType {
        t.Errorf("Expected ActorType %v, got %v", model.ActorType, result.ActorType)
    }
    if result.ActorId != model.ActorId {
        t.Errorf("Expected ActorId %v, got %v", model.ActorId, result.ActorId)
    }
    if result.ActorEmail != model.ActorEmail {
        t.Errorf("Expected ActorEmail %v, got %v", model.ActorEmail, result.ActorEmail)
    }
    if result.ActorName != model.ActorName {
        t.Errorf("Expected ActorName %v, got %v", model.ActorName, result.ActorName)
    }
    if result.ResourceParent != model.ResourceParent {
        t.Errorf("Expected ResourceParent %v, got %v", model.ResourceParent, result.ResourceParent)
    }
    if result.ResourceType != model.ResourceType {
        t.Errorf("Expected ResourceType %v, got %v", model.ResourceType, result.ResourceType)
    }
    if result.ResourceId != model.ResourceId {
        t.Errorf("Expected ResourceId %v, got %v", model.ResourceId, result.ResourceId)
    }
    if result.Resource != model.Resource {
        t.Errorf("Expected Resource %v, got %v", model.Resource, result.Resource)
    }
    if result.Event != model.Event {
        t.Errorf("Expected Event %v, got %v", model.Event, result.Event)
    }
    if result.UserAgent != model.UserAgent {
        t.Errorf("Expected UserAgent %v, got %v", model.UserAgent, result.UserAgent)
    }
    if result.Ip != model.Ip {
        t.Errorf("Expected Ip %v, got %v", model.Ip, result.Ip)
    }
    if result.Mode != model.Mode {
        t.Errorf("Expected Mode %v, got %v", model.Mode, result.Mode)
    }
    if result.Country != model.Country {
        t.Errorf("Expected Country %v, got %v", model.Country, result.Country)
    }
    if result.Time != model.Time {
        t.Errorf("Expected Time %v, got %v", model.Time, result.Time)
    }
    if result.ProjectId != model.ProjectId {
        t.Errorf("Expected ProjectId %v, got %v", model.ProjectId, result.ProjectId)
    }
    if result.TeamId != model.TeamId {
        t.Errorf("Expected TeamId %v, got %v", model.TeamId, result.TeamId)
    }
    if result.Hostname != model.Hostname {
        t.Errorf("Expected Hostname %v, got %v", model.Hostname, result.Hostname)
    }}
