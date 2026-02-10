package organizations

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// Organizations service
type Organizations struct {
	client client.Client
}

func New(clt client.Client) *Organizations {
	return &Organizations{
		client: clt,
	}
}

	
// Delete delete an organization.
func (srv *Organizations) Delete(OrganizationId string)(*interface{}, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
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
	
// EstimationDeleteOrganization get estimation for deleting an organization.
func (srv *Organizations) EstimationDeleteOrganization(OrganizationId string)(*models.EstimationDeleteOrganization, error) {
	r := strings.NewReplacer("{organizationId}", OrganizationId)
	path := r.Replace("/organizations/{organizationId}/estimations/delete-organization")
	params := map[string]interface{}{}
	params["organizationId"] = OrganizationId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.EstimationDeleteOrganization{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.EstimationDeleteOrganization
	parsed, ok := resp.Result.(models.EstimationDeleteOrganization)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
