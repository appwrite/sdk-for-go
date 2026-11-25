package monitors

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/models"
)

// Monitors service
type Monitors struct {
	client client.Client
}

func New(clt client.Client) *Monitors {
	return &Monitors{
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

func (srv *Monitors) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Monitors) WithListTotal(v bool) ListOption {
	return func(o *ListOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// List get a list of the project's monitors of every type. Use queries to
// filter by type, resource, metric or status, and `Query.search('search',
// term)` to search monitor IDs, names, metrics and resources.
func (srv *Monitors) List(optionalSetters ...ListOption) (*models.MonitorList, error) {
	path := "/monitors"
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

		parsed := models.MonitorList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MonitorList
	parsed, ok := resp.Result.(models.MonitorList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateDnsMonitorOptions struct {
	Record         string
	Resolver       string
	Rules          []string
	Interval       int
	Timeout        int
	Cooldown       int
	Enabled        bool
	Notify         []string
	Severity       string
	EscalationId   string
	enabledSetters map[string]bool
}

func (options CreateDnsMonitorOptions) New() *CreateDnsMonitorOptions {
	options.enabledSetters = map[string]bool{"Record": false, "Resolver": false, "Rules": false, "Interval": false, "Timeout": false, "Cooldown": false, "Enabled": false, "Notify": false, "Severity": false, "EscalationId": false}
	return &options
}

type CreateDnsMonitorOption func(*CreateDnsMonitorOptions)

func (srv *Monitors) WithCreateDnsMonitorRecord(v string) CreateDnsMonitorOption {
	return func(o *CreateDnsMonitorOptions) {
		o.Record = v
		o.enabledSetters["Record"] = true
	}
}
func (srv *Monitors) WithCreateDnsMonitorResolver(v string) CreateDnsMonitorOption {
	return func(o *CreateDnsMonitorOptions) {
		o.Resolver = v
		o.enabledSetters["Resolver"] = true
	}
}
func (srv *Monitors) WithCreateDnsMonitorRules(v []string) CreateDnsMonitorOption {
	return func(o *CreateDnsMonitorOptions) {
		o.Rules = v
		o.enabledSetters["Rules"] = true
	}
}
func (srv *Monitors) WithCreateDnsMonitorInterval(v int) CreateDnsMonitorOption {
	return func(o *CreateDnsMonitorOptions) {
		o.Interval = v
		o.enabledSetters["Interval"] = true
	}
}
func (srv *Monitors) WithCreateDnsMonitorTimeout(v int) CreateDnsMonitorOption {
	return func(o *CreateDnsMonitorOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
func (srv *Monitors) WithCreateDnsMonitorCooldown(v int) CreateDnsMonitorOption {
	return func(o *CreateDnsMonitorOptions) {
		o.Cooldown = v
		o.enabledSetters["Cooldown"] = true
	}
}
func (srv *Monitors) WithCreateDnsMonitorEnabled(v bool) CreateDnsMonitorOption {
	return func(o *CreateDnsMonitorOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Monitors) WithCreateDnsMonitorNotify(v []string) CreateDnsMonitorOption {
	return func(o *CreateDnsMonitorOptions) {
		o.Notify = v
		o.enabledSetters["Notify"] = true
	}
}
func (srv *Monitors) WithCreateDnsMonitorSeverity(v string) CreateDnsMonitorOption {
	return func(o *CreateDnsMonitorOptions) {
		o.Severity = v
		o.enabledSetters["Severity"] = true
	}
}
func (srv *Monitors) WithCreateDnsMonitorEscalationId(v string) CreateDnsMonitorOption {
	return func(o *CreateDnsMonitorOptions) {
		o.EscalationId = v
		o.enabledSetters["EscalationId"] = true
	}
}

// CreateDnsMonitor create a DNS monitor. It asks a resolver for a domain's
// records of one type on its own interval, and opens an alert when the
// resolver fails, or when its rules match the answers, by default when there
// is no such record. Each alert fires the
// `monitors.[monitorId].alerts.[alertId].firing` event when it opens and
// `.resolved` when a check passes again, for project webhooks, functions and
// realtime, and notifies the organization's members by email and in the
// console.
func (srv *Monitors) CreateDnsMonitor(MonitorId string, Name string, Domain string, optionalSetters ...CreateDnsMonitorOption) (*models.Monitor, error) {
	path := "/monitors/dns"
	options := CreateDnsMonitorOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["monitorId"] = MonitorId
	params["name"] = Name
	params["domain"] = Domain
	if options.enabledSetters["Record"] {
		params["record"] = options.Record
	}
	if options.enabledSetters["Resolver"] {
		params["resolver"] = options.Resolver
	}
	if options.enabledSetters["Rules"] {
		params["rules"] = options.Rules
	}
	if options.enabledSetters["Interval"] {
		params["interval"] = options.Interval
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
	}
	if options.enabledSetters["Cooldown"] {
		params["cooldown"] = options.Cooldown
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Notify"] {
		params["notify"] = options.Notify
	}
	if options.enabledSetters["Severity"] {
		params["severity"] = options.Severity
	}
	if options.enabledSetters["EscalationId"] {
		params["escalationId"] = options.EscalationId
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

		parsed := models.Monitor{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Monitor
	parsed, ok := resp.Result.(models.Monitor)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateDnsMonitorOptions struct {
	Name           string
	Domain         string
	Record         string
	Resolver       string
	Rules          []string
	Interval       int
	Timeout        int
	Cooldown       int
	Enabled        bool
	Notify         []string
	Severity       string
	EscalationId   string
	enabledSetters map[string]bool
}

func (options UpdateDnsMonitorOptions) New() *UpdateDnsMonitorOptions {
	options.enabledSetters = map[string]bool{"Name": false, "Domain": false, "Record": false, "Resolver": false, "Rules": false, "Interval": false, "Timeout": false, "Cooldown": false, "Enabled": false, "Notify": false, "Severity": false, "EscalationId": false}
	return &options
}

type UpdateDnsMonitorOption func(*UpdateDnsMonitorOptions)

func (srv *Monitors) WithUpdateDnsMonitorName(v string) UpdateDnsMonitorOption {
	return func(o *UpdateDnsMonitorOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Monitors) WithUpdateDnsMonitorDomain(v string) UpdateDnsMonitorOption {
	return func(o *UpdateDnsMonitorOptions) {
		o.Domain = v
		o.enabledSetters["Domain"] = true
	}
}
func (srv *Monitors) WithUpdateDnsMonitorRecord(v string) UpdateDnsMonitorOption {
	return func(o *UpdateDnsMonitorOptions) {
		o.Record = v
		o.enabledSetters["Record"] = true
	}
}
func (srv *Monitors) WithUpdateDnsMonitorResolver(v string) UpdateDnsMonitorOption {
	return func(o *UpdateDnsMonitorOptions) {
		o.Resolver = v
		o.enabledSetters["Resolver"] = true
	}
}
func (srv *Monitors) WithUpdateDnsMonitorRules(v []string) UpdateDnsMonitorOption {
	return func(o *UpdateDnsMonitorOptions) {
		o.Rules = v
		o.enabledSetters["Rules"] = true
	}
}
func (srv *Monitors) WithUpdateDnsMonitorInterval(v int) UpdateDnsMonitorOption {
	return func(o *UpdateDnsMonitorOptions) {
		o.Interval = v
		o.enabledSetters["Interval"] = true
	}
}
func (srv *Monitors) WithUpdateDnsMonitorTimeout(v int) UpdateDnsMonitorOption {
	return func(o *UpdateDnsMonitorOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
func (srv *Monitors) WithUpdateDnsMonitorCooldown(v int) UpdateDnsMonitorOption {
	return func(o *UpdateDnsMonitorOptions) {
		o.Cooldown = v
		o.enabledSetters["Cooldown"] = true
	}
}
func (srv *Monitors) WithUpdateDnsMonitorEnabled(v bool) UpdateDnsMonitorOption {
	return func(o *UpdateDnsMonitorOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Monitors) WithUpdateDnsMonitorNotify(v []string) UpdateDnsMonitorOption {
	return func(o *UpdateDnsMonitorOptions) {
		o.Notify = v
		o.enabledSetters["Notify"] = true
	}
}
func (srv *Monitors) WithUpdateDnsMonitorSeverity(v string) UpdateDnsMonitorOption {
	return func(o *UpdateDnsMonitorOptions) {
		o.Severity = v
		o.enabledSetters["Severity"] = true
	}
}
func (srv *Monitors) WithUpdateDnsMonitorEscalationId(v string) UpdateDnsMonitorOption {
	return func(o *UpdateDnsMonitorOptions) {
		o.EscalationId = v
		o.enabledSetters["EscalationId"] = true
	}
}

// UpdateDnsMonitor update a DNS monitor by its unique ID. Only the parameters
// passed change; the monitor's state, its last check and its open alert are
// kept. Passing `rules` replaces all of them.
func (srv *Monitors) UpdateDnsMonitor(MonitorId string, optionalSetters ...UpdateDnsMonitorOption) (*models.Monitor, error) {
	if MonitorId == "" {
		return nil, errors.New("Missing required parameter: \"monitorId\"")
	}

	r := strings.NewReplacer("{monitorId}", client.EncodePath(MonitorId))
	path := r.Replace("/monitors/dns/{monitorId}")
	options := UpdateDnsMonitorOptions{}.New()
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
	if options.enabledSetters["Record"] {
		params["record"] = options.Record
	}
	if options.enabledSetters["Resolver"] {
		params["resolver"] = options.Resolver
	}
	if options.enabledSetters["Rules"] {
		params["rules"] = options.Rules
	}
	if options.enabledSetters["Interval"] {
		params["interval"] = options.Interval
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
	}
	if options.enabledSetters["Cooldown"] {
		params["cooldown"] = options.Cooldown
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Notify"] {
		params["notify"] = options.Notify
	}
	if options.enabledSetters["Severity"] {
		params["severity"] = options.Severity
	}
	if options.enabledSetters["EscalationId"] {
		params["escalationId"] = options.EscalationId
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

		parsed := models.Monitor{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Monitor
	parsed, ok := resp.Result.(models.Monitor)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateHttpMonitorOptions struct {
	Method         string
	Headers        interface{}
	Body           string
	Rules          []string
	Interval       int
	Timeout        int
	Cooldown       int
	Enabled        bool
	Notify         []string
	Severity       string
	EscalationId   string
	enabledSetters map[string]bool
}

func (options CreateHttpMonitorOptions) New() *CreateHttpMonitorOptions {
	options.enabledSetters = map[string]bool{"Method": false, "Headers": false, "Body": false, "Rules": false, "Interval": false, "Timeout": false, "Cooldown": false, "Enabled": false, "Notify": false, "Severity": false, "EscalationId": false}
	return &options
}

type CreateHttpMonitorOption func(*CreateHttpMonitorOptions)

func (srv *Monitors) WithCreateHttpMonitorMethod(v string) CreateHttpMonitorOption {
	return func(o *CreateHttpMonitorOptions) {
		o.Method = v
		o.enabledSetters["Method"] = true
	}
}
func (srv *Monitors) WithCreateHttpMonitorHeaders(v interface{}) CreateHttpMonitorOption {
	return func(o *CreateHttpMonitorOptions) {
		o.Headers = v
		o.enabledSetters["Headers"] = true
	}
}
func (srv *Monitors) WithCreateHttpMonitorBody(v string) CreateHttpMonitorOption {
	return func(o *CreateHttpMonitorOptions) {
		o.Body = v
		o.enabledSetters["Body"] = true
	}
}
func (srv *Monitors) WithCreateHttpMonitorRules(v []string) CreateHttpMonitorOption {
	return func(o *CreateHttpMonitorOptions) {
		o.Rules = v
		o.enabledSetters["Rules"] = true
	}
}
func (srv *Monitors) WithCreateHttpMonitorInterval(v int) CreateHttpMonitorOption {
	return func(o *CreateHttpMonitorOptions) {
		o.Interval = v
		o.enabledSetters["Interval"] = true
	}
}
func (srv *Monitors) WithCreateHttpMonitorTimeout(v int) CreateHttpMonitorOption {
	return func(o *CreateHttpMonitorOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
func (srv *Monitors) WithCreateHttpMonitorCooldown(v int) CreateHttpMonitorOption {
	return func(o *CreateHttpMonitorOptions) {
		o.Cooldown = v
		o.enabledSetters["Cooldown"] = true
	}
}
func (srv *Monitors) WithCreateHttpMonitorEnabled(v bool) CreateHttpMonitorOption {
	return func(o *CreateHttpMonitorOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Monitors) WithCreateHttpMonitorNotify(v []string) CreateHttpMonitorOption {
	return func(o *CreateHttpMonitorOptions) {
		o.Notify = v
		o.enabledSetters["Notify"] = true
	}
}
func (srv *Monitors) WithCreateHttpMonitorSeverity(v string) CreateHttpMonitorOption {
	return func(o *CreateHttpMonitorOptions) {
		o.Severity = v
		o.enabledSetters["Severity"] = true
	}
}
func (srv *Monitors) WithCreateHttpMonitorEscalationId(v string) CreateHttpMonitorOption {
	return func(o *CreateHttpMonitorOptions) {
		o.EscalationId = v
		o.enabledSetters["EscalationId"] = true
	}
}

// CreateHttpMonitor create an HTTP monitor. It requests a URL on its own
// interval, following redirects, and opens an alert when the request fails or
// when its rules match the response, by default a status outside 2xx or the
// certificate of an HTTPS URL about to expire. The response body is discarded
// as it arrives. Each alert fires the
// `monitors.[monitorId].alerts.[alertId].firing` event when it opens and
// `.resolved` when a check passes again, for project webhooks, functions and
// realtime, and notifies the organization's members by email and in the
// console.
func (srv *Monitors) CreateHttpMonitor(MonitorId string, Name string, Url string, optionalSetters ...CreateHttpMonitorOption) (*models.Monitor, error) {
	path := "/monitors/http"
	options := CreateHttpMonitorOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["monitorId"] = MonitorId
	params["name"] = Name
	params["url"] = Url
	if options.enabledSetters["Method"] {
		params["method"] = options.Method
	}
	if options.enabledSetters["Headers"] {
		params["headers"] = options.Headers
	}
	if options.enabledSetters["Body"] {
		params["body"] = options.Body
	}
	if options.enabledSetters["Rules"] {
		params["rules"] = options.Rules
	}
	if options.enabledSetters["Interval"] {
		params["interval"] = options.Interval
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
	}
	if options.enabledSetters["Cooldown"] {
		params["cooldown"] = options.Cooldown
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Notify"] {
		params["notify"] = options.Notify
	}
	if options.enabledSetters["Severity"] {
		params["severity"] = options.Severity
	}
	if options.enabledSetters["EscalationId"] {
		params["escalationId"] = options.EscalationId
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

		parsed := models.Monitor{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Monitor
	parsed, ok := resp.Result.(models.Monitor)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateHttpMonitorOptions struct {
	Name           string
	Url            string
	Method         string
	Headers        interface{}
	Body           string
	Rules          []string
	Interval       int
	Timeout        int
	Cooldown       int
	Enabled        bool
	Notify         []string
	Severity       string
	EscalationId   string
	enabledSetters map[string]bool
}

func (options UpdateHttpMonitorOptions) New() *UpdateHttpMonitorOptions {
	options.enabledSetters = map[string]bool{"Name": false, "Url": false, "Method": false, "Headers": false, "Body": false, "Rules": false, "Interval": false, "Timeout": false, "Cooldown": false, "Enabled": false, "Notify": false, "Severity": false, "EscalationId": false}
	return &options
}

type UpdateHttpMonitorOption func(*UpdateHttpMonitorOptions)

func (srv *Monitors) WithUpdateHttpMonitorName(v string) UpdateHttpMonitorOption {
	return func(o *UpdateHttpMonitorOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Monitors) WithUpdateHttpMonitorUrl(v string) UpdateHttpMonitorOption {
	return func(o *UpdateHttpMonitorOptions) {
		o.Url = v
		o.enabledSetters["Url"] = true
	}
}
func (srv *Monitors) WithUpdateHttpMonitorMethod(v string) UpdateHttpMonitorOption {
	return func(o *UpdateHttpMonitorOptions) {
		o.Method = v
		o.enabledSetters["Method"] = true
	}
}
func (srv *Monitors) WithUpdateHttpMonitorHeaders(v interface{}) UpdateHttpMonitorOption {
	return func(o *UpdateHttpMonitorOptions) {
		o.Headers = v
		o.enabledSetters["Headers"] = true
	}
}
func (srv *Monitors) WithUpdateHttpMonitorBody(v string) UpdateHttpMonitorOption {
	return func(o *UpdateHttpMonitorOptions) {
		o.Body = v
		o.enabledSetters["Body"] = true
	}
}
func (srv *Monitors) WithUpdateHttpMonitorRules(v []string) UpdateHttpMonitorOption {
	return func(o *UpdateHttpMonitorOptions) {
		o.Rules = v
		o.enabledSetters["Rules"] = true
	}
}
func (srv *Monitors) WithUpdateHttpMonitorInterval(v int) UpdateHttpMonitorOption {
	return func(o *UpdateHttpMonitorOptions) {
		o.Interval = v
		o.enabledSetters["Interval"] = true
	}
}
func (srv *Monitors) WithUpdateHttpMonitorTimeout(v int) UpdateHttpMonitorOption {
	return func(o *UpdateHttpMonitorOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
func (srv *Monitors) WithUpdateHttpMonitorCooldown(v int) UpdateHttpMonitorOption {
	return func(o *UpdateHttpMonitorOptions) {
		o.Cooldown = v
		o.enabledSetters["Cooldown"] = true
	}
}
func (srv *Monitors) WithUpdateHttpMonitorEnabled(v bool) UpdateHttpMonitorOption {
	return func(o *UpdateHttpMonitorOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Monitors) WithUpdateHttpMonitorNotify(v []string) UpdateHttpMonitorOption {
	return func(o *UpdateHttpMonitorOptions) {
		o.Notify = v
		o.enabledSetters["Notify"] = true
	}
}
func (srv *Monitors) WithUpdateHttpMonitorSeverity(v string) UpdateHttpMonitorOption {
	return func(o *UpdateHttpMonitorOptions) {
		o.Severity = v
		o.enabledSetters["Severity"] = true
	}
}
func (srv *Monitors) WithUpdateHttpMonitorEscalationId(v string) UpdateHttpMonitorOption {
	return func(o *UpdateHttpMonitorOptions) {
		o.EscalationId = v
		o.enabledSetters["EscalationId"] = true
	}
}

// UpdateHttpMonitor update an HTTP monitor by its unique ID. Only the
// parameters passed change; the monitor's state, its last check and its open
// alert are kept. Passing `rules` replaces all of them.
func (srv *Monitors) UpdateHttpMonitor(MonitorId string, optionalSetters ...UpdateHttpMonitorOption) (*models.Monitor, error) {
	if MonitorId == "" {
		return nil, errors.New("Missing required parameter: \"monitorId\"")
	}

	r := strings.NewReplacer("{monitorId}", client.EncodePath(MonitorId))
	path := r.Replace("/monitors/http/{monitorId}")
	options := UpdateHttpMonitorOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Url"] {
		params["url"] = options.Url
	}
	if options.enabledSetters["Method"] {
		params["method"] = options.Method
	}
	if options.enabledSetters["Headers"] {
		params["headers"] = options.Headers
	}
	if options.enabledSetters["Body"] {
		params["body"] = options.Body
	}
	if options.enabledSetters["Rules"] {
		params["rules"] = options.Rules
	}
	if options.enabledSetters["Interval"] {
		params["interval"] = options.Interval
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
	}
	if options.enabledSetters["Cooldown"] {
		params["cooldown"] = options.Cooldown
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Notify"] {
		params["notify"] = options.Notify
	}
	if options.enabledSetters["Severity"] {
		params["severity"] = options.Severity
	}
	if options.enabledSetters["EscalationId"] {
		params["escalationId"] = options.EscalationId
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

		parsed := models.Monitor{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Monitor
	parsed, ok := resp.Result.(models.Monitor)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateSmtpMonitorOptions struct {
	TlsMode        string
	Rules          []string
	Interval       int
	Timeout        int
	Cooldown       int
	Enabled        bool
	Notify         []string
	Severity       string
	EscalationId   string
	enabledSetters map[string]bool
}

func (options CreateSmtpMonitorOptions) New() *CreateSmtpMonitorOptions {
	options.enabledSetters = map[string]bool{"TlsMode": false, "Rules": false, "Interval": false, "Timeout": false, "Cooldown": false, "Enabled": false, "Notify": false, "Severity": false, "EscalationId": false}
	return &options
}

type CreateSmtpMonitorOption func(*CreateSmtpMonitorOptions)

func (srv *Monitors) WithCreateSmtpMonitorTlsMode(v string) CreateSmtpMonitorOption {
	return func(o *CreateSmtpMonitorOptions) {
		o.TlsMode = v
		o.enabledSetters["TlsMode"] = true
	}
}
func (srv *Monitors) WithCreateSmtpMonitorRules(v []string) CreateSmtpMonitorOption {
	return func(o *CreateSmtpMonitorOptions) {
		o.Rules = v
		o.enabledSetters["Rules"] = true
	}
}
func (srv *Monitors) WithCreateSmtpMonitorInterval(v int) CreateSmtpMonitorOption {
	return func(o *CreateSmtpMonitorOptions) {
		o.Interval = v
		o.enabledSetters["Interval"] = true
	}
}
func (srv *Monitors) WithCreateSmtpMonitorTimeout(v int) CreateSmtpMonitorOption {
	return func(o *CreateSmtpMonitorOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
func (srv *Monitors) WithCreateSmtpMonitorCooldown(v int) CreateSmtpMonitorOption {
	return func(o *CreateSmtpMonitorOptions) {
		o.Cooldown = v
		o.enabledSetters["Cooldown"] = true
	}
}
func (srv *Monitors) WithCreateSmtpMonitorEnabled(v bool) CreateSmtpMonitorOption {
	return func(o *CreateSmtpMonitorOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Monitors) WithCreateSmtpMonitorNotify(v []string) CreateSmtpMonitorOption {
	return func(o *CreateSmtpMonitorOptions) {
		o.Notify = v
		o.enabledSetters["Notify"] = true
	}
}
func (srv *Monitors) WithCreateSmtpMonitorSeverity(v string) CreateSmtpMonitorOption {
	return func(o *CreateSmtpMonitorOptions) {
		o.Severity = v
		o.enabledSetters["Severity"] = true
	}
}
func (srv *Monitors) WithCreateSmtpMonitorEscalationId(v string) CreateSmtpMonitorOption {
	return func(o *CreateSmtpMonitorOptions) {
		o.EscalationId = v
		o.enabledSetters["EscalationId"] = true
	}
}

// CreateSmtpMonitor create an SMTP monitor. It reads a mail server's
// greeting, sends `EHLO` and, as its TLS mode asks, upgrades with `STARTTLS`
// or speaks TLS from the start, on its own interval. It opens an alert when
// the server cannot be reached, refuses the session or fails the TLS
// handshake, or when its rules match. It never authenticates or sends mail.
// Each alert fires the `monitors.[monitorId].alerts.[alertId].firing` event
// when it opens and `.resolved` when a check passes again, for project
// webhooks, functions and realtime, and notifies the organization's members
// by email and in the console.
func (srv *Monitors) CreateSmtpMonitor(MonitorId string, Name string, Host string, Port int, optionalSetters ...CreateSmtpMonitorOption) (*models.Monitor, error) {
	path := "/monitors/smtp"
	options := CreateSmtpMonitorOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["monitorId"] = MonitorId
	params["name"] = Name
	params["host"] = Host
	params["port"] = Port
	if options.enabledSetters["TlsMode"] {
		params["tlsMode"] = options.TlsMode
	}
	if options.enabledSetters["Rules"] {
		params["rules"] = options.Rules
	}
	if options.enabledSetters["Interval"] {
		params["interval"] = options.Interval
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
	}
	if options.enabledSetters["Cooldown"] {
		params["cooldown"] = options.Cooldown
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Notify"] {
		params["notify"] = options.Notify
	}
	if options.enabledSetters["Severity"] {
		params["severity"] = options.Severity
	}
	if options.enabledSetters["EscalationId"] {
		params["escalationId"] = options.EscalationId
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

		parsed := models.Monitor{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Monitor
	parsed, ok := resp.Result.(models.Monitor)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateSmtpMonitorOptions struct {
	Name           string
	Host           string
	Port           int
	TlsMode        string
	Rules          []string
	Interval       int
	Timeout        int
	Cooldown       int
	Enabled        bool
	Notify         []string
	Severity       string
	EscalationId   string
	enabledSetters map[string]bool
}

func (options UpdateSmtpMonitorOptions) New() *UpdateSmtpMonitorOptions {
	options.enabledSetters = map[string]bool{"Name": false, "Host": false, "Port": false, "TlsMode": false, "Rules": false, "Interval": false, "Timeout": false, "Cooldown": false, "Enabled": false, "Notify": false, "Severity": false, "EscalationId": false}
	return &options
}

type UpdateSmtpMonitorOption func(*UpdateSmtpMonitorOptions)

func (srv *Monitors) WithUpdateSmtpMonitorName(v string) UpdateSmtpMonitorOption {
	return func(o *UpdateSmtpMonitorOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Monitors) WithUpdateSmtpMonitorHost(v string) UpdateSmtpMonitorOption {
	return func(o *UpdateSmtpMonitorOptions) {
		o.Host = v
		o.enabledSetters["Host"] = true
	}
}
func (srv *Monitors) WithUpdateSmtpMonitorPort(v int) UpdateSmtpMonitorOption {
	return func(o *UpdateSmtpMonitorOptions) {
		o.Port = v
		o.enabledSetters["Port"] = true
	}
}
func (srv *Monitors) WithUpdateSmtpMonitorTlsMode(v string) UpdateSmtpMonitorOption {
	return func(o *UpdateSmtpMonitorOptions) {
		o.TlsMode = v
		o.enabledSetters["TlsMode"] = true
	}
}
func (srv *Monitors) WithUpdateSmtpMonitorRules(v []string) UpdateSmtpMonitorOption {
	return func(o *UpdateSmtpMonitorOptions) {
		o.Rules = v
		o.enabledSetters["Rules"] = true
	}
}
func (srv *Monitors) WithUpdateSmtpMonitorInterval(v int) UpdateSmtpMonitorOption {
	return func(o *UpdateSmtpMonitorOptions) {
		o.Interval = v
		o.enabledSetters["Interval"] = true
	}
}
func (srv *Monitors) WithUpdateSmtpMonitorTimeout(v int) UpdateSmtpMonitorOption {
	return func(o *UpdateSmtpMonitorOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
func (srv *Monitors) WithUpdateSmtpMonitorCooldown(v int) UpdateSmtpMonitorOption {
	return func(o *UpdateSmtpMonitorOptions) {
		o.Cooldown = v
		o.enabledSetters["Cooldown"] = true
	}
}
func (srv *Monitors) WithUpdateSmtpMonitorEnabled(v bool) UpdateSmtpMonitorOption {
	return func(o *UpdateSmtpMonitorOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Monitors) WithUpdateSmtpMonitorNotify(v []string) UpdateSmtpMonitorOption {
	return func(o *UpdateSmtpMonitorOptions) {
		o.Notify = v
		o.enabledSetters["Notify"] = true
	}
}
func (srv *Monitors) WithUpdateSmtpMonitorSeverity(v string) UpdateSmtpMonitorOption {
	return func(o *UpdateSmtpMonitorOptions) {
		o.Severity = v
		o.enabledSetters["Severity"] = true
	}
}
func (srv *Monitors) WithUpdateSmtpMonitorEscalationId(v string) UpdateSmtpMonitorOption {
	return func(o *UpdateSmtpMonitorOptions) {
		o.EscalationId = v
		o.enabledSetters["EscalationId"] = true
	}
}

// UpdateSmtpMonitor update an SMTP monitor by its unique ID. Only the
// parameters passed change; the monitor's state, its last check and its open
// alert are kept. Passing `rules` replaces all of them.
func (srv *Monitors) UpdateSmtpMonitor(MonitorId string, optionalSetters ...UpdateSmtpMonitorOption) (*models.Monitor, error) {
	if MonitorId == "" {
		return nil, errors.New("Missing required parameter: \"monitorId\"")
	}

	r := strings.NewReplacer("{monitorId}", client.EncodePath(MonitorId))
	path := r.Replace("/monitors/smtp/{monitorId}")
	options := UpdateSmtpMonitorOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Host"] {
		params["host"] = options.Host
	}
	if options.enabledSetters["Port"] {
		params["port"] = options.Port
	}
	if options.enabledSetters["TlsMode"] {
		params["tlsMode"] = options.TlsMode
	}
	if options.enabledSetters["Rules"] {
		params["rules"] = options.Rules
	}
	if options.enabledSetters["Interval"] {
		params["interval"] = options.Interval
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
	}
	if options.enabledSetters["Cooldown"] {
		params["cooldown"] = options.Cooldown
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Notify"] {
		params["notify"] = options.Notify
	}
	if options.enabledSetters["Severity"] {
		params["severity"] = options.Severity
	}
	if options.enabledSetters["EscalationId"] {
		params["escalationId"] = options.EscalationId
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

		parsed := models.Monitor{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Monitor
	parsed, ok := resp.Result.(models.Monitor)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateTcpMonitorOptions struct {
	Rules          []string
	Interval       int
	Timeout        int
	Cooldown       int
	Enabled        bool
	Notify         []string
	Severity       string
	EscalationId   string
	enabledSetters map[string]bool
}

func (options CreateTcpMonitorOptions) New() *CreateTcpMonitorOptions {
	options.enabledSetters = map[string]bool{"Rules": false, "Interval": false, "Timeout": false, "Cooldown": false, "Enabled": false, "Notify": false, "Severity": false, "EscalationId": false}
	return &options
}

type CreateTcpMonitorOption func(*CreateTcpMonitorOptions)

func (srv *Monitors) WithCreateTcpMonitorRules(v []string) CreateTcpMonitorOption {
	return func(o *CreateTcpMonitorOptions) {
		o.Rules = v
		o.enabledSetters["Rules"] = true
	}
}
func (srv *Monitors) WithCreateTcpMonitorInterval(v int) CreateTcpMonitorOption {
	return func(o *CreateTcpMonitorOptions) {
		o.Interval = v
		o.enabledSetters["Interval"] = true
	}
}
func (srv *Monitors) WithCreateTcpMonitorTimeout(v int) CreateTcpMonitorOption {
	return func(o *CreateTcpMonitorOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
func (srv *Monitors) WithCreateTcpMonitorCooldown(v int) CreateTcpMonitorOption {
	return func(o *CreateTcpMonitorOptions) {
		o.Cooldown = v
		o.enabledSetters["Cooldown"] = true
	}
}
func (srv *Monitors) WithCreateTcpMonitorEnabled(v bool) CreateTcpMonitorOption {
	return func(o *CreateTcpMonitorOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Monitors) WithCreateTcpMonitorNotify(v []string) CreateTcpMonitorOption {
	return func(o *CreateTcpMonitorOptions) {
		o.Notify = v
		o.enabledSetters["Notify"] = true
	}
}
func (srv *Monitors) WithCreateTcpMonitorSeverity(v string) CreateTcpMonitorOption {
	return func(o *CreateTcpMonitorOptions) {
		o.Severity = v
		o.enabledSetters["Severity"] = true
	}
}
func (srv *Monitors) WithCreateTcpMonitorEscalationId(v string) CreateTcpMonitorOption {
	return func(o *CreateTcpMonitorOptions) {
		o.EscalationId = v
		o.enabledSetters["EscalationId"] = true
	}
}

// CreateTcpMonitor create a TCP monitor. It opens a connection to a host and
// port on its own interval and closes it at once, and opens an alert when the
// connection cannot be made, or when its rules match. Each alert fires the
// `monitors.[monitorId].alerts.[alertId].firing` event when it opens and
// `.resolved` when a check passes again, for project webhooks, functions and
// realtime, and notifies the organization's members by email and in the
// console.
func (srv *Monitors) CreateTcpMonitor(MonitorId string, Name string, Host string, Port int, optionalSetters ...CreateTcpMonitorOption) (*models.Monitor, error) {
	path := "/monitors/tcp"
	options := CreateTcpMonitorOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["monitorId"] = MonitorId
	params["name"] = Name
	params["host"] = Host
	params["port"] = Port
	if options.enabledSetters["Rules"] {
		params["rules"] = options.Rules
	}
	if options.enabledSetters["Interval"] {
		params["interval"] = options.Interval
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
	}
	if options.enabledSetters["Cooldown"] {
		params["cooldown"] = options.Cooldown
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Notify"] {
		params["notify"] = options.Notify
	}
	if options.enabledSetters["Severity"] {
		params["severity"] = options.Severity
	}
	if options.enabledSetters["EscalationId"] {
		params["escalationId"] = options.EscalationId
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

		parsed := models.Monitor{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Monitor
	parsed, ok := resp.Result.(models.Monitor)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateTcpMonitorOptions struct {
	Name           string
	Host           string
	Port           int
	Rules          []string
	Interval       int
	Timeout        int
	Cooldown       int
	Enabled        bool
	Notify         []string
	Severity       string
	EscalationId   string
	enabledSetters map[string]bool
}

func (options UpdateTcpMonitorOptions) New() *UpdateTcpMonitorOptions {
	options.enabledSetters = map[string]bool{"Name": false, "Host": false, "Port": false, "Rules": false, "Interval": false, "Timeout": false, "Cooldown": false, "Enabled": false, "Notify": false, "Severity": false, "EscalationId": false}
	return &options
}

type UpdateTcpMonitorOption func(*UpdateTcpMonitorOptions)

func (srv *Monitors) WithUpdateTcpMonitorName(v string) UpdateTcpMonitorOption {
	return func(o *UpdateTcpMonitorOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Monitors) WithUpdateTcpMonitorHost(v string) UpdateTcpMonitorOption {
	return func(o *UpdateTcpMonitorOptions) {
		o.Host = v
		o.enabledSetters["Host"] = true
	}
}
func (srv *Monitors) WithUpdateTcpMonitorPort(v int) UpdateTcpMonitorOption {
	return func(o *UpdateTcpMonitorOptions) {
		o.Port = v
		o.enabledSetters["Port"] = true
	}
}
func (srv *Monitors) WithUpdateTcpMonitorRules(v []string) UpdateTcpMonitorOption {
	return func(o *UpdateTcpMonitorOptions) {
		o.Rules = v
		o.enabledSetters["Rules"] = true
	}
}
func (srv *Monitors) WithUpdateTcpMonitorInterval(v int) UpdateTcpMonitorOption {
	return func(o *UpdateTcpMonitorOptions) {
		o.Interval = v
		o.enabledSetters["Interval"] = true
	}
}
func (srv *Monitors) WithUpdateTcpMonitorTimeout(v int) UpdateTcpMonitorOption {
	return func(o *UpdateTcpMonitorOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
func (srv *Monitors) WithUpdateTcpMonitorCooldown(v int) UpdateTcpMonitorOption {
	return func(o *UpdateTcpMonitorOptions) {
		o.Cooldown = v
		o.enabledSetters["Cooldown"] = true
	}
}
func (srv *Monitors) WithUpdateTcpMonitorEnabled(v bool) UpdateTcpMonitorOption {
	return func(o *UpdateTcpMonitorOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Monitors) WithUpdateTcpMonitorNotify(v []string) UpdateTcpMonitorOption {
	return func(o *UpdateTcpMonitorOptions) {
		o.Notify = v
		o.enabledSetters["Notify"] = true
	}
}
func (srv *Monitors) WithUpdateTcpMonitorSeverity(v string) UpdateTcpMonitorOption {
	return func(o *UpdateTcpMonitorOptions) {
		o.Severity = v
		o.enabledSetters["Severity"] = true
	}
}
func (srv *Monitors) WithUpdateTcpMonitorEscalationId(v string) UpdateTcpMonitorOption {
	return func(o *UpdateTcpMonitorOptions) {
		o.EscalationId = v
		o.enabledSetters["EscalationId"] = true
	}
}

// UpdateTcpMonitor update a TCP monitor by its unique ID. Only the parameters
// passed change; the monitor's state, its last check and its open alert are
// kept. Passing `rules` replaces all of them.
func (srv *Monitors) UpdateTcpMonitor(MonitorId string, optionalSetters ...UpdateTcpMonitorOption) (*models.Monitor, error) {
	if MonitorId == "" {
		return nil, errors.New("Missing required parameter: \"monitorId\"")
	}

	r := strings.NewReplacer("{monitorId}", client.EncodePath(MonitorId))
	path := r.Replace("/monitors/tcp/{monitorId}")
	options := UpdateTcpMonitorOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Host"] {
		params["host"] = options.Host
	}
	if options.enabledSetters["Port"] {
		params["port"] = options.Port
	}
	if options.enabledSetters["Rules"] {
		params["rules"] = options.Rules
	}
	if options.enabledSetters["Interval"] {
		params["interval"] = options.Interval
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
	}
	if options.enabledSetters["Cooldown"] {
		params["cooldown"] = options.Cooldown
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Notify"] {
		params["notify"] = options.Notify
	}
	if options.enabledSetters["Severity"] {
		params["severity"] = options.Severity
	}
	if options.enabledSetters["EscalationId"] {
		params["escalationId"] = options.EscalationId
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

		parsed := models.Monitor{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Monitor
	parsed, ok := resp.Result.(models.Monitor)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateUsageMonitorOptions struct {
	Window         int
	Cooldown       int
	Enabled        bool
	Notify         []string
	Severity       string
	EscalationId   string
	enabledSetters map[string]bool
}

func (options CreateUsageMonitorOptions) New() *CreateUsageMonitorOptions {
	options.enabledSetters = map[string]bool{"Window": false, "Cooldown": false, "Enabled": false, "Notify": false, "Severity": false, "EscalationId": false}
	return &options
}

type CreateUsageMonitorOption func(*CreateUsageMonitorOptions)

func (srv *Monitors) WithCreateUsageMonitorWindow(v int) CreateUsageMonitorOption {
	return func(o *CreateUsageMonitorOptions) {
		o.Window = v
		o.enabledSetters["Window"] = true
	}
}
func (srv *Monitors) WithCreateUsageMonitorCooldown(v int) CreateUsageMonitorOption {
	return func(o *CreateUsageMonitorOptions) {
		o.Cooldown = v
		o.enabledSetters["Cooldown"] = true
	}
}
func (srv *Monitors) WithCreateUsageMonitorEnabled(v bool) CreateUsageMonitorOption {
	return func(o *CreateUsageMonitorOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Monitors) WithCreateUsageMonitorNotify(v []string) CreateUsageMonitorOption {
	return func(o *CreateUsageMonitorOptions) {
		o.Notify = v
		o.enabledSetters["Notify"] = true
	}
}
func (srv *Monitors) WithCreateUsageMonitorSeverity(v string) CreateUsageMonitorOption {
	return func(o *CreateUsageMonitorOptions) {
		o.Severity = v
		o.enabledSetters["Severity"] = true
	}
}
func (srv *Monitors) WithCreateUsageMonitorEscalationId(v string) CreateUsageMonitorOption {
	return func(o *CreateUsageMonitorOptions) {
		o.EscalationId = v
		o.enabledSetters["EscalationId"] = true
	}
}

// CreateUsageMonitor create a usage monitor. It reads a usage metric of the
// project, or of one of its databases, buckets, functions or sites, on a
// schedule, and opens an alert when all its rules hold. Each alert fires the
// `monitors.[monitorId].alerts.[alertId].firing` event when it opens and
// `.resolved` when usage recovers, for project webhooks, functions and
// realtime, and notifies the confirmed members of the `notify` roles by email
// and in the console, or the steps of the escalation set by `escalationId`.
func (srv *Monitors) CreateUsageMonitor(MonitorId string, Name string, Rules []string, optionalSetters ...CreateUsageMonitorOption) (*models.Monitor, error) {
	path := "/monitors/usage"
	options := CreateUsageMonitorOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["monitorId"] = MonitorId
	params["name"] = Name
	params["rules"] = Rules
	if options.enabledSetters["Window"] {
		params["window"] = options.Window
	}
	if options.enabledSetters["Cooldown"] {
		params["cooldown"] = options.Cooldown
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Notify"] {
		params["notify"] = options.Notify
	}
	if options.enabledSetters["Severity"] {
		params["severity"] = options.Severity
	}
	if options.enabledSetters["EscalationId"] {
		params["escalationId"] = options.EscalationId
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

		parsed := models.Monitor{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Monitor
	parsed, ok := resp.Result.(models.Monitor)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateUsageMonitorOptions struct {
	Name           string
	Rules          []string
	Window         int
	Cooldown       int
	Enabled        bool
	Notify         []string
	Severity       string
	EscalationId   string
	enabledSetters map[string]bool
}

func (options UpdateUsageMonitorOptions) New() *UpdateUsageMonitorOptions {
	options.enabledSetters = map[string]bool{"Name": false, "Rules": false, "Window": false, "Cooldown": false, "Enabled": false, "Notify": false, "Severity": false, "EscalationId": false}
	return &options
}

type UpdateUsageMonitorOption func(*UpdateUsageMonitorOptions)

func (srv *Monitors) WithUpdateUsageMonitorName(v string) UpdateUsageMonitorOption {
	return func(o *UpdateUsageMonitorOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Monitors) WithUpdateUsageMonitorRules(v []string) UpdateUsageMonitorOption {
	return func(o *UpdateUsageMonitorOptions) {
		o.Rules = v
		o.enabledSetters["Rules"] = true
	}
}
func (srv *Monitors) WithUpdateUsageMonitorWindow(v int) UpdateUsageMonitorOption {
	return func(o *UpdateUsageMonitorOptions) {
		o.Window = v
		o.enabledSetters["Window"] = true
	}
}
func (srv *Monitors) WithUpdateUsageMonitorCooldown(v int) UpdateUsageMonitorOption {
	return func(o *UpdateUsageMonitorOptions) {
		o.Cooldown = v
		o.enabledSetters["Cooldown"] = true
	}
}
func (srv *Monitors) WithUpdateUsageMonitorEnabled(v bool) UpdateUsageMonitorOption {
	return func(o *UpdateUsageMonitorOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Monitors) WithUpdateUsageMonitorNotify(v []string) UpdateUsageMonitorOption {
	return func(o *UpdateUsageMonitorOptions) {
		o.Notify = v
		o.enabledSetters["Notify"] = true
	}
}
func (srv *Monitors) WithUpdateUsageMonitorSeverity(v string) UpdateUsageMonitorOption {
	return func(o *UpdateUsageMonitorOptions) {
		o.Severity = v
		o.enabledSetters["Severity"] = true
	}
}
func (srv *Monitors) WithUpdateUsageMonitorEscalationId(v string) UpdateUsageMonitorOption {
	return func(o *UpdateUsageMonitorOptions) {
		o.EscalationId = v
		o.enabledSetters["EscalationId"] = true
	}
}

// UpdateUsageMonitor update a usage monitor by its unique ID. Only the
// parameters passed change; the monitor's state and its open alert are kept.
// Passing `rules` replaces all of them.
func (srv *Monitors) UpdateUsageMonitor(MonitorId string, optionalSetters ...UpdateUsageMonitorOption) (*models.Monitor, error) {
	if MonitorId == "" {
		return nil, errors.New("Missing required parameter: \"monitorId\"")
	}

	r := strings.NewReplacer("{monitorId}", client.EncodePath(MonitorId))
	path := r.Replace("/monitors/usage/{monitorId}")
	options := UpdateUsageMonitorOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Rules"] {
		params["rules"] = options.Rules
	}
	if options.enabledSetters["Window"] {
		params["window"] = options.Window
	}
	if options.enabledSetters["Cooldown"] {
		params["cooldown"] = options.Cooldown
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Notify"] {
		params["notify"] = options.Notify
	}
	if options.enabledSetters["Severity"] {
		params["severity"] = options.Severity
	}
	if options.enabledSetters["EscalationId"] {
		params["escalationId"] = options.EscalationId
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

		parsed := models.Monitor{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Monitor
	parsed, ok := resp.Result.(models.Monitor)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Get get a monitor by its unique ID, with its state as of its last
// evaluation.
func (srv *Monitors) Get(MonitorId string) (*models.Monitor, error) {
	if MonitorId == "" {
		return nil, errors.New("Missing required parameter: \"monitorId\"")
	}

	r := strings.NewReplacer("{monitorId}", client.EncodePath(MonitorId))
	path := r.Replace("/monitors/{monitorId}")
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

		parsed := models.Monitor{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Monitor
	parsed, ok := resp.Result.(models.Monitor)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Delete delete a monitor by its unique ID. It stops being evaluated; an
// alert it is firing is resolved and kept with the rest of its alert history.
func (srv *Monitors) Delete(MonitorId string) (*interface{}, error) {
	if MonitorId == "" {
		return nil, errors.New("Missing required parameter: \"monitorId\"")
	}

	r := strings.NewReplacer("{monitorId}", client.EncodePath(MonitorId))
	path := r.Replace("/monitors/{monitorId}")
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
