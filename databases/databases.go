package databases

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// Databases service
type Databases struct {
	client client.Client
}

func New(clt client.Client) *Databases {
	return &Databases{
		client: clt,
	}
}

type ListOptions struct {
	Queries []string
	Search string
	enabledSetters map[string]bool
}
func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
	}
	return &options
}
type ListOption func(*ListOptions)
func (srv *Databases) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Databases) WithListSearch(v string) ListOption {
	return func(o *ListOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
	
// List get a list of all databases from the current Appwrite project. You can
// use the search parameter to filter your results.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.list` instead.
func (srv *Databases) List(optionalSetters ...ListOption)(*models.DatabaseList, error) {
	path := "/databases"
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
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

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
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{
		"Enabled": false,
	}
	return &options
}
type CreateOption func(*CreateOptions)
func (srv *Databases) WithCreateEnabled(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
					
// Create create a new Database.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.create` instead.
func (srv *Databases) Create(DatabaseId string, Name string, optionalSetters ...CreateOption)(*models.Database, error) {
	path := "/databases"
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
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

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
type ListTransactionsOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListTransactionsOptions) New() *ListTransactionsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListTransactionsOption func(*ListTransactionsOptions)
func (srv *Databases) WithListTransactionsQueries(v []string) ListTransactionsOption {
	return func(o *ListTransactionsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
	
// ListTransactions list transactions across all databases.
func (srv *Databases) ListTransactions(optionalSetters ...ListTransactionsOption)(*models.TransactionList, error) {
	path := "/databases/transactions"
	options := ListTransactionsOptions{}.New()
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
	Ttl int
	enabledSetters map[string]bool
}
func (options CreateTransactionOptions) New() *CreateTransactionOptions {
	options.enabledSetters = map[string]bool{
		"Ttl": false,
	}
	return &options
}
type CreateTransactionOption func(*CreateTransactionOptions)
func (srv *Databases) WithCreateTransactionTtl(v int) CreateTransactionOption {
	return func(o *CreateTransactionOptions) {
		o.Ttl = v
		o.enabledSetters["Ttl"] = true
	}
}
	
// CreateTransaction create a new transaction.
func (srv *Databases) CreateTransaction(optionalSetters ...CreateTransactionOption)(*models.Transaction, error) {
	path := "/databases/transactions"
	options := CreateTransactionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Ttl"] {
		params["ttl"] = options.Ttl
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
func (srv *Databases) GetTransaction(TransactionId string)(*models.Transaction, error) {
	r := strings.NewReplacer("{transactionId}", TransactionId)
	path := r.Replace("/databases/transactions/{transactionId}")
	params := map[string]interface{}{}
	params["transactionId"] = TransactionId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

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
	Commit bool
	Rollback bool
	enabledSetters map[string]bool
}
func (options UpdateTransactionOptions) New() *UpdateTransactionOptions {
	options.enabledSetters = map[string]bool{
		"Commit": false,
		"Rollback": false,
	}
	return &options
}
type UpdateTransactionOption func(*UpdateTransactionOptions)
func (srv *Databases) WithUpdateTransactionCommit(v bool) UpdateTransactionOption {
	return func(o *UpdateTransactionOptions) {
		o.Commit = v
		o.enabledSetters["Commit"] = true
	}
}
func (srv *Databases) WithUpdateTransactionRollback(v bool) UpdateTransactionOption {
	return func(o *UpdateTransactionOptions) {
		o.Rollback = v
		o.enabledSetters["Rollback"] = true
	}
}
			
// UpdateTransaction update a transaction, to either commit or roll back its
// operations.
func (srv *Databases) UpdateTransaction(TransactionId string, optionalSetters ...UpdateTransactionOption)(*models.Transaction, error) {
	r := strings.NewReplacer("{transactionId}", TransactionId)
	path := r.Replace("/databases/transactions/{transactionId}")
	options := UpdateTransactionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["transactionId"] = TransactionId
	if options.enabledSetters["Commit"] {
		params["commit"] = options.Commit
	}
	if options.enabledSetters["Rollback"] {
		params["rollback"] = options.Rollback
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
func (srv *Databases) DeleteTransaction(TransactionId string)(*interface{}, error) {
	r := strings.NewReplacer("{transactionId}", TransactionId)
	path := r.Replace("/databases/transactions/{transactionId}")
	params := map[string]interface{}{}
	params["transactionId"] = TransactionId
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
type CreateOperationsOptions struct {
	Operations []interface{}
	enabledSetters map[string]bool
}
func (options CreateOperationsOptions) New() *CreateOperationsOptions {
	options.enabledSetters = map[string]bool{
		"Operations": false,
	}
	return &options
}
type CreateOperationsOption func(*CreateOperationsOptions)
func (srv *Databases) WithCreateOperationsOperations(v []interface{}) CreateOperationsOption {
	return func(o *CreateOperationsOptions) {
		o.Operations = v
		o.enabledSetters["Operations"] = true
	}
}
			
// CreateOperations create multiple operations in a single transaction.
func (srv *Databases) CreateOperations(TransactionId string, optionalSetters ...CreateOperationsOption)(*models.Transaction, error) {
	r := strings.NewReplacer("{transactionId}", TransactionId)
	path := r.Replace("/databases/transactions/{transactionId}/operations")
	options := CreateOperationsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["transactionId"] = TransactionId
	if options.enabledSetters["Operations"] {
		params["operations"] = options.Operations
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
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.get` instead.
func (srv *Databases) Get(DatabaseId string)(*models.Database, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

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
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOptions) New() *UpdateOptions {
	options.enabledSetters = map[string]bool{
		"Enabled": false,
	}
	return &options
}
type UpdateOption func(*UpdateOptions)
func (srv *Databases) WithUpdateEnabled(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
					
// Update update a database by its unique ID.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.update` instead.
func (srv *Databases) Update(DatabaseId string, Name string, optionalSetters ...UpdateOption)(*models.Database, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}")
	options := UpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["name"] = Name
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.delete` instead.
func (srv *Databases) Delete(DatabaseId string)(*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
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
type ListCollectionsOptions struct {
	Queries []string
	Search string
	enabledSetters map[string]bool
}
func (options ListCollectionsOptions) New() *ListCollectionsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
	}
	return &options
}
type ListCollectionsOption func(*ListCollectionsOptions)
func (srv *Databases) WithListCollectionsQueries(v []string) ListCollectionsOption {
	return func(o *ListCollectionsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Databases) WithListCollectionsSearch(v string) ListCollectionsOption {
	return func(o *ListCollectionsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
			
// ListCollections get a list of all collections that belong to the provided
// databaseId. You can use the search parameter to filter your results.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.listTables` instead.
func (srv *Databases) ListCollections(DatabaseId string, optionalSetters ...ListCollectionsOption)(*models.CollectionList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}/collections")
	options := ListCollectionsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.CollectionList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.CollectionList
	parsed, ok := resp.Result.(models.CollectionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateCollectionOptions struct {
	Permissions []string
	DocumentSecurity bool
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateCollectionOptions) New() *CreateCollectionOptions {
	options.enabledSetters = map[string]bool{
		"Permissions": false,
		"DocumentSecurity": false,
		"Enabled": false,
	}
	return &options
}
type CreateCollectionOption func(*CreateCollectionOptions)
func (srv *Databases) WithCreateCollectionPermissions(v []string) CreateCollectionOption {
	return func(o *CreateCollectionOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *Databases) WithCreateCollectionDocumentSecurity(v bool) CreateCollectionOption {
	return func(o *CreateCollectionOptions) {
		o.DocumentSecurity = v
		o.enabledSetters["DocumentSecurity"] = true
	}
}
func (srv *Databases) WithCreateCollectionEnabled(v bool) CreateCollectionOption {
	return func(o *CreateCollectionOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
							
// CreateCollection create a new Collection. Before using this route, you
// should create a new database resource using either a [server
// integration](https://appwrite.io/docs/server/databases#databasesCreateCollection)
// API or directly from your database console.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createTable` instead.
func (srv *Databases) CreateCollection(DatabaseId string, CollectionId string, Name string, optionalSetters ...CreateCollectionOption)(*models.Collection, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/databases/{databaseId}/collections")
	options := CreateCollectionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["name"] = Name
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["DocumentSecurity"] {
		params["documentSecurity"] = options.DocumentSecurity
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

		parsed := models.Collection{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Collection
	parsed, ok := resp.Result.(models.Collection)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetCollection get a collection by its unique ID. This endpoint response
// returns a JSON object with the collection metadata.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.getTable` instead.
func (srv *Databases) GetCollection(DatabaseId string, CollectionId string)(*models.Collection, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Collection{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Collection
	parsed, ok := resp.Result.(models.Collection)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateCollectionOptions struct {
	Permissions []string
	DocumentSecurity bool
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateCollectionOptions) New() *UpdateCollectionOptions {
	options.enabledSetters = map[string]bool{
		"Permissions": false,
		"DocumentSecurity": false,
		"Enabled": false,
	}
	return &options
}
type UpdateCollectionOption func(*UpdateCollectionOptions)
func (srv *Databases) WithUpdateCollectionPermissions(v []string) UpdateCollectionOption {
	return func(o *UpdateCollectionOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *Databases) WithUpdateCollectionDocumentSecurity(v bool) UpdateCollectionOption {
	return func(o *UpdateCollectionOptions) {
		o.DocumentSecurity = v
		o.enabledSetters["DocumentSecurity"] = true
	}
}
func (srv *Databases) WithUpdateCollectionEnabled(v bool) UpdateCollectionOption {
	return func(o *UpdateCollectionOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
							
// UpdateCollection update a collection by its unique ID.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updateTable` instead.
func (srv *Databases) UpdateCollection(DatabaseId string, CollectionId string, Name string, optionalSetters ...UpdateCollectionOption)(*models.Collection, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}")
	options := UpdateCollectionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["name"] = Name
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["DocumentSecurity"] {
		params["documentSecurity"] = options.DocumentSecurity
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Collection{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Collection
	parsed, ok := resp.Result.(models.Collection)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// DeleteCollection delete a collection by its unique ID. Only users with
// write permissions have access to delete this resource.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.deleteTable` instead.
func (srv *Databases) DeleteCollection(DatabaseId string, CollectionId string)(*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
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
type ListAttributesOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListAttributesOptions) New() *ListAttributesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListAttributesOption func(*ListAttributesOptions)
func (srv *Databases) WithListAttributesQueries(v []string) ListAttributesOption {
	return func(o *ListAttributesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
					
// ListAttributes list attributes in the collection.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.listColumns` instead.
func (srv *Databases) ListAttributes(DatabaseId string, CollectionId string, optionalSetters ...ListAttributesOption)(*models.AttributeList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes")
	options := ListAttributesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
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

		parsed := models.AttributeList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeList
	parsed, ok := resp.Result.(models.AttributeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateBooleanAttributeOptions struct {
	Default bool
	Array bool
	enabledSetters map[string]bool
}
func (options CreateBooleanAttributeOptions) New() *CreateBooleanAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateBooleanAttributeOption func(*CreateBooleanAttributeOptions)
func (srv *Databases) WithCreateBooleanAttributeDefault(v bool) CreateBooleanAttributeOption {
	return func(o *CreateBooleanAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *Databases) WithCreateBooleanAttributeArray(v bool) CreateBooleanAttributeOption {
	return func(o *CreateBooleanAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
									
// CreateBooleanAttribute create a boolean attribute.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createBooleanColumn` instead.
func (srv *Databases) CreateBooleanAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateBooleanAttributeOption)(*models.AttributeBoolean, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/boolean")
	options := CreateBooleanAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
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

		parsed := models.AttributeBoolean{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeBoolean
	parsed, ok := resp.Result.(models.AttributeBoolean)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateBooleanAttributeOptions struct {
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateBooleanAttributeOptions) New() *UpdateBooleanAttributeOptions {
	options.enabledSetters = map[string]bool{
		"NewKey": false,
	}
	return &options
}
type UpdateBooleanAttributeOption func(*UpdateBooleanAttributeOptions)
func (srv *Databases) WithUpdateBooleanAttributeNewKey(v string) UpdateBooleanAttributeOption {
	return func(o *UpdateBooleanAttributeOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateBooleanAttribute update a boolean attribute. Changing the `default`
// value will not update already existing documents.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updateBooleanColumn` instead.
func (srv *Databases) UpdateBooleanAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default bool, optionalSetters ...UpdateBooleanAttributeOption)(*models.AttributeBoolean, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/boolean/{key}")
	options := UpdateBooleanAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = Default
	if options.enabledSetters["NewKey"] {
		params["newKey"] = options.NewKey
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

		parsed := models.AttributeBoolean{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeBoolean
	parsed, ok := resp.Result.(models.AttributeBoolean)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateDatetimeAttributeOptions struct {
	Default string
	Array bool
	enabledSetters map[string]bool
}
func (options CreateDatetimeAttributeOptions) New() *CreateDatetimeAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateDatetimeAttributeOption func(*CreateDatetimeAttributeOptions)
func (srv *Databases) WithCreateDatetimeAttributeDefault(v string) CreateDatetimeAttributeOption {
	return func(o *CreateDatetimeAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *Databases) WithCreateDatetimeAttributeArray(v bool) CreateDatetimeAttributeOption {
	return func(o *CreateDatetimeAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
									
// CreateDatetimeAttribute create a date time attribute according to the ISO
// 8601 standard.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createDatetimeColumn` instead.
func (srv *Databases) CreateDatetimeAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateDatetimeAttributeOption)(*models.AttributeDatetime, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/datetime")
	options := CreateDatetimeAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
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

		parsed := models.AttributeDatetime{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeDatetime
	parsed, ok := resp.Result.(models.AttributeDatetime)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateDatetimeAttributeOptions struct {
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateDatetimeAttributeOptions) New() *UpdateDatetimeAttributeOptions {
	options.enabledSetters = map[string]bool{
		"NewKey": false,
	}
	return &options
}
type UpdateDatetimeAttributeOption func(*UpdateDatetimeAttributeOptions)
func (srv *Databases) WithUpdateDatetimeAttributeNewKey(v string) UpdateDatetimeAttributeOption {
	return func(o *UpdateDatetimeAttributeOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateDatetimeAttribute update a date time attribute. Changing the
// `default` value will not update already existing documents.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updateDatetimeColumn` instead.
func (srv *Databases) UpdateDatetimeAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string, optionalSetters ...UpdateDatetimeAttributeOption)(*models.AttributeDatetime, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/datetime/{key}")
	options := UpdateDatetimeAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = Default
	if options.enabledSetters["NewKey"] {
		params["newKey"] = options.NewKey
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

		parsed := models.AttributeDatetime{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeDatetime
	parsed, ok := resp.Result.(models.AttributeDatetime)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateEmailAttributeOptions struct {
	Default string
	Array bool
	enabledSetters map[string]bool
}
func (options CreateEmailAttributeOptions) New() *CreateEmailAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateEmailAttributeOption func(*CreateEmailAttributeOptions)
func (srv *Databases) WithCreateEmailAttributeDefault(v string) CreateEmailAttributeOption {
	return func(o *CreateEmailAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *Databases) WithCreateEmailAttributeArray(v bool) CreateEmailAttributeOption {
	return func(o *CreateEmailAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
									
// CreateEmailAttribute create an email attribute.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createEmailColumn` instead.
func (srv *Databases) CreateEmailAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateEmailAttributeOption)(*models.AttributeEmail, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/email")
	options := CreateEmailAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
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

		parsed := models.AttributeEmail{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeEmail
	parsed, ok := resp.Result.(models.AttributeEmail)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateEmailAttributeOptions struct {
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateEmailAttributeOptions) New() *UpdateEmailAttributeOptions {
	options.enabledSetters = map[string]bool{
		"NewKey": false,
	}
	return &options
}
type UpdateEmailAttributeOption func(*UpdateEmailAttributeOptions)
func (srv *Databases) WithUpdateEmailAttributeNewKey(v string) UpdateEmailAttributeOption {
	return func(o *UpdateEmailAttributeOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateEmailAttribute update an email attribute. Changing the `default`
// value will not update already existing documents.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updateEmailColumn` instead.
func (srv *Databases) UpdateEmailAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string, optionalSetters ...UpdateEmailAttributeOption)(*models.AttributeEmail, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/email/{key}")
	options := UpdateEmailAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = Default
	if options.enabledSetters["NewKey"] {
		params["newKey"] = options.NewKey
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

		parsed := models.AttributeEmail{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeEmail
	parsed, ok := resp.Result.(models.AttributeEmail)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateEnumAttributeOptions struct {
	Default string
	Array bool
	enabledSetters map[string]bool
}
func (options CreateEnumAttributeOptions) New() *CreateEnumAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateEnumAttributeOption func(*CreateEnumAttributeOptions)
func (srv *Databases) WithCreateEnumAttributeDefault(v string) CreateEnumAttributeOption {
	return func(o *CreateEnumAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *Databases) WithCreateEnumAttributeArray(v bool) CreateEnumAttributeOption {
	return func(o *CreateEnumAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
											
// CreateEnumAttribute create an enum attribute. The `elements` param acts as
// a white-list of accepted values for this attribute.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createEnumColumn` instead.
func (srv *Databases) CreateEnumAttribute(DatabaseId string, CollectionId string, Key string, Elements []string, Required bool, optionalSetters ...CreateEnumAttributeOption)(*models.AttributeEnum, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/enum")
	options := CreateEnumAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["elements"] = Elements
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
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

		parsed := models.AttributeEnum{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeEnum
	parsed, ok := resp.Result.(models.AttributeEnum)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateEnumAttributeOptions struct {
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateEnumAttributeOptions) New() *UpdateEnumAttributeOptions {
	options.enabledSetters = map[string]bool{
		"NewKey": false,
	}
	return &options
}
type UpdateEnumAttributeOption func(*UpdateEnumAttributeOptions)
func (srv *Databases) WithUpdateEnumAttributeNewKey(v string) UpdateEnumAttributeOption {
	return func(o *UpdateEnumAttributeOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
													
// UpdateEnumAttribute update an enum attribute. Changing the `default` value
// will not update already existing documents.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updateEnumColumn` instead.
func (srv *Databases) UpdateEnumAttribute(DatabaseId string, CollectionId string, Key string, Elements []string, Required bool, Default string, optionalSetters ...UpdateEnumAttributeOption)(*models.AttributeEnum, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/enum/{key}")
	options := UpdateEnumAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["elements"] = Elements
	params["required"] = Required
	params["default"] = Default
	if options.enabledSetters["NewKey"] {
		params["newKey"] = options.NewKey
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

		parsed := models.AttributeEnum{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeEnum
	parsed, ok := resp.Result.(models.AttributeEnum)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateFloatAttributeOptions struct {
	Min float64
	Max float64
	Default float64
	Array bool
	enabledSetters map[string]bool
}
func (options CreateFloatAttributeOptions) New() *CreateFloatAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Min": false,
		"Max": false,
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateFloatAttributeOption func(*CreateFloatAttributeOptions)
func (srv *Databases) WithCreateFloatAttributeMin(v float64) CreateFloatAttributeOption {
	return func(o *CreateFloatAttributeOptions) {
		o.Min = v
		o.enabledSetters["Min"] = true
	}
}
func (srv *Databases) WithCreateFloatAttributeMax(v float64) CreateFloatAttributeOption {
	return func(o *CreateFloatAttributeOptions) {
		o.Max = v
		o.enabledSetters["Max"] = true
	}
}
func (srv *Databases) WithCreateFloatAttributeDefault(v float64) CreateFloatAttributeOption {
	return func(o *CreateFloatAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *Databases) WithCreateFloatAttributeArray(v bool) CreateFloatAttributeOption {
	return func(o *CreateFloatAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
									
// CreateFloatAttribute create a float attribute. Optionally, minimum and
// maximum values can be provided.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createFloatColumn` instead.
func (srv *Databases) CreateFloatAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateFloatAttributeOption)(*models.AttributeFloat, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/float")
	options := CreateFloatAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Min"] {
		params["min"] = options.Min
	}
	if options.enabledSetters["Max"] {
		params["max"] = options.Max
	}
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
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

		parsed := models.AttributeFloat{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeFloat
	parsed, ok := resp.Result.(models.AttributeFloat)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateFloatAttributeOptions struct {
	Min float64
	Max float64
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateFloatAttributeOptions) New() *UpdateFloatAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Min": false,
		"Max": false,
		"NewKey": false,
	}
	return &options
}
type UpdateFloatAttributeOption func(*UpdateFloatAttributeOptions)
func (srv *Databases) WithUpdateFloatAttributeMin(v float64) UpdateFloatAttributeOption {
	return func(o *UpdateFloatAttributeOptions) {
		o.Min = v
		o.enabledSetters["Min"] = true
	}
}
func (srv *Databases) WithUpdateFloatAttributeMax(v float64) UpdateFloatAttributeOption {
	return func(o *UpdateFloatAttributeOptions) {
		o.Max = v
		o.enabledSetters["Max"] = true
	}
}
func (srv *Databases) WithUpdateFloatAttributeNewKey(v string) UpdateFloatAttributeOption {
	return func(o *UpdateFloatAttributeOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateFloatAttribute update a float attribute. Changing the `default` value
// will not update already existing documents.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updateFloatColumn` instead.
func (srv *Databases) UpdateFloatAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default float64, optionalSetters ...UpdateFloatAttributeOption)(*models.AttributeFloat, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/float/{key}")
	options := UpdateFloatAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = Default
	if options.enabledSetters["Min"] {
		params["min"] = options.Min
	}
	if options.enabledSetters["Max"] {
		params["max"] = options.Max
	}
	if options.enabledSetters["NewKey"] {
		params["newKey"] = options.NewKey
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

		parsed := models.AttributeFloat{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeFloat
	parsed, ok := resp.Result.(models.AttributeFloat)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateIntegerAttributeOptions struct {
	Min int
	Max int
	Default int
	Array bool
	enabledSetters map[string]bool
}
func (options CreateIntegerAttributeOptions) New() *CreateIntegerAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Min": false,
		"Max": false,
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateIntegerAttributeOption func(*CreateIntegerAttributeOptions)
func (srv *Databases) WithCreateIntegerAttributeMin(v int) CreateIntegerAttributeOption {
	return func(o *CreateIntegerAttributeOptions) {
		o.Min = v
		o.enabledSetters["Min"] = true
	}
}
func (srv *Databases) WithCreateIntegerAttributeMax(v int) CreateIntegerAttributeOption {
	return func(o *CreateIntegerAttributeOptions) {
		o.Max = v
		o.enabledSetters["Max"] = true
	}
}
func (srv *Databases) WithCreateIntegerAttributeDefault(v int) CreateIntegerAttributeOption {
	return func(o *CreateIntegerAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *Databases) WithCreateIntegerAttributeArray(v bool) CreateIntegerAttributeOption {
	return func(o *CreateIntegerAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
									
// CreateIntegerAttribute create an integer attribute. Optionally, minimum and
// maximum values can be provided.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createIntegerColumn` instead.
func (srv *Databases) CreateIntegerAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateIntegerAttributeOption)(*models.AttributeInteger, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/integer")
	options := CreateIntegerAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Min"] {
		params["min"] = options.Min
	}
	if options.enabledSetters["Max"] {
		params["max"] = options.Max
	}
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
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

		parsed := models.AttributeInteger{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeInteger
	parsed, ok := resp.Result.(models.AttributeInteger)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateIntegerAttributeOptions struct {
	Min int
	Max int
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateIntegerAttributeOptions) New() *UpdateIntegerAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Min": false,
		"Max": false,
		"NewKey": false,
	}
	return &options
}
type UpdateIntegerAttributeOption func(*UpdateIntegerAttributeOptions)
func (srv *Databases) WithUpdateIntegerAttributeMin(v int) UpdateIntegerAttributeOption {
	return func(o *UpdateIntegerAttributeOptions) {
		o.Min = v
		o.enabledSetters["Min"] = true
	}
}
func (srv *Databases) WithUpdateIntegerAttributeMax(v int) UpdateIntegerAttributeOption {
	return func(o *UpdateIntegerAttributeOptions) {
		o.Max = v
		o.enabledSetters["Max"] = true
	}
}
func (srv *Databases) WithUpdateIntegerAttributeNewKey(v string) UpdateIntegerAttributeOption {
	return func(o *UpdateIntegerAttributeOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateIntegerAttribute update an integer attribute. Changing the `default`
// value will not update already existing documents.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updateIntegerColumn` instead.
func (srv *Databases) UpdateIntegerAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default int, optionalSetters ...UpdateIntegerAttributeOption)(*models.AttributeInteger, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/integer/{key}")
	options := UpdateIntegerAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = Default
	if options.enabledSetters["Min"] {
		params["min"] = options.Min
	}
	if options.enabledSetters["Max"] {
		params["max"] = options.Max
	}
	if options.enabledSetters["NewKey"] {
		params["newKey"] = options.NewKey
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

		parsed := models.AttributeInteger{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeInteger
	parsed, ok := resp.Result.(models.AttributeInteger)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateIpAttributeOptions struct {
	Default string
	Array bool
	enabledSetters map[string]bool
}
func (options CreateIpAttributeOptions) New() *CreateIpAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateIpAttributeOption func(*CreateIpAttributeOptions)
func (srv *Databases) WithCreateIpAttributeDefault(v string) CreateIpAttributeOption {
	return func(o *CreateIpAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *Databases) WithCreateIpAttributeArray(v bool) CreateIpAttributeOption {
	return func(o *CreateIpAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
									
// CreateIpAttribute create IP address attribute.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createIpColumn` instead.
func (srv *Databases) CreateIpAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateIpAttributeOption)(*models.AttributeIp, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/ip")
	options := CreateIpAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
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

		parsed := models.AttributeIp{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeIp
	parsed, ok := resp.Result.(models.AttributeIp)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateIpAttributeOptions struct {
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateIpAttributeOptions) New() *UpdateIpAttributeOptions {
	options.enabledSetters = map[string]bool{
		"NewKey": false,
	}
	return &options
}
type UpdateIpAttributeOption func(*UpdateIpAttributeOptions)
func (srv *Databases) WithUpdateIpAttributeNewKey(v string) UpdateIpAttributeOption {
	return func(o *UpdateIpAttributeOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateIpAttribute update an ip attribute. Changing the `default` value will
// not update already existing documents.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updateIpColumn` instead.
func (srv *Databases) UpdateIpAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string, optionalSetters ...UpdateIpAttributeOption)(*models.AttributeIp, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/ip/{key}")
	options := UpdateIpAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = Default
	if options.enabledSetters["NewKey"] {
		params["newKey"] = options.NewKey
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

		parsed := models.AttributeIp{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeIp
	parsed, ok := resp.Result.(models.AttributeIp)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateLineAttributeOptions struct {
	Default []interface{}
	enabledSetters map[string]bool
}
func (options CreateLineAttributeOptions) New() *CreateLineAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
	}
	return &options
}
type CreateLineAttributeOption func(*CreateLineAttributeOptions)
func (srv *Databases) WithCreateLineAttributeDefault(v []interface{}) CreateLineAttributeOption {
	return func(o *CreateLineAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
									
// CreateLineAttribute create a geometric line attribute.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createLineColumn` instead.
func (srv *Databases) CreateLineAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateLineAttributeOption)(*models.AttributeLine, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/line")
	options := CreateLineAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
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

		parsed := models.AttributeLine{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeLine
	parsed, ok := resp.Result.(models.AttributeLine)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateLineAttributeOptions struct {
	Default []interface{}
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateLineAttributeOptions) New() *UpdateLineAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"NewKey": false,
	}
	return &options
}
type UpdateLineAttributeOption func(*UpdateLineAttributeOptions)
func (srv *Databases) WithUpdateLineAttributeDefault(v []interface{}) UpdateLineAttributeOption {
	return func(o *UpdateLineAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *Databases) WithUpdateLineAttributeNewKey(v string) UpdateLineAttributeOption {
	return func(o *UpdateLineAttributeOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
									
// UpdateLineAttribute update a line attribute. Changing the `default` value
// will not update already existing documents.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updateLineColumn` instead.
func (srv *Databases) UpdateLineAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...UpdateLineAttributeOption)(*models.AttributeLine, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/line/{key}")
	options := UpdateLineAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["NewKey"] {
		params["newKey"] = options.NewKey
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

		parsed := models.AttributeLine{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeLine
	parsed, ok := resp.Result.(models.AttributeLine)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreatePointAttributeOptions struct {
	Default []interface{}
	enabledSetters map[string]bool
}
func (options CreatePointAttributeOptions) New() *CreatePointAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
	}
	return &options
}
type CreatePointAttributeOption func(*CreatePointAttributeOptions)
func (srv *Databases) WithCreatePointAttributeDefault(v []interface{}) CreatePointAttributeOption {
	return func(o *CreatePointAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
									
// CreatePointAttribute create a geometric point attribute.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createPointColumn` instead.
func (srv *Databases) CreatePointAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreatePointAttributeOption)(*models.AttributePoint, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/point")
	options := CreatePointAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
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

		parsed := models.AttributePoint{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributePoint
	parsed, ok := resp.Result.(models.AttributePoint)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdatePointAttributeOptions struct {
	Default []interface{}
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdatePointAttributeOptions) New() *UpdatePointAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"NewKey": false,
	}
	return &options
}
type UpdatePointAttributeOption func(*UpdatePointAttributeOptions)
func (srv *Databases) WithUpdatePointAttributeDefault(v []interface{}) UpdatePointAttributeOption {
	return func(o *UpdatePointAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *Databases) WithUpdatePointAttributeNewKey(v string) UpdatePointAttributeOption {
	return func(o *UpdatePointAttributeOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
									
// UpdatePointAttribute update a point attribute. Changing the `default` value
// will not update already existing documents.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updatePointColumn` instead.
func (srv *Databases) UpdatePointAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...UpdatePointAttributeOption)(*models.AttributePoint, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/point/{key}")
	options := UpdatePointAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["NewKey"] {
		params["newKey"] = options.NewKey
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

		parsed := models.AttributePoint{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributePoint
	parsed, ok := resp.Result.(models.AttributePoint)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreatePolygonAttributeOptions struct {
	Default []interface{}
	enabledSetters map[string]bool
}
func (options CreatePolygonAttributeOptions) New() *CreatePolygonAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
	}
	return &options
}
type CreatePolygonAttributeOption func(*CreatePolygonAttributeOptions)
func (srv *Databases) WithCreatePolygonAttributeDefault(v []interface{}) CreatePolygonAttributeOption {
	return func(o *CreatePolygonAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
									
// CreatePolygonAttribute create a geometric polygon attribute.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createPolygonColumn` instead.
func (srv *Databases) CreatePolygonAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreatePolygonAttributeOption)(*models.AttributePolygon, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/polygon")
	options := CreatePolygonAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
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

		parsed := models.AttributePolygon{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributePolygon
	parsed, ok := resp.Result.(models.AttributePolygon)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdatePolygonAttributeOptions struct {
	Default []interface{}
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdatePolygonAttributeOptions) New() *UpdatePolygonAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"NewKey": false,
	}
	return &options
}
type UpdatePolygonAttributeOption func(*UpdatePolygonAttributeOptions)
func (srv *Databases) WithUpdatePolygonAttributeDefault(v []interface{}) UpdatePolygonAttributeOption {
	return func(o *UpdatePolygonAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *Databases) WithUpdatePolygonAttributeNewKey(v string) UpdatePolygonAttributeOption {
	return func(o *UpdatePolygonAttributeOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
									
// UpdatePolygonAttribute update a polygon attribute. Changing the `default`
// value will not update already existing documents.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updatePolygonColumn` instead.
func (srv *Databases) UpdatePolygonAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...UpdatePolygonAttributeOption)(*models.AttributePolygon, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/polygon/{key}")
	options := UpdatePolygonAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["NewKey"] {
		params["newKey"] = options.NewKey
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

		parsed := models.AttributePolygon{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributePolygon
	parsed, ok := resp.Result.(models.AttributePolygon)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateRelationshipAttributeOptions struct {
	TwoWay bool
	Key string
	TwoWayKey string
	OnDelete string
	enabledSetters map[string]bool
}
func (options CreateRelationshipAttributeOptions) New() *CreateRelationshipAttributeOptions {
	options.enabledSetters = map[string]bool{
		"TwoWay": false,
		"Key": false,
		"TwoWayKey": false,
		"OnDelete": false,
	}
	return &options
}
type CreateRelationshipAttributeOption func(*CreateRelationshipAttributeOptions)
func (srv *Databases) WithCreateRelationshipAttributeTwoWay(v bool) CreateRelationshipAttributeOption {
	return func(o *CreateRelationshipAttributeOptions) {
		o.TwoWay = v
		o.enabledSetters["TwoWay"] = true
	}
}
func (srv *Databases) WithCreateRelationshipAttributeKey(v string) CreateRelationshipAttributeOption {
	return func(o *CreateRelationshipAttributeOptions) {
		o.Key = v
		o.enabledSetters["Key"] = true
	}
}
func (srv *Databases) WithCreateRelationshipAttributeTwoWayKey(v string) CreateRelationshipAttributeOption {
	return func(o *CreateRelationshipAttributeOptions) {
		o.TwoWayKey = v
		o.enabledSetters["TwoWayKey"] = true
	}
}
func (srv *Databases) WithCreateRelationshipAttributeOnDelete(v string) CreateRelationshipAttributeOption {
	return func(o *CreateRelationshipAttributeOptions) {
		o.OnDelete = v
		o.enabledSetters["OnDelete"] = true
	}
}
									
// CreateRelationshipAttribute create relationship attribute. [Learn more
// about relationship
// attributes](https://appwrite.io/docs/databases-relationships#relationship-attributes).
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createRelationshipColumn` instead.
func (srv *Databases) CreateRelationshipAttribute(DatabaseId string, CollectionId string, RelatedCollectionId string, Type string, optionalSetters ...CreateRelationshipAttributeOption)(*models.AttributeRelationship, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/relationship")
	options := CreateRelationshipAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["relatedCollectionId"] = RelatedCollectionId
	params["type"] = Type
	if options.enabledSetters["TwoWay"] {
		params["twoWay"] = options.TwoWay
	}
	if options.enabledSetters["Key"] {
		params["key"] = options.Key
	}
	if options.enabledSetters["TwoWayKey"] {
		params["twoWayKey"] = options.TwoWayKey
	}
	if options.enabledSetters["OnDelete"] {
		params["onDelete"] = options.OnDelete
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

		parsed := models.AttributeRelationship{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeRelationship
	parsed, ok := resp.Result.(models.AttributeRelationship)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateStringAttributeOptions struct {
	Default string
	Array bool
	Encrypt bool
	enabledSetters map[string]bool
}
func (options CreateStringAttributeOptions) New() *CreateStringAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array": false,
		"Encrypt": false,
	}
	return &options
}
type CreateStringAttributeOption func(*CreateStringAttributeOptions)
func (srv *Databases) WithCreateStringAttributeDefault(v string) CreateStringAttributeOption {
	return func(o *CreateStringAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *Databases) WithCreateStringAttributeArray(v bool) CreateStringAttributeOption {
	return func(o *CreateStringAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
func (srv *Databases) WithCreateStringAttributeEncrypt(v bool) CreateStringAttributeOption {
	return func(o *CreateStringAttributeOptions) {
		o.Encrypt = v
		o.enabledSetters["Encrypt"] = true
	}
}
											
// CreateStringAttribute create a string attribute.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createStringColumn` instead.
func (srv *Databases) CreateStringAttribute(DatabaseId string, CollectionId string, Key string, Size int, Required bool, optionalSetters ...CreateStringAttributeOption)(*models.AttributeString, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/string")
	options := CreateStringAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["size"] = Size
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
	}
	if options.enabledSetters["Encrypt"] {
		params["encrypt"] = options.Encrypt
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

		parsed := models.AttributeString{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeString
	parsed, ok := resp.Result.(models.AttributeString)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateStringAttributeOptions struct {
	Size int
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateStringAttributeOptions) New() *UpdateStringAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Size": false,
		"NewKey": false,
	}
	return &options
}
type UpdateStringAttributeOption func(*UpdateStringAttributeOptions)
func (srv *Databases) WithUpdateStringAttributeSize(v int) UpdateStringAttributeOption {
	return func(o *UpdateStringAttributeOptions) {
		o.Size = v
		o.enabledSetters["Size"] = true
	}
}
func (srv *Databases) WithUpdateStringAttributeNewKey(v string) UpdateStringAttributeOption {
	return func(o *UpdateStringAttributeOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateStringAttribute update a string attribute. Changing the `default`
// value will not update already existing documents.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updateStringColumn` instead.
func (srv *Databases) UpdateStringAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string, optionalSetters ...UpdateStringAttributeOption)(*models.AttributeString, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/string/{key}")
	options := UpdateStringAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = Default
	if options.enabledSetters["Size"] {
		params["size"] = options.Size
	}
	if options.enabledSetters["NewKey"] {
		params["newKey"] = options.NewKey
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

		parsed := models.AttributeString{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeString
	parsed, ok := resp.Result.(models.AttributeString)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateUrlAttributeOptions struct {
	Default string
	Array bool
	enabledSetters map[string]bool
}
func (options CreateUrlAttributeOptions) New() *CreateUrlAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateUrlAttributeOption func(*CreateUrlAttributeOptions)
func (srv *Databases) WithCreateUrlAttributeDefault(v string) CreateUrlAttributeOption {
	return func(o *CreateUrlAttributeOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *Databases) WithCreateUrlAttributeArray(v bool) CreateUrlAttributeOption {
	return func(o *CreateUrlAttributeOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
									
// CreateUrlAttribute create a URL attribute.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createUrlColumn` instead.
func (srv *Databases) CreateUrlAttribute(DatabaseId string, CollectionId string, Key string, Required bool, optionalSetters ...CreateUrlAttributeOption)(*models.AttributeUrl, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/url")
	options := CreateUrlAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	if options.enabledSetters["Array"] {
		params["array"] = options.Array
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

		parsed := models.AttributeUrl{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeUrl
	parsed, ok := resp.Result.(models.AttributeUrl)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateUrlAttributeOptions struct {
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateUrlAttributeOptions) New() *UpdateUrlAttributeOptions {
	options.enabledSetters = map[string]bool{
		"NewKey": false,
	}
	return &options
}
type UpdateUrlAttributeOption func(*UpdateUrlAttributeOptions)
func (srv *Databases) WithUpdateUrlAttributeNewKey(v string) UpdateUrlAttributeOption {
	return func(o *UpdateUrlAttributeOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateUrlAttribute update an url attribute. Changing the `default` value
// will not update already existing documents.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updateUrlColumn` instead.
func (srv *Databases) UpdateUrlAttribute(DatabaseId string, CollectionId string, Key string, Required bool, Default string, optionalSetters ...UpdateUrlAttributeOption)(*models.AttributeUrl, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/url/{key}")
	options := UpdateUrlAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["required"] = Required
	params["default"] = Default
	if options.enabledSetters["NewKey"] {
		params["newKey"] = options.NewKey
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

		parsed := models.AttributeUrl{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeUrl
	parsed, ok := resp.Result.(models.AttributeUrl)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// GetAttribute get attribute by ID.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.getColumn` instead.
func (srv *Databases) GetAttribute(DatabaseId string, CollectionId string, Key string)(*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
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
					
// DeleteAttribute deletes an attribute.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.deleteColumn` instead.
func (srv *Databases) DeleteAttribute(DatabaseId string, CollectionId string, Key string)(*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
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
type UpdateRelationshipAttributeOptions struct {
	OnDelete string
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateRelationshipAttributeOptions) New() *UpdateRelationshipAttributeOptions {
	options.enabledSetters = map[string]bool{
		"OnDelete": false,
		"NewKey": false,
	}
	return &options
}
type UpdateRelationshipAttributeOption func(*UpdateRelationshipAttributeOptions)
func (srv *Databases) WithUpdateRelationshipAttributeOnDelete(v string) UpdateRelationshipAttributeOption {
	return func(o *UpdateRelationshipAttributeOptions) {
		o.OnDelete = v
		o.enabledSetters["OnDelete"] = true
	}
}
func (srv *Databases) WithUpdateRelationshipAttributeNewKey(v string) UpdateRelationshipAttributeOption {
	return func(o *UpdateRelationshipAttributeOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
							
// UpdateRelationshipAttribute update relationship attribute. [Learn more
// about relationship
// attributes](https://appwrite.io/docs/databases-relationships#relationship-attributes).
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updateRelationshipColumn` instead.
func (srv *Databases) UpdateRelationshipAttribute(DatabaseId string, CollectionId string, Key string, optionalSetters ...UpdateRelationshipAttributeOption)(*models.AttributeRelationship, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/attributes/{key}/relationship")
	options := UpdateRelationshipAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	if options.enabledSetters["OnDelete"] {
		params["onDelete"] = options.OnDelete
	}
	if options.enabledSetters["NewKey"] {
		params["newKey"] = options.NewKey
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

		parsed := models.AttributeRelationship{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AttributeRelationship
	parsed, ok := resp.Result.(models.AttributeRelationship)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListDocumentsOptions struct {
	Queries []string
	TransactionId string
	enabledSetters map[string]bool
}
func (options ListDocumentsOptions) New() *ListDocumentsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"TransactionId": false,
	}
	return &options
}
type ListDocumentsOption func(*ListDocumentsOptions)
func (srv *Databases) WithListDocumentsQueries(v []string) ListDocumentsOption {
	return func(o *ListDocumentsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Databases) WithListDocumentsTransactionId(v string) ListDocumentsOption {
	return func(o *ListDocumentsOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}
					
// ListDocuments get a list of all the user's documents in a given collection.
// You can use the query params to filter your results.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.listRows` instead.
func (srv *Databases) ListDocuments(DatabaseId string, CollectionId string, optionalSetters ...ListDocumentsOption)(*models.DocumentList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents")
	options := ListDocumentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

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
	Permissions []string
	TransactionId string
	enabledSetters map[string]bool
}
func (options CreateDocumentOptions) New() *CreateDocumentOptions {
	options.enabledSetters = map[string]bool{
		"Permissions": false,
		"TransactionId": false,
	}
	return &options
}
type CreateDocumentOption func(*CreateDocumentOptions)
func (srv *Databases) WithCreateDocumentPermissions(v []string) CreateDocumentOption {
	return func(o *CreateDocumentOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *Databases) WithCreateDocumentTransactionId(v string) CreateDocumentOption {
	return func(o *CreateDocumentOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}
									
// CreateDocument create a new Document. Before using this route, you should
// create a new collection resource using either a [server
// integration](https://appwrite.io/docs/server/databases#databasesCreateCollection)
// API or directly from your database console.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createRow` instead.
func (srv *Databases) CreateDocument(DatabaseId string, CollectionId string, DocumentId string, Data interface{}, optionalSetters ...CreateDocumentOption)(*models.Document, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents")
	options := CreateDocumentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documentId"] = DocumentId
	params["data"] = Data
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
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
type CreateDocumentsOptions struct {
	TransactionId string
	enabledSetters map[string]bool
}
func (options CreateDocumentsOptions) New() *CreateDocumentsOptions {
	options.enabledSetters = map[string]bool{
		"TransactionId": false,
	}
	return &options
}
type CreateDocumentsOption func(*CreateDocumentsOptions)
func (srv *Databases) WithCreateDocumentsTransactionId(v string) CreateDocumentsOption {
	return func(o *CreateDocumentsOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}
							
// CreateDocuments create new Documents. Before using this route, you should
// create a new collection resource using either a [server
// integration](https://appwrite.io/docs/server/databases#databasesCreateCollection)
// API or directly from your database console.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createRows` instead.
func (srv *Databases) CreateDocuments(DatabaseId string, CollectionId string, Documents []interface{}, optionalSetters ...CreateDocumentsOption)(*models.DocumentList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents")
	options := CreateDocumentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documents"] = Documents
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
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
	TransactionId string
	enabledSetters map[string]bool
}
func (options UpsertDocumentsOptions) New() *UpsertDocumentsOptions {
	options.enabledSetters = map[string]bool{
		"TransactionId": false,
	}
	return &options
}
type UpsertDocumentsOption func(*UpsertDocumentsOptions)
func (srv *Databases) WithUpsertDocumentsTransactionId(v string) UpsertDocumentsOption {
	return func(o *UpsertDocumentsOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}
							
// UpsertDocuments create or update Documents. Before using this route, you
// should create a new collection resource using either a [server
// integration](https://appwrite.io/docs/server/databases#databasesCreateCollection)
// API or directly from your database console.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.upsertRows` instead.
func (srv *Databases) UpsertDocuments(DatabaseId string, CollectionId string, Documents []interface{}, optionalSetters ...UpsertDocumentsOption)(*models.DocumentList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents")
	options := UpsertDocumentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documents"] = Documents
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
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
	Data interface{}
	Queries []string
	TransactionId string
	enabledSetters map[string]bool
}
func (options UpdateDocumentsOptions) New() *UpdateDocumentsOptions {
	options.enabledSetters = map[string]bool{
		"Data": false,
		"Queries": false,
		"TransactionId": false,
	}
	return &options
}
type UpdateDocumentsOption func(*UpdateDocumentsOptions)
func (srv *Databases) WithUpdateDocumentsData(v interface{}) UpdateDocumentsOption {
	return func(o *UpdateDocumentsOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func (srv *Databases) WithUpdateDocumentsQueries(v []string) UpdateDocumentsOption {
	return func(o *UpdateDocumentsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Databases) WithUpdateDocumentsTransactionId(v string) UpdateDocumentsOption {
	return func(o *UpdateDocumentsOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}
					
// UpdateDocuments update all documents that match your queries, if no queries
// are submitted then all documents are updated. You can pass only specific
// fields to be updated.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updateRows` instead.
func (srv *Databases) UpdateDocuments(DatabaseId string, CollectionId string, optionalSetters ...UpdateDocumentsOption)(*models.DocumentList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents")
	options := UpdateDocumentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
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
	Queries []string
	TransactionId string
	enabledSetters map[string]bool
}
func (options DeleteDocumentsOptions) New() *DeleteDocumentsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"TransactionId": false,
	}
	return &options
}
type DeleteDocumentsOption func(*DeleteDocumentsOptions)
func (srv *Databases) WithDeleteDocumentsQueries(v []string) DeleteDocumentsOption {
	return func(o *DeleteDocumentsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Databases) WithDeleteDocumentsTransactionId(v string) DeleteDocumentsOption {
	return func(o *DeleteDocumentsOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}
					
// DeleteDocuments bulk delete documents using queries, if no queries are
// passed then all documents are deleted.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.deleteRows` instead.
func (srv *Databases) DeleteDocuments(DatabaseId string, CollectionId string, optionalSetters ...DeleteDocumentsOption)(*models.DocumentList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents")
	options := DeleteDocumentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

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
	Queries []string
	TransactionId string
	enabledSetters map[string]bool
}
func (options GetDocumentOptions) New() *GetDocumentOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"TransactionId": false,
	}
	return &options
}
type GetDocumentOption func(*GetDocumentOptions)
func (srv *Databases) WithGetDocumentQueries(v []string) GetDocumentOption {
	return func(o *GetDocumentOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Databases) WithGetDocumentTransactionId(v string) GetDocumentOption {
	return func(o *GetDocumentOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}
							
// GetDocument get a document by its unique ID. This endpoint response returns
// a JSON object with the document data.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.getRow` instead.
func (srv *Databases) GetDocument(DatabaseId string, CollectionId string, DocumentId string, optionalSetters ...GetDocumentOption)(*models.Document, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")
	options := GetDocumentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documentId"] = DocumentId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

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
	Permissions []string
	TransactionId string
	enabledSetters map[string]bool
}
func (options UpsertDocumentOptions) New() *UpsertDocumentOptions {
	options.enabledSetters = map[string]bool{
		"Permissions": false,
		"TransactionId": false,
	}
	return &options
}
type UpsertDocumentOption func(*UpsertDocumentOptions)
func (srv *Databases) WithUpsertDocumentPermissions(v []string) UpsertDocumentOption {
	return func(o *UpsertDocumentOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *Databases) WithUpsertDocumentTransactionId(v string) UpsertDocumentOption {
	return func(o *UpsertDocumentOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}
									
// UpsertDocument create or update a Document. Before using this route, you
// should create a new collection resource using either a [server
// integration](https://appwrite.io/docs/server/databases#databasesCreateCollection)
// API or directly from your database console.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.upsertRow` instead.
func (srv *Databases) UpsertDocument(DatabaseId string, CollectionId string, DocumentId string, Data interface{}, optionalSetters ...UpsertDocumentOption)(*models.Document, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")
	options := UpsertDocumentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documentId"] = DocumentId
	params["data"] = Data
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
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
	Data interface{}
	Permissions []string
	TransactionId string
	enabledSetters map[string]bool
}
func (options UpdateDocumentOptions) New() *UpdateDocumentOptions {
	options.enabledSetters = map[string]bool{
		"Data": false,
		"Permissions": false,
		"TransactionId": false,
	}
	return &options
}
type UpdateDocumentOption func(*UpdateDocumentOptions)
func (srv *Databases) WithUpdateDocumentData(v interface{}) UpdateDocumentOption {
	return func(o *UpdateDocumentOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func (srv *Databases) WithUpdateDocumentPermissions(v []string) UpdateDocumentOption {
	return func(o *UpdateDocumentOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *Databases) WithUpdateDocumentTransactionId(v string) UpdateDocumentOption {
	return func(o *UpdateDocumentOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}
							
// UpdateDocument update a document by its unique ID. Using the patch method
// you can pass only specific fields that will get updated.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.updateRow` instead.
func (srv *Databases) UpdateDocument(DatabaseId string, CollectionId string, DocumentId string, optionalSetters ...UpdateDocumentOption)(*models.Document, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")
	options := UpdateDocumentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documentId"] = DocumentId
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
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
	TransactionId string
	enabledSetters map[string]bool
}
func (options DeleteDocumentOptions) New() *DeleteDocumentOptions {
	options.enabledSetters = map[string]bool{
		"TransactionId": false,
	}
	return &options
}
type DeleteDocumentOption func(*DeleteDocumentOptions)
func (srv *Databases) WithDeleteDocumentTransactionId(v string) DeleteDocumentOption {
	return func(o *DeleteDocumentOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}
							
// DeleteDocument delete a document by its unique ID.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.deleteRow` instead.
func (srv *Databases) DeleteDocument(DatabaseId string, CollectionId string, DocumentId string, optionalSetters ...DeleteDocumentOption)(*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")
	options := DeleteDocumentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documentId"] = DocumentId
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
	}
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
type DecrementDocumentAttributeOptions struct {
	Value float64
	Min float64
	TransactionId string
	enabledSetters map[string]bool
}
func (options DecrementDocumentAttributeOptions) New() *DecrementDocumentAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Value": false,
		"Min": false,
		"TransactionId": false,
	}
	return &options
}
type DecrementDocumentAttributeOption func(*DecrementDocumentAttributeOptions)
func (srv *Databases) WithDecrementDocumentAttributeValue(v float64) DecrementDocumentAttributeOption {
	return func(o *DecrementDocumentAttributeOptions) {
		o.Value = v
		o.enabledSetters["Value"] = true
	}
}
func (srv *Databases) WithDecrementDocumentAttributeMin(v float64) DecrementDocumentAttributeOption {
	return func(o *DecrementDocumentAttributeOptions) {
		o.Min = v
		o.enabledSetters["Min"] = true
	}
}
func (srv *Databases) WithDecrementDocumentAttributeTransactionId(v string) DecrementDocumentAttributeOption {
	return func(o *DecrementDocumentAttributeOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}
									
// DecrementDocumentAttribute decrement a specific attribute of a document by
// a given value.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.decrementRowColumn` instead.
func (srv *Databases) DecrementDocumentAttribute(DatabaseId string, CollectionId string, DocumentId string, Attribute string, optionalSetters ...DecrementDocumentAttributeOption)(*models.Document, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId, "{attribute}", Attribute)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}/{attribute}/decrement")
	options := DecrementDocumentAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documentId"] = DocumentId
	params["attribute"] = Attribute
	if options.enabledSetters["Value"] {
		params["value"] = options.Value
	}
	if options.enabledSetters["Min"] {
		params["min"] = options.Min
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
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
type IncrementDocumentAttributeOptions struct {
	Value float64
	Max float64
	TransactionId string
	enabledSetters map[string]bool
}
func (options IncrementDocumentAttributeOptions) New() *IncrementDocumentAttributeOptions {
	options.enabledSetters = map[string]bool{
		"Value": false,
		"Max": false,
		"TransactionId": false,
	}
	return &options
}
type IncrementDocumentAttributeOption func(*IncrementDocumentAttributeOptions)
func (srv *Databases) WithIncrementDocumentAttributeValue(v float64) IncrementDocumentAttributeOption {
	return func(o *IncrementDocumentAttributeOptions) {
		o.Value = v
		o.enabledSetters["Value"] = true
	}
}
func (srv *Databases) WithIncrementDocumentAttributeMax(v float64) IncrementDocumentAttributeOption {
	return func(o *IncrementDocumentAttributeOptions) {
		o.Max = v
		o.enabledSetters["Max"] = true
	}
}
func (srv *Databases) WithIncrementDocumentAttributeTransactionId(v string) IncrementDocumentAttributeOption {
	return func(o *IncrementDocumentAttributeOptions) {
		o.TransactionId = v
		o.enabledSetters["TransactionId"] = true
	}
}
									
// IncrementDocumentAttribute increment a specific attribute of a document by
// a given value.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.incrementRowColumn` instead.
func (srv *Databases) IncrementDocumentAttribute(DatabaseId string, CollectionId string, DocumentId string, Attribute string, optionalSetters ...IncrementDocumentAttributeOption)(*models.Document, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId, "{attribute}", Attribute)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}/{attribute}/increment")
	options := IncrementDocumentAttributeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["documentId"] = DocumentId
	params["attribute"] = Attribute
	if options.enabledSetters["Value"] {
		params["value"] = options.Value
	}
	if options.enabledSetters["Max"] {
		params["max"] = options.Max
	}
	if options.enabledSetters["TransactionId"] {
		params["transactionId"] = options.TransactionId
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
type ListIndexesOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListIndexesOptions) New() *ListIndexesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListIndexesOption func(*ListIndexesOptions)
func (srv *Databases) WithListIndexesQueries(v []string) ListIndexesOption {
	return func(o *ListIndexesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
					
// ListIndexes list indexes in the collection.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.listIndexes` instead.
func (srv *Databases) ListIndexes(DatabaseId string, CollectionId string, optionalSetters ...ListIndexesOption)(*models.IndexList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/indexes")
	options := ListIndexesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
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
	Orders []string
	Lengths []int
	enabledSetters map[string]bool
}
func (options CreateIndexOptions) New() *CreateIndexOptions {
	options.enabledSetters = map[string]bool{
		"Orders": false,
		"Lengths": false,
	}
	return &options
}
type CreateIndexOption func(*CreateIndexOptions)
func (srv *Databases) WithCreateIndexOrders(v []string) CreateIndexOption {
	return func(o *CreateIndexOptions) {
		o.Orders = v
		o.enabledSetters["Orders"] = true
	}
}
func (srv *Databases) WithCreateIndexLengths(v []int) CreateIndexOption {
	return func(o *CreateIndexOptions) {
		o.Lengths = v
		o.enabledSetters["Lengths"] = true
	}
}
											
// CreateIndex creates an index on the attributes listed. Your index should
// include all the attributes you will query in a single request.
// Attributes can be `key`, `fulltext`, and `unique`.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.createIndex` instead.
func (srv *Databases) CreateIndex(DatabaseId string, CollectionId string, Key string, Type string, Attributes []string, optionalSetters ...CreateIndexOption)(*models.Index, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/indexes")
	options := CreateIndexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	params["type"] = Type
	params["attributes"] = Attributes
	if options.enabledSetters["Orders"] {
		params["orders"] = options.Orders
	}
	if options.enabledSetters["Lengths"] {
		params["lengths"] = options.Lengths
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
					
// GetIndex get an index by its unique ID.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.getIndex` instead.
func (srv *Databases) GetIndex(DatabaseId string, CollectionId string, Key string)(*models.Index, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/indexes/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

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
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `TablesDB.deleteIndex` instead.
func (srv *Databases) DeleteIndex(DatabaseId string, CollectionId string, Key string)(*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{key}", Key)
	path := r.Replace("/databases/{databaseId}/collections/{collectionId}/indexes/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["collectionId"] = CollectionId
	params["key"] = Key
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
