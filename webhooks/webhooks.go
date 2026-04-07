package webhooks

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/v2/client"
	"github.com/appwrite/sdk-for-go/v2/models"
	"strings"
)

// Webhooks service
type Webhooks struct {
	client client.Client
}

func New(clt client.Client) *Webhooks {
	return &Webhooks{
		client: clt,
	}
}

type ListOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListOption func(*ListOptions)
func (srv *Webhooks) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Webhooks) WithListTotal(v bool) ListOption {
	return func(o *ListOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// List get a list of all webhooks belonging to the project. You can use the
// query params to filter your results.
func (srv *Webhooks) List(optionalSetters ...ListOption)(*models.WebhookList, error) {
	path := "/webhooks"
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
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.WebhookList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.WebhookList
	parsed, ok := resp.Result.(models.WebhookList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateOptions struct {
	Enabled bool
	Security bool
	HttpUser string
	HttpPass string
	enabledSetters map[string]bool
}
func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{
		"Enabled": false,
		"Security": false,
		"HttpUser": false,
		"HttpPass": false,
	}
	return &options
}
type CreateOption func(*CreateOptions)
func (srv *Webhooks) WithCreateEnabled(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Webhooks) WithCreateSecurity(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.Security = v
		o.enabledSetters["Security"] = true
	}
}
func (srv *Webhooks) WithCreateHttpUser(v string) CreateOption {
	return func(o *CreateOptions) {
		o.HttpUser = v
		o.enabledSetters["HttpUser"] = true
	}
}
func (srv *Webhooks) WithCreateHttpPass(v string) CreateOption {
	return func(o *CreateOptions) {
		o.HttpPass = v
		o.enabledSetters["HttpPass"] = true
	}
}
									
// Create create a new webhook. Use this endpoint to configure a URL that will
// receive events from Appwrite when specific events occur.
func (srv *Webhooks) Create(WebhookId string, Url string, Name string, Events []string, optionalSetters ...CreateOption)(*models.Webhook, error) {
	path := "/webhooks"
	options := CreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["webhookId"] = WebhookId
	params["url"] = Url
	params["name"] = Name
	params["events"] = Events
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Security"] {
		params["security"] = options.Security
	}
	if options.enabledSetters["HttpUser"] {
		params["httpUser"] = options.HttpUser
	}
	if options.enabledSetters["HttpPass"] {
		params["httpPass"] = options.HttpPass
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Webhook{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Webhook
	parsed, ok := resp.Result.(models.Webhook)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Get get a webhook by its unique ID. This endpoint returns details about a
// specific webhook configured for a project.
func (srv *Webhooks) Get(WebhookId string)(*models.Webhook, error) {
	r := strings.NewReplacer("{webhookId}", WebhookId)
	path := r.Replace("/webhooks/{webhookId}")
	params := map[string]interface{}{}
	params["webhookId"] = WebhookId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Webhook{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Webhook
	parsed, ok := resp.Result.(models.Webhook)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOptions struct {
	Enabled bool
	Security bool
	HttpUser string
	HttpPass string
	enabledSetters map[string]bool
}
func (options UpdateOptions) New() *UpdateOptions {
	options.enabledSetters = map[string]bool{
		"Enabled": false,
		"Security": false,
		"HttpUser": false,
		"HttpPass": false,
	}
	return &options
}
type UpdateOption func(*UpdateOptions)
func (srv *Webhooks) WithUpdateEnabled(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Webhooks) WithUpdateSecurity(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Security = v
		o.enabledSetters["Security"] = true
	}
}
func (srv *Webhooks) WithUpdateHttpUser(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.HttpUser = v
		o.enabledSetters["HttpUser"] = true
	}
}
func (srv *Webhooks) WithUpdateHttpPass(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.HttpPass = v
		o.enabledSetters["HttpPass"] = true
	}
}
									
// Update update a webhook by its unique ID. Use this endpoint to update the
// URL, events, or status of an existing webhook.
func (srv *Webhooks) Update(WebhookId string, Name string, Url string, Events []string, optionalSetters ...UpdateOption)(*models.Webhook, error) {
	r := strings.NewReplacer("{webhookId}", WebhookId)
	path := r.Replace("/webhooks/{webhookId}")
	options := UpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["webhookId"] = WebhookId
	params["name"] = Name
	params["url"] = Url
	params["events"] = Events
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Security"] {
		params["security"] = options.Security
	}
	if options.enabledSetters["HttpUser"] {
		params["httpUser"] = options.HttpUser
	}
	if options.enabledSetters["HttpPass"] {
		params["httpPass"] = options.HttpPass
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Webhook{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Webhook
	parsed, ok := resp.Result.(models.Webhook)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Delete delete a webhook by its unique ID. Once deleted, the webhook will no
// longer receive project events.
func (srv *Webhooks) Delete(WebhookId string)(*interface{}, error) {
	r := strings.NewReplacer("{webhookId}", WebhookId)
	path := r.Replace("/webhooks/{webhookId}")
	params := map[string]interface{}{}
	params["webhookId"] = WebhookId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

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
	
// UpdateSignature update the webhook signature key. This endpoint can be used
// to regenerate the signature key used to sign and validate payload
// deliveries for a specific webhook.
func (srv *Webhooks) UpdateSignature(WebhookId string)(*models.Webhook, error) {
	r := strings.NewReplacer("{webhookId}", WebhookId)
	path := r.Replace("/webhooks/{webhookId}/signature")
	params := map[string]interface{}{}
	params["webhookId"] = WebhookId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Webhook{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Webhook
	parsed, ok := resp.Result.(models.Webhook)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
