package analytics

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/models"
)

// Analytics service
type Analytics struct {
	client client.Client
}

func New(clt client.Client) *Analytics {
	return &Analytics{
		client: clt,
	}
}

type ListPropertiesOptions struct {
	Queries        []string
	Search         string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListPropertiesOptions) New() *ListPropertiesOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Search": false, "Total": false}
	return &options
}

type ListPropertiesOption func(*ListPropertiesOptions)

func (srv *Analytics) WithListPropertiesQueries(v []string) ListPropertiesOption {
	return func(o *ListPropertiesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Analytics) WithListPropertiesSearch(v string) ListPropertiesOption {
	return func(o *ListPropertiesOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Analytics) WithListPropertiesTotal(v bool) ListPropertiesOption {
	return func(o *ListPropertiesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListProperties list analytics properties for the current project.
func (srv *Analytics) ListProperties(optionalSetters ...ListPropertiesOption) (*models.AnalyticsPropertyList, error) {
	path := "/analytics/properties"
	options := ListPropertiesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AnalyticsPropertyList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AnalyticsPropertyList
	parsed, ok := resp.Result.(models.AnalyticsPropertyList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreatePropertyOptions struct {
	Domain         string
	Timezone       string
	Enabled        bool
	Public         bool
	AllowedOrigins []string
	enabledSetters map[string]bool
}

func (options CreatePropertyOptions) New() *CreatePropertyOptions {
	options.enabledSetters = map[string]bool{"Domain": false, "Timezone": false, "Enabled": false, "Public": false, "AllowedOrigins": false}
	return &options
}

type CreatePropertyOption func(*CreatePropertyOptions)

func (srv *Analytics) WithCreatePropertyDomain(v string) CreatePropertyOption {
	return func(o *CreatePropertyOptions) {
		o.Domain = v
		o.enabledSetters["Domain"] = true
	}
}
func (srv *Analytics) WithCreatePropertyTimezone(v string) CreatePropertyOption {
	return func(o *CreatePropertyOptions) {
		o.Timezone = v
		o.enabledSetters["Timezone"] = true
	}
}
func (srv *Analytics) WithCreatePropertyEnabled(v bool) CreatePropertyOption {
	return func(o *CreatePropertyOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Analytics) WithCreatePropertyPublic(v bool) CreatePropertyOption {
	return func(o *CreatePropertyOptions) {
		o.Public = v
		o.enabledSetters["Public"] = true
	}
}
func (srv *Analytics) WithCreatePropertyAllowedOrigins(v []string) CreatePropertyOption {
	return func(o *CreatePropertyOptions) {
		o.AllowedOrigins = v
		o.enabledSetters["AllowedOrigins"] = true
	}
}

// CreateProperty create a new analytics property to track a website or
// application.
func (srv *Analytics) CreateProperty(PropertyId string, Name string, optionalSetters ...CreatePropertyOption) (*models.AnalyticsProperty, error) {
	path := "/analytics/properties"
	options := CreatePropertyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["propertyId"] = PropertyId
	params["name"] = Name
	if options.enabledSetters["Domain"] {
		params["domain"] = options.Domain
	}
	if options.enabledSetters["Timezone"] {
		params["timezone"] = options.Timezone
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Public"] {
		params["public"] = options.Public
	}
	if options.enabledSetters["AllowedOrigins"] {
		params["allowedOrigins"] = options.AllowedOrigins
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AnalyticsProperty{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AnalyticsProperty
	parsed, ok := resp.Result.(models.AnalyticsProperty)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetProperty get an analytics property by ID.
func (srv *Analytics) GetProperty(PropertyId string) (*models.AnalyticsProperty, error) {
	if PropertyId == "" {
		return nil, errors.New("Missing required parameter: \"propertyId\"")
	}

	r := strings.NewReplacer("{propertyId}", client.EncodePath(PropertyId))
	path := r.Replace("/analytics/properties/{propertyId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AnalyticsProperty{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AnalyticsProperty
	parsed, ok := resp.Result.(models.AnalyticsProperty)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdatePropertyOptions struct {
	Name           string
	Domain         string
	Timezone       string
	Enabled        bool
	Public         bool
	AllowedOrigins []string
	enabledSetters map[string]bool
}

func (options UpdatePropertyOptions) New() *UpdatePropertyOptions {
	options.enabledSetters = map[string]bool{"Name": false, "Domain": false, "Timezone": false, "Enabled": false, "Public": false, "AllowedOrigins": false}
	return &options
}

type UpdatePropertyOption func(*UpdatePropertyOptions)

func (srv *Analytics) WithUpdatePropertyName(v string) UpdatePropertyOption {
	return func(o *UpdatePropertyOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Analytics) WithUpdatePropertyDomain(v string) UpdatePropertyOption {
	return func(o *UpdatePropertyOptions) {
		o.Domain = v
		o.enabledSetters["Domain"] = true
	}
}
func (srv *Analytics) WithUpdatePropertyTimezone(v string) UpdatePropertyOption {
	return func(o *UpdatePropertyOptions) {
		o.Timezone = v
		o.enabledSetters["Timezone"] = true
	}
}
func (srv *Analytics) WithUpdatePropertyEnabled(v bool) UpdatePropertyOption {
	return func(o *UpdatePropertyOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Analytics) WithUpdatePropertyPublic(v bool) UpdatePropertyOption {
	return func(o *UpdatePropertyOptions) {
		o.Public = v
		o.enabledSetters["Public"] = true
	}
}
func (srv *Analytics) WithUpdatePropertyAllowedOrigins(v []string) UpdatePropertyOption {
	return func(o *UpdatePropertyOptions) {
		o.AllowedOrigins = v
		o.enabledSetters["AllowedOrigins"] = true
	}
}

// UpdateProperty update an analytics property. Only the attributes you pass
// are changed; omitted attributes keep their current value.
func (srv *Analytics) UpdateProperty(PropertyId string, optionalSetters ...UpdatePropertyOption) (*models.AnalyticsProperty, error) {
	if PropertyId == "" {
		return nil, errors.New("Missing required parameter: \"propertyId\"")
	}

	r := strings.NewReplacer("{propertyId}", client.EncodePath(PropertyId))
	path := r.Replace("/analytics/properties/{propertyId}")
	options := UpdatePropertyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Domain"] {
		params["domain"] = options.Domain
	}
	if options.enabledSetters["Timezone"] {
		params["timezone"] = options.Timezone
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Public"] {
		params["public"] = options.Public
	}
	if options.enabledSetters["AllowedOrigins"] {
		params["allowedOrigins"] = options.AllowedOrigins
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AnalyticsProperty{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AnalyticsProperty
	parsed, ok := resp.Result.(models.AnalyticsProperty)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteProperty delete an analytics property along with every event and
// session collected for it. This cannot be undone.
func (srv *Analytics) DeleteProperty(PropertyId string) (*interface{}, error) {
	if PropertyId == "" {
		return nil, errors.New("Missing required parameter: \"propertyId\"")
	}

	r := strings.NewReplacer("{propertyId}", client.EncodePath(PropertyId))
	path := r.Replace("/analytics/properties/{propertyId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		var parsed interface{}

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed interface{}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateEventOptions struct {
	Domain         string
	Referrer       string
	ScreenWidth    int
	SessionHash    string
	ScrollDepth    int
	EngagementTime int
	Props          []string
	UserId         string
	Ip             string
	UserAgent      string
	enabledSetters map[string]bool
}

func (options CreateEventOptions) New() *CreateEventOptions {
	options.enabledSetters = map[string]bool{"Domain": false, "Referrer": false, "ScreenWidth": false, "SessionHash": false, "ScrollDepth": false, "EngagementTime": false, "Props": false, "UserId": false, "Ip": false, "UserAgent": false}
	return &options
}

type CreateEventOption func(*CreateEventOptions)

func (srv *Analytics) WithCreateEventDomain(v string) CreateEventOption {
	return func(o *CreateEventOptions) {
		o.Domain = v
		o.enabledSetters["Domain"] = true
	}
}
func (srv *Analytics) WithCreateEventReferrer(v string) CreateEventOption {
	return func(o *CreateEventOptions) {
		o.Referrer = v
		o.enabledSetters["Referrer"] = true
	}
}
func (srv *Analytics) WithCreateEventScreenWidth(v int) CreateEventOption {
	return func(o *CreateEventOptions) {
		o.ScreenWidth = v
		o.enabledSetters["ScreenWidth"] = true
	}
}
func (srv *Analytics) WithCreateEventSessionHash(v string) CreateEventOption {
	return func(o *CreateEventOptions) {
		o.SessionHash = v
		o.enabledSetters["SessionHash"] = true
	}
}
func (srv *Analytics) WithCreateEventScrollDepth(v int) CreateEventOption {
	return func(o *CreateEventOptions) {
		o.ScrollDepth = v
		o.enabledSetters["ScrollDepth"] = true
	}
}
func (srv *Analytics) WithCreateEventEngagementTime(v int) CreateEventOption {
	return func(o *CreateEventOptions) {
		o.EngagementTime = v
		o.enabledSetters["EngagementTime"] = true
	}
}
func (srv *Analytics) WithCreateEventProps(v []string) CreateEventOption {
	return func(o *CreateEventOptions) {
		o.Props = v
		o.enabledSetters["Props"] = true
	}
}
func (srv *Analytics) WithCreateEventUserId(v string) CreateEventOption {
	return func(o *CreateEventOptions) {
		o.UserId = v
		o.enabledSetters["UserId"] = true
	}
}
func (srv *Analytics) WithCreateEventIp(v string) CreateEventOption {
	return func(o *CreateEventOptions) {
		o.Ip = v
		o.enabledSetters["Ip"] = true
	}
}
func (srv *Analytics) WithCreateEventUserAgent(v string) CreateEventOption {
	return func(o *CreateEventOptions) {
		o.UserAgent = v
		o.enabledSetters["UserAgent"] = true
	}
}

// CreateEvent send a tracking event from a browser, native app, or
// server-side SDK.
func (srv *Analytics) CreateEvent(PropertyId string, Name string, Url string, optionalSetters ...CreateEventOption) (*interface{}, error) {
	if PropertyId == "" {
		return nil, errors.New("Missing required parameter: \"propertyId\"")
	}

	r := strings.NewReplacer("{propertyId}", client.EncodePath(PropertyId))
	path := r.Replace("/analytics/properties/{propertyId}/events")
	options := CreateEventOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["url"] = Url
	if options.enabledSetters["Domain"] {
		params["domain"] = options.Domain
	}
	if options.enabledSetters["Referrer"] {
		params["referrer"] = options.Referrer
	}
	if options.enabledSetters["ScreenWidth"] {
		params["screenWidth"] = options.ScreenWidth
	}
	if options.enabledSetters["SessionHash"] {
		params["sessionHash"] = options.SessionHash
	}
	if options.enabledSetters["ScrollDepth"] {
		params["scrollDepth"] = options.ScrollDepth
	}
	if options.enabledSetters["EngagementTime"] {
		params["engagementTime"] = options.EngagementTime
	}
	if options.enabledSetters["Props"] {
		params["props"] = options.Props
	}
	if options.enabledSetters["UserId"] {
		params["userId"] = options.UserId
	}
	if options.enabledSetters["Ip"] {
		params["ip"] = options.Ip
	}
	if options.enabledSetters["UserAgent"] {
		params["userAgent"] = options.UserAgent
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		var parsed interface{}

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed interface{}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListMetricsOptions struct {
	Queries        []string
	Interval       string
	Dimensions     []string
	DateRange      string
	StartAt        string
	EndAt          string
	Limit          int
	enabledSetters map[string]bool
}

func (options ListMetricsOptions) New() *ListMetricsOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Interval": false, "Dimensions": false, "DateRange": false, "StartAt": false, "EndAt": false, "Limit": false}
	return &options
}

type ListMetricsOption func(*ListMetricsOptions)

func (srv *Analytics) WithListMetricsQueries(v []string) ListMetricsOption {
	return func(o *ListMetricsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Analytics) WithListMetricsInterval(v string) ListMetricsOption {
	return func(o *ListMetricsOptions) {
		o.Interval = v
		o.enabledSetters["Interval"] = true
	}
}
func (srv *Analytics) WithListMetricsDimensions(v []string) ListMetricsOption {
	return func(o *ListMetricsOptions) {
		o.Dimensions = v
		o.enabledSetters["Dimensions"] = true
	}
}
func (srv *Analytics) WithListMetricsDateRange(v string) ListMetricsOption {
	return func(o *ListMetricsOptions) {
		o.DateRange = v
		o.enabledSetters["DateRange"] = true
	}
}
func (srv *Analytics) WithListMetricsStartAt(v string) ListMetricsOption {
	return func(o *ListMetricsOptions) {
		o.StartAt = v
		o.enabledSetters["StartAt"] = true
	}
}
func (srv *Analytics) WithListMetricsEndAt(v string) ListMetricsOption {
	return func(o *ListMetricsOptions) {
		o.EndAt = v
		o.enabledSetters["EndAt"] = true
	}
}
func (srv *Analytics) WithListMetricsLimit(v int) ListMetricsOption {
	return func(o *ListMetricsOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}

// ListMetrics read analytics metrics (visitors, sessions, pageviews, events,
// bounceRate, …) for a property over a date range.
//
// **Three response shapes**, chosen by `dimensions[]` and `interval`:
// - Neither: one row aggregating the whole window, with `value` and `date`
// null. Only this shape carries `pageviews`, `visits`, `bounceRate`,
// `visitDuration`, `viewsPerVisit`, `scrollDepth` and `engagementTime`.
// - `dimensions[]`: one row per dimension value, ranked by visitors, with
// `value` set and `date` null.
// - `interval`: one row per time bucket in chronological order, with `date`
// set and `value` null.
//
// Combining `dimensions[]` with `interval` is not supported yet. `queries[]`
// filters the underlying events using standard Utopia query syntax.
func (srv *Analytics) ListMetrics(PropertyId string, optionalSetters ...ListMetricsOption) (*models.AnalyticsMetricList, error) {
	if PropertyId == "" {
		return nil, errors.New("Missing required parameter: \"propertyId\"")
	}

	r := strings.NewReplacer("{propertyId}", client.EncodePath(PropertyId))
	path := r.Replace("/analytics/properties/{propertyId}/metrics")
	options := ListMetricsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Interval"] {
		params["interval"] = options.Interval
	}
	if options.enabledSetters["Dimensions"] {
		params["dimensions"] = options.Dimensions
	}
	if options.enabledSetters["DateRange"] {
		params["dateRange"] = options.DateRange
	}
	if options.enabledSetters["StartAt"] {
		params["startAt"] = options.StartAt
	}
	if options.enabledSetters["EndAt"] {
		params["endAt"] = options.EndAt
	}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AnalyticsMetricList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AnalyticsMetricList
	parsed, ok := resp.Result.(models.AnalyticsMetricList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
