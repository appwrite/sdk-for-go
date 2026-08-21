package vectorsdb

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/models"
)

// VectorsDB service
type VectorsDB struct {
	client client.Client
}

func New(clt client.Client) *VectorsDB {
	return &VectorsDB{
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

func (srv *VectorsDB) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *VectorsDB) WithListTotal(v bool) ListOption {
	return func(o *ListOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// List get a list of all databases from the current Appwrite project. You can
// use the search parameter to filter your results.
func (srv *VectorsDB) List(optionalSetters ...ListOption) (*models.DatabaseList, error) {
	path := "/vectorsdb"
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

		parsed := models.DatabaseList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DatabaseList
	parsed, ok := resp.Result.(models.DatabaseList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateOptions struct {
	Enabled        bool
	Specification  string
	Replicas       int
	SyncMode       string
	enabledSetters map[string]bool
}

func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{"Enabled": false, "Specification": false, "Replicas": false, "SyncMode": false}
	return &options
}

type CreateOption func(*CreateOptions)

func (srv *VectorsDB) WithCreateEnabled(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *VectorsDB) WithCreateSpecification(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Specification = v
		o.enabledSetters["Specification"] = true
	}
}
func (srv *VectorsDB) WithCreateReplicas(v int) CreateOption {
	return func(o *CreateOptions) {
		o.Replicas = v
		o.enabledSetters["Replicas"] = true
	}
}
func (srv *VectorsDB) WithCreateSyncMode(v string) CreateOption {
	return func(o *CreateOptions) {
		o.SyncMode = v
		o.enabledSetters["SyncMode"] = true
	}
}

// Create create a new Database.
func (srv *VectorsDB) Create(DatabaseId string, Name string, optionalSetters ...CreateOption) (*models.Database, error) {
	path := "/vectorsdb"
	options := CreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["name"] = Name
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Database{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Database
	parsed, ok := resp.Result.(models.Database)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListSpecifications list the dedicated database specifications available on
// the current plan. Each specification reports its resource limits, pricing,
// and whether it is enabled for the organization.
func (srv *VectorsDB) ListSpecifications() (*models.DedicatedDatabaseSpecificationList, error) {
	path := "/vectorsdb/specifications"
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

type ListTransactionsOptions struct {
	Queries        []string
	enabledSetters map[string]bool
}

func (options ListTransactionsOptions) New() *ListTransactionsOptions {
	options.enabledSetters = map[string]bool{"Queries": false}
	return &options
}

type ListTransactionsOption func(*ListTransactionsOptions)

func (srv *VectorsDB) WithListTransactionsQueries(v []string) ListTransactionsOption {
	return func(o *ListTransactionsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListTransactions list transactions across all databases.
func (srv *VectorsDB) ListTransactions(optionalSetters ...ListTransactionsOption) (*models.TransactionList, error) {
	path := "/vectorsdb/transactions"
	options := ListTransactionsOptions{}.New()
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

		parsed := models.TransactionList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.TransactionList
	parsed, ok := resp.Result.(models.TransactionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateTransactionOptions struct {
	Ttl            int
	enabledSetters map[string]bool
}

func (options CreateTransactionOptions) New() *CreateTransactionOptions {
	options.enabledSetters = map[string]bool{"Ttl": false}
	return &options
}

type CreateTransactionOption func(*CreateTransactionOptions)

func (srv *VectorsDB) WithCreateTransactionTtl(v int) CreateTransactionOption {
	return func(o *CreateTransactionOptions) {
		o.Ttl = v
		o.enabledSetters["Ttl"] = true
	}
}

// CreateTransaction create a new transaction.
func (srv *VectorsDB) CreateTransaction(optionalSetters ...CreateTransactionOption) (*models.Transaction, error) {
	path := "/vectorsdb/transactions"
	options := CreateTransactionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
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

		parsed := models.Transaction{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Transaction
	parsed, ok := resp.Result.(models.Transaction)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetTransaction get a transaction by its unique ID.
func (srv *VectorsDB) GetTransaction(TransactionId string) (*models.Transaction, error) {
	r := strings.NewReplacer("{transactionId}", client.EncodePath(TransactionId))
	path := r.Replace("/vectorsdb/transactions/{transactionId}")
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

		parsed := models.Transaction{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Transaction
	parsed, ok := resp.Result.(models.Transaction)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateTransactionOptions struct {
	Commit         bool
	Rollback       bool
	enabledSetters map[string]bool
}

func (options UpdateTransactionOptions) New() *UpdateTransactionOptions {
	options.enabledSetters = map[string]bool{"Commit": false, "Rollback": false}
	return &options
}

type UpdateTransactionOption func(*UpdateTransactionOptions)

func (srv *VectorsDB) WithUpdateTransactionCommit(v bool) UpdateTransactionOption {
	return func(o *UpdateTransactionOptions) {
		o.Commit = v
		o.enabledSetters["Commit"] = true
	}
}
func (srv *VectorsDB) WithUpdateTransactionRollback(v bool) UpdateTransactionOption {
	return func(o *UpdateTransactionOptions) {
		o.Rollback = v
		o.enabledSetters["Rollback"] = true
	}
}

// UpdateTransaction update a transaction, to either commit or roll back its
// operations.
func (srv *VectorsDB) UpdateTransaction(TransactionId string, optionalSetters ...UpdateTransactionOption) (*models.Transaction, error) {
	r := strings.NewReplacer("{transactionId}", client.EncodePath(TransactionId))
	path := r.Replace("/vectorsdb/transactions/{transactionId}")
	options := UpdateTransactionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Commit"] {
		params["commit"] = options.Commit
	}
	if options.enabledSetters["Rollback"] {
		params["rollback"] = options.Rollback
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

		parsed := models.Transaction{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Transaction
	parsed, ok := resp.Result.(models.Transaction)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteTransaction delete a transaction by its unique ID.
func (srv *VectorsDB) DeleteTransaction(TransactionId string) (*interface{}, error) {
	r := strings.NewReplacer("{transactionId}", client.EncodePath(TransactionId))
	path := r.Replace("/vectorsdb/transactions/{transactionId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"

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

type CreateOperationsOptions struct {
	Operations     []interface{}
	enabledSetters map[string]bool
}

func (options CreateOperationsOptions) New() *CreateOperationsOptions {
	options.enabledSetters = map[string]bool{"Operations": false}
	return &options
}

type CreateOperationsOption func(*CreateOperationsOptions)

func (srv *VectorsDB) WithCreateOperationsOperations(v []interface{}) CreateOperationsOption {
	return func(o *CreateOperationsOptions) {
		o.Operations = v
		o.enabledSetters["Operations"] = true
	}
}

// CreateOperations create multiple operations in a single transaction.
func (srv *VectorsDB) CreateOperations(TransactionId string, optionalSetters ...CreateOperationsOption) (*models.Transaction, error) {
	r := strings.NewReplacer("{transactionId}", client.EncodePath(TransactionId))
	path := r.Replace("/vectorsdb/transactions/{transactionId}/operations")
	options := CreateOperationsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Operations"] {
		params["operations"] = options.Operations
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

		parsed := models.Transaction{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Transaction
	parsed, ok := resp.Result.(models.Transaction)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Get get a database by its unique ID. This endpoint response returns a JSON
// object with the database metadata.
func (srv *VectorsDB) Get(DatabaseId string) (*models.Database, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/vectorsdb/{databaseId}")
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

		parsed := models.Database{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Database
	parsed, ok := resp.Result.(models.Database)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateOptions struct {
	Enabled        bool
	Specification  string
	Replicas       int
	SyncMode       string
	enabledSetters map[string]bool
}

func (options UpdateOptions) New() *UpdateOptions {
	options.enabledSetters = map[string]bool{"Enabled": false, "Specification": false, "Replicas": false, "SyncMode": false}
	return &options
}

type UpdateOption func(*UpdateOptions)

func (srv *VectorsDB) WithUpdateEnabled(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *VectorsDB) WithUpdateSpecification(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Specification = v
		o.enabledSetters["Specification"] = true
	}
}
func (srv *VectorsDB) WithUpdateReplicas(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.Replicas = v
		o.enabledSetters["Replicas"] = true
	}
}
func (srv *VectorsDB) WithUpdateSyncMode(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.SyncMode = v
		o.enabledSetters["SyncMode"] = true
	}
}

// Update update a database by its unique ID.
func (srv *VectorsDB) Update(DatabaseId string, Name string, optionalSetters ...UpdateOption) (*models.Database, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/vectorsdb/{databaseId}")
	options := UpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Database{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Database
	parsed, ok := resp.Result.(models.Database)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Delete delete a database by its unique ID. Only API keys with with
// databases.write scope can delete a database.
func (srv *VectorsDB) Delete(DatabaseId string) (*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/vectorsdb/{databaseId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"

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

type ListCollectionsOptions struct {
	Queries        []string
	Search         string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListCollectionsOptions) New() *ListCollectionsOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Search": false, "Total": false}
	return &options
}

type ListCollectionsOption func(*ListCollectionsOptions)

func (srv *VectorsDB) WithListCollectionsQueries(v []string) ListCollectionsOption {
	return func(o *ListCollectionsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *VectorsDB) WithListCollectionsSearch(v string) ListCollectionsOption {
	return func(o *ListCollectionsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *VectorsDB) WithListCollectionsTotal(v bool) ListCollectionsOption {
	return func(o *ListCollectionsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListCollections get a list of all collections that belong to the provided
// databaseId. You can use the search parameter to filter your results.
func (srv *VectorsDB) ListCollections(DatabaseId string, optionalSetters ...ListCollectionsOption) (*models.VectorsdbCollectionList, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/vectorsdb/{databaseId}/collections")
	options := ListCollectionsOptions{}.New()
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

		parsed := models.VectorsdbCollectionList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VectorsdbCollectionList
	parsed, ok := resp.Result.(models.VectorsdbCollectionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateCollectionOptions struct {
	Permissions      []string
	DocumentSecurity bool
	Enabled          bool
	enabledSetters   map[string]bool
}

func (options CreateCollectionOptions) New() *CreateCollectionOptions {
	options.enabledSetters = map[string]bool{"Permissions": false, "DocumentSecurity": false, "Enabled": false}
	return &options
}

type CreateCollectionOption func(*CreateCollectionOptions)

func (srv *VectorsDB) WithCreateCollectionPermissions(v []string) CreateCollectionOption {
	return func(o *CreateCollectionOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *VectorsDB) WithCreateCollectionDocumentSecurity(v bool) CreateCollectionOption {
	return func(o *CreateCollectionOptions) {
		o.DocumentSecurity = v
		o.enabledSetters["DocumentSecurity"] = true
	}
}
func (srv *VectorsDB) WithCreateCollectionEnabled(v bool) CreateCollectionOption {
	return func(o *CreateCollectionOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}

// CreateCollection create a new Collection. Before using this route, you
// should create a new database resource using either a [server
// integration](https://appwrite.io/docs/server/databases#documentsDBCreateCollection)
// API or directly from your database console.
func (srv *VectorsDB) CreateCollection(DatabaseId string, CollectionId string, Name string, Dimension int, optionalSetters ...CreateCollectionOption) (*models.VectorsdbCollection, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/vectorsdb/{databaseId}/collections")
	options := CreateCollectionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["collectionId"] = CollectionId
	params["name"] = Name
	params["dimension"] = Dimension
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["DocumentSecurity"] {
		params["documentSecurity"] = options.DocumentSecurity
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

		parsed := models.VectorsdbCollection{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VectorsdbCollection
	parsed, ok := resp.Result.(models.VectorsdbCollection)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetCollection get a collection by its unique ID. This endpoint response
// returns a JSON object with the collection metadata.
func (srv *VectorsDB) GetCollection(DatabaseId string, CollectionId string) (*models.VectorsdbCollection, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}")
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

		parsed := models.VectorsdbCollection{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VectorsdbCollection
	parsed, ok := resp.Result.(models.VectorsdbCollection)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateCollectionOptions struct {
	Dimension        int
	Permissions      []string
	DocumentSecurity bool
	Enabled          bool
	enabledSetters   map[string]bool
}

func (options UpdateCollectionOptions) New() *UpdateCollectionOptions {
	options.enabledSetters = map[string]bool{"Dimension": false, "Permissions": false, "DocumentSecurity": false, "Enabled": false}
	return &options
}

type UpdateCollectionOption func(*UpdateCollectionOptions)

func (srv *VectorsDB) WithUpdateCollectionDimension(v int) UpdateCollectionOption {
	return func(o *UpdateCollectionOptions) {
		o.Dimension = v
		o.enabledSetters["Dimension"] = true
	}
}
func (srv *VectorsDB) WithUpdateCollectionPermissions(v []string) UpdateCollectionOption {
	return func(o *UpdateCollectionOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *VectorsDB) WithUpdateCollectionDocumentSecurity(v bool) UpdateCollectionOption {
	return func(o *UpdateCollectionOptions) {
		o.DocumentSecurity = v
		o.enabledSetters["DocumentSecurity"] = true
	}
}
func (srv *VectorsDB) WithUpdateCollectionEnabled(v bool) UpdateCollectionOption {
	return func(o *UpdateCollectionOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}

// UpdateCollection update a collection by its unique ID.
func (srv *VectorsDB) UpdateCollection(DatabaseId string, CollectionId string, Name string, optionalSetters ...UpdateCollectionOption) (*models.VectorsdbCollection, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}")
	options := UpdateCollectionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	if options.enabledSetters["Dimension"] {
		params["dimension"] = options.Dimension
	}
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["DocumentSecurity"] {
		params["documentSecurity"] = options.DocumentSecurity
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
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

		parsed := models.VectorsdbCollection{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VectorsdbCollection
	parsed, ok := resp.Result.(models.VectorsdbCollection)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteCollection delete a collection by its unique ID. Only users with
// write permissions have access to delete this resource.
func (srv *VectorsDB) DeleteCollection(DatabaseId string, CollectionId string) (*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"

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

type ListDocumentsOptions struct {
	Queries        []string
	TransactionId  string
	Total          bool
	Ttl            int
	enabledSetters map[string]bool
}

func (options ListDocumentsOptions) New() *ListDocumentsOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "TransactionId": false, "Total": false, "Ttl": false}
	return &options
}

type ListDocumentsOption func(*ListDocumentsOptions)

func (srv *VectorsDB) WithListDocumentsQueries(v []string) ListDocumentsOption {
	return func(o *ListDocumentsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *VectorsDB) WithListDocumentsTransactionId(v string) ListDocumentsOption {
	return func(o *ListDocumentsOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}
func (srv *VectorsDB) WithListDocumentsTotal(v bool) ListDocumentsOption {
	return func(o *ListDocumentsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
func (srv *VectorsDB) WithListDocumentsTtl(v int) ListDocumentsOption {
	return func(o *ListDocumentsOptions) {
		o.Ttl = v
		o.enabledSetters["Ttl"] = true
	}
}

// ListDocuments get a list of all the user's documents in a given collection.
// You can use the query params to filter your results.
func (srv *VectorsDB) ListDocuments(DatabaseId string, CollectionId string, optionalSetters ...ListDocumentsOption) (*models.DocumentList, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/documents")
	options := ListDocumentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	if options.enabledSetters["Ttl"] {
		params["ttl"] = options.Ttl
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

		parsed := models.DocumentList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DocumentList
	parsed, ok := resp.Result.(models.DocumentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateDocumentOptions struct {
	Permissions    []string
	enabledSetters map[string]bool
}

func (options CreateDocumentOptions) New() *CreateDocumentOptions {
	options.enabledSetters = map[string]bool{"Permissions": false}
	return &options
}

type CreateDocumentOption func(*CreateDocumentOptions)

func (srv *VectorsDB) WithCreateDocumentPermissions(v []string) CreateDocumentOption {
	return func(o *CreateDocumentOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}

// CreateDocument create a new Document. Before using this route, you should
// create a new collection resource using either a [server
// integration](https://appwrite.io/docs/server/databases#documentsDBCreateCollection)
// API or directly from your database console.
func (srv *VectorsDB) CreateDocument(DatabaseId string, CollectionId string, DocumentId string, Data interface{}, optionalSetters ...CreateDocumentOption) (*models.Document, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/documents")
	options := CreateDocumentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["documentId"] = DocumentId
	params["data"] = Data
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
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

		parsed := models.Document{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Document
	parsed, ok := resp.Result.(models.Document)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreateDocuments create new Documents. Before using this route, you should
// create a new collection resource using either a [server
// integration](https://appwrite.io/docs/server/databases#documentsDBCreateCollection)
// API or directly from your database console.
func (srv *VectorsDB) CreateDocuments(DatabaseId string, CollectionId string, Documents []interface{}) (*models.DocumentList, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/documents")
	params := map[string]interface{}{}
	params["documents"] = Documents
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

		parsed := models.DocumentList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DocumentList
	parsed, ok := resp.Result.(models.DocumentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpsertDocumentsOptions struct {
	TransactionId  string
	enabledSetters map[string]bool
}

func (options UpsertDocumentsOptions) New() *UpsertDocumentsOptions {
	options.enabledSetters = map[string]bool{"TransactionId": false}
	return &options
}

type UpsertDocumentsOption func(*UpsertDocumentsOptions)

func (srv *VectorsDB) WithUpsertDocumentsTransactionId(v string) UpsertDocumentsOption {
	return func(o *UpsertDocumentsOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}

// UpsertDocuments create or update Documents. Before using this route, you
// should create a new collection resource using either a [server
// integration](https://appwrite.io/docs/server/databases#documentsDBCreateCollection)
// API or directly from your database console.
func (srv *VectorsDB) UpsertDocuments(DatabaseId string, CollectionId string, Documents []interface{}, optionalSetters ...UpsertDocumentsOption) (*models.DocumentList, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/documents")
	options := UpsertDocumentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["documents"] = Documents
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
	}
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

		parsed := models.DocumentList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DocumentList
	parsed, ok := resp.Result.(models.DocumentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateDocumentsOptions struct {
	Data           interface{}
	Queries        []string
	TransactionId  string
	enabledSetters map[string]bool
}

func (options UpdateDocumentsOptions) New() *UpdateDocumentsOptions {
	options.enabledSetters = map[string]bool{"Data": false, "Queries": false, "TransactionId": false}
	return &options
}

type UpdateDocumentsOption func(*UpdateDocumentsOptions)

func (srv *VectorsDB) WithUpdateDocumentsData(v interface{}) UpdateDocumentsOption {
	return func(o *UpdateDocumentsOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func (srv *VectorsDB) WithUpdateDocumentsQueries(v []string) UpdateDocumentsOption {
	return func(o *UpdateDocumentsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *VectorsDB) WithUpdateDocumentsTransactionId(v string) UpdateDocumentsOption {
	return func(o *UpdateDocumentsOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}

// UpdateDocuments update all documents that match your queries, if no queries
// are submitted then all documents are updated. You can pass only specific
// fields to be updated.
func (srv *VectorsDB) UpdateDocuments(DatabaseId string, CollectionId string, optionalSetters ...UpdateDocumentsOption) (*models.DocumentList, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/documents")
	options := UpdateDocumentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
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

		parsed := models.DocumentList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DocumentList
	parsed, ok := resp.Result.(models.DocumentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type DeleteDocumentsOptions struct {
	Queries        []string
	TransactionId  string
	enabledSetters map[string]bool
}

func (options DeleteDocumentsOptions) New() *DeleteDocumentsOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "TransactionId": false}
	return &options
}

type DeleteDocumentsOption func(*DeleteDocumentsOptions)

func (srv *VectorsDB) WithDeleteDocumentsQueries(v []string) DeleteDocumentsOption {
	return func(o *DeleteDocumentsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *VectorsDB) WithDeleteDocumentsTransactionId(v string) DeleteDocumentsOption {
	return func(o *DeleteDocumentsOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}

// DeleteDocuments bulk delete documents using queries, if no queries are
// passed then all documents are deleted.
func (srv *VectorsDB) DeleteDocuments(DatabaseId string, CollectionId string, optionalSetters ...DeleteDocumentsOption) (*models.DocumentList, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/documents")
	options := DeleteDocumentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
	}
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

		parsed := models.DocumentList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DocumentList
	parsed, ok := resp.Result.(models.DocumentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateQueryOptions struct {
	Queries        []string
	TransactionId  string
	Total          bool
	Ttl            int
	enabledSetters map[string]bool
}

func (options CreateQueryOptions) New() *CreateQueryOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "TransactionId": false, "Total": false, "Ttl": false}
	return &options
}

type CreateQueryOption func(*CreateQueryOptions)

func (srv *VectorsDB) WithCreateQueryQueries(v []string) CreateQueryOption {
	return func(o *CreateQueryOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *VectorsDB) WithCreateQueryTransactionId(v string) CreateQueryOption {
	return func(o *CreateQueryOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}
func (srv *VectorsDB) WithCreateQueryTotal(v bool) CreateQueryOption {
	return func(o *CreateQueryOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
func (srv *VectorsDB) WithCreateQueryTtl(v int) CreateQueryOption {
	return func(o *CreateQueryOptions) {
		o.Ttl = v
		o.enabledSetters["Ttl"] = true
	}
}

// CreateQuery get a list of all the user's documents in a given collection
// using a POST request. This behaves identically to the list documents
// endpoint but accepts the queries in the request body, allowing much larger
// `queries` arrays than can fit in a URL query string.
func (srv *VectorsDB) CreateQuery(DatabaseId string, CollectionId string, optionalSetters ...CreateQueryOption) (*models.DocumentList, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/documents/query")
	options := CreateQueryOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
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

		parsed := models.DocumentList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DocumentList
	parsed, ok := resp.Result.(models.DocumentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type GetDocumentOptions struct {
	Queries        []string
	TransactionId  string
	enabledSetters map[string]bool
}

func (options GetDocumentOptions) New() *GetDocumentOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "TransactionId": false}
	return &options
}

type GetDocumentOption func(*GetDocumentOptions)

func (srv *VectorsDB) WithGetDocumentQueries(v []string) GetDocumentOption {
	return func(o *GetDocumentOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *VectorsDB) WithGetDocumentTransactionId(v string) GetDocumentOption {
	return func(o *GetDocumentOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}

// GetDocument get a document by its unique ID. This endpoint response returns
// a JSON object with the document data.
func (srv *VectorsDB) GetDocument(DatabaseId string, CollectionId string, DocumentId string, optionalSetters ...GetDocumentOption) (*models.Document, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId), "{documentId}", client.EncodePath(DocumentId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/documents/{documentId}")
	options := GetDocumentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
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

		parsed := models.Document{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Document
	parsed, ok := resp.Result.(models.Document)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpsertDocumentOptions struct {
	Data           interface{}
	Permissions    []string
	TransactionId  string
	enabledSetters map[string]bool
}

func (options UpsertDocumentOptions) New() *UpsertDocumentOptions {
	options.enabledSetters = map[string]bool{"Data": false, "Permissions": false, "TransactionId": false}
	return &options
}

type UpsertDocumentOption func(*UpsertDocumentOptions)

func (srv *VectorsDB) WithUpsertDocumentData(v interface{}) UpsertDocumentOption {
	return func(o *UpsertDocumentOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func (srv *VectorsDB) WithUpsertDocumentPermissions(v []string) UpsertDocumentOption {
	return func(o *UpsertDocumentOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *VectorsDB) WithUpsertDocumentTransactionId(v string) UpsertDocumentOption {
	return func(o *UpsertDocumentOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}

// UpsertDocument create or update a Document. Before using this route, you
// should create a new collection resource using either a [server
// integration](https://appwrite.io/docs/server/databases#documentsDBCreateCollection)
// API or directly from your database console.
func (srv *VectorsDB) UpsertDocument(DatabaseId string, CollectionId string, DocumentId string, optionalSetters ...UpsertDocumentOption) (*models.Document, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId), "{documentId}", client.EncodePath(DocumentId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/documents/{documentId}")
	options := UpsertDocumentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
	}
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

		parsed := models.Document{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Document
	parsed, ok := resp.Result.(models.Document)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateDocumentOptions struct {
	Data           interface{}
	Permissions    []string
	TransactionId  string
	enabledSetters map[string]bool
}

func (options UpdateDocumentOptions) New() *UpdateDocumentOptions {
	options.enabledSetters = map[string]bool{"Data": false, "Permissions": false, "TransactionId": false}
	return &options
}

type UpdateDocumentOption func(*UpdateDocumentOptions)

func (srv *VectorsDB) WithUpdateDocumentData(v interface{}) UpdateDocumentOption {
	return func(o *UpdateDocumentOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func (srv *VectorsDB) WithUpdateDocumentPermissions(v []string) UpdateDocumentOption {
	return func(o *UpdateDocumentOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *VectorsDB) WithUpdateDocumentTransactionId(v string) UpdateDocumentOption {
	return func(o *UpdateDocumentOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}

// UpdateDocument update a document by its unique ID. Using the patch method
// you can pass only specific fields that will get updated.
func (srv *VectorsDB) UpdateDocument(DatabaseId string, CollectionId string, DocumentId string, optionalSetters ...UpdateDocumentOption) (*models.Document, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId), "{documentId}", client.EncodePath(DocumentId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/documents/{documentId}")
	options := UpdateDocumentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
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

		parsed := models.Document{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Document
	parsed, ok := resp.Result.(models.Document)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type DeleteDocumentOptions struct {
	TransactionId  string
	enabledSetters map[string]bool
}

func (options DeleteDocumentOptions) New() *DeleteDocumentOptions {
	options.enabledSetters = map[string]bool{"TransactionId": false}
	return &options
}

type DeleteDocumentOption func(*DeleteDocumentOptions)

func (srv *VectorsDB) WithDeleteDocumentTransactionId(v string) DeleteDocumentOption {
	return func(o *DeleteDocumentOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}

// DeleteDocument delete a document by its unique ID.
func (srv *VectorsDB) DeleteDocument(DatabaseId string, CollectionId string, DocumentId string, optionalSetters ...DeleteDocumentOption) (*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId), "{documentId}", client.EncodePath(DocumentId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/documents/{documentId}")
	options := DeleteDocumentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"

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

type ListIndexesOptions struct {
	Queries        []string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListIndexesOptions) New() *ListIndexesOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Total": false}
	return &options
}

type ListIndexesOption func(*ListIndexesOptions)

func (srv *VectorsDB) WithListIndexesQueries(v []string) ListIndexesOption {
	return func(o *ListIndexesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *VectorsDB) WithListIndexesTotal(v bool) ListIndexesOption {
	return func(o *ListIndexesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListIndexes list indexes in the collection.
func (srv *VectorsDB) ListIndexes(DatabaseId string, CollectionId string, optionalSetters ...ListIndexesOption) (*models.IndexList, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/indexes")
	options := ListIndexesOptions{}.New()
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

		parsed := models.IndexList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.IndexList
	parsed, ok := resp.Result.(models.IndexList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateIndexOptions struct {
	Orders         []string
	Lengths        []int
	enabledSetters map[string]bool
}

func (options CreateIndexOptions) New() *CreateIndexOptions {
	options.enabledSetters = map[string]bool{"Orders": false, "Lengths": false}
	return &options
}

type CreateIndexOption func(*CreateIndexOptions)

func (srv *VectorsDB) WithCreateIndexOrders(v []string) CreateIndexOption {
	return func(o *CreateIndexOptions) {
		o.Orders = v
		o.enabledSetters["Orders"] = true
	}
}
func (srv *VectorsDB) WithCreateIndexLengths(v []int) CreateIndexOption {
	return func(o *CreateIndexOptions) {
		o.Lengths = v
		o.enabledSetters["Lengths"] = true
	}
}

// CreateIndex creates an index on the attributes listed. Your index should
// include all the attributes you will query in a single request.
// Attributes can be `key`, `fulltext`, and `unique`.
func (srv *VectorsDB) CreateIndex(DatabaseId string, CollectionId string, Key string, Type string, Attributes []string, optionalSetters ...CreateIndexOption) (*models.Index, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/indexes")
	options := CreateIndexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["key"] = Key
	params["type"] = Type
	params["attributes"] = Attributes
	if options.enabledSetters["Orders"] {
		params["orders"] = options.Orders
	}
	if options.enabledSetters["Lengths"] {
		params["lengths"] = options.Lengths
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

		parsed := models.Index{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Index
	parsed, ok := resp.Result.(models.Index)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetIndex get index by ID.
func (srv *VectorsDB) GetIndex(DatabaseId string, CollectionId string, Key string) (*models.Index, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId), "{key}", client.EncodePath(Key))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/indexes/{key}")
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

		parsed := models.Index{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Index
	parsed, ok := resp.Result.(models.Index)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteIndex delete an index.
func (srv *VectorsDB) DeleteIndex(DatabaseId string, CollectionId string, Key string) (*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId), "{collectionId}", client.EncodePath(CollectionId), "{key}", client.EncodePath(Key))
	path := r.Replace("/vectorsdb/{databaseId}/collections/{collectionId}/indexes/{key}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"

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

type CreateFailoverOptions struct {
	TargetReplicaId string
	enabledSetters  map[string]bool
}

func (options CreateFailoverOptions) New() *CreateFailoverOptions {
	options.enabledSetters = map[string]bool{"TargetReplicaId": false}
	return &options
}

type CreateFailoverOption func(*CreateFailoverOptions)

func (srv *VectorsDB) WithCreateFailoverTargetReplicaId(v string) CreateFailoverOption {
	return func(o *CreateFailoverOptions) {
		o.TargetReplicaId = v
		o.enabledSetters["TargetReplicaId"] = true
	}
}

// CreateFailover trigger a manual failover for a dedicated database with high
// availability enabled. Promotes a replica to primary. The failover runs
// asynchronously; poll the database document for status updates. A database
// left mid-operation also accepts this call as a repair once nothing is
// driving the operation it is stuck in. Repairing a failover that did not
// finish, a `failed` database, a stranded upgrade or migrate, or a stranded
// compute resize additionally requires `targetReplicaId` to name the member
// to promote, because the default target may be the member that operation
// already promoted.
func (srv *VectorsDB) CreateFailover(DatabaseId string, optionalSetters ...CreateFailoverOption) (*models.DedicatedDatabase, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/vectorsdb/{databaseId}/failovers")
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

func (srv *VectorsDB) WithListOperationsStatus(v string) ListOperationsOption {
	return func(o *ListOperationsOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *VectorsDB) WithListOperationsLimit(v int) ListOperationsOption {
	return func(o *ListOperationsOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *VectorsDB) WithListOperationsOffset(v int) ListOperationsOption {
	return func(o *ListOperationsOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}

// ListOperations list the lifecycle operations recorded for a dedicated
// database, newest first. Every provision, update, restore, backup and
// replication action is recorded here with its outcome, including an attempt
// that was abandoned because another worker took over the database.
func (srv *VectorsDB) ListOperations(DatabaseId string, optionalSetters ...ListOperationsOption) (*models.DedicatedDatabaseOperationList, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/vectorsdb/{databaseId}/operations")
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

// GetReplicas get high availability status for a dedicated database. Returns
// replica statuses, replication lag, and sync mode.
func (srv *VectorsDB) GetReplicas(DatabaseId string) (*models.DedicatedDatabaseReplicas, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/vectorsdb/{databaseId}/replicas")
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

// GetStatus get real-time health and status information for a dedicated
// database. Returns health status, readiness, uptime, connection info,
// replica status, and volume information.
func (srv *VectorsDB) GetStatus(DatabaseId string) (*models.DatabaseStatus, error) {
	r := strings.NewReplacer("{databaseId}", client.EncodePath(DatabaseId))
	path := r.Replace("/vectorsdb/{databaseId}/status")
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
