package models

import (
	"encoding/json"
	"errors"
)

// VideoRendition Model
type VideoRendition struct {
	// Rendition ID.
	Id string `json:"$id"`
	// Rendition creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Rendition update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Video ID.
	VideoId string `json:"videoId"`
	// Video profile ID this rendition was encoded against.
	ProfileId string `json:"profileId"`
	// Rendition name, derived from its dimensions and bitrate.
	Name string `json:"name"`
	// Transcoding start time in ISO 8601 format.
	StartedAt string `json:"startedAt"`
	// Transcoding end time in ISO 8601 format.
	EndedAt string `json:"endedAt"`
	// Rendition width in pixels.
	Width int `json:"width"`
	// Rendition height in pixels.
	Height int `json:"height"`
	// Video bitrate in bits per second.
	VideoBitRate int `json:"videoBitRate"`
	// Audio bitrate in bits per second.
	AudioBitRate int `json:"audioBitRate"`
	// Longest segment duration in seconds.
	TargetDuration float64 `json:"targetDuration"`
	// Transcoding status: one of `pending`, `started`, `ended`, `ready`, `error`
	// or `aborted`.
	Status string `json:"status"`
	// Why the rendition failed or was aborted. Empty unless `status` is `error`
	// or `aborted`.
	Error string `json:"error"`
	// Transcoding progress as a percentage, from 0 to 100.
	Progress int `json:"progress"`
	// Streaming output format: `hls`, `dash`, or `cmaf`.
	Output string `json:"output"`
	// Video encode codec snapshotted from the profile at create time: one of
	// `h264`, `hevc`, or `vp9`.
	Codec string `json:"codec"`
	// Stored rendition size in bytes.
	Size int `json:"size"`

	// Used by Decode() method
	data []byte
}

func (model VideoRendition) New(data []byte) *VideoRendition {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *VideoRendition) UnmarshalJSON(data []byte) error {
	type alias VideoRendition
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id             json.RawMessage `json:"$id"`
		CreatedAt      json.RawMessage `json:"$createdAt"`
		UpdatedAt      json.RawMessage `json:"$updatedAt"`
		VideoId        json.RawMessage `json:"videoId"`
		ProfileId      json.RawMessage `json:"profileId"`
		Name           json.RawMessage `json:"name"`
		StartedAt      json.RawMessage `json:"startedAt"`
		EndedAt        json.RawMessage `json:"endedAt"`
		Width          json.RawMessage `json:"width"`
		Height         json.RawMessage `json:"height"`
		VideoBitRate   json.RawMessage `json:"videoBitRate"`
		AudioBitRate   json.RawMessage `json:"audioBitRate"`
		TargetDuration json.RawMessage `json:"targetDuration"`
		Status         json.RawMessage `json:"status"`
		Error          json.RawMessage `json:"error"`
		Progress       json.RawMessage `json:"progress"`
		Output         json.RawMessage `json:"output"`
		Codec          json.RawMessage `json:"codec"`
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
	if present.ProfileId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProfileId); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.StartedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.StartedAt); err != nil {
			return err
		}
	}
	if present.EndedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.EndedAt); err != nil {
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
	if present.Progress == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Progress); err != nil {
			return err
		}
	}
	if present.Output == nil {
		if err := json.Unmarshal([]byte("\"hls\""), &decoded.Output); err != nil {
			return err
		}
	}
	if present.Codec == nil {
		if err := json.Unmarshal([]byte("\"h264\""), &decoded.Codec); err != nil {
			return err
		}
	}
	if present.Size == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Size); err != nil {
			return err
		}
	}

	*model = VideoRendition(decoded)
	return nil
}

func (model *VideoRendition) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
