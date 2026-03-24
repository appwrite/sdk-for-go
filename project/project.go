package project

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// Project service
type Project struct {
	client client.Client
}

func New(clt client.Client) *Project {
	return &Project{
		client: clt,
	}
}

type ListVariablesOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListVariablesOptions) New() *ListVariablesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListVariablesOption func(*ListVariablesOptions)
func (srv *Project) WithListVariablesQueries(v []string) ListVariablesOption {
	return func(o *ListVariablesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Project) WithListVariablesTotal(v bool) ListVariablesOption {
	return func(o *ListVariablesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListVariables get a list of all project environment variables.
func (srv *Project) ListVariables(optionalSetters ...ListVariablesOption)(*models.VariableList, error) {
	path := "/project/variables"
	options := ListVariablesOptions{}.New()
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

		parsed := models.VariableList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VariableList
	parsed, ok := resp.Result.(models.VariableList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateVariableOptions struct {
	Secret bool
	enabledSetters map[string]bool
}
func (options CreateVariableOptions) New() *CreateVariableOptions {
	options.enabledSetters = map[string]bool{
		"Secret": false,
	}
	return &options
}
type CreateVariableOption func(*CreateVariableOptions)
func (srv *Project) WithCreateVariableSecret(v bool) CreateVariableOption {
	return func(o *CreateVariableOptions) {
		o.Secret = v
		o.enabledSetters["Secret"] = true
	}
}
							
// CreateVariable create a new project environment variable. These variables
// can be accessed by all functions and sites in the project.
func (srv *Project) CreateVariable(VariableId string, Key string, Value string, optionalSetters ...CreateVariableOption)(*models.Variable, error) {
	path := "/project/variables"
	options := CreateVariableOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["variableId"] = VariableId
	params["key"] = Key
	params["value"] = Value
	if options.enabledSetters["Secret"] {
		params["secret"] = options.Secret
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

		parsed := models.Variable{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Variable
	parsed, ok := resp.Result.(models.Variable)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetVariable get a variable by its unique ID.
func (srv *Project) GetVariable(VariableId string)(*models.Variable, error) {
	r := strings.NewReplacer("{variableId}", VariableId)
	path := r.Replace("/project/variables/{variableId}")
	params := map[string]interface{}{}
	params["variableId"] = VariableId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Variable{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Variable
	parsed, ok := resp.Result.(models.Variable)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateVariableOptions struct {
	Key string
	Value string
	Secret bool
	enabledSetters map[string]bool
}
func (options UpdateVariableOptions) New() *UpdateVariableOptions {
	options.enabledSetters = map[string]bool{
		"Key": false,
		"Value": false,
		"Secret": false,
	}
	return &options
}
type UpdateVariableOption func(*UpdateVariableOptions)
func (srv *Project) WithUpdateVariableKey(v string) UpdateVariableOption {
	return func(o *UpdateVariableOptions) {
		o.Key = v
		o.enabledSetters["Key"] = true
	}
}
func (srv *Project) WithUpdateVariableValue(v string) UpdateVariableOption {
	return func(o *UpdateVariableOptions) {
		o.Value = v
		o.enabledSetters["Value"] = true
	}
}
func (srv *Project) WithUpdateVariableSecret(v bool) UpdateVariableOption {
	return func(o *UpdateVariableOptions) {
		o.Secret = v
		o.enabledSetters["Secret"] = true
	}
}
			
// UpdateVariable update variable by its unique ID.
func (srv *Project) UpdateVariable(VariableId string, optionalSetters ...UpdateVariableOption)(*models.Variable, error) {
	r := strings.NewReplacer("{variableId}", VariableId)
	path := r.Replace("/project/variables/{variableId}")
	options := UpdateVariableOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["variableId"] = VariableId
	if options.enabledSetters["Key"] {
		params["key"] = options.Key
	}
	if options.enabledSetters["Value"] {
		params["value"] = options.Value
	}
	if options.enabledSetters["Secret"] {
		params["secret"] = options.Secret
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

		parsed := models.Variable{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Variable
	parsed, ok := resp.Result.(models.Variable)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteVariable delete a variable by its unique ID.
func (srv *Project) DeleteVariable(VariableId string)(*interface{}, error) {
	r := strings.NewReplacer("{variableId}", VariableId)
	path := r.Replace("/project/variables/{variableId}")
	params := map[string]interface{}{}
	params["variableId"] = VariableId
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
