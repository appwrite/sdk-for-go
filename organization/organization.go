package organization

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/models"
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
func (srv *Organization) Get() (*models.Organization, error) {
	path := "/organization"
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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
func (srv *Organization) Update(Name string) (*models.Organization, error) {
	path := "/organization"
	params := map[string]interface{}{}
	params["name"] = Name
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

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
func (srv *Organization) Delete() (*interface{}, error) {
	path := "/organization"
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

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

type ListEscalationsOptions struct {
	Queries        []string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListEscalationsOptions) New() *ListEscalationsOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Total": false}
	return &options
}

type ListEscalationsOption func(*ListEscalationsOptions)

func (srv *Organization) WithListEscalationsQueries(v []string) ListEscalationsOption {
	return func(o *ListEscalationsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Organization) WithListEscalationsTotal(v bool) ListEscalationsOption {
	return func(o *ListEscalationsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListEscalations get a list of your organization's escalations. Use
// `Query.search('search', term)` to search escalation IDs and names.
func (srv *Organization) ListEscalations(optionalSetters ...ListEscalationsOption) (*models.EscalationList, error) {
	path := "/organization/escalations"
	options := ListEscalationsOptions{}.New()
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
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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

		parsed := models.EscalationList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.EscalationList
	parsed, ok := resp.Result.(models.EscalationList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateEscalationOptions struct {
	Repeat         int
	enabledSetters map[string]bool
}

func (options CreateEscalationOptions) New() *CreateEscalationOptions {
	options.enabledSetters = map[string]bool{"Repeat": false}
	return &options
}

type CreateEscalationOption func(*CreateEscalationOptions)

func (srv *Organization) WithCreateEscalationRepeat(v int) CreateEscalationOption {
	return func(o *CreateEscalationOptions) {
		o.Repeat = v
		o.enabledSetters["Repeat"] = true
	}
}

// CreateEscalation create an escalation for your organization. A monitor
// using it notifies the escalation's steps one after another while its alert
// fires, each by email, push and in the console, until a member acknowledges
// the alert or it resolves. Each step waits its `delay` in seconds, after the
// alert fires for the first step and after the previous step for the others,
// then notifies the members holding its `roles` and those named in `userIds`.
// With `repeat`, the steps run again from the first, a monitor's cooldown
// after the last step.
func (srv *Organization) CreateEscalation(EscalationId string, Name string, Steps []interface{}, optionalSetters ...CreateEscalationOption) (*models.Escalation, error) {
	path := "/organization/escalations"
	options := CreateEscalationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["escalationId"] = EscalationId
	params["name"] = Name
	params["steps"] = Steps
	if options.enabledSetters["Repeat"] {
		params["repeat"] = options.Repeat
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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

		parsed := models.Escalation{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Escalation
	parsed, ok := resp.Result.(models.Escalation)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetEscalation get one of your organization's escalations by its ID.
func (srv *Organization) GetEscalation(EscalationId string) (*models.Escalation, error) {
	if EscalationId == "" {
		return nil, errors.New("Missing required parameter: \"escalationId\"")
	}

	r := strings.NewReplacer("{escalationId}", client.EncodePath(EscalationId))
	path := r.Replace("/organization/escalations/{escalationId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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

		parsed := models.Escalation{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Escalation
	parsed, ok := resp.Result.(models.Escalation)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateEscalationOptions struct {
	Name           string
	Steps          []interface{}
	Repeat         int
	enabledSetters map[string]bool
}

func (options UpdateEscalationOptions) New() *UpdateEscalationOptions {
	options.enabledSetters = map[string]bool{"Name": false, "Steps": false, "Repeat": false}
	return &options
}

type UpdateEscalationOption func(*UpdateEscalationOptions)

func (srv *Organization) WithUpdateEscalationName(v string) UpdateEscalationOption {
	return func(o *UpdateEscalationOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Organization) WithUpdateEscalationSteps(v []interface{}) UpdateEscalationOption {
	return func(o *UpdateEscalationOptions) {
		o.Steps = v
		o.enabledSetters["Steps"] = true
	}
}
func (srv *Organization) WithUpdateEscalationRepeat(v int) UpdateEscalationOption {
	return func(o *UpdateEscalationOptions) {
		o.Repeat = v
		o.enabledSetters["Repeat"] = true
	}
}

// UpdateEscalation update an escalation's name, steps or repeats; what you
// leave out stays as it is. Alerts already escalating keep the steps they
// started with; the next alert uses the new ones.
func (srv *Organization) UpdateEscalation(EscalationId string, optionalSetters ...UpdateEscalationOption) (*models.Escalation, error) {
	if EscalationId == "" {
		return nil, errors.New("Missing required parameter: \"escalationId\"")
	}

	r := strings.NewReplacer("{escalationId}", client.EncodePath(EscalationId))
	path := r.Replace("/organization/escalations/{escalationId}")
	options := UpdateEscalationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Steps"] {
		params["steps"] = options.Steps
	}
	if options.enabledSetters["Repeat"] {
		params["repeat"] = options.Repeat
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.Escalation{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Escalation
	parsed, ok := resp.Result.(models.Escalation)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteEscalation delete an escalation. An escalation that a monitor uses
// cannot be deleted: move the monitor to another escalation or to `notify`
// roles first. Open alerts keep the steps they started with.
func (srv *Organization) DeleteEscalation(EscalationId string) (*interface{}, error) {
	if EscalationId == "" {
		return nil, errors.New("Missing required parameter: \"escalationId\"")
	}

	r := strings.NewReplacer("{escalationId}", client.EncodePath(EscalationId))
	path := r.Replace("/organization/escalations/{escalationId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

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
	Queries        []string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListInstallationsOptions) New() *ListInstallationsOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Total": false}
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
func (srv *Organization) ListInstallations(optionalSetters ...ListInstallationsOption) (*models.AppInstallationList, error) {
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
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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
	enabledSetters       map[string]bool
}

func (options CreateInstallationOptions) New() *CreateInstallationOptions {
	options.enabledSetters = map[string]bool{"AuthorizationDetails": false}
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
func (srv *Organization) CreateInstallation(AppId string, optionalSetters ...CreateInstallationOption) (*models.AppInstallation, error) {
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
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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
func (srv *Organization) GetInstallation(InstallationId string) (*models.AppInstallation, error) {
	if InstallationId == "" {
		return nil, errors.New("Missing required parameter: \"installationId\"")
	}

	r := strings.NewReplacer("{installationId}", client.EncodePath(InstallationId))
	path := r.Replace("/organization/installations/{installationId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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
	enabledSetters       map[string]bool
}

func (options UpdateInstallationOptions) New() *UpdateInstallationOptions {
	options.enabledSetters = map[string]bool{"AuthorizationDetails": false}
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
func (srv *Organization) UpdateInstallation(InstallationId string, optionalSetters ...UpdateInstallationOption) (*models.AppInstallation, error) {
	if InstallationId == "" {
		return nil, errors.New("Missing required parameter: \"installationId\"")
	}

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
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

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
func (srv *Organization) DeleteInstallation(InstallationId string) (*interface{}, error) {
	if InstallationId == "" {
		return nil, errors.New("Missing required parameter: \"installationId\"")
	}

	r := strings.NewReplacer("{installationId}", client.EncodePath(InstallationId))
	path := r.Replace("/organization/installations/{installationId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

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
	Queries        []string
	Search         string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListMembershipsOptions) New() *ListMembershipsOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Search": false, "Total": false}
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
func (srv *Organization) ListMemberships(optionalSetters ...ListMembershipsOption) (*models.MembershipList, error) {
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
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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
	Email          string
	UserId         string
	Phone          string
	Url            string
	Name           string
	enabledSetters map[string]bool
}

func (options CreateMembershipOptions) New() *CreateMembershipOptions {
	options.enabledSetters = map[string]bool{"Email": false, "UserId": false, "Phone": false, "Url": false, "Name": false}
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
func (srv *Organization) CreateMembership(Roles []string, optionalSetters ...CreateMembershipOption) (*models.Membership, error) {
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
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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
func (srv *Organization) GetMembership(MembershipId string) (*models.Membership, error) {
	if MembershipId == "" {
		return nil, errors.New("Missing required parameter: \"membershipId\"")
	}

	r := strings.NewReplacer("{membershipId}", client.EncodePath(MembershipId))
	path := r.Replace("/organization/memberships/{membershipId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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
func (srv *Organization) UpdateMembership(MembershipId string, Roles []string) (*models.Membership, error) {
	if MembershipId == "" {
		return nil, errors.New("Missing required parameter: \"membershipId\"")
	}

	r := strings.NewReplacer("{membershipId}", client.EncodePath(MembershipId))
	path := r.Replace("/organization/memberships/{membershipId}")
	params := map[string]interface{}{}
	params["roles"] = Roles
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

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
func (srv *Organization) DeleteMembership(MembershipId string) (*interface{}, error) {
	if MembershipId == "" {
		return nil, errors.New("Missing required parameter: \"membershipId\"")
	}

	r := strings.NewReplacer("{membershipId}", client.EncodePath(MembershipId))
	path := r.Replace("/organization/memberships/{membershipId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

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
	Queries        []string
	Search         string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListProjectsOptions) New() *ListProjectsOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Search": false, "Total": false}
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
func (srv *Organization) ListProjects(optionalSetters ...ListProjectsOption) (*models.ProjectList, error) {
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
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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
	Region         string
	enabledSetters map[string]bool
}

func (options CreateProjectOptions) New() *CreateProjectOptions {
	options.enabledSetters = map[string]bool{"Region": false}
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
func (srv *Organization) CreateProject(ProjectId string, Name string, optionalSetters ...CreateProjectOption) (*models.Project, error) {
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
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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
func (srv *Organization) GetProject(ProjectId string) (*models.Project, error) {
	if ProjectId == "" {
		return nil, errors.New("Missing required parameter: \"projectId\"")
	}

	r := strings.NewReplacer("{projectId}", client.EncodePath(ProjectId))
	path := r.Replace("/organization/projects/{projectId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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
func (srv *Organization) UpdateProject(ProjectId string, Name string) (*models.Project, error) {
	if ProjectId == "" {
		return nil, errors.New("Missing required parameter: \"projectId\"")
	}

	r := strings.NewReplacer("{projectId}", client.EncodePath(ProjectId))
	path := r.Replace("/organization/projects/{projectId}")
	params := map[string]interface{}{}
	params["name"] = Name
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

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
func (srv *Organization) DeleteProject(ProjectId string) (*interface{}, error) {
	if ProjectId == "" {
		return nil, errors.New("Missing required parameter: \"projectId\"")
	}

	r := strings.NewReplacer("{projectId}", client.EncodePath(ProjectId))
	path := r.Replace("/organization/projects/{projectId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

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

type ListProjectKeysOptions struct {
	Queries        []string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListProjectKeysOptions) New() *ListProjectKeysOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Total": false}
	return &options
}

type ListProjectKeysOption func(*ListProjectKeysOptions)

func (srv *Organization) WithListProjectKeysQueries(v []string) ListProjectKeysOption {
	return func(o *ListProjectKeysOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Organization) WithListProjectKeysTotal(v bool) ListProjectKeysOption {
	return func(o *ListProjectKeysOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListProjectKeys get a list of all API keys of a project in your
// organization.
func (srv *Organization) ListProjectKeys(ProjectId string, optionalSetters ...ListProjectKeysOption) (*models.KeyList, error) {
	if ProjectId == "" {
		return nil, errors.New("Missing required parameter: \"projectId\"")
	}

	r := strings.NewReplacer("{projectId}", client.EncodePath(ProjectId))
	path := r.Replace("/organization/projects/{projectId}/keys")
	options := ListProjectKeysOptions{}.New()
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
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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

// CreateEphemeralProjectKey create a new ephemeral API key for a project in
// your organization. It's recommended to have multiple API keys with strict
// scopes for separate functions within your project.
//
// You can also create a standard API key if you need a longer-lived key
// instead.
func (srv *Organization) CreateEphemeralProjectKey(ProjectId string, Scopes []string, Duration int) (*models.EphemeralKey, error) {
	if ProjectId == "" {
		return nil, errors.New("Missing required parameter: \"projectId\"")
	}

	r := strings.NewReplacer("{projectId}", client.EncodePath(ProjectId))
	path := r.Replace("/organization/projects/{projectId}/keys/ephemeral")
	params := map[string]interface{}{}
	params["scopes"] = Scopes
	params["duration"] = Duration
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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

// GetProjectKey get a project key by its unique ID.
func (srv *Organization) GetProjectKey(ProjectId string, KeyId string) (*models.Key, error) {
	if ProjectId == "" {
		return nil, errors.New("Missing required parameter: \"projectId\"")
	}
	if KeyId == "" {
		return nil, errors.New("Missing required parameter: \"keyId\"")
	}

	r := strings.NewReplacer("{projectId}", client.EncodePath(ProjectId), "{keyId}", client.EncodePath(KeyId))
	path := r.Replace("/organization/projects/{projectId}/keys/{keyId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
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

type UpdateProjectKeyOptions struct {
	Expire         string
	Description    string
	enabledSetters map[string]bool
}

func (options UpdateProjectKeyOptions) New() *UpdateProjectKeyOptions {
	options.enabledSetters = map[string]bool{"Expire": false, "Description": false}
	return &options
}

type UpdateProjectKeyOption func(*UpdateProjectKeyOptions)

func (srv *Organization) WithUpdateProjectKeyExpire(v string) UpdateProjectKeyOption {
	return func(o *UpdateProjectKeyOptions) {
		o.Expire = v
		o.enabledSetters["Expire"] = true
	}
}
func (srv *Organization) WithUpdateProjectKeyDescription(v string) UpdateProjectKeyOption {
	return func(o *UpdateProjectKeyOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}

// UpdateProjectKey update a project key by its unique ID. Use this endpoint
// to update the name, scopes, or expiration time of an API key.
func (srv *Organization) UpdateProjectKey(ProjectId string, KeyId string, Name string, Scopes []string, optionalSetters ...UpdateProjectKeyOption) (*models.Key, error) {
	if ProjectId == "" {
		return nil, errors.New("Missing required parameter: \"projectId\"")
	}
	if KeyId == "" {
		return nil, errors.New("Missing required parameter: \"keyId\"")
	}

	r := strings.NewReplacer("{projectId}", client.EncodePath(ProjectId), "{keyId}", client.EncodePath(KeyId))
	path := r.Replace("/organization/projects/{projectId}/keys/{keyId}")
	options := UpdateProjectKeyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["scopes"] = Scopes
	if options.enabledSetters["Expire"] {
		params["expire"] = options.Expire
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

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

// DeleteProjectKey delete a project key by its unique ID. Once deleted, the
// key can no longer be used to authenticate API calls.
func (srv *Organization) DeleteProjectKey(ProjectId string, KeyId string) (*interface{}, error) {
	if ProjectId == "" {
		return nil, errors.New("Missing required parameter: \"projectId\"")
	}
	if KeyId == "" {
		return nil, errors.New("Missing required parameter: \"keyId\"")
	}

	r := strings.NewReplacer("{projectId}", client.EncodePath(ProjectId), "{keyId}", client.EncodePath(KeyId))
	path := r.Replace("/organization/projects/{projectId}/keys/{keyId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

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
