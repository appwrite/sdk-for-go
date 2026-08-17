package oauth2

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/models"
)

// Oauth2 service
type Oauth2 struct {
	client client.Client
}

func New(clt client.Client) *Oauth2 {
	return &Oauth2{
		client: clt,
	}
}

type ApproveOptions struct {
	AuthorizationDetails string
	Scope                string
	enabledSetters       map[string]bool
}

func (options ApproveOptions) New() *ApproveOptions {
	options.enabledSetters = map[string]bool{"AuthorizationDetails": false, "Scope": false}
	return &options
}

type ApproveOption func(*ApproveOptions)

func (srv *Oauth2) WithApproveAuthorizationDetails(v string) ApproveOption {
	return func(o *ApproveOptions) {
		o.AuthorizationDetails = v
		o.enabledSetters["AuthorizationDetails"] = true
	}
}
func (srv *Oauth2) WithApproveScope(v string) ApproveOption {
	return func(o *ApproveOptions) {
		o.Scope = v
		o.enabledSetters["Scope"] = true
	}
}

// Approve approve an OAuth2 grant after the user gives consent. Returns the
// `redirectUrl` the end user should be sent to. The consent screen may
// optionally pass enriched `authorization_details` to record the concrete
// resources the user selected. You can pass Accept header of
// `application/json` to receive a JSON response instead of a redirect.
func (srv *Oauth2) Approve(GrantId string, optionalSetters ...ApproveOption) (*models.Oauth2Approve, error) {
	r := strings.NewReplacer("{project_id}", client.EncodePath(srv.client.Config["project"]))
	path := r.Replace("/oauth2/{project_id}/approve")
	options := ApproveOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["grant_id"] = GrantId
	if options.enabledSetters["AuthorizationDetails"] {
		params["authorization_details"] = options.AuthorizationDetails
	}
	if options.enabledSetters["Scope"] {
		params["scope"] = options.Scope
	}
	headers := map[string]interface{}{}
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

		parsed := models.Oauth2Approve{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Oauth2Approve
	parsed, ok := resp.Result.(models.Oauth2Approve)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type AuthorizeOptions struct {
	ClientId             string
	RedirectUri          string
	ResponseType         string
	Scope                string
	State                string
	Nonce                string
	CodeChallenge        string
	CodeChallengeMethod  string
	Prompt               string
	MaxAge               int
	AuthorizationDetails string
	Resource             string
	Audience             string
	RequestUri           string
	enabledSetters       map[string]bool
}

func (options AuthorizeOptions) New() *AuthorizeOptions {
	options.enabledSetters = map[string]bool{"ClientId": false, "RedirectUri": false, "ResponseType": false, "Scope": false, "State": false, "Nonce": false, "CodeChallenge": false, "CodeChallengeMethod": false, "Prompt": false, "MaxAge": false, "AuthorizationDetails": false, "Resource": false, "Audience": false, "RequestUri": false}
	return &options
}

type AuthorizeOption func(*AuthorizeOptions)

func (srv *Oauth2) WithAuthorizeClientId(v string) AuthorizeOption {
	return func(o *AuthorizeOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Oauth2) WithAuthorizeRedirectUri(v string) AuthorizeOption {
	return func(o *AuthorizeOptions) {
		o.RedirectUri = v
		o.enabledSetters["RedirectUri"] = true
	}
}
func (srv *Oauth2) WithAuthorizeResponseType(v string) AuthorizeOption {
	return func(o *AuthorizeOptions) {
		o.ResponseType = v
		o.enabledSetters["ResponseType"] = true
	}
}
func (srv *Oauth2) WithAuthorizeScope(v string) AuthorizeOption {
	return func(o *AuthorizeOptions) {
		o.Scope = v
		o.enabledSetters["Scope"] = true
	}
}
func (srv *Oauth2) WithAuthorizeState(v string) AuthorizeOption {
	return func(o *AuthorizeOptions) {
		o.State = v
		o.enabledSetters["State"] = true
	}
}
func (srv *Oauth2) WithAuthorizeNonce(v string) AuthorizeOption {
	return func(o *AuthorizeOptions) {
		o.Nonce = v
		o.enabledSetters["Nonce"] = true
	}
}
func (srv *Oauth2) WithAuthorizeCodeChallenge(v string) AuthorizeOption {
	return func(o *AuthorizeOptions) {
		o.CodeChallenge = v
		o.enabledSetters["CodeChallenge"] = true
	}
}
func (srv *Oauth2) WithAuthorizeCodeChallengeMethod(v string) AuthorizeOption {
	return func(o *AuthorizeOptions) {
		o.CodeChallengeMethod = v
		o.enabledSetters["CodeChallengeMethod"] = true
	}
}
func (srv *Oauth2) WithAuthorizePrompt(v string) AuthorizeOption {
	return func(o *AuthorizeOptions) {
		o.Prompt = v
		o.enabledSetters["Prompt"] = true
	}
}
func (srv *Oauth2) WithAuthorizeMaxAge(v int) AuthorizeOption {
	return func(o *AuthorizeOptions) {
		o.MaxAge = v
		o.enabledSetters["MaxAge"] = true
	}
}
func (srv *Oauth2) WithAuthorizeAuthorizationDetails(v string) AuthorizeOption {
	return func(o *AuthorizeOptions) {
		o.AuthorizationDetails = v
		o.enabledSetters["AuthorizationDetails"] = true
	}
}
func (srv *Oauth2) WithAuthorizeResource(v string) AuthorizeOption {
	return func(o *AuthorizeOptions) {
		o.Resource = v
		o.enabledSetters["Resource"] = true
	}
}
func (srv *Oauth2) WithAuthorizeAudience(v string) AuthorizeOption {
	return func(o *AuthorizeOptions) {
		o.Audience = v
		o.enabledSetters["Audience"] = true
	}
}
func (srv *Oauth2) WithAuthorizeRequestUri(v string) AuthorizeOption {
	return func(o *AuthorizeOptions) {
		o.RequestUri = v
		o.enabledSetters["RequestUri"] = true
	}
}

// Authorize begin the OAuth2 authorization flow. When called without a
// session, the user is redirected to the consent screen without grant ID.
// When called with a session, the redirect URL includes param for grant ID.
// You can pass Accept header of `application/json` to receive a JSON response
// instead of a redirect.
func (srv *Oauth2) Authorize(optionalSetters ...AuthorizeOption) (*models.Oauth2Authorize, error) {
	r := strings.NewReplacer("{project_id}", client.EncodePath(srv.client.Config["project"]))
	path := r.Replace("/oauth2/{project_id}/authorize")
	options := AuthorizeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["client_id"] = options.ClientId
	}
	if options.enabledSetters["RedirectUri"] {
		params["redirect_uri"] = options.RedirectUri
	}
	if options.enabledSetters["ResponseType"] {
		params["response_type"] = options.ResponseType
	}
	if options.enabledSetters["Scope"] {
		params["scope"] = options.Scope
	}
	if options.enabledSetters["State"] {
		params["state"] = options.State
	}
	if options.enabledSetters["Nonce"] {
		params["nonce"] = options.Nonce
	}
	if options.enabledSetters["CodeChallenge"] {
		params["code_challenge"] = options.CodeChallenge
	}
	if options.enabledSetters["CodeChallengeMethod"] {
		params["code_challenge_method"] = options.CodeChallengeMethod
	}
	if options.enabledSetters["Prompt"] {
		params["prompt"] = options.Prompt
	}
	if options.enabledSetters["MaxAge"] {
		params["max_age"] = options.MaxAge
	}
	if options.enabledSetters["AuthorizationDetails"] {
		params["authorization_details"] = options.AuthorizationDetails
	}
	if options.enabledSetters["Resource"] {
		params["resource"] = options.Resource
	}
	if options.enabledSetters["Audience"] {
		params["audience"] = options.Audience
	}
	if options.enabledSetters["RequestUri"] {
		params["request_uri"] = options.RequestUri
	}
	headers := map[string]interface{}{}
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

		parsed := models.Oauth2Authorize{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Oauth2Authorize
	parsed, ok := resp.Result.(models.Oauth2Authorize)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type AuthorizePostOptions struct {
	ClientId             string
	RedirectUri          string
	ResponseType         string
	Scope                string
	State                string
	Nonce                string
	CodeChallenge        string
	CodeChallengeMethod  string
	Prompt               string
	MaxAge               int
	AuthorizationDetails string
	Resource             string
	Audience             string
	RequestUri           string
	enabledSetters       map[string]bool
}

func (options AuthorizePostOptions) New() *AuthorizePostOptions {
	options.enabledSetters = map[string]bool{"ClientId": false, "RedirectUri": false, "ResponseType": false, "Scope": false, "State": false, "Nonce": false, "CodeChallenge": false, "CodeChallengeMethod": false, "Prompt": false, "MaxAge": false, "AuthorizationDetails": false, "Resource": false, "Audience": false, "RequestUri": false}
	return &options
}

type AuthorizePostOption func(*AuthorizePostOptions)

func (srv *Oauth2) WithAuthorizePostClientId(v string) AuthorizePostOption {
	return func(o *AuthorizePostOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Oauth2) WithAuthorizePostRedirectUri(v string) AuthorizePostOption {
	return func(o *AuthorizePostOptions) {
		o.RedirectUri = v
		o.enabledSetters["RedirectUri"] = true
	}
}
func (srv *Oauth2) WithAuthorizePostResponseType(v string) AuthorizePostOption {
	return func(o *AuthorizePostOptions) {
		o.ResponseType = v
		o.enabledSetters["ResponseType"] = true
	}
}
func (srv *Oauth2) WithAuthorizePostScope(v string) AuthorizePostOption {
	return func(o *AuthorizePostOptions) {
		o.Scope = v
		o.enabledSetters["Scope"] = true
	}
}
func (srv *Oauth2) WithAuthorizePostState(v string) AuthorizePostOption {
	return func(o *AuthorizePostOptions) {
		o.State = v
		o.enabledSetters["State"] = true
	}
}
func (srv *Oauth2) WithAuthorizePostNonce(v string) AuthorizePostOption {
	return func(o *AuthorizePostOptions) {
		o.Nonce = v
		o.enabledSetters["Nonce"] = true
	}
}
func (srv *Oauth2) WithAuthorizePostCodeChallenge(v string) AuthorizePostOption {
	return func(o *AuthorizePostOptions) {
		o.CodeChallenge = v
		o.enabledSetters["CodeChallenge"] = true
	}
}
func (srv *Oauth2) WithAuthorizePostCodeChallengeMethod(v string) AuthorizePostOption {
	return func(o *AuthorizePostOptions) {
		o.CodeChallengeMethod = v
		o.enabledSetters["CodeChallengeMethod"] = true
	}
}
func (srv *Oauth2) WithAuthorizePostPrompt(v string) AuthorizePostOption {
	return func(o *AuthorizePostOptions) {
		o.Prompt = v
		o.enabledSetters["Prompt"] = true
	}
}
func (srv *Oauth2) WithAuthorizePostMaxAge(v int) AuthorizePostOption {
	return func(o *AuthorizePostOptions) {
		o.MaxAge = v
		o.enabledSetters["MaxAge"] = true
	}
}
func (srv *Oauth2) WithAuthorizePostAuthorizationDetails(v string) AuthorizePostOption {
	return func(o *AuthorizePostOptions) {
		o.AuthorizationDetails = v
		o.enabledSetters["AuthorizationDetails"] = true
	}
}
func (srv *Oauth2) WithAuthorizePostResource(v string) AuthorizePostOption {
	return func(o *AuthorizePostOptions) {
		o.Resource = v
		o.enabledSetters["Resource"] = true
	}
}
func (srv *Oauth2) WithAuthorizePostAudience(v string) AuthorizePostOption {
	return func(o *AuthorizePostOptions) {
		o.Audience = v
		o.enabledSetters["Audience"] = true
	}
}
func (srv *Oauth2) WithAuthorizePostRequestUri(v string) AuthorizePostOption {
	return func(o *AuthorizePostOptions) {
		o.RequestUri = v
		o.enabledSetters["RequestUri"] = true
	}
}

// AuthorizePost begin the OAuth2 authorization flow. When called without a
// session, the user is redirected to the consent screen without grant ID.
// When called with a session, the redirect URL includes param for grant ID.
// You can pass Accept header of `application/json` to receive a JSON response
// instead of a redirect.
func (srv *Oauth2) AuthorizePost(optionalSetters ...AuthorizePostOption) (*models.Oauth2Authorize, error) {
	r := strings.NewReplacer("{project_id}", client.EncodePath(srv.client.Config["project"]))
	path := r.Replace("/oauth2/{project_id}/authorize")
	options := AuthorizePostOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["client_id"] = options.ClientId
	}
	if options.enabledSetters["RedirectUri"] {
		params["redirect_uri"] = options.RedirectUri
	}
	if options.enabledSetters["ResponseType"] {
		params["response_type"] = options.ResponseType
	}
	if options.enabledSetters["Scope"] {
		params["scope"] = options.Scope
	}
	if options.enabledSetters["State"] {
		params["state"] = options.State
	}
	if options.enabledSetters["Nonce"] {
		params["nonce"] = options.Nonce
	}
	if options.enabledSetters["CodeChallenge"] {
		params["code_challenge"] = options.CodeChallenge
	}
	if options.enabledSetters["CodeChallengeMethod"] {
		params["code_challenge_method"] = options.CodeChallengeMethod
	}
	if options.enabledSetters["Prompt"] {
		params["prompt"] = options.Prompt
	}
	if options.enabledSetters["MaxAge"] {
		params["max_age"] = options.MaxAge
	}
	if options.enabledSetters["AuthorizationDetails"] {
		params["authorization_details"] = options.AuthorizationDetails
	}
	if options.enabledSetters["Resource"] {
		params["resource"] = options.Resource
	}
	if options.enabledSetters["Audience"] {
		params["audience"] = options.Audience
	}
	if options.enabledSetters["RequestUri"] {
		params["request_uri"] = options.RequestUri
	}
	headers := map[string]interface{}{}
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

		parsed := models.Oauth2Authorize{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Oauth2Authorize
	parsed, ok := resp.Result.(models.Oauth2Authorize)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateDeviceAuthorizationOptions struct {
	ClientId             string
	Scope                string
	AuthorizationDetails string
	Resource             string
	Audience             string
	enabledSetters       map[string]bool
}

func (options CreateDeviceAuthorizationOptions) New() *CreateDeviceAuthorizationOptions {
	options.enabledSetters = map[string]bool{"ClientId": false, "Scope": false, "AuthorizationDetails": false, "Resource": false, "Audience": false}
	return &options
}

type CreateDeviceAuthorizationOption func(*CreateDeviceAuthorizationOptions)

func (srv *Oauth2) WithCreateDeviceAuthorizationClientId(v string) CreateDeviceAuthorizationOption {
	return func(o *CreateDeviceAuthorizationOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Oauth2) WithCreateDeviceAuthorizationScope(v string) CreateDeviceAuthorizationOption {
	return func(o *CreateDeviceAuthorizationOptions) {
		o.Scope = v
		o.enabledSetters["Scope"] = true
	}
}
func (srv *Oauth2) WithCreateDeviceAuthorizationAuthorizationDetails(v string) CreateDeviceAuthorizationOption {
	return func(o *CreateDeviceAuthorizationOptions) {
		o.AuthorizationDetails = v
		o.enabledSetters["AuthorizationDetails"] = true
	}
}
func (srv *Oauth2) WithCreateDeviceAuthorizationResource(v string) CreateDeviceAuthorizationOption {
	return func(o *CreateDeviceAuthorizationOptions) {
		o.Resource = v
		o.enabledSetters["Resource"] = true
	}
}
func (srv *Oauth2) WithCreateDeviceAuthorizationAudience(v string) CreateDeviceAuthorizationOption {
	return func(o *CreateDeviceAuthorizationOptions) {
		o.Audience = v
		o.enabledSetters["Audience"] = true
	}
}

// CreateDeviceAuthorization start the OAuth2 Device Authorization Grant.
// Returns the device code, user code, verification URL, expiration, and
// polling interval.
func (srv *Oauth2) CreateDeviceAuthorization(optionalSetters ...CreateDeviceAuthorizationOption) (*models.Oauth2DeviceAuthorization, error) {
	r := strings.NewReplacer("{project_id}", client.EncodePath(srv.client.Config["project"]))
	path := r.Replace("/oauth2/{project_id}/device_authorization")
	options := CreateDeviceAuthorizationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ClientId"] {
		params["client_id"] = options.ClientId
	}
	if options.enabledSetters["Scope"] {
		params["scope"] = options.Scope
	}
	if options.enabledSetters["AuthorizationDetails"] {
		params["authorization_details"] = options.AuthorizationDetails
	}
	if options.enabledSetters["Resource"] {
		params["resource"] = options.Resource
	}
	if options.enabledSetters["Audience"] {
		params["audience"] = options.Audience
	}
	headers := map[string]interface{}{}
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

		parsed := models.Oauth2DeviceAuthorization{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Oauth2DeviceAuthorization
	parsed, ok := resp.Result.(models.Oauth2DeviceAuthorization)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreateGrant exchange a device flow user code for an OAuth2 grant. The
// authenticated user is bound to the pending grant. Pass the returned grant
// ID to the get grant endpoint to render the consent screen, then to the
// approve or reject endpoint to complete the flow.
func (srv *Oauth2) CreateGrant(UserCode string) (*models.Oauth2Grant, error) {
	r := strings.NewReplacer("{project_id}", client.EncodePath(srv.client.Config["project"]))
	path := r.Replace("/oauth2/{project_id}/grants")
	params := map[string]interface{}{}
	params["user_code"] = UserCode
	headers := map[string]interface{}{}
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

		parsed := models.Oauth2Grant{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Oauth2Grant
	parsed, ok := resp.Result.(models.Oauth2Grant)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetGrant get an OAuth2 grant by its ID. Used by the consent screen to
// display the details of the authorization the user is being asked to
// approve. A grant can only be read by the user it belongs to, or by server
// SDK.
func (srv *Oauth2) GetGrant(GrantId string) (*models.Oauth2Grant, error) {
	r := strings.NewReplacer("{project_id}", client.EncodePath(srv.client.Config["project"]), "{grant_id}", client.EncodePath(GrantId))
	path := r.Replace("/oauth2/{project_id}/grants/{grant_id}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
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

		parsed := models.Oauth2Grant{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Oauth2Grant
	parsed, ok := resp.Result.(models.Oauth2Grant)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListOrganizationsOptions struct {
	Limit          int
	Offset         int
	Search         string
	enabledSetters map[string]bool
}

func (options ListOrganizationsOptions) New() *ListOrganizationsOptions {
	options.enabledSetters = map[string]bool{"Limit": false, "Offset": false, "Search": false}
	return &options
}

type ListOrganizationsOption func(*ListOrganizationsOptions)

func (srv *Oauth2) WithListOrganizationsLimit(v int) ListOrganizationsOption {
	return func(o *ListOrganizationsOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Oauth2) WithListOrganizationsOffset(v int) ListOrganizationsOption {
	return func(o *ListOrganizationsOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Oauth2) WithListOrganizationsSearch(v string) ListOrganizationsOption {
	return func(o *ListOrganizationsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}

// ListOrganizations list the organizations the OAuth2 access token can
// access. Resolves the token's `organization` authorization details,
// expanding the `*` wildcard into the concrete set of organizations the user
// can see.
func (srv *Oauth2) ListOrganizations(optionalSetters ...ListOrganizationsOption) (*models.Oauth2OrganizationList, error) {
	r := strings.NewReplacer("{project_id}", client.EncodePath(srv.client.Config["project"]))
	path := r.Replace("/oauth2/{project_id}/organizations")
	options := ListOrganizationsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	headers := map[string]interface{}{}
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

		parsed := models.Oauth2OrganizationList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Oauth2OrganizationList
	parsed, ok := resp.Result.(models.Oauth2OrganizationList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreatePAROptions struct {
	Scope                string
	State                string
	Nonce                string
	CodeChallenge        string
	CodeChallengeMethod  string
	Prompt               string
	MaxAge               int
	AuthorizationDetails string
	Resource             string
	Audience             string
	enabledSetters       map[string]bool
}

func (options CreatePAROptions) New() *CreatePAROptions {
	options.enabledSetters = map[string]bool{"Scope": false, "State": false, "Nonce": false, "CodeChallenge": false, "CodeChallengeMethod": false, "Prompt": false, "MaxAge": false, "AuthorizationDetails": false, "Resource": false, "Audience": false}
	return &options
}

type CreatePAROption func(*CreatePAROptions)

func (srv *Oauth2) WithCreatePARScope(v string) CreatePAROption {
	return func(o *CreatePAROptions) {
		o.Scope = v
		o.enabledSetters["Scope"] = true
	}
}
func (srv *Oauth2) WithCreatePARState(v string) CreatePAROption {
	return func(o *CreatePAROptions) {
		o.State = v
		o.enabledSetters["State"] = true
	}
}
func (srv *Oauth2) WithCreatePARNonce(v string) CreatePAROption {
	return func(o *CreatePAROptions) {
		o.Nonce = v
		o.enabledSetters["Nonce"] = true
	}
}
func (srv *Oauth2) WithCreatePARCodeChallenge(v string) CreatePAROption {
	return func(o *CreatePAROptions) {
		o.CodeChallenge = v
		o.enabledSetters["CodeChallenge"] = true
	}
}
func (srv *Oauth2) WithCreatePARCodeChallengeMethod(v string) CreatePAROption {
	return func(o *CreatePAROptions) {
		o.CodeChallengeMethod = v
		o.enabledSetters["CodeChallengeMethod"] = true
	}
}
func (srv *Oauth2) WithCreatePARPrompt(v string) CreatePAROption {
	return func(o *CreatePAROptions) {
		o.Prompt = v
		o.enabledSetters["Prompt"] = true
	}
}
func (srv *Oauth2) WithCreatePARMaxAge(v int) CreatePAROption {
	return func(o *CreatePAROptions) {
		o.MaxAge = v
		o.enabledSetters["MaxAge"] = true
	}
}
func (srv *Oauth2) WithCreatePARAuthorizationDetails(v string) CreatePAROption {
	return func(o *CreatePAROptions) {
		o.AuthorizationDetails = v
		o.enabledSetters["AuthorizationDetails"] = true
	}
}
func (srv *Oauth2) WithCreatePARResource(v string) CreatePAROption {
	return func(o *CreatePAROptions) {
		o.Resource = v
		o.enabledSetters["Resource"] = true
	}
}
func (srv *Oauth2) WithCreatePARAudience(v string) CreatePAROption {
	return func(o *CreatePAROptions) {
		o.Audience = v
		o.enabledSetters["Audience"] = true
	}
}

// CreatePAR store an OAuth2 authorization request server-side and receive a
// short-lived request_uri handle for the authorize endpoint.
func (srv *Oauth2) CreatePAR(ClientId string, RedirectUri string, ResponseType string, optionalSetters ...CreatePAROption) (*models.Oauth2PAR, error) {
	r := strings.NewReplacer("{project_id}", client.EncodePath(srv.client.Config["project"]))
	path := r.Replace("/oauth2/{project_id}/par")
	options := CreatePAROptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["client_id"] = ClientId
	params["redirect_uri"] = RedirectUri
	params["response_type"] = ResponseType
	if options.enabledSetters["Scope"] {
		params["scope"] = options.Scope
	}
	if options.enabledSetters["State"] {
		params["state"] = options.State
	}
	if options.enabledSetters["Nonce"] {
		params["nonce"] = options.Nonce
	}
	if options.enabledSetters["CodeChallenge"] {
		params["code_challenge"] = options.CodeChallenge
	}
	if options.enabledSetters["CodeChallengeMethod"] {
		params["code_challenge_method"] = options.CodeChallengeMethod
	}
	if options.enabledSetters["Prompt"] {
		params["prompt"] = options.Prompt
	}
	if options.enabledSetters["MaxAge"] {
		params["max_age"] = options.MaxAge
	}
	if options.enabledSetters["AuthorizationDetails"] {
		params["authorization_details"] = options.AuthorizationDetails
	}
	if options.enabledSetters["Resource"] {
		params["resource"] = options.Resource
	}
	if options.enabledSetters["Audience"] {
		params["audience"] = options.Audience
	}
	headers := map[string]interface{}{}
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

		parsed := models.Oauth2PAR{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Oauth2PAR
	parsed, ok := resp.Result.(models.Oauth2PAR)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListProjectsOptions struct {
	Limit          int
	Offset         int
	Search         string
	enabledSetters map[string]bool
}

func (options ListProjectsOptions) New() *ListProjectsOptions {
	options.enabledSetters = map[string]bool{"Limit": false, "Offset": false, "Search": false}
	return &options
}

type ListProjectsOption func(*ListProjectsOptions)

func (srv *Oauth2) WithListProjectsLimit(v int) ListProjectsOption {
	return func(o *ListProjectsOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Oauth2) WithListProjectsOffset(v int) ListProjectsOption {
	return func(o *ListProjectsOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Oauth2) WithListProjectsSearch(v string) ListProjectsOption {
	return func(o *ListProjectsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}

// ListProjects list the projects the OAuth2 access token can access. Resolves
// the token's `project` authorization details, expanding the `*` wildcard
// into the concrete set of projects the user can see.
func (srv *Oauth2) ListProjects(optionalSetters ...ListProjectsOption) (*models.Oauth2ProjectList, error) {
	r := strings.NewReplacer("{project_id}", client.EncodePath(srv.client.Config["project"]))
	path := r.Replace("/oauth2/{project_id}/projects")
	options := ListProjectsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	headers := map[string]interface{}{}
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

		parsed := models.Oauth2ProjectList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Oauth2ProjectList
	parsed, ok := resp.Result.(models.Oauth2ProjectList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Reject reject an OAuth2 grant when the user denies consent. Returns the
// `redirectUrl` the end user should be sent to with an `access_denied` error.
// You can pass Accept header of `application/json` to receive a JSON response
// instead of a redirect.
func (srv *Oauth2) Reject(GrantId string) (*models.Oauth2Reject, error) {
	r := strings.NewReplacer("{project_id}", client.EncodePath(srv.client.Config["project"]))
	path := r.Replace("/oauth2/{project_id}/reject")
	params := map[string]interface{}{}
	params["grant_id"] = GrantId
	headers := map[string]interface{}{}
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

		parsed := models.Oauth2Reject{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Oauth2Reject
	parsed, ok := resp.Result.(models.Oauth2Reject)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type RevokeOptions struct {
	TokenTypeHint  string
	ClientId       string
	ClientSecret   string
	enabledSetters map[string]bool
}

func (options RevokeOptions) New() *RevokeOptions {
	options.enabledSetters = map[string]bool{"TokenTypeHint": false, "ClientId": false, "ClientSecret": false}
	return &options
}

type RevokeOption func(*RevokeOptions)

func (srv *Oauth2) WithRevokeTokenTypeHint(v string) RevokeOption {
	return func(o *RevokeOptions) {
		o.TokenTypeHint = v
		o.enabledSetters["TokenTypeHint"] = true
	}
}
func (srv *Oauth2) WithRevokeClientId(v string) RevokeOption {
	return func(o *RevokeOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Oauth2) WithRevokeClientSecret(v string) RevokeOption {
	return func(o *RevokeOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}

// Revoke revoke an OAuth2 access token or refresh token.
func (srv *Oauth2) Revoke(Token string, optionalSetters ...RevokeOption) (*interface{}, error) {
	r := strings.NewReplacer("{project_id}", client.EncodePath(srv.client.Config["project"]))
	path := r.Replace("/oauth2/{project_id}/revoke")
	options := RevokeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["token"] = Token
	if options.enabledSetters["TokenTypeHint"] {
		params["token_type_hint"] = options.TokenTypeHint
	}
	if options.enabledSetters["ClientId"] {
		params["client_id"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["client_secret"] = options.ClientSecret
	}
	headers := map[string]interface{}{}
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

type CreateTokenOptions struct {
	Code           string
	RefreshToken   string
	DeviceCode     string
	ClientId       string
	ClientSecret   string
	CodeVerifier   string
	RedirectUri    string
	Resource       string
	Audience       string
	enabledSetters map[string]bool
}

func (options CreateTokenOptions) New() *CreateTokenOptions {
	options.enabledSetters = map[string]bool{"Code": false, "RefreshToken": false, "DeviceCode": false, "ClientId": false, "ClientSecret": false, "CodeVerifier": false, "RedirectUri": false, "Resource": false, "Audience": false}
	return &options
}

type CreateTokenOption func(*CreateTokenOptions)

func (srv *Oauth2) WithCreateTokenCode(v string) CreateTokenOption {
	return func(o *CreateTokenOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Oauth2) WithCreateTokenRefreshToken(v string) CreateTokenOption {
	return func(o *CreateTokenOptions) {
		o.RefreshToken = v
		o.enabledSetters["RefreshToken"] = true
	}
}
func (srv *Oauth2) WithCreateTokenDeviceCode(v string) CreateTokenOption {
	return func(o *CreateTokenOptions) {
		o.DeviceCode = v
		o.enabledSetters["DeviceCode"] = true
	}
}
func (srv *Oauth2) WithCreateTokenClientId(v string) CreateTokenOption {
	return func(o *CreateTokenOptions) {
		o.ClientId = v
		o.enabledSetters["ClientId"] = true
	}
}
func (srv *Oauth2) WithCreateTokenClientSecret(v string) CreateTokenOption {
	return func(o *CreateTokenOptions) {
		o.ClientSecret = v
		o.enabledSetters["ClientSecret"] = true
	}
}
func (srv *Oauth2) WithCreateTokenCodeVerifier(v string) CreateTokenOption {
	return func(o *CreateTokenOptions) {
		o.CodeVerifier = v
		o.enabledSetters["CodeVerifier"] = true
	}
}
func (srv *Oauth2) WithCreateTokenRedirectUri(v string) CreateTokenOption {
	return func(o *CreateTokenOptions) {
		o.RedirectUri = v
		o.enabledSetters["RedirectUri"] = true
	}
}
func (srv *Oauth2) WithCreateTokenResource(v string) CreateTokenOption {
	return func(o *CreateTokenOptions) {
		o.Resource = v
		o.enabledSetters["Resource"] = true
	}
}
func (srv *Oauth2) WithCreateTokenAudience(v string) CreateTokenOption {
	return func(o *CreateTokenOptions) {
		o.Audience = v
		o.enabledSetters["Audience"] = true
	}
}

// CreateToken exchange an OAuth2 authorization code, refresh token, or device
// code for access and refresh tokens.
func (srv *Oauth2) CreateToken(GrantType string, optionalSetters ...CreateTokenOption) (*models.Oauth2Token, error) {
	r := strings.NewReplacer("{project_id}", client.EncodePath(srv.client.Config["project"]))
	path := r.Replace("/oauth2/{project_id}/token")
	options := CreateTokenOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["grant_type"] = GrantType
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["RefreshToken"] {
		params["refresh_token"] = options.RefreshToken
	}
	if options.enabledSetters["DeviceCode"] {
		params["device_code"] = options.DeviceCode
	}
	if options.enabledSetters["ClientId"] {
		params["client_id"] = options.ClientId
	}
	if options.enabledSetters["ClientSecret"] {
		params["client_secret"] = options.ClientSecret
	}
	if options.enabledSetters["CodeVerifier"] {
		params["code_verifier"] = options.CodeVerifier
	}
	if options.enabledSetters["RedirectUri"] {
		params["redirect_uri"] = options.RedirectUri
	}
	if options.enabledSetters["Resource"] {
		params["resource"] = options.Resource
	}
	if options.enabledSetters["Audience"] {
		params["audience"] = options.Audience
	}
	headers := map[string]interface{}{}
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
