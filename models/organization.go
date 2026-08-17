package models

import (
	"encoding/json"
	"errors"
)

// Organization Model
type Organization struct {
	// Team ID.
	Id string `json:"$id"`
	// Team creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Team update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Team name.
	Name string `json:"name"`
	// Total number of team members.
	Total int `json:"total"`
	// Team preferences as a key-value object
	Prefs Preferences `json:"prefs"`
	// Project budget limit. Null when no budget is set.
	BillingBudget int `json:"billingBudget"`
	// Project budget limit
	BudgetAlerts []int `json:"budgetAlerts"`
	// Organization's billing plan ID.
	BillingPlan string `json:"billingPlan"`
	// Organization's billing plan ID.
	BillingPlanId string `json:"billingPlanId"`
	// Organization's billing plan.
	BillingPlanDetails BillingPlan `json:"billingPlanDetails"`
	// Billing email set for the organization.
	BillingEmail string `json:"billingEmail"`
	// Billing cycle start date.
	BillingStartDate string `json:"billingStartDate"`
	// Current invoice cycle start date.
	BillingCurrentInvoiceDate string `json:"billingCurrentInvoiceDate"`
	// Next invoice cycle start date.
	BillingNextInvoiceDate string `json:"billingNextInvoiceDate"`
	// Start date of trial.
	BillingTrialStartDate string `json:"billingTrialStartDate"`
	// Number of trial days.
	BillingTrialDays int `json:"billingTrialDays"`
	// Current active aggregation id.
	BillingAggregationId string `json:"billingAggregationId"`
	// Current active aggregation id.
	BillingInvoiceId string `json:"billingInvoiceId"`
	// Default payment method.
	PaymentMethodId string `json:"paymentMethodId"`
	// Default payment method.
	BillingAddressId string `json:"billingAddressId"`
	// Backup payment method.
	BackupPaymentMethodId string `json:"backupPaymentMethodId"`
	// Team status.
	Status string `json:"status"`
	// Remarks on team status.
	Remarks string `json:"remarks"`
	// Organization agreements
	AgreementBAA string `json:"agreementBAA"`
	// Program manager's name.
	ProgramManagerName string `json:"programManagerName"`
	// Program manager's calendar link.
	ProgramManagerCalendar string `json:"programManagerCalendar"`
	// Program's discord channel name.
	ProgramDiscordChannelName string `json:"programDiscordChannelName"`
	// Program's discord channel URL.
	ProgramDiscordChannelUrl string `json:"programDiscordChannelUrl"`
	// Billing limits reached
	BillingLimits BillingLimits `json:"billingLimits"`
	// Billing plan selected for downgrade.
	BillingPlanDowngrade string `json:"billingPlanDowngrade"`
	// Tax Id
	BillingTaxId string `json:"billingTaxId"`
	// Marked for deletion
	MarkedForDeletion bool `json:"markedForDeletion"`
	// Product with which the organization is associated (appwrite or imagine)
	Platform string `json:"platform"`
	// Selected projects
	Projects []string `json:"projects"`

	// Used by Decode() method
	data []byte
}

func (model Organization) New(data []byte) *Organization {
	model.data = data
	return &model
}

func (model *Organization) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
