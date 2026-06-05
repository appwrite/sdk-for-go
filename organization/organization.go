package organization

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/v5/client"
	"github.com/appwrite/sdk-for-go/v5/models"
	"strings"
)

// Organization service
type Organization struct {
	client client.Client
}

func New(clt client.Client) *Organization {
	return &Organization{
		client: clt,
	}
}

type ListKeysOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListKeysOptions) New() *ListKeysOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListKeysOption func(*ListKeysOptions)
func (srv *Organization) WithListKeysQueries(v []string) ListKeysOption {
	return func(o *ListKeysOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Organization) WithListKeysTotal(v bool) ListKeysOption {
	return func(o *ListKeysOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListKeys get a list of all API keys from the current organization.
func (srv *Organization) ListKeys(optionalSetters ...ListKeysOption)(*models.KeyList, error) {
	path := "/organization/keys"
	options := ListKeysOptions{}.New()
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
		"X-Appwrite-Project": srv.client.Config["project"],
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.KeyList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.KeyList
	parsed, ok := resp.Result.(models.KeyList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateKeyOptions struct {
	Expire string
	enabledSetters map[string]bool
}
func (options CreateKeyOptions) New() *CreateKeyOptions {
	options.enabledSetters = map[string]bool{
		"Expire": false,
	}
	return &options
}
type CreateKeyOption func(*CreateKeyOptions)
func (srv *Organization) WithCreateKeyExpire(v string) CreateKeyOption {
	return func(o *CreateKeyOptions) {
		o.Expire = v
		o.enabledSetters["Expire"] = true
	}
}
							
// CreateKey create a new organization API key.
func (srv *Organization) CreateKey(KeyId string, Name string, Scopes []string, optionalSetters ...CreateKeyOption)(*models.Key, error) {
	path := "/organization/keys"
	options := CreateKeyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["keyId"] = KeyId
	params["name"] = Name
	params["scopes"] = Scopes
	if options.enabledSetters["Expire"] {
		params["expire"] = options.Expire
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Key{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Key
	parsed, ok := resp.Result.(models.Key)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetKey get a key by its unique ID. This endpoint returns details about a
// specific API key in your organization including its scopes.
func (srv *Organization) GetKey(KeyId string)(*models.Key, error) {
	r := strings.NewReplacer("{keyId}", KeyId)
	path := r.Replace("/organization/keys/{keyId}")
	params := map[string]interface{}{}
	params["keyId"] = KeyId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Key{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Key
	parsed, ok := resp.Result.(models.Key)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateKeyOptions struct {
	Expire string
	enabledSetters map[string]bool
}
func (options UpdateKeyOptions) New() *UpdateKeyOptions {
	options.enabledSetters = map[string]bool{
		"Expire": false,
	}
	return &options
}
type UpdateKeyOption func(*UpdateKeyOptions)
func (srv *Organization) WithUpdateKeyExpire(v string) UpdateKeyOption {
	return func(o *UpdateKeyOptions) {
		o.Expire = v
		o.enabledSetters["Expire"] = true
	}
}
							
// UpdateKey update a key by its unique ID. Use this endpoint to update the
// name, scopes, or expiration time of an API key.
func (srv *Organization) UpdateKey(KeyId string, Name string, Scopes []string, optionalSetters ...UpdateKeyOption)(*models.Key, error) {
	r := strings.NewReplacer("{keyId}", KeyId)
	path := r.Replace("/organization/keys/{keyId}")
	options := UpdateKeyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["keyId"] = KeyId
	params["name"] = Name
	params["scopes"] = Scopes
	if options.enabledSetters["Expire"] {
		params["expire"] = options.Expire
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Key{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Key
	parsed, ok := resp.Result.(models.Key)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteKey delete a key by its unique ID. Once deleted, the key can no
// longer be used to authenticate API calls.
func (srv *Organization) DeleteKey(KeyId string)(*interface{}, error) {
	r := strings.NewReplacer("{keyId}", KeyId)
	path := r.Replace("/organization/keys/{keyId}")
	params := map[string]interface{}{}
	params["keyId"] = KeyId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
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
type ListProjectsOptions struct {
	Queries []string
	Search string
	Total bool
	enabledSetters map[string]bool
}
func (options ListProjectsOptions) New() *ListProjectsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
		"Total": false,
	}
	return &options
}
type ListProjectsOption func(*ListProjectsOptions)
func (srv *Organization) WithListProjectsQueries(v []string) ListProjectsOption {
	return func(o *ListProjectsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Organization) WithListProjectsSearch(v string) ListProjectsOption {
	return func(o *ListProjectsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Organization) WithListProjectsTotal(v bool) ListProjectsOption {
	return func(o *ListProjectsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListProjects get a list of all projects. You can use the query params to
// filter your results.
func (srv *Organization) ListProjects(optionalSetters ...ListProjectsOption)(*models.ProjectList, error) {
	path := "/organization/projects"
	options := ListProjectsOptions{}.New()
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
		"X-Appwrite-Project": srv.client.Config["project"],
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ProjectList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProjectList
	parsed, ok := resp.Result.(models.ProjectList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateProjectOptions struct {
	Region string
	enabledSetters map[string]bool
}
func (options CreateProjectOptions) New() *CreateProjectOptions {
	options.enabledSetters = map[string]bool{
		"Region": false,
	}
	return &options
}
type CreateProjectOption func(*CreateProjectOptions)
func (srv *Organization) WithCreateProjectRegion(v string) CreateProjectOption {
	return func(o *CreateProjectOptions) {
		o.Region = v
		o.enabledSetters["Region"] = true
	}
}
					
// CreateProject create a new project.
func (srv *Organization) CreateProject(ProjectId string, Name string, optionalSetters ...CreateProjectOption)(*models.Project, error) {
	path := "/organization/projects"
	options := CreateProjectOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["projectId"] = ProjectId
	params["name"] = Name
	if options.enabledSetters["Region"] {
		params["region"] = options.Region
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Project{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Project
	parsed, ok := resp.Result.(models.Project)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetProject get a project.
func (srv *Organization) GetProject(ProjectId string)(*models.Project, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/organization/projects/{projectId}")
	params := map[string]interface{}{}
	params["projectId"] = ProjectId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Project{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Project
	parsed, ok := resp.Result.(models.Project)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateProject update a project by its unique ID.
func (srv *Organization) UpdateProject(ProjectId string, Name string)(*models.Project, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/organization/projects/{projectId}")
	params := map[string]interface{}{}
	params["projectId"] = ProjectId
	params["name"] = Name
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Project{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Project
	parsed, ok := resp.Result.(models.Project)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteProject delete a project by its unique ID.
func (srv *Organization) DeleteProject(ProjectId string)(*interface{}, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/organization/projects/{projectId}")
	params := map[string]interface{}{}
	params["projectId"] = ProjectId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
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
