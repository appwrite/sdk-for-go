package tablesdb

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// TablesDb service
type TablesDb struct {
	client client.Client
}

func New(clt client.Client) *TablesDb {
	return &TablesDb{
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
func (srv *TablesDb) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *TablesDb) WithListSearch(v string) ListOption {
	return func(o *ListOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
	
// List get a list of all databases from the current Appwrite project. You can
// use the search parameter to filter your results.
func (srv *TablesDb) List(optionalSetters ...ListOption)(*models.DatabaseList, error) {
	path := "/tablesdb"
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
	Type string
	enabledSetters map[string]bool
}
func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{
		"Enabled": false,
		"Type": false,
	}
	return &options
}
type CreateOption func(*CreateOptions)
func (srv *TablesDb) WithCreateEnabled(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *TablesDb) WithCreateType(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
					
// Create create a new Database.
func (srv *TablesDb) Create(DatabaseId string, Name string, optionalSetters ...CreateOption)(*models.Database, error) {
	path := "/tablesdb"
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
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
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
	
// Get get a database by its unique ID. This endpoint response returns a JSON
// object with the database metadata.
func (srv *TablesDb) Get(DatabaseId string)(*models.Database, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/tablesdb/{databaseId}")
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
func (srv *TablesDb) WithUpdateEnabled(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
					
// Update update a database by its unique ID.
func (srv *TablesDb) Update(DatabaseId string, Name string, optionalSetters ...UpdateOption)(*models.Database, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/tablesdb/{databaseId}")
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
func (srv *TablesDb) Delete(DatabaseId string)(*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/tablesdb/{databaseId}")
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
type ListTablesOptions struct {
	Queries []string
	Search string
	enabledSetters map[string]bool
}
func (options ListTablesOptions) New() *ListTablesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
	}
	return &options
}
type ListTablesOption func(*ListTablesOptions)
func (srv *TablesDb) WithListTablesQueries(v []string) ListTablesOption {
	return func(o *ListTablesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *TablesDb) WithListTablesSearch(v string) ListTablesOption {
	return func(o *ListTablesOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
			
// ListTables get a list of all tables that belong to the provided databaseId.
// You can use the search parameter to filter your results.
func (srv *TablesDb) ListTables(DatabaseId string, optionalSetters ...ListTablesOption)(*models.TableList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/tablesdb/{databaseId}/tables")
	options := ListTablesOptions{}.New()
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

		parsed := models.TableList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.TableList
	parsed, ok := resp.Result.(models.TableList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateTableOptions struct {
	Permissions []string
	RowSecurity bool
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateTableOptions) New() *CreateTableOptions {
	options.enabledSetters = map[string]bool{
		"Permissions": false,
		"RowSecurity": false,
		"Enabled": false,
	}
	return &options
}
type CreateTableOption func(*CreateTableOptions)
func (srv *TablesDb) WithCreateTablePermissions(v []string) CreateTableOption {
	return func(o *CreateTableOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *TablesDb) WithCreateTableRowSecurity(v bool) CreateTableOption {
	return func(o *CreateTableOptions) {
		o.RowSecurity = v
		o.enabledSetters["RowSecurity"] = true
	}
}
func (srv *TablesDb) WithCreateTableEnabled(v bool) CreateTableOption {
	return func(o *CreateTableOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
							
// CreateTable create a new Table. Before using this route, you should create
// a new database resource using either a [server
// integration](https://appwrite.io/docs/server/databases#databasesCreateTable)
// API or directly from your database console.
func (srv *TablesDb) CreateTable(DatabaseId string, TableId string, Name string, optionalSetters ...CreateTableOption)(*models.Table, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId)
	path := r.Replace("/tablesdb/{databaseId}/tables")
	options := CreateTableOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	params["name"] = Name
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["RowSecurity"] {
		params["rowSecurity"] = options.RowSecurity
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

		parsed := models.Table{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Table
	parsed, ok := resp.Result.(models.Table)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetTable get a table by its unique ID. This endpoint response returns a
// JSON object with the table metadata.
func (srv *TablesDb) GetTable(DatabaseId string, TableId string)(*models.Table, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Table{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Table
	parsed, ok := resp.Result.(models.Table)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateTableOptions struct {
	Permissions []string
	RowSecurity bool
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateTableOptions) New() *UpdateTableOptions {
	options.enabledSetters = map[string]bool{
		"Permissions": false,
		"RowSecurity": false,
		"Enabled": false,
	}
	return &options
}
type UpdateTableOption func(*UpdateTableOptions)
func (srv *TablesDb) WithUpdateTablePermissions(v []string) UpdateTableOption {
	return func(o *UpdateTableOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *TablesDb) WithUpdateTableRowSecurity(v bool) UpdateTableOption {
	return func(o *UpdateTableOptions) {
		o.RowSecurity = v
		o.enabledSetters["RowSecurity"] = true
	}
}
func (srv *TablesDb) WithUpdateTableEnabled(v bool) UpdateTableOption {
	return func(o *UpdateTableOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
							
// UpdateTable update a table by its unique ID.
func (srv *TablesDb) UpdateTable(DatabaseId string, TableId string, Name string, optionalSetters ...UpdateTableOption)(*models.Table, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}")
	options := UpdateTableOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	params["name"] = Name
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["RowSecurity"] {
		params["rowSecurity"] = options.RowSecurity
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

		parsed := models.Table{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Table
	parsed, ok := resp.Result.(models.Table)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// DeleteTable delete a table by its unique ID. Only users with write
// permissions have access to delete this resource.
func (srv *TablesDb) DeleteTable(DatabaseId string, TableId string)(*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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
type ListColumnsOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListColumnsOptions) New() *ListColumnsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListColumnsOption func(*ListColumnsOptions)
func (srv *TablesDb) WithListColumnsQueries(v []string) ListColumnsOption {
	return func(o *ListColumnsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
					
// ListColumns list columns in the table.
func (srv *TablesDb) ListColumns(DatabaseId string, TableId string, optionalSetters ...ListColumnsOption)(*models.ColumnList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns")
	options := ListColumnsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnList
	parsed, ok := resp.Result.(models.ColumnList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateBooleanColumnOptions struct {
	Default bool
	Array bool
	enabledSetters map[string]bool
}
func (options CreateBooleanColumnOptions) New() *CreateBooleanColumnOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateBooleanColumnOption func(*CreateBooleanColumnOptions)
func (srv *TablesDb) WithCreateBooleanColumnDefault(v bool) CreateBooleanColumnOption {
	return func(o *CreateBooleanColumnOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *TablesDb) WithCreateBooleanColumnArray(v bool) CreateBooleanColumnOption {
	return func(o *CreateBooleanColumnOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
									
// CreateBooleanColumn create a boolean column.
func (srv *TablesDb) CreateBooleanColumn(DatabaseId string, TableId string, Key string, Required bool, optionalSetters ...CreateBooleanColumnOption)(*models.ColumnBoolean, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/boolean")
	options := CreateBooleanColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnBoolean{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnBoolean
	parsed, ok := resp.Result.(models.ColumnBoolean)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateBooleanColumnOptions struct {
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateBooleanColumnOptions) New() *UpdateBooleanColumnOptions {
	options.enabledSetters = map[string]bool{
		"NewKey": false,
	}
	return &options
}
type UpdateBooleanColumnOption func(*UpdateBooleanColumnOptions)
func (srv *TablesDb) WithUpdateBooleanColumnNewKey(v string) UpdateBooleanColumnOption {
	return func(o *UpdateBooleanColumnOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateBooleanColumn update a boolean column. Changing the `default` value
// will not update already existing rows.
func (srv *TablesDb) UpdateBooleanColumn(DatabaseId string, TableId string, Key string, Required bool, Default bool, optionalSetters ...UpdateBooleanColumnOption)(*models.ColumnBoolean, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{key}", Key)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/boolean/{key}")
	options := UpdateBooleanColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnBoolean{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnBoolean
	parsed, ok := resp.Result.(models.ColumnBoolean)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateDatetimeColumnOptions struct {
	Default string
	Array bool
	enabledSetters map[string]bool
}
func (options CreateDatetimeColumnOptions) New() *CreateDatetimeColumnOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateDatetimeColumnOption func(*CreateDatetimeColumnOptions)
func (srv *TablesDb) WithCreateDatetimeColumnDefault(v string) CreateDatetimeColumnOption {
	return func(o *CreateDatetimeColumnOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *TablesDb) WithCreateDatetimeColumnArray(v bool) CreateDatetimeColumnOption {
	return func(o *CreateDatetimeColumnOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
									
// CreateDatetimeColumn create a date time column according to the ISO 8601
// standard.
func (srv *TablesDb) CreateDatetimeColumn(DatabaseId string, TableId string, Key string, Required bool, optionalSetters ...CreateDatetimeColumnOption)(*models.ColumnDatetime, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/datetime")
	options := CreateDatetimeColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnDatetime{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnDatetime
	parsed, ok := resp.Result.(models.ColumnDatetime)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateDatetimeColumnOptions struct {
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateDatetimeColumnOptions) New() *UpdateDatetimeColumnOptions {
	options.enabledSetters = map[string]bool{
		"NewKey": false,
	}
	return &options
}
type UpdateDatetimeColumnOption func(*UpdateDatetimeColumnOptions)
func (srv *TablesDb) WithUpdateDatetimeColumnNewKey(v string) UpdateDatetimeColumnOption {
	return func(o *UpdateDatetimeColumnOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateDatetimeColumn update a date time column. Changing the `default`
// value will not update already existing rows.
func (srv *TablesDb) UpdateDatetimeColumn(DatabaseId string, TableId string, Key string, Required bool, Default string, optionalSetters ...UpdateDatetimeColumnOption)(*models.ColumnDatetime, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{key}", Key)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/datetime/{key}")
	options := UpdateDatetimeColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnDatetime{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnDatetime
	parsed, ok := resp.Result.(models.ColumnDatetime)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateEmailColumnOptions struct {
	Default string
	Array bool
	enabledSetters map[string]bool
}
func (options CreateEmailColumnOptions) New() *CreateEmailColumnOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateEmailColumnOption func(*CreateEmailColumnOptions)
func (srv *TablesDb) WithCreateEmailColumnDefault(v string) CreateEmailColumnOption {
	return func(o *CreateEmailColumnOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *TablesDb) WithCreateEmailColumnArray(v bool) CreateEmailColumnOption {
	return func(o *CreateEmailColumnOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
									
// CreateEmailColumn create an email column.
func (srv *TablesDb) CreateEmailColumn(DatabaseId string, TableId string, Key string, Required bool, optionalSetters ...CreateEmailColumnOption)(*models.ColumnEmail, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/email")
	options := CreateEmailColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnEmail{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnEmail
	parsed, ok := resp.Result.(models.ColumnEmail)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateEmailColumnOptions struct {
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateEmailColumnOptions) New() *UpdateEmailColumnOptions {
	options.enabledSetters = map[string]bool{
		"NewKey": false,
	}
	return &options
}
type UpdateEmailColumnOption func(*UpdateEmailColumnOptions)
func (srv *TablesDb) WithUpdateEmailColumnNewKey(v string) UpdateEmailColumnOption {
	return func(o *UpdateEmailColumnOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateEmailColumn update an email column. Changing the `default` value will
// not update already existing rows.
func (srv *TablesDb) UpdateEmailColumn(DatabaseId string, TableId string, Key string, Required bool, Default string, optionalSetters ...UpdateEmailColumnOption)(*models.ColumnEmail, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{key}", Key)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/email/{key}")
	options := UpdateEmailColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnEmail{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnEmail
	parsed, ok := resp.Result.(models.ColumnEmail)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateEnumColumnOptions struct {
	Default string
	Array bool
	enabledSetters map[string]bool
}
func (options CreateEnumColumnOptions) New() *CreateEnumColumnOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateEnumColumnOption func(*CreateEnumColumnOptions)
func (srv *TablesDb) WithCreateEnumColumnDefault(v string) CreateEnumColumnOption {
	return func(o *CreateEnumColumnOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *TablesDb) WithCreateEnumColumnArray(v bool) CreateEnumColumnOption {
	return func(o *CreateEnumColumnOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
											
// CreateEnumColumn create an enumeration column. The `elements` param acts as
// a white-list of accepted values for this column.
func (srv *TablesDb) CreateEnumColumn(DatabaseId string, TableId string, Key string, Elements []string, Required bool, optionalSetters ...CreateEnumColumnOption)(*models.ColumnEnum, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/enum")
	options := CreateEnumColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnEnum{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnEnum
	parsed, ok := resp.Result.(models.ColumnEnum)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateEnumColumnOptions struct {
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateEnumColumnOptions) New() *UpdateEnumColumnOptions {
	options.enabledSetters = map[string]bool{
		"NewKey": false,
	}
	return &options
}
type UpdateEnumColumnOption func(*UpdateEnumColumnOptions)
func (srv *TablesDb) WithUpdateEnumColumnNewKey(v string) UpdateEnumColumnOption {
	return func(o *UpdateEnumColumnOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
													
// UpdateEnumColumn update an enum column. Changing the `default` value will
// not update already existing rows.
func (srv *TablesDb) UpdateEnumColumn(DatabaseId string, TableId string, Key string, Elements []string, Required bool, Default string, optionalSetters ...UpdateEnumColumnOption)(*models.ColumnEnum, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{key}", Key)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/enum/{key}")
	options := UpdateEnumColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnEnum{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnEnum
	parsed, ok := resp.Result.(models.ColumnEnum)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateFloatColumnOptions struct {
	Min float64
	Max float64
	Default float64
	Array bool
	enabledSetters map[string]bool
}
func (options CreateFloatColumnOptions) New() *CreateFloatColumnOptions {
	options.enabledSetters = map[string]bool{
		"Min": false,
		"Max": false,
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateFloatColumnOption func(*CreateFloatColumnOptions)
func (srv *TablesDb) WithCreateFloatColumnMin(v float64) CreateFloatColumnOption {
	return func(o *CreateFloatColumnOptions) {
		o.Min = v
		o.enabledSetters["Min"] = true
	}
}
func (srv *TablesDb) WithCreateFloatColumnMax(v float64) CreateFloatColumnOption {
	return func(o *CreateFloatColumnOptions) {
		o.Max = v
		o.enabledSetters["Max"] = true
	}
}
func (srv *TablesDb) WithCreateFloatColumnDefault(v float64) CreateFloatColumnOption {
	return func(o *CreateFloatColumnOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *TablesDb) WithCreateFloatColumnArray(v bool) CreateFloatColumnOption {
	return func(o *CreateFloatColumnOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
									
// CreateFloatColumn create a float column. Optionally, minimum and maximum
// values can be provided.
func (srv *TablesDb) CreateFloatColumn(DatabaseId string, TableId string, Key string, Required bool, optionalSetters ...CreateFloatColumnOption)(*models.ColumnFloat, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/float")
	options := CreateFloatColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnFloat{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnFloat
	parsed, ok := resp.Result.(models.ColumnFloat)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateFloatColumnOptions struct {
	Min float64
	Max float64
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateFloatColumnOptions) New() *UpdateFloatColumnOptions {
	options.enabledSetters = map[string]bool{
		"Min": false,
		"Max": false,
		"NewKey": false,
	}
	return &options
}
type UpdateFloatColumnOption func(*UpdateFloatColumnOptions)
func (srv *TablesDb) WithUpdateFloatColumnMin(v float64) UpdateFloatColumnOption {
	return func(o *UpdateFloatColumnOptions) {
		o.Min = v
		o.enabledSetters["Min"] = true
	}
}
func (srv *TablesDb) WithUpdateFloatColumnMax(v float64) UpdateFloatColumnOption {
	return func(o *UpdateFloatColumnOptions) {
		o.Max = v
		o.enabledSetters["Max"] = true
	}
}
func (srv *TablesDb) WithUpdateFloatColumnNewKey(v string) UpdateFloatColumnOption {
	return func(o *UpdateFloatColumnOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateFloatColumn update a float column. Changing the `default` value will
// not update already existing rows.
func (srv *TablesDb) UpdateFloatColumn(DatabaseId string, TableId string, Key string, Required bool, Default float64, optionalSetters ...UpdateFloatColumnOption)(*models.ColumnFloat, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{key}", Key)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/float/{key}")
	options := UpdateFloatColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnFloat{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnFloat
	parsed, ok := resp.Result.(models.ColumnFloat)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateIntegerColumnOptions struct {
	Min int
	Max int
	Default int
	Array bool
	enabledSetters map[string]bool
}
func (options CreateIntegerColumnOptions) New() *CreateIntegerColumnOptions {
	options.enabledSetters = map[string]bool{
		"Min": false,
		"Max": false,
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateIntegerColumnOption func(*CreateIntegerColumnOptions)
func (srv *TablesDb) WithCreateIntegerColumnMin(v int) CreateIntegerColumnOption {
	return func(o *CreateIntegerColumnOptions) {
		o.Min = v
		o.enabledSetters["Min"] = true
	}
}
func (srv *TablesDb) WithCreateIntegerColumnMax(v int) CreateIntegerColumnOption {
	return func(o *CreateIntegerColumnOptions) {
		o.Max = v
		o.enabledSetters["Max"] = true
	}
}
func (srv *TablesDb) WithCreateIntegerColumnDefault(v int) CreateIntegerColumnOption {
	return func(o *CreateIntegerColumnOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *TablesDb) WithCreateIntegerColumnArray(v bool) CreateIntegerColumnOption {
	return func(o *CreateIntegerColumnOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
									
// CreateIntegerColumn create an integer column. Optionally, minimum and
// maximum values can be provided.
func (srv *TablesDb) CreateIntegerColumn(DatabaseId string, TableId string, Key string, Required bool, optionalSetters ...CreateIntegerColumnOption)(*models.ColumnInteger, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/integer")
	options := CreateIntegerColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnInteger{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnInteger
	parsed, ok := resp.Result.(models.ColumnInteger)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateIntegerColumnOptions struct {
	Min int
	Max int
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateIntegerColumnOptions) New() *UpdateIntegerColumnOptions {
	options.enabledSetters = map[string]bool{
		"Min": false,
		"Max": false,
		"NewKey": false,
	}
	return &options
}
type UpdateIntegerColumnOption func(*UpdateIntegerColumnOptions)
func (srv *TablesDb) WithUpdateIntegerColumnMin(v int) UpdateIntegerColumnOption {
	return func(o *UpdateIntegerColumnOptions) {
		o.Min = v
		o.enabledSetters["Min"] = true
	}
}
func (srv *TablesDb) WithUpdateIntegerColumnMax(v int) UpdateIntegerColumnOption {
	return func(o *UpdateIntegerColumnOptions) {
		o.Max = v
		o.enabledSetters["Max"] = true
	}
}
func (srv *TablesDb) WithUpdateIntegerColumnNewKey(v string) UpdateIntegerColumnOption {
	return func(o *UpdateIntegerColumnOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateIntegerColumn update an integer column. Changing the `default` value
// will not update already existing rows.
func (srv *TablesDb) UpdateIntegerColumn(DatabaseId string, TableId string, Key string, Required bool, Default int, optionalSetters ...UpdateIntegerColumnOption)(*models.ColumnInteger, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{key}", Key)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/integer/{key}")
	options := UpdateIntegerColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnInteger{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnInteger
	parsed, ok := resp.Result.(models.ColumnInteger)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateIpColumnOptions struct {
	Default string
	Array bool
	enabledSetters map[string]bool
}
func (options CreateIpColumnOptions) New() *CreateIpColumnOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateIpColumnOption func(*CreateIpColumnOptions)
func (srv *TablesDb) WithCreateIpColumnDefault(v string) CreateIpColumnOption {
	return func(o *CreateIpColumnOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *TablesDb) WithCreateIpColumnArray(v bool) CreateIpColumnOption {
	return func(o *CreateIpColumnOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
									
// CreateIpColumn create IP address column.
func (srv *TablesDb) CreateIpColumn(DatabaseId string, TableId string, Key string, Required bool, optionalSetters ...CreateIpColumnOption)(*models.ColumnIp, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/ip")
	options := CreateIpColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnIp{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnIp
	parsed, ok := resp.Result.(models.ColumnIp)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateIpColumnOptions struct {
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateIpColumnOptions) New() *UpdateIpColumnOptions {
	options.enabledSetters = map[string]bool{
		"NewKey": false,
	}
	return &options
}
type UpdateIpColumnOption func(*UpdateIpColumnOptions)
func (srv *TablesDb) WithUpdateIpColumnNewKey(v string) UpdateIpColumnOption {
	return func(o *UpdateIpColumnOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateIpColumn update an ip column. Changing the `default` value will not
// update already existing rows.
func (srv *TablesDb) UpdateIpColumn(DatabaseId string, TableId string, Key string, Required bool, Default string, optionalSetters ...UpdateIpColumnOption)(*models.ColumnIp, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{key}", Key)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/ip/{key}")
	options := UpdateIpColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnIp{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnIp
	parsed, ok := resp.Result.(models.ColumnIp)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateRelationshipColumnOptions struct {
	TwoWay bool
	Key string
	TwoWayKey string
	OnDelete string
	enabledSetters map[string]bool
}
func (options CreateRelationshipColumnOptions) New() *CreateRelationshipColumnOptions {
	options.enabledSetters = map[string]bool{
		"TwoWay": false,
		"Key": false,
		"TwoWayKey": false,
		"OnDelete": false,
	}
	return &options
}
type CreateRelationshipColumnOption func(*CreateRelationshipColumnOptions)
func (srv *TablesDb) WithCreateRelationshipColumnTwoWay(v bool) CreateRelationshipColumnOption {
	return func(o *CreateRelationshipColumnOptions) {
		o.TwoWay = v
		o.enabledSetters["TwoWay"] = true
	}
}
func (srv *TablesDb) WithCreateRelationshipColumnKey(v string) CreateRelationshipColumnOption {
	return func(o *CreateRelationshipColumnOptions) {
		o.Key = v
		o.enabledSetters["Key"] = true
	}
}
func (srv *TablesDb) WithCreateRelationshipColumnTwoWayKey(v string) CreateRelationshipColumnOption {
	return func(o *CreateRelationshipColumnOptions) {
		o.TwoWayKey = v
		o.enabledSetters["TwoWayKey"] = true
	}
}
func (srv *TablesDb) WithCreateRelationshipColumnOnDelete(v string) CreateRelationshipColumnOption {
	return func(o *CreateRelationshipColumnOptions) {
		o.OnDelete = v
		o.enabledSetters["OnDelete"] = true
	}
}
									
// CreateRelationshipColumn create relationship column. [Learn more about
// relationship
// columns](https://appwrite.io/docs/databases-relationships#relationship-columns).
func (srv *TablesDb) CreateRelationshipColumn(DatabaseId string, TableId string, RelatedTableId string, Type string, optionalSetters ...CreateRelationshipColumnOption)(*models.ColumnRelationship, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/relationship")
	options := CreateRelationshipColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	params["relatedTableId"] = RelatedTableId
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

		parsed := models.ColumnRelationship{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnRelationship
	parsed, ok := resp.Result.(models.ColumnRelationship)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateStringColumnOptions struct {
	Default string
	Array bool
	Encrypt bool
	enabledSetters map[string]bool
}
func (options CreateStringColumnOptions) New() *CreateStringColumnOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array": false,
		"Encrypt": false,
	}
	return &options
}
type CreateStringColumnOption func(*CreateStringColumnOptions)
func (srv *TablesDb) WithCreateStringColumnDefault(v string) CreateStringColumnOption {
	return func(o *CreateStringColumnOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *TablesDb) WithCreateStringColumnArray(v bool) CreateStringColumnOption {
	return func(o *CreateStringColumnOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
func (srv *TablesDb) WithCreateStringColumnEncrypt(v bool) CreateStringColumnOption {
	return func(o *CreateStringColumnOptions) {
		o.Encrypt = v
		o.enabledSetters["Encrypt"] = true
	}
}
											
// CreateStringColumn create a string column.
func (srv *TablesDb) CreateStringColumn(DatabaseId string, TableId string, Key string, Size int, Required bool, optionalSetters ...CreateStringColumnOption)(*models.ColumnString, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/string")
	options := CreateStringColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnString{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnString
	parsed, ok := resp.Result.(models.ColumnString)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateStringColumnOptions struct {
	Size int
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateStringColumnOptions) New() *UpdateStringColumnOptions {
	options.enabledSetters = map[string]bool{
		"Size": false,
		"NewKey": false,
	}
	return &options
}
type UpdateStringColumnOption func(*UpdateStringColumnOptions)
func (srv *TablesDb) WithUpdateStringColumnSize(v int) UpdateStringColumnOption {
	return func(o *UpdateStringColumnOptions) {
		o.Size = v
		o.enabledSetters["Size"] = true
	}
}
func (srv *TablesDb) WithUpdateStringColumnNewKey(v string) UpdateStringColumnOption {
	return func(o *UpdateStringColumnOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateStringColumn update a string column. Changing the `default` value
// will not update already existing rows.
func (srv *TablesDb) UpdateStringColumn(DatabaseId string, TableId string, Key string, Required bool, Default string, optionalSetters ...UpdateStringColumnOption)(*models.ColumnString, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{key}", Key)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/string/{key}")
	options := UpdateStringColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnString{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnString
	parsed, ok := resp.Result.(models.ColumnString)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateUrlColumnOptions struct {
	Default string
	Array bool
	enabledSetters map[string]bool
}
func (options CreateUrlColumnOptions) New() *CreateUrlColumnOptions {
	options.enabledSetters = map[string]bool{
		"Default": false,
		"Array": false,
	}
	return &options
}
type CreateUrlColumnOption func(*CreateUrlColumnOptions)
func (srv *TablesDb) WithCreateUrlColumnDefault(v string) CreateUrlColumnOption {
	return func(o *CreateUrlColumnOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}
func (srv *TablesDb) WithCreateUrlColumnArray(v bool) CreateUrlColumnOption {
	return func(o *CreateUrlColumnOptions) {
		o.Array = v
		o.enabledSetters["Array"] = true
	}
}
									
// CreateUrlColumn create a URL column.
func (srv *TablesDb) CreateUrlColumn(DatabaseId string, TableId string, Key string, Required bool, optionalSetters ...CreateUrlColumnOption)(*models.ColumnUrl, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/url")
	options := CreateUrlColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnUrl{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnUrl
	parsed, ok := resp.Result.(models.ColumnUrl)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateUrlColumnOptions struct {
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateUrlColumnOptions) New() *UpdateUrlColumnOptions {
	options.enabledSetters = map[string]bool{
		"NewKey": false,
	}
	return &options
}
type UpdateUrlColumnOption func(*UpdateUrlColumnOptions)
func (srv *TablesDb) WithUpdateUrlColumnNewKey(v string) UpdateUrlColumnOption {
	return func(o *UpdateUrlColumnOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
											
// UpdateUrlColumn update an url column. Changing the `default` value will not
// update already existing rows.
func (srv *TablesDb) UpdateUrlColumn(DatabaseId string, TableId string, Key string, Required bool, Default string, optionalSetters ...UpdateUrlColumnOption)(*models.ColumnUrl, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{key}", Key)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/url/{key}")
	options := UpdateUrlColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnUrl{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnUrl
	parsed, ok := resp.Result.(models.ColumnUrl)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// GetColumn get column by ID.
func (srv *TablesDb) GetColumn(DatabaseId string, TableId string, Key string)(*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{key}", Key)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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
					
// DeleteColumn deletes a column.
func (srv *TablesDb) DeleteColumn(DatabaseId string, TableId string, Key string)(*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{key}", Key)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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
type UpdateRelationshipColumnOptions struct {
	OnDelete string
	NewKey string
	enabledSetters map[string]bool
}
func (options UpdateRelationshipColumnOptions) New() *UpdateRelationshipColumnOptions {
	options.enabledSetters = map[string]bool{
		"OnDelete": false,
		"NewKey": false,
	}
	return &options
}
type UpdateRelationshipColumnOption func(*UpdateRelationshipColumnOptions)
func (srv *TablesDb) WithUpdateRelationshipColumnOnDelete(v string) UpdateRelationshipColumnOption {
	return func(o *UpdateRelationshipColumnOptions) {
		o.OnDelete = v
		o.enabledSetters["OnDelete"] = true
	}
}
func (srv *TablesDb) WithUpdateRelationshipColumnNewKey(v string) UpdateRelationshipColumnOption {
	return func(o *UpdateRelationshipColumnOptions) {
		o.NewKey = v
		o.enabledSetters["NewKey"] = true
	}
}
							
// UpdateRelationshipColumn update relationship column. [Learn more about
// relationship
// columns](https://appwrite.io/docs/databases-relationships#relationship-columns).
func (srv *TablesDb) UpdateRelationshipColumn(DatabaseId string, TableId string, Key string, optionalSetters ...UpdateRelationshipColumnOption)(*models.ColumnRelationship, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{key}", Key)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/columns/{key}/relationship")
	options := UpdateRelationshipColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnRelationship{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnRelationship
	parsed, ok := resp.Result.(models.ColumnRelationship)
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
func (srv *TablesDb) WithListIndexesQueries(v []string) ListIndexesOption {
	return func(o *ListIndexesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
					
// ListIndexes list indexes in the collection.
func (srv *TablesDb) ListIndexes(DatabaseId string, TableId string, optionalSetters ...ListIndexesOption)(*models.ColumnIndexList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/indexes")
	options := ListIndexesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.ColumnIndexList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnIndexList
	parsed, ok := resp.Result.(models.ColumnIndexList)
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
func (srv *TablesDb) WithCreateIndexOrders(v []string) CreateIndexOption {
	return func(o *CreateIndexOptions) {
		o.Orders = v
		o.enabledSetters["Orders"] = true
	}
}
func (srv *TablesDb) WithCreateIndexLengths(v []int) CreateIndexOption {
	return func(o *CreateIndexOptions) {
		o.Lengths = v
		o.enabledSetters["Lengths"] = true
	}
}
											
// CreateIndex creates an index on the columns listed. Your index should
// include all the columns you will query in a single request.
// Attributes can be `key`, `fulltext`, and `unique`.
func (srv *TablesDb) CreateIndex(DatabaseId string, TableId string, Key string, Type string, Columns []string, optionalSetters ...CreateIndexOption)(*models.ColumnIndex, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/indexes")
	options := CreateIndexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	params["key"] = Key
	params["type"] = Type
	params["columns"] = Columns
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

		parsed := models.ColumnIndex{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnIndex
	parsed, ok := resp.Result.(models.ColumnIndex)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// GetIndex get index by ID.
func (srv *TablesDb) GetIndex(DatabaseId string, TableId string, Key string)(*models.ColumnIndex, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{key}", Key)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/indexes/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	params["key"] = Key
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ColumnIndex{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ColumnIndex
	parsed, ok := resp.Result.(models.ColumnIndex)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// DeleteIndex delete an index.
func (srv *TablesDb) DeleteIndex(DatabaseId string, TableId string, Key string)(*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{key}", Key)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/indexes/{key}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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
type ListRowsOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListRowsOptions) New() *ListRowsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListRowsOption func(*ListRowsOptions)
func (srv *TablesDb) WithListRowsQueries(v []string) ListRowsOption {
	return func(o *ListRowsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
					
// ListRows get a list of all the user's rows in a given table. You can use
// the query params to filter your results.
func (srv *TablesDb) ListRows(DatabaseId string, TableId string, optionalSetters ...ListRowsOption)(*models.RowList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/rows")
	options := ListRowsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
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

		parsed := models.RowList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.RowList
	parsed, ok := resp.Result.(models.RowList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateRowOptions struct {
	Permissions []string
	enabledSetters map[string]bool
}
func (options CreateRowOptions) New() *CreateRowOptions {
	options.enabledSetters = map[string]bool{
		"Permissions": false,
	}
	return &options
}
type CreateRowOption func(*CreateRowOptions)
func (srv *TablesDb) WithCreateRowPermissions(v []string) CreateRowOption {
	return func(o *CreateRowOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
									
// CreateRow create a new Row. Before using this route, you should create a
// new table resource using either a [server
// integration](https://appwrite.io/docs/server/databases#databasesCreateTable)
// API or directly from your database console.
func (srv *TablesDb) CreateRow(DatabaseId string, TableId string, RowId string, Data interface{}, optionalSetters ...CreateRowOption)(*models.Row, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/rows")
	options := CreateRowOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	params["rowId"] = RowId
	params["data"] = Data
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
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

		parsed := models.Row{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Row
	parsed, ok := resp.Result.(models.Row)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// CreateRows create new Rows. Before using this route, you should create a
// new table resource using either a [server
// integration](https://appwrite.io/docs/server/databases#databasesCreateTable)
// API or directly from your database console.
func (srv *TablesDb) CreateRows(DatabaseId string, TableId string, Rows []interface{})(*models.RowList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/rows")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	params["rows"] = Rows
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.RowList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.RowList
	parsed, ok := resp.Result.(models.RowList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// UpsertRows create or update Rows. Before using this route, you should
// create a new table resource using either a [server
// integration](https://appwrite.io/docs/server/databases#databasesCreateTable)
// API or directly from your database console.
func (srv *TablesDb) UpsertRows(DatabaseId string, TableId string, Rows []interface{})(*models.RowList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/rows")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	params["rows"] = Rows
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.RowList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.RowList
	parsed, ok := resp.Result.(models.RowList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateRowsOptions struct {
	Data interface{}
	Queries []string
	enabledSetters map[string]bool
}
func (options UpdateRowsOptions) New() *UpdateRowsOptions {
	options.enabledSetters = map[string]bool{
		"Data": false,
		"Queries": false,
	}
	return &options
}
type UpdateRowsOption func(*UpdateRowsOptions)
func (srv *TablesDb) WithUpdateRowsData(v interface{}) UpdateRowsOption {
	return func(o *UpdateRowsOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func (srv *TablesDb) WithUpdateRowsQueries(v []string) UpdateRowsOption {
	return func(o *UpdateRowsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
					
// UpdateRows update all rows that match your queries, if no queries are
// submitted then all rows are updated. You can pass only specific fields to
// be updated.
func (srv *TablesDb) UpdateRows(DatabaseId string, TableId string, optionalSetters ...UpdateRowsOption)(*models.RowList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/rows")
	options := UpdateRowsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
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

		parsed := models.RowList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.RowList
	parsed, ok := resp.Result.(models.RowList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type DeleteRowsOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options DeleteRowsOptions) New() *DeleteRowsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type DeleteRowsOption func(*DeleteRowsOptions)
func (srv *TablesDb) WithDeleteRowsQueries(v []string) DeleteRowsOption {
	return func(o *DeleteRowsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
					
// DeleteRows bulk delete rows using queries, if no queries are passed then
// all rows are deleted.
func (srv *TablesDb) DeleteRows(DatabaseId string, TableId string, optionalSetters ...DeleteRowsOption)(*models.RowList, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/rows")
	options := DeleteRowsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
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

		parsed := models.RowList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.RowList
	parsed, ok := resp.Result.(models.RowList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetRowOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options GetRowOptions) New() *GetRowOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type GetRowOption func(*GetRowOptions)
func (srv *TablesDb) WithGetRowQueries(v []string) GetRowOption {
	return func(o *GetRowOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
							
// GetRow get a row by its unique ID. This endpoint response returns a JSON
// object with the row data.
func (srv *TablesDb) GetRow(DatabaseId string, TableId string, RowId string, optionalSetters ...GetRowOption)(*models.Row, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{rowId}", RowId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/rows/{rowId}")
	options := GetRowOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	params["rowId"] = RowId
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

		parsed := models.Row{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Row
	parsed, ok := resp.Result.(models.Row)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpsertRowOptions struct {
	Data interface{}
	Permissions []string
	enabledSetters map[string]bool
}
func (options UpsertRowOptions) New() *UpsertRowOptions {
	options.enabledSetters = map[string]bool{
		"Data": false,
		"Permissions": false,
	}
	return &options
}
type UpsertRowOption func(*UpsertRowOptions)
func (srv *TablesDb) WithUpsertRowData(v interface{}) UpsertRowOption {
	return func(o *UpsertRowOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func (srv *TablesDb) WithUpsertRowPermissions(v []string) UpsertRowOption {
	return func(o *UpsertRowOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
							
// UpsertRow create or update a Row. Before using this route, you should
// create a new table resource using either a [server
// integration](https://appwrite.io/docs/server/databases#databasesCreateTable)
// API or directly from your database console.
func (srv *TablesDb) UpsertRow(DatabaseId string, TableId string, RowId string, optionalSetters ...UpsertRowOption)(*models.Row, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{rowId}", RowId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/rows/{rowId}")
	options := UpsertRowOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	params["rowId"] = RowId
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
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

		parsed := models.Row{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Row
	parsed, ok := resp.Result.(models.Row)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateRowOptions struct {
	Data interface{}
	Permissions []string
	enabledSetters map[string]bool
}
func (options UpdateRowOptions) New() *UpdateRowOptions {
	options.enabledSetters = map[string]bool{
		"Data": false,
		"Permissions": false,
	}
	return &options
}
type UpdateRowOption func(*UpdateRowOptions)
func (srv *TablesDb) WithUpdateRowData(v interface{}) UpdateRowOption {
	return func(o *UpdateRowOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func (srv *TablesDb) WithUpdateRowPermissions(v []string) UpdateRowOption {
	return func(o *UpdateRowOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
							
// UpdateRow update a row by its unique ID. Using the patch method you can
// pass only specific fields that will get updated.
func (srv *TablesDb) UpdateRow(DatabaseId string, TableId string, RowId string, optionalSetters ...UpdateRowOption)(*models.Row, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{rowId}", RowId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/rows/{rowId}")
	options := UpdateRowOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	params["rowId"] = RowId
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
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

		parsed := models.Row{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Row
	parsed, ok := resp.Result.(models.Row)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// DeleteRow delete a row by its unique ID.
func (srv *TablesDb) DeleteRow(DatabaseId string, TableId string, RowId string)(*interface{}, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{rowId}", RowId)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/rows/{rowId}")
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	params["rowId"] = RowId
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
type DecrementRowColumnOptions struct {
	Value float64
	Min float64
	enabledSetters map[string]bool
}
func (options DecrementRowColumnOptions) New() *DecrementRowColumnOptions {
	options.enabledSetters = map[string]bool{
		"Value": false,
		"Min": false,
	}
	return &options
}
type DecrementRowColumnOption func(*DecrementRowColumnOptions)
func (srv *TablesDb) WithDecrementRowColumnValue(v float64) DecrementRowColumnOption {
	return func(o *DecrementRowColumnOptions) {
		o.Value = v
		o.enabledSetters["Value"] = true
	}
}
func (srv *TablesDb) WithDecrementRowColumnMin(v float64) DecrementRowColumnOption {
	return func(o *DecrementRowColumnOptions) {
		o.Min = v
		o.enabledSetters["Min"] = true
	}
}
									
// DecrementRowColumn decrement a specific column of a row by a given value.
func (srv *TablesDb) DecrementRowColumn(DatabaseId string, TableId string, RowId string, Column string, optionalSetters ...DecrementRowColumnOption)(*models.Row, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{rowId}", RowId, "{column}", Column)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/rows/{rowId}/{column}/decrement")
	options := DecrementRowColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	params["rowId"] = RowId
	params["column"] = Column
	if options.enabledSetters["Value"] {
		params["value"] = options.Value
	}
	if options.enabledSetters["Min"] {
		params["min"] = options.Min
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

		parsed := models.Row{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Row
	parsed, ok := resp.Result.(models.Row)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type IncrementRowColumnOptions struct {
	Value float64
	Max float64
	enabledSetters map[string]bool
}
func (options IncrementRowColumnOptions) New() *IncrementRowColumnOptions {
	options.enabledSetters = map[string]bool{
		"Value": false,
		"Max": false,
	}
	return &options
}
type IncrementRowColumnOption func(*IncrementRowColumnOptions)
func (srv *TablesDb) WithIncrementRowColumnValue(v float64) IncrementRowColumnOption {
	return func(o *IncrementRowColumnOptions) {
		o.Value = v
		o.enabledSetters["Value"] = true
	}
}
func (srv *TablesDb) WithIncrementRowColumnMax(v float64) IncrementRowColumnOption {
	return func(o *IncrementRowColumnOptions) {
		o.Max = v
		o.enabledSetters["Max"] = true
	}
}
									
// IncrementRowColumn increment a specific column of a row by a given value.
func (srv *TablesDb) IncrementRowColumn(DatabaseId string, TableId string, RowId string, Column string, optionalSetters ...IncrementRowColumnOption)(*models.Row, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{tableId}", TableId, "{rowId}", RowId, "{column}", Column)
	path := r.Replace("/tablesdb/{databaseId}/tables/{tableId}/rows/{rowId}/{column}/increment")
	options := IncrementRowColumnOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["databaseId"] = DatabaseId
	params["tableId"] = TableId
	params["rowId"] = RowId
	params["column"] = Column
	if options.enabledSetters["Value"] {
		params["value"] = options.Value
	}
	if options.enabledSetters["Max"] {
		params["max"] = options.Max
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

		parsed := models.Row{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Row
	parsed, ok := resp.Result.(models.Row)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
