package models

import (
	"encoding/json"
	"testing"
)

func TestOrganizationModel(t *testing.T) {
	model := Organization{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "VIP", Total: 7, Prefs: Preferences{}, BudgetAlerts: []int{1}, BillingPlan: "tier-1", BillingPlanId: "tier-1", BillingPlanDetails: BillingPlan{Id: "tier-0", Name: "Hobby", Desc: "Hobby plan", Order: 0, Price: 25, Trial: 14, Bandwidth: 25, Storage: 25, ImageTransformations: 100, ScreenshotsGenerated: 50, Webhooks: 25, WafRules: 2, Projects: 2, Platforms: 3, Users: 25, Teams: 25, Databases: 25, DatabasesReads: 500000, DatabasesWrites: 250000, DatabasesBatchSize: 100, Buckets: 25, FileSize: 25, Functions: 25, Sites: 1, Executions: 25, ExecutionsRetentionCount: 10000, GBHours: 100, Realtime: 25, RealtimeMessages: 100000, Messages: 1000, Topics: 1, AuthPhone: 10, Domains: 5, UsageLogs: 30, ProjectInactivityDays: 7, AlertLimit: 80, Usage: UsageBillingPlan{Bandwidth: AdditionalResource{Name: "string", Unit: "GB", Currency: "USD", Price: 5, Value: 25, InvoiceDesc: "string"}, Executions: AdditionalResource{Name: "string", Unit: "GB", Currency: "USD", Price: 5, Value: 25, InvoiceDesc: "string"}, Realtime: AdditionalResource{Name: "string", Unit: "GB", Currency: "USD", Price: 5, Value: 25, InvoiceDesc: "string"}, RealtimeMessages: AdditionalResource{Name: "string", Unit: "GB", Currency: "USD", Price: 5, Value: 25, InvoiceDesc: "string"}, Storage: AdditionalResource{Name: "string", Unit: "GB", Currency: "USD", Price: 5, Value: 25, InvoiceDesc: "string"}, Users: AdditionalResource{Name: "string", Unit: "GB", Currency: "USD", Price: 5, Value: 25, InvoiceDesc: "string"}, GBHours: AdditionalResource{Name: "string", Unit: "GB", Currency: "USD", Price: 5, Value: 25, InvoiceDesc: "string"}, ImageTransformations: AdditionalResource{Name: "string", Unit: "GB", Currency: "USD", Price: 5, Value: 25, InvoiceDesc: "string"}}, Addons: BillingPlanAddon{}, BudgetCapEnabled: true, CustomSmtp: true, EmailBranding: true, RequiresPaymentMethod: true, RequiresBillingAddress: true, IsAvailable: true, SelfService: true, PremiumSupport: true, Budgeting: true, SupportsMockNumbers: true, SupportsOrganizationRoles: true, SupportsCredits: true, SupportsDisposableEmailValidation: true, SupportsCanonicalEmailValidation: true, SupportsFreeEmailValidation: true, SupportsCorporateEmailValidation: true, SupportsProjectSpecificRoles: true, UsagePerProject: true, SupportedAddons: BillingPlanSupportedAddons{Baa: true, PremiumGeoDB: true, PremiumGeoDBOrg: true}, DeploymentSize: 30, BuildSize: 2000, DatabasesAllowEncrypt: true, Group: "pro"}, BillingEmail: "billing@org.example", BillingStartDate: "2020-10-15T06:38:00.000+00:00", BillingCurrentInvoiceDate: "2020-10-15T06:38:00.000+00:00", BillingNextInvoiceDate: "2020-10-15T06:38:00.000+00:00", BillingTrialDays: 14, BillingAggregationId: "adbc3de4rddfsd", BillingInvoiceId: "adbc3de4rddfsd", PaymentMethodId: "adbc3de4rddfsd", Status: "active", MarkedForDeletion: true, Platform: "imagine", Projects: []string{"test"}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Organization
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
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
	if result.BillingPlan != model.BillingPlan {
		t.Errorf("Expected BillingPlan %v, got %v", model.BillingPlan, result.BillingPlan)
	}
	if result.BillingPlanId != model.BillingPlanId {
		t.Errorf("Expected BillingPlanId %v, got %v", model.BillingPlanId, result.BillingPlanId)
	}
	if result.BillingEmail != model.BillingEmail {
		t.Errorf("Expected BillingEmail %v, got %v", model.BillingEmail, result.BillingEmail)
	}
	if result.BillingStartDate != model.BillingStartDate {
		t.Errorf("Expected BillingStartDate %v, got %v", model.BillingStartDate, result.BillingStartDate)
	}
	if result.BillingCurrentInvoiceDate != model.BillingCurrentInvoiceDate {
		t.Errorf("Expected BillingCurrentInvoiceDate %v, got %v", model.BillingCurrentInvoiceDate, result.BillingCurrentInvoiceDate)
	}
	if result.BillingNextInvoiceDate != model.BillingNextInvoiceDate {
		t.Errorf("Expected BillingNextInvoiceDate %v, got %v", model.BillingNextInvoiceDate, result.BillingNextInvoiceDate)
	}
	if result.BillingTrialDays != model.BillingTrialDays {
		t.Errorf("Expected BillingTrialDays %v, got %v", model.BillingTrialDays, result.BillingTrialDays)
	}
	if result.BillingAggregationId != model.BillingAggregationId {
		t.Errorf("Expected BillingAggregationId %v, got %v", model.BillingAggregationId, result.BillingAggregationId)
	}
	if result.BillingInvoiceId != model.BillingInvoiceId {
		t.Errorf("Expected BillingInvoiceId %v, got %v", model.BillingInvoiceId, result.BillingInvoiceId)
	}
	if result.PaymentMethodId != model.PaymentMethodId {
		t.Errorf("Expected PaymentMethodId %v, got %v", model.PaymentMethodId, result.PaymentMethodId)
	}
	if result.Status != model.Status {
		t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
	}
	if result.MarkedForDeletion != model.MarkedForDeletion {
		t.Errorf("Expected MarkedForDeletion %v, got %v", model.MarkedForDeletion, result.MarkedForDeletion)
	}
	if result.Platform != model.Platform {
		t.Errorf("Expected Platform %v, got %v", model.Platform, result.Platform)
	}
}
