package organization

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/v6/client"
	"github.com/appwrite/sdk-for-go/v6/models"
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


// Get get the current organization.
func (srv *Organization) Get()(*models.Organization, error) {
	path := "/organization"
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

		parsed := models.Organization{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Organization
	parsed, ok := resp.Result.(models.Organization)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Update update the current organization's name.
func (srv *Organization) Update(Name string)(*models.Organization, error) {
	path := "/organization"
	params := map[string]interface{}{}
	params["name"] = Name
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

		parsed := models.Organization{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Organization
	parsed, ok := resp.Result.(models.Organization)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Delete delete the current organization. All projects that belong to the
// organization are deleted as well.
func (srv *Organization) Delete()(*interface{}, error) {
	path := "/organization"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
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
func (srv *Organization) WithListInstallationsQueries(v []string) ListInstallationsOption {
	return func(o *ListInstallationsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Organization) WithListInstallationsTotal(v bool) ListInstallationsOption {
	return func(o *ListInstallationsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListInstallations list app installations on the organization. Any
// organization member can read installations.
func (srv *Organization) ListInstallations(optionalSetters ...ListInstallationsOption)(*models.AppInstallationList, error) {
	path := "/organization/installations"
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
type CreateInstallationOptions struct {
	AuthorizationDetails string
	enabledSetters map[string]bool
}
func (options CreateInstallationOptions) New() *CreateInstallationOptions {
	options.enabledSetters = map[string]bool{
		"AuthorizationDetails": false,
	}
	return &options
}
type CreateInstallationOption func(*CreateInstallationOptions)
func (srv *Organization) WithCreateInstallationAuthorizationDetails(v string) CreateInstallationOption {
	return func(o *CreateInstallationOptions) {
		o.AuthorizationDetails = v
		o.enabledSetters["AuthorizationDetails"] = true
	}
}
			
// CreateInstallation install an app on the organization. Only organization
// members with the owner role can install apps. The installation is granted
// the scopes the app currently requests.
func (srv *Organization) CreateInstallation(AppId string, optionalSetters ...CreateInstallationOption)(*models.AppInstallation, error) {
	path := "/organization/installations"
	options := CreateInstallationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["appId"] = AppId
	if options.enabledSetters["AuthorizationDetails"] {
		params["authorizationDetails"] = options.AuthorizationDetails
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
	
// GetInstallation get an app installation on the organization by its unique
// ID. Any organization member can read installations.
func (srv *Organization) GetInstallation(InstallationId string)(*models.AppInstallation, error) {
	r := strings.NewReplacer("{installationId}", client.EncodePath(InstallationId))
	path := r.Replace("/organization/installations/{installationId}")
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
type UpdateInstallationOptions struct {
	AuthorizationDetails string
	enabledSetters map[string]bool
}
func (options UpdateInstallationOptions) New() *UpdateInstallationOptions {
	options.enabledSetters = map[string]bool{
		"AuthorizationDetails": false,
	}
	return &options
}
type UpdateInstallationOption func(*UpdateInstallationOptions)
func (srv *Organization) WithUpdateInstallationAuthorizationDetails(v string) UpdateInstallationOption {
	return func(o *UpdateInstallationOptions) {
		o.AuthorizationDetails = v
		o.enabledSetters["AuthorizationDetails"] = true
	}
}
			
// UpdateInstallation update an app installation on the organization. Only
// organization members with the owner role can update installations. The
// installation's granted scopes are refreshed to the scopes the app currently
// requests; previously issued installation access tokens are revoked.
func (srv *Organization) UpdateInstallation(InstallationId string, optionalSetters ...UpdateInstallationOption)(*models.AppInstallation, error) {
	r := strings.NewReplacer("{installationId}", client.EncodePath(InstallationId))
	path := r.Replace("/organization/installations/{installationId}")
	options := UpdateInstallationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["AuthorizationDetails"] {
		params["authorizationDetails"] = options.AuthorizationDetails
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
	
// DeleteInstallation uninstall an app from the organization by its
// installation ID. Only organization members with the owner role can remove
// installations. Previously issued installation access tokens are revoked.
func (srv *Organization) DeleteInstallation(InstallationId string)(*interface{}, error) {
	r := strings.NewReplacer("{installationId}", client.EncodePath(InstallationId))
	path := r.Replace("/organization/installations/{installationId}")
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
	r := strings.NewReplacer("{keyId}", client.EncodePath(KeyId))
	path := r.Replace("/organization/keys/{keyId}")
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
	r := strings.NewReplacer("{keyId}", client.EncodePath(KeyId))
	path := r.Replace("/organization/keys/{keyId}")
	options := UpdateKeyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["scopes"] = Scopes
	if options.enabledSetters["Expire"] {
		params["expire"] = options.Expire
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
	r := strings.NewReplacer("{keyId}", client.EncodePath(KeyId))
	path := r.Replace("/organization/keys/{keyId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
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
type ListMembershipsOptions struct {
	Queries []string
	Search string
	Total bool
	enabledSetters map[string]bool
}
func (options ListMembershipsOptions) New() *ListMembershipsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
		"Total": false,
	}
	return &options
}
type ListMembershipsOption func(*ListMembershipsOptions)
func (srv *Organization) WithListMembershipsQueries(v []string) ListMembershipsOption {
	return func(o *ListMembershipsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Organization) WithListMembershipsSearch(v string) ListMembershipsOption {
	return func(o *ListMembershipsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Organization) WithListMembershipsTotal(v bool) ListMembershipsOption {
	return func(o *ListMembershipsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListMemberships get a list of all memberships from the current
// organization.
func (srv *Organization) ListMemberships(optionalSetters ...ListMembershipsOption)(*models.MembershipList, error) {
	path := "/organization/memberships"
	options := ListMembershipsOptions{}.New()
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

		parsed := models.MembershipList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MembershipList
	parsed, ok := resp.Result.(models.MembershipList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateMembershipOptions struct {
	Email string
	UserId string
	Phone string
	Url string
	Name string
	enabledSetters map[string]bool
}
func (options CreateMembershipOptions) New() *CreateMembershipOptions {
	options.enabledSetters = map[string]bool{
		"Email": false,
		"UserId": false,
		"Phone": false,
		"Url": false,
		"Name": false,
	}
	return &options
}
type CreateMembershipOption func(*CreateMembershipOptions)
func (srv *Organization) WithCreateMembershipEmail(v string) CreateMembershipOption {
	return func(o *CreateMembershipOptions) {
		o.Email = v
		o.enabledSetters["Email"] = true
	}
}
func (srv *Organization) WithCreateMembershipUserId(v string) CreateMembershipOption {
	return func(o *CreateMembershipOptions) {
		o.UserId = v
		o.enabledSetters["UserId"] = true
	}
}
func (srv *Organization) WithCreateMembershipPhone(v string) CreateMembershipOption {
	return func(o *CreateMembershipOptions) {
		o.Phone = v
		o.enabledSetters["Phone"] = true
	}
}
func (srv *Organization) WithCreateMembershipUrl(v string) CreateMembershipOption {
	return func(o *CreateMembershipOptions) {
		o.Url = v
		o.enabledSetters["Url"] = true
	}
}
func (srv *Organization) WithCreateMembershipName(v string) CreateMembershipOption {
	return func(o *CreateMembershipOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
			
// CreateMembership invite a new member to join the current organization. An
// email with a link to join the organization will be sent to the new member's
// email address. If member doesn't exist in the project it will be
// automatically created.
func (srv *Organization) CreateMembership(Roles []string, optionalSetters ...CreateMembershipOption)(*models.Membership, error) {
	path := "/organization/memberships"
	options := CreateMembershipOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Email"] {
		params["email"] = options.Email
	}
	if options.enabledSetters["UserId"] {
		params["userId"] = options.UserId
	}
	if options.enabledSetters["Phone"] {
		params["phone"] = options.Phone
	}
	params["roles"] = Roles
	if options.enabledSetters["Url"] {
		params["url"] = options.Url
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
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

		parsed := models.Membership{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Membership
	parsed, ok := resp.Result.(models.Membership)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetMembership get a membership from the current organization by its unique
// ID.
func (srv *Organization) GetMembership(MembershipId string)(*models.Membership, error) {
	r := strings.NewReplacer("{membershipId}", client.EncodePath(MembershipId))
	path := r.Replace("/organization/memberships/{membershipId}")
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

		parsed := models.Membership{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Membership
	parsed, ok := resp.Result.(models.Membership)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateMembership modify the roles of a member in the current organization.
func (srv *Organization) UpdateMembership(MembershipId string, Roles []string)(*models.Membership, error) {
	r := strings.NewReplacer("{membershipId}", client.EncodePath(MembershipId))
	path := r.Replace("/organization/memberships/{membershipId}")
	params := map[string]interface{}{}
	params["roles"] = Roles
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

		parsed := models.Membership{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Membership
	parsed, ok := resp.Result.(models.Membership)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteMembership remove a member from the current organization. The member
// is removed whether they accepted the invitation or not; a pending
// invitation is revoked.
func (srv *Organization) DeleteMembership(MembershipId string)(*interface{}, error) {
	r := strings.NewReplacer("{membershipId}", client.EncodePath(MembershipId))
	path := r.Replace("/organization/memberships/{membershipId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
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
	r := strings.NewReplacer("{projectId}", client.EncodePath(ProjectId))
	path := r.Replace("/organization/projects/{projectId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
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
	r := strings.NewReplacer("{projectId}", client.EncodePath(ProjectId))
	path := r.Replace("/organization/projects/{projectId}")
	params := map[string]interface{}{}
	params["name"] = Name
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
	r := strings.NewReplacer("{projectId}", client.EncodePath(ProjectId))
	path := r.Replace("/organization/projects/{projectId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"X-Appwrite-Project": srv.client.Config["project"],
		"content-type": "application/json",
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
