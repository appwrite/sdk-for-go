package sites

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/file"
	"strings"
)

// Sites service
type Sites struct {
	client client.Client
}

func New(clt client.Client) *Sites {
	return &Sites{
		client: clt,
	}
}

type ListOptions struct {
	Queries []string
	Search string
	Total bool
	enabledSetters map[string]bool
}
func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
		"Total": false,
	}
	return &options
}
type ListOption func(*ListOptions)
func (srv *Sites) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Sites) WithListSearch(v string) ListOption {
	return func(o *ListOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Sites) WithListTotal(v bool) ListOption {
	return func(o *ListOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// List get a list of all the project's sites. You can use the query params to
// filter your results.
func (srv *Sites) List(optionalSetters ...ListOption)(*models.SiteList, error) {
	path := "/sites"
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
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.SiteList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.SiteList
	parsed, ok := resp.Result.(models.SiteList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateOptions struct {
	Enabled bool
	Logging bool
	Timeout int
	InstallCommand string
	BuildCommand string
	OutputDirectory string
	Adapter string
	InstallationId string
	FallbackFile string
	ProviderRepositoryId string
	ProviderBranch string
	ProviderSilentMode bool
	ProviderRootDirectory string
	Specification string
	enabledSetters map[string]bool
}
func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{
		"Enabled": false,
		"Logging": false,
		"Timeout": false,
		"InstallCommand": false,
		"BuildCommand": false,
		"OutputDirectory": false,
		"Adapter": false,
		"InstallationId": false,
		"FallbackFile": false,
		"ProviderRepositoryId": false,
		"ProviderBranch": false,
		"ProviderSilentMode": false,
		"ProviderRootDirectory": false,
		"Specification": false,
	}
	return &options
}
type CreateOption func(*CreateOptions)
func (srv *Sites) WithCreateEnabled(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Sites) WithCreateLogging(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.Logging = v
		o.enabledSetters["Logging"] = true
	}
}
func (srv *Sites) WithCreateTimeout(v int) CreateOption {
	return func(o *CreateOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
func (srv *Sites) WithCreateInstallCommand(v string) CreateOption {
	return func(o *CreateOptions) {
		o.InstallCommand = v
		o.enabledSetters["InstallCommand"] = true
	}
}
func (srv *Sites) WithCreateBuildCommand(v string) CreateOption {
	return func(o *CreateOptions) {
		o.BuildCommand = v
		o.enabledSetters["BuildCommand"] = true
	}
}
func (srv *Sites) WithCreateOutputDirectory(v string) CreateOption {
	return func(o *CreateOptions) {
		o.OutputDirectory = v
		o.enabledSetters["OutputDirectory"] = true
	}
}
func (srv *Sites) WithCreateAdapter(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Adapter = v
		o.enabledSetters["Adapter"] = true
	}
}
func (srv *Sites) WithCreateInstallationId(v string) CreateOption {
	return func(o *CreateOptions) {
		o.InstallationId = v
		o.enabledSetters["InstallationId"] = true
	}
}
func (srv *Sites) WithCreateFallbackFile(v string) CreateOption {
	return func(o *CreateOptions) {
		o.FallbackFile = v
		o.enabledSetters["FallbackFile"] = true
	}
}
func (srv *Sites) WithCreateProviderRepositoryId(v string) CreateOption {
	return func(o *CreateOptions) {
		o.ProviderRepositoryId = v
		o.enabledSetters["ProviderRepositoryId"] = true
	}
}
func (srv *Sites) WithCreateProviderBranch(v string) CreateOption {
	return func(o *CreateOptions) {
		o.ProviderBranch = v
		o.enabledSetters["ProviderBranch"] = true
	}
}
func (srv *Sites) WithCreateProviderSilentMode(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.ProviderSilentMode = v
		o.enabledSetters["ProviderSilentMode"] = true
	}
}
func (srv *Sites) WithCreateProviderRootDirectory(v string) CreateOption {
	return func(o *CreateOptions) {
		o.ProviderRootDirectory = v
		o.enabledSetters["ProviderRootDirectory"] = true
	}
}
func (srv *Sites) WithCreateSpecification(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Specification = v
		o.enabledSetters["Specification"] = true
	}
}
									
// Create create a new site.
func (srv *Sites) Create(SiteId string, Name string, Framework string, BuildRuntime string, optionalSetters ...CreateOption)(*models.Site, error) {
	path := "/sites"
	options := CreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["name"] = Name
	params["framework"] = Framework
	params["buildRuntime"] = BuildRuntime
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Logging"] {
		params["logging"] = options.Logging
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
	}
	if options.enabledSetters["InstallCommand"] {
		params["installCommand"] = options.InstallCommand
	}
	if options.enabledSetters["BuildCommand"] {
		params["buildCommand"] = options.BuildCommand
	}
	if options.enabledSetters["OutputDirectory"] {
		params["outputDirectory"] = options.OutputDirectory
	}
	if options.enabledSetters["Adapter"] {
		params["adapter"] = options.Adapter
	}
	if options.enabledSetters["InstallationId"] {
		params["installationId"] = options.InstallationId
	}
	if options.enabledSetters["FallbackFile"] {
		params["fallbackFile"] = options.FallbackFile
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

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Site{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Site
	parsed, ok := resp.Result.(models.Site)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListFrameworks get a list of all frameworks that are currently available on
// the server instance.
func (srv *Sites) ListFrameworks()(*models.FrameworkList, error) {
	path := "/sites/frameworks"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.FrameworkList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.FrameworkList
	parsed, ok := resp.Result.(models.FrameworkList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListSpecifications list allowed site specifications for this instance.
func (srv *Sites) ListSpecifications()(*models.SpecificationList, error) {
	path := "/sites/specifications"
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
	
// Get get a site by its unique ID.
func (srv *Sites) Get(SiteId string)(*models.Site, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/sites/{siteId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Site{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Site
	parsed, ok := resp.Result.(models.Site)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOptions struct {
	Enabled bool
	Logging bool
	Timeout int
	InstallCommand string
	BuildCommand string
	OutputDirectory string
	BuildRuntime string
	Adapter string
	FallbackFile string
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
		"Enabled": false,
		"Logging": false,
		"Timeout": false,
		"InstallCommand": false,
		"BuildCommand": false,
		"OutputDirectory": false,
		"BuildRuntime": false,
		"Adapter": false,
		"FallbackFile": false,
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
func (srv *Sites) WithUpdateEnabled(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Sites) WithUpdateLogging(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Logging = v
		o.enabledSetters["Logging"] = true
	}
}
func (srv *Sites) WithUpdateTimeout(v int) UpdateOption {
	return func(o *UpdateOptions) {
		o.Timeout = v
		o.enabledSetters["Timeout"] = true
	}
}
func (srv *Sites) WithUpdateInstallCommand(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.InstallCommand = v
		o.enabledSetters["InstallCommand"] = true
	}
}
func (srv *Sites) WithUpdateBuildCommand(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.BuildCommand = v
		o.enabledSetters["BuildCommand"] = true
	}
}
func (srv *Sites) WithUpdateOutputDirectory(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.OutputDirectory = v
		o.enabledSetters["OutputDirectory"] = true
	}
}
func (srv *Sites) WithUpdateBuildRuntime(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.BuildRuntime = v
		o.enabledSetters["BuildRuntime"] = true
	}
}
func (srv *Sites) WithUpdateAdapter(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Adapter = v
		o.enabledSetters["Adapter"] = true
	}
}
func (srv *Sites) WithUpdateFallbackFile(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.FallbackFile = v
		o.enabledSetters["FallbackFile"] = true
	}
}
func (srv *Sites) WithUpdateInstallationId(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.InstallationId = v
		o.enabledSetters["InstallationId"] = true
	}
}
func (srv *Sites) WithUpdateProviderRepositoryId(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.ProviderRepositoryId = v
		o.enabledSetters["ProviderRepositoryId"] = true
	}
}
func (srv *Sites) WithUpdateProviderBranch(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.ProviderBranch = v
		o.enabledSetters["ProviderBranch"] = true
	}
}
func (srv *Sites) WithUpdateProviderSilentMode(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.ProviderSilentMode = v
		o.enabledSetters["ProviderSilentMode"] = true
	}
}
func (srv *Sites) WithUpdateProviderRootDirectory(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.ProviderRootDirectory = v
		o.enabledSetters["ProviderRootDirectory"] = true
	}
}
func (srv *Sites) WithUpdateSpecification(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Specification = v
		o.enabledSetters["Specification"] = true
	}
}
							
// Update update site by its unique ID.
func (srv *Sites) Update(SiteId string, Name string, Framework string, optionalSetters ...UpdateOption)(*models.Site, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/sites/{siteId}")
	options := UpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["name"] = Name
	params["framework"] = Framework
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Logging"] {
		params["logging"] = options.Logging
	}
	if options.enabledSetters["Timeout"] {
		params["timeout"] = options.Timeout
	}
	if options.enabledSetters["InstallCommand"] {
		params["installCommand"] = options.InstallCommand
	}
	if options.enabledSetters["BuildCommand"] {
		params["buildCommand"] = options.BuildCommand
	}
	if options.enabledSetters["OutputDirectory"] {
		params["outputDirectory"] = options.OutputDirectory
	}
	if options.enabledSetters["BuildRuntime"] {
		params["buildRuntime"] = options.BuildRuntime
	}
	if options.enabledSetters["Adapter"] {
		params["adapter"] = options.Adapter
	}
	if options.enabledSetters["FallbackFile"] {
		params["fallbackFile"] = options.FallbackFile
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

		parsed := models.Site{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Site
	parsed, ok := resp.Result.(models.Site)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Delete delete a site by its unique ID.
func (srv *Sites) Delete(SiteId string)(*interface{}, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/sites/{siteId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
			
// UpdateSiteDeployment update the site active deployment. Use this endpoint
// to switch the code deployment that should be used when visitor opens your
// site.
func (srv *Sites) UpdateSiteDeployment(SiteId string, DeploymentId string)(*models.Site, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/sites/{siteId}/deployment")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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

		parsed := models.Site{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Site
	parsed, ok := resp.Result.(models.Site)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListDeploymentsOptions struct {
	Queries []string
	Search string
	Total bool
	enabledSetters map[string]bool
}
func (options ListDeploymentsOptions) New() *ListDeploymentsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
		"Total": false,
	}
	return &options
}
type ListDeploymentsOption func(*ListDeploymentsOptions)
func (srv *Sites) WithListDeploymentsQueries(v []string) ListDeploymentsOption {
	return func(o *ListDeploymentsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Sites) WithListDeploymentsSearch(v string) ListDeploymentsOption {
	return func(o *ListDeploymentsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Sites) WithListDeploymentsTotal(v bool) ListDeploymentsOption {
	return func(o *ListDeploymentsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// ListDeployments get a list of all the site's code deployments. You can use
// the query params to filter your results.
func (srv *Sites) ListDeployments(SiteId string, optionalSetters ...ListDeploymentsOption)(*models.DeploymentList, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/sites/{siteId}/deployments")
	options := ListDeploymentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
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
	InstallCommand string
	BuildCommand string
	OutputDirectory string
	enabledSetters map[string]bool
}
func (options CreateDeploymentOptions) New() *CreateDeploymentOptions {
	options.enabledSetters = map[string]bool{
		"InstallCommand": false,
		"BuildCommand": false,
		"OutputDirectory": false,
	}
	return &options
}
type CreateDeploymentOption func(*CreateDeploymentOptions)
func (srv *Sites) WithCreateDeploymentInstallCommand(v string) CreateDeploymentOption {
	return func(o *CreateDeploymentOptions) {
		o.InstallCommand = v
		o.enabledSetters["InstallCommand"] = true
	}
}
func (srv *Sites) WithCreateDeploymentBuildCommand(v string) CreateDeploymentOption {
	return func(o *CreateDeploymentOptions) {
		o.BuildCommand = v
		o.enabledSetters["BuildCommand"] = true
	}
}
func (srv *Sites) WithCreateDeploymentOutputDirectory(v string) CreateDeploymentOption {
	return func(o *CreateDeploymentOptions) {
		o.OutputDirectory = v
		o.enabledSetters["OutputDirectory"] = true
	}
}
							
// CreateDeployment create a new site code deployment. Use this endpoint to
// upload a new version of your site code. To activate your newly uploaded
// code, you'll need to update the site's deployment to use your new
// deployment ID.
func (srv *Sites) CreateDeployment(SiteId string, Code file.InputFile, Activate bool, optionalSetters ...CreateDeploymentOption)(*models.Deployment, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/sites/{siteId}/deployments")
	options := CreateDeploymentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["code"] = Code
	params["activate"] = Activate
	if options.enabledSetters["InstallCommand"] {
		params["installCommand"] = options.InstallCommand
	}
	if options.enabledSetters["BuildCommand"] {
		params["buildCommand"] = options.BuildCommand
	}
	if options.enabledSetters["OutputDirectory"] {
		params["outputDirectory"] = options.OutputDirectory
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
			
// CreateDuplicateDeployment create a new build for an existing site
// deployment. This endpoint allows you to rebuild a deployment with the
// updated site configuration, including its commands and output directory if
// they have been modified. The build process will be queued and executed
// asynchronously. The original deployment's code will be preserved and used
// for the new build.
func (srv *Sites) CreateDuplicateDeployment(SiteId string, DeploymentId string)(*models.Deployment, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/sites/{siteId}/deployments/duplicate")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["deploymentId"] = DeploymentId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
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
type CreateTemplateDeploymentOptions struct {
	Activate bool
	enabledSetters map[string]bool
}
func (options CreateTemplateDeploymentOptions) New() *CreateTemplateDeploymentOptions {
	options.enabledSetters = map[string]bool{
		"Activate": false,
	}
	return &options
}
type CreateTemplateDeploymentOption func(*CreateTemplateDeploymentOptions)
func (srv *Sites) WithCreateTemplateDeploymentActivate(v bool) CreateTemplateDeploymentOption {
	return func(o *CreateTemplateDeploymentOptions) {
		o.Activate = v
		o.enabledSetters["Activate"] = true
	}
}
													
// CreateTemplateDeployment create a deployment based on a template.
// 
// Use this endpoint with combination of
// [listTemplates](https://appwrite.io/docs/products/sites/templates) to find
// the template details.
func (srv *Sites) CreateTemplateDeployment(SiteId string, Repository string, Owner string, RootDirectory string, Type string, Reference string, optionalSetters ...CreateTemplateDeploymentOption)(*models.Deployment, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/sites/{siteId}/deployments/template")
	options := CreateTemplateDeploymentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["repository"] = Repository
	params["owner"] = Owner
	params["rootDirectory"] = RootDirectory
	params["type"] = Type
	params["reference"] = Reference
	if options.enabledSetters["Activate"] {
		params["activate"] = options.Activate
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
type CreateVcsDeploymentOptions struct {
	Activate bool
	enabledSetters map[string]bool
}
func (options CreateVcsDeploymentOptions) New() *CreateVcsDeploymentOptions {
	options.enabledSetters = map[string]bool{
		"Activate": false,
	}
	return &options
}
type CreateVcsDeploymentOption func(*CreateVcsDeploymentOptions)
func (srv *Sites) WithCreateVcsDeploymentActivate(v bool) CreateVcsDeploymentOption {
	return func(o *CreateVcsDeploymentOptions) {
		o.Activate = v
		o.enabledSetters["Activate"] = true
	}
}
							
// CreateVcsDeployment create a deployment when a site is connected to VCS.
// 
// This endpoint lets you create deployment from a branch, commit, or a tag.
func (srv *Sites) CreateVcsDeployment(SiteId string, Type string, Reference string, optionalSetters ...CreateVcsDeploymentOption)(*models.Deployment, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/sites/{siteId}/deployments/vcs")
	options := CreateVcsDeploymentOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["type"] = Type
	params["reference"] = Reference
	if options.enabledSetters["Activate"] {
		params["activate"] = options.Activate
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
			
// GetDeployment get a site deployment by its unique ID.
func (srv *Sites) GetDeployment(SiteId string, DeploymentId string)(*models.Deployment, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{deploymentId}", DeploymentId)
	path := r.Replace("/sites/{siteId}/deployments/{deploymentId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
			
// DeleteDeployment delete a site deployment by its unique ID.
func (srv *Sites) DeleteDeployment(SiteId string, DeploymentId string)(*interface{}, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{deploymentId}", DeploymentId)
	path := r.Replace("/sites/{siteId}/deployments/{deploymentId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
type GetDeploymentDownloadOptions struct {
	Type string
	enabledSetters map[string]bool
}
func (options GetDeploymentDownloadOptions) New() *GetDeploymentDownloadOptions {
	options.enabledSetters = map[string]bool{
		"Type": false,
	}
	return &options
}
type GetDeploymentDownloadOption func(*GetDeploymentDownloadOptions)
func (srv *Sites) WithGetDeploymentDownloadType(v string) GetDeploymentDownloadOption {
	return func(o *GetDeploymentDownloadOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
					
// GetDeploymentDownload get a site deployment content by its unique ID. The
// endpoint response return with a 'Content-Disposition: attachment' header
// that tells the browser to start downloading the file to user downloads
// directory.
func (srv *Sites) GetDeploymentDownload(SiteId string, DeploymentId string, optionalSetters ...GetDeploymentDownloadOption)(*[]byte, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{deploymentId}", DeploymentId)
	path := r.Replace("/sites/{siteId}/deployments/{deploymentId}/download")
	options := GetDeploymentDownloadOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["deploymentId"] = DeploymentId
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
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
			
// UpdateDeploymentStatus cancel an ongoing site deployment build. If the
// build is already in progress, it will be stopped and marked as canceled. If
// the build hasn't started yet, it will be marked as canceled without
// executing. You cannot cancel builds that have already completed (status
// 'ready') or failed. The response includes the final build status and
// details.
func (srv *Sites) UpdateDeploymentStatus(SiteId string, DeploymentId string)(*models.Deployment, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{deploymentId}", DeploymentId)
	path := r.Replace("/sites/{siteId}/deployments/{deploymentId}/status")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
type ListLogsOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListLogsOptions) New() *ListLogsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListLogsOption func(*ListLogsOptions)
func (srv *Sites) WithListLogsQueries(v []string) ListLogsOption {
	return func(o *ListLogsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Sites) WithListLogsTotal(v bool) ListLogsOption {
	return func(o *ListLogsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// ListLogs get a list of all site logs. You can use the query params to
// filter your results.
func (srv *Sites) ListLogs(SiteId string, optionalSetters ...ListLogsOption)(*models.ExecutionList, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/sites/{siteId}/logs")
	options := ListLogsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
			
// GetLog get a site request log by its unique ID.
func (srv *Sites) GetLog(SiteId string, LogId string)(*models.Execution, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{logId}", LogId)
	path := r.Replace("/sites/{siteId}/logs/{logId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["logId"] = LogId
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
			
// DeleteLog delete a site log by its unique ID.
func (srv *Sites) DeleteLog(SiteId string, LogId string)(*interface{}, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{logId}", LogId)
	path := r.Replace("/sites/{siteId}/logs/{logId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["logId"] = LogId
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
	
// ListVariables get a list of all variables of a specific site.
func (srv *Sites) ListVariables(SiteId string)(*models.VariableList, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/sites/{siteId}/variables")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
func (srv *Sites) WithCreateVariableSecret(v bool) CreateVariableOption {
	return func(o *CreateVariableOptions) {
		o.Secret = v
		o.enabledSetters["Secret"] = true
	}
}
							
// CreateVariable create a new site variable. These variables can be accessed
// during build and runtime (server-side rendering) as environment variables.
func (srv *Sites) CreateVariable(SiteId string, Key string, Value string, optionalSetters ...CreateVariableOption)(*models.Variable, error) {
	r := strings.NewReplacer("{siteId}", SiteId)
	path := r.Replace("/sites/{siteId}/variables")
	options := CreateVariableOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
func (srv *Sites) GetVariable(SiteId string, VariableId string)(*models.Variable, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{variableId}", VariableId)
	path := r.Replace("/sites/{siteId}/variables/{variableId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
	Secret bool
	enabledSetters map[string]bool
}
func (options UpdateVariableOptions) New() *UpdateVariableOptions {
	options.enabledSetters = map[string]bool{
		"Value": false,
		"Secret": false,
	}
	return &options
}
type UpdateVariableOption func(*UpdateVariableOptions)
func (srv *Sites) WithUpdateVariableValue(v string) UpdateVariableOption {
	return func(o *UpdateVariableOptions) {
		o.Value = v
		o.enabledSetters["Value"] = true
	}
}
func (srv *Sites) WithUpdateVariableSecret(v bool) UpdateVariableOption {
	return func(o *UpdateVariableOptions) {
		o.Secret = v
		o.enabledSetters["Secret"] = true
	}
}
							
// UpdateVariable update variable by its unique ID.
func (srv *Sites) UpdateVariable(SiteId string, VariableId string, Key string, optionalSetters ...UpdateVariableOption)(*models.Variable, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{variableId}", VariableId)
	path := r.Replace("/sites/{siteId}/variables/{variableId}")
	options := UpdateVariableOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["siteId"] = SiteId
	params["variableId"] = VariableId
	params["key"] = Key
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
func (srv *Sites) DeleteVariable(SiteId string, VariableId string)(*interface{}, error) {
	r := strings.NewReplacer("{siteId}", SiteId, "{variableId}", VariableId)
	path := r.Replace("/sites/{siteId}/variables/{variableId}")
	params := map[string]interface{}{}
	params["siteId"] = SiteId
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
