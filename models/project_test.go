package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestProjectModel(t *testing.T) {
	model := Project{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "New Project", TeamId: "1592981250", Region: "fra", SmtpEnabled: true, SmtpSenderName: "John Appwrite", SmtpSenderEmail: "john@appwrite.io", SmtpReplyToName: "Support Team", SmtpReplyToEmail: "support@appwrite.io", SmtpHost: "mail.appwrite.io", SmtpPort: 25, SmtpUsername: "emailuser", SmtpPassword: "smtp-password", SmtpSecure: "tls", LdapEnabled: true, LdapHosts: []string{"test"}, LdapPort: 636, LdapEncryption: "ldaps", LdapVerification: true, LdapCertificate: "string", LdapTimeout: 10, LdapBaseDn: "ou=people,dc=example,dc=com", LdapBindDn: "cn=appwrite,ou=services,dc=example,dc=com", LdapBindPassword: "ldap-password", LdapUserFilter: "(uid={username})", LdapRequiredGroupDn: "cn=staff,ou=groups,dc=example,dc=com", LdapIdAttribute: "entryUUID", LdapEmailAttribute: "mail", LdapNameAttribute: "cn", LdapPhoneAttribute: "mobile", PingCount: 1, PingedAt: "2020-10-15T06:38:00.000+00:00", Labels: []string{"test"}, Status: "active", Onboarding: map[string]interface{}{}, AuthMethods: []ProjectAuthMethod{ProjectAuthMethod{Id: "email-password", Enabled: true}}, Services: []ProjectService{ProjectService{Id: "sites", Enabled: true}}, Protocols: []ProjectProtocol{ProjectProtocol{Id: "graphql", Enabled: true}}, Blocks: []Block{Block{CreatedAt: "2020-10-15T06:38:00.000+00:00", ResourceType: "project", ResourceId: "5e5ea5c16897e", Mode: "readOnly"}}, ConsoleAccessedAt: "2020-10-15T06:38:00.000+00:00"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Project
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
	if result.Name != model.Name {
		t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
	}
	if result.TeamId != model.TeamId {
		t.Errorf("Expected TeamId %v, got %v", model.TeamId, result.TeamId)
	}
	if result.Region != model.Region {
		t.Errorf("Expected Region %v, got %v", model.Region, result.Region)
	}
	if result.SmtpEnabled != model.SmtpEnabled {
		t.Errorf("Expected SmtpEnabled %v, got %v", model.SmtpEnabled, result.SmtpEnabled)
	}
	if result.SmtpSenderName != model.SmtpSenderName {
		t.Errorf("Expected SmtpSenderName %v, got %v", model.SmtpSenderName, result.SmtpSenderName)
	}
	if result.SmtpSenderEmail != model.SmtpSenderEmail {
		t.Errorf("Expected SmtpSenderEmail %v, got %v", model.SmtpSenderEmail, result.SmtpSenderEmail)
	}
	if result.SmtpReplyToName != model.SmtpReplyToName {
		t.Errorf("Expected SmtpReplyToName %v, got %v", model.SmtpReplyToName, result.SmtpReplyToName)
	}
	if result.SmtpReplyToEmail != model.SmtpReplyToEmail {
		t.Errorf("Expected SmtpReplyToEmail %v, got %v", model.SmtpReplyToEmail, result.SmtpReplyToEmail)
	}
	if result.SmtpHost != model.SmtpHost {
		t.Errorf("Expected SmtpHost %v, got %v", model.SmtpHost, result.SmtpHost)
	}
	if result.SmtpPort != model.SmtpPort {
		t.Errorf("Expected SmtpPort %v, got %v", model.SmtpPort, result.SmtpPort)
	}
	if result.SmtpUsername != model.SmtpUsername {
		t.Errorf("Expected SmtpUsername %v, got %v", model.SmtpUsername, result.SmtpUsername)
	}
	if result.SmtpPassword != model.SmtpPassword {
		t.Errorf("Expected SmtpPassword %v, got %v", model.SmtpPassword, result.SmtpPassword)
	}
	if result.SmtpSecure != model.SmtpSecure {
		t.Errorf("Expected SmtpSecure %v, got %v", model.SmtpSecure, result.SmtpSecure)
	}
	if result.LdapEnabled != model.LdapEnabled {
		t.Errorf("Expected LdapEnabled %v, got %v", model.LdapEnabled, result.LdapEnabled)
	}
	if result.LdapPort != model.LdapPort {
		t.Errorf("Expected LdapPort %v, got %v", model.LdapPort, result.LdapPort)
	}
	if result.LdapEncryption != model.LdapEncryption {
		t.Errorf("Expected LdapEncryption %v, got %v", model.LdapEncryption, result.LdapEncryption)
	}
	if result.LdapVerification != model.LdapVerification {
		t.Errorf("Expected LdapVerification %v, got %v", model.LdapVerification, result.LdapVerification)
	}
	if result.LdapCertificate != model.LdapCertificate {
		t.Errorf("Expected LdapCertificate %v, got %v", model.LdapCertificate, result.LdapCertificate)
	}
	if result.LdapTimeout != model.LdapTimeout {
		t.Errorf("Expected LdapTimeout %v, got %v", model.LdapTimeout, result.LdapTimeout)
	}
	if result.LdapBaseDn != model.LdapBaseDn {
		t.Errorf("Expected LdapBaseDn %v, got %v", model.LdapBaseDn, result.LdapBaseDn)
	}
	if result.LdapBindDn != model.LdapBindDn {
		t.Errorf("Expected LdapBindDn %v, got %v", model.LdapBindDn, result.LdapBindDn)
	}
	if result.LdapBindPassword != model.LdapBindPassword {
		t.Errorf("Expected LdapBindPassword %v, got %v", model.LdapBindPassword, result.LdapBindPassword)
	}
	if result.LdapUserFilter != model.LdapUserFilter {
		t.Errorf("Expected LdapUserFilter %v, got %v", model.LdapUserFilter, result.LdapUserFilter)
	}
	if result.LdapRequiredGroupDn != model.LdapRequiredGroupDn {
		t.Errorf("Expected LdapRequiredGroupDn %v, got %v", model.LdapRequiredGroupDn, result.LdapRequiredGroupDn)
	}
	if result.LdapIdAttribute != model.LdapIdAttribute {
		t.Errorf("Expected LdapIdAttribute %v, got %v", model.LdapIdAttribute, result.LdapIdAttribute)
	}
	if result.LdapEmailAttribute != model.LdapEmailAttribute {
		t.Errorf("Expected LdapEmailAttribute %v, got %v", model.LdapEmailAttribute, result.LdapEmailAttribute)
	}
	if result.LdapNameAttribute != model.LdapNameAttribute {
		t.Errorf("Expected LdapNameAttribute %v, got %v", model.LdapNameAttribute, result.LdapNameAttribute)
	}
	if result.LdapPhoneAttribute != model.LdapPhoneAttribute {
		t.Errorf("Expected LdapPhoneAttribute %v, got %v", model.LdapPhoneAttribute, result.LdapPhoneAttribute)
	}
	if result.PingCount != model.PingCount {
		t.Errorf("Expected PingCount %v, got %v", model.PingCount, result.PingCount)
	}
	if result.PingedAt != model.PingedAt {
		t.Errorf("Expected PingedAt %v, got %v", model.PingedAt, result.PingedAt)
	}
	if result.Status != model.Status {
		t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
	}
	if result.ConsoleAccessedAt != model.ConsoleAccessedAt {
		t.Errorf("Expected ConsoleAccessedAt %v, got %v", model.ConsoleAccessedAt, result.ConsoleAccessedAt)
	}
}

func TestProjectModelDefaults(t *testing.T) {
	var fresh Project
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
		encoded, err := json.Marshal(fresh.Name)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing name to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.TeamId)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing teamId to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"fra\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Region)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing region to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SmtpEnabled)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing smtpEnabled to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SmtpSenderName)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing smtpSenderName to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SmtpSenderEmail)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing smtpSenderEmail to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SmtpReplyToName)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing smtpReplyToName to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SmtpReplyToEmail)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing smtpReplyToEmail to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SmtpHost)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing smtpHost to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SmtpPort)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing smtpPort to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SmtpUsername)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing smtpUsername to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SmtpPassword)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing smtpPassword to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SmtpSecure)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing smtpSecure to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapEnabled)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapEnabled to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[]"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapHosts)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapHosts to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("636"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapPort)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapPort to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"ldaps\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapEncryption)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapEncryption to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapVerification)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapVerification to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapCertificate)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapCertificate to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("10"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapTimeout)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapTimeout to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapBaseDn)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapBaseDn to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapBindDn)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapBindDn to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapBindPassword)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapBindPassword to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"(uid={username})\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapUserFilter)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapUserFilter to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapRequiredGroupDn)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapRequiredGroupDn to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"entryUUID\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapIdAttribute)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapIdAttribute to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"mail\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapEmailAttribute)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapEmailAttribute to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"cn\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapNameAttribute)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapNameAttribute to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.LdapPhoneAttribute)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ldapPhoneAttribute to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.PingCount)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing pingCount to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.PingedAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing pingedAt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[]"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Labels)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing labels to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"active\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Status)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing status to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[]"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Onboarding)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing onboarding to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ConsoleAccessedAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing consoleAccessedAt to default to %v, got %v", want, got)
		}
	}

	reused := Project{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "New Project", TeamId: "1592981250", Region: "fra", SmtpEnabled: true, SmtpSenderName: "John Appwrite", SmtpSenderEmail: "john@appwrite.io", SmtpReplyToName: "Support Team", SmtpReplyToEmail: "support@appwrite.io", SmtpHost: "mail.appwrite.io", SmtpPort: 25, SmtpUsername: "emailuser", SmtpPassword: "smtp-password", SmtpSecure: "tls", LdapEnabled: true, LdapHosts: []string{"test"}, LdapPort: 636, LdapEncryption: "ldaps", LdapVerification: true, LdapCertificate: "string", LdapTimeout: 10, LdapBaseDn: "ou=people,dc=example,dc=com", LdapBindDn: "cn=appwrite,ou=services,dc=example,dc=com", LdapBindPassword: "ldap-password", LdapUserFilter: "(uid={username})", LdapRequiredGroupDn: "cn=staff,ou=groups,dc=example,dc=com", LdapIdAttribute: "entryUUID", LdapEmailAttribute: "mail", LdapNameAttribute: "cn", LdapPhoneAttribute: "mobile", PingCount: 1, PingedAt: "2020-10-15T06:38:00.000+00:00", Labels: []string{"test"}, Status: "active", Onboarding: map[string]interface{}{}, AuthMethods: []ProjectAuthMethod{ProjectAuthMethod{Id: "email-password", Enabled: true}}, Services: []ProjectService{ProjectService{Id: "sites", Enabled: true}}, Protocols: []ProjectProtocol{ProjectProtocol{Id: "graphql", Enabled: true}}, Blocks: []Block{Block{CreatedAt: "2020-10-15T06:38:00.000+00:00", ResourceType: "project", ResourceId: "5e5ea5c16897e", Mode: "readOnly"}}, ConsoleAccessedAt: "2020-10-15T06:38:00.000+00:00"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
