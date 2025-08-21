package users

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// Users service
type Users struct {
	client client.Client
}

func New(clt client.Client) *Users {
	return &Users{
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
func (srv *Users) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Users) WithListSearch(v string) ListOption {
	return func(o *ListOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
	
// List get a list of all the project's users. You can use the query params to
// filter your results.
func (srv *Users) List(optionalSetters ...ListOption)(*models.UserList, error) {
	path := "/users"
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
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.UserList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.UserList
	parsed, ok := resp.Result.(models.UserList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateOptions struct {
	Email string
	Phone string
	Password string
	Name string
	enabledSetters map[string]bool
}
func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{
		"Email": false,
		"Phone": false,
		"Password": false,
		"Name": false,
	}
	return &options
}
type CreateOption func(*CreateOptions)
func (srv *Users) WithCreateEmail(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Email = v
		o.enabledSetters["Email"] = true
	}
}
func (srv *Users) WithCreatePhone(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Phone = v
		o.enabledSetters["Phone"] = true
	}
}
func (srv *Users) WithCreatePassword(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Password = v
		o.enabledSetters["Password"] = true
	}
}
func (srv *Users) WithCreateName(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
			
// Create create a new user.
func (srv *Users) Create(UserId string, optionalSetters ...CreateOption)(*models.User, error) {
	path := "/users"
	options := CreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	if options.enabledSetters["Email"] {
		params["email"] = options.Email
	}
	if options.enabledSetters["Phone"] {
		params["phone"] = options.Phone
	}
	if options.enabledSetters["Password"] {
		params["password"] = options.Password
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
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateArgon2UserOptions struct {
	Name string
	enabledSetters map[string]bool
}
func (options CreateArgon2UserOptions) New() *CreateArgon2UserOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
	}
	return &options
}
type CreateArgon2UserOption func(*CreateArgon2UserOptions)
func (srv *Users) WithCreateArgon2UserName(v string) CreateArgon2UserOption {
	return func(o *CreateArgon2UserOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
							
// CreateArgon2User create a new user. Password provided must be hashed with
// the [Argon2](https://en.wikipedia.org/wiki/Argon2) algorithm. Use the [POST
// /users](https://appwrite.io/docs/server/users#usersCreate) endpoint to
// create users with a plain text password.
func (srv *Users) CreateArgon2User(UserId string, Email string, Password string, optionalSetters ...CreateArgon2UserOption)(*models.User, error) {
	path := "/users/argon2"
	options := CreateArgon2UserOptions{}.New()
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
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateBcryptUserOptions struct {
	Name string
	enabledSetters map[string]bool
}
func (options CreateBcryptUserOptions) New() *CreateBcryptUserOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
	}
	return &options
}
type CreateBcryptUserOption func(*CreateBcryptUserOptions)
func (srv *Users) WithCreateBcryptUserName(v string) CreateBcryptUserOption {
	return func(o *CreateBcryptUserOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
							
// CreateBcryptUser create a new user. Password provided must be hashed with
// the [Bcrypt](https://en.wikipedia.org/wiki/Bcrypt) algorithm. Use the [POST
// /users](https://appwrite.io/docs/server/users#usersCreate) endpoint to
// create users with a plain text password.
func (srv *Users) CreateBcryptUser(UserId string, Email string, Password string, optionalSetters ...CreateBcryptUserOption)(*models.User, error) {
	path := "/users/bcrypt"
	options := CreateBcryptUserOptions{}.New()
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
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListIdentitiesOptions struct {
	Queries []string
	Search string
	enabledSetters map[string]bool
}
func (options ListIdentitiesOptions) New() *ListIdentitiesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
	}
	return &options
}
type ListIdentitiesOption func(*ListIdentitiesOptions)
func (srv *Users) WithListIdentitiesQueries(v []string) ListIdentitiesOption {
	return func(o *ListIdentitiesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Users) WithListIdentitiesSearch(v string) ListIdentitiesOption {
	return func(o *ListIdentitiesOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
	
// ListIdentities get identities for all users.
func (srv *Users) ListIdentities(optionalSetters ...ListIdentitiesOption)(*models.IdentityList, error) {
	path := "/users/identities"
	options := ListIdentitiesOptions{}.New()
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
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.IdentityList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.IdentityList
	parsed, ok := resp.Result.(models.IdentityList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteIdentity delete an identity by its unique ID.
func (srv *Users) DeleteIdentity(IdentityId string)(*interface{}, error) {
	r := strings.NewReplacer("{identityId}", IdentityId)
	path := r.Replace("/users/identities/{identityId}")
	params := map[string]interface{}{}
	params["identityId"] = IdentityId
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
type CreateMD5UserOptions struct {
	Name string
	enabledSetters map[string]bool
}
func (options CreateMD5UserOptions) New() *CreateMD5UserOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
	}
	return &options
}
type CreateMD5UserOption func(*CreateMD5UserOptions)
func (srv *Users) WithCreateMD5UserName(v string) CreateMD5UserOption {
	return func(o *CreateMD5UserOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
							
// CreateMD5User create a new user. Password provided must be hashed with the
// [MD5](https://en.wikipedia.org/wiki/MD5) algorithm. Use the [POST
// /users](https://appwrite.io/docs/server/users#usersCreate) endpoint to
// create users with a plain text password.
func (srv *Users) CreateMD5User(UserId string, Email string, Password string, optionalSetters ...CreateMD5UserOption)(*models.User, error) {
	path := "/users/md5"
	options := CreateMD5UserOptions{}.New()
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
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreatePHPassUserOptions struct {
	Name string
	enabledSetters map[string]bool
}
func (options CreatePHPassUserOptions) New() *CreatePHPassUserOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
	}
	return &options
}
type CreatePHPassUserOption func(*CreatePHPassUserOptions)
func (srv *Users) WithCreatePHPassUserName(v string) CreatePHPassUserOption {
	return func(o *CreatePHPassUserOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
							
// CreatePHPassUser create a new user. Password provided must be hashed with
// the [PHPass](https://www.openwall.com/phpass/) algorithm. Use the [POST
// /users](https://appwrite.io/docs/server/users#usersCreate) endpoint to
// create users with a plain text password.
func (srv *Users) CreatePHPassUser(UserId string, Email string, Password string, optionalSetters ...CreatePHPassUserOption)(*models.User, error) {
	path := "/users/phpass"
	options := CreatePHPassUserOptions{}.New()
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
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateScryptUserOptions struct {
	Name string
	enabledSetters map[string]bool
}
func (options CreateScryptUserOptions) New() *CreateScryptUserOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
	}
	return &options
}
type CreateScryptUserOption func(*CreateScryptUserOptions)
func (srv *Users) WithCreateScryptUserName(v string) CreateScryptUserOption {
	return func(o *CreateScryptUserOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
																	
// CreateScryptUser create a new user. Password provided must be hashed with
// the [Scrypt](https://github.com/Tarsnap/scrypt) algorithm. Use the [POST
// /users](https://appwrite.io/docs/server/users#usersCreate) endpoint to
// create users with a plain text password.
func (srv *Users) CreateScryptUser(UserId string, Email string, Password string, PasswordSalt string, PasswordCpu int, PasswordMemory int, PasswordParallel int, PasswordLength int, optionalSetters ...CreateScryptUserOption)(*models.User, error) {
	path := "/users/scrypt"
	options := CreateScryptUserOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	params["passwordSalt"] = PasswordSalt
	params["passwordCpu"] = PasswordCpu
	params["passwordMemory"] = PasswordMemory
	params["passwordParallel"] = PasswordParallel
	params["passwordLength"] = PasswordLength
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
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateScryptModifiedUserOptions struct {
	Name string
	enabledSetters map[string]bool
}
func (options CreateScryptModifiedUserOptions) New() *CreateScryptModifiedUserOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
	}
	return &options
}
type CreateScryptModifiedUserOption func(*CreateScryptModifiedUserOptions)
func (srv *Users) WithCreateScryptModifiedUserName(v string) CreateScryptModifiedUserOption {
	return func(o *CreateScryptModifiedUserOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
													
// CreateScryptModifiedUser create a new user. Password provided must be
// hashed with the [Scrypt
// Modified](https://gist.github.com/Meldiron/eecf84a0225eccb5a378d45bb27462cc)
// algorithm. Use the [POST
// /users](https://appwrite.io/docs/server/users#usersCreate) endpoint to
// create users with a plain text password.
func (srv *Users) CreateScryptModifiedUser(UserId string, Email string, Password string, PasswordSalt string, PasswordSaltSeparator string, PasswordSignerKey string, optionalSetters ...CreateScryptModifiedUserOption)(*models.User, error) {
	path := "/users/scrypt-modified"
	options := CreateScryptModifiedUserOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	params["passwordSalt"] = PasswordSalt
	params["passwordSaltSeparator"] = PasswordSaltSeparator
	params["passwordSignerKey"] = PasswordSignerKey
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
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateSHAUserOptions struct {
	PasswordVersion string
	Name string
	enabledSetters map[string]bool
}
func (options CreateSHAUserOptions) New() *CreateSHAUserOptions {
	options.enabledSetters = map[string]bool{
		"PasswordVersion": false,
		"Name": false,
	}
	return &options
}
type CreateSHAUserOption func(*CreateSHAUserOptions)
func (srv *Users) WithCreateSHAUserPasswordVersion(v string) CreateSHAUserOption {
	return func(o *CreateSHAUserOptions) {
		o.PasswordVersion = v
		o.enabledSetters["PasswordVersion"] = true
	}
}
func (srv *Users) WithCreateSHAUserName(v string) CreateSHAUserOption {
	return func(o *CreateSHAUserOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
							
// CreateSHAUser create a new user. Password provided must be hashed with the
// [SHA](https://en.wikipedia.org/wiki/Secure_Hash_Algorithm) algorithm. Use
// the [POST /users](https://appwrite.io/docs/server/users#usersCreate)
// endpoint to create users with a plain text password.
func (srv *Users) CreateSHAUser(UserId string, Email string, Password string, optionalSetters ...CreateSHAUserOption)(*models.User, error) {
	path := "/users/sha"
	options := CreateSHAUserOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	params["password"] = Password
	if options.enabledSetters["PasswordVersion"] {
		params["passwordVersion"] = options.PasswordVersion
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
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Get get a user by its unique ID.
func (srv *Users) Get(UserId string)(*models.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}")
	params := map[string]interface{}{}
	params["userId"] = UserId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Delete delete a user by its unique ID, thereby releasing it's ID. Since ID
// is released and can be reused, all user-related resources like documents or
// storage files should be deleted before user deletion. If you want to keep
// ID reserved, use the
// [updateStatus](https://appwrite.io/docs/server/users#usersUpdateStatus)
// endpoint instead.
func (srv *Users) Delete(UserId string)(*interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}")
	params := map[string]interface{}{}
	params["userId"] = UserId
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
			
// UpdateEmail update the user email by its unique ID.
func (srv *Users) UpdateEmail(UserId string, Email string)(*models.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/email")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["email"] = Email
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateJWTOptions struct {
	SessionId string
	Duration int
	enabledSetters map[string]bool
}
func (options CreateJWTOptions) New() *CreateJWTOptions {
	options.enabledSetters = map[string]bool{
		"SessionId": false,
		"Duration": false,
	}
	return &options
}
type CreateJWTOption func(*CreateJWTOptions)
func (srv *Users) WithCreateJWTSessionId(v string) CreateJWTOption {
	return func(o *CreateJWTOptions) {
		o.SessionId = v
		o.enabledSetters["SessionId"] = true
	}
}
func (srv *Users) WithCreateJWTDuration(v int) CreateJWTOption {
	return func(o *CreateJWTOptions) {
		o.Duration = v
		o.enabledSetters["Duration"] = true
	}
}
			
// CreateJWT use this endpoint to create a JSON Web Token for user by its
// unique ID. You can use the resulting JWT to authenticate on behalf of the
// user. The JWT secret will become invalid if the session it uses gets
// deleted.
func (srv *Users) CreateJWT(UserId string, optionalSetters ...CreateJWTOption)(*models.Jwt, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/jwts")
	options := CreateJWTOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	if options.enabledSetters["SessionId"] {
		params["sessionId"] = options.SessionId
	}
	if options.enabledSetters["Duration"] {
		params["duration"] = options.Duration
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

		parsed := models.Jwt{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Jwt
	parsed, ok := resp.Result.(models.Jwt)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateLabels update the user labels by its unique ID.
// 
// Labels can be used to grant access to resources. While teams are a way for
// user's to share access to a resource, labels can be defined by the
// developer to grant access without an invitation. See the [Permissions
// docs](https://appwrite.io/docs/permissions) for more info.
func (srv *Users) UpdateLabels(UserId string, Labels []string)(*models.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/labels")
	params := map[string]interface{}{}
	params["userId"] = UserId
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

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListLogsOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListLogsOptions) New() *ListLogsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListLogsOption func(*ListLogsOptions)
func (srv *Users) WithListLogsQueries(v []string) ListLogsOption {
	return func(o *ListLogsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
			
// ListLogs get the user activity logs list by its unique ID.
func (srv *Users) ListLogs(UserId string, optionalSetters ...ListLogsOption)(*models.LogList, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/logs")
	options := ListLogsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.LogList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LogList
	parsed, ok := resp.Result.(models.LogList)
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
func (srv *Users) WithListMembershipsQueries(v []string) ListMembershipsOption {
	return func(o *ListMembershipsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Users) WithListMembershipsSearch(v string) ListMembershipsOption {
	return func(o *ListMembershipsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
			
// ListMemberships get the user membership list by its unique ID.
func (srv *Users) ListMemberships(UserId string, optionalSetters ...ListMembershipsOption)(*models.MembershipList, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/memberships")
	options := ListMembershipsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

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
			
// UpdateMfa enable or disable MFA on a user account.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `UpdateMFA` instead.
func (srv *Users) UpdateMfa(UserId string, Mfa bool)(*models.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/mfa")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["mfa"] = Mfa
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateMFA enable or disable MFA on a user account.
func (srv *Users) UpdateMFA(UserId string, Mfa bool)(*models.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/mfa")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["mfa"] = Mfa
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// DeleteMfaAuthenticator delete an authenticator app.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `DeleteMFAAuthenticator` instead.
func (srv *Users) DeleteMfaAuthenticator(UserId string, Type string)(*interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId, "{type}", Type)
	path := r.Replace("/users/{userId}/mfa/authenticators/{type}")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["type"] = Type
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
			
// DeleteMFAAuthenticator delete an authenticator app.
func (srv *Users) DeleteMFAAuthenticator(UserId string, Type string)(*interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId, "{type}", Type)
	path := r.Replace("/users/{userId}/mfa/authenticators/{type}")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["type"] = Type
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
	
// ListMfaFactors list the factors available on the account to be used as a
// MFA challange.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `ListMFAFactors` instead.
func (srv *Users) ListMfaFactors(UserId string)(*models.MfaFactors, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/mfa/factors")
	params := map[string]interface{}{}
	params["userId"] = UserId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MfaFactors{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MfaFactors
	parsed, ok := resp.Result.(models.MfaFactors)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ListMFAFactors list the factors available on the account to be used as a
// MFA challange.
func (srv *Users) ListMFAFactors(UserId string)(*models.MfaFactors, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/mfa/factors")
	params := map[string]interface{}{}
	params["userId"] = UserId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MfaFactors{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MfaFactors
	parsed, ok := resp.Result.(models.MfaFactors)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetMfaRecoveryCodes get recovery codes that can be used as backup for MFA
// flow by User ID. Before getting codes, they must be generated using
// [createMfaRecoveryCodes](/docs/references/cloud/client-web/account#createMfaRecoveryCodes)
// method.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `GetMFARecoveryCodes` instead.
func (srv *Users) GetMfaRecoveryCodes(UserId string)(*models.MfaRecoveryCodes, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/mfa/recovery-codes")
	params := map[string]interface{}{}
	params["userId"] = UserId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MfaRecoveryCodes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MfaRecoveryCodes
	parsed, ok := resp.Result.(models.MfaRecoveryCodes)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetMFARecoveryCodes get recovery codes that can be used as backup for MFA
// flow by User ID. Before getting codes, they must be generated using
// [createMfaRecoveryCodes](/docs/references/cloud/client-web/account#createMfaRecoveryCodes)
// method.
func (srv *Users) GetMFARecoveryCodes(UserId string)(*models.MfaRecoveryCodes, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/mfa/recovery-codes")
	params := map[string]interface{}{}
	params["userId"] = UserId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MfaRecoveryCodes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MfaRecoveryCodes
	parsed, ok := resp.Result.(models.MfaRecoveryCodes)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// UpdateMfaRecoveryCodes regenerate recovery codes that can be used as backup
// for MFA flow by User ID. Before regenerating codes, they must be first
// generated using
// [createMfaRecoveryCodes](/docs/references/cloud/client-web/account#createMfaRecoveryCodes)
// method.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `UpdateMFARecoveryCodes` instead.
func (srv *Users) UpdateMfaRecoveryCodes(UserId string)(*models.MfaRecoveryCodes, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/mfa/recovery-codes")
	params := map[string]interface{}{}
	params["userId"] = UserId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MfaRecoveryCodes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MfaRecoveryCodes
	parsed, ok := resp.Result.(models.MfaRecoveryCodes)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// UpdateMFARecoveryCodes regenerate recovery codes that can be used as backup
// for MFA flow by User ID. Before regenerating codes, they must be first
// generated using
// [createMfaRecoveryCodes](/docs/references/cloud/client-web/account#createMfaRecoveryCodes)
// method.
func (srv *Users) UpdateMFARecoveryCodes(UserId string)(*models.MfaRecoveryCodes, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/mfa/recovery-codes")
	params := map[string]interface{}{}
	params["userId"] = UserId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MfaRecoveryCodes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MfaRecoveryCodes
	parsed, ok := resp.Result.(models.MfaRecoveryCodes)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CreateMfaRecoveryCodes generate recovery codes used as backup for MFA flow
// for User ID. Recovery codes can be used as a MFA verification type in
// [createMfaChallenge](/docs/references/cloud/client-web/account#createMfaChallenge)
// method by client SDK.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `CreateMFARecoveryCodes` instead.
func (srv *Users) CreateMfaRecoveryCodes(UserId string)(*models.MfaRecoveryCodes, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/mfa/recovery-codes")
	params := map[string]interface{}{}
	params["userId"] = UserId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MfaRecoveryCodes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MfaRecoveryCodes
	parsed, ok := resp.Result.(models.MfaRecoveryCodes)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CreateMFARecoveryCodes generate recovery codes used as backup for MFA flow
// for User ID. Recovery codes can be used as a MFA verification type in
// [createMfaChallenge](/docs/references/cloud/client-web/account#createMfaChallenge)
// method by client SDK.
func (srv *Users) CreateMFARecoveryCodes(UserId string)(*models.MfaRecoveryCodes, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/mfa/recovery-codes")
	params := map[string]interface{}{}
	params["userId"] = UserId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MfaRecoveryCodes{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MfaRecoveryCodes
	parsed, ok := resp.Result.(models.MfaRecoveryCodes)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateName update the user name by its unique ID.
func (srv *Users) UpdateName(UserId string, Name string)(*models.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/name")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["name"] = Name
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdatePassword update the user password by its unique ID.
func (srv *Users) UpdatePassword(UserId string, Password string)(*models.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/password")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["password"] = Password
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdatePhone update the user phone by its unique ID.
func (srv *Users) UpdatePhone(UserId string, Number string)(*models.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/phone")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["number"] = Number
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetPrefs get the user preferences by its unique ID.
func (srv *Users) GetPrefs(UserId string)(*models.Preferences, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/prefs")
	params := map[string]interface{}{}
	params["userId"] = UserId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Preferences{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
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
			
// UpdatePrefs update the user preferences by its unique ID. The object you
// pass is stored as is, and replaces any previous value. The maximum allowed
// prefs size is 64kB and throws error if exceeded.
func (srv *Users) UpdatePrefs(UserId string, Prefs interface{})(*models.Preferences, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/prefs")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["prefs"] = Prefs
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Preferences{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
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
	
// ListSessions get the user sessions list by its unique ID.
func (srv *Users) ListSessions(UserId string)(*models.SessionList, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/sessions")
	params := map[string]interface{}{}
	params["userId"] = UserId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.SessionList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.SessionList
	parsed, ok := resp.Result.(models.SessionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CreateSession creates a session for a user. Returns an immediately usable
// session object.
// 
// If you want to generate a token for a custom authentication flow, use the
// [POST
// /users/{userId}/tokens](https://appwrite.io/docs/server/users#createToken)
// endpoint.
func (srv *Users) CreateSession(UserId string)(*models.Session, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/sessions")
	params := map[string]interface{}{}
	params["userId"] = UserId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Session{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Session
	parsed, ok := resp.Result.(models.Session)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteSessions delete all user's sessions by using the user's unique ID.
func (srv *Users) DeleteSessions(UserId string)(*interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/sessions")
	params := map[string]interface{}{}
	params["userId"] = UserId
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
			
// DeleteSession delete a user sessions by its unique ID.
func (srv *Users) DeleteSession(UserId string, SessionId string)(*interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId, "{sessionId}", SessionId)
	path := r.Replace("/users/{userId}/sessions/{sessionId}")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["sessionId"] = SessionId
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
			
// UpdateStatus update the user status by its unique ID. Use this endpoint as
// an alternative to deleting a user if you want to keep user's ID reserved.
func (srv *Users) UpdateStatus(UserId string, Status bool)(*models.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/status")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["status"] = Status
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListTargetsOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListTargetsOptions) New() *ListTargetsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListTargetsOption func(*ListTargetsOptions)
func (srv *Users) WithListTargetsQueries(v []string) ListTargetsOption {
	return func(o *ListTargetsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
			
// ListTargets list the messaging targets that are associated with a user.
func (srv *Users) ListTargets(UserId string, optionalSetters ...ListTargetsOption)(*models.TargetList, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/targets")
	options := ListTargetsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.TargetList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.TargetList
	parsed, ok := resp.Result.(models.TargetList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateTargetOptions struct {
	ProviderId string
	Name string
	enabledSetters map[string]bool
}
func (options CreateTargetOptions) New() *CreateTargetOptions {
	options.enabledSetters = map[string]bool{
		"ProviderId": false,
		"Name": false,
	}
	return &options
}
type CreateTargetOption func(*CreateTargetOptions)
func (srv *Users) WithCreateTargetProviderId(v string) CreateTargetOption {
	return func(o *CreateTargetOptions) {
		o.ProviderId = v
		o.enabledSetters["ProviderId"] = true
	}
}
func (srv *Users) WithCreateTargetName(v string) CreateTargetOption {
	return func(o *CreateTargetOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
									
// CreateTarget create a messaging target.
func (srv *Users) CreateTarget(UserId string, TargetId string, ProviderType string, Identifier string, optionalSetters ...CreateTargetOption)(*models.Target, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/targets")
	options := CreateTargetOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["targetId"] = TargetId
	params["providerType"] = ProviderType
	params["identifier"] = Identifier
	if options.enabledSetters["ProviderId"] {
		params["providerId"] = options.ProviderId
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
		bytes := []byte(resp.Result.(string))

		parsed := models.Target{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Target
	parsed, ok := resp.Result.(models.Target)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetTarget get a user's push notification target by ID.
func (srv *Users) GetTarget(UserId string, TargetId string)(*models.Target, error) {
	r := strings.NewReplacer("{userId}", UserId, "{targetId}", TargetId)
	path := r.Replace("/users/{userId}/targets/{targetId}")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["targetId"] = TargetId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Target{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Target
	parsed, ok := resp.Result.(models.Target)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateTargetOptions struct {
	Identifier string
	ProviderId string
	Name string
	enabledSetters map[string]bool
}
func (options UpdateTargetOptions) New() *UpdateTargetOptions {
	options.enabledSetters = map[string]bool{
		"Identifier": false,
		"ProviderId": false,
		"Name": false,
	}
	return &options
}
type UpdateTargetOption func(*UpdateTargetOptions)
func (srv *Users) WithUpdateTargetIdentifier(v string) UpdateTargetOption {
	return func(o *UpdateTargetOptions) {
		o.Identifier = v
		o.enabledSetters["Identifier"] = true
	}
}
func (srv *Users) WithUpdateTargetProviderId(v string) UpdateTargetOption {
	return func(o *UpdateTargetOptions) {
		o.ProviderId = v
		o.enabledSetters["ProviderId"] = true
	}
}
func (srv *Users) WithUpdateTargetName(v string) UpdateTargetOption {
	return func(o *UpdateTargetOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
					
// UpdateTarget update a messaging target.
func (srv *Users) UpdateTarget(UserId string, TargetId string, optionalSetters ...UpdateTargetOption)(*models.Target, error) {
	r := strings.NewReplacer("{userId}", UserId, "{targetId}", TargetId)
	path := r.Replace("/users/{userId}/targets/{targetId}")
	options := UpdateTargetOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["targetId"] = TargetId
	if options.enabledSetters["Identifier"] {
		params["identifier"] = options.Identifier
	}
	if options.enabledSetters["ProviderId"] {
		params["providerId"] = options.ProviderId
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
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

		parsed := models.Target{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Target
	parsed, ok := resp.Result.(models.Target)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// DeleteTarget delete a messaging target.
func (srv *Users) DeleteTarget(UserId string, TargetId string)(*interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId, "{targetId}", TargetId)
	path := r.Replace("/users/{userId}/targets/{targetId}")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["targetId"] = TargetId
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
type CreateTokenOptions struct {
	Length int
	Expire int
	enabledSetters map[string]bool
}
func (options CreateTokenOptions) New() *CreateTokenOptions {
	options.enabledSetters = map[string]bool{
		"Length": false,
		"Expire": false,
	}
	return &options
}
type CreateTokenOption func(*CreateTokenOptions)
func (srv *Users) WithCreateTokenLength(v int) CreateTokenOption {
	return func(o *CreateTokenOptions) {
		o.Length = v
		o.enabledSetters["Length"] = true
	}
}
func (srv *Users) WithCreateTokenExpire(v int) CreateTokenOption {
	return func(o *CreateTokenOptions) {
		o.Expire = v
		o.enabledSetters["Expire"] = true
	}
}
			
// CreateToken returns a token with a secret key for creating a session. Use
// the user ID and secret and submit a request to the [PUT
// /account/sessions/token](https://appwrite.io/docs/references/cloud/client-web/account#createSession)
// endpoint to complete the login process.
func (srv *Users) CreateToken(UserId string, optionalSetters ...CreateTokenOption)(*models.Token, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/tokens")
	options := CreateTokenOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["userId"] = UserId
	if options.enabledSetters["Length"] {
		params["length"] = options.Length
	}
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

		parsed := models.Token{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Token
	parsed, ok := resp.Result.(models.Token)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdateEmailVerification update the user email verification status by its
// unique ID.
func (srv *Users) UpdateEmailVerification(UserId string, EmailVerification bool)(*models.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/verification")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["emailVerification"] = EmailVerification
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// UpdatePhoneVerification update the user phone verification status by its
// unique ID.
func (srv *Users) UpdatePhoneVerification(UserId string, PhoneVerification bool)(*models.User, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/verification/phone")
	params := map[string]interface{}{}
	params["userId"] = UserId
	params["phoneVerification"] = PhoneVerification
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.User{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.User
	parsed, ok := resp.Result.(models.User)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
