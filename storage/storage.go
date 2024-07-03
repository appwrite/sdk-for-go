package storage

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/file"
	"strings"
)

// Storage service
type Storage struct {
	client client.Client
}

func NewStorage(clt client.Client) *Storage {
	return &Storage{
		client: clt,
	}
}


type ListBucketsOptions struct {
	Queries []interface{}
	Search string
	enabledSetters map[string]bool
}
func (options ListBucketsOptions) New() *ListBucketsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
	}
	return &options
}
type ListBucketsOption func(*ListBucketsOptions)
func WithListBucketsQueries(v []interface{}) ListBucketsOption {
	return func(o *ListBucketsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func WithListBucketsSearch(v string) ListBucketsOption {
	return func(o *ListBucketsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
	
// ListBuckets get a list of all the storage buckets. You can use the query
// params to filter your results.
func (srv *Storage) ListBuckets(optionalSetters ...ListBucketsOption)(*models.BucketList, error) {
	path := "/storage/buckets"
	options := ListBucketsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.BucketList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.BucketList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateBucketOptions struct {
	Permissions []interface{}
	FileSecurity bool
	Enabled bool
	MaximumFileSize int
	AllowedFileExtensions []interface{}
	Compression string
	Encryption bool
	Antivirus bool
	enabledSetters map[string]bool
}
func (options CreateBucketOptions) New() *CreateBucketOptions {
	options.enabledSetters = map[string]bool{
		"Permissions": false,
		"FileSecurity": false,
		"Enabled": false,
		"MaximumFileSize": false,
		"AllowedFileExtensions": false,
		"Compression": false,
		"Encryption": false,
		"Antivirus": false,
	}
	return &options
}
type CreateBucketOption func(*CreateBucketOptions)
func WithCreateBucketPermissions(v []interface{}) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func WithCreateBucketFileSecurity(v bool) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.FileSecurity = v
		o.enabledSetters["FileSecurity"] = true
	}
}
func WithCreateBucketEnabled(v bool) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func WithCreateBucketMaximumFileSize(v int) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.MaximumFileSize = v
		o.enabledSetters["MaximumFileSize"] = true
	}
}
func WithCreateBucketAllowedFileExtensions(v []interface{}) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.AllowedFileExtensions = v
		o.enabledSetters["AllowedFileExtensions"] = true
	}
}
func WithCreateBucketCompression(v string) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.Compression = v
		o.enabledSetters["Compression"] = true
	}
}
func WithCreateBucketEncryption(v bool) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.Encryption = v
		o.enabledSetters["Encryption"] = true
	}
}
func WithCreateBucketAntivirus(v bool) CreateBucketOption {
	return func(o *CreateBucketOptions) {
		o.Antivirus = v
		o.enabledSetters["Antivirus"] = true
	}
}
					
// CreateBucket create a new storage bucket.
func (srv *Storage) CreateBucket(BucketId string, Name string, optionalSetters ...CreateBucketOption)(*models.Bucket, error) {
	path := "/storage/buckets"
	options := CreateBucketOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["name"] = Name
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["FileSecurity"] {
		params["fileSecurity"] = options.FileSecurity
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["MaximumFileSize"] {
		params["maximumFileSize"] = options.MaximumFileSize
	}
	if options.enabledSetters["AllowedFileExtensions"] {
		params["allowedFileExtensions"] = options.AllowedFileExtensions
	}
	if options.enabledSetters["Compression"] {
		params["compression"] = options.Compression
	}
	if options.enabledSetters["Encryption"] {
		params["encryption"] = options.Encryption
	}
	if options.enabledSetters["Antivirus"] {
		params["antivirus"] = options.Antivirus
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Bucket
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Bucket)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetBucket get a storage bucket by its unique ID. This endpoint response
// returns a JSON object with the storage bucket metadata.
func (srv *Storage) GetBucket(BucketId string)(*models.Bucket, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}")
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Bucket
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Bucket)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateBucketOptions struct {
	Permissions []interface{}
	FileSecurity bool
	Enabled bool
	MaximumFileSize int
	AllowedFileExtensions []interface{}
	Compression string
	Encryption bool
	Antivirus bool
	enabledSetters map[string]bool
}
func (options UpdateBucketOptions) New() *UpdateBucketOptions {
	options.enabledSetters = map[string]bool{
		"Permissions": false,
		"FileSecurity": false,
		"Enabled": false,
		"MaximumFileSize": false,
		"AllowedFileExtensions": false,
		"Compression": false,
		"Encryption": false,
		"Antivirus": false,
	}
	return &options
}
type UpdateBucketOption func(*UpdateBucketOptions)
func WithUpdateBucketPermissions(v []interface{}) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func WithUpdateBucketFileSecurity(v bool) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.FileSecurity = v
		o.enabledSetters["FileSecurity"] = true
	}
}
func WithUpdateBucketEnabled(v bool) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func WithUpdateBucketMaximumFileSize(v int) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.MaximumFileSize = v
		o.enabledSetters["MaximumFileSize"] = true
	}
}
func WithUpdateBucketAllowedFileExtensions(v []interface{}) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.AllowedFileExtensions = v
		o.enabledSetters["AllowedFileExtensions"] = true
	}
}
func WithUpdateBucketCompression(v string) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.Compression = v
		o.enabledSetters["Compression"] = true
	}
}
func WithUpdateBucketEncryption(v bool) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.Encryption = v
		o.enabledSetters["Encryption"] = true
	}
}
func WithUpdateBucketAntivirus(v bool) UpdateBucketOption {
	return func(o *UpdateBucketOptions) {
		o.Antivirus = v
		o.enabledSetters["Antivirus"] = true
	}
}
					
// UpdateBucket update a storage bucket by its unique ID.
func (srv *Storage) UpdateBucket(BucketId string, Name string, optionalSetters ...UpdateBucketOption)(*models.Bucket, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}")
	options := UpdateBucketOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["name"] = Name
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["FileSecurity"] {
		params["fileSecurity"] = options.FileSecurity
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["MaximumFileSize"] {
		params["maximumFileSize"] = options.MaximumFileSize
	}
	if options.enabledSetters["AllowedFileExtensions"] {
		params["allowedFileExtensions"] = options.AllowedFileExtensions
	}
	if options.enabledSetters["Compression"] {
		params["compression"] = options.Compression
	}
	if options.enabledSetters["Encryption"] {
		params["encryption"] = options.Encryption
	}
	if options.enabledSetters["Antivirus"] {
		params["antivirus"] = options.Antivirus
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.Bucket
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.Bucket)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteBucket delete a storage bucket by its unique ID.
func (srv *Storage) DeleteBucket(BucketId string)(*interface{}, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}")
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListFilesOptions struct {
	Queries []interface{}
	Search string
	enabledSetters map[string]bool
}
func (options ListFilesOptions) New() *ListFilesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
	}
	return &options
}
type ListFilesOption func(*ListFilesOptions)
func WithListFilesQueries(v []interface{}) ListFilesOption {
	return func(o *ListFilesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func WithListFilesSearch(v string) ListFilesOption {
	return func(o *ListFilesOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
			
// ListFiles get a list of all the user files. You can use the query params to
// filter your results.
func (srv *Storage) ListFiles(BucketId string, optionalSetters ...ListFilesOption)(*models.FileList, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}/files")
	options := ListFilesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.FileList
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.FileList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateFileOptions struct {
	Permissions []interface{}
	enabledSetters map[string]bool
}
func (options CreateFileOptions) New() *CreateFileOptions {
	options.enabledSetters = map[string]bool{
		"Permissions": false,
	}
	return &options
}
type CreateFileOption func(*CreateFileOptions)
func WithCreateFilePermissions(v []interface{}) CreateFileOption {
	return func(o *CreateFileOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
							
// CreateFile create a new file. Before using this route, you should create a
// new bucket resource using either a [server
// integration](https://appwrite.io/docs/server/storage#storageCreateBucket)
// API or directly from your Appwrite console.
// 
// Larger files should be uploaded using multiple requests with the
// [content-range](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Range)
// header to send a partial request with a maximum supported chunk of `5MB`.
// The `content-range` header values should always be in bytes.
// 
// When the first request is sent, the server will return the **File** object,
// and the subsequent part request must include the file's **id** in
// `x-appwrite-id` header to allow the server to know that the partial upload
// is for the existing file and not for a new one.
// 
// If you're creating a new file using one of the Appwrite SDKs, all the
// chunking logic will be managed by the SDK internally.
func (srv *Storage) CreateFile(BucketId string, FileId string, File file.InputFile, optionalSetters ...CreateFileOption)(*models.File, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	path := r.Replace("/storage/buckets/{bucketId}/files")
	options := CreateFileOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	params["file"] = File
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	headers := map[string]interface{}{
		"content-type": "multipart/form-data",
	}

    paramName := "file"


    uploadId := ""
    uploadId = FileId

    resp, err := srv.client.FileUpload(path, headers, params, paramName, uploadId)
    if err != nil {
		return nil, err
	}
	var parsed models.File
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.File)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil
}
			
// GetFile get a file by its unique ID. This endpoint response returns a JSON
// object with the file metadata.
func (srv *Storage) GetFile(BucketId string, FileId string)(*models.File, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}")
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.File
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.File)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateFileOptions struct {
	Name string
	Permissions []interface{}
	enabledSetters map[string]bool
}
func (options UpdateFileOptions) New() *UpdateFileOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Permissions": false,
	}
	return &options
}
type UpdateFileOption func(*UpdateFileOptions)
func WithUpdateFileName(v string) UpdateFileOption {
	return func(o *UpdateFileOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func WithUpdateFilePermissions(v []interface{}) UpdateFileOption {
	return func(o *UpdateFileOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
					
// UpdateFile update a file by its unique ID. Only users with write
// permissions have access to update this resource.
func (srv *Storage) UpdateFile(BucketId string, FileId string, optionalSetters ...UpdateFileOption)(*models.File, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}")
	options := UpdateFileOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed models.File
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(models.File)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// DeleteFile delete a file by its unique ID. Only users with write
// permissions have access to delete this resource.
func (srv *Storage) DeleteFile(BucketId string, FileId string)(*interface{}, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}")
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetFileDownload get a file content by its unique ID. The endpoint response
// return with a 'Content-Disposition: attachment' header that tells the
// browser to start downloading the file to user downloads directory.
func (srv *Storage) GetFileDownload(BucketId string, FileId string)(*[]byte, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}/download")
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed []byte
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type GetFilePreviewOptions struct {
	Width int
	Height int
	Gravity string
	Quality int
	BorderWidth int
	BorderColor string
	BorderRadius int
	Opacity float64
	Rotation int
	Background string
	Output string
	enabledSetters map[string]bool
}
func (options GetFilePreviewOptions) New() *GetFilePreviewOptions {
	options.enabledSetters = map[string]bool{
		"Width": false,
		"Height": false,
		"Gravity": false,
		"Quality": false,
		"BorderWidth": false,
		"BorderColor": false,
		"BorderRadius": false,
		"Opacity": false,
		"Rotation": false,
		"Background": false,
		"Output": false,
	}
	return &options
}
type GetFilePreviewOption func(*GetFilePreviewOptions)
func WithGetFilePreviewWidth(v int) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func WithGetFilePreviewHeight(v int) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}
func WithGetFilePreviewGravity(v string) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Gravity = v
		o.enabledSetters["Gravity"] = true
	}
}
func WithGetFilePreviewQuality(v int) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Quality = v
		o.enabledSetters["Quality"] = true
	}
}
func WithGetFilePreviewBorderWidth(v int) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.BorderWidth = v
		o.enabledSetters["BorderWidth"] = true
	}
}
func WithGetFilePreviewBorderColor(v string) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.BorderColor = v
		o.enabledSetters["BorderColor"] = true
	}
}
func WithGetFilePreviewBorderRadius(v int) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.BorderRadius = v
		o.enabledSetters["BorderRadius"] = true
	}
}
func WithGetFilePreviewOpacity(v float64) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Opacity = v
		o.enabledSetters["Opacity"] = true
	}
}
func WithGetFilePreviewRotation(v int) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Rotation = v
		o.enabledSetters["Rotation"] = true
	}
}
func WithGetFilePreviewBackground(v string) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Background = v
		o.enabledSetters["Background"] = true
	}
}
func WithGetFilePreviewOutput(v string) GetFilePreviewOption {
	return func(o *GetFilePreviewOptions) {
		o.Output = v
		o.enabledSetters["Output"] = true
	}
}
					
// GetFilePreview get a file preview image. Currently, this method supports
// preview for image files (jpg, png, and gif), other supported formats, like
// pdf, docs, slides, and spreadsheets, will return the file icon image. You
// can also pass query string arguments for cutting and resizing your preview
// image. Preview is supported only for image files smaller than 10MB.
func (srv *Storage) GetFilePreview(BucketId string, FileId string, optionalSetters ...GetFilePreviewOption)(*[]byte, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}/preview")
	options := GetFilePreviewOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	if options.enabledSetters["Width"] {
		params["width"] = options.Width
	}
	if options.enabledSetters["Height"] {
		params["height"] = options.Height
	}
	if options.enabledSetters["Gravity"] {
		params["gravity"] = options.Gravity
	}
	if options.enabledSetters["Quality"] {
		params["quality"] = options.Quality
	}
	if options.enabledSetters["BorderWidth"] {
		params["borderWidth"] = options.BorderWidth
	}
	if options.enabledSetters["BorderColor"] {
		params["borderColor"] = options.BorderColor
	}
	if options.enabledSetters["BorderRadius"] {
		params["borderRadius"] = options.BorderRadius
	}
	if options.enabledSetters["Opacity"] {
		params["opacity"] = options.Opacity
	}
	if options.enabledSetters["Rotation"] {
		params["rotation"] = options.Rotation
	}
	if options.enabledSetters["Background"] {
		params["background"] = options.Background
	}
	if options.enabledSetters["Output"] {
		params["output"] = options.Output
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed []byte
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetFileView get a file content by its unique ID. This endpoint is similar
// to the download method but returns with no  'Content-Disposition:
// attachment' header.
func (srv *Storage) GetFileView(BucketId string, FileId string)(*[]byte, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/storage/buckets/{bucketId}/files/{fileId}/view")
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	var parsed []byte
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.([]byte)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
