package logs

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/models"
)

// Logs service
type Logs struct {
	client client.Client
}

func New(clt client.Client) *Logs {
	return &Logs{
		client: clt,
	}
}

type ListOptions struct {
	ResourceType   string
	ResourceId     string
	Queries        []string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{"ResourceType": false, "ResourceId": false, "Queries": false, "Total": false}
	return &options
}

type ListOption func(*ListOptions)

func (srv *Logs) WithListResourceType(v string) ListOption {
	return func(o *ListOptions) {
		o.ResourceType = v
		o.enabledSetters["ResourceType"] = true
	}
}
func (srv *Logs) WithListResourceId(v string) ListOption {
	return func(o *ListOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
func (srv *Logs) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Logs) WithListTotal(v bool) ListOption {
	return func(o *ListOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// List list the log lines a resource wrote to stdout and stderr, newest
// first. Leave out resourceType and resourceId to list every resource in the
// project. Each row is stamped with the region that served it. Results are
// limited to your plan's log retention.
func (srv *Logs) List(optionalSetters ...ListOption) (*models.RuntimeLogList, error) {
	path := "/logs"
	options := ListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ResourceType"] {
		params["resourceType"] = options.ResourceType
	}
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
	}
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

		parsed := models.RuntimeLogList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.RuntimeLogList
	parsed, ok := resp.Result.(models.RuntimeLogList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Get get a log line by its unique ID. The log must belong to the given
// resource and fall within your plan's log retention.
func (srv *Logs) Get(LogId string, ResourceType string, ResourceId string) (*models.RuntimeLog, error) {
	if LogId == "" {
		return nil, errors.New("Missing required parameter: \"logId\"")
	}

	r := strings.NewReplacer("{logId}", client.EncodePath(LogId))
	path := r.Replace("/logs/{logId}")
	params := map[string]interface{}{}
	params["resourceType"] = ResourceType
	params["resourceId"] = ResourceId
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

		parsed := models.RuntimeLog{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.RuntimeLog
	parsed, ok := resp.Result.(models.RuntimeLog)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
