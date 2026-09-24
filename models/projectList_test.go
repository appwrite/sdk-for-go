package models

import (
	"encoding/json"
	"testing"
)

func TestProjectListModel(t *testing.T) {
	model := ProjectList{Total: 5, Projects: []Project{Project{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "New Project", TeamId: "1592981250", Region: "fra", SmtpEnabled: true, SmtpSenderName: "John Appwrite", SmtpSenderEmail: "john@appwrite.io", SmtpReplyToName: "Support Team", SmtpReplyToEmail: "support@appwrite.io", SmtpHost: "mail.appwrite.io", SmtpPort: 25, SmtpUsername: "emailuser", SmtpPassword: "smtp-password", SmtpSecure: "tls", PingCount: 1, PingedAt: "2020-10-15T06:38:00.000+00:00", Labels: []string{"test"}, Status: "active", Onboarding: map[string]interface{}{}, AuthMethods: []ProjectAuthMethod{ProjectAuthMethod{Id: "email-password", Enabled: true}}, Services: []ProjectService{ProjectService{Id: "sites", Enabled: true}}, Protocols: []ProjectProtocol{ProjectProtocol{Id: "graphql", Enabled: true}}, Blocks: []Block{Block{CreatedAt: "2020-10-15T06:38:00.000+00:00", ResourceType: "project", ResourceId: "5e5ea5c16897e", Mode: "readOnly"}}, ConsoleAccessedAt: "2020-10-15T06:38:00.000+00:00"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result ProjectList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}
