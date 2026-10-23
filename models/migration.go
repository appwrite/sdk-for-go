package models

import (
	"encoding/json"
	"errors"
)

// Migration Model
type Migration struct {
	// Migration ID.
	Id string `json:"$id"`
	// Migration creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Variable creation date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Migration status ( pending, processing, failed, completed )
	Status string `json:"status"`
	// Migration stage ( init, processing, source-check, destination-check,
	// migrating, finished )
	Stage string `json:"stage"`
	// A string containing the type of source of the migration.
	Source string `json:"source"`
	// A string containing the type of destination of the migration.
	Destination string `json:"destination"`
	// Resources to migrate.
	Resources []string `json:"resources"`
	// ID of the resource being migrated.
	ResourceId string `json:"resourceId"`
	// Internal ID of the resource being migrated.
	ResourceInternalId string `json:"resourceInternalId"`
	// Type of the resource being migrated.
	ResourceType string `json:"resourceType"`
	// ID of the parent resource that contains the migrated resource.
	ParentResourceId string `json:"parentResourceId"`
	// Internal ID of the parent resource that contains the migrated resource.
	ParentResourceInternalId string `json:"parentResourceInternalId"`
	// Type of the parent resource that contains the migrated resource.
	ParentResourceType string `json:"parentResourceType"`
	// ID of the destination resource created or overwritten by the migration.
	DestinationResourceId string `json:"destinationResourceId"`
	// Internal ID of the destination resource created or overwritten by the
	// migration.
	DestinationResourceInternalId string `json:"destinationResourceInternalId"`
	// Type of the destination resource created or overwritten by the migration.
	DestinationResourceType string `json:"destinationResourceType"`
	// A group of counters that represent the total progress of the migration.
	StatusCounters interface{} `json:"statusCounters"`
	// An array of objects containing the report data of the resources that were
	// migrated.
	ResourceData []interface{} `json:"resourceData"`
	// All errors that occurred during the migration process.
	Errors []string `json:"errors"`
	// Migration options used during the migration process.
	Options interface{} `json:"options"`

	// Used by Decode() method
	data []byte
}

func (model Migration) New(data []byte) *Migration {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Migration) UnmarshalJSON(data []byte) error {
	type alias Migration
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                            json.RawMessage `json:"$id"`
		CreatedAt                     json.RawMessage `json:"$createdAt"`
		UpdatedAt                     json.RawMessage `json:"$updatedAt"`
		Status                        json.RawMessage `json:"status"`
		Stage                         json.RawMessage `json:"stage"`
		Source                        json.RawMessage `json:"source"`
		Destination                   json.RawMessage `json:"destination"`
		Resources                     json.RawMessage `json:"resources"`
		ResourceId                    json.RawMessage `json:"resourceId"`
		ResourceInternalId            json.RawMessage `json:"resourceInternalId"`
		ResourceType                  json.RawMessage `json:"resourceType"`
		ParentResourceId              json.RawMessage `json:"parentResourceId"`
		ParentResourceInternalId      json.RawMessage `json:"parentResourceInternalId"`
		ParentResourceType            json.RawMessage `json:"parentResourceType"`
		DestinationResourceId         json.RawMessage `json:"destinationResourceId"`
		DestinationResourceInternalId json.RawMessage `json:"destinationResourceInternalId"`
		DestinationResourceType       json.RawMessage `json:"destinationResourceType"`
		StatusCounters                json.RawMessage `json:"statusCounters"`
		ResourceData                  json.RawMessage `json:"resourceData"`
		Errors                        json.RawMessage `json:"errors"`
		Options                       json.RawMessage `json:"options"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.UpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UpdatedAt); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Status); err != nil {
			return err
		}
	}
	if present.Stage == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Stage); err != nil {
			return err
		}
	}
	if present.Source == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Source); err != nil {
			return err
		}
	}
	if present.Destination == nil {
		if err := json.Unmarshal([]byte("\"Appwrite\""), &decoded.Destination); err != nil {
			return err
		}
	}
	if present.Resources == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Resources); err != nil {
			return err
		}
	}
	if present.ResourceId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceId); err != nil {
			return err
		}
	}
	if present.ResourceInternalId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceInternalId); err != nil {
			return err
		}
	}
	if present.ResourceType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ResourceType); err != nil {
			return err
		}
	}
	if present.ParentResourceId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ParentResourceId); err != nil {
			return err
		}
	}
	if present.ParentResourceInternalId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ParentResourceInternalId); err != nil {
			return err
		}
	}
	if present.ParentResourceType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ParentResourceType); err != nil {
			return err
		}
	}
	if present.DestinationResourceId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DestinationResourceId); err != nil {
			return err
		}
	}
	if present.DestinationResourceInternalId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DestinationResourceInternalId); err != nil {
			return err
		}
	}
	if present.DestinationResourceType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DestinationResourceType); err != nil {
			return err
		}
	}
	if present.StatusCounters == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.StatusCounters); err != nil {
			return err
		}
	}
	if present.ResourceData == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.ResourceData); err != nil {
			return err
		}
	}
	if present.Errors == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Errors); err != nil {
			return err
		}
	}
	if present.Options == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Options); err != nil {
			return err
		}
	}

	*model = Migration(decoded)
	return nil
}

func (model *Migration) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
