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
	Members int `json:"members"`
	// Webhooks
	Webhooks int `json:"webhooks"`
	// Maximum WAF rules per project
	WafRules int `json:"wafRules"`
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
	ActivityLogs int `json:"activityLogs"`
	// Usage history days
	UsageLogs int `json:"usageLogs"`
	// Usage log time intervals allowed for this plan (e.g. 15m, 1h, 1d).
	UsageLogsIntervals []string `json:"usageLogsIntervals"`
	// Number of days of console inactivity before a project is paused. 0 means
	// pausing is disabled.
	ProjectInactivityDays int `json:"projectInactivityDays"`
	// Alert threshold percentage
	AlertLimit int `json:"alertLimit"`
	// Additional resources
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
	BackupsEnabled bool `json:"backupsEnabled"`
	// Whether usage addons are calculated per project.
	UsagePerProject bool `json:"usagePerProject"`
	// Supported addons for this plan
	SupportedAddons BillingPlanSupportedAddons `json:"supportedAddons"`
	// How many policies does plan support
	BackupPolicies int `json:"backupPolicies"`
	// Maximum function and site deployment size in MB
	DeploymentSize int `json:"deploymentSize"`
	// Maximum function and site deployment size in MB
	BuildSize int `json:"buildSize"`
	// Does the plan support encrypted string attributes or not.
	DatabasesAllowEncrypt bool `json:"databasesAllowEncrypt"`
	// Plan specific limits
	Limits BillingPlanLimits `json:"limits"`
	// Group of this billing plan for variants
	Group string `json:"group"`
	// Details of the program this plan is a part of.
	Program Program `json:"program"`
	// Dedicated database limits available to this plan.
	DedicatedDatabases BillingPlanDedicatedDatabaseLimits `json:"dedicatedDatabases"`

	// Used by Decode() method
	data []byte
}

func (model BillingPlan) New(data []byte) *BillingPlan {
	model.data = data
	return &model
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
