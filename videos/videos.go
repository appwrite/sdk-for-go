package videos

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/models"
)

// Videos service
type Videos struct {
	client client.Client
}

func New(clt client.Client) *Videos {
	return &Videos{
		client: clt,
	}
}

type ListOptions struct {
	Queries        []string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Total": false}
	return &options
}

type ListOption func(*ListOptions)

func (srv *Videos) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Videos) WithListTotal(v bool) ListOption {
	return func(o *ListOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// List get a list of videos in the current project. Sessions receive only
// videos they can read according to each video's `$permissions`. API keys and
// admins receive the full project catalog. Playback uses the same document
// read ACL plus the `videos.play` (or `videos.read`) scope: a guest can
// stream a video that grants `read("any")` or `read("guests")` even though
// guests lack `videos.read` and will not see it in this list.
func (srv *Videos) List(optionalSetters ...ListOption) (*models.VideoList, error) {
	path := "/videos"
	options := ListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.VideoList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VideoList
	parsed, ok := resp.Result.(models.VideoList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateOptions struct {
	Name           string
	Permissions    []string
	enabledSetters map[string]bool
}

func (options CreateOptions) New() *CreateOptions {
	options.enabledSetters = map[string]bool{"Name": false, "Permissions": false}
	return &options
}

type CreateOption func(*CreateOptions)

func (srv *Videos) WithCreateName(v string) CreateOption {
	return func(o *CreateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Videos) WithCreatePermissions(v []string) CreateOption {
	return func(o *CreateOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}

// Create create a video resource from an existing file in a storage bucket.
// The source file must be a video or audio file, and the caller must be able
// to read that file. Creating a video stores the document and queues timeline
// generation. The worker downloads the source, probes metadata, extracts
// embedded captions once, and deletes its working copy when finished. Sources
// with a video track also get a thumbnail timeline. Call create-rendition to
// package a streaming output.
//
// An optional `name` defaults to the source file name. An optional
// `permissions` array sets the video document ACL (list, get, update, delete,
// manage child resources, and playback). By default, only the current user is
// granted all permissions; API keys and admins create with empty permissions
// when omitted. Guests with `videos.play` can stream when the video grants
// them read (`read("any")` or `read("guests")`).
//
// Uploaded caption files stay listed alongside auto-extracted tracks for the
// same language; when the current default is an auto-extracted track with
// that code, the upload takes the default flag.
func (srv *Videos) Create(VideoId string, BucketId string, FileId string, optionalSetters ...CreateOption) (*models.Video, error) {
	path := "/videos"
	options := CreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["videoId"] = VideoId
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
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

		parsed := models.Video{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Video
	parsed, ok := resp.Result.(models.Video)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ListCodecs get a list of video encode codecs enabled on this instance. Each
// codec includes the output formats it may be packaged into.
func (srv *Videos) ListCodecs() (*models.VideoCodecList, error) {
	path := "/videos/codecs"
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.VideoCodecList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VideoCodecList
	parsed, ok := resp.Result.(models.VideoCodecList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Get get a video by its unique ID. This endpoint response returns a JSON
// object with the video metadata.
func (srv *Videos) Get(VideoId string) (*models.Video, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.Video{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Video
	parsed, ok := resp.Result.(models.Video)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateOptions struct {
	Name           string
	Permissions    []string
	enabledSetters map[string]bool
}

func (options UpdateOptions) New() *UpdateOptions {
	options.enabledSetters = map[string]bool{"Name": false, "Permissions": false}
	return &options
}

type UpdateOption func(*UpdateOptions)

func (srv *Videos) WithUpdateName(v string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Videos) WithUpdatePermissions(v []string) UpdateOption {
	return func(o *UpdateOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}

// Update update a video's name, permissions, or both. Only the fields you
// pass are changed: omit `name` to keep the current name and `permissions` to
// keep the current ACL. The caller must hold update permission on the video
// document. The source file cannot be replaced; delete the video and create a
// new one to point at a different file.
func (srv *Videos) Update(VideoId string, optionalSetters ...UpdateOption) (*models.Video, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}")
	options := UpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.Video{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Video
	parsed, ok := resp.Result.(models.Video)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Delete delete a video by its unique ID. This also removes every rendition,
// caption, preview and transcoded artifact derived from it. Fails with
// `video_rendition_in_progress` while a rendition is started or ended, unless
// that rendition has not been updated within the rendition retention window.
// Fails with `video_timeline_in_progress` while the video's timeline is being
// generated, unless the video has not been updated within the same window.
func (srv *Videos) Delete(VideoId string) (*interface{}, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

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

type ListCaptionsOptions struct {
	Queries        []string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListCaptionsOptions) New() *ListCaptionsOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Total": false}
	return &options
}

type ListCaptionsOption func(*ListCaptionsOptions)

func (srv *Videos) WithListCaptionsQueries(v []string) ListCaptionsOption {
	return func(o *ListCaptionsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Videos) WithListCaptionsTotal(v bool) ListCaptionsOption {
	return func(o *ListCaptionsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListCaptions get a list of all the caption tracks attached to a video. You
// can use the query params to filter your results.
func (srv *Videos) ListCaptions(VideoId string, optionalSetters ...ListCaptionsOption) (*models.VideoCaptionList, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}/captions")
	options := ListCaptionsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.VideoCaptionList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VideoCaptionList
	parsed, ok := resp.Result.(models.VideoCaptionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateCaptionOptions struct {
	Default        bool
	enabledSetters map[string]bool
}

func (options CreateCaptionOptions) New() *CreateCaptionOptions {
	options.enabledSetters = map[string]bool{"Default": false}
	return &options
}

type CreateCaptionOption func(*CreateCaptionOptions)

func (srv *Videos) WithCreateCaptionDefault(v bool) CreateCaptionOption {
	return func(o *CreateCaptionOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}

// CreateCaption add a caption track to a video from a WebVTT or SubRip file
// in a storage bucket. The file is normalised to WebVTT and segmented in the
// background. Subscribe to `videos.[videoId].captions.[captionId].update` for
// `started`, `ready`, and `error` status changes instead of only polling the
// caption document.
//
// Uploaded tracks are listed alongside auto-extracted embedded tracks
// (`embedded: true`). When the current default is an auto-extracted track for
// the same language code, the upload takes the default flag so players pick
// the authored file first. Auto-extracted tracks are never removed
// automatically — delete a track explicitly to remove it. Extraction runs
// once per video, so a deleted extracted track is not re-created.
func (srv *Videos) CreateCaption(VideoId string, BucketId string, FileId string, Name string, Language string, optionalSetters ...CreateCaptionOption) (*models.VideoCaption, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}/captions")
	options := CreateCaptionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	params["name"] = Name
	params["language"] = Language
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
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

		parsed := models.VideoCaption{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VideoCaption
	parsed, ok := resp.Result.(models.VideoCaption)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateCaptionOptions struct {
	BucketId       string
	FileId         string
	Name           string
	Language       string
	Default        bool
	enabledSetters map[string]bool
}

func (options UpdateCaptionOptions) New() *UpdateCaptionOptions {
	options.enabledSetters = map[string]bool{"BucketId": false, "FileId": false, "Name": false, "Language": false, "Default": false}
	return &options
}

type UpdateCaptionOption func(*UpdateCaptionOptions)

func (srv *Videos) WithUpdateCaptionBucketId(v string) UpdateCaptionOption {
	return func(o *UpdateCaptionOptions) {
		o.BucketId = v
		o.enabledSetters["BucketId"] = true
	}
}
func (srv *Videos) WithUpdateCaptionFileId(v string) UpdateCaptionOption {
	return func(o *UpdateCaptionOptions) {
		o.FileId = v
		o.enabledSetters["FileId"] = true
	}
}
func (srv *Videos) WithUpdateCaptionName(v string) UpdateCaptionOption {
	return func(o *UpdateCaptionOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Videos) WithUpdateCaptionLanguage(v string) UpdateCaptionOption {
	return func(o *UpdateCaptionOptions) {
		o.Language = v
		o.enabledSetters["Language"] = true
	}
}
func (srv *Videos) WithUpdateCaptionDefault(v bool) UpdateCaptionOption {
	return func(o *UpdateCaptionOptions) {
		o.Default = v
		o.enabledSetters["Default"] = true
	}
}

// UpdateCaption update a caption track by its unique ID. Pass only `name`,
// `language`, and/or `default` to retag an extracted track (for example `und`
// → `heb` for Hebrew) without replacing the file. Changing
// `bucketId`/`fileId` re-packages the track in the background.
func (srv *Videos) UpdateCaption(VideoId string, CaptionId string, optionalSetters ...UpdateCaptionOption) (*models.VideoCaption, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}
	if CaptionId == "" {
		return nil, errors.New("Missing required parameter: \"captionId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId), "{captionId}", client.EncodePath(CaptionId))
	path := r.Replace("/videos/{videoId}/captions/{captionId}")
	options := UpdateCaptionOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["BucketId"] {
		params["bucketId"] = options.BucketId
	}
	if options.enabledSetters["FileId"] {
		params["fileId"] = options.FileId
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Language"] {
		params["language"] = options.Language
	}
	if options.enabledSetters["Default"] {
		params["default"] = options.Default
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.VideoCaption{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VideoCaption
	parsed, ok := resp.Result.(models.VideoCaption)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteCaption delete a caption track by its unique ID, along with its
// packaged segments.
func (srv *Videos) DeleteCaption(VideoId string, CaptionId string) (*interface{}, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}
	if CaptionId == "" {
		return nil, errors.New("Missing required parameter: \"captionId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId), "{captionId}", client.EncodePath(CaptionId))
	path := r.Replace("/videos/{videoId}/captions/{captionId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

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

func (srv *Videos) getCmafHlsManifestParams() map[string]interface{} {
	params := map[string]interface{}{}
	return params
}

// GetCmafHlsManifest get the top-level streaming manifest for a video output:
// an HLS master playlist (`master.m3u8`) or a DASH MPD (`master.mpd`). Hand
// this URL to a player to begin adaptive playback.
//
// Guests can request this URL when they hold `videos.play` (or `videos.read`)
// **and** the video document grants them read. A video with `read("any")` or
// `read("guests")` is playable without a session.
//
// For CMAF renditions, use `getCmafHlsManifest` or `getCmafDashManifest`.
// Both masters point at the same fMP4 segment set.
func (srv *Videos) GetCmafHlsManifest(VideoId string) (*[]byte, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}/outputs/cmaf/master.m3u8")
	params := srv.getCmafHlsManifestParams()
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "*/*"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		var parsed []byte

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed []byte
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetCmafHlsManifestURL get the top-level streaming manifest for a video
// output: an HLS master playlist (`master.m3u8`) or a DASH MPD
// (`master.mpd`). Hand this URL to a player to begin adaptive playback.
//
// Guests can request this URL when they hold `videos.play` (or `videos.read`)
// **and** the video document grants them read. A video with `read("any")` or
// `read("guests")` is playable without a session.
//
// For CMAF renditions, use `getCmafHlsManifest` or `getCmafDashManifest`.
// Both masters point at the same fMP4 segment set.
// Returns the URL for the resource instead of the content.
func (srv *Videos) GetCmafHlsManifestURL(VideoId string) (*string, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}/outputs/cmaf/master.m3u8")
	params := srv.getCmafHlsManifestParams()
	u, err := url.Parse(srv.client.Endpoint + path)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	for key, value := range params {
		client.AddQueryParam(q, key, value)
	}
	u.RawQuery = q.Encode()
	result := u.String()
	return &result, nil
}

func (srv *Videos) getCmafDashManifestParams() map[string]interface{} {
	params := map[string]interface{}{}
	return params
}

// GetCmafDashManifest get the top-level streaming manifest for a video
// output: an HLS master playlist (`master.m3u8`) or a DASH MPD
// (`master.mpd`). Hand this URL to a player to begin adaptive playback.
//
// Guests can request this URL when they hold `videos.play` (or `videos.read`)
// **and** the video document grants them read. A video with `read("any")` or
// `read("guests")` is playable without a session.
//
// For CMAF renditions, use `getCmafHlsManifest` or `getCmafDashManifest`.
// Both masters point at the same fMP4 segment set.
func (srv *Videos) GetCmafDashManifest(VideoId string) (*[]byte, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}/outputs/cmaf/master.mpd")
	params := srv.getCmafDashManifestParams()
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "*/*"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		var parsed []byte

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed []byte
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetCmafDashManifestURL get the top-level streaming manifest for a video
// output: an HLS master playlist (`master.m3u8`) or a DASH MPD
// (`master.mpd`). Hand this URL to a player to begin adaptive playback.
//
// Guests can request this URL when they hold `videos.play` (or `videos.read`)
// **and** the video document grants them read. A video with `read("any")` or
// `read("guests")` is playable without a session.
//
// For CMAF renditions, use `getCmafHlsManifest` or `getCmafDashManifest`.
// Both masters point at the same fMP4 segment set.
// Returns the URL for the resource instead of the content.
func (srv *Videos) GetCmafDashManifestURL(VideoId string) (*string, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}/outputs/cmaf/master.mpd")
	params := srv.getCmafDashManifestParams()
	u, err := url.Parse(srv.client.Endpoint + path)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	for key, value := range params {
		client.AddQueryParam(q, key, value)
	}
	u.RawQuery = q.Encode()
	result := u.String()
	return &result, nil
}

func (srv *Videos) getDashManifestParams() map[string]interface{} {
	params := map[string]interface{}{}
	return params
}

// GetDashManifest get the top-level streaming manifest for a video output: an
// HLS master playlist (`master.m3u8`) or a DASH MPD (`master.mpd`). Hand this
// URL to a player to begin adaptive playback.
//
// Guests can request this URL when they hold `videos.play` (or `videos.read`)
// **and** the video document grants them read. A video with `read("any")` or
// `read("guests")` is playable without a session.
//
// For CMAF renditions, use `getCmafHlsManifest` or `getCmafDashManifest`.
// Both masters point at the same fMP4 segment set.
func (srv *Videos) GetDashManifest(VideoId string) (*[]byte, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}/outputs/dash/master.mpd")
	params := srv.getDashManifestParams()
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "*/*"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		var parsed []byte

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed []byte
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetDashManifestURL get the top-level streaming manifest for a video output:
// an HLS master playlist (`master.m3u8`) or a DASH MPD (`master.mpd`). Hand
// this URL to a player to begin adaptive playback.
//
// Guests can request this URL when they hold `videos.play` (or `videos.read`)
// **and** the video document grants them read. A video with `read("any")` or
// `read("guests")` is playable without a session.
//
// For CMAF renditions, use `getCmafHlsManifest` or `getCmafDashManifest`.
// Both masters point at the same fMP4 segment set.
// Returns the URL for the resource instead of the content.
func (srv *Videos) GetDashManifestURL(VideoId string) (*string, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}/outputs/dash/master.mpd")
	params := srv.getDashManifestParams()
	u, err := url.Parse(srv.client.Endpoint + path)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	for key, value := range params {
		client.AddQueryParam(q, key, value)
	}
	u.RawQuery = q.Encode()
	result := u.String()
	return &result, nil
}

func (srv *Videos) getHlsManifestParams() map[string]interface{} {
	params := map[string]interface{}{}
	return params
}

// GetHlsManifest get the top-level streaming manifest for a video output: an
// HLS master playlist (`master.m3u8`) or a DASH MPD (`master.mpd`). Hand this
// URL to a player to begin adaptive playback.
//
// Guests can request this URL when they hold `videos.play` (or `videos.read`)
// **and** the video document grants them read. A video with `read("any")` or
// `read("guests")` is playable without a session.
//
// For CMAF renditions, use `getCmafHlsManifest` or `getCmafDashManifest`.
// Both masters point at the same fMP4 segment set.
func (srv *Videos) GetHlsManifest(VideoId string) (*[]byte, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}/outputs/hls/master.m3u8")
	params := srv.getHlsManifestParams()
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "*/*"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		var parsed []byte

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed []byte
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetHlsManifestURL get the top-level streaming manifest for a video output:
// an HLS master playlist (`master.m3u8`) or a DASH MPD (`master.mpd`). Hand
// this URL to a player to begin adaptive playback.
//
// Guests can request this URL when they hold `videos.play` (or `videos.read`)
// **and** the video document grants them read. A video with `read("any")` or
// `read("guests")` is playable without a session.
//
// For CMAF renditions, use `getCmafHlsManifest` or `getCmafDashManifest`.
// Both masters point at the same fMP4 segment set.
// Returns the URL for the resource instead of the content.
func (srv *Videos) GetHlsManifestURL(VideoId string) (*string, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}/outputs/hls/master.m3u8")
	params := srv.getHlsManifestParams()
	u, err := url.Parse(srv.client.Endpoint + path)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	for key, value := range params {
		client.AddQueryParam(q, key, value)
	}
	u.RawQuery = q.Encode()
	result := u.String()
	return &result, nil
}

func (srv *Videos) getCaptionManifestParams() map[string]interface{} {
	params := map[string]interface{}{}
	return params
}

// GetCaptionManifest get the playable caption resource for a video output: an
// HLS captions playlist, or the WebVTT file itself for DASH.
//
// Guests can request this URL when they hold `videos.play` (or `videos.read`)
// **and** the video document grants them read.
func (srv *Videos) GetCaptionManifest(VideoId string, Output string, CaptionId string) (*[]byte, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}
	if CaptionId == "" {
		return nil, errors.New("Missing required parameter: \"captionId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId), "{output}", client.EncodePath(Output), "{captionId}", client.EncodePath(CaptionId))
	path := r.Replace("/videos/{videoId}/outputs/{output}/captions/{captionId}/manifest")
	params := srv.getCaptionManifestParams()
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "*/*"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		var parsed []byte

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed []byte
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetCaptionManifestURL get the playable caption resource for a video output:
// an HLS captions playlist, or the WebVTT file itself for DASH.
//
// Guests can request this URL when they hold `videos.play` (or `videos.read`)
// **and** the video document grants them read.
// Returns the URL for the resource instead of the content.
func (srv *Videos) GetCaptionManifestURL(VideoId string, Output string, CaptionId string) (*string, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}
	if CaptionId == "" {
		return nil, errors.New("Missing required parameter: \"captionId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId), "{output}", client.EncodePath(Output), "{captionId}", client.EncodePath(CaptionId))
	path := r.Replace("/videos/{videoId}/outputs/{output}/captions/{captionId}/manifest")
	params := srv.getCaptionManifestParams()
	u, err := url.Parse(srv.client.Endpoint + path)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	for key, value := range params {
		client.AddQueryParam(q, key, value)
	}
	u.RawQuery = q.Encode()
	result := u.String()
	return &result, nil
}

type ListRenditionsOptions struct {
	Queries        []string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListRenditionsOptions) New() *ListRenditionsOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Total": false}
	return &options
}

type ListRenditionsOption func(*ListRenditionsOptions)

func (srv *Videos) WithListRenditionsQueries(v []string) ListRenditionsOption {
	return func(o *ListRenditionsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Videos) WithListRenditionsTotal(v bool) ListRenditionsOption {
	return func(o *ListRenditionsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListRenditions get a list of all the renditions of a video. You can use the
// query params to filter your results.
func (srv *Videos) ListRenditions(VideoId string, optionalSetters ...ListRenditionsOption) (*models.VideoRenditionList, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}/renditions")
	options := ListRenditionsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.VideoRenditionList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VideoRenditionList
	parsed, ok := resp.Result.(models.VideoRenditionList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreateRendition request a new rendition of a video, encoded against a video
// profile and packaged for HLS, DASH, or CMAF. The rendition is created
// immediately with a `pending` status and transcoded in the background; poll
// it or subscribe to realtime events to follow its progress. The worker
// downloads the source file for this job, probes metadata and extracts
// embedded captions when needed, then deletes the working copy when the job
// ends.
//
// The profile codec must be enabled on this instance. A disabled codec fails
// with `video_codec_disabled`.
//
// Each video may have only one rendition per profile and output combination.
// Creating a duplicate fails with `video_rendition_already_exists`. After an
// encode fails or is aborted, delete the rendition and create it again. The
// same profile may still be encoded for different outputs (for example HLS
// and DASH).
//
// A video accepts 10 rendition creates per hour. The next create returns
// `general_rate_limit_exceeded`.
//
// CMAF packs shared fMP4 segments once and exposes both HLS
// (`/outputs/cmaf/master.m3u8`) and DASH (`/outputs/cmaf/master.mpd`) masters
// for the same encode.
//
// Renditions can only be created from the Console or a server SDK with an API
// key that has the `videos.write` scope. Each encode spends the project's
// compute, so client sessions receive `user_unauthorized`.
func (srv *Videos) CreateRendition(VideoId string, ProfileId string, Output string) (*models.VideoRendition, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}/renditions")
	params := map[string]interface{}{}
	params["profileId"] = ProfileId
	params["output"] = Output
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

		parsed := models.VideoRendition{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VideoRendition
	parsed, ok := resp.Result.(models.VideoRendition)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetRendition get a rendition by its unique ID, including its current
// transcoding status and progress.
func (srv *Videos) GetRendition(VideoId string, RenditionId string) (*models.VideoRendition, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}
	if RenditionId == "" {
		return nil, errors.New("Missing required parameter: \"renditionId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId), "{renditionId}", client.EncodePath(RenditionId))
	path := r.Replace("/videos/{videoId}/renditions/{renditionId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.VideoRendition{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.VideoRendition
	parsed, ok := resp.Result.(models.VideoRendition)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteRendition delete a rendition by its unique ID, along with its
// packaged segments and transcoded output. Fails with
// `video_rendition_in_progress` while the rendition is started or ended,
// unless it has not been updated within the rendition retention window.
func (srv *Videos) DeleteRendition(VideoId string, RenditionId string) (*interface{}, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}
	if RenditionId == "" {
		return nil, errors.New("Missing required parameter: \"renditionId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId), "{renditionId}", client.EncodePath(RenditionId))
	path := r.Replace("/videos/{videoId}/renditions/{renditionId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

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

type GetThumbnailOptions struct {
	Width          int
	Height         int
	Output         string
	enabledSetters map[string]bool
}

func (options GetThumbnailOptions) New() *GetThumbnailOptions {
	options.enabledSetters = map[string]bool{"Width": false, "Height": false, "Output": false}
	return &options
}

type GetThumbnailOption func(*GetThumbnailOptions)

func (srv *Videos) WithGetThumbnailWidth(v int) GetThumbnailOption {
	return func(o *GetThumbnailOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func (srv *Videos) WithGetThumbnailHeight(v int) GetThumbnailOption {
	return func(o *GetThumbnailOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}
func (srv *Videos) WithGetThumbnailOutput(v string) GetThumbnailOption {
	return func(o *GetThumbnailOptions) {
		o.Output = v
		o.enabledSetters["Output"] = true
	}
}

func (srv *Videos) getThumbnailParams(optionalSetters ...GetThumbnailOption) map[string]interface{} {
	options := GetThumbnailOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Width"] {
		params["width"] = options.Width
	}
	if options.enabledSetters["Height"] {
		params["height"] = options.Height
	}
	if options.enabledSetters["Output"] {
		params["output"] = options.Output
	}
	return params
}

// GetThumbnail get a single thumbnail sheet from a video's timeline as an
// image. You can resize the returned image and choose its output format.
//
// Guests can request this URL when they hold `videos.play` (or `videos.read`)
// **and** the video document grants them read.
func (srv *Videos) GetThumbnail(VideoId string, ThumbnailId string, optionalSetters ...GetThumbnailOption) (*[]byte, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}
	if ThumbnailId == "" {
		return nil, errors.New("Missing required parameter: \"thumbnailId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId), "{thumbnailId}", client.EncodePath(ThumbnailId))
	path := r.Replace("/videos/{videoId}/thumbnails/{thumbnailId}")
	params := srv.getThumbnailParams(optionalSetters...)
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "image/*"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		var parsed []byte

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed []byte
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetThumbnailURL get a single thumbnail sheet from a video's timeline as an
// image. You can resize the returned image and choose its output format.
//
// Guests can request this URL when they hold `videos.play` (or `videos.read`)
// **and** the video document grants them read.
// Returns the URL for the resource instead of the content.
func (srv *Videos) GetThumbnailURL(VideoId string, ThumbnailId string, optionalSetters ...GetThumbnailOption) (*string, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}
	if ThumbnailId == "" {
		return nil, errors.New("Missing required parameter: \"thumbnailId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId), "{thumbnailId}", client.EncodePath(ThumbnailId))
	path := r.Replace("/videos/{videoId}/thumbnails/{thumbnailId}")
	params := srv.getThumbnailParams(optionalSetters...)
	u, err := url.Parse(srv.client.Endpoint + path)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	for key, value := range params {
		client.AddQueryParam(q, key, value)
	}
	u.RawQuery = q.Encode()
	result := u.String()
	return &result, nil
}

func (srv *Videos) getTimelineParams() map[string]interface{} {
	params := map[string]interface{}{}
	return params
}

// GetTimeline get the WebVTT thumbnail timeline for a video, used to render
// scrubbing thumbnails in a player. Timeline generation is queued when the
// video is created. Subscribe to `videos.[videoId].timelines.update` for
// probe and terminal updates, or poll this endpoint until it returns a WebVTT
// document. Audio-only sources produce no timeline and this endpoint returns
// `video_timelines_not_found`.
//
// Guests can request this URL when they hold `videos.play` (or `videos.read`)
// **and** the video document grants them read. A video with `read("any")` or
// `read("guests")` is playable without a session.
func (srv *Videos) GetTimeline(VideoId string) (*[]byte, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}/timeline")
	params := srv.getTimelineParams()
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "*/*"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		var parsed []byte

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed []byte
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetTimelineURL get the WebVTT thumbnail timeline for a video, used to
// render scrubbing thumbnails in a player. Timeline generation is queued when
// the video is created. Subscribe to `videos.[videoId].timelines.update` for
// probe and terminal updates, or poll this endpoint until it returns a WebVTT
// document. Audio-only sources produce no timeline and this endpoint returns
// `video_timelines_not_found`.
//
// Guests can request this URL when they hold `videos.play` (or `videos.read`)
// **and** the video document grants them read. A video with `read("any")` or
// `read("guests")` is playable without a session.
// Returns the URL for the resource instead of the content.
func (srv *Videos) GetTimelineURL(VideoId string) (*string, error) {
	if VideoId == "" {
		return nil, errors.New("Missing required parameter: \"videoId\"")
	}

	r := strings.NewReplacer("{videoId}", client.EncodePath(VideoId))
	path := r.Replace("/videos/{videoId}/timeline")
	params := srv.getTimelineParams()
	u, err := url.Parse(srv.client.Endpoint + path)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	for key, value := range params {
		client.AddQueryParam(q, key, value)
	}
	u.RawQuery = q.Encode()
	result := u.String()
	return &result, nil
}
