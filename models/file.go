package models

import (
	"encoding/json"
	"errors"
)

// File Model
type File struct {
	// File ID.
	Id string `json:"$id"`
	// Bucket ID.
	BucketId string `json:"bucketId"`
	// File creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// File update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// File permissions. [Learn more about
	// permissions](https://appwrite.io/docs/permissions).
	Permissions []string `json:"$permissions"`
	// File name.
	Name string `json:"name"`
	// Virtual folder containing the file, with a trailing slash. Empty for the
	// bucket root.
	Folder string `json:"folder"`
	// Full virtual path of the file: the folder followed by the file name.
	Key string `json:"key"`
	// File MD5 signature.
	Signature string `json:"signature"`
	// File mime type.
	MimeType string `json:"mimeType"`
	// File original size in bytes.
	SizeOriginal int `json:"sizeOriginal"`
	// File actual stored size in bytes after compression and/or encryption.
	SizeActual int `json:"sizeActual"`
	// Total number of chunks available
	ChunksTotal int `json:"chunksTotal"`
	// Total number of chunks uploaded
	ChunksUploaded int `json:"chunksUploaded"`
	// Whether file contents are encrypted at rest.
	Encryption bool `json:"encryption"`
	// Compression algorithm used for the file. Will be one of none,
	// [gzip](https://en.wikipedia.org/wiki/Gzip), or
	// [zstd](https://en.wikipedia.org/wiki/Zstd).
	Compression string `json:"compression"`

	// Used by Decode() method
	data []byte
}

func (model File) New(data []byte) *File {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *File) UnmarshalJSON(data []byte) error {
	type alias File
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id             json.RawMessage `json:"$id"`
		BucketId       json.RawMessage `json:"bucketId"`
		CreatedAt      json.RawMessage `json:"$createdAt"`
		UpdatedAt      json.RawMessage `json:"$updatedAt"`
		Permissions    json.RawMessage `json:"$permissions"`
		Name           json.RawMessage `json:"name"`
		Folder         json.RawMessage `json:"folder"`
		Key            json.RawMessage `json:"key"`
		Signature      json.RawMessage `json:"signature"`
		MimeType       json.RawMessage `json:"mimeType"`
		SizeOriginal   json.RawMessage `json:"sizeOriginal"`
		SizeActual     json.RawMessage `json:"sizeActual"`
		ChunksTotal    json.RawMessage `json:"chunksTotal"`
		ChunksUploaded json.RawMessage `json:"chunksUploaded"`
		Encryption     json.RawMessage `json:"encryption"`
		Compression    json.RawMessage `json:"compression"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.BucketId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.BucketId); err != nil {
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
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Folder == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Folder); err != nil {
			return err
		}
	}
	if present.Key == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Key); err != nil {
			return err
		}
	}
	if present.Signature == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Signature); err != nil {
			return err
		}
	}
	if present.MimeType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.MimeType); err != nil {
			return err
		}
	}
	if present.SizeOriginal == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.SizeOriginal); err != nil {
			return err
		}
	}
	if present.SizeActual == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.SizeActual); err != nil {
			return err
		}
	}
	if present.ChunksTotal == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ChunksTotal); err != nil {
			return err
		}
	}
	if present.ChunksUploaded == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ChunksUploaded); err != nil {
			return err
		}
	}
	if present.Encryption == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Encryption); err != nil {
			return err
		}
	}
	if present.Compression == nil {
		if err := json.Unmarshal([]byte("\"none\""), &decoded.Compression); err != nil {
			return err
		}
	}

	*model = File(decoded)
	return nil
}

func (model *File) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
