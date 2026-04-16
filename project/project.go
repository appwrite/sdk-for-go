package project

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/v3/client"
	"github.com/appwrite/sdk-for-go/v3/models"
	"fmt"
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
func (srv *Project) WithListKeysQueries(v []string) ListKeysOption {
	return func(o *ListKeysOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Project) WithListKeysTotal(v bool) ListKeysOption {
	return func(o *ListKeysOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListKeys get a list of all API keys from the current project.
func (srv *Project) ListKeys(optionalSetters ...ListKeysOption)(*models.KeyList, error) {
	path := "/project/keys"
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
func (srv *Project) WithCreateKeyExpire(v string) CreateKeyOption {
	return func(o *CreateKeyOptions) {
		o.Expire = v
		o.enabledSetters["Expire"] = true
	}
}
							
// CreateKey create a new API key. It's recommended to have multiple API keys
// with strict scopes for separate functions within your project.
func (srv *Project) CreateKey(KeyId string, Name string, Scopes []string, optionalSetters ...CreateKeyOption)(*models.Key, error) {
	path := "/project/keys"
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
	
// GetKey get a key by its unique ID.
func (srv *Project) GetKey(KeyId string)(*models.Key, error) {
	r := strings.NewReplacer("{keyId}", KeyId)
	path := r.Replace("/project/keys/{keyId}")
	params := map[string]interface{}{}
	params["keyId"] = KeyId
	headers := map[string]interface{}{
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
func (srv *Project) WithUpdateKeyExpire(v string) UpdateKeyOption {
	return func(o *UpdateKeyOptions) {
		o.Expire = v
		o.enabledSetters["Expire"] = true
	}
}
							
// UpdateKey update a key by its unique ID. Use this endpoint to update the
// name, scopes, or expiration time of an API key.
func (srv *Project) UpdateKey(KeyId string, Name string, Scopes []string, optionalSetters ...UpdateKeyOption)(*models.Key, error) {
	r := strings.NewReplacer("{keyId}", KeyId)
	path := r.Replace("/project/keys/{keyId}")
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
func (srv *Project) DeleteKey(KeyId string)(*interface{}, error) {
	r := strings.NewReplacer("{keyId}", KeyId)
	path := r.Replace("/project/keys/{keyId}")
	params := map[string]interface{}{}
	params["keyId"] = KeyId
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
	
// UpdateLabels update the project labels. Labels can be used to easily filter
// projects in an organization.
func (srv *Project) UpdateLabels(Labels []string)(*models.Project, error) {
	path := "/project/labels"
	params := map[string]interface{}{}
	params["labels"] = Labels
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
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
type ListPlatformsOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListPlatformsOptions) New() *ListPlatformsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListPlatformsOption func(*ListPlatformsOptions)
func (srv *Project) WithListPlatformsQueries(v []string) ListPlatformsOption {
	return func(o *ListPlatformsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Project) WithListPlatformsTotal(v bool) ListPlatformsOption {
	return func(o *ListPlatformsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListPlatforms get a list of all platforms in the project. This endpoint
// returns an array of all platforms and their configurations.
func (srv *Project) ListPlatforms(optionalSetters ...ListPlatformsOption)(*models.PlatformList, error) {
	path := "/project/platforms"
	options := ListPlatformsOptions{}.New()
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

		parsed := models.PlatformList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PlatformList
	parsed, ok := resp.Result.(models.PlatformList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// CreateAndroidPlatform create a new Android platform for your project. Use
// this endpoint to register a new Android platform where your users will run
// your application which will interact with the Appwrite API.
func (srv *Project) CreateAndroidPlatform(PlatformId string, Name string, ApplicationId string)(*models.PlatformAndroid, error) {
	path := "/project/platforms/android"
	params := map[string]interface{}{}
	params["platformId"] = PlatformId
	params["name"] = Name
	params["applicationId"] = ApplicationId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PlatformAndroid{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PlatformAndroid
	parsed, ok := resp.Result.(models.PlatformAndroid)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// UpdateAndroidPlatform update an Android platform by its unique ID. Use this
// endpoint to update the platform's name or application ID.
func (srv *Project) UpdateAndroidPlatform(PlatformId string, Name string, ApplicationId string)(*models.PlatformAndroid, error) {
	r := strings.NewReplacer("{platformId}", PlatformId)
	path := r.Replace("/project/platforms/android/{platformId}")
	params := map[string]interface{}{}
	params["platformId"] = PlatformId
	params["name"] = Name
	params["applicationId"] = ApplicationId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PlatformAndroid{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PlatformAndroid
	parsed, ok := resp.Result.(models.PlatformAndroid)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// CreateApplePlatform create a new Apple platform for your project. Use this
// endpoint to register a new Apple platform where your users will run your
// application which will interact with the Appwrite API.
func (srv *Project) CreateApplePlatform(PlatformId string, Name string, BundleIdentifier string)(*models.PlatformApple, error) {
	path := "/project/platforms/apple"
	params := map[string]interface{}{}
	params["platformId"] = PlatformId
	params["name"] = Name
	params["bundleIdentifier"] = BundleIdentifier
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PlatformApple{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PlatformApple
	parsed, ok := resp.Result.(models.PlatformApple)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// UpdateApplePlatform update an Apple platform by its unique ID. Use this
// endpoint to update the platform's name or bundle identifier.
func (srv *Project) UpdateApplePlatform(PlatformId string, Name string, BundleIdentifier string)(*models.PlatformApple, error) {
	r := strings.NewReplacer("{platformId}", PlatformId)
	path := r.Replace("/project/platforms/apple/{platformId}")
	params := map[string]interface{}{}
	params["platformId"] = PlatformId
	params["name"] = Name
	params["bundleIdentifier"] = BundleIdentifier
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PlatformApple{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PlatformApple
	parsed, ok := resp.Result.(models.PlatformApple)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// CreateLinuxPlatform create a new Linux platform for your project. Use this
// endpoint to register a new Linux platform where your users will run your
// application which will interact with the Appwrite API.
func (srv *Project) CreateLinuxPlatform(PlatformId string, Name string, PackageName string)(*models.PlatformLinux, error) {
	path := "/project/platforms/linux"
	params := map[string]interface{}{}
	params["platformId"] = PlatformId
	params["name"] = Name
	params["packageName"] = PackageName
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PlatformLinux{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PlatformLinux
	parsed, ok := resp.Result.(models.PlatformLinux)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// UpdateLinuxPlatform update a Linux platform by its unique ID. Use this
// endpoint to update the platform's name or package name.
func (srv *Project) UpdateLinuxPlatform(PlatformId string, Name string, PackageName string)(*models.PlatformLinux, error) {
	r := strings.NewReplacer("{platformId}", PlatformId)
	path := r.Replace("/project/platforms/linux/{platformId}")
	params := map[string]interface{}{}
	params["platformId"] = PlatformId
	params["name"] = Name
	params["packageName"] = PackageName
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PlatformLinux{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PlatformLinux
	parsed, ok := resp.Result.(models.PlatformLinux)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// CreateWebPlatform create a new web platform for your project. Use this
// endpoint to register a new platform where your users will run your
// application which will interact with the Appwrite API.
func (srv *Project) CreateWebPlatform(PlatformId string, Name string, Hostname string)(*models.PlatformWeb, error) {
	path := "/project/platforms/web"
	params := map[string]interface{}{}
	params["platformId"] = PlatformId
	params["name"] = Name
	params["hostname"] = Hostname
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PlatformWeb{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PlatformWeb
	parsed, ok := resp.Result.(models.PlatformWeb)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// UpdateWebPlatform update a web platform by its unique ID. Use this endpoint
// to update the platform's name or hostname.
func (srv *Project) UpdateWebPlatform(PlatformId string, Name string, Hostname string)(*models.PlatformWeb, error) {
	r := strings.NewReplacer("{platformId}", PlatformId)
	path := r.Replace("/project/platforms/web/{platformId}")
	params := map[string]interface{}{}
	params["platformId"] = PlatformId
	params["name"] = Name
	params["hostname"] = Hostname
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PlatformWeb{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PlatformWeb
	parsed, ok := resp.Result.(models.PlatformWeb)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// CreateWindowsPlatform create a new Windows platform for your project. Use
// this endpoint to register a new Windows platform where your users will run
// your application which will interact with the Appwrite API.
func (srv *Project) CreateWindowsPlatform(PlatformId string, Name string, PackageIdentifierName string)(*models.PlatformWindows, error) {
	path := "/project/platforms/windows"
	params := map[string]interface{}{}
	params["platformId"] = PlatformId
	params["name"] = Name
	params["packageIdentifierName"] = PackageIdentifierName
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PlatformWindows{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PlatformWindows
	parsed, ok := resp.Result.(models.PlatformWindows)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// UpdateWindowsPlatform update a Windows platform by its unique ID. Use this
// endpoint to update the platform's name or package identifier name.
func (srv *Project) UpdateWindowsPlatform(PlatformId string, Name string, PackageIdentifierName string)(*models.PlatformWindows, error) {
	r := strings.NewReplacer("{platformId}", PlatformId)
	path := r.Replace("/project/platforms/windows/{platformId}")
	params := map[string]interface{}{}
	params["platformId"] = PlatformId
	params["name"] = Name
	params["packageIdentifierName"] = PackageIdentifierName
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PlatformWindows{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PlatformWindows
	parsed, ok := resp.Result.(models.PlatformWindows)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetPlatform get a platform by its unique ID. This endpoint returns the
// platform's details, including its name, type, and key configurations.
func (srv *Project) GetPlatform(PlatformId string)(models.Model, error) {
	r := strings.NewReplacer("{platformId}", PlatformId)
	path := r.Replace("/project/platforms/{platformId}")
	params := map[string]interface{}{}
	params["platformId"] = PlatformId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		var response map[string]interface{}
		if err := json.Unmarshal(bytes, &response); err != nil {
			return nil, err
		}
		if fmt.Sprint(response["type"]) == "web" {
			parsed := models.PlatformWeb{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["type"]) == "apple" {
			parsed := models.PlatformApple{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["type"]) == "android" {
			parsed := models.PlatformAndroid{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["type"]) == "windows" {
			parsed := models.PlatformWindows{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["type"]) == "linux" {
			parsed := models.PlatformLinux{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}

		return nil, errors.New("unable to match response to any expected response model")
	}
	parsed, ok := resp.Result.(models.Model)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return parsed, nil

}
	
// DeletePlatform delete a platform by its unique ID. This endpoint removes
// the platform and all its configurations from the project.
func (srv *Project) DeletePlatform(PlatformId string)(*interface{}, error) {
	r := strings.NewReplacer("{platformId}", PlatformId)
	path := r.Replace("/project/platforms/{platformId}")
	params := map[string]interface{}{}
	params["platformId"] = PlatformId
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
			
// UpdateProtocolStatus update the status of a specific protocol. Use this
// endpoint to enable or disable a protocol in your project.
func (srv *Project) UpdateProtocolStatus(ProtocolId string, Enabled bool)(*models.Project, error) {
	r := strings.NewReplacer("{protocolId}", ProtocolId)
	path := r.Replace("/project/protocols/{protocolId}/status")
	params := map[string]interface{}{}
	params["protocolId"] = ProtocolId
	params["enabled"] = Enabled
	headers := map[string]interface{}{
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
			
// UpdateServiceStatus update the status of a specific service. Use this
// endpoint to enable or disable a service in your project.
func (srv *Project) UpdateServiceStatus(ServiceId string, Enabled bool)(*models.Project, error) {
	r := strings.NewReplacer("{serviceId}", ServiceId)
	path := r.Replace("/project/services/{serviceId}/status")
	params := map[string]interface{}{}
	params["serviceId"] = ServiceId
	params["enabled"] = Enabled
	headers := map[string]interface{}{
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
