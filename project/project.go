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


// Delete delete a project.
func (srv *Project) Delete()(*interface{}, error) {
	path := "/project"
	params := map[string]interface{}{}
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
			
// UpdateAuthMethod update properties of a specific auth method. Use this
// endpoint to enable or disable a method in your project.
func (srv *Project) UpdateAuthMethod(MethodId string, Enabled bool)(*models.Project, error) {
	r := strings.NewReplacer("{methodId}", MethodId)
	path := r.Replace("/project/auth-methods/{methodId}")
	params := map[string]interface{}{}
	params["methodId"] = MethodId
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
// 
// You can also create an ephemeral API key if you need a short-lived key
// instead.
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
			
// CreateEphemeralKey create a new ephemeral API key. It's recommended to have
// multiple API keys with strict scopes for separate functions within your
// project.
// 
// You can also create a standard API key if you need a longer-lived key
// instead.
func (srv *Project) CreateEphemeralKey(Scopes []string, Duration int)(*models.EphemeralKey, error) {
	path := "/project/keys/ephemeral"
	params := map[string]interface{}{}
	params["scopes"] = Scopes
	params["duration"] = Duration
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.EphemeralKey{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.EphemeralKey
	parsed, ok := resp.Result.(models.EphemeralKey)
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
type ListMockPhonesOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListMockPhonesOptions) New() *ListMockPhonesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListMockPhonesOption func(*ListMockPhonesOptions)
func (srv *Project) WithListMockPhonesQueries(v []string) ListMockPhonesOption {
	return func(o *ListMockPhonesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Project) WithListMockPhonesTotal(v bool) ListMockPhonesOption {
	return func(o *ListMockPhonesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListMockPhones get a list of all mock phones in the project. This endpoint
// returns an array of all mock phones and their OTPs.
func (srv *Project) ListMockPhones(optionalSetters ...ListMockPhonesOption)(*models.MockNumberList, error) {
	path := "/project/mock-phones"
	options := ListMockPhonesOptions{}.New()
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

		parsed := models.MockNumberList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MockNumberList
	parsed, ok := resp.Result.(models.MockNumberList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CreateMockPhone create a new mock phone for your project. Use this endpoint
// to register a mock phone number and its sign-in OTP for your testers.
func (srv *Project) CreateMockPhone(Number string, Otp string)(*models.MockNumber, error) {
	path := "/project/mock-phones"
	params := map[string]interface{}{}
	params["number"] = Number
	params["otp"] = Otp
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MockNumber{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MockNumber
	parsed, ok := resp.Result.(models.MockNumber)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetMockPhone get a mock phone by its unique number. This endpoint returns
// the mock phone's OTP.
func (srv *Project) GetMockPhone(Number string)(*models.MockNumber, error) {
	r := strings.NewReplacer("{number}", Number)
	path := r.Replace("/project/mock-phones/{number}")
	params := map[string]interface{}{}
	params["number"] = Number
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MockNumber{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MockNumber
	parsed, ok := resp.Result.(models.MockNumber)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateMockPhone update a mock phone by its unique number. Use this endpoint
// to update the mock phone's OTP.
func (srv *Project) UpdateMockPhone(Number string, Otp string)(*models.MockNumber, error) {
	r := strings.NewReplacer("{number}", Number)
	path := r.Replace("/project/mock-phones/{number}")
	params := map[string]interface{}{}
	params["number"] = Number
	params["otp"] = Otp
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MockNumber{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MockNumber
	parsed, ok := resp.Result.(models.MockNumber)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteMockPhone delete a mock phone by its unique number. This endpoint
// removes the mock phone and its OTP configuration from the project.
func (srv *Project) DeleteMockPhone(Number string)(*interface{}, error) {
	r := strings.NewReplacer("{number}", Number)
	path := r.Replace("/project/mock-phones/{number}")
	params := map[string]interface{}{}
	params["number"] = Number
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
type ListOAuth2ProvidersOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListOAuth2ProvidersOptions) New() *ListOAuth2ProvidersOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListOAuth2ProvidersOption func(*ListOAuth2ProvidersOptions)
func (srv *Project) WithListOAuth2ProvidersQueries(v []string) ListOAuth2ProvidersOption {
	return func(o *ListOAuth2ProvidersOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Project) WithListOAuth2ProvidersTotal(v bool) ListOAuth2ProvidersOption {
	return func(o *ListOAuth2ProvidersOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListOAuth2Providers get a list of all OAuth2 providers supported by the
// server, along with the project's configuration for each. Credential fields
// are write-only and always returned empty.
func (srv *Project) ListOAuth2Providers(optionalSetters ...ListOAuth2ProvidersOption)(*models.OAuth2ProviderList, error) {
	path := "/project/oauth2"
	options := ListOAuth2ProvidersOptions{}.New()
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

		parsed := models.OAuth2ProviderList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2ProviderList
	parsed, ok := resp.Result.(models.OAuth2ProviderList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetOAuth2Provider get a single OAuth2 provider configuration. Credential
// fields (client secret, p8 file, key/team IDs) are write-only and always
// returned empty.
func (srv *Project) GetOAuth2Provider(ProviderId string)(models.Model, error) {
	path := "/project/oauth2/:provider"
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
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
		if fmt.Sprint(response["$id"]) == "github" {
			parsed := models.OAuth2Github{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "discord" {
			parsed := models.OAuth2Discord{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "figma" {
			parsed := models.OAuth2Figma{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "dropbox" {
			parsed := models.OAuth2Dropbox{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "dailymotion" {
			parsed := models.OAuth2Dailymotion{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "bitbucket" {
			parsed := models.OAuth2Bitbucket{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "bitly" {
			parsed := models.OAuth2Bitly{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "box" {
			parsed := models.OAuth2Box{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "autodesk" {
			parsed := models.OAuth2Autodesk{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "google" {
			parsed := models.OAuth2Google{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "zoom" {
			parsed := models.OAuth2Zoom{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "zoho" {
			parsed := models.OAuth2Zoho{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "yandex" {
			parsed := models.OAuth2Yandex{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "x" {
			parsed := models.OAuth2X{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "wordpress" {
			parsed := models.OAuth2WordPress{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "twitch" {
			parsed := models.OAuth2Twitch{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "stripe" {
			parsed := models.OAuth2Stripe{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "spotify" {
			parsed := models.OAuth2Spotify{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "slack" {
			parsed := models.OAuth2Slack{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "podio" {
			parsed := models.OAuth2Podio{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "notion" {
			parsed := models.OAuth2Notion{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "salesforce" {
			parsed := models.OAuth2Salesforce{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "yahoo" {
			parsed := models.OAuth2Yahoo{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "linkedin" {
			parsed := models.OAuth2Linkedin{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "disqus" {
			parsed := models.OAuth2Disqus{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "amazon" {
			parsed := models.OAuth2Amazon{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "etsy" {
			parsed := models.OAuth2Etsy{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "facebook" {
			parsed := models.OAuth2Facebook{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "tradeshiftBox" {
			parsed := models.OAuth2Tradeshift{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "paypalSandbox" {
			parsed := models.OAuth2Paypal{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "gitlab" {
			parsed := models.OAuth2Gitlab{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "authentik" {
			parsed := models.OAuth2Authentik{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "auth0" {
			parsed := models.OAuth2Auth0{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "fusionauth" {
			parsed := models.OAuth2FusionAuth{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "keycloak" {
			parsed := models.OAuth2Keycloak{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "oidc" {
			parsed := models.OAuth2Oidc{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "apple" {
			parsed := models.OAuth2Apple{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "okta" {
			parsed := models.OAuth2Okta{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "kick" {
			parsed := models.OAuth2Kick{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "microsoft" {
			parsed := models.OAuth2Microsoft{}.New(bytes)
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
type UpdateOAuth2AmazonOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2AmazonOptions) New() *UpdateOAuth2AmazonOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2AmazonOption func(*UpdateOAuth2AmazonOptions)
func (srv *Project) WithUpdateOAuth2AmazonClientId(v string) UpdateOAuth2AmazonOption {
	return func(o *UpdateOAuth2AmazonOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2AmazonClientSecret(v string) UpdateOAuth2AmazonOption {
	return func(o *UpdateOAuth2AmazonOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2AmazonEnabled(v bool) UpdateOAuth2AmazonOption {
	return func(o *UpdateOAuth2AmazonOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Amazon update the project OAuth2 Amazon configuration.
func (srv *Project) UpdateOAuth2Amazon(optionalSetters ...UpdateOAuth2AmazonOption)(*models.OAuth2Amazon, error) {
	path := "/project/oauth2/amazon"
	options := UpdateOAuth2AmazonOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Amazon{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Amazon
	parsed, ok := resp.Result.(models.OAuth2Amazon)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2AppleOptions struct {
	ServiceId string
	KeyId string
	TeamId string
	P8File string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2AppleOptions) New() *UpdateOAuth2AppleOptions {
	options.enabledSetters = map[string]bool{
		"ServiceId": false,
		"KeyId": false,
		"TeamId": false,
		"P8File": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2AppleOption func(*UpdateOAuth2AppleOptions)
func (srv *Project) WithUpdateOAuth2AppleServiceId(v string) UpdateOAuth2AppleOption {
	return func(o *UpdateOAuth2AppleOptions) {
		o.ServiceId = v
		o.enabledSetters["ServiceId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2AppleKeyId(v string) UpdateOAuth2AppleOption {
	return func(o *UpdateOAuth2AppleOptions) {
		o.KeyId = v
		o.enabledSetters["KeyId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2AppleTeamId(v string) UpdateOAuth2AppleOption {
	return func(o *UpdateOAuth2AppleOptions) {
		o.TeamId = v
		o.enabledSetters["TeamId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2AppleP8File(v string) UpdateOAuth2AppleOption {
	return func(o *UpdateOAuth2AppleOptions) {
		o.P8File = v
		o.enabledSetters["P8File"] = true
	}
}
func (srv *Project) WithUpdateOAuth2AppleEnabled(v bool) UpdateOAuth2AppleOption {
	return func(o *UpdateOAuth2AppleOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Apple update the project OAuth2 Apple configuration.
func (srv *Project) UpdateOAuth2Apple(optionalSetters ...UpdateOAuth2AppleOption)(*models.OAuth2Apple, error) {
	path := "/project/oauth2/apple"
	options := UpdateOAuth2AppleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ServiceId"] {
		params["serviceId"] = options.ServiceId
	}
	if options.enabledSetters["KeyId"] {
		params["keyId"] = options.KeyId
	}
	if options.enabledSetters["TeamId"] {
		params["teamId"] = options.TeamId
	}
	if options.enabledSetters["P8File"] {
		params["p8File"] = options.P8File
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Apple{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Apple
	parsed, ok := resp.Result.(models.OAuth2Apple)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2Auth0Options struct {
	ClientId string
	ClientSecret string
	Endpoint string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2Auth0Options) New() *UpdateOAuth2Auth0Options {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Endpoint": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2Auth0Option func(*UpdateOAuth2Auth0Options)
func (srv *Project) WithUpdateOAuth2Auth0ClientId(v string) UpdateOAuth2Auth0Option {
	return func(o *UpdateOAuth2Auth0Options) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2Auth0ClientSecret(v string) UpdateOAuth2Auth0Option {
	return func(o *UpdateOAuth2Auth0Options) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2Auth0Endpoint(v string) UpdateOAuth2Auth0Option {
	return func(o *UpdateOAuth2Auth0Options) {
		o.Endpoint = v
		o.enabledSetters["Endpoint"] = true
	}
}
func (srv *Project) WithUpdateOAuth2Auth0Enabled(v bool) UpdateOAuth2Auth0Option {
	return func(o *UpdateOAuth2Auth0Options) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Auth0 update the project OAuth2 Auth0 configuration.
func (srv *Project) UpdateOAuth2Auth0(optionalSetters ...UpdateOAuth2Auth0Option)(*models.OAuth2Auth0, error) {
	path := "/project/oauth2/auth0"
	options := UpdateOAuth2Auth0Options{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Endpoint"] {
		params["endpoint"] = options.Endpoint
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Auth0{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Auth0
	parsed, ok := resp.Result.(models.OAuth2Auth0)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2AuthentikOptions struct {
	ClientId string
	ClientSecret string
	Endpoint string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2AuthentikOptions) New() *UpdateOAuth2AuthentikOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Endpoint": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2AuthentikOption func(*UpdateOAuth2AuthentikOptions)
func (srv *Project) WithUpdateOAuth2AuthentikClientId(v string) UpdateOAuth2AuthentikOption {
	return func(o *UpdateOAuth2AuthentikOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2AuthentikClientSecret(v string) UpdateOAuth2AuthentikOption {
	return func(o *UpdateOAuth2AuthentikOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2AuthentikEndpoint(v string) UpdateOAuth2AuthentikOption {
	return func(o *UpdateOAuth2AuthentikOptions) {
		o.Endpoint = v
		o.enabledSetters["Endpoint"] = true
	}
}
func (srv *Project) WithUpdateOAuth2AuthentikEnabled(v bool) UpdateOAuth2AuthentikOption {
	return func(o *UpdateOAuth2AuthentikOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Authentik update the project OAuth2 Authentik configuration.
func (srv *Project) UpdateOAuth2Authentik(optionalSetters ...UpdateOAuth2AuthentikOption)(*models.OAuth2Authentik, error) {
	path := "/project/oauth2/authentik"
	options := UpdateOAuth2AuthentikOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Endpoint"] {
		params["endpoint"] = options.Endpoint
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Authentik{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Authentik
	parsed, ok := resp.Result.(models.OAuth2Authentik)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2AutodeskOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2AutodeskOptions) New() *UpdateOAuth2AutodeskOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2AutodeskOption func(*UpdateOAuth2AutodeskOptions)
func (srv *Project) WithUpdateOAuth2AutodeskClientId(v string) UpdateOAuth2AutodeskOption {
	return func(o *UpdateOAuth2AutodeskOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2AutodeskClientSecret(v string) UpdateOAuth2AutodeskOption {
	return func(o *UpdateOAuth2AutodeskOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2AutodeskEnabled(v bool) UpdateOAuth2AutodeskOption {
	return func(o *UpdateOAuth2AutodeskOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Autodesk update the project OAuth2 Autodesk configuration.
func (srv *Project) UpdateOAuth2Autodesk(optionalSetters ...UpdateOAuth2AutodeskOption)(*models.OAuth2Autodesk, error) {
	path := "/project/oauth2/autodesk"
	options := UpdateOAuth2AutodeskOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Autodesk{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Autodesk
	parsed, ok := resp.Result.(models.OAuth2Autodesk)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2BitbucketOptions struct {
	Key string
	Secret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2BitbucketOptions) New() *UpdateOAuth2BitbucketOptions {
	options.enabledSetters = map[string]bool{
		"Key": false,
		"Secret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2BitbucketOption func(*UpdateOAuth2BitbucketOptions)
func (srv *Project) WithUpdateOAuth2BitbucketKey(v string) UpdateOAuth2BitbucketOption {
	return func(o *UpdateOAuth2BitbucketOptions) {
		o.Key = v
		o.enabledSetters["Key"] = true
	}
}
func (srv *Project) WithUpdateOAuth2BitbucketSecret(v string) UpdateOAuth2BitbucketOption {
	return func(o *UpdateOAuth2BitbucketOptions) {
		o.Secret = v
		o.enabledSetters["Secret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2BitbucketEnabled(v bool) UpdateOAuth2BitbucketOption {
	return func(o *UpdateOAuth2BitbucketOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Bitbucket update the project OAuth2 Bitbucket configuration.
func (srv *Project) UpdateOAuth2Bitbucket(optionalSetters ...UpdateOAuth2BitbucketOption)(*models.OAuth2Bitbucket, error) {
	path := "/project/oauth2/bitbucket"
	options := UpdateOAuth2BitbucketOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Key"] {
		params["key"] = options.Key
	}
	if options.enabledSetters["Secret"] {
		params["secret"] = options.Secret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Bitbucket{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Bitbucket
	parsed, ok := resp.Result.(models.OAuth2Bitbucket)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2BitlyOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2BitlyOptions) New() *UpdateOAuth2BitlyOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2BitlyOption func(*UpdateOAuth2BitlyOptions)
func (srv *Project) WithUpdateOAuth2BitlyClientId(v string) UpdateOAuth2BitlyOption {
	return func(o *UpdateOAuth2BitlyOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2BitlyClientSecret(v string) UpdateOAuth2BitlyOption {
	return func(o *UpdateOAuth2BitlyOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2BitlyEnabled(v bool) UpdateOAuth2BitlyOption {
	return func(o *UpdateOAuth2BitlyOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Bitly update the project OAuth2 Bitly configuration.
func (srv *Project) UpdateOAuth2Bitly(optionalSetters ...UpdateOAuth2BitlyOption)(*models.OAuth2Bitly, error) {
	path := "/project/oauth2/bitly"
	options := UpdateOAuth2BitlyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Bitly{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Bitly
	parsed, ok := resp.Result.(models.OAuth2Bitly)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2BoxOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2BoxOptions) New() *UpdateOAuth2BoxOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2BoxOption func(*UpdateOAuth2BoxOptions)
func (srv *Project) WithUpdateOAuth2BoxClientId(v string) UpdateOAuth2BoxOption {
	return func(o *UpdateOAuth2BoxOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2BoxClientSecret(v string) UpdateOAuth2BoxOption {
	return func(o *UpdateOAuth2BoxOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2BoxEnabled(v bool) UpdateOAuth2BoxOption {
	return func(o *UpdateOAuth2BoxOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Box update the project OAuth2 Box configuration.
func (srv *Project) UpdateOAuth2Box(optionalSetters ...UpdateOAuth2BoxOption)(*models.OAuth2Box, error) {
	path := "/project/oauth2/box"
	options := UpdateOAuth2BoxOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Box{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Box
	parsed, ok := resp.Result.(models.OAuth2Box)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2DailymotionOptions struct {
	ApiKey string
	ApiSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2DailymotionOptions) New() *UpdateOAuth2DailymotionOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"ApiSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2DailymotionOption func(*UpdateOAuth2DailymotionOptions)
func (srv *Project) WithUpdateOAuth2DailymotionApiKey(v string) UpdateOAuth2DailymotionOption {
	return func(o *UpdateOAuth2DailymotionOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Project) WithUpdateOAuth2DailymotionApiSecret(v string) UpdateOAuth2DailymotionOption {
	return func(o *UpdateOAuth2DailymotionOptions) {
		o.ApiSecret = v
		o.enabledSetters["ApiSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2DailymotionEnabled(v bool) UpdateOAuth2DailymotionOption {
	return func(o *UpdateOAuth2DailymotionOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Dailymotion update the project OAuth2 Dailymotion
// configuration.
func (srv *Project) UpdateOAuth2Dailymotion(optionalSetters ...UpdateOAuth2DailymotionOption)(*models.OAuth2Dailymotion, error) {
	path := "/project/oauth2/dailymotion"
	options := UpdateOAuth2DailymotionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["ApiSecret"] {
		params["apiSecret"] = options.ApiSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Dailymotion{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Dailymotion
	parsed, ok := resp.Result.(models.OAuth2Dailymotion)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2DiscordOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2DiscordOptions) New() *UpdateOAuth2DiscordOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2DiscordOption func(*UpdateOAuth2DiscordOptions)
func (srv *Project) WithUpdateOAuth2DiscordClientId(v string) UpdateOAuth2DiscordOption {
	return func(o *UpdateOAuth2DiscordOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2DiscordClientSecret(v string) UpdateOAuth2DiscordOption {
	return func(o *UpdateOAuth2DiscordOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2DiscordEnabled(v bool) UpdateOAuth2DiscordOption {
	return func(o *UpdateOAuth2DiscordOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Discord update the project OAuth2 Discord configuration.
func (srv *Project) UpdateOAuth2Discord(optionalSetters ...UpdateOAuth2DiscordOption)(*models.OAuth2Discord, error) {
	path := "/project/oauth2/discord"
	options := UpdateOAuth2DiscordOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Discord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Discord
	parsed, ok := resp.Result.(models.OAuth2Discord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2DisqusOptions struct {
	PublicKey string
	SecretKey string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2DisqusOptions) New() *UpdateOAuth2DisqusOptions {
	options.enabledSetters = map[string]bool{
		"PublicKey": false,
		"SecretKey": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2DisqusOption func(*UpdateOAuth2DisqusOptions)
func (srv *Project) WithUpdateOAuth2DisqusPublicKey(v string) UpdateOAuth2DisqusOption {
	return func(o *UpdateOAuth2DisqusOptions) {
		o.PublicKey = v
		o.enabledSetters["PublicKey"] = true
	}
}
func (srv *Project) WithUpdateOAuth2DisqusSecretKey(v string) UpdateOAuth2DisqusOption {
	return func(o *UpdateOAuth2DisqusOptions) {
		o.SecretKey = v
		o.enabledSetters["SecretKey"] = true
	}
}
func (srv *Project) WithUpdateOAuth2DisqusEnabled(v bool) UpdateOAuth2DisqusOption {
	return func(o *UpdateOAuth2DisqusOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Disqus update the project OAuth2 Disqus configuration.
func (srv *Project) UpdateOAuth2Disqus(optionalSetters ...UpdateOAuth2DisqusOption)(*models.OAuth2Disqus, error) {
	path := "/project/oauth2/disqus"
	options := UpdateOAuth2DisqusOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["PublicKey"] {
		params["publicKey"] = options.PublicKey
	}
	if options.enabledSetters["SecretKey"] {
		params["secretKey"] = options.SecretKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Disqus{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Disqus
	parsed, ok := resp.Result.(models.OAuth2Disqus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2DropboxOptions struct {
	AppKey string
	AppSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2DropboxOptions) New() *UpdateOAuth2DropboxOptions {
	options.enabledSetters = map[string]bool{
		"AppKey": false,
		"AppSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2DropboxOption func(*UpdateOAuth2DropboxOptions)
func (srv *Project) WithUpdateOAuth2DropboxAppKey(v string) UpdateOAuth2DropboxOption {
	return func(o *UpdateOAuth2DropboxOptions) {
		o.AppKey = v
		o.enabledSetters["AppKey"] = true
	}
}
func (srv *Project) WithUpdateOAuth2DropboxAppSecret(v string) UpdateOAuth2DropboxOption {
	return func(o *UpdateOAuth2DropboxOptions) {
		o.AppSecret = v
		o.enabledSetters["AppSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2DropboxEnabled(v bool) UpdateOAuth2DropboxOption {
	return func(o *UpdateOAuth2DropboxOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Dropbox update the project OAuth2 Dropbox configuration.
func (srv *Project) UpdateOAuth2Dropbox(optionalSetters ...UpdateOAuth2DropboxOption)(*models.OAuth2Dropbox, error) {
	path := "/project/oauth2/dropbox"
	options := UpdateOAuth2DropboxOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["AppKey"] {
		params["appKey"] = options.AppKey
	}
	if options.enabledSetters["AppSecret"] {
		params["appSecret"] = options.AppSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Dropbox{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Dropbox
	parsed, ok := resp.Result.(models.OAuth2Dropbox)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2EtsyOptions struct {
	KeyString string
	SharedSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2EtsyOptions) New() *UpdateOAuth2EtsyOptions {
	options.enabledSetters = map[string]bool{
		"KeyString": false,
		"SharedSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2EtsyOption func(*UpdateOAuth2EtsyOptions)
func (srv *Project) WithUpdateOAuth2EtsyKeyString(v string) UpdateOAuth2EtsyOption {
	return func(o *UpdateOAuth2EtsyOptions) {
		o.KeyString = v
		o.enabledSetters["KeyString"] = true
	}
}
func (srv *Project) WithUpdateOAuth2EtsySharedSecret(v string) UpdateOAuth2EtsyOption {
	return func(o *UpdateOAuth2EtsyOptions) {
		o.SharedSecret = v
		o.enabledSetters["SharedSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2EtsyEnabled(v bool) UpdateOAuth2EtsyOption {
	return func(o *UpdateOAuth2EtsyOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Etsy update the project OAuth2 Etsy configuration.
func (srv *Project) UpdateOAuth2Etsy(optionalSetters ...UpdateOAuth2EtsyOption)(*models.OAuth2Etsy, error) {
	path := "/project/oauth2/etsy"
	options := UpdateOAuth2EtsyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["KeyString"] {
		params["keyString"] = options.KeyString
	}
	if options.enabledSetters["SharedSecret"] {
		params["sharedSecret"] = options.SharedSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Etsy{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Etsy
	parsed, ok := resp.Result.(models.OAuth2Etsy)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2FacebookOptions struct {
	AppId string
	AppSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2FacebookOptions) New() *UpdateOAuth2FacebookOptions {
	options.enabledSetters = map[string]bool{
		"AppId": false,
		"AppSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2FacebookOption func(*UpdateOAuth2FacebookOptions)
func (srv *Project) WithUpdateOAuth2FacebookAppId(v string) UpdateOAuth2FacebookOption {
	return func(o *UpdateOAuth2FacebookOptions) {
		o.AppId = v
		o.enabledSetters["AppId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2FacebookAppSecret(v string) UpdateOAuth2FacebookOption {
	return func(o *UpdateOAuth2FacebookOptions) {
		o.AppSecret = v
		o.enabledSetters["AppSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2FacebookEnabled(v bool) UpdateOAuth2FacebookOption {
	return func(o *UpdateOAuth2FacebookOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Facebook update the project OAuth2 Facebook configuration.
func (srv *Project) UpdateOAuth2Facebook(optionalSetters ...UpdateOAuth2FacebookOption)(*models.OAuth2Facebook, error) {
	path := "/project/oauth2/facebook"
	options := UpdateOAuth2FacebookOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["AppId"] {
		params["appId"] = options.AppId
	}
	if options.enabledSetters["AppSecret"] {
		params["appSecret"] = options.AppSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Facebook{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Facebook
	parsed, ok := resp.Result.(models.OAuth2Facebook)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2FigmaOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2FigmaOptions) New() *UpdateOAuth2FigmaOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2FigmaOption func(*UpdateOAuth2FigmaOptions)
func (srv *Project) WithUpdateOAuth2FigmaClientId(v string) UpdateOAuth2FigmaOption {
	return func(o *UpdateOAuth2FigmaOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2FigmaClientSecret(v string) UpdateOAuth2FigmaOption {
	return func(o *UpdateOAuth2FigmaOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2FigmaEnabled(v bool) UpdateOAuth2FigmaOption {
	return func(o *UpdateOAuth2FigmaOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Figma update the project OAuth2 Figma configuration.
func (srv *Project) UpdateOAuth2Figma(optionalSetters ...UpdateOAuth2FigmaOption)(*models.OAuth2Figma, error) {
	path := "/project/oauth2/figma"
	options := UpdateOAuth2FigmaOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Figma{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Figma
	parsed, ok := resp.Result.(models.OAuth2Figma)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2FusionAuthOptions struct {
	ClientId string
	ClientSecret string
	Endpoint string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2FusionAuthOptions) New() *UpdateOAuth2FusionAuthOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Endpoint": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2FusionAuthOption func(*UpdateOAuth2FusionAuthOptions)
func (srv *Project) WithUpdateOAuth2FusionAuthClientId(v string) UpdateOAuth2FusionAuthOption {
	return func(o *UpdateOAuth2FusionAuthOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2FusionAuthClientSecret(v string) UpdateOAuth2FusionAuthOption {
	return func(o *UpdateOAuth2FusionAuthOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2FusionAuthEndpoint(v string) UpdateOAuth2FusionAuthOption {
	return func(o *UpdateOAuth2FusionAuthOptions) {
		o.Endpoint = v
		o.enabledSetters["Endpoint"] = true
	}
}
func (srv *Project) WithUpdateOAuth2FusionAuthEnabled(v bool) UpdateOAuth2FusionAuthOption {
	return func(o *UpdateOAuth2FusionAuthOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2FusionAuth update the project OAuth2 FusionAuth configuration.
func (srv *Project) UpdateOAuth2FusionAuth(optionalSetters ...UpdateOAuth2FusionAuthOption)(*models.OAuth2FusionAuth, error) {
	path := "/project/oauth2/fusionauth"
	options := UpdateOAuth2FusionAuthOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Endpoint"] {
		params["endpoint"] = options.Endpoint
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2FusionAuth{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2FusionAuth
	parsed, ok := resp.Result.(models.OAuth2FusionAuth)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2GitHubOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2GitHubOptions) New() *UpdateOAuth2GitHubOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2GitHubOption func(*UpdateOAuth2GitHubOptions)
func (srv *Project) WithUpdateOAuth2GitHubClientId(v string) UpdateOAuth2GitHubOption {
	return func(o *UpdateOAuth2GitHubOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2GitHubClientSecret(v string) UpdateOAuth2GitHubOption {
	return func(o *UpdateOAuth2GitHubOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2GitHubEnabled(v bool) UpdateOAuth2GitHubOption {
	return func(o *UpdateOAuth2GitHubOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2GitHub update the project OAuth2 GitHub configuration.
func (srv *Project) UpdateOAuth2GitHub(optionalSetters ...UpdateOAuth2GitHubOption)(*models.OAuth2Github, error) {
	path := "/project/oauth2/github"
	options := UpdateOAuth2GitHubOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Github{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Github
	parsed, ok := resp.Result.(models.OAuth2Github)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2GitlabOptions struct {
	ApplicationId string
	Secret string
	Endpoint string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2GitlabOptions) New() *UpdateOAuth2GitlabOptions {
	options.enabledSetters = map[string]bool{
		"ApplicationId": false,
		"Secret": false,
		"Endpoint": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2GitlabOption func(*UpdateOAuth2GitlabOptions)
func (srv *Project) WithUpdateOAuth2GitlabApplicationId(v string) UpdateOAuth2GitlabOption {
	return func(o *UpdateOAuth2GitlabOptions) {
		o.ApplicationId = v
		o.enabledSetters["ApplicationId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2GitlabSecret(v string) UpdateOAuth2GitlabOption {
	return func(o *UpdateOAuth2GitlabOptions) {
		o.Secret = v
		o.enabledSetters["Secret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2GitlabEndpoint(v string) UpdateOAuth2GitlabOption {
	return func(o *UpdateOAuth2GitlabOptions) {
		o.Endpoint = v
		o.enabledSetters["Endpoint"] = true
	}
}
func (srv *Project) WithUpdateOAuth2GitlabEnabled(v bool) UpdateOAuth2GitlabOption {
	return func(o *UpdateOAuth2GitlabOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Gitlab update the project OAuth2 Gitlab configuration.
func (srv *Project) UpdateOAuth2Gitlab(optionalSetters ...UpdateOAuth2GitlabOption)(*models.OAuth2Gitlab, error) {
	path := "/project/oauth2/gitlab"
	options := UpdateOAuth2GitlabOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ApplicationId"] {
		params["applicationId"] = options.ApplicationId
	}
	if options.enabledSetters["Secret"] {
		params["secret"] = options.Secret
	}
	if options.enabledSetters["Endpoint"] {
		params["endpoint"] = options.Endpoint
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Gitlab{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Gitlab
	parsed, ok := resp.Result.(models.OAuth2Gitlab)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2GoogleOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2GoogleOptions) New() *UpdateOAuth2GoogleOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2GoogleOption func(*UpdateOAuth2GoogleOptions)
func (srv *Project) WithUpdateOAuth2GoogleClientId(v string) UpdateOAuth2GoogleOption {
	return func(o *UpdateOAuth2GoogleOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2GoogleClientSecret(v string) UpdateOAuth2GoogleOption {
	return func(o *UpdateOAuth2GoogleOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2GoogleEnabled(v bool) UpdateOAuth2GoogleOption {
	return func(o *UpdateOAuth2GoogleOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Google update the project OAuth2 Google configuration.
func (srv *Project) UpdateOAuth2Google(optionalSetters ...UpdateOAuth2GoogleOption)(*models.OAuth2Google, error) {
	path := "/project/oauth2/google"
	options := UpdateOAuth2GoogleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Google{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Google
	parsed, ok := resp.Result.(models.OAuth2Google)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2KeycloakOptions struct {
	ClientId string
	ClientSecret string
	Endpoint string
	RealmName string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2KeycloakOptions) New() *UpdateOAuth2KeycloakOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Endpoint": false,
		"RealmName": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2KeycloakOption func(*UpdateOAuth2KeycloakOptions)
func (srv *Project) WithUpdateOAuth2KeycloakClientId(v string) UpdateOAuth2KeycloakOption {
	return func(o *UpdateOAuth2KeycloakOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2KeycloakClientSecret(v string) UpdateOAuth2KeycloakOption {
	return func(o *UpdateOAuth2KeycloakOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2KeycloakEndpoint(v string) UpdateOAuth2KeycloakOption {
	return func(o *UpdateOAuth2KeycloakOptions) {
		o.Endpoint = v
		o.enabledSetters["Endpoint"] = true
	}
}
func (srv *Project) WithUpdateOAuth2KeycloakRealmName(v string) UpdateOAuth2KeycloakOption {
	return func(o *UpdateOAuth2KeycloakOptions) {
		o.RealmName = v
		o.enabledSetters["RealmName"] = true
	}
}
func (srv *Project) WithUpdateOAuth2KeycloakEnabled(v bool) UpdateOAuth2KeycloakOption {
	return func(o *UpdateOAuth2KeycloakOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Keycloak update the project OAuth2 Keycloak configuration.
func (srv *Project) UpdateOAuth2Keycloak(optionalSetters ...UpdateOAuth2KeycloakOption)(*models.OAuth2Keycloak, error) {
	path := "/project/oauth2/keycloak"
	options := UpdateOAuth2KeycloakOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Endpoint"] {
		params["endpoint"] = options.Endpoint
	}
	if options.enabledSetters["RealmName"] {
		params["realmName"] = options.RealmName
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Keycloak{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Keycloak
	parsed, ok := resp.Result.(models.OAuth2Keycloak)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2KickOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2KickOptions) New() *UpdateOAuth2KickOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2KickOption func(*UpdateOAuth2KickOptions)
func (srv *Project) WithUpdateOAuth2KickClientId(v string) UpdateOAuth2KickOption {
	return func(o *UpdateOAuth2KickOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2KickClientSecret(v string) UpdateOAuth2KickOption {
	return func(o *UpdateOAuth2KickOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2KickEnabled(v bool) UpdateOAuth2KickOption {
	return func(o *UpdateOAuth2KickOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Kick update the project OAuth2 Kick configuration.
func (srv *Project) UpdateOAuth2Kick(optionalSetters ...UpdateOAuth2KickOption)(*models.OAuth2Kick, error) {
	path := "/project/oauth2/kick"
	options := UpdateOAuth2KickOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Kick{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Kick
	parsed, ok := resp.Result.(models.OAuth2Kick)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2LinkedinOptions struct {
	ClientId string
	PrimaryClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2LinkedinOptions) New() *UpdateOAuth2LinkedinOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"PrimaryClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2LinkedinOption func(*UpdateOAuth2LinkedinOptions)
func (srv *Project) WithUpdateOAuth2LinkedinClientId(v string) UpdateOAuth2LinkedinOption {
	return func(o *UpdateOAuth2LinkedinOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2LinkedinPrimaryClientSecret(v string) UpdateOAuth2LinkedinOption {
	return func(o *UpdateOAuth2LinkedinOptions) {
		o.PrimaryClientSecret = v
		o.enabledSetters["PrimaryClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2LinkedinEnabled(v bool) UpdateOAuth2LinkedinOption {
	return func(o *UpdateOAuth2LinkedinOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Linkedin update the project OAuth2 Linkedin configuration.
func (srv *Project) UpdateOAuth2Linkedin(optionalSetters ...UpdateOAuth2LinkedinOption)(*models.OAuth2Linkedin, error) {
	path := "/project/oauth2/linkedin"
	options := UpdateOAuth2LinkedinOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["PrimaryClientSecret"] {
		params["primaryClientSecret"] = options.PrimaryClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Linkedin{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Linkedin
	parsed, ok := resp.Result.(models.OAuth2Linkedin)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2MicrosoftOptions struct {
	ApplicationId string
	ApplicationSecret string
	Tenant string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2MicrosoftOptions) New() *UpdateOAuth2MicrosoftOptions {
	options.enabledSetters = map[string]bool{
		"ApplicationId": false,
		"ApplicationSecret": false,
		"Tenant": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2MicrosoftOption func(*UpdateOAuth2MicrosoftOptions)
func (srv *Project) WithUpdateOAuth2MicrosoftApplicationId(v string) UpdateOAuth2MicrosoftOption {
	return func(o *UpdateOAuth2MicrosoftOptions) {
		o.ApplicationId = v
		o.enabledSetters["ApplicationId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2MicrosoftApplicationSecret(v string) UpdateOAuth2MicrosoftOption {
	return func(o *UpdateOAuth2MicrosoftOptions) {
		o.ApplicationSecret = v
		o.enabledSetters["ApplicationSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2MicrosoftTenant(v string) UpdateOAuth2MicrosoftOption {
	return func(o *UpdateOAuth2MicrosoftOptions) {
		o.Tenant = v
		o.enabledSetters["Tenant"] = true
	}
}
func (srv *Project) WithUpdateOAuth2MicrosoftEnabled(v bool) UpdateOAuth2MicrosoftOption {
	return func(o *UpdateOAuth2MicrosoftOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Microsoft update the project OAuth2 Microsoft configuration.
func (srv *Project) UpdateOAuth2Microsoft(optionalSetters ...UpdateOAuth2MicrosoftOption)(*models.OAuth2Microsoft, error) {
	path := "/project/oauth2/microsoft"
	options := UpdateOAuth2MicrosoftOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ApplicationId"] {
		params["applicationId"] = options.ApplicationId
	}
	if options.enabledSetters["ApplicationSecret"] {
		params["applicationSecret"] = options.ApplicationSecret
	}
	if options.enabledSetters["Tenant"] {
		params["tenant"] = options.Tenant
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Microsoft{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Microsoft
	parsed, ok := resp.Result.(models.OAuth2Microsoft)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2NotionOptions struct {
	OauthClientId string
	OauthClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2NotionOptions) New() *UpdateOAuth2NotionOptions {
	options.enabledSetters = map[string]bool{
		"OauthClientId": false,
		"OauthClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2NotionOption func(*UpdateOAuth2NotionOptions)
func (srv *Project) WithUpdateOAuth2NotionOauthClientId(v string) UpdateOAuth2NotionOption {
	return func(o *UpdateOAuth2NotionOptions) {
		o.OauthClientId = v
		o.enabledSetters["OauthClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2NotionOauthClientSecret(v string) UpdateOAuth2NotionOption {
	return func(o *UpdateOAuth2NotionOptions) {
		o.OauthClientSecret = v
		o.enabledSetters["OauthClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2NotionEnabled(v bool) UpdateOAuth2NotionOption {
	return func(o *UpdateOAuth2NotionOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Notion update the project OAuth2 Notion configuration.
func (srv *Project) UpdateOAuth2Notion(optionalSetters ...UpdateOAuth2NotionOption)(*models.OAuth2Notion, error) {
	path := "/project/oauth2/notion"
	options := UpdateOAuth2NotionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["OauthClientId"] {
		params["oauthClientId"] = options.OauthClientId
	}
	if options.enabledSetters["OauthClientSecret"] {
		params["oauthClientSecret"] = options.OauthClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Notion{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Notion
	parsed, ok := resp.Result.(models.OAuth2Notion)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2OidcOptions struct {
	ClientId string
	ClientSecret string
	WellKnownURL string
	AuthorizationURL string
	TokenURL string
	UserInfoURL string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2OidcOptions) New() *UpdateOAuth2OidcOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"WellKnownURL": false,
		"AuthorizationURL": false,
		"TokenURL": false,
		"UserInfoURL": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2OidcOption func(*UpdateOAuth2OidcOptions)
func (srv *Project) WithUpdateOAuth2OidcClientId(v string) UpdateOAuth2OidcOption {
	return func(o *UpdateOAuth2OidcOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2OidcClientSecret(v string) UpdateOAuth2OidcOption {
	return func(o *UpdateOAuth2OidcOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2OidcWellKnownURL(v string) UpdateOAuth2OidcOption {
	return func(o *UpdateOAuth2OidcOptions) {
		o.WellKnownURL = v
		o.enabledSetters["WellKnownURL"] = true
	}
}
func (srv *Project) WithUpdateOAuth2OidcAuthorizationURL(v string) UpdateOAuth2OidcOption {
	return func(o *UpdateOAuth2OidcOptions) {
		o.AuthorizationURL = v
		o.enabledSetters["AuthorizationURL"] = true
	}
}
func (srv *Project) WithUpdateOAuth2OidcTokenURL(v string) UpdateOAuth2OidcOption {
	return func(o *UpdateOAuth2OidcOptions) {
		o.TokenURL = v
		o.enabledSetters["TokenURL"] = true
	}
}
func (srv *Project) WithUpdateOAuth2OidcUserInfoURL(v string) UpdateOAuth2OidcOption {
	return func(o *UpdateOAuth2OidcOptions) {
		o.UserInfoURL = v
		o.enabledSetters["UserInfoURL"] = true
	}
}
func (srv *Project) WithUpdateOAuth2OidcEnabled(v bool) UpdateOAuth2OidcOption {
	return func(o *UpdateOAuth2OidcOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Oidc update the project OAuth2 Oidc configuration.
func (srv *Project) UpdateOAuth2Oidc(optionalSetters ...UpdateOAuth2OidcOption)(*models.OAuth2Oidc, error) {
	path := "/project/oauth2/oidc"
	options := UpdateOAuth2OidcOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["WellKnownURL"] {
		params["wellKnownURL"] = options.WellKnownURL
	}
	if options.enabledSetters["AuthorizationURL"] {
		params["authorizationURL"] = options.AuthorizationURL
	}
	if options.enabledSetters["TokenURL"] {
		params["tokenURL"] = options.TokenURL
	}
	if options.enabledSetters["UserInfoURL"] {
		params["userInfoURL"] = options.UserInfoURL
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Oidc{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Oidc
	parsed, ok := resp.Result.(models.OAuth2Oidc)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2OktaOptions struct {
	ClientId string
	ClientSecret string
	Domain string
	AuthorizationServerId string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2OktaOptions) New() *UpdateOAuth2OktaOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Domain": false,
		"AuthorizationServerId": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2OktaOption func(*UpdateOAuth2OktaOptions)
func (srv *Project) WithUpdateOAuth2OktaClientId(v string) UpdateOAuth2OktaOption {
	return func(o *UpdateOAuth2OktaOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2OktaClientSecret(v string) UpdateOAuth2OktaOption {
	return func(o *UpdateOAuth2OktaOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2OktaDomain(v string) UpdateOAuth2OktaOption {
	return func(o *UpdateOAuth2OktaOptions) {
		o.Domain = v
		o.enabledSetters["Domain"] = true
	}
}
func (srv *Project) WithUpdateOAuth2OktaAuthorizationServerId(v string) UpdateOAuth2OktaOption {
	return func(o *UpdateOAuth2OktaOptions) {
		o.AuthorizationServerId = v
		o.enabledSetters["AuthorizationServerId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2OktaEnabled(v bool) UpdateOAuth2OktaOption {
	return func(o *UpdateOAuth2OktaOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Okta update the project OAuth2 Okta configuration.
func (srv *Project) UpdateOAuth2Okta(optionalSetters ...UpdateOAuth2OktaOption)(*models.OAuth2Okta, error) {
	path := "/project/oauth2/okta"
	options := UpdateOAuth2OktaOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Domain"] {
		params["domain"] = options.Domain
	}
	if options.enabledSetters["AuthorizationServerId"] {
		params["authorizationServerId"] = options.AuthorizationServerId
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Okta{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Okta
	parsed, ok := resp.Result.(models.OAuth2Okta)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2PaypalOptions struct {
	ClientId string
	SecretKey string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2PaypalOptions) New() *UpdateOAuth2PaypalOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"SecretKey": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2PaypalOption func(*UpdateOAuth2PaypalOptions)
func (srv *Project) WithUpdateOAuth2PaypalClientId(v string) UpdateOAuth2PaypalOption {
	return func(o *UpdateOAuth2PaypalOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2PaypalSecretKey(v string) UpdateOAuth2PaypalOption {
	return func(o *UpdateOAuth2PaypalOptions) {
		o.SecretKey = v
		o.enabledSetters["SecretKey"] = true
	}
}
func (srv *Project) WithUpdateOAuth2PaypalEnabled(v bool) UpdateOAuth2PaypalOption {
	return func(o *UpdateOAuth2PaypalOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Paypal update the project OAuth2 Paypal configuration.
func (srv *Project) UpdateOAuth2Paypal(optionalSetters ...UpdateOAuth2PaypalOption)(*models.OAuth2Paypal, error) {
	path := "/project/oauth2/paypal"
	options := UpdateOAuth2PaypalOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["SecretKey"] {
		params["secretKey"] = options.SecretKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Paypal{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Paypal
	parsed, ok := resp.Result.(models.OAuth2Paypal)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2PaypalSandboxOptions struct {
	ClientId string
	SecretKey string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2PaypalSandboxOptions) New() *UpdateOAuth2PaypalSandboxOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"SecretKey": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2PaypalSandboxOption func(*UpdateOAuth2PaypalSandboxOptions)
func (srv *Project) WithUpdateOAuth2PaypalSandboxClientId(v string) UpdateOAuth2PaypalSandboxOption {
	return func(o *UpdateOAuth2PaypalSandboxOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2PaypalSandboxSecretKey(v string) UpdateOAuth2PaypalSandboxOption {
	return func(o *UpdateOAuth2PaypalSandboxOptions) {
		o.SecretKey = v
		o.enabledSetters["SecretKey"] = true
	}
}
func (srv *Project) WithUpdateOAuth2PaypalSandboxEnabled(v bool) UpdateOAuth2PaypalSandboxOption {
	return func(o *UpdateOAuth2PaypalSandboxOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2PaypalSandbox update the project OAuth2 PaypalSandbox
// configuration.
func (srv *Project) UpdateOAuth2PaypalSandbox(optionalSetters ...UpdateOAuth2PaypalSandboxOption)(*models.OAuth2Paypal, error) {
	path := "/project/oauth2/paypalSandbox"
	options := UpdateOAuth2PaypalSandboxOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["SecretKey"] {
		params["secretKey"] = options.SecretKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Paypal{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Paypal
	parsed, ok := resp.Result.(models.OAuth2Paypal)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2PodioOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2PodioOptions) New() *UpdateOAuth2PodioOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2PodioOption func(*UpdateOAuth2PodioOptions)
func (srv *Project) WithUpdateOAuth2PodioClientId(v string) UpdateOAuth2PodioOption {
	return func(o *UpdateOAuth2PodioOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2PodioClientSecret(v string) UpdateOAuth2PodioOption {
	return func(o *UpdateOAuth2PodioOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2PodioEnabled(v bool) UpdateOAuth2PodioOption {
	return func(o *UpdateOAuth2PodioOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Podio update the project OAuth2 Podio configuration.
func (srv *Project) UpdateOAuth2Podio(optionalSetters ...UpdateOAuth2PodioOption)(*models.OAuth2Podio, error) {
	path := "/project/oauth2/podio"
	options := UpdateOAuth2PodioOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Podio{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Podio
	parsed, ok := resp.Result.(models.OAuth2Podio)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2SalesforceOptions struct {
	CustomerKey string
	CustomerSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2SalesforceOptions) New() *UpdateOAuth2SalesforceOptions {
	options.enabledSetters = map[string]bool{
		"CustomerKey": false,
		"CustomerSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2SalesforceOption func(*UpdateOAuth2SalesforceOptions)
func (srv *Project) WithUpdateOAuth2SalesforceCustomerKey(v string) UpdateOAuth2SalesforceOption {
	return func(o *UpdateOAuth2SalesforceOptions) {
		o.CustomerKey = v
		o.enabledSetters["CustomerKey"] = true
	}
}
func (srv *Project) WithUpdateOAuth2SalesforceCustomerSecret(v string) UpdateOAuth2SalesforceOption {
	return func(o *UpdateOAuth2SalesforceOptions) {
		o.CustomerSecret = v
		o.enabledSetters["CustomerSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2SalesforceEnabled(v bool) UpdateOAuth2SalesforceOption {
	return func(o *UpdateOAuth2SalesforceOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Salesforce update the project OAuth2 Salesforce configuration.
func (srv *Project) UpdateOAuth2Salesforce(optionalSetters ...UpdateOAuth2SalesforceOption)(*models.OAuth2Salesforce, error) {
	path := "/project/oauth2/salesforce"
	options := UpdateOAuth2SalesforceOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["CustomerKey"] {
		params["customerKey"] = options.CustomerKey
	}
	if options.enabledSetters["CustomerSecret"] {
		params["customerSecret"] = options.CustomerSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Salesforce{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Salesforce
	parsed, ok := resp.Result.(models.OAuth2Salesforce)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2SlackOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2SlackOptions) New() *UpdateOAuth2SlackOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2SlackOption func(*UpdateOAuth2SlackOptions)
func (srv *Project) WithUpdateOAuth2SlackClientId(v string) UpdateOAuth2SlackOption {
	return func(o *UpdateOAuth2SlackOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2SlackClientSecret(v string) UpdateOAuth2SlackOption {
	return func(o *UpdateOAuth2SlackOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2SlackEnabled(v bool) UpdateOAuth2SlackOption {
	return func(o *UpdateOAuth2SlackOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Slack update the project OAuth2 Slack configuration.
func (srv *Project) UpdateOAuth2Slack(optionalSetters ...UpdateOAuth2SlackOption)(*models.OAuth2Slack, error) {
	path := "/project/oauth2/slack"
	options := UpdateOAuth2SlackOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Slack{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Slack
	parsed, ok := resp.Result.(models.OAuth2Slack)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2SpotifyOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2SpotifyOptions) New() *UpdateOAuth2SpotifyOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2SpotifyOption func(*UpdateOAuth2SpotifyOptions)
func (srv *Project) WithUpdateOAuth2SpotifyClientId(v string) UpdateOAuth2SpotifyOption {
	return func(o *UpdateOAuth2SpotifyOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2SpotifyClientSecret(v string) UpdateOAuth2SpotifyOption {
	return func(o *UpdateOAuth2SpotifyOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2SpotifyEnabled(v bool) UpdateOAuth2SpotifyOption {
	return func(o *UpdateOAuth2SpotifyOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Spotify update the project OAuth2 Spotify configuration.
func (srv *Project) UpdateOAuth2Spotify(optionalSetters ...UpdateOAuth2SpotifyOption)(*models.OAuth2Spotify, error) {
	path := "/project/oauth2/spotify"
	options := UpdateOAuth2SpotifyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Spotify{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Spotify
	parsed, ok := resp.Result.(models.OAuth2Spotify)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2StripeOptions struct {
	ClientId string
	ApiSecretKey string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2StripeOptions) New() *UpdateOAuth2StripeOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ApiSecretKey": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2StripeOption func(*UpdateOAuth2StripeOptions)
func (srv *Project) WithUpdateOAuth2StripeClientId(v string) UpdateOAuth2StripeOption {
	return func(o *UpdateOAuth2StripeOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2StripeApiSecretKey(v string) UpdateOAuth2StripeOption {
	return func(o *UpdateOAuth2StripeOptions) {
		o.ApiSecretKey = v
		o.enabledSetters["ApiSecretKey"] = true
	}
}
func (srv *Project) WithUpdateOAuth2StripeEnabled(v bool) UpdateOAuth2StripeOption {
	return func(o *UpdateOAuth2StripeOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Stripe update the project OAuth2 Stripe configuration.
func (srv *Project) UpdateOAuth2Stripe(optionalSetters ...UpdateOAuth2StripeOption)(*models.OAuth2Stripe, error) {
	path := "/project/oauth2/stripe"
	options := UpdateOAuth2StripeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ApiSecretKey"] {
		params["apiSecretKey"] = options.ApiSecretKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Stripe{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Stripe
	parsed, ok := resp.Result.(models.OAuth2Stripe)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2TradeshiftOptions struct {
	Oauth2ClientId string
	Oauth2ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2TradeshiftOptions) New() *UpdateOAuth2TradeshiftOptions {
	options.enabledSetters = map[string]bool{
		"Oauth2ClientId": false,
		"Oauth2ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2TradeshiftOption func(*UpdateOAuth2TradeshiftOptions)
func (srv *Project) WithUpdateOAuth2TradeshiftOauth2ClientId(v string) UpdateOAuth2TradeshiftOption {
	return func(o *UpdateOAuth2TradeshiftOptions) {
		o.Oauth2ClientId = v
		o.enabledSetters["Oauth2ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2TradeshiftOauth2ClientSecret(v string) UpdateOAuth2TradeshiftOption {
	return func(o *UpdateOAuth2TradeshiftOptions) {
		o.Oauth2ClientSecret = v
		o.enabledSetters["Oauth2ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2TradeshiftEnabled(v bool) UpdateOAuth2TradeshiftOption {
	return func(o *UpdateOAuth2TradeshiftOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Tradeshift update the project OAuth2 Tradeshift configuration.
func (srv *Project) UpdateOAuth2Tradeshift(optionalSetters ...UpdateOAuth2TradeshiftOption)(*models.OAuth2Tradeshift, error) {
	path := "/project/oauth2/tradeshift"
	options := UpdateOAuth2TradeshiftOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Oauth2ClientId"] {
		params["oauth2ClientId"] = options.Oauth2ClientId
	}
	if options.enabledSetters["Oauth2ClientSecret"] {
		params["oauth2ClientSecret"] = options.Oauth2ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Tradeshift{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Tradeshift
	parsed, ok := resp.Result.(models.OAuth2Tradeshift)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2TradeshiftSandboxOptions struct {
	Oauth2ClientId string
	Oauth2ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2TradeshiftSandboxOptions) New() *UpdateOAuth2TradeshiftSandboxOptions {
	options.enabledSetters = map[string]bool{
		"Oauth2ClientId": false,
		"Oauth2ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2TradeshiftSandboxOption func(*UpdateOAuth2TradeshiftSandboxOptions)
func (srv *Project) WithUpdateOAuth2TradeshiftSandboxOauth2ClientId(v string) UpdateOAuth2TradeshiftSandboxOption {
	return func(o *UpdateOAuth2TradeshiftSandboxOptions) {
		o.Oauth2ClientId = v
		o.enabledSetters["Oauth2ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2TradeshiftSandboxOauth2ClientSecret(v string) UpdateOAuth2TradeshiftSandboxOption {
	return func(o *UpdateOAuth2TradeshiftSandboxOptions) {
		o.Oauth2ClientSecret = v
		o.enabledSetters["Oauth2ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2TradeshiftSandboxEnabled(v bool) UpdateOAuth2TradeshiftSandboxOption {
	return func(o *UpdateOAuth2TradeshiftSandboxOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2TradeshiftSandbox update the project OAuth2 Tradeshift Sandbox
// configuration.
func (srv *Project) UpdateOAuth2TradeshiftSandbox(optionalSetters ...UpdateOAuth2TradeshiftSandboxOption)(*models.OAuth2Tradeshift, error) {
	path := "/project/oauth2/tradeshiftBox"
	options := UpdateOAuth2TradeshiftSandboxOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Oauth2ClientId"] {
		params["oauth2ClientId"] = options.Oauth2ClientId
	}
	if options.enabledSetters["Oauth2ClientSecret"] {
		params["oauth2ClientSecret"] = options.Oauth2ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Tradeshift{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Tradeshift
	parsed, ok := resp.Result.(models.OAuth2Tradeshift)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2TwitchOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2TwitchOptions) New() *UpdateOAuth2TwitchOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2TwitchOption func(*UpdateOAuth2TwitchOptions)
func (srv *Project) WithUpdateOAuth2TwitchClientId(v string) UpdateOAuth2TwitchOption {
	return func(o *UpdateOAuth2TwitchOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2TwitchClientSecret(v string) UpdateOAuth2TwitchOption {
	return func(o *UpdateOAuth2TwitchOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2TwitchEnabled(v bool) UpdateOAuth2TwitchOption {
	return func(o *UpdateOAuth2TwitchOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Twitch update the project OAuth2 Twitch configuration.
func (srv *Project) UpdateOAuth2Twitch(optionalSetters ...UpdateOAuth2TwitchOption)(*models.OAuth2Twitch, error) {
	path := "/project/oauth2/twitch"
	options := UpdateOAuth2TwitchOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Twitch{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Twitch
	parsed, ok := resp.Result.(models.OAuth2Twitch)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2WordPressOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2WordPressOptions) New() *UpdateOAuth2WordPressOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2WordPressOption func(*UpdateOAuth2WordPressOptions)
func (srv *Project) WithUpdateOAuth2WordPressClientId(v string) UpdateOAuth2WordPressOption {
	return func(o *UpdateOAuth2WordPressOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2WordPressClientSecret(v string) UpdateOAuth2WordPressOption {
	return func(o *UpdateOAuth2WordPressOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2WordPressEnabled(v bool) UpdateOAuth2WordPressOption {
	return func(o *UpdateOAuth2WordPressOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2WordPress update the project OAuth2 WordPress configuration.
func (srv *Project) UpdateOAuth2WordPress(optionalSetters ...UpdateOAuth2WordPressOption)(*models.OAuth2WordPress, error) {
	path := "/project/oauth2/wordpress"
	options := UpdateOAuth2WordPressOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2WordPress{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2WordPress
	parsed, ok := resp.Result.(models.OAuth2WordPress)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2XOptions struct {
	CustomerKey string
	SecretKey string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2XOptions) New() *UpdateOAuth2XOptions {
	options.enabledSetters = map[string]bool{
		"CustomerKey": false,
		"SecretKey": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2XOption func(*UpdateOAuth2XOptions)
func (srv *Project) WithUpdateOAuth2XCustomerKey(v string) UpdateOAuth2XOption {
	return func(o *UpdateOAuth2XOptions) {
		o.CustomerKey = v
		o.enabledSetters["CustomerKey"] = true
	}
}
func (srv *Project) WithUpdateOAuth2XSecretKey(v string) UpdateOAuth2XOption {
	return func(o *UpdateOAuth2XOptions) {
		o.SecretKey = v
		o.enabledSetters["SecretKey"] = true
	}
}
func (srv *Project) WithUpdateOAuth2XEnabled(v bool) UpdateOAuth2XOption {
	return func(o *UpdateOAuth2XOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2X update the project OAuth2 X configuration.
func (srv *Project) UpdateOAuth2X(optionalSetters ...UpdateOAuth2XOption)(*models.OAuth2X, error) {
	path := "/project/oauth2/x"
	options := UpdateOAuth2XOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["CustomerKey"] {
		params["customerKey"] = options.CustomerKey
	}
	if options.enabledSetters["SecretKey"] {
		params["secretKey"] = options.SecretKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2X{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2X
	parsed, ok := resp.Result.(models.OAuth2X)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2YahooOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2YahooOptions) New() *UpdateOAuth2YahooOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2YahooOption func(*UpdateOAuth2YahooOptions)
func (srv *Project) WithUpdateOAuth2YahooClientId(v string) UpdateOAuth2YahooOption {
	return func(o *UpdateOAuth2YahooOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2YahooClientSecret(v string) UpdateOAuth2YahooOption {
	return func(o *UpdateOAuth2YahooOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2YahooEnabled(v bool) UpdateOAuth2YahooOption {
	return func(o *UpdateOAuth2YahooOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Yahoo update the project OAuth2 Yahoo configuration.
func (srv *Project) UpdateOAuth2Yahoo(optionalSetters ...UpdateOAuth2YahooOption)(*models.OAuth2Yahoo, error) {
	path := "/project/oauth2/yahoo"
	options := UpdateOAuth2YahooOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Yahoo{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Yahoo
	parsed, ok := resp.Result.(models.OAuth2Yahoo)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2YandexOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2YandexOptions) New() *UpdateOAuth2YandexOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2YandexOption func(*UpdateOAuth2YandexOptions)
func (srv *Project) WithUpdateOAuth2YandexClientId(v string) UpdateOAuth2YandexOption {
	return func(o *UpdateOAuth2YandexOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2YandexClientSecret(v string) UpdateOAuth2YandexOption {
	return func(o *UpdateOAuth2YandexOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2YandexEnabled(v bool) UpdateOAuth2YandexOption {
	return func(o *UpdateOAuth2YandexOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Yandex update the project OAuth2 Yandex configuration.
func (srv *Project) UpdateOAuth2Yandex(optionalSetters ...UpdateOAuth2YandexOption)(*models.OAuth2Yandex, error) {
	path := "/project/oauth2/yandex"
	options := UpdateOAuth2YandexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Yandex{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Yandex
	parsed, ok := resp.Result.(models.OAuth2Yandex)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2ZohoOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2ZohoOptions) New() *UpdateOAuth2ZohoOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2ZohoOption func(*UpdateOAuth2ZohoOptions)
func (srv *Project) WithUpdateOAuth2ZohoClientId(v string) UpdateOAuth2ZohoOption {
	return func(o *UpdateOAuth2ZohoOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2ZohoClientSecret(v string) UpdateOAuth2ZohoOption {
	return func(o *UpdateOAuth2ZohoOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2ZohoEnabled(v bool) UpdateOAuth2ZohoOption {
	return func(o *UpdateOAuth2ZohoOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Zoho update the project OAuth2 Zoho configuration.
func (srv *Project) UpdateOAuth2Zoho(optionalSetters ...UpdateOAuth2ZohoOption)(*models.OAuth2Zoho, error) {
	path := "/project/oauth2/zoho"
	options := UpdateOAuth2ZohoOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Zoho{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Zoho
	parsed, ok := resp.Result.(models.OAuth2Zoho)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOAuth2ZoomOptions struct {
	ClientId string
	ClientSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateOAuth2ZoomOptions) New() *UpdateOAuth2ZoomOptions {
	options.enabledSetters = map[string]bool{
		"ClientId": false,
		"ClientSecret": false,
		"Enabled": false,
	}
	return &options
}
type UpdateOAuth2ZoomOption func(*UpdateOAuth2ZoomOptions)
func (srv *Project) WithUpdateOAuth2ZoomClientId(v string) UpdateOAuth2ZoomOption {
	return func(o *UpdateOAuth2ZoomOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Project) WithUpdateOAuth2ZoomClientSecret(v string) UpdateOAuth2ZoomOption {
	return func(o *UpdateOAuth2ZoomOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Project) WithUpdateOAuth2ZoomEnabled(v bool) UpdateOAuth2ZoomOption {
	return func(o *UpdateOAuth2ZoomOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateOAuth2Zoom update the project OAuth2 Zoom configuration.
func (srv *Project) UpdateOAuth2Zoom(optionalSetters ...UpdateOAuth2ZoomOption)(*models.OAuth2Zoom, error) {
	path := "/project/oauth2/zoom"
	options := UpdateOAuth2ZoomOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["clientId"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["clientSecret"] = options.ClientSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.OAuth2Zoom{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OAuth2Zoom
	parsed, ok := resp.Result.(models.OAuth2Zoom)
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
type ListPoliciesOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListPoliciesOptions) New() *ListPoliciesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListPoliciesOption func(*ListPoliciesOptions)
func (srv *Project) WithListPoliciesQueries(v []string) ListPoliciesOption {
	return func(o *ListPoliciesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Project) WithListPoliciesTotal(v bool) ListPoliciesOption {
	return func(o *ListPoliciesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListPolicies get a list of all project policies and their current
// configuration.
func (srv *Project) ListPolicies(optionalSetters ...ListPoliciesOption)(*models.PolicyList, error) {
	path := "/project/policies"
	options := ListPoliciesOptions{}.New()
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

		parsed := models.PolicyList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PolicyList
	parsed, ok := resp.Result.(models.PolicyList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateMembershipPrivacyPolicyOptions struct {
	UserId bool
	UserEmail bool
	UserPhone bool
	UserName bool
	UserMFA bool
	enabledSetters map[string]bool
}
func (options UpdateMembershipPrivacyPolicyOptions) New() *UpdateMembershipPrivacyPolicyOptions {
	options.enabledSetters = map[string]bool{
		"UserId": false,
		"UserEmail": false,
		"UserPhone": false,
		"UserName": false,
		"UserMFA": false,
	}
	return &options
}
type UpdateMembershipPrivacyPolicyOption func(*UpdateMembershipPrivacyPolicyOptions)
func (srv *Project) WithUpdateMembershipPrivacyPolicyUserId(v bool) UpdateMembershipPrivacyPolicyOption {
	return func(o *UpdateMembershipPrivacyPolicyOptions) {
		o.UserId = v
		o.enabledSetters["UserId"] = true
	}
}
func (srv *Project) WithUpdateMembershipPrivacyPolicyUserEmail(v bool) UpdateMembershipPrivacyPolicyOption {
	return func(o *UpdateMembershipPrivacyPolicyOptions) {
		o.UserEmail = v
		o.enabledSetters["UserEmail"] = true
	}
}
func (srv *Project) WithUpdateMembershipPrivacyPolicyUserPhone(v bool) UpdateMembershipPrivacyPolicyOption {
	return func(o *UpdateMembershipPrivacyPolicyOptions) {
		o.UserPhone = v
		o.enabledSetters["UserPhone"] = true
	}
}
func (srv *Project) WithUpdateMembershipPrivacyPolicyUserName(v bool) UpdateMembershipPrivacyPolicyOption {
	return func(o *UpdateMembershipPrivacyPolicyOptions) {
		o.UserName = v
		o.enabledSetters["UserName"] = true
	}
}
func (srv *Project) WithUpdateMembershipPrivacyPolicyUserMFA(v bool) UpdateMembershipPrivacyPolicyOption {
	return func(o *UpdateMembershipPrivacyPolicyOptions) {
		o.UserMFA = v
		o.enabledSetters["UserMFA"] = true
	}
}
	
// UpdateMembershipPrivacyPolicy updating this policy allows you to control if
// team members can see other members information. When enabled, all team
// members can see ID, name, email, phone number, and MFA status of other
// members..
func (srv *Project) UpdateMembershipPrivacyPolicy(optionalSetters ...UpdateMembershipPrivacyPolicyOption)(*models.Project, error) {
	path := "/project/policies/membership-privacy"
	options := UpdateMembershipPrivacyPolicyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["UserId"] {
		params["userId"] = options.UserId
	}
	if options.enabledSetters["UserEmail"] {
		params["userEmail"] = options.UserEmail
	}
	if options.enabledSetters["UserPhone"] {
		params["userPhone"] = options.UserPhone
	}
	if options.enabledSetters["UserName"] {
		params["userName"] = options.UserName
	}
	if options.enabledSetters["UserMFA"] {
		params["userMFA"] = options.UserMFA
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
	
// UpdatePasswordDictionaryPolicy updating this policy allows you to control
// if new passwords are checked against most common passwords dictionary. When
// enabled, and user changes their password, password must not be contained in
// the dictionary.
func (srv *Project) UpdatePasswordDictionaryPolicy(Enabled bool)(*models.Project, error) {
	path := "/project/policies/password-dictionary"
	params := map[string]interface{}{}
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
	
// UpdatePasswordHistoryPolicy updates one of password strength policies.
// Based on total length configured, previous password hashes are stored, and
// users cannot choose a new password that is already stored in the passwird
// history list, when updating an user password, or setting new one through
// password recovery.
// 
// Keep in mind, while password history policy is disabled, the history is not
// being stored. Enabling the policy will not have any history on existing
// users, and it will only start to collect and enforce the policy on password
// changes since the policy is enabled.
func (srv *Project) UpdatePasswordHistoryPolicy(Total int)(*models.Project, error) {
	path := "/project/policies/password-history"
	params := map[string]interface{}{}
	params["total"] = Total
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
	
// UpdatePasswordPersonalDataPolicy updating this policy allows you to control
// if password strength is checked against personal data. When enabled, and
// user sets or changes their password, the password must not contain user ID,
// name, email or phone number.
func (srv *Project) UpdatePasswordPersonalDataPolicy(Enabled bool)(*models.Project, error) {
	path := "/project/policies/password-personal-data"
	params := map[string]interface{}{}
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
	
// UpdateSessionAlertPolicy updating this policy allows you to control if
// email alert is sent upon session creation. When enabled, and user signs
// into their account, they will be sent an email notification. There is an
// exception, the first session after a new sign up does not trigger an alert,
// even if the policy is enabled.
func (srv *Project) UpdateSessionAlertPolicy(Enabled bool)(*models.Project, error) {
	path := "/project/policies/session-alert"
	params := map[string]interface{}{}
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
	
// UpdateSessionDurationPolicy update maximum duration how long sessions
// created within a project should stay active for.
func (srv *Project) UpdateSessionDurationPolicy(Duration int)(*models.Project, error) {
	path := "/project/policies/session-duration"
	params := map[string]interface{}{}
	params["duration"] = Duration
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
	
// UpdateSessionInvalidationPolicy updating this policy allows you to control
// if existing sessions should be invalidated when a password of a user is
// changed. When enabled, and user changes their password, they will be logged
// out of all their devices.
func (srv *Project) UpdateSessionInvalidationPolicy(Enabled bool)(*models.Project, error) {
	path := "/project/policies/session-invalidation"
	params := map[string]interface{}{}
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
	
// UpdateSessionLimitPolicy update the maximum number of sessions allowed per
// user. When the limit is hit, the oldest session will be deleted to make
// room for new one.
func (srv *Project) UpdateSessionLimitPolicy(Total int)(*models.Project, error) {
	path := "/project/policies/session-limit"
	params := map[string]interface{}{}
	params["total"] = Total
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
	
// UpdateUserLimitPolicy update the maximum number of users in the project.
// When the limit is hit or amount of existing users already exceeded the
// limit, all users remain active, but new user sign up will be prohibited.
func (srv *Project) UpdateUserLimitPolicy(Total int)(*models.Project, error) {
	path := "/project/policies/user-limit"
	params := map[string]interface{}{}
	params["total"] = Total
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
	
// GetPolicy get a policy by its unique ID. This endpoint returns the current
// configuration for the requested project policy.
func (srv *Project) GetPolicy(PolicyId string)(models.Model, error) {
	r := strings.NewReplacer("{policyId}", PolicyId)
	path := r.Replace("/project/policies/{policyId}")
	params := map[string]interface{}{}
	params["policyId"] = PolicyId
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
		if fmt.Sprint(response["$id"]) == "password-dictionary" {
			parsed := models.PolicyPasswordDictionary{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "password-history" {
			parsed := models.PolicyPasswordHistory{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "password-personal-data" {
			parsed := models.PolicyPasswordPersonalData{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "session-alert" {
			parsed := models.PolicySessionAlert{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "session-duration" {
			parsed := models.PolicySessionDuration{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "session-invalidation" {
			parsed := models.PolicySessionInvalidation{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "session-limit" {
			parsed := models.PolicySessionLimit{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "user-limit" {
			parsed := models.PolicyUserLimit{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["$id"]) == "membership-privacy" {
			parsed := models.PolicyMembershipPrivacy{}.New(bytes)
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
			
// UpdateProtocol update properties of a specific protocol. Use this endpoint
// to enable or disable a protocol in your project.
func (srv *Project) UpdateProtocol(ProtocolId string, Enabled bool)(*models.Project, error) {
	r := strings.NewReplacer("{protocolId}", ProtocolId)
	path := r.Replace("/project/protocols/{protocolId}")
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
			
// UpdateService update properties of a specific service. Use this endpoint to
// enable or disable a service in your project.
func (srv *Project) UpdateService(ServiceId string, Enabled bool)(*models.Project, error) {
	r := strings.NewReplacer("{serviceId}", ServiceId)
	path := r.Replace("/project/services/{serviceId}")
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
type UpdateSMTPOptions struct {
	Host string
	Port int
	Username string
	Password string
	SenderEmail string
	SenderName string
	ReplyToEmail string
	ReplyToName string
	Secure string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateSMTPOptions) New() *UpdateSMTPOptions {
	options.enabledSetters = map[string]bool{
		"Host": false,
		"Port": false,
		"Username": false,
		"Password": false,
		"SenderEmail": false,
		"SenderName": false,
		"ReplyToEmail": false,
		"ReplyToName": false,
		"Secure": false,
		"Enabled": false,
	}
	return &options
}
type UpdateSMTPOption func(*UpdateSMTPOptions)
func (srv *Project) WithUpdateSMTPHost(v string) UpdateSMTPOption {
	return func(o *UpdateSMTPOptions) {
		o.Host = v
		o.enabledSetters["Host"] = true
	}
}
func (srv *Project) WithUpdateSMTPPort(v int) UpdateSMTPOption {
	return func(o *UpdateSMTPOptions) {
		o.Port = v
		o.enabledSetters["Port"] = true
	}
}
func (srv *Project) WithUpdateSMTPUsername(v string) UpdateSMTPOption {
	return func(o *UpdateSMTPOptions) {
		o.Username = v
		o.enabledSetters["Username"] = true
	}
}
func (srv *Project) WithUpdateSMTPPassword(v string) UpdateSMTPOption {
	return func(o *UpdateSMTPOptions) {
		o.Password = v
		o.enabledSetters["Password"] = true
	}
}
func (srv *Project) WithUpdateSMTPSenderEmail(v string) UpdateSMTPOption {
	return func(o *UpdateSMTPOptions) {
		o.SenderEmail = v
		o.enabledSetters["SenderEmail"] = true
	}
}
func (srv *Project) WithUpdateSMTPSenderName(v string) UpdateSMTPOption {
	return func(o *UpdateSMTPOptions) {
		o.SenderName = v
		o.enabledSetters["SenderName"] = true
	}
}
func (srv *Project) WithUpdateSMTPReplyToEmail(v string) UpdateSMTPOption {
	return func(o *UpdateSMTPOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
func (srv *Project) WithUpdateSMTPReplyToName(v string) UpdateSMTPOption {
	return func(o *UpdateSMTPOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
func (srv *Project) WithUpdateSMTPSecure(v string) UpdateSMTPOption {
	return func(o *UpdateSMTPOptions) {
		o.Secure = v
		o.enabledSetters["Secure"] = true
	}
}
func (srv *Project) WithUpdateSMTPEnabled(v bool) UpdateSMTPOption {
	return func(o *UpdateSMTPOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
	
// UpdateSMTP update the SMTP configuration for your project. Use this
// endpoint to configure your project's SMTP provider with your custom
// settings for sending transactional emails.
func (srv *Project) UpdateSMTP(optionalSetters ...UpdateSMTPOption)(*models.Project, error) {
	path := "/project/smtp"
	options := UpdateSMTPOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Host"] {
		params["host"] = options.Host
	}
	if options.enabledSetters["Port"] {
		params["port"] = options.Port
	}
	if options.enabledSetters["Username"] {
		params["username"] = options.Username
	}
	if options.enabledSetters["Password"] {
		params["password"] = options.Password
	}
	if options.enabledSetters["SenderEmail"] {
		params["senderEmail"] = options.SenderEmail
	}
	if options.enabledSetters["SenderName"] {
		params["senderName"] = options.SenderName
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
	}
	if options.enabledSetters["Secure"] {
		params["secure"] = options.Secure
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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
	
// CreateSMTPTest send a test email to verify SMTP configuration.
func (srv *Project) CreateSMTPTest(Emails []string)(*interface{}, error) {
	path := "/project/smtp/tests"
	params := map[string]interface{}{}
	params["emails"] = Emails
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
type ListEmailTemplatesOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListEmailTemplatesOptions) New() *ListEmailTemplatesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListEmailTemplatesOption func(*ListEmailTemplatesOptions)
func (srv *Project) WithListEmailTemplatesQueries(v []string) ListEmailTemplatesOption {
	return func(o *ListEmailTemplatesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Project) WithListEmailTemplatesTotal(v bool) ListEmailTemplatesOption {
	return func(o *ListEmailTemplatesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListEmailTemplates get a list of all custom email templates configured for
// the project. This endpoint returns an array of all configured email
// templates and their locales.
func (srv *Project) ListEmailTemplates(optionalSetters ...ListEmailTemplatesOption)(*models.EmailTemplateList, error) {
	path := "/project/templates/email"
	options := ListEmailTemplatesOptions{}.New()
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

		parsed := models.EmailTemplateList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.EmailTemplateList
	parsed, ok := resp.Result.(models.EmailTemplateList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateEmailTemplateOptions struct {
	Locale string
	Subject string
	Message string
	SenderName string
	SenderEmail string
	ReplyToEmail string
	ReplyToName string
	enabledSetters map[string]bool
}
func (options UpdateEmailTemplateOptions) New() *UpdateEmailTemplateOptions {
	options.enabledSetters = map[string]bool{
		"Locale": false,
		"Subject": false,
		"Message": false,
		"SenderName": false,
		"SenderEmail": false,
		"ReplyToEmail": false,
		"ReplyToName": false,
	}
	return &options
}
type UpdateEmailTemplateOption func(*UpdateEmailTemplateOptions)
func (srv *Project) WithUpdateEmailTemplateLocale(v string) UpdateEmailTemplateOption {
	return func(o *UpdateEmailTemplateOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *Project) WithUpdateEmailTemplateSubject(v string) UpdateEmailTemplateOption {
	return func(o *UpdateEmailTemplateOptions) {
		o.Subject = v
		o.enabledSetters["Subject"] = true
	}
}
func (srv *Project) WithUpdateEmailTemplateMessage(v string) UpdateEmailTemplateOption {
	return func(o *UpdateEmailTemplateOptions) {
		o.Message = v
		o.enabledSetters["Message"] = true
	}
}
func (srv *Project) WithUpdateEmailTemplateSenderName(v string) UpdateEmailTemplateOption {
	return func(o *UpdateEmailTemplateOptions) {
		o.SenderName = v
		o.enabledSetters["SenderName"] = true
	}
}
func (srv *Project) WithUpdateEmailTemplateSenderEmail(v string) UpdateEmailTemplateOption {
	return func(o *UpdateEmailTemplateOptions) {
		o.SenderEmail = v
		o.enabledSetters["SenderEmail"] = true
	}
}
func (srv *Project) WithUpdateEmailTemplateReplyToEmail(v string) UpdateEmailTemplateOption {
	return func(o *UpdateEmailTemplateOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
func (srv *Project) WithUpdateEmailTemplateReplyToName(v string) UpdateEmailTemplateOption {
	return func(o *UpdateEmailTemplateOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
			
// UpdateEmailTemplate update a custom email template for the specified locale
// and type. Use this endpoint to modify the content of your email templates.
func (srv *Project) UpdateEmailTemplate(TemplateId string, optionalSetters ...UpdateEmailTemplateOption)(*models.EmailTemplate, error) {
	path := "/project/templates/email"
	options := UpdateEmailTemplateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["templateId"] = TemplateId
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["Subject"] {
		params["subject"] = options.Subject
	}
	if options.enabledSetters["Message"] {
		params["message"] = options.Message
	}
	if options.enabledSetters["SenderName"] {
		params["senderName"] = options.SenderName
	}
	if options.enabledSetters["SenderEmail"] {
		params["senderEmail"] = options.SenderEmail
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
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

		parsed := models.EmailTemplate{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.EmailTemplate
	parsed, ok := resp.Result.(models.EmailTemplate)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetEmailTemplateOptions struct {
	Locale string
	enabledSetters map[string]bool
}
func (options GetEmailTemplateOptions) New() *GetEmailTemplateOptions {
	options.enabledSetters = map[string]bool{
		"Locale": false,
	}
	return &options
}
type GetEmailTemplateOption func(*GetEmailTemplateOptions)
func (srv *Project) WithGetEmailTemplateLocale(v string) GetEmailTemplateOption {
	return func(o *GetEmailTemplateOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
			
// GetEmailTemplate get a custom email template for the specified locale and
// type. This endpoint returns the template content, subject, and other
// configuration details.
func (srv *Project) GetEmailTemplate(TemplateId string, optionalSetters ...GetEmailTemplateOption)(*models.EmailTemplate, error) {
	r := strings.NewReplacer("{templateId}", TemplateId)
	path := r.Replace("/project/templates/email/{templateId}")
	options := GetEmailTemplateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["templateId"] = TemplateId
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.EmailTemplate{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.EmailTemplate
	parsed, ok := resp.Result.(models.EmailTemplate)
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
