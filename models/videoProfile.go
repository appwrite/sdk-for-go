package models

import (
	"encoding/json"
	"errors"
)

// VideoProfile Model
type VideoProfile struct {
	// Video profile ID.
	Id string `json:"$id"`
	// Video profile creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Video profile update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Video profile name.
	Name string `json:"name"`
	// Video encode codec: one of `h264`, `hevc`, or `vp9`.
	Codec string `json:"codec"`
	// Target video bitrate in bits per second.
	VideoBitRate int `json:"videoBitRate"`
	// Target audio bitrate in bits per second.
	AudioBitRate int `json:"audioBitRate"`
	// Target video width in pixels.
	Width int `json:"width"`
	// Target video height in pixels.
	Height int `json:"height"`

	// Used by Decode() method
	data []byte
}

func (model VideoProfile) New(data []byte) *VideoProfile {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *VideoProfile) UnmarshalJSON(data []byte) error {
	type alias VideoProfile
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id           json.RawMessage `json:"$id"`
		CreatedAt    json.RawMessage `json:"$createdAt"`
		UpdatedAt    json.RawMessage `json:"$updatedAt"`
		Name         json.RawMessage `json:"name"`
		Codec        json.RawMessage `json:"codec"`
		VideoBitRate json.RawMessage `json:"videoBitRate"`
		AudioBitRate json.RawMessage `json:"audioBitRate"`
		Width        json.RawMessage `json:"width"`
		Height       json.RawMessage `json:"height"`
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
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Codec == nil {
		if err := json.Unmarshal([]byte("\"h264\""), &decoded.Codec); err != nil {
			return err
		}
	}
	if present.VideoBitRate == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.VideoBitRate); err != nil {
			return err
		}
	}
	if present.AudioBitRate == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.AudioBitRate); err != nil {
			return err
		}
	}
	if present.Width == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Width); err != nil {
			return err
		}
	}
	if present.Height == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Height); err != nil {
			return err
		}
	}

	*model = VideoProfile(decoded)
	return nil
}

func (model *VideoProfile) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
