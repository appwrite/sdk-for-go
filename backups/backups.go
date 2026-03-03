package backups

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// Backups service
type Backups struct {
	client client.Client
}

func New(clt client.Client) *Backups {
	return &Backups{
		client: clt,
	}
}

type ListArchivesOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListArchivesOptions) New() *ListArchivesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListArchivesOption func(*ListArchivesOptions)
func (srv *Backups) WithListArchivesQueries(v []string) ListArchivesOption {
	return func(o *ListArchivesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
	
// ListArchives list all archives for a project.
func (srv *Backups) ListArchives(optionalSetters ...ListArchivesOption)(*models.BackupArchiveList, error) {
	path := "/backups/archives"
	options := ListArchivesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.BackupArchiveList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BackupArchiveList
	parsed, ok := resp.Result.(models.BackupArchiveList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateArchiveOptions struct {
	ResourceId string
	enabledSetters map[string]bool
}
func (options CreateArchiveOptions) New() *CreateArchiveOptions {
	options.enabledSetters = map[string]bool{
		"ResourceId": false,
	}
	return &options
}
type CreateArchiveOption func(*CreateArchiveOptions)
func (srv *Backups) WithCreateArchiveResourceId(v string) CreateArchiveOption {
	return func(o *CreateArchiveOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
			
// CreateArchive create a new archive asynchronously for a project.
func (srv *Backups) CreateArchive(Services []string, optionalSetters ...CreateArchiveOption)(*models.BackupArchive, error) {
	path := "/backups/archives"
	options := CreateArchiveOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["services"] = Services
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
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

		parsed := models.BackupArchive{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BackupArchive
	parsed, ok := resp.Result.(models.BackupArchive)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetArchive get a backup archive using it's ID.
func (srv *Backups) GetArchive(ArchiveId string)(*models.BackupArchive, error) {
	r := strings.NewReplacer("{archiveId}", ArchiveId)
	path := r.Replace("/backups/archives/{archiveId}")
	params := map[string]interface{}{}
	params["archiveId"] = ArchiveId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.BackupArchive{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BackupArchive
	parsed, ok := resp.Result.(models.BackupArchive)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteArchive delete an existing archive for a project.
func (srv *Backups) DeleteArchive(ArchiveId string)(*interface{}, error) {
	r := strings.NewReplacer("{archiveId}", ArchiveId)
	path := r.Replace("/backups/archives/{archiveId}")
	params := map[string]interface{}{}
	params["archiveId"] = ArchiveId
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
type ListPoliciesOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListPoliciesOptions) New() *ListPoliciesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListPoliciesOption func(*ListPoliciesOptions)
func (srv *Backups) WithListPoliciesQueries(v []string) ListPoliciesOption {
	return func(o *ListPoliciesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
	
// ListPolicies list all policies for a project.
func (srv *Backups) ListPolicies(optionalSetters ...ListPoliciesOption)(*models.BackupPolicyList, error) {
	path := "/backups/policies"
	options := ListPoliciesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.BackupPolicyList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BackupPolicyList
	parsed, ok := resp.Result.(models.BackupPolicyList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreatePolicyOptions struct {
	Name string
	ResourceId string
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreatePolicyOptions) New() *CreatePolicyOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"ResourceId": false,
		"Enabled": false,
	}
	return &options
}
type CreatePolicyOption func(*CreatePolicyOptions)
func (srv *Backups) WithCreatePolicyName(v string) CreatePolicyOption {
	return func(o *CreatePolicyOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Backups) WithCreatePolicyResourceId(v string) CreatePolicyOption {
	return func(o *CreatePolicyOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
func (srv *Backups) WithCreatePolicyEnabled(v bool) CreatePolicyOption {
	return func(o *CreatePolicyOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
									
// CreatePolicy create a new backup policy.
func (srv *Backups) CreatePolicy(PolicyId string, Services []string, Retention int, Schedule string, optionalSetters ...CreatePolicyOption)(*models.BackupPolicy, error) {
	path := "/backups/policies"
	options := CreatePolicyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["policyId"] = PolicyId
	params["services"] = Services
	params["retention"] = Retention
	params["schedule"] = Schedule
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.BackupPolicy{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BackupPolicy
	parsed, ok := resp.Result.(models.BackupPolicy)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetPolicy get a backup policy using it's ID.
func (srv *Backups) GetPolicy(PolicyId string)(*models.BackupPolicy, error) {
	r := strings.NewReplacer("{policyId}", PolicyId)
	path := r.Replace("/backups/policies/{policyId}")
	params := map[string]interface{}{}
	params["policyId"] = PolicyId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.BackupPolicy{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BackupPolicy
	parsed, ok := resp.Result.(models.BackupPolicy)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdatePolicyOptions struct {
	Name string
	Retention int
	Schedule string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdatePolicyOptions) New() *UpdatePolicyOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Retention": false,
		"Schedule": false,
		"Enabled": false,
	}
	return &options
}
type UpdatePolicyOption func(*UpdatePolicyOptions)
func (srv *Backups) WithUpdatePolicyName(v string) UpdatePolicyOption {
	return func(o *UpdatePolicyOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Backups) WithUpdatePolicyRetention(v int) UpdatePolicyOption {
	return func(o *UpdatePolicyOptions) {
		o.Retention = v
		o.enabledSetters["Retention"] = true
	}
}
func (srv *Backups) WithUpdatePolicySchedule(v string) UpdatePolicyOption {
	return func(o *UpdatePolicyOptions) {
		o.Schedule = v
		o.enabledSetters["Schedule"] = true
	}
}
func (srv *Backups) WithUpdatePolicyEnabled(v bool) UpdatePolicyOption {
	return func(o *UpdatePolicyOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
			
// UpdatePolicy update an existing policy using it's ID.
func (srv *Backups) UpdatePolicy(PolicyId string, optionalSetters ...UpdatePolicyOption)(*models.BackupPolicy, error) {
	r := strings.NewReplacer("{policyId}", PolicyId)
	path := r.Replace("/backups/policies/{policyId}")
	options := UpdatePolicyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["policyId"] = PolicyId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Retention"] {
		params["retention"] = options.Retention
	}
	if options.enabledSetters["Schedule"] {
		params["schedule"] = options.Schedule
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.BackupPolicy{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BackupPolicy
	parsed, ok := resp.Result.(models.BackupPolicy)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeletePolicy delete a policy using it's ID.
func (srv *Backups) DeletePolicy(PolicyId string)(*interface{}, error) {
	r := strings.NewReplacer("{policyId}", PolicyId)
	path := r.Replace("/backups/policies/{policyId}")
	params := map[string]interface{}{}
	params["policyId"] = PolicyId
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
type CreateRestorationOptions struct {
	NewResourceId string
	NewResourceName string
	enabledSetters map[string]bool
}
func (options CreateRestorationOptions) New() *CreateRestorationOptions {
	options.enabledSetters = map[string]bool{
		"NewResourceId": false,
		"NewResourceName": false,
	}
	return &options
}
type CreateRestorationOption func(*CreateRestorationOptions)
func (srv *Backups) WithCreateRestorationNewResourceId(v string) CreateRestorationOption {
	return func(o *CreateRestorationOptions) {
		o.NewResourceId = v
		o.enabledSetters["NewResourceId"] = true
	}
}
func (srv *Backups) WithCreateRestorationNewResourceName(v string) CreateRestorationOption {
	return func(o *CreateRestorationOptions) {
		o.NewResourceName = v
		o.enabledSetters["NewResourceName"] = true
	}
}
					
// CreateRestoration create and trigger a new restoration for a backup on a
// project.
func (srv *Backups) CreateRestoration(ArchiveId string, Services []string, optionalSetters ...CreateRestorationOption)(*models.BackupRestoration, error) {
	path := "/backups/restoration"
	options := CreateRestorationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["archiveId"] = ArchiveId
	params["services"] = Services
	if options.enabledSetters["NewResourceId"] {
		params["newResourceId"] = options.NewResourceId
	}
	if options.enabledSetters["NewResourceName"] {
		params["newResourceName"] = options.NewResourceName
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

		parsed := models.BackupRestoration{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BackupRestoration
	parsed, ok := resp.Result.(models.BackupRestoration)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListRestorationsOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListRestorationsOptions) New() *ListRestorationsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListRestorationsOption func(*ListRestorationsOptions)
func (srv *Backups) WithListRestorationsQueries(v []string) ListRestorationsOption {
	return func(o *ListRestorationsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
	
// ListRestorations list all backup restorations for a project.
func (srv *Backups) ListRestorations(optionalSetters ...ListRestorationsOption)(*models.BackupRestorationList, error) {
	path := "/backups/restorations"
	options := ListRestorationsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.BackupRestorationList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BackupRestorationList
	parsed, ok := resp.Result.(models.BackupRestorationList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetRestoration get the current status of a backup restoration.
func (srv *Backups) GetRestoration(RestorationId string)(*models.BackupRestoration, error) {
	r := strings.NewReplacer("{restorationId}", RestorationId)
	path := r.Replace("/backups/restorations/{restorationId}")
	params := map[string]interface{}{}
	params["restorationId"] = RestorationId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.BackupRestoration{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.BackupRestoration
	parsed, ok := resp.Result.(models.BackupRestoration)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
