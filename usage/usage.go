package usage

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/v4/client"
	"github.com/appwrite/sdk-for-go/v4/models"
	"strings"
)

// Usage service
type Usage struct {
	client client.Client
}

func New(clt client.Client) *Usage {
	return &Usage{
		client: clt,
	}
}

type ListEventsOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListEventsOptions) New() *ListEventsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListEventsOption func(*ListEventsOptions)
func (srv *Usage) WithListEventsQueries(v []string) ListEventsOption {
	return func(o *ListEventsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Usage) WithListEventsTotal(v bool) ListEventsOption {
	return func(o *ListEventsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListEvents query usage event metrics from the usage database. Returns
// individual event rows with full metadata. Pass Query objects as JSON
// strings to filter, paginate, and order results. Supported query methods:
// equal, greaterThanEqual, lessThanEqual, orderAsc, orderDesc, limit, offset.
// Supported filter attributes: metric, path, method, status, resource,
// resourceId, country, userAgent, time (these match the underlying column
// names — note that the response surfaces `resource` as `resourceType` and
// `country` as `countryCode`). When no time filter is supplied the endpoint
// defaults to the last 7 days. Default `limit(100)` is applied if none is
// given; user-supplied limits are capped at 500. The `total` field is capped
// at 5000 to keep counts predictable — pass `total=false` to skip the count
// entirely.
func (srv *Usage) ListEvents(optionalSetters ...ListEventsOption)(*models.UsageEventList, error) {
	path := "/usage/events"
	options := ListEventsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.UsageEventList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.UsageEventList
	parsed, ok := resp.Result.(models.UsageEventList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListGaugesOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListGaugesOptions) New() *ListGaugesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListGaugesOption func(*ListGaugesOptions)
func (srv *Usage) WithListGaugesQueries(v []string) ListGaugesOption {
	return func(o *ListGaugesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Usage) WithListGaugesTotal(v bool) ListGaugesOption {
	return func(o *ListGaugesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListGauges query usage gauge metrics (point-in-time resource snapshots)
// from the usage database. Returns individual gauge snapshots with metric,
// value, and timestamp. Pass Query objects as JSON strings to filter,
// paginate, and order results. Supported query methods: equal,
// greaterThanEqual, lessThanEqual, orderAsc, orderDesc, limit, offset.
// Supported filter attributes: metric, time. Use `orderDesc("time"),
// limit(1)` to fetch the most recent snapshot. When no time filter is
// supplied the endpoint defaults to the last 7 days. Default `limit(100)` is
// applied if none is given; user-supplied limits are capped at 500. The
// `total` field is capped at 5000 to keep counts predictable — pass
// `total=false` to skip the count entirely.
func (srv *Usage) ListGauges(optionalSetters ...ListGaugesOption)(*models.UsageGaugeList, error) {
	path := "/usage/gauges"
	options := ListGaugesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.UsageGaugeList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.UsageGaugeList
	parsed, ok := resp.Result.(models.UsageGaugeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
