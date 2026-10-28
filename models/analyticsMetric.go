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
	// Total page and screen view events (events named "pageview" or
	// "screen_view") in the requested range. The flat aggregate and the time
	// series compute it; null on breakdown rows.
	Pageviews *int `json:"pageviews"`
	// Total events in the requested range.
	Events int `json:"events"`
	// Total sessions with activity in the requested range. Only the flat
	// aggregate computes it; null on breakdown and time-series rows.
	Visits *int `json:"visits"`
	// Share of bounced visits, as a percentage; a visit is a bounce unless it
	// records two or more page or screen views, at least one interaction, or ten
	// seconds or more of activity. An interaction is any event that is neither a
	// view nor automatic measurement such as "app_foregrounded",
	// "engagement_time" or "scroll_depth". Only the flat aggregate computes it;
	// null on breakdown and time-series rows.
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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AnalyticsMetric) UnmarshalJSON(data []byte) error {
	type alias AnalyticsMetric
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Visitors json.RawMessage `json:"visitors"`
		Sessions json.RawMessage `json:"sessions"`
		Events   json.RawMessage `json:"events"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Visitors == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Visitors); err != nil {
			return err
		}
	}
	if present.Sessions == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Sessions); err != nil {
			return err
		}
	}
	if present.Events == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Events); err != nil {
			return err
		}
	}

	*model = AnalyticsMetric(decoded)
	return nil
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
