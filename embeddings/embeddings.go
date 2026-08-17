package embeddings

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/models"
)

// Embeddings service
type Embeddings struct {
	client client.Client
}

func New(clt client.Client) *Embeddings {
	return &Embeddings{
		client: clt,
	}
}

type CreateTextEmbeddingsOptions struct {
	Model          string
	enabledSetters map[string]bool
}

func (options CreateTextEmbeddingsOptions) New() *CreateTextEmbeddingsOptions {
	options.enabledSetters = map[string]bool{"Model": false}
	return &options
}

type CreateTextEmbeddingsOption func(*CreateTextEmbeddingsOptions)

func (srv *Embeddings) WithCreateTextEmbeddingsModel(v string) CreateTextEmbeddingsOption {
	return func(o *CreateTextEmbeddingsOptions) {
		o.Model = v
		o.enabledSetters["Model"] = true
	}
}

// CreateTextEmbeddings generate vector embeddings for an array of text using
// the selected embedding model. Use the returned vectors to power semantic
// search and similarity queries against your vector collections.
func (srv *Embeddings) CreateTextEmbeddings(Texts []string, optionalSetters ...CreateTextEmbeddingsOption) (*models.EmbeddingList, error) {
	path := "/embeddings/text"
	options := CreateTextEmbeddingsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["texts"] = Texts
	if options.enabledSetters["Model"] {
		params["model"] = options.Model
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

		parsed := models.EmbeddingList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.EmbeddingList
	parsed, ok := resp.Result.(models.EmbeddingList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
