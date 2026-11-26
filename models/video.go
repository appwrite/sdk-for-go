package models

import (
	"encoding/json"
	"errors"
)

// Video Model
type Video struct {
	// Video ID.
	Id string `json:"$id"`
	// Video creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Video update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Video permissions. [Learn more about
	// permissions](https://appwrite.io/docs/permissions).
	Permissions []string `json:"$permissions"`
	// Storage bucket ID holding the source file.
	BucketId string `json:"bucketId"`
	// Source file ID.
	FileId string `json:"fileId"`
	// Video name.
	Name string `json:"name"`
	// Source file size in bytes.
	Size int `json:"size"`
	// Container format.
	Format string `json:"format"`
	// Video duration in milliseconds.
	Duration int `json:"duration"`
	// Video width in pixels.
	Width int `json:"width"`
	// Video height in pixels.
	Height int `json:"height"`
	// Video aspect ratio.
	AspectRatio string `json:"aspectRatio"`
	// Video codec.
	VideoCodec string `json:"videoCodec"`
	// Video format.
	VideoFormat string `json:"videoFormat"`
	// Video format profile.
	VideoFormatProfile string `json:"videoFormatProfile"`
	// Video bitrate in bits per second.
	VideoBitRate int `json:"videoBitRate"`
	// Video frame rate in frames per second.
	VideoFrameRate float64 `json:"videoFrameRate"`
	// Video frame rate mode.
	VideoFrameRateMode string `json:"videoFrameRateMode"`
	// Audio codec.
	AudioCodec string `json:"audioCodec"`
	// Audio format.
	AudioFormat string `json:"audioFormat"`
	// Audio bitrate in bits per second.
	AudioBitRate int `json:"audioBitRate"`
	// Audio sample rate in hertz.
	AudioSampleRate int `json:"audioSampleRate"`
	// Thumbnail timeline job status: one of `pending`, `started`, `ready`,
	// `skipped` (source has no video track) or `error`.
	TimelineStatus string `json:"timelineStatus"`
	// Why the timeline job failed. Empty unless `timelineStatus` is `error`.
	TimelineError string `json:"timelineError"`

	// Used by Decode() method
	data []byte
}

func (model Video) New(data []byte) *Video {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Video) UnmarshalJSON(data []byte) error {
	type alias Video
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                 json.RawMessage `json:"$id"`
		CreatedAt          json.RawMessage `json:"$createdAt"`
		UpdatedAt          json.RawMessage `json:"$updatedAt"`
		Permissions        json.RawMessage `json:"$permissions"`
		BucketId           json.RawMessage `json:"bucketId"`
		FileId             json.RawMessage `json:"fileId"`
		Name               json.RawMessage `json:"name"`
		Size               json.RawMessage `json:"size"`
		Format             json.RawMessage `json:"format"`
		Duration           json.RawMessage `json:"duration"`
		Width              json.RawMessage `json:"width"`
		Height             json.RawMessage `json:"height"`
		AspectRatio        json.RawMessage `json:"aspectRatio"`
		VideoCodec         json.RawMessage `json:"videoCodec"`
		VideoFormat        json.RawMessage `json:"videoFormat"`
		VideoFormatProfile json.RawMessage `json:"videoFormatProfile"`
		VideoBitRate       json.RawMessage `json:"videoBitRate"`
		VideoFrameRate     json.RawMessage `json:"videoFrameRate"`
		VideoFrameRateMode json.RawMessage `json:"videoFrameRateMode"`
		AudioCodec         json.RawMessage `json:"audioCodec"`
		AudioFormat        json.RawMessage `json:"audioFormat"`
		AudioBitRate       json.RawMessage `json:"audioBitRate"`
		AudioSampleRate    json.RawMessage `json:"audioSampleRate"`
		TimelineStatus     json.RawMessage `json:"timelineStatus"`
		TimelineError      json.RawMessage `json:"timelineError"`
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
	if present.Size == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Size); err != nil {
			return err
		}
	}
	if present.Format == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Format); err != nil {
			return err
		}
	}
	if present.Duration == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Duration); err != nil {
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
	if present.AspectRatio == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AspectRatio); err != nil {
			return err
		}
	}
	if present.VideoCodec == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.VideoCodec); err != nil {
			return err
		}
	}
	if present.VideoFormat == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.VideoFormat); err != nil {
			return err
		}
	}
	if present.VideoFormatProfile == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.VideoFormatProfile); err != nil {
			return err
		}
	}
	if present.VideoBitRate == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.VideoBitRate); err != nil {
			return err
		}
	}
	if present.VideoFrameRate == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.VideoFrameRate); err != nil {
			return err
		}
	}
	if present.VideoFrameRateMode == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.VideoFrameRateMode); err != nil {
			return err
		}
	}
	if present.AudioCodec == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AudioCodec); err != nil {
			return err
		}
	}
	if present.AudioFormat == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AudioFormat); err != nil {
			return err
		}
	}
	if present.AudioBitRate == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.AudioBitRate); err != nil {
			return err
		}
	}
	if present.AudioSampleRate == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.AudioSampleRate); err != nil {
			return err
		}
	}
	if present.TimelineStatus == nil {
		if err := json.Unmarshal([]byte("\"pending\""), &decoded.TimelineStatus); err != nil {
			return err
		}
	}
	if present.TimelineError == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TimelineError); err != nil {
			return err
		}
	}

	*model = Video(decoded)
	return nil
}

func (model *Video) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
