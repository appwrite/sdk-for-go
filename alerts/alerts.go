package alerts

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/models"
)

// Alerts service
type Alerts struct {
	client client.Client
}

func New(clt client.Client) *Alerts {
	return &Alerts{
		client: clt,
	}
}

type ListOptions struct {
	Queries        []string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Total": false}
	return &options
}

type ListOption func(*ListOptions)

func (srv *Alerts) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Alerts) WithListTotal(v bool) ListOption {
	return func(o *ListOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// List get a list of the alerts the project's monitors raised, open and
// resolved. An alert is created when a monitor starts firing and resolved
// when it recovers. Use queries to filter by monitor, status, type or
// resource, order by `$createdAt` for the latest first, and
// `Query.search('search', term)` to search alert and monitor IDs, names and
// metrics.
func (srv *Alerts) List(optionalSetters ...ListOption) (*models.AlertList, error) {
	path := "/alerts"
	options := ListOptions{}.New()
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

		parsed := models.AlertList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AlertList
	parsed, ok := resp.Result.(models.AlertList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Get get an alert by its unique ID: the monitor that raised it, when it
// fired and resolved, and the readings while it fired.
func (srv *Alerts) Get(AlertId string) (*models.Alert, error) {
	if AlertId == "" {
		return nil, errors.New("Missing required parameter: \"alertId\"")
	}

	r := strings.NewReplacer("{alertId}", client.EncodePath(AlertId))
	path := r.Replace("/alerts/{alertId}")
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

		parsed := models.Alert{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Alert
	parsed, ok := resp.Result.(models.Alert)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Update acknowledge a firing alert, or undo that. Everyone who can read the
// alert sees who acknowledged it and when. The alert keeps firing until its
// monitor recovers; a resolved alert keeps its acknowledgement and cannot be
// changed.
func (srv *Alerts) Update(AlertId string, Acknowledged bool) (*models.Alert, error) {
	if AlertId == "" {
		return nil, errors.New("Missing required parameter: \"alertId\"")
	}

	r := strings.NewReplacer("{alertId}", client.EncodePath(AlertId))
	path := r.Replace("/alerts/{alertId}")
	params := map[string]interface{}{}
	params["acknowledged"] = Acknowledged
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

		parsed := models.Alert{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Alert
	parsed, ok := resp.Result.(models.Alert)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
