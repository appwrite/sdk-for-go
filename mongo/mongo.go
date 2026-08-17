package mongo

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/models"
)

// Mongo service
type Mongo struct {
	client client.Client
}

func New(clt client.Client) *Mongo {
	return &Mongo{
		client: clt,
	}
}

type ListOptions struct {
	Queries        []string
	enabledSetters map[string]bool
}

func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{"Queries": false}
	return &options
}

type ListOption func(*ListOptions)

func (srv *Mongo) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// List list all dedicated databases. Results support pagination.
func (srv *Mongo) List(optionalSetters ...ListOption) (*models.DedicatedDatabaseList, error) {
	path := "/mongo"
	options := ListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
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

		parsed := models.DedicatedDatabaseList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabaseList
	parsed, ok := resp.Result.(models.DedicatedDatabaseList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateOptions struct {
	Version                            string
	Specification                      string
	Replicas                           int
	SyncMode                           string
	NetworkIdleTimeoutSeconds          int
	NetworkIPAllowlist                 []string
	IdleTimeoutMinutes                 int
	Pitr                               bool
	PitrRetentionDays                  int
	StorageAutoscaling                 bool
	StorageAutoscalingThresholdPercent int
	StorageAutoscalingMaxGb            int
	enabledSetters                     map[string]bool
}

func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{"Version": false, "Specification": false, "Replicas": false, "SyncMode": false, "NetworkIdleTimeoutSeconds": false, "NetworkIPAllowlist": false, "IdleTimeoutMinutes": false, "Pitr": false, "PitrRetentionDays": false, "StorageAutoscaling": false, "StorageAutoscalingThresholdPercent": false, "StorageAutoscalingMaxGb": false}
	return &options
}

type CreateOption func(*CreateOptions)

func (srv *Mongo) WithCreateVersion(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Version = v
		o.enabledSetters["Version"] = true
	}
}
func (srv *Mongo) WithCreateSpecification(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Specification = v
		o.enabledSetters["Specification"] = true
	}
}
func (srv *Mongo) WithCreateReplicas(v int) CreateOption {
	return func(o *CreateOptions) {
		o.Replicas = v
		o.enabledSetters["Replicas"] = true
	}
}
func (srv *Mongo) WithCreateSyncMode(v string) CreateOption {
	return func(o *CreateOptions) {
		o.SyncMode = v
		o.enabledSetters["SyncMode"] = true
	}
}
func (srv *Mongo) WithCreateNetworkIdleTimeoutSeconds(v int) CreateOption {
	return func(o *CreateOptions) {
		o.NetworkIdleTimeoutSeconds = v
		o.enabledSetters["NetworkIdleTimeoutSeconds"] = true
	}
}
func (srv *Mongo) WithCreateNetworkIPAllowlist(v []string) CreateOption {
	return func(o *CreateOptions) {
		o.NetworkIPAllowlist = v
		o.enabledSetters["NetworkIPAllowlist"] = true
	}
}
func (srv *Mongo) WithCreateIdleTimeoutMinutes(v int) CreateOption {
	return func(o *CreateOptions) {
		o.IdleTimeoutMinutes = v
		o.enabledSetters["IdleTimeoutMinutes"] = true
	}
}
func (srv *Mongo) WithCreatePitr(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.Pitr = v
		o.enabledSetters["Pitr"] = true
	}
}
func (srv *Mongo) WithCreatePitrRetentionDays(v int) CreateOption {
	return func(o *CreateOptions) {
		o.PitrRetentionDays = v
		o.enabledSetters["PitrRetentionDays"] = true
	}
}
func (srv *Mongo) WithCreateStorageAutoscaling(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.StorageAutoscaling = v
		o.enabledSetters["StorageAutoscaling"] = true
	}
}
func (srv *Mongo) WithCreateStorageAutoscalingThresholdPercent(v int) CreateOption {
	return func(o *CreateOptions) {
		o.StorageAutoscalingThresholdPercent = v
		o.enabledSetters["StorageAutoscalingThresholdPercent"] = true
	}
}
func (srv *Mongo) WithCreateStorageAutoscalingMaxGb(v int) CreateOption {
	return func(o *CreateOptions) {
		o.StorageAutoscalingMaxGb = v
		o.enabledSetters["StorageAutoscalingMaxGb"] = true
	}
}

// Create create a new dedicated database with the chosen engine and
// configuration. Status will be 'provisioning' until the database is ready.
func (srv *Mongo) Create(DatabaseId string, Name string, optionalSetters ...CreateOption) (*models.DedicatedDatabase, error) {
	path := "/mongo"
	options := CreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["name"] = Name
	if options.enabledSetters["Version"] {
		params["version"] = options.Version
	}
	if options.enabledSetters["Specification"] {
		params["specification"] = options.Specification
	}
	if options.enabledSetters["Replicas"] {
		params["replicas"] = options.Replicas
	}
	if options.enabledSetters["SyncMode"] {
		params["syncMode"] = options.SyncMode
	}
	if options.enabledSetters["NetworkIdleTimeoutSeconds"] {
		params["networkIdleTimeoutSeconds"] = options.NetworkIdleTimeoutSeconds
	}
	if options.enabledSetters["NetworkIPAllowlist"] {
		params["networkIPAllowlist"] = options.NetworkIPAllowlist
	}
	if options.enabledSetters["IdleTimeoutMinutes"] {
		params["idleTimeoutMinutes"] = options.IdleTimeoutMinutes
	}
	if options.enabledSetters["Pitr"] {
		params["pitr"] = options.Pitr
	}
	if options.enabledSetters["PitrRetentionDays"] {
		params["pitrRetentionDays"] = options.PitrRetentionDays
	}
	if options.enabledSetters["StorageAutoscaling"] {
		params["storageAutoscaling"] = options.StorageAutoscaling
	}
	if options.enabledSetters["StorageAutoscalingThresholdPercent"] {
		params["storageAutoscalingThresholdPercent"] = options.StorageAutoscalingThresholdPercent
	}
	if options.enabledSetters["StorageAutoscalingMaxGb"] {
		params["storageAutoscalingMaxGb"] = options.StorageAutoscalingMaxGb
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

		parsed := models.DedicatedDatabase{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabase
	parsed, ok := resp.Result.(models.DedicatedDatabase)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListSpecifications list the dedicated database specifications available on
// the current plan. Each specification reports its resource limits, pricing,
// and whether it is enabled for the organization.
func (srv *Mongo) ListSpecifications() (*models.DedicatedDatabaseSpecificationList, error) {
	path := "/mongo/specifications"
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

		parsed := models.DedicatedDatabaseSpecificationList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabaseSpecificationList
	parsed, ok := resp.Result.(models.DedicatedDatabaseSpecificationList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Get get a dedicated database by its unique ID. Returns the database
// configuration and current status.
func (srv *Mongo) Get(DatabaseId string) (*models.DedicatedDatabase, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}")
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

		parsed := models.DedicatedDatabase{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabase
	parsed, ok := resp.Result.(models.DedicatedDatabase)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateOptions struct {
	Name                               string
	Status                             string
	Specification                      string
	Replicas                           int
	SyncMode                           string
	NetworkIdleTimeoutSeconds          int
	NetworkIPAllowlist                 []string
	IdleTimeoutMinutes                 int
	Pitr                               bool
	PitrRetentionDays                  int
	StorageAutoscaling                 bool
	StorageAutoscalingThresholdPercent int
	StorageAutoscalingMaxGb            int
	MetricsTraceSampleRate             float64
	MetricsSlowQueryLogThresholdMs     int
	SqlApiEnabled                      bool
	SqlApiAllowedStatements            []string
	SqlApiMaxRows                      int
	SqlApiMaxBytes                     int
	SqlApiTimeoutSeconds               int
	enabledSetters                     map[string]bool
}

func (options UpdateOptions) New() *UpdateOptions {
	options.enabledSetters = map[string]bool{"Name": false, "Status": false, "Specification": false, "Replicas": false, "SyncMode": false, "NetworkIdleTimeoutSeconds": false, "NetworkIPAllowlist": false, "IdleTimeoutMinutes": false, "Pitr": false, "PitrRetentionDays": false, "StorageAutoscaling": false, "StorageAutoscalingThresholdPercent": false, "StorageAutoscalingMaxGb": false, "MetricsTraceSampleRate": false, "MetricsSlowQueryLogThresholdMs": false, "SqlApiEnabled": false, "SqlApiAllowedStatements": false, "SqlApiMaxRows": false, "SqlApiMaxBytes": false, "SqlApiTimeoutSeconds": false}
	return &options
}

type UpdateOption func(*UpdateOptions)

func (srv *Mongo) WithUpdateName(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Mongo) WithUpdateStatus(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Mongo) WithUpdateSpecification(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Specification = v
		o.enabledSetters["Specification"] = true
	}
}
func (srv *Mongo) WithUpdateReplicas(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.Replicas = v
		o.enabledSetters["Replicas"] = true
	}
}
func (srv *Mongo) WithUpdateSyncMode(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.SyncMode = v
		o.enabledSetters["SyncMode"] = true
	}
}
func (srv *Mongo) WithUpdateNetworkIdleTimeoutSeconds(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.NetworkIdleTimeoutSeconds = v
		o.enabledSetters["NetworkIdleTimeoutSeconds"] = true
	}
}
func (srv *Mongo) WithUpdateNetworkIPAllowlist(v []string) UpdateOption {
	return func(o *UpdateOptions) {
		o.NetworkIPAllowlist = v
		o.enabledSetters["NetworkIPAllowlist"] = true
	}
}
func (srv *Mongo) WithUpdateIdleTimeoutMinutes(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.IdleTimeoutMinutes = v
		o.enabledSetters["IdleTimeoutMinutes"] = true
	}
}
func (srv *Mongo) WithUpdatePitr(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Pitr = v
		o.enabledSetters["Pitr"] = true
	}
}
func (srv *Mongo) WithUpdatePitrRetentionDays(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.PitrRetentionDays = v
		o.enabledSetters["PitrRetentionDays"] = true
	}
}
func (srv *Mongo) WithUpdateStorageAutoscaling(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.StorageAutoscaling = v
		o.enabledSetters["StorageAutoscaling"] = true
	}
}
func (srv *Mongo) WithUpdateStorageAutoscalingThresholdPercent(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.StorageAutoscalingThresholdPercent = v
		o.enabledSetters["StorageAutoscalingThresholdPercent"] = true
	}
}
func (srv *Mongo) WithUpdateStorageAutoscalingMaxGb(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.StorageAutoscalingMaxGb = v
		o.enabledSetters["StorageAutoscalingMaxGb"] = true
	}
}
func (srv *Mongo) WithUpdateMetricsTraceSampleRate(v float64) UpdateOption {
	return func(o *UpdateOptions) {
		o.MetricsTraceSampleRate = v
		o.enabledSetters["MetricsTraceSampleRate"] = true
	}
}
func (srv *Mongo) WithUpdateMetricsSlowQueryLogThresholdMs(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.MetricsSlowQueryLogThresholdMs = v
		o.enabledSetters["MetricsSlowQueryLogThresholdMs"] = true
	}
}
func (srv *Mongo) WithUpdateSqlApiEnabled(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.SqlApiEnabled = v
		o.enabledSetters["SqlApiEnabled"] = true
	}
}
func (srv *Mongo) WithUpdateSqlApiAllowedStatements(v []string) UpdateOption {
	return func(o *UpdateOptions) {
		o.SqlApiAllowedStatements = v
		o.enabledSetters["SqlApiAllowedStatements"] = true
	}
}
func (srv *Mongo) WithUpdateSqlApiMaxRows(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.SqlApiMaxRows = v
		o.enabledSetters["SqlApiMaxRows"] = true
	}
}
func (srv *Mongo) WithUpdateSqlApiMaxBytes(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.SqlApiMaxBytes = v
		o.enabledSetters["SqlApiMaxBytes"] = true
	}
}
func (srv *Mongo) WithUpdateSqlApiTimeoutSeconds(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.SqlApiTimeoutSeconds = v
		o.enabledSetters["SqlApiTimeoutSeconds"] = true
	}
}

// Update update a dedicated database configuration. All changes are applied
// with zero downtime. Specification changes (cpu, memory, storage) are
// handled via rolling cutover. Storage expansion is done online. All other
// settings are applied in-place.
func (srv *Mongo) Update(DatabaseId string, optionalSetters ...UpdateOption) (*models.DedicatedDatabase, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}")
	options := UpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["Specification"] {
		params["specification"] = options.Specification
	}
	if options.enabledSetters["Replicas"] {
		params["replicas"] = options.Replicas
	}
	if options.enabledSetters["SyncMode"] {
		params["syncMode"] = options.SyncMode
	}
	if options.enabledSetters["NetworkIdleTimeoutSeconds"] {
		params["networkIdleTimeoutSeconds"] = options.NetworkIdleTimeoutSeconds
	}
	if options.enabledSetters["NetworkIPAllowlist"] {
		params["networkIPAllowlist"] = options.NetworkIPAllowlist
	}
	if options.enabledSetters["IdleTimeoutMinutes"] {
		params["idleTimeoutMinutes"] = options.IdleTimeoutMinutes
	}
	if options.enabledSetters["Pitr"] {
		params["pitr"] = options.Pitr
	}
	if options.enabledSetters["PitrRetentionDays"] {
		params["pitrRetentionDays"] = options.PitrRetentionDays
	}
	if options.enabledSetters["StorageAutoscaling"] {
		params["storageAutoscaling"] = options.StorageAutoscaling
	}
	if options.enabledSetters["StorageAutoscalingThresholdPercent"] {
		params["storageAutoscalingThresholdPercent"] = options.StorageAutoscalingThresholdPercent
	}
	if options.enabledSetters["StorageAutoscalingMaxGb"] {
		params["storageAutoscalingMaxGb"] = options.StorageAutoscalingMaxGb
	}
	if options.enabledSetters["MetricsTraceSampleRate"] {
		params["metricsTraceSampleRate"] = options.MetricsTraceSampleRate
	}
	if options.enabledSetters["MetricsSlowQueryLogThresholdMs"] {
		params["metricsSlowQueryLogThresholdMs"] = options.MetricsSlowQueryLogThresholdMs
	}
	if options.enabledSetters["SqlApiEnabled"] {
		params["sqlApiEnabled"] = options.SqlApiEnabled
	}
	if options.enabledSetters["SqlApiAllowedStatements"] {
		params["sqlApiAllowedStatements"] = options.SqlApiAllowedStatements
	}
	if options.enabledSetters["SqlApiMaxRows"] {
		params["sqlApiMaxRows"] = options.SqlApiMaxRows
	}
	if options.enabledSetters["SqlApiMaxBytes"] {
		params["sqlApiMaxBytes"] = options.SqlApiMaxBytes
	}
	if options.enabledSetters["SqlApiTimeoutSeconds"] {
		params["sqlApiTimeoutSeconds"] = options.SqlApiTimeoutSeconds
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

		parsed := models.DedicatedDatabase{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabase
	parsed, ok := resp.Result.(models.DedicatedDatabase)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Delete delete a dedicated database. This action is irreversible. The
// database status will be set to 'deleting' and all resources will be cleaned
// up. Deletion is allowed from any state, and repeating the call
// re-dispatches the cleanup.
func (srv *Mongo) Delete(DatabaseId string) (*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}")
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

type ListBackupsOptions struct {
	Queries        []string
	enabledSetters map[string]bool
}

func (options ListBackupsOptions) New() *ListBackupsOptions {
	options.enabledSetters = map[string]bool{"Queries": false}
	return &options
}

type ListBackupsOption func(*ListBackupsOptions)

func (srv *Mongo) WithListBackupsQueries(v []string) ListBackupsOption {
	return func(o *ListBackupsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListBackups list all backups for a dedicated database. Results can be
// filtered by status and type.
func (srv *Mongo) ListBackups(DatabaseId string, optionalSetters ...ListBackupsOption) (*models.DedicatedDatabaseBackupList, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/backups")
	options := ListBackupsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
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

		parsed := models.DedicatedDatabaseBackupList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabaseBackupList
	parsed, ok := resp.Result.(models.DedicatedDatabaseBackupList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateBackupOptions struct {
	Type           string
	enabledSetters map[string]bool
}

func (options CreateBackupOptions) New() *CreateBackupOptions {
	options.enabledSetters = map[string]bool{"Type": false}
	return &options
}

type CreateBackupOption func(*CreateBackupOptions)

func (srv *Mongo) WithCreateBackupType(v string) CreateBackupOption {
	return func(o *CreateBackupOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}

// CreateBackup create a manual backup of a dedicated database. The backup
// will be created asynchronously and its status can be checked via the get
// backup endpoint.
func (srv *Mongo) CreateBackup(DatabaseId string, optionalSetters ...CreateBackupOption) (*models.DedicatedDatabaseBackup, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/backups")
	options := CreateBackupOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
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

		parsed := models.DedicatedDatabaseBackup{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabaseBackup
	parsed, ok := resp.Result.(models.DedicatedDatabaseBackup)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListBackupPoliciesOptions struct {
	Queries        []string
	enabledSetters map[string]bool
}

func (options ListBackupPoliciesOptions) New() *ListBackupPoliciesOptions {
	options.enabledSetters = map[string]bool{"Queries": false}
	return &options
}

type ListBackupPoliciesOption func(*ListBackupPoliciesOptions)

func (srv *Mongo) WithListBackupPoliciesQueries(v []string) ListBackupPoliciesOption {
	return func(o *ListBackupPoliciesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListBackupPolicies list scheduled backup policies for a dedicated database.
func (srv *Mongo) ListBackupPolicies(DatabaseId string, optionalSetters ...ListBackupPoliciesOption) (*models.BackupPolicyList, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/backups/policies")
	options := ListBackupPoliciesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
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

type CreateBackupPolicyOptions struct {
	Type           string
	Enabled        bool
	enabledSetters map[string]bool
}

func (options CreateBackupPolicyOptions) New() *CreateBackupPolicyOptions {
	options.enabledSetters = map[string]bool{"Type": false, "Enabled": false}
	return &options
}

type CreateBackupPolicyOption func(*CreateBackupPolicyOptions)

func (srv *Mongo) WithCreateBackupPolicyType(v string) CreateBackupPolicyOption {
	return func(o *CreateBackupPolicyOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *Mongo) WithCreateBackupPolicyEnabled(v bool) CreateBackupPolicyOption {
	return func(o *CreateBackupPolicyOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}

// CreateBackupPolicy create a scheduled backup policy for a dedicated
// database.
func (srv *Mongo) CreateBackupPolicy(DatabaseId string, PolicyId string, Name string, Schedule string, Retention int, optionalSetters ...CreateBackupPolicyOption) (*models.BackupPolicy, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/backups/policies")
	options := CreateBackupPolicyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["policyId"] = PolicyId
	params["name"] = Name
	params["schedule"] = Schedule
	params["retention"] = Retention
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

// GetBackupPolicy get a scheduled backup policy for a dedicated database.
func (srv *Mongo) GetBackupPolicy(DatabaseId string, PolicyId string) (*models.BackupPolicy, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{policyId}", client.EncodePath(PolicyId))
	path := r.Replace("/mongo/{databaseId}/backups/policies/{policyId}")
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

type UpdateBackupPolicyOptions struct {
	Name           string
	Schedule       string
	Retention      int
	Enabled        bool
	enabledSetters map[string]bool
}

func (options UpdateBackupPolicyOptions) New() *UpdateBackupPolicyOptions {
	options.enabledSetters = map[string]bool{"Name": false, "Schedule": false, "Retention": false, "Enabled": false}
	return &options
}

type UpdateBackupPolicyOption func(*UpdateBackupPolicyOptions)

func (srv *Mongo) WithUpdateBackupPolicyName(v string) UpdateBackupPolicyOption {
	return func(o *UpdateBackupPolicyOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Mongo) WithUpdateBackupPolicySchedule(v string) UpdateBackupPolicyOption {
	return func(o *UpdateBackupPolicyOptions) {
		o.Schedule = v
		o.enabledSetters["Schedule"] = true
	}
}
func (srv *Mongo) WithUpdateBackupPolicyRetention(v int) UpdateBackupPolicyOption {
	return func(o *UpdateBackupPolicyOptions) {
		o.Retention = v
		o.enabledSetters["Retention"] = true
	}
}
func (srv *Mongo) WithUpdateBackupPolicyEnabled(v bool) UpdateBackupPolicyOption {
	return func(o *UpdateBackupPolicyOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}

// UpdateBackupPolicy update a scheduled backup policy for a dedicated
// database.
func (srv *Mongo) UpdateBackupPolicy(DatabaseId string, PolicyId string, optionalSetters ...UpdateBackupPolicyOption) (*models.BackupPolicy, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{policyId}", client.EncodePath(PolicyId))
	path := r.Replace("/mongo/{databaseId}/backups/policies/{policyId}")
	options := UpdateBackupPolicyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Schedule"] {
		params["schedule"] = options.Schedule
	}
	if options.enabledSetters["Retention"] {
		params["retention"] = options.Retention
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

// DeleteBackupPolicy delete a scheduled backup policy for a dedicated
// database. Backups already taken by the policy are kept until their
// retention expires.
func (srv *Mongo) DeleteBackupPolicy(DatabaseId string, PolicyId string) (*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{policyId}", client.EncodePath(PolicyId))
	path := r.Replace("/mongo/{databaseId}/backups/policies/{policyId}")
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

type UpdateBackupStorageOptions struct {
	Region         string
	Prefix         string
	Endpoint       string
	enabledSetters map[string]bool
}

func (options UpdateBackupStorageOptions) New() *UpdateBackupStorageOptions {
	options.enabledSetters = map[string]bool{"Region": false, "Prefix": false, "Endpoint": false}
	return &options
}

type UpdateBackupStorageOption func(*UpdateBackupStorageOptions)

func (srv *Mongo) WithUpdateBackupStorageRegion(v string) UpdateBackupStorageOption {
	return func(o *UpdateBackupStorageOptions) {
		o.Region = v
		o.enabledSetters["Region"] = true
	}
}
func (srv *Mongo) WithUpdateBackupStoragePrefix(v string) UpdateBackupStorageOption {
	return func(o *UpdateBackupStorageOptions) {
		o.Prefix = v
		o.enabledSetters["Prefix"] = true
	}
}
func (srv *Mongo) WithUpdateBackupStorageEndpoint(v string) UpdateBackupStorageOption {
	return func(o *UpdateBackupStorageOptions) {
		o.Endpoint = v
		o.enabledSetters["Endpoint"] = true
	}
}

// UpdateBackupStorage configure off-cluster backup storage for a dedicated
// database. Supports S3, GCS, and Azure Blob Storage destinations. Backups
// will be stored to the configured destination in addition to on-cluster
// storage.
func (srv *Mongo) UpdateBackupStorage(DatabaseId string, Provider string, Bucket string, AccessKey string, SecretKey string, optionalSetters ...UpdateBackupStorageOption) (*models.DedicatedDatabaseBackupStorage, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/backups/storage")
	options := UpdateBackupStorageOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["provider"] = Provider
	params["bucket"] = Bucket
	if options.enabledSetters["Region"] {
		params["region"] = options.Region
	}
	if options.enabledSetters["Prefix"] {
		params["prefix"] = options.Prefix
	}
	if options.enabledSetters["Endpoint"] {
		params["endpoint"] = options.Endpoint
	}
	params["accessKey"] = AccessKey
	params["secretKey"] = SecretKey
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DedicatedDatabaseBackupStorage{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabaseBackupStorage
	parsed, ok := resp.Result.(models.DedicatedDatabaseBackupStorage)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetBackup get details of a specific database backup including its status,
// size, and timestamps.
func (srv *Mongo) GetBackup(DatabaseId string, BackupId string) (*models.DedicatedDatabaseBackup, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{backupId}", client.EncodePath(BackupId))
	path := r.Replace("/mongo/{databaseId}/backups/{backupId}")
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

		parsed := models.DedicatedDatabaseBackup{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabaseBackup
	parsed, ok := resp.Result.(models.DedicatedDatabaseBackup)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteBackup delete a database backup. This will permanently remove the
// backup from storage and cannot be undone.
func (srv *Mongo) DeleteBackup(DatabaseId string, BackupId string) (*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{backupId}", client.EncodePath(BackupId))
	path := r.Replace("/mongo/{databaseId}/backups/{backupId}")
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

// ListBranches list all ephemeral branches for a dedicated database. Returns
// branch metadata including ID, name, namespace, and expiration time.
func (srv *Mongo) ListBranches(DatabaseId string) (*models.DedicatedDatabaseBranchList, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/branches")
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

		parsed := models.DedicatedDatabaseBranchList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabaseBranchList
	parsed, ok := resp.Result.(models.DedicatedDatabaseBranchList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateBranchOptions struct {
	BranchId       string
	Ttl            int
	enabledSetters map[string]bool
}

func (options CreateBranchOptions) New() *CreateBranchOptions {
	options.enabledSetters = map[string]bool{"BranchId": false, "Ttl": false}
	return &options
}

type CreateBranchOption func(*CreateBranchOptions)

func (srv *Mongo) WithCreateBranchBranchId(v string) CreateBranchOption {
	return func(o *CreateBranchOptions) {
		o.BranchId = v
		o.enabledSetters["BranchId"] = true
	}
}
func (srv *Mongo) WithCreateBranchTtl(v int) CreateBranchOption {
	return func(o *CreateBranchOptions) {
		o.Ttl = v
		o.enabledSetters["Ttl"] = true
	}
}

// CreateBranch create an ephemeral database branch from the primary via PVC
// snapshot. The branch is a full copy of the database at the current point in
// time, useful for testing schema migrations or running experiments without
// affecting production data. Branches expire after the configured TTL
// (default 24 hours). The branch is created asynchronously.
func (srv *Mongo) CreateBranch(DatabaseId string, optionalSetters ...CreateBranchOption) (*models.DedicatedDatabase, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/branches")
	options := CreateBranchOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["BranchId"] {
		params["branchId"] = options.BranchId
	}
	if options.enabledSetters["Ttl"] {
		params["ttl"] = options.Ttl
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

		parsed := models.DedicatedDatabase{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabase
	parsed, ok := resp.Result.(models.DedicatedDatabase)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteBranch delete an ephemeral database branch. This removes the branch
// namespace, its PVC, and the associated VolumeSnapshot. The deletion runs
// asynchronously and is irreversible.
func (srv *Mongo) DeleteBranch(DatabaseId string, BranchId string) (*models.DedicatedDatabase, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{branchId}", client.EncodePath(BranchId))
	path := r.Replace("/mongo/{databaseId}/branches/{branchId}")
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

		parsed := models.DedicatedDatabase{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabase
	parsed, ok := resp.Result.(models.DedicatedDatabase)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateCredentials rotate the primary connection credentials for a dedicated
// database. Generates a new password and updates the database atomically.
// Previous credentials stop working immediately. Returns the database with a
// refreshed connection string carrying the new password.
func (srv *Mongo) UpdateCredentials(DatabaseId string) (*models.DedicatedDatabase, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/credentials")
	params := map[string]interface{}{}
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

		parsed := models.DedicatedDatabase{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabase
	parsed, ok := resp.Result.(models.DedicatedDatabase)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateFailoverOptions struct {
	TargetReplicaId string
	enabledSetters  map[string]bool
}

func (options CreateFailoverOptions) New() *CreateFailoverOptions {
	options.enabledSetters = map[string]bool{"TargetReplicaId": false}
	return &options
}

type CreateFailoverOption func(*CreateFailoverOptions)

func (srv *Mongo) WithCreateFailoverTargetReplicaId(v string) CreateFailoverOption {
	return func(o *CreateFailoverOptions) {
		o.TargetReplicaId = v
		o.enabledSetters["TargetReplicaId"] = true
	}
}

// CreateFailover trigger a manual failover for a dedicated database with high
// availability enabled. Promotes a replica to primary. The failover runs
// asynchronously; poll the database document for status updates. A database
// left mid-operation by a failover that did not finish also accepts this call
// as a repair, provided `targetReplicaId` names the member to promote.
func (srv *Mongo) CreateFailover(DatabaseId string, optionalSetters ...CreateFailoverOption) (*models.DedicatedDatabase, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/failovers")
	options := CreateFailoverOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["TargetReplicaId"] {
		params["targetReplicaId"] = options.TargetReplicaId
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

		parsed := models.DedicatedDatabase{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabase
	parsed, ok := resp.Result.(models.DedicatedDatabase)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateMaintenance update the maintenance window for a dedicated database.
// Maintenance operations like minor version upgrades will be performed during
// this window.
func (srv *Mongo) UpdateMaintenance(DatabaseId string, Day string, HourUtc int) (*models.DedicatedDatabase, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/maintenance")
	params := map[string]interface{}{}
	params["day"] = Day
	params["hourUtc"] = HourUtc
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

		parsed := models.DedicatedDatabase{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabase
	parsed, ok := resp.Result.(models.DedicatedDatabase)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateMigrationOptions struct {
	Specification  string
	enabledSetters map[string]bool
}

func (options CreateMigrationOptions) New() *CreateMigrationOptions {
	options.enabledSetters = map[string]bool{"Specification": false}
	return &options
}

type CreateMigrationOption func(*CreateMigrationOptions)

func (srv *Mongo) WithCreateMigrationSpecification(v string) CreateMigrationOption {
	return func(o *CreateMigrationOptions) {
		o.Specification = v
		o.enabledSetters["Specification"] = true
	}
}

// CreateMigration migrate a database between shared and dedicated types.
// Shared to dedicated provisions an always-on dedicated instance; dedicated
// to shared converts to a serverless instance that scales to zero when idle.
// Data is copied to the target with a brief read-only window during cutover.
func (srv *Mongo) CreateMigration(DatabaseId string, TargetType string, optionalSetters ...CreateMigrationOption) (*models.DedicatedDatabase, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/migrations")
	options := CreateMigrationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["targetType"] = TargetType
	if options.enabledSetters["Specification"] {
		params["specification"] = options.Specification
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

		parsed := models.DedicatedDatabase{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabase
	parsed, ok := resp.Result.(models.DedicatedDatabase)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListOperationsOptions struct {
	Status         string
	Limit          int
	Offset         int
	enabledSetters map[string]bool
}

func (options ListOperationsOptions) New() *ListOperationsOptions {
	options.enabledSetters = map[string]bool{"Status": false, "Limit": false, "Offset": false}
	return &options
}

type ListOperationsOption func(*ListOperationsOptions)

func (srv *Mongo) WithListOperationsStatus(v string) ListOperationsOption {
	return func(o *ListOperationsOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Mongo) WithListOperationsLimit(v int) ListOperationsOption {
	return func(o *ListOperationsOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Mongo) WithListOperationsOffset(v int) ListOperationsOption {
	return func(o *ListOperationsOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}

// ListOperations list the lifecycle operations recorded for a dedicated
// database, newest first. Every provision, update, restore, backup and
// replication action is recorded here with its outcome, including an attempt
// that was abandoned because another worker took over the database.
func (srv *Mongo) ListOperations(DatabaseId string, optionalSetters ...ListOperationsOption) (*models.DedicatedDatabaseOperationList, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/operations")
	options := ListOperationsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
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

		parsed := models.DedicatedDatabaseOperationList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabaseOperationList
	parsed, ok := resp.Result.(models.DedicatedDatabaseOperationList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetPitr get available point-in-time recovery windows for a dedicated
// database. Returns the earliest and latest recovery points.
func (srv *Mongo) GetPitr(DatabaseId string) (*models.DedicatedDatabasePITRWindows, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/pitr")
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

		parsed := models.DedicatedDatabasePITRWindows{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabasePITRWindows
	parsed, ok := resp.Result.(models.DedicatedDatabasePITRWindows)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetReplicas get high availability status for a dedicated database. Returns
// replica statuses, replication lag, and sync mode.
func (srv *Mongo) GetReplicas(DatabaseId string) (*models.DedicatedDatabaseReplicas, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/replicas")
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

		parsed := models.DedicatedDatabaseReplicas{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabaseReplicas
	parsed, ok := resp.Result.(models.DedicatedDatabaseReplicas)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListRestorationsOptions struct {
	Status         string
	Type           string
	Limit          int
	Offset         int
	enabledSetters map[string]bool
}

func (options ListRestorationsOptions) New() *ListRestorationsOptions {
	options.enabledSetters = map[string]bool{"Status": false, "Type": false, "Limit": false, "Offset": false}
	return &options
}

type ListRestorationsOption func(*ListRestorationsOptions)

func (srv *Mongo) WithListRestorationsStatus(v string) ListRestorationsOption {
	return func(o *ListRestorationsOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Mongo) WithListRestorationsType(v string) ListRestorationsOption {
	return func(o *ListRestorationsOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *Mongo) WithListRestorationsLimit(v int) ListRestorationsOption {
	return func(o *ListRestorationsOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Mongo) WithListRestorationsOffset(v int) ListRestorationsOption {
	return func(o *ListRestorationsOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}

// ListRestorations list all restorations for a dedicated database. Results
// can be filtered by status and type.
func (srv *Mongo) ListRestorations(DatabaseId string, optionalSetters ...ListRestorationsOption) (*models.DedicatedDatabaseRestorationList, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/restorations")
	options := ListRestorationsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
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

		parsed := models.DedicatedDatabaseRestorationList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabaseRestorationList
	parsed, ok := resp.Result.(models.DedicatedDatabaseRestorationList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateRestorationOptions struct {
	Type             string
	BackupId         string
	TargetDatabaseId string
	TargetTime       string
	enabledSetters   map[string]bool
}

func (options CreateRestorationOptions) New() *CreateRestorationOptions {
	options.enabledSetters = map[string]bool{"Type": false, "BackupId": false, "TargetDatabaseId": false, "TargetTime": false}
	return &options
}

type CreateRestorationOption func(*CreateRestorationOptions)

func (srv *Mongo) WithCreateRestorationType(v string) CreateRestorationOption {
	return func(o *CreateRestorationOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *Mongo) WithCreateRestorationBackupId(v string) CreateRestorationOption {
	return func(o *CreateRestorationOptions) {
		o.BackupId = v
		o.enabledSetters["BackupId"] = true
	}
}
func (srv *Mongo) WithCreateRestorationTargetDatabaseId(v string) CreateRestorationOption {
	return func(o *CreateRestorationOptions) {
		o.TargetDatabaseId = v
		o.enabledSetters["TargetDatabaseId"] = true
	}
}
func (srv *Mongo) WithCreateRestorationTargetTime(v string) CreateRestorationOption {
	return func(o *CreateRestorationOptions) {
		o.TargetTime = v
		o.enabledSetters["TargetTime"] = true
	}
}

// CreateRestoration restore a database from a backup or to a specific point
// in time (PITR). For backup restoration, provide a backupId. For PITR,
// provide a targetTime as an ISO 8601 datetime. PITR requires the database to
// have PITR enabled and is only available for enterprise databases.
func (srv *Mongo) CreateRestoration(DatabaseId string, optionalSetters ...CreateRestorationOption) (*models.DedicatedDatabaseRestoration, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/restorations")
	options := CreateRestorationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["BackupId"] {
		params["backupId"] = options.BackupId
	}
	if options.enabledSetters["TargetDatabaseId"] {
		params["targetDatabaseId"] = options.TargetDatabaseId
	}
	if options.enabledSetters["TargetTime"] {
		params["targetTime"] = options.TargetTime
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

		parsed := models.DedicatedDatabaseRestoration{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabaseRestoration
	parsed, ok := resp.Result.(models.DedicatedDatabaseRestoration)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetRestoration get details of a specific database restoration including its
// status, type, and timestamps.
func (srv *Mongo) GetRestoration(DatabaseId string, RestorationId string) (*models.DedicatedDatabaseRestoration, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{restorationId}", client.EncodePath(RestorationId))
	path := r.Replace("/mongo/{databaseId}/restorations/{restorationId}")
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

		parsed := models.DedicatedDatabaseRestoration{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabaseRestoration
	parsed, ok := resp.Result.(models.DedicatedDatabaseRestoration)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetStatus get real-time health and status information for a dedicated
// database. Returns health status, readiness, uptime, connection info,
// replica status, and volume information.
func (srv *Mongo) GetStatus(DatabaseId string) (*models.DatabaseStatus, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/status")
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

		parsed := models.DatabaseStatus{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DatabaseStatus
	parsed, ok := resp.Result.(models.DatabaseStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreateUpgrade upgrade a dedicated database to a new engine version. Uses
// blue-green deployment for zero-downtime cutover.
func (srv *Mongo) CreateUpgrade(DatabaseId string, TargetVersion string) (*models.DedicatedDatabase, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/mongo/{databaseId}/upgrades")
	params := map[string]interface{}{}
	params["targetVersion"] = TargetVersion
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

		parsed := models.DedicatedDatabase{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DedicatedDatabase
	parsed, ok := resp.Result.(models.DedicatedDatabase)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
