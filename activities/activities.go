package activities

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// Activities service
type Activities struct {
	client client.Client
}

func New(clt client.Client) *Activities {
	return &Activities{
		client: clt,
	}
}

type ListEventsOptions struct {
	Queries string
	enabledSetters map[string]bool
}
func (options ListEventsOptions) New() *ListEventsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListEventsOption func(*ListEventsOptions)
func (srv *Activities) WithListEventsQueries(v string) ListEventsOption {
	return func(o *ListEventsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
	
// ListEvents list all events for selected filters.
func (srv *Activities) ListEvents(optionalSetters ...ListEventsOption)(*models.ActivityEventList, error) {
	path := "/activities/events"
	options := ListEventsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
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

		parsed := models.ActivityEventList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ActivityEventList
	parsed, ok := resp.Result.(models.ActivityEventList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetEvent get event by ID.
func (srv *Activities) GetEvent(EventId string)(*models.ActivityEvent, error) {
	r := strings.NewReplacer("{eventId}", EventId)
	path := r.Replace("/activities/events/{eventId}")
	params := map[string]interface{}{}
	params["eventId"] = EventId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ActivityEvent{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ActivityEvent
	parsed, ok := resp.Result.(models.ActivityEvent)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
