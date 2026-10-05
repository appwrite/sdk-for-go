package models

import (
	"encoding/json"
	"errors"
)

// AnalyticsMetric Model
type AnalyticsMetric struct {
	// Dimension value this row covers. Null when no dimension was requested,
	// empty when the dimension could not be derived for those events.
	Value *string `json:"value"`
	// Start of the time bucket this row covers, in ISO 8601. Null when no
	// interval was requested.
	Date *string `json:"date"`
	// Unique visitors in the requested range.
	Visitors int `json:"visitors"`
	// Unique sessions in the requested range.
	Sessions int `json:"sessions"`
	// Total pageview events (events with name="pageview") in the requested range.
	// Only the flat aggregate computes it; null on breakdown and time-series
	// rows.
	Pageviews *int `json:"pageviews"`
	// Total events in the requested range.
	Events int `json:"events"`
	// Total sessions with activity in the requested range. Only the flat
	// aggregate computes it; null on breakdown and time-series rows.
	Visits *int `json:"visits"`
	// Share of single-event sessions, as a percentage. Only the flat aggregate
	// computes it; null on breakdown and time-series rows.
	BounceRate *float64 `json:"bounceRate"`
	// Average session length in seconds, measured first event to last event.
	// Single-event sessions count as 0. Only the flat aggregate computes it; null
	// on breakdown and time-series rows.
	VisitDuration *float64 `json:"visitDuration"`
	// Average pageviews per session (pageviews divided by visits). Only the flat
	// aggregate computes it; null on breakdown and time-series rows.
	ViewsPerVisit *float64 `json:"viewsPerVisit"`
	// Average scroll depth across events, as a percentage. Only the flat
	// aggregate computes it; null on breakdown and time-series rows.
	ScrollDepth *float64 `json:"scrollDepth"`
	// Average engaged time per session in seconds, counting only foreground time.
	// Sessions that reported no engagement are excluded. Only the flat aggregate
	// computes it; null on breakdown and time-series rows.
	EngagementTime *float64 `json:"engagementTime"`

	// Used by Decode() method
	data []byte
}

func (model AnalyticsMetric) New(data []byte) *AnalyticsMetric {
	model.data = data
	return &model
}

func (model *AnalyticsMetric) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
