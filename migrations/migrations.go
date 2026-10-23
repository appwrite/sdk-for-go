package migrations

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/models"
)

// Migrations service
type Migrations struct {
	client client.Client
}

func New(clt client.Client) *Migrations {
	return &Migrations{
		client: clt,
	}
}

type ListOptions struct {
	Queries        []string
	Search         string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Search": false, "Total": false}
	return &options
}

type ListOption func(*ListOptions)

func (srv *Migrations) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Migrations) WithListSearch(v string) ListOption {
	return func(o *ListOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Migrations) WithListTotal(v bool) ListOption {
	return func(o *ListOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// List list all migrations in the current project. This endpoint returns a
// list of all migrations including their status, progress, and any errors
// that occurred during the migration process.
func (srv *Migrations) List(optionalSetters ...ListOption) (*models.MigrationList, error) {
	path := "/migrations"
	options := ListOptions{}.New()
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

		parsed := models.MigrationList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MigrationList
	parsed, ok := resp.Result.(models.MigrationList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateAppwriteMigrationOptions struct {
	OnDuplicate    string
	enabledSetters map[string]bool
}

func (options CreateAppwriteMigrationOptions) New() *CreateAppwriteMigrationOptions {
	options.enabledSetters = map[string]bool{"OnDuplicate": false}
	return &options
}

type CreateAppwriteMigrationOption func(*CreateAppwriteMigrationOptions)

func (srv *Migrations) WithCreateAppwriteMigrationOnDuplicate(v string) CreateAppwriteMigrationOption {
	return func(o *CreateAppwriteMigrationOptions) {
		o.OnDuplicate = v
		o.enabledSetters["OnDuplicate"] = true
	}
}

// CreateAppwriteMigration migrate data from another Appwrite project to your
// current project. This endpoint allows you to migrate resources like
// databases, collections, documents, users, and files from an existing
// Appwrite project.
func (srv *Migrations) CreateAppwriteMigration(Resources []string, Endpoint string, ProjectId string, ApiKey string, optionalSetters ...CreateAppwriteMigrationOption) (*models.Migration, error) {
	path := "/migrations/appwrite"
	options := CreateAppwriteMigrationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["resources"] = Resources
	params["endpoint"] = Endpoint
	params["projectId"] = ProjectId
	params["apiKey"] = ApiKey
	if options.enabledSetters["OnDuplicate"] {
		params["onDuplicate"] = options.OnDuplicate
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

		parsed := models.Migration{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Migration
	parsed, ok := resp.Result.(models.Migration)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetAppwriteReport generate a report of the data in an Appwrite project
// before migrating. This endpoint analyzes the source project and returns
// information about the resources that can be migrated.
func (srv *Migrations) GetAppwriteReport(Resources []string, Endpoint string, ProjectID string, Key string) (*models.MigrationReport, error) {
	path := "/migrations/appwrite/report"
	params := map[string]interface{}{}
	params["resources"] = Resources
	params["endpoint"] = Endpoint
	params["projectID"] = ProjectID
	params["key"] = Key
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

		parsed := models.MigrationReport{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MigrationReport
	parsed, ok := resp.Result.(models.MigrationReport)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateCSVExportOptions struct {
	Columns        []string
	Queries        []string
	Delimiter      string
	Enclosure      string
	Escape         string
	Header         bool
	Notify         bool
	enabledSetters map[string]bool
}

func (options CreateCSVExportOptions) New() *CreateCSVExportOptions {
	options.enabledSetters = map[string]bool{"Columns": false, "Queries": false, "Delimiter": false, "Enclosure": false, "Escape": false, "Header": false, "Notify": false}
	return &options
}

type CreateCSVExportOption func(*CreateCSVExportOptions)

func (srv *Migrations) WithCreateCSVExportColumns(v []string) CreateCSVExportOption {
	return func(o *CreateCSVExportOptions) {
		o.Columns = v
		o.enabledSetters["Columns"] = true
	}
}
func (srv *Migrations) WithCreateCSVExportQueries(v []string) CreateCSVExportOption {
	return func(o *CreateCSVExportOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Migrations) WithCreateCSVExportDelimiter(v string) CreateCSVExportOption {
	return func(o *CreateCSVExportOptions) {
		o.Delimiter = v
		o.enabledSetters["Delimiter"] = true
	}
}
func (srv *Migrations) WithCreateCSVExportEnclosure(v string) CreateCSVExportOption {
	return func(o *CreateCSVExportOptions) {
		o.Enclosure = v
		o.enabledSetters["Enclosure"] = true
	}
}
func (srv *Migrations) WithCreateCSVExportEscape(v string) CreateCSVExportOption {
	return func(o *CreateCSVExportOptions) {
		o.Escape = v
		o.enabledSetters["Escape"] = true
	}
}
func (srv *Migrations) WithCreateCSVExportHeader(v bool) CreateCSVExportOption {
	return func(o *CreateCSVExportOptions) {
		o.Header = v
		o.enabledSetters["Header"] = true
	}
}
func (srv *Migrations) WithCreateCSVExportNotify(v bool) CreateCSVExportOption {
	return func(o *CreateCSVExportOptions) {
		o.Notify = v
		o.enabledSetters["Notify"] = true
	}
}

// CreateCSVExport export documents to a CSV file from your Appwrite database.
// This endpoint allows you to export documents to a CSV file stored in a
// secure internal bucket. You'll receive an email with a download link when
// the export is complete.
func (srv *Migrations) CreateCSVExport(DatabaseId string, CollectionId string, Filename string, optionalSetters ...CreateCSVExportOption) (*models.Migration, error) {
	path := "/migrations/csv/exports"
	options := CreateCSVExportOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["filename"] = Filename
	if options.enabledSetters["Columns"] {
		params["columns"] = options.Columns
	}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Delimiter"] {
		params["delimiter"] = options.Delimiter
	}
	if options.enabledSetters["Enclosure"] {
		params["enclosure"] = options.Enclosure
	}
	if options.enabledSetters["Escape"] {
		params["escape"] = options.Escape
	}
	if options.enabledSetters["Header"] {
		params["header"] = options.Header
	}
	if options.enabledSetters["Notify"] {
		params["notify"] = options.Notify
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

		parsed := models.Migration{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Migration
	parsed, ok := resp.Result.(models.Migration)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateCSVImportOptions struct {
	InternalFile   bool
	OnDuplicate    string
	enabledSetters map[string]bool
}

func (options CreateCSVImportOptions) New() *CreateCSVImportOptions {
	options.enabledSetters = map[string]bool{"InternalFile": false, "OnDuplicate": false}
	return &options
}

type CreateCSVImportOption func(*CreateCSVImportOptions)

func (srv *Migrations) WithCreateCSVImportInternalFile(v bool) CreateCSVImportOption {
	return func(o *CreateCSVImportOptions) {
		o.InternalFile = v
		o.enabledSetters["InternalFile"] = true
	}
}
func (srv *Migrations) WithCreateCSVImportOnDuplicate(v string) CreateCSVImportOption {
	return func(o *CreateCSVImportOptions) {
		o.OnDuplicate = v
		o.enabledSetters["OnDuplicate"] = true
	}
}

// CreateCSVImport import documents from a CSV file into your Appwrite
// database. This endpoint allows you to import documents from a CSV file
// uploaded to Appwrite Storage bucket.
func (srv *Migrations) CreateCSVImport(BucketId string, FileId string, DatabaseId string, CollectionId string, optionalSetters ...CreateCSVImportOption) (*models.Migration, error) {
	path := "/migrations/csv/imports"
	options := CreateCSVImportOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	if options.enabledSetters["InternalFile"] {
		params["internalFile"] = options.InternalFile
	}
	if options.enabledSetters["OnDuplicate"] {
		params["onDuplicate"] = options.OnDuplicate
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

		parsed := models.Migration{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Migration
	parsed, ok := resp.Result.(models.Migration)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreateFirebaseMigration migrate data from a Firebase project to your
// Appwrite project. This endpoint allows you to migrate resources like
// authentication and other supported services from a Firebase project.
func (srv *Migrations) CreateFirebaseMigration(Resources []string, ServiceAccount string) (*models.Migration, error) {
	path := "/migrations/firebase"
	params := map[string]interface{}{}
	params["resources"] = Resources
	params["serviceAccount"] = ServiceAccount
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

		parsed := models.Migration{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Migration
	parsed, ok := resp.Result.(models.Migration)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetFirebaseReport generate a report of the data in a Firebase project
// before migrating. This endpoint analyzes the source project and returns
// information about the resources that can be migrated.
func (srv *Migrations) GetFirebaseReport(Resources []string, ServiceAccount string) (*models.MigrationReport, error) {
	path := "/migrations/firebase/report"
	params := map[string]interface{}{}
	params["resources"] = Resources
	params["serviceAccount"] = ServiceAccount
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

		parsed := models.MigrationReport{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MigrationReport
	parsed, ok := resp.Result.(models.MigrationReport)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateJSONExportOptions struct {
	Columns        []string
	Queries        []string
	Notify         bool
	enabledSetters map[string]bool
}

func (options CreateJSONExportOptions) New() *CreateJSONExportOptions {
	options.enabledSetters = map[string]bool{"Columns": false, "Queries": false, "Notify": false}
	return &options
}

type CreateJSONExportOption func(*CreateJSONExportOptions)

func (srv *Migrations) WithCreateJSONExportColumns(v []string) CreateJSONExportOption {
	return func(o *CreateJSONExportOptions) {
		o.Columns = v
		o.enabledSetters["Columns"] = true
	}
}
func (srv *Migrations) WithCreateJSONExportQueries(v []string) CreateJSONExportOption {
	return func(o *CreateJSONExportOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Migrations) WithCreateJSONExportNotify(v bool) CreateJSONExportOption {
	return func(o *CreateJSONExportOptions) {
		o.Notify = v
		o.enabledSetters["Notify"] = true
	}
}

// CreateJSONExport export documents to a JSON file from your Appwrite
// database. This endpoint allows you to export documents to a JSON file
// stored in a secure internal bucket. You'll receive an email with a download
// link when the export is complete.
func (srv *Migrations) CreateJSONExport(DatabaseId string, CollectionId string, Filename string, optionalSetters ...CreateJSONExportOption) (*models.Migration, error) {
	path := "/migrations/json/exports"
	options := CreateJSONExportOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["filename"] = Filename
	if options.enabledSetters["Columns"] {
		params["columns"] = options.Columns
	}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Notify"] {
		params["notify"] = options.Notify
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

		parsed := models.Migration{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Migration
	parsed, ok := resp.Result.(models.Migration)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateJSONImportOptions struct {
	InternalFile   bool
	OnDuplicate    string
	enabledSetters map[string]bool
}

func (options CreateJSONImportOptions) New() *CreateJSONImportOptions {
	options.enabledSetters = map[string]bool{"InternalFile": false, "OnDuplicate": false}
	return &options
}

type CreateJSONImportOption func(*CreateJSONImportOptions)

func (srv *Migrations) WithCreateJSONImportInternalFile(v bool) CreateJSONImportOption {
	return func(o *CreateJSONImportOptions) {
		o.InternalFile = v
		o.enabledSetters["InternalFile"] = true
	}
}
func (srv *Migrations) WithCreateJSONImportOnDuplicate(v string) CreateJSONImportOption {
	return func(o *CreateJSONImportOptions) {
		o.OnDuplicate = v
		o.enabledSetters["OnDuplicate"] = true
	}
}

// CreateJSONImport import documents from a JSON file into your Appwrite
// database. This endpoint allows you to import documents from a JSON file
// uploaded to Appwrite Storage bucket.
func (srv *Migrations) CreateJSONImport(BucketId string, FileId string, DatabaseId string, CollectionId string, optionalSetters ...CreateJSONImportOption) (*models.Migration, error) {
	path := "/migrations/json/imports"
	options := CreateJSONImportOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	if options.enabledSetters["InternalFile"] {
		params["internalFile"] = options.InternalFile
	}
	if options.enabledSetters["OnDuplicate"] {
		params["onDuplicate"] = options.OnDuplicate
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

		parsed := models.Migration{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Migration
	parsed, ok := resp.Result.(models.Migration)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateNHostMigrationOptions struct {
	Port           int
	enabledSetters map[string]bool
}

func (options CreateNHostMigrationOptions) New() *CreateNHostMigrationOptions {
	options.enabledSetters = map[string]bool{"Port": false}
	return &options
}

type CreateNHostMigrationOption func(*CreateNHostMigrationOptions)

func (srv *Migrations) WithCreateNHostMigrationPort(v int) CreateNHostMigrationOption {
	return func(o *CreateNHostMigrationOptions) {
		o.Port = v
		o.enabledSetters["Port"] = true
	}
}

// CreateNHostMigration migrate data from an NHost project to your Appwrite
// project. This endpoint allows you to migrate resources like authentication,
// databases, and other supported services from an NHost project.
func (srv *Migrations) CreateNHostMigration(Resources []string, Subdomain string, Region string, AdminSecret string, Database string, Username string, Password string, optionalSetters ...CreateNHostMigrationOption) (*models.Migration, error) {
	path := "/migrations/nhost"
	options := CreateNHostMigrationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["resources"] = Resources
	params["subdomain"] = Subdomain
	params["region"] = Region
	params["adminSecret"] = AdminSecret
	params["database"] = Database
	params["username"] = Username
	params["password"] = Password
	if options.enabledSetters["Port"] {
		params["port"] = options.Port
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

		parsed := models.Migration{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Migration
	parsed, ok := resp.Result.(models.Migration)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type GetNHostReportOptions struct {
	Port           int
	enabledSetters map[string]bool
}

func (options GetNHostReportOptions) New() *GetNHostReportOptions {
	options.enabledSetters = map[string]bool{"Port": false}
	return &options
}

type GetNHostReportOption func(*GetNHostReportOptions)

func (srv *Migrations) WithGetNHostReportPort(v int) GetNHostReportOption {
	return func(o *GetNHostReportOptions) {
		o.Port = v
		o.enabledSetters["Port"] = true
	}
}

// GetNHostReport generate a detailed report of the data in an NHost project
// before migrating. This endpoint analyzes the source project and returns
// information about the resources that can be migrated.
func (srv *Migrations) GetNHostReport(Resources []string, Subdomain string, Region string, AdminSecret string, Database string, Username string, Password string, optionalSetters ...GetNHostReportOption) (*models.MigrationReport, error) {
	path := "/migrations/nhost/report"
	options := GetNHostReportOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["resources"] = Resources
	params["subdomain"] = Subdomain
	params["region"] = Region
	params["adminSecret"] = AdminSecret
	params["database"] = Database
	params["username"] = Username
	params["password"] = Password
	if options.enabledSetters["Port"] {
		params["port"] = options.Port
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

		parsed := models.MigrationReport{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MigrationReport
	parsed, ok := resp.Result.(models.MigrationReport)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateSupabaseMigrationOptions struct {
	Port           int
	enabledSetters map[string]bool
}

func (options CreateSupabaseMigrationOptions) New() *CreateSupabaseMigrationOptions {
	options.enabledSetters = map[string]bool{"Port": false}
	return &options
}

type CreateSupabaseMigrationOption func(*CreateSupabaseMigrationOptions)

func (srv *Migrations) WithCreateSupabaseMigrationPort(v int) CreateSupabaseMigrationOption {
	return func(o *CreateSupabaseMigrationOptions) {
		o.Port = v
		o.enabledSetters["Port"] = true
	}
}

// CreateSupabaseMigration migrate data from a Supabase project to your
// Appwrite project. This endpoint allows you to migrate resources like
// authentication, databases, and other supported services from a Supabase
// project.
func (srv *Migrations) CreateSupabaseMigration(Resources []string, Endpoint string, ApiKey string, DatabaseHost string, Username string, Password string, optionalSetters ...CreateSupabaseMigrationOption) (*models.Migration, error) {
	path := "/migrations/supabase"
	options := CreateSupabaseMigrationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["resources"] = Resources
	params["endpoint"] = Endpoint
	params["apiKey"] = ApiKey
	params["databaseHost"] = DatabaseHost
	params["username"] = Username
	params["password"] = Password
	if options.enabledSetters["Port"] {
		params["port"] = options.Port
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

		parsed := models.Migration{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Migration
	parsed, ok := resp.Result.(models.Migration)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type GetSupabaseReportOptions struct {
	Port           int
	enabledSetters map[string]bool
}

func (options GetSupabaseReportOptions) New() *GetSupabaseReportOptions {
	options.enabledSetters = map[string]bool{"Port": false}
	return &options
}

type GetSupabaseReportOption func(*GetSupabaseReportOptions)

func (srv *Migrations) WithGetSupabaseReportPort(v int) GetSupabaseReportOption {
	return func(o *GetSupabaseReportOptions) {
		o.Port = v
		o.enabledSetters["Port"] = true
	}
}

// GetSupabaseReport generate a report of the data in a Supabase project
// before migrating. This endpoint analyzes the source project and returns
// information about the resources that can be migrated.
func (srv *Migrations) GetSupabaseReport(Resources []string, Endpoint string, ApiKey string, DatabaseHost string, Username string, Password string, optionalSetters ...GetSupabaseReportOption) (*models.MigrationReport, error) {
	path := "/migrations/supabase/report"
	options := GetSupabaseReportOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["resources"] = Resources
	params["endpoint"] = Endpoint
	params["apiKey"] = ApiKey
	params["databaseHost"] = DatabaseHost
	params["username"] = Username
	params["password"] = Password
	if options.enabledSetters["Port"] {
		params["port"] = options.Port
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

		parsed := models.MigrationReport{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MigrationReport
	parsed, ok := resp.Result.(models.MigrationReport)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Get get a migration by its unique ID. This endpoint returns detailed
// information about a specific migration including its current status,
// progress, and any errors that occurred during the migration process.
func (srv *Migrations) Get(MigrationId string) (*models.Migration, error) {
	if MigrationId == "" {
		return nil, errors.New("Missing required parameter: \"migrationId\"")
	}

	r := strings.NewReplacer("{migrationId}", client.EncodePath(MigrationId))
	path := r.Replace("/migrations/{migrationId}")
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

		parsed := models.Migration{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Migration
	parsed, ok := resp.Result.(models.Migration)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Retry retry a failed migration. This endpoint allows you to retry a
// migration that has previously failed.
func (srv *Migrations) Retry(MigrationId string) (*models.Migration, error) {
	if MigrationId == "" {
		return nil, errors.New("Missing required parameter: \"migrationId\"")
	}

	r := strings.NewReplacer("{migrationId}", client.EncodePath(MigrationId))
	path := r.Replace("/migrations/{migrationId}")
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

		parsed := models.Migration{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Migration
	parsed, ok := resp.Result.(models.Migration)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Delete delete a migration by its unique ID. This endpoint allows you to
// remove a migration from your project's migration history.
func (srv *Migrations) Delete(MigrationId string) (*interface{}, error) {
	if MigrationId == "" {
		return nil, errors.New("Missing required parameter: \"migrationId\"")
	}

	r := strings.NewReplacer("{migrationId}", client.EncodePath(MigrationId))
	path := r.Replace("/migrations/{migrationId}")
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
