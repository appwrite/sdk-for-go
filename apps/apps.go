package apps

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/v6/client"
	"github.com/appwrite/sdk-for-go/v6/models"
	"strings"
)

// Apps service
type Apps struct {
	client client.Client
}

func New(clt client.Client) *Apps {
	return &Apps{
		client: clt,
	}
}

type ListOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListOption func(*ListOptions)
func (srv *Apps) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Apps) WithListTotal(v bool) ListOption {
	return func(o *ListOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// List list applications.
func (srv *Apps) List(optionalSetters ...ListOption)(*models.AppsList, error) {
	path := "/apps"
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
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AppsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AppsList
	parsed, ok := resp.Result.(models.AppsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateOptions struct {
	Description string
	ClientUri string
	LogoUri string
	PrivacyPolicyUrl string
	TermsUrl string
	Contacts []string
	Tagline string
	Tags []string
	Images []string
	SupportUrl string
	DataDeletionUrl string
	PostLogoutRedirectUris []string
	Enabled bool
	Type string
	DeviceFlow bool
	TeamId string
	enabledSetters map[string]bool
}
func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"ClientUri": false,
		"LogoUri": false,
		"PrivacyPolicyUrl": false,
		"TermsUrl": false,
		"Contacts": false,
		"Tagline": false,
		"Tags": false,
		"Images": false,
		"SupportUrl": false,
		"DataDeletionUrl": false,
		"PostLogoutRedirectUris": false,
		"Enabled": false,
		"Type": false,
		"DeviceFlow": false,
		"TeamId": false,
	}
	return &options
}
type CreateOption func(*CreateOptions)
func (srv *Apps) WithCreateDescription(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Apps) WithCreateClientUri(v string) CreateOption {
	return func(o *CreateOptions) {
		o.ClientUri = v
		o.enabledSetters["ClientUri"] = true
	}
}
func (srv *Apps) WithCreateLogoUri(v string) CreateOption {
	return func(o *CreateOptions) {
		o.LogoUri = v
		o.enabledSetters["LogoUri"] = true
	}
}
func (srv *Apps) WithCreatePrivacyPolicyUrl(v string) CreateOption {
	return func(o *CreateOptions) {
		o.PrivacyPolicyUrl = v
		o.enabledSetters["PrivacyPolicyUrl"] = true
	}
}
func (srv *Apps) WithCreateTermsUrl(v string) CreateOption {
	return func(o *CreateOptions) {
		o.TermsUrl = v
		o.enabledSetters["TermsUrl"] = true
	}
}
func (srv *Apps) WithCreateContacts(v []string) CreateOption {
	return func(o *CreateOptions) {
		o.Contacts = v
		o.enabledSetters["Contacts"] = true
	}
}
func (srv *Apps) WithCreateTagline(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Tagline = v
		o.enabledSetters["Tagline"] = true
	}
}
func (srv *Apps) WithCreateTags(v []string) CreateOption {
	return func(o *CreateOptions) {
		o.Tags = v
		o.enabledSetters["Tags"] = true
	}
}
func (srv *Apps) WithCreateImages(v []string) CreateOption {
	return func(o *CreateOptions) {
		o.Images = v
		o.enabledSetters["Images"] = true
	}
}
func (srv *Apps) WithCreateSupportUrl(v string) CreateOption {
	return func(o *CreateOptions) {
		o.SupportUrl = v
		o.enabledSetters["SupportUrl"] = true
	}
}
func (srv *Apps) WithCreateDataDeletionUrl(v string) CreateOption {
	return func(o *CreateOptions) {
		o.DataDeletionUrl = v
		o.enabledSetters["DataDeletionUrl"] = true
	}
}
func (srv *Apps) WithCreatePostLogoutRedirectUris(v []string) CreateOption {
	return func(o *CreateOptions) {
		o.PostLogoutRedirectUris = v
		o.enabledSetters["PostLogoutRedirectUris"] = true
	}
}
func (srv *Apps) WithCreateEnabled(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Apps) WithCreateType(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *Apps) WithCreateDeviceFlow(v bool) CreateOption {
	return func(o *CreateOptions) {
		o.DeviceFlow = v
		o.enabledSetters["DeviceFlow"] = true
	}
}
func (srv *Apps) WithCreateTeamId(v string) CreateOption {
	return func(o *CreateOptions) {
		o.TeamId = v
		o.enabledSetters["TeamId"] = true
	}
}
							
// Create create a new application.
func (srv *Apps) Create(AppId string, Name string, RedirectUris []string, optionalSetters ...CreateOption)(*models.App, error) {
	path := "/apps"
	options := CreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["appId"] = AppId
	params["name"] = Name
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["ClientUri"] {
		params["clientUri"] = options.ClientUri
	}
	if options.enabledSetters["LogoUri"] {
		params["logoUri"] = options.LogoUri
	}
	if options.enabledSetters["PrivacyPolicyUrl"] {
		params["privacyPolicyUrl"] = options.PrivacyPolicyUrl
	}
	if options.enabledSetters["TermsUrl"] {
		params["termsUrl"] = options.TermsUrl
	}
	if options.enabledSetters["Contacts"] {
		params["contacts"] = options.Contacts
	}
	if options.enabledSetters["Tagline"] {
		params["tagline"] = options.Tagline
	}
	if options.enabledSetters["Tags"] {
		params["tags"] = options.Tags
	}
	if options.enabledSetters["Images"] {
		params["images"] = options.Images
	}
	if options.enabledSetters["SupportUrl"] {
		params["supportUrl"] = options.SupportUrl
	}
	if options.enabledSetters["DataDeletionUrl"] {
		params["dataDeletionUrl"] = options.DataDeletionUrl
	}
	params["redirectUris"] = RedirectUris
	if options.enabledSetters["PostLogoutRedirectUris"] {
		params["postLogoutRedirectUris"] = options.PostLogoutRedirectUris
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["DeviceFlow"] {
		params["deviceFlow"] = options.DeviceFlow
	}
	if options.enabledSetters["TeamId"] {
		params["teamId"] = options.TeamId
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.App{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.App
	parsed, ok := resp.Result.(models.App)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListInstallationScopes list scopes an application can request when
// installed on a team.
func (srv *Apps) ListInstallationScopes()(*models.AppScopeList, error) {
	path := "/apps/scopes/installations"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AppScopeList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AppScopeList
	parsed, ok := resp.Result.(models.AppScopeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListOAuth2Scopes list scopes an application can request during the OAuth2
// flow.
func (srv *Apps) ListOAuth2Scopes()(*models.AppScopeList, error) {
	path := "/apps/scopes/oauth2"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AppScopeList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AppScopeList
	parsed, ok := resp.Result.(models.AppScopeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Get get an application by its unique ID.
func (srv *Apps) Get(AppId string)(*models.App, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId))
	path := r.Replace("/apps/{appId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.App{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.App
	parsed, ok := resp.Result.(models.App)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOptions struct {
	Description string
	ClientUri string
	LogoUri string
	PrivacyPolicyUrl string
	TermsUrl string
	Contacts []string
	Tagline string
	Tags []string
	Images []string
	SupportUrl string
	DataDeletionUrl string
	Enabled bool
	RedirectUris []string
	PostLogoutRedirectUris []string
	Type string
	DeviceFlow bool
	InstallationScopes []string
	InstallationRedirectUrl string
	enabledSetters map[string]bool
}
func (options UpdateOptions) New() *UpdateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"ClientUri": false,
		"LogoUri": false,
		"PrivacyPolicyUrl": false,
		"TermsUrl": false,
		"Contacts": false,
		"Tagline": false,
		"Tags": false,
		"Images": false,
		"SupportUrl": false,
		"DataDeletionUrl": false,
		"Enabled": false,
		"RedirectUris": false,
		"PostLogoutRedirectUris": false,
		"Type": false,
		"DeviceFlow": false,
		"InstallationScopes": false,
		"InstallationRedirectUrl": false,
	}
	return &options
}
type UpdateOption func(*UpdateOptions)
func (srv *Apps) WithUpdateDescription(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Apps) WithUpdateClientUri(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.ClientUri = v
		o.enabledSetters["ClientUri"] = true
	}
}
func (srv *Apps) WithUpdateLogoUri(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.LogoUri = v
		o.enabledSetters["LogoUri"] = true
	}
}
func (srv *Apps) WithUpdatePrivacyPolicyUrl(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.PrivacyPolicyUrl = v
		o.enabledSetters["PrivacyPolicyUrl"] = true
	}
}
func (srv *Apps) WithUpdateTermsUrl(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.TermsUrl = v
		o.enabledSetters["TermsUrl"] = true
	}
}
func (srv *Apps) WithUpdateContacts(v []string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Contacts = v
		o.enabledSetters["Contacts"] = true
	}
}
func (srv *Apps) WithUpdateTagline(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Tagline = v
		o.enabledSetters["Tagline"] = true
	}
}
func (srv *Apps) WithUpdateTags(v []string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Tags = v
		o.enabledSetters["Tags"] = true
	}
}
func (srv *Apps) WithUpdateImages(v []string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Images = v
		o.enabledSetters["Images"] = true
	}
}
func (srv *Apps) WithUpdateSupportUrl(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.SupportUrl = v
		o.enabledSetters["SupportUrl"] = true
	}
}
func (srv *Apps) WithUpdateDataDeletionUrl(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.DataDeletionUrl = v
		o.enabledSetters["DataDeletionUrl"] = true
	}
}
func (srv *Apps) WithUpdateEnabled(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Apps) WithUpdateRedirectUris(v []string) UpdateOption {
	return func(o *UpdateOptions) {
		o.RedirectUris = v
		o.enabledSetters["RedirectUris"] = true
	}
}
func (srv *Apps) WithUpdatePostLogoutRedirectUris(v []string) UpdateOption {
	return func(o *UpdateOptions) {
		o.PostLogoutRedirectUris = v
		o.enabledSetters["PostLogoutRedirectUris"] = true
	}
}
func (srv *Apps) WithUpdateType(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *Apps) WithUpdateDeviceFlow(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.DeviceFlow = v
		o.enabledSetters["DeviceFlow"] = true
	}
}
func (srv *Apps) WithUpdateInstallationScopes(v []string) UpdateOption {
	return func(o *UpdateOptions) {
		o.InstallationScopes = v
		o.enabledSetters["InstallationScopes"] = true
	}
}
func (srv *Apps) WithUpdateInstallationRedirectUrl(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.InstallationRedirectUrl = v
		o.enabledSetters["InstallationRedirectUrl"] = true
	}
}
					
// Update update an application by its unique ID.
func (srv *Apps) Update(AppId string, Name string, optionalSetters ...UpdateOption)(*models.App, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId))
	path := r.Replace("/apps/{appId}")
	options := UpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["ClientUri"] {
		params["clientUri"] = options.ClientUri
	}
	if options.enabledSetters["LogoUri"] {
		params["logoUri"] = options.LogoUri
	}
	if options.enabledSetters["PrivacyPolicyUrl"] {
		params["privacyPolicyUrl"] = options.PrivacyPolicyUrl
	}
	if options.enabledSetters["TermsUrl"] {
		params["termsUrl"] = options.TermsUrl
	}
	if options.enabledSetters["Contacts"] {
		params["contacts"] = options.Contacts
	}
	if options.enabledSetters["Tagline"] {
		params["tagline"] = options.Tagline
	}
	if options.enabledSetters["Tags"] {
		params["tags"] = options.Tags
	}
	if options.enabledSetters["Images"] {
		params["images"] = options.Images
	}
	if options.enabledSetters["SupportUrl"] {
		params["supportUrl"] = options.SupportUrl
	}
	if options.enabledSetters["DataDeletionUrl"] {
		params["dataDeletionUrl"] = options.DataDeletionUrl
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["RedirectUris"] {
		params["redirectUris"] = options.RedirectUris
	}
	if options.enabledSetters["PostLogoutRedirectUris"] {
		params["postLogoutRedirectUris"] = options.PostLogoutRedirectUris
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["DeviceFlow"] {
		params["deviceFlow"] = options.DeviceFlow
	}
	if options.enabledSetters["InstallationScopes"] {
		params["installationScopes"] = options.InstallationScopes
	}
	if options.enabledSetters["InstallationRedirectUrl"] {
		params["installationRedirectUrl"] = options.InstallationRedirectUrl
	}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.App{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.App
	parsed, ok := resp.Result.(models.App)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Delete delete an application by its unique ID.
func (srv *Apps) Delete(AppId string)(*interface{}, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId))
	path := r.Replace("/apps/{appId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

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
type ListInstallationsOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListInstallationsOptions) New() *ListInstallationsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListInstallationsOption func(*ListInstallationsOptions)
func (srv *Apps) WithListInstallationsQueries(v []string) ListInstallationsOption {
	return func(o *ListInstallationsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Apps) WithListInstallationsTotal(v bool) ListInstallationsOption {
	return func(o *ListInstallationsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// ListInstallations list installations of an application. Requires an app key
// sent in the `X-Appwrite-Key` header alongside the `X-Appwrite-App` header,
// or a caller with update access to the app.
func (srv *Apps) ListInstallations(AppId string, optionalSetters ...ListInstallationsOption)(*models.AppInstallationList, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId))
	path := r.Replace("/apps/{appId}/installations")
	options := ListInstallationsOptions{}.New()
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
		"accept": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AppInstallationList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AppInstallationList
	parsed, ok := resp.Result.(models.AppInstallationList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetInstallation get an installation of an application by its unique ID.
// Requires an app key sent in the `X-Appwrite-Key` header alongside the
// `X-Appwrite-App` header, or a caller with update access to the app.
func (srv *Apps) GetInstallation(AppId string, InstallationId string)(*models.AppInstallation, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId), "{installationId}", client.EncodePath(InstallationId))
	path := r.Replace("/apps/{appId}/installations/{installationId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AppInstallation{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AppInstallation
	parsed, ok := resp.Result.(models.AppInstallation)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// DeleteInstallation delete an installation of an application by its unique
// ID. Requires a caller with update access to the app. Previously issued
// installation access tokens are revoked.
func (srv *Apps) DeleteInstallation(AppId string, InstallationId string)(*interface{}, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId), "{installationId}", client.EncodePath(InstallationId))
	path := r.Replace("/apps/{appId}/installations/{installationId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

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
			
// CreateInstallationToken create a token for an installation of an
// application. Requires an app key sent in the `X-Appwrite-Key` header
// alongside the `X-Appwrite-App` header, or a caller with update access to
// the app. The returned token carries the scopes and authorization details
// granted to the installation, and can be used as an `Authorization: Bearer`
// header everywhere OAuth2 access tokens are accepted. Multiple tokens can be
// active for the same installation at once; each token stays valid until it
// expires or the installation is updated or deleted.
func (srv *Apps) CreateInstallationToken(AppId string, InstallationId string)(*models.Oauth2Token, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId), "{installationId}", client.EncodePath(InstallationId))
	path := r.Replace("/apps/{appId}/installations/{installationId}/tokens")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.Oauth2Token{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Oauth2Token
	parsed, ok := resp.Result.(models.Oauth2Token)
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
func (srv *Apps) WithListKeysQueries(v []string) ListKeysOption {
	return func(o *ListKeysOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Apps) WithListKeysTotal(v bool) ListKeysOption {
	return func(o *ListKeysOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// ListKeys list app keys for an application.
func (srv *Apps) ListKeys(AppId string, optionalSetters ...ListKeysOption)(*models.AppKeyList, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId))
	path := r.Replace("/apps/{appId}/keys")
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
		"accept": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AppKeyList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AppKeyList
	parsed, ok := resp.Result.(models.AppKeyList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CreateKey create a new app key for an application. App keys carry no
// scopes; send one in the `X-Appwrite-Key` header alongside the
// `X-Appwrite-App` header to list the application's installations and create
// installation access tokens.
func (srv *Apps) CreateKey(AppId string)(*models.AppKey, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId))
	path := r.Replace("/apps/{appId}/keys")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AppKey{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AppKey
	parsed, ok := resp.Result.(models.AppKey)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetKey get an app key by its unique ID.
func (srv *Apps) GetKey(AppId string, KeyId string)(*models.AppKey, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId), "{keyId}", client.EncodePath(KeyId))
	path := r.Replace("/apps/{appId}/keys/{keyId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AppKey{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AppKey
	parsed, ok := resp.Result.(models.AppKey)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// DeleteKey delete an app key by its unique ID.
func (srv *Apps) DeleteKey(AppId string, KeyId string)(*interface{}, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId), "{keyId}", client.EncodePath(KeyId))
	path := r.Replace("/apps/{appId}/keys/{keyId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

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
			
// UpdateLabels update the labels of an application. Labels are read-only for
// clients; only a server SDK using a project API key can set them. Replaces
// the previous labels.
func (srv *Apps) UpdateLabels(AppId string, Labels []string)(*models.App, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId))
	path := r.Replace("/apps/{appId}/labels")
	params := map[string]interface{}{}
	params["labels"] = Labels
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.App{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.App
	parsed, ok := resp.Result.(models.App)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListSecretsOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListSecretsOptions) New() *ListSecretsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListSecretsOption func(*ListSecretsOptions)
func (srv *Apps) WithListSecretsQueries(v []string) ListSecretsOption {
	return func(o *ListSecretsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Apps) WithListSecretsTotal(v bool) ListSecretsOption {
	return func(o *ListSecretsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// ListSecrets list client secrets for an application.
func (srv *Apps) ListSecrets(AppId string, optionalSetters ...ListSecretsOption)(*models.AppSecretList, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId))
	path := r.Replace("/apps/{appId}/secrets")
	options := ListSecretsOptions{}.New()
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
		"accept": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AppSecretList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AppSecretList
	parsed, ok := resp.Result.(models.AppSecretList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CreateSecret create a new client secret for an application.
func (srv *Apps) CreateSecret(AppId string)(*models.AppSecretPlaintext, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId))
	path := r.Replace("/apps/{appId}/secrets")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AppSecretPlaintext{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AppSecretPlaintext
	parsed, ok := resp.Result.(models.AppSecretPlaintext)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetSecret get an application client secret by its unique ID.
func (srv *Apps) GetSecret(AppId string, SecretId string)(*models.AppSecret, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId), "{secretId}", client.EncodePath(SecretId))
	path := r.Replace("/apps/{appId}/secrets/{secretId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"accept": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.AppSecret{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.AppSecret
	parsed, ok := resp.Result.(models.AppSecret)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// DeleteSecret delete an application client secret by its unique ID.
func (srv *Apps) DeleteSecret(AppId string, SecretId string)(*interface{}, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId), "{secretId}", client.EncodePath(SecretId))
	path := r.Replace("/apps/{appId}/secrets/{secretId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

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
			
// UpdateTeam transfer an application to another team by its unique ID.
func (srv *Apps) UpdateTeam(AppId string, TeamId string)(*models.App, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId))
	path := r.Replace("/apps/{appId}/team")
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.App{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.App
	parsed, ok := resp.Result.(models.App)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteTokens revoke all tokens for an application by its unique ID.
func (srv *Apps) DeleteTokens(AppId string)(*interface{}, error) {
	r := strings.NewReplacer("{appId}", client.EncodePath(AppId))
	path := r.Replace("/apps/{appId}/tokens")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
		"accept": "application/json",
	}

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
