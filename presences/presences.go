package presences

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/v4/client"
	"github.com/appwrite/sdk-for-go/v4/models"
	"strings"
)

// Presences service
type Presences struct {
	client client.Client
}

func New(clt client.Client) *Presences {
	return &Presences{
		client: clt,
	}
}

type ListOptions struct {
	Queries []string
	Total bool
	Ttl int
	enabledSetters map[string]bool
}
func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
		"Ttl": false,
	}
	return &options
}
type ListOption func(*ListOptions)
func (srv *Presences) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Presences) WithListTotal(v bool) ListOption {
	return func(o *ListOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
func (srv *Presences) WithListTtl(v int) ListOption {
	return func(o *ListOptions) {
		o.Ttl = v
		o.enabledSetters["Ttl"] = true
	}
}
	
// List list presence logs. Expired entries are filtered out automatically.
func (srv *Presences) List(optionalSetters ...ListOption)(*models.PresenceList, error) {
	path := "/presences"
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
	if options.enabledSetters["Ttl"] {
		params["ttl"] = options.Ttl
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PresenceList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PresenceList
	parsed, ok := resp.Result.(models.PresenceList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Get get a presence log by its unique ID. Entries whose `expiresAt` is in
// the past are treated as not found.
func (srv *Presences) Get(PresenceId string)(*models.Presence, error) {
	r := strings.NewReplacer("{presenceId}", PresenceId)
	path := r.Replace("/presences/{presenceId}")
	params := map[string]interface{}{}
	params["presenceId"] = PresenceId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Presence{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Presence
	parsed, ok := resp.Result.(models.Presence)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpsertOptions struct {
	Permissions []string
	ExpiresAt string
	Metadata interface{}
	enabledSetters map[string]bool
}
func (options UpsertOptions) New() *UpsertOptions {
	options.enabledSetters = map[string]bool{
		"Permissions": false,
		"ExpiresAt": false,
		"Metadata": false,
	}
	return &options
}
type UpsertOption func(*UpsertOptions)
func (srv *Presences) WithUpsertPermissions(v []string) UpsertOption {
	return func(o *UpsertOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *Presences) WithUpsertExpiresAt(v string) UpsertOption {
	return func(o *UpsertOptions) {
		o.ExpiresAt = v
		o.enabledSetters["ExpiresAt"] = true
	}
}
func (srv *Presences) WithUpsertMetadata(v interface{}) UpsertOption {
	return func(o *UpsertOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
							
// Upsert create or update a presence log by its user ID.
func (srv *Presences) Upsert(PresenceId string, UserId string, Status string, optionalSetters ...UpsertOption)(*models.Presence, error) {
	r := strings.NewReplacer("{presenceId}", PresenceId)
	path := r.Replace("/presences/{presenceId}")
	options := UpsertOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["presenceId"] = PresenceId
	params["userId"] = UserId
	params["status"] = Status
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["ExpiresAt"] {
		params["expiresAt"] = options.ExpiresAt
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
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

		parsed := models.Presence{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Presence
	parsed, ok := resp.Result.(models.Presence)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdatePresenceOptions struct {
	Status string
	ExpiresAt string
	Metadata interface{}
	Permissions []string
	Purge bool
	enabledSetters map[string]bool
}
func (options UpdatePresenceOptions) New() *UpdatePresenceOptions {
	options.enabledSetters = map[string]bool{
		"Status": false,
		"ExpiresAt": false,
		"Metadata": false,
		"Permissions": false,
		"Purge": false,
	}
	return &options
}
type UpdatePresenceOption func(*UpdatePresenceOptions)
func (srv *Presences) WithUpdatePresenceStatus(v string) UpdatePresenceOption {
	return func(o *UpdatePresenceOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Presences) WithUpdatePresenceExpiresAt(v string) UpdatePresenceOption {
	return func(o *UpdatePresenceOptions) {
		o.ExpiresAt = v
		o.enabledSetters["ExpiresAt"] = true
	}
}
func (srv *Presences) WithUpdatePresenceMetadata(v interface{}) UpdatePresenceOption {
	return func(o *UpdatePresenceOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Presences) WithUpdatePresencePermissions(v []string) UpdatePresenceOption {
	return func(o *UpdatePresenceOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *Presences) WithUpdatePresencePurge(v bool) UpdatePresenceOption {
	return func(o *UpdatePresenceOptions) {
		o.Purge = v
		o.enabledSetters["Purge"] = true
	}
}
					
// UpdatePresence update a presence log by its unique ID. Using the patch
// method you can pass only specific fields that will get updated.
func (srv *Presences) UpdatePresence(PresenceId string, UserId string, optionalSetters ...UpdatePresenceOption)(*models.Presence, error) {
	r := strings.NewReplacer("{presenceId}", PresenceId)
	path := r.Replace("/presences/{presenceId}")
	options := UpdatePresenceOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["presenceId"] = PresenceId
	params["userId"] = UserId
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["ExpiresAt"] {
		params["expiresAt"] = options.ExpiresAt
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["Purge"] {
		params["purge"] = options.Purge
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Presence{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Presence
	parsed, ok := resp.Result.(models.Presence)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Delete delete a presence log by its unique ID.
func (srv *Presences) Delete(PresenceId string)(*interface{}, error) {
	r := strings.NewReplacer("{presenceId}", PresenceId)
	path := r.Replace("/presences/{presenceId}")
	params := map[string]interface{}{}
	params["presenceId"] = PresenceId
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
