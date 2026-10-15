package models

import (
	"encoding/json"
	"errors"
)

// BillingPlan Model
type BillingPlan struct {
	// Plan ID.
	Id string `json:"$id"`
	// Plan name
	Name string `json:"name"`
	// Plan description
	Desc string `json:"desc"`
	// Plan order
	Order int `json:"order"`
	// Price
	Price float64 `json:"price"`
	// Trial days
	Trial int `json:"trial"`
	// Bandwidth
	Bandwidth int `json:"bandwidth"`
	// Storage
	Storage int `json:"storage"`
	// Image Transformations
	ImageTransformations int `json:"imageTransformations"`
	// Screenshots generated
	ScreenshotsGenerated int `json:"screenshotsGenerated"`
	// Members
	Members *int `json:"members"`
	// Webhooks
	Webhooks int `json:"webhooks"`
	// Maximum WAF rules per project
	WafRules int `json:"wafRules"`
	// Maximum environments per project. 0 when the plan does not offer
	// environments.
	Environments int `json:"environments"`
	// Projects
	Projects int `json:"projects"`
	// Platforms
	Platforms int `json:"platforms"`
	// Users
	Users int `json:"users"`
	// Teams
	Teams int `json:"teams"`
	// Databases
	Databases int `json:"databases"`
	// Database reads per month
	DatabasesReads int `json:"databasesReads"`
	// Database writes per month
	DatabasesWrites int `json:"databasesWrites"`
	// Database batch size limit
	DatabasesBatchSize int `json:"databasesBatchSize"`
	// Buckets
	Buckets int `json:"buckets"`
	// File size
	FileSize int `json:"fileSize"`
	// Functions
	Functions int `json:"functions"`
	// Sites
	Sites int `json:"sites"`
	// Function executions
	Executions int `json:"executions"`
	// Rolling max executions retained per function/site
	ExecutionsRetentionCount int `json:"executionsRetentionCount"`
	// GB hours for functions
	GBHours int `json:"GBHours"`
	// Realtime connections
	Realtime int `json:"realtime"`
	// Realtime messages
	RealtimeMessages int `json:"realtimeMessages"`
	// Messages per month
	Messages int `json:"messages"`
	// Topics for messaging
	Topics int `json:"topics"`
	// SMS authentications per month
	AuthPhone int `json:"authPhone"`
	// Custom domains
	Domains int `json:"domains"`
	// Activity log days
	ActivityLogs *int `json:"activityLogs"`
	// Usage history days
	UsageLogs int `json:"usageLogs"`
	// Usage log time intervals allowed for this plan (e.g. 15m, 1h, 1d).
	UsageLogsIntervals []string `json:"usageLogsIntervals"`
	// Metrics this plan only records as a total. They cannot be broken down by
	// dimension or filtered, because the stored events cover a fraction of the
	// real traffic.
	UsageAggregateOnlyMetrics []string `json:"usageAggregateOnlyMetrics"`
	// Number of days of console inactivity before a project is paused. 0 means
	// pausing is disabled.
	ProjectInactivityDays int `json:"projectInactivityDays"`
	// Alert threshold percentage
	AlertLimit int `json:"alertLimit"`
	// Additional resources. Null when the plan does not support usage billing.
	Usage UsageBillingPlan `json:"usage"`
	// Addons
	Addons BillingPlanAddon `json:"addons"`
	// Budget cap enabled or disabled.
	BudgetCapEnabled bool `json:"budgetCapEnabled"`
	// Custom SMTP
	CustomSmtp bool `json:"customSmtp"`
	// Appwrite branding in email
	EmailBranding bool `json:"emailBranding"`
	// Does plan require payment method
	RequiresPaymentMethod bool `json:"requiresPaymentMethod"`
	// Does plan require billing address
	RequiresBillingAddress bool `json:"requiresBillingAddress"`
	// ISO country codes eligible for this regional plan. Empty means no
	// restriction.
	EligibleCountries []string `json:"eligibleCountries"`
	// Is the billing plan available
	IsAvailable bool `json:"isAvailable"`
	// Can user change the plan themselves
	SelfService bool `json:"selfService"`
	// Does plan enable premium support
	PremiumSupport bool `json:"premiumSupport"`
	// Does plan support budget cap
	Budgeting bool `json:"budgeting"`
	// Does plan support mock numbers
	SupportsMockNumbers bool `json:"supportsMockNumbers"`
	// Does plan support organization roles
	SupportsOrganizationRoles bool `json:"supportsOrganizationRoles"`
	// Does plan support credit
	SupportsCredits bool `json:"supportsCredits"`
	// Does plan support dedicated databases.
	SupportsDedicatedDatabases bool `json:"supportsDedicatedDatabases"`
	// Does plan support blocking disposable email addresses.
	SupportsDisposableEmailValidation bool `json:"supportsDisposableEmailValidation"`
	// Does plan support requiring canonical email addresses.
	SupportsCanonicalEmailValidation bool `json:"supportsCanonicalEmailValidation"`
	// Does plan support blocking free email addresses.
	SupportsFreeEmailValidation bool `json:"supportsFreeEmailValidation"`
	// Does plan support restricting sign-ups to corporate email addresses only.
	SupportsCorporateEmailValidation bool `json:"supportsCorporateEmailValidation"`
	// Does plan support project-specific member roles.
	SupportsProjectSpecificRoles bool `json:"supportsProjectSpecificRoles"`
	// Does plan support backup policies.
	BackupsEnabled *bool `json:"backupsEnabled"`
	// Whether usage addons are calculated per project.
	UsagePerProject bool `json:"usagePerProject"`
	// Supported addons for this plan
	SupportedAddons BillingPlanSupportedAddons `json:"supportedAddons"`
	// How many policies does plan support
	BackupPolicies *int `json:"backupPolicies"`
	// Maximum function and site deployment size in MB
	DeploymentSize int `json:"deploymentSize"`
	// Maximum function and site deployment size in MB
	BuildSize int `json:"buildSize"`
	// Shortest function schedule interval allowed, in minutes. 0 allows every
	// interval.
	FunctionsIntervalMinimum int `json:"functionsIntervalMinimum"`
	// Does the plan support encrypted string attributes or not.
	DatabasesAllowEncrypt bool `json:"databasesAllowEncrypt"`
	// Plan specific limits
	Limits BillingPlanLimits `json:"limits"`
	// Group of this billing plan for variants
	Group string `json:"group"`
	// Details of the program this plan is a part of.
	Program Program `json:"program"`
	// Included monthly dedicated-database compute credit in USD. Resets each
	// billing cycle with no roll-over.
	DatabaseComputeCredit float64 `json:"databaseComputeCredit"`
	// Dedicated database limits available to this plan.
	DedicatedDatabases BillingPlanDedicatedDatabaseLimits `json:"dedicatedDatabases"`

	// Used by Decode() method
	data []byte
}

func (model BillingPlan) New(data []byte) *BillingPlan {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *BillingPlan) UnmarshalJSON(data []byte) error {
	type alias BillingPlan
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                                json.RawMessage `json:"$id"`
		Name                              json.RawMessage `json:"name"`
		Desc                              json.RawMessage `json:"desc"`
		Order                             json.RawMessage `json:"order"`
		Price                             json.RawMessage `json:"price"`
		Trial                             json.RawMessage `json:"trial"`
		Bandwidth                         json.RawMessage `json:"bandwidth"`
		Storage                           json.RawMessage `json:"storage"`
		ImageTransformations              json.RawMessage `json:"imageTransformations"`
		ScreenshotsGenerated              json.RawMessage `json:"screenshotsGenerated"`
		Webhooks                          json.RawMessage `json:"webhooks"`
		WafRules                          json.RawMessage `json:"wafRules"`
		Environments                      json.RawMessage `json:"environments"`
		Projects                          json.RawMessage `json:"projects"`
		Platforms                         json.RawMessage `json:"platforms"`
		Users                             json.RawMessage `json:"users"`
		Teams                             json.RawMessage `json:"teams"`
		Databases                         json.RawMessage `json:"databases"`
		DatabasesReads                    json.RawMessage `json:"databasesReads"`
		DatabasesWrites                   json.RawMessage `json:"databasesWrites"`
		DatabasesBatchSize                json.RawMessage `json:"databasesBatchSize"`
		Buckets                           json.RawMessage `json:"buckets"`
		FileSize                          json.RawMessage `json:"fileSize"`
		Functions                         json.RawMessage `json:"functions"`
		Sites                             json.RawMessage `json:"sites"`
		Executions                        json.RawMessage `json:"executions"`
		ExecutionsRetentionCount          json.RawMessage `json:"executionsRetentionCount"`
		GBHours                           json.RawMessage `json:"GBHours"`
		Realtime                          json.RawMessage `json:"realtime"`
		RealtimeMessages                  json.RawMessage `json:"realtimeMessages"`
		Messages                          json.RawMessage `json:"messages"`
		Topics                            json.RawMessage `json:"topics"`
		AuthPhone                         json.RawMessage `json:"authPhone"`
		Domains                           json.RawMessage `json:"domains"`
		UsageLogs                         json.RawMessage `json:"usageLogs"`
		ProjectInactivityDays             json.RawMessage `json:"projectInactivityDays"`
		AlertLimit                        json.RawMessage `json:"alertLimit"`
		BudgetCapEnabled                  json.RawMessage `json:"budgetCapEnabled"`
		CustomSmtp                        json.RawMessage `json:"customSmtp"`
		EmailBranding                     json.RawMessage `json:"emailBranding"`
		RequiresPaymentMethod             json.RawMessage `json:"requiresPaymentMethod"`
		RequiresBillingAddress            json.RawMessage `json:"requiresBillingAddress"`
		IsAvailable                       json.RawMessage `json:"isAvailable"`
		SelfService                       json.RawMessage `json:"selfService"`
		PremiumSupport                    json.RawMessage `json:"premiumSupport"`
		Budgeting                         json.RawMessage `json:"budgeting"`
		SupportsMockNumbers               json.RawMessage `json:"supportsMockNumbers"`
		SupportsOrganizationRoles         json.RawMessage `json:"supportsOrganizationRoles"`
		SupportsCredits                   json.RawMessage `json:"supportsCredits"`
		SupportsDedicatedDatabases        json.RawMessage `json:"supportsDedicatedDatabases"`
		SupportsDisposableEmailValidation json.RawMessage `json:"supportsDisposableEmailValidation"`
		SupportsCanonicalEmailValidation  json.RawMessage `json:"supportsCanonicalEmailValidation"`
		SupportsFreeEmailValidation       json.RawMessage `json:"supportsFreeEmailValidation"`
		SupportsCorporateEmailValidation  json.RawMessage `json:"supportsCorporateEmailValidation"`
		SupportsProjectSpecificRoles      json.RawMessage `json:"supportsProjectSpecificRoles"`
		UsagePerProject                   json.RawMessage `json:"usagePerProject"`
		DeploymentSize                    json.RawMessage `json:"deploymentSize"`
		BuildSize                         json.RawMessage `json:"buildSize"`
		FunctionsIntervalMinimum          json.RawMessage `json:"functionsIntervalMinimum"`
		DatabasesAllowEncrypt             json.RawMessage `json:"databasesAllowEncrypt"`
		DatabaseComputeCredit             json.RawMessage `json:"databaseComputeCredit"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Desc == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Desc); err != nil {
			return err
		}
	}
	if present.Order == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Order); err != nil {
			return err
		}
	}
	if present.Price == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Price); err != nil {
			return err
		}
	}
	if present.Trial == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Trial); err != nil {
			return err
		}
	}
	if present.Bandwidth == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Bandwidth); err != nil {
			return err
		}
	}
	if present.Storage == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Storage); err != nil {
			return err
		}
	}
	if present.ImageTransformations == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ImageTransformations); err != nil {
			return err
		}
	}
	if present.ScreenshotsGenerated == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ScreenshotsGenerated); err != nil {
			return err
		}
	}
	if present.Webhooks == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Webhooks); err != nil {
			return err
		}
	}
	if present.WafRules == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.WafRules); err != nil {
			return err
		}
	}
	if present.Environments == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Environments); err != nil {
			return err
		}
	}
	if present.Projects == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Projects); err != nil {
			return err
		}
	}
	if present.Platforms == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Platforms); err != nil {
			return err
		}
	}
	if present.Users == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Users); err != nil {
			return err
		}
	}
	if present.Teams == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Teams); err != nil {
			return err
		}
	}
	if present.Databases == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Databases); err != nil {
			return err
		}
	}
	if present.DatabasesReads == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.DatabasesReads); err != nil {
			return err
		}
	}
	if present.DatabasesWrites == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.DatabasesWrites); err != nil {
			return err
		}
	}
	if present.DatabasesBatchSize == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.DatabasesBatchSize); err != nil {
			return err
		}
	}
	if present.Buckets == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Buckets); err != nil {
			return err
		}
	}
	if present.FileSize == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.FileSize); err != nil {
			return err
		}
	}
	if present.Functions == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Functions); err != nil {
			return err
		}
	}
	if present.Sites == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Sites); err != nil {
			return err
		}
	}
	if present.Executions == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Executions); err != nil {
			return err
		}
	}
	if present.ExecutionsRetentionCount == nil {
		if err := json.Unmarshal([]byte("1000"), &decoded.ExecutionsRetentionCount); err != nil {
			return err
		}
	}
	if present.GBHours == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.GBHours); err != nil {
			return err
		}
	}
	if present.Realtime == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Realtime); err != nil {
			return err
		}
	}
	if present.RealtimeMessages == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.RealtimeMessages); err != nil {
			return err
		}
	}
	if present.Messages == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Messages); err != nil {
			return err
		}
	}
	if present.Topics == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Topics); err != nil {
			return err
		}
	}
	if present.AuthPhone == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.AuthPhone); err != nil {
			return err
		}
	}
	if present.Domains == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Domains); err != nil {
			return err
		}
	}
	if present.UsageLogs == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.UsageLogs); err != nil {
			return err
		}
	}
	if present.ProjectInactivityDays == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ProjectInactivityDays); err != nil {
			return err
		}
	}
	if present.AlertLimit == nil {
		if err := json.Unmarshal([]byte("80"), &decoded.AlertLimit); err != nil {
			return err
		}
	}
	if present.BudgetCapEnabled == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.BudgetCapEnabled); err != nil {
			return err
		}
	}
	if present.CustomSmtp == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.CustomSmtp); err != nil {
			return err
		}
	}
	if present.EmailBranding == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.EmailBranding); err != nil {
			return err
		}
	}
	if present.RequiresPaymentMethod == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.RequiresPaymentMethod); err != nil {
			return err
		}
	}
	if present.RequiresBillingAddress == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.RequiresBillingAddress); err != nil {
			return err
		}
	}
	if present.IsAvailable == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.IsAvailable); err != nil {
			return err
		}
	}
	if present.SelfService == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.SelfService); err != nil {
			return err
		}
	}
	if present.PremiumSupport == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.PremiumSupport); err != nil {
			return err
		}
	}
	if present.Budgeting == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Budgeting); err != nil {
			return err
		}
	}
	if present.SupportsMockNumbers == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.SupportsMockNumbers); err != nil {
			return err
		}
	}
	if present.SupportsOrganizationRoles == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.SupportsOrganizationRoles); err != nil {
			return err
		}
	}
	if present.SupportsCredits == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.SupportsCredits); err != nil {
			return err
		}
	}
	if present.SupportsDedicatedDatabases == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.SupportsDedicatedDatabases); err != nil {
			return err
		}
	}
	if present.SupportsDisposableEmailValidation == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.SupportsDisposableEmailValidation); err != nil {
			return err
		}
	}
	if present.SupportsCanonicalEmailValidation == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.SupportsCanonicalEmailValidation); err != nil {
			return err
		}
	}
	if present.SupportsFreeEmailValidation == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.SupportsFreeEmailValidation); err != nil {
			return err
		}
	}
	if present.SupportsCorporateEmailValidation == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.SupportsCorporateEmailValidation); err != nil {
			return err
		}
	}
	if present.SupportsProjectSpecificRoles == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.SupportsProjectSpecificRoles); err != nil {
			return err
		}
	}
	if present.UsagePerProject == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.UsagePerProject); err != nil {
			return err
		}
	}
	if present.DeploymentSize == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.DeploymentSize); err != nil {
			return err
		}
	}
	if present.BuildSize == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.BuildSize); err != nil {
			return err
		}
	}
	if present.FunctionsIntervalMinimum == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.FunctionsIntervalMinimum); err != nil {
			return err
		}
	}
	if present.DatabasesAllowEncrypt == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.DatabasesAllowEncrypt); err != nil {
			return err
		}
	}
	if present.DatabaseComputeCredit == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.DatabaseComputeCredit); err != nil {
			return err
		}
	}

	*model = BillingPlan(decoded)
	return nil
}

func (model *BillingPlan) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
