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
	BillingBudget *int `json:"billingBudget"`
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
	BillingStartDate *string `json:"billingStartDate"`
	// Current invoice cycle start date.
	BillingCurrentInvoiceDate *string `json:"billingCurrentInvoiceDate"`
	// Next invoice cycle start date.
	BillingNextInvoiceDate *string `json:"billingNextInvoiceDate"`
	// Start date of trial.
	BillingTrialStartDate *string `json:"billingTrialStartDate"`
	// Number of trial days.
	BillingTrialDays *int `json:"billingTrialDays"`
	// Current active aggregation id.
	BillingAggregationId *string `json:"billingAggregationId"`
	// Current active aggregation id.
	BillingInvoiceId *string `json:"billingInvoiceId"`
	// Default payment method.
	PaymentMethodId *string `json:"paymentMethodId"`
	// Default payment method.
	BillingAddressId *string `json:"billingAddressId"`
	// Backup payment method.
	BackupPaymentMethodId *string `json:"backupPaymentMethodId"`
	// Team status.
	Status string `json:"status"`
	// Remarks on team status.
	Remarks *string `json:"remarks"`
	// Organization agreements
	AgreementBAA *string `json:"agreementBAA"`
	// Program manager's name.
	ProgramManagerName *string `json:"programManagerName"`
	// Program manager's calendar link.
	ProgramManagerCalendar *string `json:"programManagerCalendar"`
	// Program's discord channel name.
	ProgramDiscordChannelName *string `json:"programDiscordChannelName"`
	// Program's discord channel URL.
	ProgramDiscordChannelUrl *string `json:"programDiscordChannelUrl"`
	// Billing limits reached
	BillingLimits BillingLimits `json:"billingLimits"`
	// Billing plan selected for downgrade.
	BillingPlanDowngrade *string `json:"billingPlanDowngrade"`
	// Tax Id
	BillingTaxId *string `json:"billingTaxId"`
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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Organization) UnmarshalJSON(data []byte) error {
	type alias Organization
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                json.RawMessage `json:"$id"`
		CreatedAt         json.RawMessage `json:"$createdAt"`
		UpdatedAt         json.RawMessage `json:"$updatedAt"`
		Name              json.RawMessage `json:"name"`
		Total             json.RawMessage `json:"total"`
		BudgetAlerts      json.RawMessage `json:"budgetAlerts"`
		BillingPlan       json.RawMessage `json:"billingPlan"`
		BillingPlanId     json.RawMessage `json:"billingPlanId"`
		BillingEmail      json.RawMessage `json:"billingEmail"`
		Status            json.RawMessage `json:"status"`
		MarkedForDeletion json.RawMessage `json:"markedForDeletion"`
		Platform          json.RawMessage `json:"platform"`
		Projects          json.RawMessage `json:"projects"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.UpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UpdatedAt); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Total == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Total); err != nil {
			return err
		}
	}
	if present.BudgetAlerts == nil {
		if err := json.Unmarshal([]byte("[75]"), &decoded.BudgetAlerts); err != nil {
			return err
		}
	}
	if present.BillingPlan == nil {
		if err := json.Unmarshal([]byte("\"tier-0\""), &decoded.BillingPlan); err != nil {
			return err
		}
	}
	if present.BillingPlanId == nil {
		if err := json.Unmarshal([]byte("\"tier-0\""), &decoded.BillingPlanId); err != nil {
			return err
		}
	}
	if present.BillingEmail == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.BillingEmail); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"draft\""), &decoded.Status); err != nil {
			return err
		}
	}
	if present.MarkedForDeletion == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.MarkedForDeletion); err != nil {
			return err
		}
	}
	if present.Platform == nil {
		if err := json.Unmarshal([]byte("\"appwrite\""), &decoded.Platform); err != nil {
			return err
		}
	}
	if present.Projects == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Projects); err != nil {
			return err
		}
	}

	*model = Organization(decoded)
	return nil
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
