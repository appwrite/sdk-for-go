package functions

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/file"
	"strings"
)

// Functions service
type Functions struct {
	client client.Client
}

func New(clt client.Client) *Functions {
	return &Functions{
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
func (srv *Functions) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Functions) WithListSearch(v string) ListOption {
	return func(o *ListOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
	
// List get a list of all the project's functions. You can use the query
// params to filter your results.
func (srv *Functions) List(optionalSetters ...ListOption)(*models.FunctionList, error) {
	path := "/functions"
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

		parsed := models.FunctionList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.FunctionList
	parsed, ok := resp.Result.(models.FunctionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateOptions struct {
	Execute []string
	Events []string
	Schedule string
	Timeout int
	Enabled bool
	Logging bool
	Entrypoint string
	Commands string
	Scopes []string
	InstallationId string
	ProviderRepositoryId string
	ProviderBranch string
	ProviderSilentMode bool
	ProviderRootDirectory string
	TemplateRepository string
	TemplateOwner string
	TemplateRootDirectory string
	TemplateVersion string
	Specification string
	enabledSetters map[string]bool
}
func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{
		"Execute": false,
		"Events": false,
		"Schedule": false,
		"Timeout": false,
		"Enabled": false,
		"Logging": false,
		"Entrypoint": false,
		"Commands": false,
		"Scopes": false,
		"InstallationId": false,
		"ProviderRepositoryId": false,
		"ProviderBranch": false,
		"ProviderSilentMode": false,
		"ProviderRootDirectory": false,
		"TemplateRepository": false,
		"TemplateOwner": false,
		"TemplateRootDirectory": false,
		"TemplateVersion": false,
		"Specification": false,
	}
	return &options
}
type CreateOption func(*CreateOptions)
func (srv *Functions) WithCreateExecute(v []string) CreateOption {
	return func(o *CreateOptions) {
		o.Execute = v
		o.enabledSetters["Execute"] = true
	}
}
func (srv *Functions) WithCreateEvents(v []string) CreateOption {
	return func(o *CreateOptions) {
		o.Events = v
		o.enabledSetters["Events"] = true
	}
}
func (srv *Functions) WithCreateSchedule(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Schedule = v
		o.enabledSetters["Schedule"] = true
	}
}
func (srv *Functions) WithCreateTimeout(v int) CreateOption {
	return func(o *CreateOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
func (srv *Functions) WithCreateEnabled(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Functions) WithCreateLogging(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.Logging = v
		o.enabledSetters["Logging"] = true
	}
}
func (srv *Functions) WithCreateEntrypoint(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Entrypoint = v
		o.enabledSetters["Entrypoint"] = true
	}
}
func (srv *Functions) WithCreateCommands(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Commands = v
		o.enabledSetters["Commands"] = true
	}
}
func (srv *Functions) WithCreateScopes(v []string) CreateOption {
	return func(o *CreateOptions) {
		o.Scopes = v
		o.enabledSetters["Scopes"] = true
	}
}
func (srv *Functions) WithCreateInstallationId(v string) CreateOption {
	return func(o *CreateOptions) {
		o.InstallationId = v
		o.enabledSetters["InstallationId"] = true
	}
}
func (srv *Functions) WithCreateProviderRepositoryId(v string) CreateOption {
	return func(o *CreateOptions) {
		o.ProviderRepositoryId = v
		o.enabledSetters["ProviderRepositoryId"] = true
	}
}
func (srv *Functions) WithCreateProviderBranch(v string) CreateOption {
	return func(o *CreateOptions) {
		o.ProviderBranch = v
		o.enabledSetters["ProviderBranch"] = true
	}
}
func (srv *Functions) WithCreateProviderSilentMode(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.ProviderSilentMode = v
		o.enabledSetters["ProviderSilentMode"] = true
	}
}
func (srv *Functions) WithCreateProviderRootDirectory(v string) CreateOption {
	return func(o *CreateOptions) {
		o.ProviderRootDirectory = v
		o.enabledSetters["ProviderRootDirectory"] = true
	}
}
func (srv *Functions) WithCreateTemplateRepository(v string) CreateOption {
	return func(o *CreateOptions) {
		o.TemplateRepository = v
		o.enabledSetters["TemplateRepository"] = true
	}
}
func (srv *Functions) WithCreateTemplateOwner(v string) CreateOption {
	return func(o *CreateOptions) {
		o.TemplateOwner = v
		o.enabledSetters["TemplateOwner"] = true
	}
}
func (srv *Functions) WithCreateTemplateRootDirectory(v string) CreateOption {
	return func(o *CreateOptions) {
		o.TemplateRootDirectory = v
		o.enabledSetters["TemplateRootDirectory"] = true
	}
}
func (srv *Functions) WithCreateTemplateVersion(v string) CreateOption {
	return func(o *CreateOptions) {
		o.TemplateVersion = v
		o.enabledSetters["TemplateVersion"] = true
	}
}
func (srv *Functions) WithCreateSpecification(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Specification = v
		o.enabledSetters["Specification"] = true
	}
}
							
// Create create a new function. You can pass a list of
// [permissions](https://appwrite.io/docs/permissions) to allow different
// project users or team with access to execute the function using the client
// API.
func (srv *Functions) Create(FunctionId string, Name string, Runtime string, optionalSetters ...CreateOption)(*models.Function, error) {
	path := "/functions"
	options := CreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["name"] = Name
	params["runtime"] = Runtime
	if options.enabledSetters["Execute"] {
		params["execute"] = options.Execute
	}
	if options.enabledSetters["Events"] {
		params["events"] = options.Events
	}
	if options.enabledSetters["Schedule"] {
		params["schedule"] = options.Schedule
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Logging"] {
		params["logging"] = options.Logging
	}
	if options.enabledSetters["Entrypoint"] {
		params["entrypoint"] = options.Entrypoint
	}
	if options.enabledSetters["Commands"] {
		params["commands"] = options.Commands
	}
	if options.enabledSetters["Scopes"] {
		params["scopes"] = options.Scopes
	}
	if options.enabledSetters["InstallationId"] {
		params["installationId"] = options.InstallationId
	}
	if options.enabledSetters["ProviderRepositoryId"] {
		params["providerRepositoryId"] = options.ProviderRepositoryId
	}
	if options.enabledSetters["ProviderBranch"] {
		params["providerBranch"] = options.ProviderBranch
	}
	if options.enabledSetters["ProviderSilentMode"] {
		params["providerSilentMode"] = options.ProviderSilentMode
	}
	if options.enabledSetters["ProviderRootDirectory"] {
		params["providerRootDirectory"] = options.ProviderRootDirectory
	}
	if options.enabledSetters["TemplateRepository"] {
		params["templateRepository"] = options.TemplateRepository
	}
	if options.enabledSetters["TemplateOwner"] {
		params["templateOwner"] = options.TemplateOwner
	}
	if options.enabledSetters["TemplateRootDirectory"] {
		params["templateRootDirectory"] = options.TemplateRootDirectory
	}
	if options.enabledSetters["TemplateVersion"] {
		params["templateVersion"] = options.TemplateVersion
	}
	if options.enabledSetters["Specification"] {
		params["specification"] = options.Specification
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

		parsed := models.Function{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Function
	parsed, ok := resp.Result.(models.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListRuntimes get a list of all runtimes that are currently active on your
// instance.
func (srv *Functions) ListRuntimes()(*models.RuntimeList, error) {
	path := "/functions/runtimes"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.RuntimeList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.RuntimeList
	parsed, ok := resp.Result.(models.RuntimeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListSpecifications list allowed function specifications for this instance.
func (srv *Functions) ListSpecifications()(*models.SpecificationList, error) {
	path := "/functions/specifications"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.SpecificationList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.SpecificationList
	parsed, ok := resp.Result.(models.SpecificationList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Get get a function by its unique ID.
func (srv *Functions) Get(FunctionId string)(*models.Function, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Function{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Function
	parsed, ok := resp.Result.(models.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOptions struct {
	Runtime string
	Execute []string
	Events []string
	Schedule string
	Timeout int
	Enabled bool
	Logging bool
	Entrypoint string
	Commands string
	Scopes []string
	InstallationId string
	ProviderRepositoryId string
	ProviderBranch string
	ProviderSilentMode bool
	ProviderRootDirectory string
	Specification string
	enabledSetters map[string]bool
}
func (options UpdateOptions) New() *UpdateOptions {
	options.enabledSetters = map[string]bool{
		"Runtime": false,
		"Execute": false,
		"Events": false,
		"Schedule": false,
		"Timeout": false,
		"Enabled": false,
		"Logging": false,
		"Entrypoint": false,
		"Commands": false,
		"Scopes": false,
		"InstallationId": false,
		"ProviderRepositoryId": false,
		"ProviderBranch": false,
		"ProviderSilentMode": false,
		"ProviderRootDirectory": false,
		"Specification": false,
	}
	return &options
}
type UpdateOption func(*UpdateOptions)
func (srv *Functions) WithUpdateRuntime(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Runtime = v
		o.enabledSetters["Runtime"] = true
	}
}
func (srv *Functions) WithUpdateExecute(v []string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Execute = v
		o.enabledSetters["Execute"] = true
	}
}
func (srv *Functions) WithUpdateEvents(v []string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Events = v
		o.enabledSetters["Events"] = true
	}
}
func (srv *Functions) WithUpdateSchedule(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Schedule = v
		o.enabledSetters["Schedule"] = true
	}
}
func (srv *Functions) WithUpdateTimeout(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
func (srv *Functions) WithUpdateEnabled(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Functions) WithUpdateLogging(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Logging = v
		o.enabledSetters["Logging"] = true
	}
}
func (srv *Functions) WithUpdateEntrypoint(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Entrypoint = v
		o.enabledSetters["Entrypoint"] = true
	}
}
func (srv *Functions) WithUpdateCommands(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Commands = v
		o.enabledSetters["Commands"] = true
	}
}
func (srv *Functions) WithUpdateScopes(v []string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Scopes = v
		o.enabledSetters["Scopes"] = true
	}
}
func (srv *Functions) WithUpdateInstallationId(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.InstallationId = v
		o.enabledSetters["InstallationId"] = true
	}
}
func (srv *Functions) WithUpdateProviderRepositoryId(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.ProviderRepositoryId = v
		o.enabledSetters["ProviderRepositoryId"] = true
	}
}
func (srv *Functions) WithUpdateProviderBranch(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.ProviderBranch = v
		o.enabledSetters["ProviderBranch"] = true
	}
}
func (srv *Functions) WithUpdateProviderSilentMode(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.ProviderSilentMode = v
		o.enabledSetters["ProviderSilentMode"] = true
	}
}
func (srv *Functions) WithUpdateProviderRootDirectory(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.ProviderRootDirectory = v
		o.enabledSetters["ProviderRootDirectory"] = true
	}
}
func (srv *Functions) WithUpdateSpecification(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Specification = v
		o.enabledSetters["Specification"] = true
	}
}
					
// Update update function by its unique ID.
func (srv *Functions) Update(FunctionId string, Name string, optionalSetters ...UpdateOption)(*models.Function, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}")
	options := UpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["name"] = Name
	if options.enabledSetters["Runtime"] {
		params["runtime"] = options.Runtime
	}
	if options.enabledSetters["Execute"] {
		params["execute"] = options.Execute
	}
	if options.enabledSetters["Events"] {
		params["events"] = options.Events
	}
	if options.enabledSetters["Schedule"] {
		params["schedule"] = options.Schedule
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Logging"] {
		params["logging"] = options.Logging
	}
	if options.enabledSetters["Entrypoint"] {
		params["entrypoint"] = options.Entrypoint
	}
	if options.enabledSetters["Commands"] {
		params["commands"] = options.Commands
	}
	if options.enabledSetters["Scopes"] {
		params["scopes"] = options.Scopes
	}
	if options.enabledSetters["InstallationId"] {
		params["installationId"] = options.InstallationId
	}
	if options.enabledSetters["ProviderRepositoryId"] {
		params["providerRepositoryId"] = options.ProviderRepositoryId
	}
	if options.enabledSetters["ProviderBranch"] {
		params["providerBranch"] = options.ProviderBranch
	}
	if options.enabledSetters["ProviderSilentMode"] {
		params["providerSilentMode"] = options.ProviderSilentMode
	}
	if options.enabledSetters["ProviderRootDirectory"] {
		params["providerRootDirectory"] = options.ProviderRootDirectory
	}
	if options.enabledSetters["Specification"] {
		params["specification"] = options.Specification
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

		parsed := models.Function{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Function
	parsed, ok := resp.Result.(models.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Delete delete a function by its unique ID.
func (srv *Functions) Delete(FunctionId string)(*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
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
type ListDeploymentsOptions struct {
	Queries []string
	Search string
	enabledSetters map[string]bool
}
func (options ListDeploymentsOptions) New() *ListDeploymentsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
	}
	return &options
}
type ListDeploymentsOption func(*ListDeploymentsOptions)
func (srv *Functions) WithListDeploymentsQueries(v []string) ListDeploymentsOption {
	return func(o *ListDeploymentsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Functions) WithListDeploymentsSearch(v string) ListDeploymentsOption {
	return func(o *ListDeploymentsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
			
// ListDeployments get a list of all the project's code deployments. You can
// use the query params to filter your results.
func (srv *Functions) ListDeployments(FunctionId string, optionalSetters ...ListDeploymentsOption)(*models.DeploymentList, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/deployments")
	options := ListDeploymentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
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

		parsed := models.DeploymentList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DeploymentList
	parsed, ok := resp.Result.(models.DeploymentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateDeploymentOptions struct {
	Entrypoint string
	Commands string
	enabledSetters map[string]bool
}
func (options CreateDeploymentOptions) New() *CreateDeploymentOptions {
	options.enabledSetters = map[string]bool{
		"Entrypoint": false,
		"Commands": false,
	}
	return &options
}
type CreateDeploymentOption func(*CreateDeploymentOptions)
func (srv *Functions) WithCreateDeploymentEntrypoint(v string) CreateDeploymentOption {
	return func(o *CreateDeploymentOptions) {
		o.Entrypoint = v
		o.enabledSetters["Entrypoint"] = true
	}
}
func (srv *Functions) WithCreateDeploymentCommands(v string) CreateDeploymentOption {
	return func(o *CreateDeploymentOptions) {
		o.Commands = v
		o.enabledSetters["Commands"] = true
	}
}
							
// CreateDeployment create a new function code deployment. Use this endpoint
// to upload a new version of your code function. To execute your newly
// uploaded code, you'll need to update the function's deployment to use your
// new deployment UID.
// 
// This endpoint accepts a tar.gz file compressed with your code. Make sure to
// include any dependencies your code has within the compressed file. You can
// learn more about code packaging in the [Appwrite Cloud Functions
// tutorial](https://appwrite.io/docs/functions).
// 
// Use the "command" param to set the entrypoint used to execute your code.
func (srv *Functions) CreateDeployment(FunctionId string, Code file.InputFile, Activate bool, optionalSetters ...CreateDeploymentOption)(*models.Deployment, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/deployments")
	options := CreateDeploymentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["code"] = Code
	params["activate"] = Activate
	if options.enabledSetters["Entrypoint"] {
		params["entrypoint"] = options.Entrypoint
	}
	if options.enabledSetters["Commands"] {
		params["commands"] = options.Commands
	}
	headers := map[string]interface{}{
		"content-type": "multipart/form-data",
	}

    paramName := "code"


    uploadId := ""

    resp, err := srv.client.FileUpload(path, headers, params, paramName, uploadId)
    if err != nil {
		return nil, err
	}
	var parsed models.Deployment
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Deployment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil
}
			
// GetDeployment get a code deployment by its unique ID.
func (srv *Functions) GetDeployment(FunctionId string, DeploymentId string)(*models.Deployment, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Deployment{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Deployment
	parsed, ok := resp.Result.(models.Deployment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateDeployment update the function code deployment ID using the unique
// function ID. Use this endpoint to switch the code deployment that should be
// executed by the execution endpoint.
func (srv *Functions) UpdateDeployment(FunctionId string, DeploymentId string)(*models.Function, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Function{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Function
	parsed, ok := resp.Result.(models.Function)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// DeleteDeployment delete a code deployment by its unique ID.
func (srv *Functions) DeleteDeployment(FunctionId string, DeploymentId string)(*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
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
type CreateBuildOptions struct {
	BuildId string
	enabledSetters map[string]bool
}
func (options CreateBuildOptions) New() *CreateBuildOptions {
	options.enabledSetters = map[string]bool{
		"BuildId": false,
	}
	return &options
}
type CreateBuildOption func(*CreateBuildOptions)
func (srv *Functions) WithCreateBuildBuildId(v string) CreateBuildOption {
	return func(o *CreateBuildOptions) {
		o.BuildId = v
		o.enabledSetters["BuildId"] = true
	}
}
					
// CreateBuild create a new build for an existing function deployment. This
// endpoint allows you to rebuild a deployment with the updated function
// configuration, including its entrypoint and build commands if they have
// been modified The build process will be queued and executed asynchronously.
// The original deployment's code will be preserved and used for the new
// build.
func (srv *Functions) CreateBuild(FunctionId string, DeploymentId string, optionalSetters ...CreateBuildOption)(*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}/build")
	options := CreateBuildOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
	if options.enabledSetters["BuildId"] {
		params["buildId"] = options.BuildId
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
			
// UpdateDeploymentBuild cancel an ongoing function deployment build. If the
// build is already in progress, it will be stopped and marked as canceled. If
// the build hasn't started yet, it will be marked as canceled without
// executing. You cannot cancel builds that have already completed (status
// 'ready') or failed. The response includes the final build status and
// details.
func (srv *Functions) UpdateDeploymentBuild(FunctionId string, DeploymentId string)(*models.Build, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}/build")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Build{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Build
	parsed, ok := resp.Result.(models.Build)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetDeploymentDownload get a Deployment's contents by its unique ID. This
// endpoint supports range requests for partial or streaming file download.
func (srv *Functions) GetDeploymentDownload(FunctionId string, DeploymentId string)(*[]byte, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{deploymentId}", DeploymentId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}/download")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["deploymentId"] = DeploymentId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		var parsed []byte

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed []byte
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListExecutionsOptions struct {
	Queries []string
	Search string
	enabledSetters map[string]bool
}
func (options ListExecutionsOptions) New() *ListExecutionsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
	}
	return &options
}
type ListExecutionsOption func(*ListExecutionsOptions)
func (srv *Functions) WithListExecutionsQueries(v []string) ListExecutionsOption {
	return func(o *ListExecutionsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Functions) WithListExecutionsSearch(v string) ListExecutionsOption {
	return func(o *ListExecutionsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
			
// ListExecutions get a list of all the current user function execution logs.
// You can use the query params to filter your results.
func (srv *Functions) ListExecutions(FunctionId string, optionalSetters ...ListExecutionsOption)(*models.ExecutionList, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/executions")
	options := ListExecutionsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
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

		parsed := models.ExecutionList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ExecutionList
	parsed, ok := resp.Result.(models.ExecutionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateExecutionOptions struct {
	Body string
	Async bool
	Path string
	Method string
	Headers interface{}
	ScheduledAt string
	enabledSetters map[string]bool
}
func (options CreateExecutionOptions) New() *CreateExecutionOptions {
	options.enabledSetters = map[string]bool{
		"Body": false,
		"Async": false,
		"Path": false,
		"Method": false,
		"Headers": false,
		"ScheduledAt": false,
	}
	return &options
}
type CreateExecutionOption func(*CreateExecutionOptions)
func (srv *Functions) WithCreateExecutionBody(v string) CreateExecutionOption {
	return func(o *CreateExecutionOptions) {
		o.Body = v
		o.enabledSetters["Body"] = true
	}
}
func (srv *Functions) WithCreateExecutionAsync(v bool) CreateExecutionOption {
	return func(o *CreateExecutionOptions) {
		o.Async = v
		o.enabledSetters["Async"] = true
	}
}
func (srv *Functions) WithCreateExecutionPath(v string) CreateExecutionOption {
	return func(o *CreateExecutionOptions) {
		o.Path = v
		o.enabledSetters["Path"] = true
	}
}
func (srv *Functions) WithCreateExecutionMethod(v string) CreateExecutionOption {
	return func(o *CreateExecutionOptions) {
		o.Method = v
		o.enabledSetters["Method"] = true
	}
}
func (srv *Functions) WithCreateExecutionHeaders(v interface{}) CreateExecutionOption {
	return func(o *CreateExecutionOptions) {
		o.Headers = v
		o.enabledSetters["Headers"] = true
	}
}
func (srv *Functions) WithCreateExecutionScheduledAt(v string) CreateExecutionOption {
	return func(o *CreateExecutionOptions) {
		o.ScheduledAt = v
		o.enabledSetters["ScheduledAt"] = true
	}
}
			
// CreateExecution trigger a function execution. The returned object will
// return you the current execution status. You can ping the `Get Execution`
// endpoint to get updates on the current execution status. Once this endpoint
// is called, your function execution process will start asynchronously.
func (srv *Functions) CreateExecution(FunctionId string, optionalSetters ...CreateExecutionOption)(*models.Execution, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/executions")
	options := CreateExecutionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	if options.enabledSetters["Body"] {
		params["body"] = options.Body
	}
	if options.enabledSetters["Async"] {
		params["async"] = options.Async
	}
	if options.enabledSetters["Path"] {
		params["path"] = options.Path
	}
	if options.enabledSetters["Method"] {
		params["method"] = options.Method
	}
	if options.enabledSetters["Headers"] {
		params["headers"] = options.Headers
	}
	if options.enabledSetters["ScheduledAt"] {
		params["scheduledAt"] = options.ScheduledAt
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

		parsed := models.Execution{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Execution
	parsed, ok := resp.Result.(models.Execution)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetExecution get a function execution log by its unique ID.
func (srv *Functions) GetExecution(FunctionId string, ExecutionId string)(*models.Execution, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{executionId}", ExecutionId)
	path := r.Replace("/functions/{functionId}/executions/{executionId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["executionId"] = ExecutionId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Execution{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Execution
	parsed, ok := resp.Result.(models.Execution)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// DeleteExecution delete a function execution by its unique ID.
func (srv *Functions) DeleteExecution(FunctionId string, ExecutionId string)(*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{executionId}", ExecutionId)
	path := r.Replace("/functions/{functionId}/executions/{executionId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["executionId"] = ExecutionId
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
	
// ListVariables get a list of all variables of a specific function.
func (srv *Functions) ListVariables(FunctionId string)(*models.VariableList, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/variables")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
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
					
// CreateVariable create a new function environment variable. These variables
// can be accessed in the function at runtime as environment variables.
func (srv *Functions) CreateVariable(FunctionId string, Key string, Value string)(*models.Variable, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	path := r.Replace("/functions/{functionId}/variables")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["key"] = Key
	params["value"] = Value
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
func (srv *Functions) GetVariable(FunctionId string, VariableId string)(*models.Variable, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{variableId}", VariableId)
	path := r.Replace("/functions/{functionId}/variables/{variableId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
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
	Value string
	enabledSetters map[string]bool
}
func (options UpdateVariableOptions) New() *UpdateVariableOptions {
	options.enabledSetters = map[string]bool{
		"Value": false,
	}
	return &options
}
type UpdateVariableOption func(*UpdateVariableOptions)
func (srv *Functions) WithUpdateVariableValue(v string) UpdateVariableOption {
	return func(o *UpdateVariableOptions) {
		o.Value = v
		o.enabledSetters["Value"] = true
	}
}
							
// UpdateVariable update variable by its unique ID.
func (srv *Functions) UpdateVariable(FunctionId string, VariableId string, Key string, optionalSetters ...UpdateVariableOption)(*models.Variable, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{variableId}", VariableId)
	path := r.Replace("/functions/{functionId}/variables/{variableId}")
	options := UpdateVariableOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
	params["variableId"] = VariableId
	params["key"] = Key
	if options.enabledSetters["Value"] {
		params["value"] = options.Value
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
func (srv *Functions) DeleteVariable(FunctionId string, VariableId string)(*interface{}, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{variableId}", VariableId)
	path := r.Replace("/functions/{functionId}/variables/{variableId}")
	params := map[string]interface{}{}
	params["functionId"] = FunctionId
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
