package models

import (
	"encoding/json"
	"errors"
)

// VideoCaption Model
type VideoCaption struct {
	// Caption ID.
	Id string `json:"$id"`
	// Caption creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Caption update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Video ID.
	VideoId string `json:"videoId"`
	// Storage bucket ID holding the caption file.
	BucketId string `json:"bucketId"`
	// Caption file ID.
	FileId string `json:"fileId"`
	// Caption display name.
	Name string `json:"name"`
	// Caption ISO 639-2 language code.
	Language string `json:"language"`
	// Is this the default caption track?
	Default bool `json:"default"`
	// Was this track auto-extracted from the source container? Extracted tracks
	// have no backing file; extraction runs once per video, so a deleted
	// extracted track is not re-created.
	Embedded bool `json:"embedded"`
	// Longest segment duration in seconds.
	TargetDuration float64 `json:"targetDuration"`
	// Packaging status: one of `pending`, `started`, `ready`, `error` or
	// `aborted`.
	Status string `json:"status"`
	// Why the caption failed or was aborted. Empty unless `status` is `error` or
	// `aborted`.
	Error string `json:"error"`
	// Stored caption size in bytes.
	Size int `json:"size"`

	// Used by Decode() method
	data []byte
}

func (model VideoCaption) New(data []byte) *VideoCaption {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *VideoCaption) UnmarshalJSON(data []byte) error {
	type alias VideoCaption
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id             json.RawMessage `json:"$id"`
		CreatedAt      json.RawMessage `json:"$createdAt"`
		UpdatedAt      json.RawMessage `json:"$updatedAt"`
		VideoId        json.RawMessage `json:"videoId"`
		BucketId       json.RawMessage `json:"bucketId"`
		FileId         json.RawMessage `json:"fileId"`
		Name           json.RawMessage `json:"name"`
		Language       json.RawMessage `json:"language"`
		Default        json.RawMessage `json:"default"`
		Embedded       json.RawMessage `json:"embedded"`
		TargetDuration json.RawMessage `json:"targetDuration"`
		Status         json.RawMessage `json:"status"`
		Error          json.RawMessage `json:"error"`
		Size           json.RawMessage `json:"size"`
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
	if present.VideoId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.VideoId); err != nil {
			return err
		}
	}
	if present.BucketId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.BucketId); err != nil {
			return err
		}
	}
	if present.FileId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.FileId); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Language == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Language); err != nil {
			return err
		}
	}
	if present.Default == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Default); err != nil {
			return err
		}
	}
	if present.Embedded == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Embedded); err != nil {
			return err
		}
	}
	if present.TargetDuration == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.TargetDuration); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"pending\""), &decoded.Status); err != nil {
			return err
		}
	}
	if present.Error == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Error); err != nil {
			return err
		}
	}
	if present.Size == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Size); err != nil {
			return err
		}
	}

	*model = VideoCaption(decoded)
	return nil
}

func (model *VideoCaption) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
