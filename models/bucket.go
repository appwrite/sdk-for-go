package models

import (
	"encoding/json"
	"errors"
)

// Bucket Model
type Bucket struct {
	// Bucket ID.
	Id string `json:"$id"`
	// Bucket creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Bucket update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Bucket permissions. [Learn more about
	// permissions](https://appwrite.io/docs/permissions).
	Permissions []string `json:"$permissions"`
	// Whether file-level security is enabled. [Learn more about
	// permissions](https://appwrite.io/docs/permissions).
	FileSecurity bool `json:"fileSecurity"`
	// Bucket name.
	Name string `json:"name"`
	// Optional description for the bucket.
	Description string `json:"description"`
	// Bucket enabled.
	Enabled bool `json:"enabled"`
	// Maximum file size supported.
	MaximumFileSize int `json:"maximumFileSize"`
	// Allowed file extensions.
	AllowedFileExtensions []string `json:"allowedFileExtensions"`
	// Compression algorithm chosen for compression. Will be one of none,
	// [gzip](https://en.wikipedia.org/wiki/Gzip), or
	// [zstd](https://en.wikipedia.org/wiki/Zstd).
	Compression string `json:"compression"`
	// Bucket is encrypted.
	Encryption bool `json:"encryption"`
	// Virus scanning is enabled.
	Antivirus bool `json:"antivirus"`
	// Image transformations are enabled.
	Transformations bool `json:"transformations"`
	// Total size of this bucket in bytes.
	TotalSize int `json:"totalSize"`

	// Used by Decode() method
	data []byte
}

func (model Bucket) New(data []byte) *Bucket {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Bucket) UnmarshalJSON(data []byte) error {
	type alias Bucket
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                    json.RawMessage `json:"$id"`
		CreatedAt             json.RawMessage `json:"$createdAt"`
		UpdatedAt             json.RawMessage `json:"$updatedAt"`
		Permissions           json.RawMessage `json:"$permissions"`
		FileSecurity          json.RawMessage `json:"fileSecurity"`
		Name                  json.RawMessage `json:"name"`
		Description           json.RawMessage `json:"description"`
		Enabled               json.RawMessage `json:"enabled"`
		MaximumFileSize       json.RawMessage `json:"maximumFileSize"`
		AllowedFileExtensions json.RawMessage `json:"allowedFileExtensions"`
		Compression           json.RawMessage `json:"compression"`
		Encryption            json.RawMessage `json:"encryption"`
		Antivirus             json.RawMessage `json:"antivirus"`
		Transformations       json.RawMessage `json:"transformations"`
		TotalSize             json.RawMessage `json:"totalSize"`
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
	if present.Permissions == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Permissions); err != nil {
			return err
		}
	}
	if present.FileSecurity == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.FileSecurity); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Description == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Description); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.MaximumFileSize == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.MaximumFileSize); err != nil {
			return err
		}
	}
	if present.AllowedFileExtensions == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.AllowedFileExtensions); err != nil {
			return err
		}
	}
	if present.Compression == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Compression); err != nil {
			return err
		}
	}
	if present.Encryption == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Encryption); err != nil {
			return err
		}
	}
	if present.Antivirus == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Antivirus); err != nil {
			return err
		}
	}
	if present.Transformations == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Transformations); err != nil {
			return err
		}
	}
	if present.TotalSize == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.TotalSize); err != nil {
			return err
		}
	}

	*model = Bucket(decoded)
	return nil
}

func (model *Bucket) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
