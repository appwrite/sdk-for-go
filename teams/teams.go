package teams

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// Teams service
type Teams struct {
	client client.Client
}

func New(clt client.Client) *Teams {
	return &Teams{
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
func (srv *Teams) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Teams) WithListSearch(v string) ListOption {
	return func(o *ListOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
	
// List get a list of all the teams in which the current user is a member. You
// can use the parameters to filter your results.
func (srv *Teams) List(optionalSetters ...ListOption)(*models.TeamList, error) {
	path := "/teams"
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
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.TeamList{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.TeamList
	parsed, ok := resp.Result.(models.TeamList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateOptions struct {
	Roles []string
	enabledSetters map[string]bool
}
func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{
		"Roles": false,
	}
	return &options
}
type CreateOption func(*CreateOptions)
func (srv *Teams) WithCreateRoles(v []string) CreateOption {
	return func(o *CreateOptions) {
		o.Roles = v
		o.enabledSetters["Roles"] = true
	}
}
					
// Create create a new team. The user who creates the team will automatically
// be assigned as the owner of the team. Only the users with the owner role
// can invite new members, add new owners and delete or update the team.
func (srv *Teams) Create(TeamId string, Name string, optionalSetters ...CreateOption)(*models.Team, error) {
	path := "/teams"
	options := CreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["name"] = Name
	if options.enabledSetters["Roles"] {
		params["roles"] = options.Roles
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.Team{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Team
	parsed, ok := resp.Result.(models.Team)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Get get a team by its ID. All team members have read access for this
// resource.
func (srv *Teams) Get(TeamId string)(*models.Team, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}")
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.Team{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Team
	parsed, ok := resp.Result.(models.Team)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateName update the team's name by its unique ID.
func (srv *Teams) UpdateName(TeamId string, Name string)(*models.Team, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}")
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["name"] = Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.Team{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Team
	parsed, ok := resp.Result.(models.Team)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Delete delete a team using its ID. Only team members with the owner role
// can delete the team.
func (srv *Teams) Delete(TeamId string)(*interface{}, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}")
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		var parsed interface{}

		err = json.Unmarshal(bytesData, &parsed)
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
	enabledSetters map[string]bool
}
func (options ListMembershipsOptions) New() *ListMembershipsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
	}
	return &options
}
type ListMembershipsOption func(*ListMembershipsOptions)
func (srv *Teams) WithListMembershipsQueries(v []string) ListMembershipsOption {
	return func(o *ListMembershipsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Teams) WithListMembershipsSearch(v string) ListMembershipsOption {
	return func(o *ListMembershipsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
			
// ListMemberships use this endpoint to list a team's members using the team's
// ID. All team members have read access to this endpoint.
func (srv *Teams) ListMemberships(TeamId string, optionalSetters ...ListMembershipsOption)(*models.MembershipList, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}/memberships")
	options := ListMembershipsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.MembershipList{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
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
func (srv *Teams) WithCreateMembershipEmail(v string) CreateMembershipOption {
	return func(o *CreateMembershipOptions) {
		o.Email = v
		o.enabledSetters["Email"] = true
	}
}
func (srv *Teams) WithCreateMembershipUserId(v string) CreateMembershipOption {
	return func(o *CreateMembershipOptions) {
		o.UserId = v
		o.enabledSetters["UserId"] = true
	}
}
func (srv *Teams) WithCreateMembershipPhone(v string) CreateMembershipOption {
	return func(o *CreateMembershipOptions) {
		o.Phone = v
		o.enabledSetters["Phone"] = true
	}
}
func (srv *Teams) WithCreateMembershipUrl(v string) CreateMembershipOption {
	return func(o *CreateMembershipOptions) {
		o.Url = v
		o.enabledSetters["Url"] = true
	}
}
func (srv *Teams) WithCreateMembershipName(v string) CreateMembershipOption {
	return func(o *CreateMembershipOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
					
// CreateMembership invite a new member to join your team. Provide an ID for
// existing users, or invite unregistered users using an email or phone
// number. If initiated from a Client SDK, Appwrite will send an email or sms
// with a link to join the team to the invited user, and an account will be
// created for them if one doesn't exist. If initiated from a Server SDK, the
// new member will be added automatically to the team.
// 
// You only need to provide one of a user ID, email, or phone number. Appwrite
// will prioritize accepting the user ID > email > phone number if you provide
// more than one of these parameters.
// 
// Use the `url` parameter to redirect the user from the invitation email to
// your app. After the user is redirected, use the [Update Team Membership
// Status](https://appwrite.io/docs/references/cloud/client-web/teams#updateMembershipStatus)
// endpoint to allow the user to accept the invitation to the team.
// 
// Please note that to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md)
// Appwrite will accept the only redirect URLs under the domains you have
// added as a platform on the Appwrite Console.
func (srv *Teams) CreateMembership(TeamId string, Roles []string, optionalSetters ...CreateMembershipOption)(*models.Membership, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}/memberships")
	options := CreateMembershipOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["roles"] = Roles
	if options.enabledSetters["Email"] {
		params["email"] = options.Email
	}
	if options.enabledSetters["UserId"] {
		params["userId"] = options.UserId
	}
	if options.enabledSetters["Phone"] {
		params["phone"] = options.Phone
	}
	if options.enabledSetters["Url"] {
		params["url"] = options.Url
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.Membership{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
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
			
// GetMembership get a team member by the membership unique id. All team
// members have read access for this resource.
func (srv *Teams) GetMembership(TeamId string, MembershipId string)(*models.Membership, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	path := r.Replace("/teams/{teamId}/memberships/{membershipId}")
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["membershipId"] = MembershipId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.Membership{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
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
					
// UpdateMembership modify the roles of a team member. Only team members with
// the owner role have access to this endpoint. Learn more about [roles and
// permissions](https://appwrite.io/docs/permissions).
func (srv *Teams) UpdateMembership(TeamId string, MembershipId string, Roles []string)(*models.Membership, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	path := r.Replace("/teams/{teamId}/memberships/{membershipId}")
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["membershipId"] = MembershipId
	params["roles"] = Roles
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.Membership{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
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
			
// DeleteMembership this endpoint allows a user to leave a team or for a team
// owner to delete the membership of any other team member. You can also use
// this endpoint to delete a user membership even if it is not accepted.
func (srv *Teams) DeleteMembership(TeamId string, MembershipId string)(*interface{}, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	path := r.Replace("/teams/{teamId}/memberships/{membershipId}")
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["membershipId"] = MembershipId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		var parsed interface{}

		err = json.Unmarshal(bytesData, &parsed)
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
							
// UpdateMembershipStatus use this endpoint to allow a user to accept an
// invitation to join a team after being redirected back to your app from the
// invitation email received by the user.
// 
// If the request is successful, a session for the user is automatically
// created.
func (srv *Teams) UpdateMembershipStatus(TeamId string, MembershipId string, UserId string, Secret string)(*models.Membership, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	path := r.Replace("/teams/{teamId}/memberships/{membershipId}/status")
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["membershipId"] = MembershipId
	params["userId"] = UserId
	params["secret"] = Secret
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.Membership{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
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
	
// GetPrefs get the team's shared preferences by its unique ID. If a
// preference doesn't need to be shared by all team members, prefer storing
// them in [user
// preferences](https://appwrite.io/docs/references/cloud/client-web/account#getPrefs).
func (srv *Teams) GetPrefs(TeamId string)(*models.Preferences, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}/prefs")
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.Preferences{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Preferences
	parsed, ok := resp.Result.(models.Preferences)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdatePrefs update the team's preferences by its unique ID. The object you
// pass is stored as is and replaces any previous value. The maximum allowed
// prefs size is 64kB and throws an error if exceeded.
func (srv *Teams) UpdatePrefs(TeamId string, Prefs interface{})(*models.Preferences, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}/prefs")
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["prefs"] = Prefs
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.Preferences{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Preferences
	parsed, ok := resp.Result.(models.Preferences)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
