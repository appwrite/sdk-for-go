package graphql

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"strings"
)

// Graphql service
type Graphql struct {
	client client.Client
}

func New(clt client.Client) *Graphql {
	return &Graphql{
		client: clt,
	}
}

	
// Query execute a GraphQL mutation.
func (srv *Graphql) Query(Query interface{})(*interface{}, error) {
	path := "/graphql"
	params := map[string]interface{}{}
	params["query"] = Query
	headers := map[string]interface{}{
		"x-sdk-graphql": "true",
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
	
// Mutation execute a GraphQL mutation.
func (srv *Graphql) Mutation(Query interface{})(*interface{}, error) {
	path := "/graphql/mutation"
	params := map[string]interface{}{}
	params["query"] = Query
	headers := map[string]interface{}{
		"x-sdk-graphql": "true",
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
