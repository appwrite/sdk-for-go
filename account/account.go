package account

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// Account service
type Account struct {
	client client.Client
}

func NewAccount(clt client.Client) *Account {
	return &Account{
		client: clt,
	}
}


// Get get the currently logged in user.
func (srv *Account) Get()(*models.User, error) {
	path := "/account"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateOptions struct {
	Name string
	enabledSetters map[string]bool
}
func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
	}
	return &options
}
type CreateOption func(*CreateOptions)
func WithCreateName(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
							
// Create use this endpoint to allow a new user to register a new account in
// your project. After the user registration completes successfully, you can
// use the
// [/account/verfication](https://appwrite.io/docs/references/cloud/client-web/account#createVerification)
// route to start verifying the user email address. To allow the new user to
// login to their new account, you need to create a new [account
// session](https://appwrite.io/docs/references/cloud/client-web/account#createEmailSession).
func (srv *Account) Create(UserId string, Email string, Password string, optionalSetters ...CreateOption)(*models.User, error) {
	path := "/account"
	options := CreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
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
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateEmail update currently logged in user account email address. After
// changing user address, the user confirmation status will get reset. A new
// confirmation email is not sent automatically however you can use the send
// confirmation email endpoint again to send the confirmation email. For
// security measures, user password is required to complete this request.
// This endpoint can also be used to convert an anonymous account to a normal
// one, by passing an email address and a new password.
func (srv *Account) UpdateEmail(Email string, Password string)(*models.User, error) {
	path := "/account/email"
	params := map[string]interface{}{}
	params["email"] = Email
	params["password"] = Password
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListIdentitiesOptions struct {
	Queries []interface{}
	enabledSetters map[string]bool
}
func (options ListIdentitiesOptions) New() *ListIdentitiesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListIdentitiesOption func(*ListIdentitiesOptions)
func WithListIdentitiesQueries(v []interface{}) ListIdentitiesOption {
	return func(o *ListIdentitiesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
	
// ListIdentities get the list of identities for the currently logged in user.
func (srv *Account) ListIdentities(optionalSetters ...ListIdentitiesOption)(*models.IdentityList, error) {
	path := "/account/identities"
	options := ListIdentitiesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.IdentityList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.IdentityList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteIdentity delete an identity by its unique ID.
func (srv *Account) DeleteIdentity(IdentityId string)(*interface{}, error) {
	r := strings.NewReplacer("{identityId}", IdentityId)
	path := r.Replace("/account/identities/{identityId}")
	params := map[string]interface{}{}
	params["identityId"] = IdentityId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreateJWT use this endpoint to create a JSON Web Token. You can use the
// resulting JWT to authenticate on behalf of the current user when working
// with the Appwrite server-side API and SDKs. The JWT secret is valid for 15
// minutes from its creation and will be invalid if the user will logout in
// that time frame.
func (srv *Account) CreateJWT()(*models.Jwt, error) {
	path := "/account/jwts"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Jwt
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Jwt)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListLogsOptions struct {
	Queries []interface{}
	enabledSetters map[string]bool
}
func (options ListLogsOptions) New() *ListLogsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListLogsOption func(*ListLogsOptions)
func WithListLogsQueries(v []interface{}) ListLogsOption {
	return func(o *ListLogsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
	
// ListLogs get the list of latest security activity logs for the currently
// logged in user. Each log returns user IP address, location and date and
// time of log.
func (srv *Account) ListLogs(optionalSetters ...ListLogsOption)(*models.LogList, error) {
	path := "/account/logs"
	options := ListLogsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.LogList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.LogList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// UpdateMFA enable or disable MFA on an account.
func (srv *Account) UpdateMFA(Mfa bool)(*models.User, error) {
	path := "/account/mfa"
	params := map[string]interface{}{}
	params["mfa"] = Mfa
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CreateMfaAuthenticator add an authenticator app to be used as an MFA
// factor. Verify the authenticator using the [verify
// authenticator](/docs/references/cloud/client-web/account#updateMfaAuthenticator)
// method.
func (srv *Account) CreateMfaAuthenticator(Type string)(*models.MfaType, error) {
	r := strings.NewReplacer("{type}", Type)
	path := r.Replace("/account/mfa/authenticators/{type}")
	params := map[string]interface{}{}
	params["type"] = Type
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.MfaType
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.MfaType)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateMfaAuthenticator verify an authenticator app after adding it using
// the [add
// authenticator](/docs/references/cloud/client-web/account#createMfaAuthenticator)
// method.
func (srv *Account) UpdateMfaAuthenticator(Type string, Otp string)(*models.User, error) {
	r := strings.NewReplacer("{type}", Type)
	path := r.Replace("/account/mfa/authenticators/{type}")
	params := map[string]interface{}{}
	params["type"] = Type
	params["otp"] = Otp
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// DeleteMfaAuthenticator delete an authenticator for a user by ID.
func (srv *Account) DeleteMfaAuthenticator(Type string, Otp string)(*interface{}, error) {
	r := strings.NewReplacer("{type}", Type)
	path := r.Replace("/account/mfa/authenticators/{type}")
	params := map[string]interface{}{}
	params["type"] = Type
	params["otp"] = Otp
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CreateMfaChallenge begin the process of MFA verification after sign-in.
// Finish the flow with
// [updateMfaChallenge](/docs/references/cloud/client-web/account#updateMfaChallenge)
// method.
func (srv *Account) CreateMfaChallenge(Factor string)(*models.MfaChallenge, error) {
	path := "/account/mfa/challenge"
	params := map[string]interface{}{}
	params["factor"] = Factor
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.MfaChallenge
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.MfaChallenge)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateMfaChallenge complete the MFA challenge by providing the one-time
// password. Finish the process of MFA verification by providing the one-time
// password. To begin the flow, use
// [createMfaChallenge](/docs/references/cloud/client-web/account#createMfaChallenge)
// method.
func (srv *Account) UpdateMfaChallenge(ChallengeId string, Otp string)(*interface{}, error) {
	path := "/account/mfa/challenge"
	params := map[string]interface{}{}
	params["challengeId"] = ChallengeId
	params["otp"] = Otp
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListMfaFactors list the factors available on the account to be used as a
// MFA challange.
func (srv *Account) ListMfaFactors()(*models.MfaFactors, error) {
	path := "/account/mfa/factors"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.MfaFactors
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.MfaFactors)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetMfaRecoveryCodes get recovery codes that can be used as backup for MFA
// flow. Before getting codes, they must be generated using
// [createMfaRecoveryCodes](/docs/references/cloud/client-web/account#createMfaRecoveryCodes)
// method. An OTP challenge is required to read recovery codes.
func (srv *Account) GetMfaRecoveryCodes()(*models.MfaRecoveryCodes, error) {
	path := "/account/mfa/recovery-codes"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.MfaRecoveryCodes
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.MfaRecoveryCodes)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreateMfaRecoveryCodes generate recovery codes as backup for MFA flow. It's
// recommended to generate and show then immediately after user successfully
// adds their authehticator. Recovery codes can be used as a MFA verification
// type in
// [createMfaChallenge](/docs/references/cloud/client-web/account#createMfaChallenge)
// method.
func (srv *Account) CreateMfaRecoveryCodes()(*models.MfaRecoveryCodes, error) {
	path := "/account/mfa/recovery-codes"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.MfaRecoveryCodes
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.MfaRecoveryCodes)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateMfaRecoveryCodes regenerate recovery codes that can be used as backup
// for MFA flow. Before regenerating codes, they must be first generated using
// [createMfaRecoveryCodes](/docs/references/cloud/client-web/account#createMfaRecoveryCodes)
// method. An OTP challenge is required to regenreate recovery codes.
func (srv *Account) UpdateMfaRecoveryCodes()(*models.MfaRecoveryCodes, error) {
	path := "/account/mfa/recovery-codes"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.MfaRecoveryCodes
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.MfaRecoveryCodes)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// UpdateName update currently logged in user account name.
func (srv *Account) UpdateName(Name string)(*models.User, error) {
	path := "/account/name"
	params := map[string]interface{}{}
	params["name"] = Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdatePasswordOptions struct {
	OldPassword string
	enabledSetters map[string]bool
}
func (options UpdatePasswordOptions) New() *UpdatePasswordOptions {
	options.enabledSetters = map[string]bool{
		"OldPassword": false,
	}
	return &options
}
type UpdatePasswordOption func(*UpdatePasswordOptions)
func WithUpdatePasswordOldPassword(v string) UpdatePasswordOption {
	return func(o *UpdatePasswordOptions) {
		o.OldPassword = v
		o.enabledSetters["OldPassword"] = true
	}
}
			
// UpdatePassword update currently logged in user password. For validation,
// user is required to pass in the new password, and the old password. For
// users created with OAuth, Team Invites and Magic URL, oldPassword is
// optional.
func (srv *Account) UpdatePassword(Password string, optionalSetters ...UpdatePasswordOption)(*models.User, error) {
	path := "/account/password"
	options := UpdatePasswordOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["password"] = Password
	if options.enabledSetters["OldPassword"] {
		params["oldPassword"] = options.OldPassword
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdatePhone update the currently logged in user's phone number. After
// updating the phone number, the phone verification status will be reset. A
// confirmation SMS is not sent automatically, however you can use the [POST
// /account/verification/phone](https://appwrite.io/docs/references/cloud/client-web/account#createPhoneVerification)
// endpoint to send a confirmation SMS.
func (srv *Account) UpdatePhone(Phone string, Password string)(*models.User, error) {
	path := "/account/phone"
	params := map[string]interface{}{}
	params["phone"] = Phone
	params["password"] = Password
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetPrefs get the preferences as a key-value object for the currently logged
// in user.
func (srv *Account) GetPrefs()(*models.Preferences, error) {
	path := "/account/prefs"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Preferences
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Preferences)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// UpdatePrefs update currently logged in user account preferences. The object
// you pass is stored as is, and replaces any previous value. The maximum
// allowed prefs size is 64kB and throws error if exceeded.
func (srv *Account) UpdatePrefs(Prefs interface{})(*models.User, error) {
	path := "/account/prefs"
	params := map[string]interface{}{}
	params["prefs"] = Prefs
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CreateRecovery sends the user an email with a temporary secret key for
// password reset. When the user clicks the confirmation link he is redirected
// back to your app password reset URL with the secret key and email address
// values attached to the URL query string. Use the query string params to
// submit a request to the [PUT
// /account/recovery](https://appwrite.io/docs/references/cloud/client-web/account#updateRecovery)
// endpoint to complete the process. The verification link sent to the user's
// email address is valid for 1 hour.
func (srv *Account) CreateRecovery(Email string, Url string)(*models.Token, error) {
	path := "/account/recovery"
	params := map[string]interface{}{}
	params["email"] = Email
	params["url"] = Url
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// UpdateRecovery use this endpoint to complete the user account password
// reset. Both the **userId** and **secret** arguments will be passed as query
// parameters to the redirect URL you have provided when sending your request
// to the [POST
// /account/recovery](https://appwrite.io/docs/references/cloud/client-web/account#createRecovery)
// endpoint.
// 
// Please note that in order to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md)
// the only valid redirect URLs are the ones from domains you have set when
// adding your platforms in the console interface.
func (srv *Account) UpdateRecovery(UserId string, Secret string, Password string)(*models.Token, error) {
	path := "/account/recovery"
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["secret"] = Secret
	params["password"] = Password
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListSessions get the list of active sessions across different devices for
// the currently logged in user.
func (srv *Account) ListSessions()(*models.SessionList, error) {
	path := "/account/sessions"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.SessionList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.SessionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteSessions delete all sessions from the user account and remove any
// sessions cookies from the end client.
func (srv *Account) DeleteSessions()(*interface{}, error) {
	path := "/account/sessions"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreateAnonymousSession use this endpoint to allow a new user to register an
// anonymous account in your project. This route will also create a new
// session for the user. To allow the new user to convert an anonymous account
// to a normal account, you need to update its [email and
// password](https://appwrite.io/docs/references/cloud/client-web/account#updateEmail)
// or create an [OAuth2
// session](https://appwrite.io/docs/references/cloud/client-web/account#CreateOAuth2Session).
func (srv *Account) CreateAnonymousSession()(*models.Session, error) {
	path := "/account/sessions/anonymous"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Session
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Session)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CreateEmailPasswordSession allow the user to login into their account by
// providing a valid email and password combination. This route will create a
// new session for the user.
// 
// A user is limited to 10 active sessions at a time by default. [Learn more
// about session
// limits](https://appwrite.io/docs/authentication-security#limits).
func (srv *Account) CreateEmailPasswordSession(Email string, Password string)(*models.Session, error) {
	path := "/account/sessions/email"
	params := map[string]interface{}{}
	params["email"] = Email
	params["password"] = Password
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Session
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Session)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateMagicURLSession use this endpoint to create a session from token.
// Provide the **userId** and **secret** parameters from the successful
// response of authentication flows initiated by token creation. For example,
// magic URL and phone login.
func (srv *Account) UpdateMagicURLSession(UserId string, Secret string)(*models.Session, error) {
	path := "/account/sessions/magic-url"
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["secret"] = Secret
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Session
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Session)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdatePhoneSession use this endpoint to create a session from token.
// Provide the **userId** and **secret** parameters from the successful
// response of authentication flows initiated by token creation. For example,
// magic URL and phone login.
func (srv *Account) UpdatePhoneSession(UserId string, Secret string)(*models.Session, error) {
	path := "/account/sessions/phone"
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["secret"] = Secret
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Session
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Session)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CreateSession use this endpoint to create a session from token. Provide the
// **userId** and **secret** parameters from the successful response of
// authentication flows initiated by token creation. For example, magic URL
// and phone login.
func (srv *Account) CreateSession(UserId string, Secret string)(*models.Session, error) {
	path := "/account/sessions/token"
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["secret"] = Secret
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Session
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Session)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetSession use this endpoint to get a logged in user's session using a
// Session ID. Inputting 'current' will return the current session being used.
func (srv *Account) GetSession(SessionId string)(*models.Session, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	path := r.Replace("/account/sessions/{sessionId}")
	params := map[string]interface{}{}
	params["sessionId"] = SessionId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Session
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Session)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// UpdateSession use this endpoint to extend a session's length. Extending a
// session is useful when session expiry is short. If the session was created
// using an OAuth provider, this endpoint refreshes the access token from the
// provider.
func (srv *Account) UpdateSession(SessionId string)(*models.Session, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	path := r.Replace("/account/sessions/{sessionId}")
	params := map[string]interface{}{}
	params["sessionId"] = SessionId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Session
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Session)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteSession logout the user. Use 'current' as the session ID to logout on
// this device, use a session ID to logout on another device. If you're
// looking to logout the user on all devices, use [Delete
// Sessions](https://appwrite.io/docs/references/cloud/client-web/account#deleteSessions)
// instead.
func (srv *Account) DeleteSession(SessionId string)(*interface{}, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	path := r.Replace("/account/sessions/{sessionId}")
	params := map[string]interface{}{}
	params["sessionId"] = SessionId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateStatus block the currently logged in user account. Behind the scene,
// the user record is not deleted but permanently blocked from any access. To
// completely delete a user, use the Users API instead.
func (srv *Account) UpdateStatus()(*models.User, error) {
	path := "/account/status"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.User
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateEmailTokenOptions struct {
	Phrase bool
	enabledSetters map[string]bool
}
func (options CreateEmailTokenOptions) New() *CreateEmailTokenOptions {
	options.enabledSetters = map[string]bool{
		"Phrase": false,
	}
	return &options
}
type CreateEmailTokenOption func(*CreateEmailTokenOptions)
func WithCreateEmailTokenPhrase(v bool) CreateEmailTokenOption {
	return func(o *CreateEmailTokenOptions) {
		o.Phrase = v
		o.enabledSetters["Phrase"] = true
	}
}
					
// CreateEmailToken sends the user an email with a secret key for creating a
// session. If the provided user ID has not be registered, a new user will be
// created. Use the returned user ID and secret and submit a request to the
// [POST
// /v1/account/sessions/token](https://appwrite.io/docs/references/cloud/client-web/account#createSession)
// endpoint to complete the login process. The secret sent to the user's email
// is valid for 15 minutes.
// 
// A user is limited to 10 active sessions at a time by default. [Learn more
// about session
// limits](https://appwrite.io/docs/authentication-security#limits).
func (srv *Account) CreateEmailToken(UserId string, Email string, optionalSetters ...CreateEmailTokenOption)(*models.Token, error) {
	path := "/account/tokens/email"
	options := CreateEmailTokenOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	if options.enabledSetters["Phrase"] {
		params["phrase"] = options.Phrase
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateMagicURLTokenOptions struct {
	Url string
	Phrase bool
	enabledSetters map[string]bool
}
func (options CreateMagicURLTokenOptions) New() *CreateMagicURLTokenOptions {
	options.enabledSetters = map[string]bool{
		"Url": false,
		"Phrase": false,
	}
	return &options
}
type CreateMagicURLTokenOption func(*CreateMagicURLTokenOptions)
func WithCreateMagicURLTokenUrl(v string) CreateMagicURLTokenOption {
	return func(o *CreateMagicURLTokenOptions) {
		o.Url = v
		o.enabledSetters["Url"] = true
	}
}
func WithCreateMagicURLTokenPhrase(v bool) CreateMagicURLTokenOption {
	return func(o *CreateMagicURLTokenOptions) {
		o.Phrase = v
		o.enabledSetters["Phrase"] = true
	}
}
					
// CreateMagicURLToken sends the user an email with a secret key for creating
// a session. If the provided user ID has not been registered, a new user will
// be created. When the user clicks the link in the email, the user is
// redirected back to the URL you provided with the secret key and userId
// values attached to the URL query string. Use the query string parameters to
// submit a request to the [POST
// /v1/account/sessions/token](https://appwrite.io/docs/references/cloud/client-web/account#createSession)
// endpoint to complete the login process. The link sent to the user's email
// address is valid for 1 hour. If you are on a mobile device you can leave
// the URL parameter empty, so that the login completion will be handled by
// your Appwrite instance by default.
// 
// A user is limited to 10 active sessions at a time by default. [Learn more
// about session
// limits](https://appwrite.io/docs/authentication-security#limits).
func (srv *Account) CreateMagicURLToken(UserId string, Email string, optionalSetters ...CreateMagicURLTokenOption)(*models.Token, error) {
	path := "/account/tokens/magic-url"
	options := CreateMagicURLTokenOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	if options.enabledSetters["Url"] {
		params["url"] = options.Url
	}
	if options.enabledSetters["Phrase"] {
		params["phrase"] = options.Phrase
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateOAuth2TokenOptions struct {
	Success string
	Failure string
	Scopes []interface{}
	enabledSetters map[string]bool
}
func (options CreateOAuth2TokenOptions) New() *CreateOAuth2TokenOptions {
	options.enabledSetters = map[string]bool{
		"Success": false,
		"Failure": false,
		"Scopes": false,
	}
	return &options
}
type CreateOAuth2TokenOption func(*CreateOAuth2TokenOptions)
func WithCreateOAuth2TokenSuccess(v string) CreateOAuth2TokenOption {
	return func(o *CreateOAuth2TokenOptions) {
		o.Success = v
		o.enabledSetters["Success"] = true
	}
}
func WithCreateOAuth2TokenFailure(v string) CreateOAuth2TokenOption {
	return func(o *CreateOAuth2TokenOptions) {
		o.Failure = v
		o.enabledSetters["Failure"] = true
	}
}
func WithCreateOAuth2TokenScopes(v []interface{}) CreateOAuth2TokenOption {
	return func(o *CreateOAuth2TokenOptions) {
		o.Scopes = v
		o.enabledSetters["Scopes"] = true
	}
}
			
// CreateOAuth2Token allow the user to login to their account using the OAuth2
// provider of their choice. Each OAuth2 provider should be enabled from the
// Appwrite console first. Use the success and failure arguments to provide a
// redirect URL's back to your app when login is completed.
// 
// If authentication succeeds, `userId` and `secret` of a token will be
// appended to the success URL as query parameters. These can be used to
// create a new session using the [Create
// session](https://appwrite.io/docs/references/cloud/client-web/account#createSession)
// endpoint.
// 
// A user is limited to 10 active sessions at a time by default. [Learn more
// about session
// limits](https://appwrite.io/docs/authentication-security#limits).
func (srv *Account) CreateOAuth2Token(Provider string, optionalSetters ...CreateOAuth2TokenOption)(*bool, error) {
	r := strings.NewReplacer("{provider}", Provider)
	path := r.Replace("/account/tokens/oauth2/{provider}")
	options := CreateOAuth2TokenOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["provider"] = Provider
	if options.enabledSetters["Success"] {
		params["success"] = options.Success
	}
	if options.enabledSetters["Failure"] {
		params["failure"] = options.Failure
	}
	if options.enabledSetters["Scopes"] {
		params["scopes"] = options.Scopes
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed bool
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(bool)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CreatePhoneToken sends the user an SMS with a secret key for creating a
// session. If the provided user ID has not be registered, a new user will be
// created. Use the returned user ID and secret and submit a request to the
// [POST
// /v1/account/sessions/token](https://appwrite.io/docs/references/cloud/client-web/account#createSession)
// endpoint to complete the login process. The secret sent to the user's phone
// is valid for 15 minutes.
// 
// A user is limited to 10 active sessions at a time by default. [Learn more
// about session
// limits](https://appwrite.io/docs/authentication-security#limits).
func (srv *Account) CreatePhoneToken(UserId string, Phone string)(*models.Token, error) {
	path := "/account/tokens/phone"
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["phone"] = Phone
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CreateVerification use this endpoint to send a verification message to your
// user email address to confirm they are the valid owners of that address.
// Both the **userId** and **secret** arguments will be passed as query
// parameters to the URL you have provided to be attached to the verification
// email. The provided URL should redirect the user back to your app and allow
// you to complete the verification process by verifying both the **userId**
// and **secret** parameters. Learn more about how to [complete the
// verification
// process](https://appwrite.io/docs/references/cloud/client-web/account#updateVerification).
// The verification link sent to the user's email address is valid for 7 days.
// 
// Please note that in order to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md),
// the only valid redirect URLs are the ones from domains you have set when
// adding your platforms in the console interface.
func (srv *Account) CreateVerification(Url string)(*models.Token, error) {
	path := "/account/verification"
	params := map[string]interface{}{}
	params["url"] = Url
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateVerification use this endpoint to complete the user email
// verification process. Use both the **userId** and **secret** parameters
// that were attached to your app URL to verify the user email ownership. If
// confirmed this route will return a 200 status code.
func (srv *Account) UpdateVerification(UserId string, Secret string)(*models.Token, error) {
	path := "/account/verification"
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["secret"] = Secret
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreatePhoneVerification use this endpoint to send a verification SMS to the
// currently logged in user. This endpoint is meant for use after updating a
// user's phone number using the
// [accountUpdatePhone](https://appwrite.io/docs/references/cloud/client-web/account#updatePhone)
// endpoint. Learn more about how to [complete the verification
// process](https://appwrite.io/docs/references/cloud/client-web/account#updatePhoneVerification).
// The verification code sent to the user's phone number is valid for 15
// minutes.
func (srv *Account) CreatePhoneVerification()(*models.Token, error) {
	path := "/account/verification/phone"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdatePhoneVerification use this endpoint to complete the user phone
// verification process. Use the **userId** and **secret** that were sent to
// your user's phone number to verify the user email ownership. If confirmed
// this route will return a 200 status code.
func (srv *Account) UpdatePhoneVerification(UserId string, Secret string)(*models.Token, error) {
	path := "/account/verification/phone"
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["secret"] = Secret
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Token
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
