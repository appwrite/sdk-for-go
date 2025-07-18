package tokens

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// Tokens service
type Tokens struct {
	client client.Client
}

func New(clt client.Client) *Tokens {
	return &Tokens{
		client: clt,
	}
}

type ListOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListOption func(*ListOptions)
func (srv *Tokens) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
					
// List list all the tokens created for a specific file or bucket. You can use
// the query params to filter your results.
func (srv *Tokens) List(BucketId string, FileId string, optionalSetters ...ListOption)(*models.ResourceTokenList, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/tokens/buckets/{bucketId}/files/{fileId}")
	options := ListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
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

		parsed := models.ResourceTokenList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ResourceTokenList
	parsed, ok := resp.Result.(models.ResourceTokenList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateFileTokenOptions struct {
	Expire string
	enabledSetters map[string]bool
}
func (options CreateFileTokenOptions) New() *CreateFileTokenOptions {
	options.enabledSetters = map[string]bool{
		"Expire": false,
	}
	return &options
}
type CreateFileTokenOption func(*CreateFileTokenOptions)
func (srv *Tokens) WithCreateFileTokenExpire(v string) CreateFileTokenOption {
	return func(o *CreateFileTokenOptions) {
		o.Expire = v
		o.enabledSetters["Expire"] = true
	}
}
					
// CreateFileToken create a new token. A token is linked to a file. Token can
// be passed as a request URL search parameter.
func (srv *Tokens) CreateFileToken(BucketId string, FileId string, optionalSetters ...CreateFileTokenOption)(*models.ResourceToken, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/tokens/buckets/{bucketId}/files/{fileId}")
	options := CreateFileTokenOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
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

		parsed := models.ResourceToken{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ResourceToken
	parsed, ok := resp.Result.(models.ResourceToken)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Get get a token by its unique ID.
func (srv *Tokens) Get(TokenId string)(*models.ResourceToken, error) {
	r := strings.NewReplacer("{tokenId}", TokenId)
	path := r.Replace("/tokens/{tokenId}")
	params := map[string]interface{}{}
	params["tokenId"] = TokenId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ResourceToken{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ResourceToken
	parsed, ok := resp.Result.(models.ResourceToken)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateOptions struct {
	Expire string
	enabledSetters map[string]bool
}
func (options UpdateOptions) New() *UpdateOptions {
	options.enabledSetters = map[string]bool{
		"Expire": false,
	}
	return &options
}
type UpdateOption func(*UpdateOptions)
func (srv *Tokens) WithUpdateExpire(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Expire = v
		o.enabledSetters["Expire"] = true
	}
}
			
// Update update a token by its unique ID. Use this endpoint to update a
// token's expiry date.
func (srv *Tokens) Update(TokenId string, optionalSetters ...UpdateOption)(*models.ResourceToken, error) {
	r := strings.NewReplacer("{tokenId}", TokenId)
	path := r.Replace("/tokens/{tokenId}")
	options := UpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["tokenId"] = TokenId
	if options.enabledSetters["Expire"] {
		params["expire"] = options.Expire
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

		parsed := models.ResourceToken{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ResourceToken
	parsed, ok := resp.Result.(models.ResourceToken)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Delete delete a token by its unique ID.
func (srv *Tokens) Delete(TokenId string)(*interface{}, error) {
	r := strings.NewReplacer("{tokenId}", TokenId)
	path := r.Replace("/tokens/{tokenId}")
	params := map[string]interface{}{}
	params["tokenId"] = TokenId
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
