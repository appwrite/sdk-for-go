package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBillingPlanModel(t *testing.T) {
	model := BillingPlan{Id: "tier-0", Name: "Hobby", Desc: "Hobby plan", Order: 0, Price: 25, Trial: 14, Bandwidth: 25, Storage: 25, ImageTransformations: 100, ScreenshotsGenerated: 50, Webhooks: 25, WafRules: 2, Environments: 5, Projects: 2, Platforms: 3, Users: 25, Teams: 25, Databases: 25, DatabasesReads: 500000, DatabasesWrites: 250000, DatabasesBatchSize: 100, Buckets: 25, FileSize: 25, Functions: 25, Sites: 1, Executions: 25, ExecutionsRetentionCount: 10000, GBHours: 100, Realtime: 25, RealtimeMessages: 100000, Messages: 1000, Topics: 1, AuthPhone: 10, Domains: 5, UsageLogs: 30, ProjectInactivityDays: 7, AlertLimit: 80, Addons: BillingPlanAddon{}, BudgetCapEnabled: true, CustomSmtp: true, EmailBranding: true, RequiresPaymentMethod: true, RequiresBillingAddress: true, IsAvailable: true, SelfService: true, PremiumSupport: true, Budgeting: true, SupportsMockNumbers: true, SupportsOrganizationRoles: true, SupportsCredits: true, SupportsDedicatedDatabases: true, SupportsDisposableEmailValidation: true, SupportsCanonicalEmailValidation: true, SupportsFreeEmailValidation: true, SupportsCorporateEmailValidation: true, SupportsProjectSpecificRoles: true, UsagePerProject: true, SupportedAddons: BillingPlanSupportedAddons{Baa: true, PremiumGeoDB: true, PremiumGeoDBOrg: true}, DeploymentSize: 30, BuildSize: 2000, FunctionsIntervalMinimum: 60, DatabasesAllowEncrypt: true, Group: "pro", DatabaseComputeCredit: 10}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result BillingPlan
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
	}
	if result.Desc != model.Desc {
		t.Errorf("Expected Desc %v, got %v", model.Desc, result.Desc)
	}
	if result.Order != model.Order {
		t.Errorf("Expected Order %v, got %v", model.Order, result.Order)
	}
	if result.Price != model.Price {
		t.Errorf("Expected Price %v, got %v", model.Price, result.Price)
	}
	if result.Trial != model.Trial {
		t.Errorf("Expected Trial %v, got %v", model.Trial, result.Trial)
	}
	if result.Bandwidth != model.Bandwidth {
		t.Errorf("Expected Bandwidth %v, got %v", model.Bandwidth, result.Bandwidth)
	}
	if result.Storage != model.Storage {
		t.Errorf("Expected Storage %v, got %v", model.Storage, result.Storage)
	}
	if result.ImageTransformations != model.ImageTransformations {
		t.Errorf("Expected ImageTransformations %v, got %v", model.ImageTransformations, result.ImageTransformations)
	}
	if result.ScreenshotsGenerated != model.ScreenshotsGenerated {
		t.Errorf("Expected ScreenshotsGenerated %v, got %v", model.ScreenshotsGenerated, result.ScreenshotsGenerated)
	}
	if result.Webhooks != model.Webhooks {
		t.Errorf("Expected Webhooks %v, got %v", model.Webhooks, result.Webhooks)
	}
	if result.WafRules != model.WafRules {
		t.Errorf("Expected WafRules %v, got %v", model.WafRules, result.WafRules)
	}
	if result.Environments != model.Environments {
		t.Errorf("Expected Environments %v, got %v", model.Environments, result.Environments)
	}
	if result.Projects != model.Projects {
		t.Errorf("Expected Projects %v, got %v", model.Projects, result.Projects)
	}
	if result.Platforms != model.Platforms {
		t.Errorf("Expected Platforms %v, got %v", model.Platforms, result.Platforms)
	}
	if result.Users != model.Users {
		t.Errorf("Expected Users %v, got %v", model.Users, result.Users)
	}
	if result.Teams != model.Teams {
		t.Errorf("Expected Teams %v, got %v", model.Teams, result.Teams)
	}
	if result.Databases != model.Databases {
		t.Errorf("Expected Databases %v, got %v", model.Databases, result.Databases)
	}
	if result.DatabasesReads != model.DatabasesReads {
		t.Errorf("Expected DatabasesReads %v, got %v", model.DatabasesReads, result.DatabasesReads)
	}
	if result.DatabasesWrites != model.DatabasesWrites {
		t.Errorf("Expected DatabasesWrites %v, got %v", model.DatabasesWrites, result.DatabasesWrites)
	}
	if result.DatabasesBatchSize != model.DatabasesBatchSize {
		t.Errorf("Expected DatabasesBatchSize %v, got %v", model.DatabasesBatchSize, result.DatabasesBatchSize)
	}
	if result.Buckets != model.Buckets {
		t.Errorf("Expected Buckets %v, got %v", model.Buckets, result.Buckets)
	}
	if result.FileSize != model.FileSize {
		t.Errorf("Expected FileSize %v, got %v", model.FileSize, result.FileSize)
	}
	if result.Functions != model.Functions {
		t.Errorf("Expected Functions %v, got %v", model.Functions, result.Functions)
	}
	if result.Sites != model.Sites {
		t.Errorf("Expected Sites %v, got %v", model.Sites, result.Sites)
	}
	if result.Executions != model.Executions {
		t.Errorf("Expected Executions %v, got %v", model.Executions, result.Executions)
	}
	if result.ExecutionsRetentionCount != model.ExecutionsRetentionCount {
		t.Errorf("Expected ExecutionsRetentionCount %v, got %v", model.ExecutionsRetentionCount, result.ExecutionsRetentionCount)
	}
	if result.GBHours != model.GBHours {
		t.Errorf("Expected GBHours %v, got %v", model.GBHours, result.GBHours)
	}
	if result.Realtime != model.Realtime {
		t.Errorf("Expected Realtime %v, got %v", model.Realtime, result.Realtime)
	}
	if result.RealtimeMessages != model.RealtimeMessages {
		t.Errorf("Expected RealtimeMessages %v, got %v", model.RealtimeMessages, result.RealtimeMessages)
	}
	if result.Messages != model.Messages {
		t.Errorf("Expected Messages %v, got %v", model.Messages, result.Messages)
	}
	if result.Topics != model.Topics {
		t.Errorf("Expected Topics %v, got %v", model.Topics, result.Topics)
	}
	if result.AuthPhone != model.AuthPhone {
		t.Errorf("Expected AuthPhone %v, got %v", model.AuthPhone, result.AuthPhone)
	}
	if result.Domains != model.Domains {
		t.Errorf("Expected Domains %v, got %v", model.Domains, result.Domains)
	}
	if result.UsageLogs != model.UsageLogs {
		t.Errorf("Expected UsageLogs %v, got %v", model.UsageLogs, result.UsageLogs)
	}
	if result.ProjectInactivityDays != model.ProjectInactivityDays {
		t.Errorf("Expected ProjectInactivityDays %v, got %v", model.ProjectInactivityDays, result.ProjectInactivityDays)
	}
	if result.AlertLimit != model.AlertLimit {
		t.Errorf("Expected AlertLimit %v, got %v", model.AlertLimit, result.AlertLimit)
	}
	if result.BudgetCapEnabled != model.BudgetCapEnabled {
		t.Errorf("Expected BudgetCapEnabled %v, got %v", model.BudgetCapEnabled, result.BudgetCapEnabled)
	}
	if result.CustomSmtp != model.CustomSmtp {
		t.Errorf("Expected CustomSmtp %v, got %v", model.CustomSmtp, result.CustomSmtp)
	}
	if result.EmailBranding != model.EmailBranding {
		t.Errorf("Expected EmailBranding %v, got %v", model.EmailBranding, result.EmailBranding)
	}
	if result.RequiresPaymentMethod != model.RequiresPaymentMethod {
		t.Errorf("Expected RequiresPaymentMethod %v, got %v", model.RequiresPaymentMethod, result.RequiresPaymentMethod)
	}
	if result.RequiresBillingAddress != model.RequiresBillingAddress {
		t.Errorf("Expected RequiresBillingAddress %v, got %v", model.RequiresBillingAddress, result.RequiresBillingAddress)
	}
	if result.IsAvailable != model.IsAvailable {
		t.Errorf("Expected IsAvailable %v, got %v", model.IsAvailable, result.IsAvailable)
	}
	if result.SelfService != model.SelfService {
		t.Errorf("Expected SelfService %v, got %v", model.SelfService, result.SelfService)
	}
	if result.PremiumSupport != model.PremiumSupport {
		t.Errorf("Expected PremiumSupport %v, got %v", model.PremiumSupport, result.PremiumSupport)
	}
	if result.Budgeting != model.Budgeting {
		t.Errorf("Expected Budgeting %v, got %v", model.Budgeting, result.Budgeting)
	}
	if result.SupportsMockNumbers != model.SupportsMockNumbers {
		t.Errorf("Expected SupportsMockNumbers %v, got %v", model.SupportsMockNumbers, result.SupportsMockNumbers)
	}
	if result.SupportsOrganizationRoles != model.SupportsOrganizationRoles {
		t.Errorf("Expected SupportsOrganizationRoles %v, got %v", model.SupportsOrganizationRoles, result.SupportsOrganizationRoles)
	}
	if result.SupportsCredits != model.SupportsCredits {
		t.Errorf("Expected SupportsCredits %v, got %v", model.SupportsCredits, result.SupportsCredits)
	}
	if result.SupportsDedicatedDatabases != model.SupportsDedicatedDatabases {
		t.Errorf("Expected SupportsDedicatedDatabases %v, got %v", model.SupportsDedicatedDatabases, result.SupportsDedicatedDatabases)
	}
	if result.SupportsDisposableEmailValidation != model.SupportsDisposableEmailValidation {
		t.Errorf("Expected SupportsDisposableEmailValidation %v, got %v", model.SupportsDisposableEmailValidation, result.SupportsDisposableEmailValidation)
	}
	if result.SupportsCanonicalEmailValidation != model.SupportsCanonicalEmailValidation {
		t.Errorf("Expected SupportsCanonicalEmailValidation %v, got %v", model.SupportsCanonicalEmailValidation, result.SupportsCanonicalEmailValidation)
	}
	if result.SupportsFreeEmailValidation != model.SupportsFreeEmailValidation {
		t.Errorf("Expected SupportsFreeEmailValidation %v, got %v", model.SupportsFreeEmailValidation, result.SupportsFreeEmailValidation)
	}
	if result.SupportsCorporateEmailValidation != model.SupportsCorporateEmailValidation {
		t.Errorf("Expected SupportsCorporateEmailValidation %v, got %v", model.SupportsCorporateEmailValidation, result.SupportsCorporateEmailValidation)
	}
	if result.SupportsProjectSpecificRoles != model.SupportsProjectSpecificRoles {
		t.Errorf("Expected SupportsProjectSpecificRoles %v, got %v", model.SupportsProjectSpecificRoles, result.SupportsProjectSpecificRoles)
	}
	if result.UsagePerProject != model.UsagePerProject {
		t.Errorf("Expected UsagePerProject %v, got %v", model.UsagePerProject, result.UsagePerProject)
	}
	if result.DeploymentSize != model.DeploymentSize {
		t.Errorf("Expected DeploymentSize %v, got %v", model.DeploymentSize, result.DeploymentSize)
	}
	if result.BuildSize != model.BuildSize {
		t.Errorf("Expected BuildSize %v, got %v", model.BuildSize, result.BuildSize)
	}
	if result.FunctionsIntervalMinimum != model.FunctionsIntervalMinimum {
		t.Errorf("Expected FunctionsIntervalMinimum %v, got %v", model.FunctionsIntervalMinimum, result.FunctionsIntervalMinimum)
	}
	if result.DatabasesAllowEncrypt != model.DatabasesAllowEncrypt {
		t.Errorf("Expected DatabasesAllowEncrypt %v, got %v", model.DatabasesAllowEncrypt, result.DatabasesAllowEncrypt)
	}
	if result.Group != model.Group {
		t.Errorf("Expected Group %v, got %v", model.Group, result.Group)
	}
	if result.DatabaseComputeCredit != model.DatabaseComputeCredit {
		t.Errorf("Expected DatabaseComputeCredit %v, got %v", model.DatabaseComputeCredit, result.DatabaseComputeCredit)
	}
}

func TestBillingPlanModelDefaults(t *testing.T) {
	var fresh BillingPlan
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
		encoded, err := json.Marshal(fresh.Desc)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing desc to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Order)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing order to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Price)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing price to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Trial)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing trial to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Bandwidth)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing bandwidth to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Storage)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing storage to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ImageTransformations)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing imageTransformations to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ScreenshotsGenerated)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing screenshotsGenerated to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Webhooks)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing webhooks to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.WafRules)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing wafRules to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Environments)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing environments to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Projects)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing projects to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Platforms)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing platforms to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Users)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing users to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Teams)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing teams to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Databases)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing databases to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DatabasesReads)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing databasesReads to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DatabasesWrites)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing databasesWrites to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DatabasesBatchSize)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing databasesBatchSize to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Buckets)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing buckets to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.FileSize)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing fileSize to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Functions)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing functions to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Sites)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing sites to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Executions)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing executions to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("1000"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ExecutionsRetentionCount)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing executionsRetentionCount to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.GBHours)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing GBHours to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Realtime)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing realtime to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.RealtimeMessages)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing realtimeMessages to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Messages)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing messages to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Topics)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing topics to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.AuthPhone)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing authPhone to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Domains)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing domains to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.UsageLogs)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing usageLogs to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ProjectInactivityDays)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing projectInactivityDays to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("80"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.AlertLimit)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing alertLimit to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.BudgetCapEnabled)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing budgetCapEnabled to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.CustomSmtp)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing customSmtp to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.EmailBranding)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing emailBranding to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.RequiresPaymentMethod)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing requiresPaymentMethod to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.RequiresBillingAddress)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing requiresBillingAddress to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.IsAvailable)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing isAvailable to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SelfService)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing selfService to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.PremiumSupport)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing premiumSupport to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Budgeting)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing budgeting to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SupportsMockNumbers)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing supportsMockNumbers to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SupportsOrganizationRoles)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing supportsOrganizationRoles to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SupportsCredits)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing supportsCredits to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SupportsDedicatedDatabases)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing supportsDedicatedDatabases to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SupportsDisposableEmailValidation)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing supportsDisposableEmailValidation to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SupportsCanonicalEmailValidation)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing supportsCanonicalEmailValidation to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SupportsFreeEmailValidation)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing supportsFreeEmailValidation to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SupportsCorporateEmailValidation)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing supportsCorporateEmailValidation to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SupportsProjectSpecificRoles)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing supportsProjectSpecificRoles to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("true"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.UsagePerProject)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing usagePerProject to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DeploymentSize)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing deploymentSize to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.BuildSize)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing buildSize to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.FunctionsIntervalMinimum)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing functionsIntervalMinimum to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DatabasesAllowEncrypt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing databasesAllowEncrypt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DatabaseComputeCredit)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing databaseComputeCredit to default to %v, got %v", want, got)
		}
	}

	reused := BillingPlan{Id: "tier-0", Name: "Hobby", Desc: "Hobby plan", Order: 0, Price: 25, Trial: 14, Bandwidth: 25, Storage: 25, ImageTransformations: 100, ScreenshotsGenerated: 50, Webhooks: 25, WafRules: 2, Environments: 5, Projects: 2, Platforms: 3, Users: 25, Teams: 25, Databases: 25, DatabasesReads: 500000, DatabasesWrites: 250000, DatabasesBatchSize: 100, Buckets: 25, FileSize: 25, Functions: 25, Sites: 1, Executions: 25, ExecutionsRetentionCount: 10000, GBHours: 100, Realtime: 25, RealtimeMessages: 100000, Messages: 1000, Topics: 1, AuthPhone: 10, Domains: 5, UsageLogs: 30, ProjectInactivityDays: 7, AlertLimit: 80, Addons: BillingPlanAddon{}, BudgetCapEnabled: true, CustomSmtp: true, EmailBranding: true, RequiresPaymentMethod: true, RequiresBillingAddress: true, IsAvailable: true, SelfService: true, PremiumSupport: true, Budgeting: true, SupportsMockNumbers: true, SupportsOrganizationRoles: true, SupportsCredits: true, SupportsDedicatedDatabases: true, SupportsDisposableEmailValidation: true, SupportsCanonicalEmailValidation: true, SupportsFreeEmailValidation: true, SupportsCorporateEmailValidation: true, SupportsProjectSpecificRoles: true, UsagePerProject: true, SupportedAddons: BillingPlanSupportedAddons{Baa: true, PremiumGeoDB: true, PremiumGeoDBOrg: true}, DeploymentSize: 30, BuildSize: 2000, FunctionsIntervalMinimum: 60, DatabasesAllowEncrypt: true, Group: "pro", DatabaseComputeCredit: 10}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
