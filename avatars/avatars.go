package avatars

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
)

// Avatars service
type Avatars struct {
	client client.Client
}

func New(clt client.Client) *Avatars {
	return &Avatars{
		client: clt,
	}
}

type GetBrowserOptions struct {
	Width          int
	Height         int
	Quality        int
	enabledSetters map[string]bool
}

func (options GetBrowserOptions) New() *GetBrowserOptions {
	options.enabledSetters = map[string]bool{"Width": false, "Height": false, "Quality": false}
	return &options
}

type GetBrowserOption func(*GetBrowserOptions)

func (srv *Avatars) WithGetBrowserWidth(v int) GetBrowserOption {
	return func(o *GetBrowserOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func (srv *Avatars) WithGetBrowserHeight(v int) GetBrowserOption {
	return func(o *GetBrowserOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}
func (srv *Avatars) WithGetBrowserQuality(v int) GetBrowserOption {
	return func(o *GetBrowserOptions) {
		o.Quality = v
		o.enabledSetters["Quality"] = true
	}
}

func (srv *Avatars) getBrowserParams(optionalSetters ...GetBrowserOption) map[string]interface{} {
	options := GetBrowserOptions{}.New()
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
	if options.enabledSetters["Quality"] {
		params["quality"] = options.Quality
	}
	return params
}

// GetBrowser you can use this endpoint to show different browser icons to
// your users. The code argument receives the browser code as it appears in
// your user [GET
// /account/sessions](https://appwrite.io/docs/references/cloud/client-web/account#getSessions)
// endpoint. Use width, height and quality arguments to change the output
// settings.
//
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
func (srv *Avatars) GetBrowser(Code string, optionalSetters ...GetBrowserOption) (*[]byte, error) {
	r := strings.NewReplacer("{code}", client.EncodePath(Code))
	path := r.Replace("/avatars/browsers/{code}")
	params := srv.getBrowserParams(optionalSetters...)
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "image/png"

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

// GetBrowserURL you can use this endpoint to show different browser icons to
// your users. The code argument receives the browser code as it appears in
// your user [GET
// /account/sessions](https://appwrite.io/docs/references/cloud/client-web/account#getSessions)
// endpoint. Use width, height and quality arguments to change the output
// settings.
//
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
// Returns the URL for the resource instead of the content.
func (srv *Avatars) GetBrowserURL(Code string, optionalSetters ...GetBrowserOption) (*string, error) {
	r := strings.NewReplacer("{code}", client.EncodePath(Code))
	path := r.Replace("/avatars/browsers/{code}")
	params := srv.getBrowserParams(optionalSetters...)
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

type GetCreditCardOptions struct {
	Width          int
	Height         int
	Quality        int
	enabledSetters map[string]bool
}

func (options GetCreditCardOptions) New() *GetCreditCardOptions {
	options.enabledSetters = map[string]bool{"Width": false, "Height": false, "Quality": false}
	return &options
}

type GetCreditCardOption func(*GetCreditCardOptions)

func (srv *Avatars) WithGetCreditCardWidth(v int) GetCreditCardOption {
	return func(o *GetCreditCardOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func (srv *Avatars) WithGetCreditCardHeight(v int) GetCreditCardOption {
	return func(o *GetCreditCardOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}
func (srv *Avatars) WithGetCreditCardQuality(v int) GetCreditCardOption {
	return func(o *GetCreditCardOptions) {
		o.Quality = v
		o.enabledSetters["Quality"] = true
	}
}

func (srv *Avatars) getCreditCardParams(optionalSetters ...GetCreditCardOption) map[string]interface{} {
	options := GetCreditCardOptions{}.New()
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
	if options.enabledSetters["Quality"] {
		params["quality"] = options.Quality
	}
	return params
}

// GetCreditCard the credit card endpoint will return you the icon of the
// credit card provider you need. Use width, height and quality arguments to
// change the output settings.
//
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
func (srv *Avatars) GetCreditCard(Code string, optionalSetters ...GetCreditCardOption) (*[]byte, error) {
	r := strings.NewReplacer("{code}", client.EncodePath(Code))
	path := r.Replace("/avatars/credit-cards/{code}")
	params := srv.getCreditCardParams(optionalSetters...)
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "image/png"

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

// GetCreditCardURL the credit card endpoint will return you the icon of the
// credit card provider you need. Use width, height and quality arguments to
// change the output settings.
//
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
// Returns the URL for the resource instead of the content.
func (srv *Avatars) GetCreditCardURL(Code string, optionalSetters ...GetCreditCardOption) (*string, error) {
	r := strings.NewReplacer("{code}", client.EncodePath(Code))
	path := r.Replace("/avatars/credit-cards/{code}")
	params := srv.getCreditCardParams(optionalSetters...)
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

func (srv *Avatars) getFaviconParams(Url string) map[string]interface{} {
	params := map[string]interface{}{}
	params["url"] = Url
	return params
}

// GetFavicon use this endpoint to fetch the favorite icon (AKA favicon) of
// any remote website URL.
//
// This endpoint does not follow HTTP redirects.
func (srv *Avatars) GetFavicon(Url string) (*[]byte, error) {
	path := "/avatars/favicon"
	params := srv.getFaviconParams(Url)
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

// GetFaviconURL use this endpoint to fetch the favorite icon (AKA favicon) of
// any remote website URL.
//
// This endpoint does not follow HTTP redirects.
// Returns the URL for the resource instead of the content.
func (srv *Avatars) GetFaviconURL(Url string) (*string, error) {
	path := "/avatars/favicon"
	params := srv.getFaviconParams(Url)
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

type GetFlagOptions struct {
	Width          int
	Height         int
	Quality        int
	enabledSetters map[string]bool
}

func (options GetFlagOptions) New() *GetFlagOptions {
	options.enabledSetters = map[string]bool{"Width": false, "Height": false, "Quality": false}
	return &options
}

type GetFlagOption func(*GetFlagOptions)

func (srv *Avatars) WithGetFlagWidth(v int) GetFlagOption {
	return func(o *GetFlagOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func (srv *Avatars) WithGetFlagHeight(v int) GetFlagOption {
	return func(o *GetFlagOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}
func (srv *Avatars) WithGetFlagQuality(v int) GetFlagOption {
	return func(o *GetFlagOptions) {
		o.Quality = v
		o.enabledSetters["Quality"] = true
	}
}

func (srv *Avatars) getFlagParams(optionalSetters ...GetFlagOption) map[string]interface{} {
	options := GetFlagOptions{}.New()
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
	if options.enabledSetters["Quality"] {
		params["quality"] = options.Quality
	}
	return params
}

// GetFlag you can use this endpoint to show different country flags icons to
// your users. The code argument receives the 2 letter country code. Use
// width, height and quality arguments to change the output settings. Country
// codes follow the [ISO 3166-1](https://en.wikipedia.org/wiki/ISO_3166-1)
// standard.
//
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
func (srv *Avatars) GetFlag(Code string, optionalSetters ...GetFlagOption) (*[]byte, error) {
	r := strings.NewReplacer("{code}", client.EncodePath(Code))
	path := r.Replace("/avatars/flags/{code}")
	params := srv.getFlagParams(optionalSetters...)
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "image/png"

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

// GetFlagURL you can use this endpoint to show different country flags icons
// to your users. The code argument receives the 2 letter country code. Use
// width, height and quality arguments to change the output settings. Country
// codes follow the [ISO 3166-1](https://en.wikipedia.org/wiki/ISO_3166-1)
// standard.
//
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
// Returns the URL for the resource instead of the content.
func (srv *Avatars) GetFlagURL(Code string, optionalSetters ...GetFlagOption) (*string, error) {
	r := strings.NewReplacer("{code}", client.EncodePath(Code))
	path := r.Replace("/avatars/flags/{code}")
	params := srv.getFlagParams(optionalSetters...)
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

type GetImageOptions struct {
	Width          int
	Height         int
	enabledSetters map[string]bool
}

func (options GetImageOptions) New() *GetImageOptions {
	options.enabledSetters = map[string]bool{"Width": false, "Height": false}
	return &options
}

type GetImageOption func(*GetImageOptions)

func (srv *Avatars) WithGetImageWidth(v int) GetImageOption {
	return func(o *GetImageOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func (srv *Avatars) WithGetImageHeight(v int) GetImageOption {
	return func(o *GetImageOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}

func (srv *Avatars) getImageParams(Url string, optionalSetters ...GetImageOption) map[string]interface{} {
	options := GetImageOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["url"] = Url
	if options.enabledSetters["Width"] {
		params["width"] = options.Width
	}
	if options.enabledSetters["Height"] {
		params["height"] = options.Height
	}
	return params
}

// GetImage use this endpoint to fetch a remote image URL and crop it to any
// image size you want. This endpoint is very useful if you need to crop and
// display remote images in your app or in case you want to make sure a 3rd
// party image is properly served using a TLS protocol.
//
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 400x400px.
//
// This endpoint does not follow HTTP redirects.
func (srv *Avatars) GetImage(Url string, optionalSetters ...GetImageOption) (*[]byte, error) {
	path := "/avatars/image"
	params := srv.getImageParams(Url, optionalSetters...)
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

// GetImageURL use this endpoint to fetch a remote image URL and crop it to
// any image size you want. This endpoint is very useful if you need to crop
// and display remote images in your app or in case you want to make sure a
// 3rd party image is properly served using a TLS protocol.
//
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 400x400px.
//
// This endpoint does not follow HTTP redirects.
// Returns the URL for the resource instead of the content.
func (srv *Avatars) GetImageURL(Url string, optionalSetters ...GetImageOption) (*string, error) {
	path := "/avatars/image"
	params := srv.getImageParams(Url, optionalSetters...)
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

type GetInitialsOptions struct {
	Name           string
	Width          int
	Height         int
	Background     string
	enabledSetters map[string]bool
}

func (options GetInitialsOptions) New() *GetInitialsOptions {
	options.enabledSetters = map[string]bool{"Name": false, "Width": false, "Height": false, "Background": false}
	return &options
}

type GetInitialsOption func(*GetInitialsOptions)

func (srv *Avatars) WithGetInitialsName(v string) GetInitialsOption {
	return func(o *GetInitialsOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Avatars) WithGetInitialsWidth(v int) GetInitialsOption {
	return func(o *GetInitialsOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func (srv *Avatars) WithGetInitialsHeight(v int) GetInitialsOption {
	return func(o *GetInitialsOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}
func (srv *Avatars) WithGetInitialsBackground(v string) GetInitialsOption {
	return func(o *GetInitialsOptions) {
		o.Background = v
		o.enabledSetters["Background"] = true
	}
}

func (srv *Avatars) getInitialsParams(optionalSetters ...GetInitialsOption) map[string]interface{} {
	options := GetInitialsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Width"] {
		params["width"] = options.Width
	}
	if options.enabledSetters["Height"] {
		params["height"] = options.Height
	}
	if options.enabledSetters["Background"] {
		params["background"] = options.Background
	}
	return params
}

// GetInitials use this endpoint to show your user initials avatar icon on
// your website or app. By default, this route will try to print your
// logged-in user name or email initials. You can also overwrite the user name
// if you pass the 'name' parameter. If no name is given and no user is
// logged, an empty avatar will be returned.
//
// You can use the color and background params to change the avatar colors. By
// default, a random theme will be selected. The random theme will persist for
// the user's initials when reloading the same theme will always return for
// the same initials.
//
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
func (srv *Avatars) GetInitials(optionalSetters ...GetInitialsOption) (*[]byte, error) {
	path := "/avatars/initials"
	params := srv.getInitialsParams(optionalSetters...)
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "image/png"

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

// GetInitialsURL use this endpoint to show your user initials avatar icon on
// your website or app. By default, this route will try to print your
// logged-in user name or email initials. You can also overwrite the user name
// if you pass the 'name' parameter. If no name is given and no user is
// logged, an empty avatar will be returned.
//
// You can use the color and background params to change the avatar colors. By
// default, a random theme will be selected. The random theme will persist for
// the user's initials when reloading the same theme will always return for
// the same initials.
//
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
// Returns the URL for the resource instead of the content.
func (srv *Avatars) GetInitialsURL(optionalSetters ...GetInitialsOption) (*string, error) {
	path := "/avatars/initials"
	params := srv.getInitialsParams(optionalSetters...)
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

type GetPhotoOptions struct {
	Width          int
	Height         int
	Quality        int
	Output         string
	Rating         string
	enabledSetters map[string]bool
}

func (options GetPhotoOptions) New() *GetPhotoOptions {
	options.enabledSetters = map[string]bool{"Width": false, "Height": false, "Quality": false, "Output": false, "Rating": false}
	return &options
}

type GetPhotoOption func(*GetPhotoOptions)

func (srv *Avatars) WithGetPhotoWidth(v int) GetPhotoOption {
	return func(o *GetPhotoOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func (srv *Avatars) WithGetPhotoHeight(v int) GetPhotoOption {
	return func(o *GetPhotoOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}
func (srv *Avatars) WithGetPhotoQuality(v int) GetPhotoOption {
	return func(o *GetPhotoOptions) {
		o.Quality = v
		o.enabledSetters["Quality"] = true
	}
}
func (srv *Avatars) WithGetPhotoOutput(v string) GetPhotoOption {
	return func(o *GetPhotoOptions) {
		o.Output = v
		o.enabledSetters["Output"] = true
	}
}
func (srv *Avatars) WithGetPhotoRating(v string) GetPhotoOption {
	return func(o *GetPhotoOptions) {
		o.Rating = v
		o.enabledSetters["Rating"] = true
	}
}

func (srv *Avatars) getPhotoParams(optionalSetters ...GetPhotoOption) map[string]interface{} {
	options := GetPhotoOptions{}.New()
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
	if options.enabledSetters["Quality"] {
		params["quality"] = options.Quality
	}
	if options.enabledSetters["Output"] {
		params["output"] = options.Output
	}
	if options.enabledSetters["Rating"] {
		params["rating"] = options.Rating
	}
	return params
}

// GetPhoto returns the best available profile photo for the currently
// authenticated user. The endpoint tries each source in priority order and
// returns the first successful result: Gravatar, Libavatar, Appwrite
// Initials, built-in static fallback file.
func (srv *Avatars) GetPhoto(optionalSetters ...GetPhotoOption) (*[]byte, error) {
	path := "/avatars/photo"
	params := srv.getPhotoParams(optionalSetters...)
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

// GetPhotoURL returns the best available profile photo for the currently
// authenticated user. The endpoint tries each source in priority order and
// returns the first successful result: Gravatar, Libavatar, Appwrite
// Initials, built-in static fallback file.
// Returns the URL for the resource instead of the content.
func (srv *Avatars) GetPhotoURL(optionalSetters ...GetPhotoOption) (*string, error) {
	path := "/avatars/photo"
	params := srv.getPhotoParams(optionalSetters...)
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

type GetQROptions struct {
	Size           int
	Margin         int
	Download       bool
	enabledSetters map[string]bool
}

func (options GetQROptions) New() *GetQROptions {
	options.enabledSetters = map[string]bool{"Size": false, "Margin": false, "Download": false}
	return &options
}

type GetQROption func(*GetQROptions)

func (srv *Avatars) WithGetQRSize(v int) GetQROption {
	return func(o *GetQROptions) {
		o.Size = v
		o.enabledSetters["Size"] = true
	}
}
func (srv *Avatars) WithGetQRMargin(v int) GetQROption {
	return func(o *GetQROptions) {
		o.Margin = v
		o.enabledSetters["Margin"] = true
	}
}
func (srv *Avatars) WithGetQRDownload(v bool) GetQROption {
	return func(o *GetQROptions) {
		o.Download = v
		o.enabledSetters["Download"] = true
	}
}

func (srv *Avatars) getQRParams(Text string, optionalSetters ...GetQROption) map[string]interface{} {
	options := GetQROptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["text"] = Text
	if options.enabledSetters["Size"] {
		params["size"] = options.Size
	}
	if options.enabledSetters["Margin"] {
		params["margin"] = options.Margin
	}
	if options.enabledSetters["Download"] {
		params["download"] = options.Download
	}
	return params
}

// GetQR converts a given plain text to a QR code image. You can use the query
// parameters to change the size and style of the resulting image.
func (srv *Avatars) GetQR(Text string, optionalSetters ...GetQROption) (*[]byte, error) {
	path := "/avatars/qr"
	params := srv.getQRParams(Text, optionalSetters...)
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "image/png"

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

// GetQRURL converts a given plain text to a QR code image. You can use the
// query parameters to change the size and style of the resulting image.
// Returns the URL for the resource instead of the content.
func (srv *Avatars) GetQRURL(Text string, optionalSetters ...GetQROption) (*string, error) {
	path := "/avatars/qr"
	params := srv.getQRParams(Text, optionalSetters...)
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

type GetScreenshotOptions struct {
	Headers        interface{}
	ViewportWidth  int
	ViewportHeight int
	Scale          float64
	Theme          string
	UserAgent      string
	Fullpage       bool
	Locale         string
	Timezone       string
	Latitude       float64
	Longitude      float64
	Accuracy       float64
	Touch          bool
	Permissions    []string
	Sleep          int
	Width          int
	Height         int
	Quality        int
	Output         string
	enabledSetters map[string]bool
}

func (options GetScreenshotOptions) New() *GetScreenshotOptions {
	options.enabledSetters = map[string]bool{"Headers": false, "ViewportWidth": false, "ViewportHeight": false, "Scale": false, "Theme": false, "UserAgent": false, "Fullpage": false, "Locale": false, "Timezone": false, "Latitude": false, "Longitude": false, "Accuracy": false, "Touch": false, "Permissions": false, "Sleep": false, "Width": false, "Height": false, "Quality": false, "Output": false}
	return &options
}

type GetScreenshotOption func(*GetScreenshotOptions)

func (srv *Avatars) WithGetScreenshotHeaders(v interface{}) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Headers = v
		o.enabledSetters["Headers"] = true
	}
}
func (srv *Avatars) WithGetScreenshotViewportWidth(v int) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.ViewportWidth = v
		o.enabledSetters["ViewportWidth"] = true
	}
}
func (srv *Avatars) WithGetScreenshotViewportHeight(v int) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.ViewportHeight = v
		o.enabledSetters["ViewportHeight"] = true
	}
}
func (srv *Avatars) WithGetScreenshotScale(v float64) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Scale = v
		o.enabledSetters["Scale"] = true
	}
}
func (srv *Avatars) WithGetScreenshotTheme(v string) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Theme = v
		o.enabledSetters["Theme"] = true
	}
}
func (srv *Avatars) WithGetScreenshotUserAgent(v string) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.UserAgent = v
		o.enabledSetters["UserAgent"] = true
	}
}
func (srv *Avatars) WithGetScreenshotFullpage(v bool) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Fullpage = v
		o.enabledSetters["Fullpage"] = true
	}
}
func (srv *Avatars) WithGetScreenshotLocale(v string) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *Avatars) WithGetScreenshotTimezone(v string) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Timezone = v
		o.enabledSetters["Timezone"] = true
	}
}
func (srv *Avatars) WithGetScreenshotLatitude(v float64) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Latitude = v
		o.enabledSetters["Latitude"] = true
	}
}
func (srv *Avatars) WithGetScreenshotLongitude(v float64) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Longitude = v
		o.enabledSetters["Longitude"] = true
	}
}
func (srv *Avatars) WithGetScreenshotAccuracy(v float64) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Accuracy = v
		o.enabledSetters["Accuracy"] = true
	}
}
func (srv *Avatars) WithGetScreenshotTouch(v bool) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Touch = v
		o.enabledSetters["Touch"] = true
	}
}
func (srv *Avatars) WithGetScreenshotPermissions(v []string) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *Avatars) WithGetScreenshotSleep(v int) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Sleep = v
		o.enabledSetters["Sleep"] = true
	}
}
func (srv *Avatars) WithGetScreenshotWidth(v int) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func (srv *Avatars) WithGetScreenshotHeight(v int) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}
func (srv *Avatars) WithGetScreenshotQuality(v int) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Quality = v
		o.enabledSetters["Quality"] = true
	}
}
func (srv *Avatars) WithGetScreenshotOutput(v string) GetScreenshotOption {
	return func(o *GetScreenshotOptions) {
		o.Output = v
		o.enabledSetters["Output"] = true
	}
}

func (srv *Avatars) getScreenshotParams(Url string, optionalSetters ...GetScreenshotOption) map[string]interface{} {
	options := GetScreenshotOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["url"] = Url
	if options.enabledSetters["Headers"] {
		params["headers"] = options.Headers
	}
	if options.enabledSetters["ViewportWidth"] {
		params["viewportWidth"] = options.ViewportWidth
	}
	if options.enabledSetters["ViewportHeight"] {
		params["viewportHeight"] = options.ViewportHeight
	}
	if options.enabledSetters["Scale"] {
		params["scale"] = options.Scale
	}
	if options.enabledSetters["Theme"] {
		params["theme"] = options.Theme
	}
	if options.enabledSetters["UserAgent"] {
		params["userAgent"] = options.UserAgent
	}
	if options.enabledSetters["Fullpage"] {
		params["fullpage"] = options.Fullpage
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["Timezone"] {
		params["timezone"] = options.Timezone
	}
	if options.enabledSetters["Latitude"] {
		params["latitude"] = options.Latitude
	}
	if options.enabledSetters["Longitude"] {
		params["longitude"] = options.Longitude
	}
	if options.enabledSetters["Accuracy"] {
		params["accuracy"] = options.Accuracy
	}
	if options.enabledSetters["Touch"] {
		params["touch"] = options.Touch
	}
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["Sleep"] {
		params["sleep"] = options.Sleep
	}
	if options.enabledSetters["Width"] {
		params["width"] = options.Width
	}
	if options.enabledSetters["Height"] {
		params["height"] = options.Height
	}
	if options.enabledSetters["Quality"] {
		params["quality"] = options.Quality
	}
	if options.enabledSetters["Output"] {
		params["output"] = options.Output
	}
	return params
}

// GetScreenshot use this endpoint to capture a screenshot of any website URL.
// This endpoint uses a headless browser to render the webpage and capture it
// as an image.
//
// You can configure the browser viewport size, theme, user agent,
// geolocation, permissions, and more. Capture either just the viewport or the
// full page scroll.
//
// When width and height are specified, the image is resized accordingly. If
// both dimensions are 0, the API provides an image at original size. If
// dimensions are not specified, the default viewport size is 1280x720px.
func (srv *Avatars) GetScreenshot(Url string, optionalSetters ...GetScreenshotOption) (*[]byte, error) {
	path := "/avatars/screenshots"
	params := srv.getScreenshotParams(Url, optionalSetters...)
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "image/png"

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

// GetScreenshotURL use this endpoint to capture a screenshot of any website
// URL. This endpoint uses a headless browser to render the webpage and
// capture it as an image.
//
// You can configure the browser viewport size, theme, user agent,
// geolocation, permissions, and more. Capture either just the viewport or the
// full page scroll.
//
// When width and height are specified, the image is resized accordingly. If
// both dimensions are 0, the API provides an image at original size. If
// dimensions are not specified, the default viewport size is 1280x720px.
// Returns the URL for the resource instead of the content.
func (srv *Avatars) GetScreenshotURL(Url string, optionalSetters ...GetScreenshotOption) (*string, error) {
	path := "/avatars/screenshots"
	params := srv.getScreenshotParams(Url, optionalSetters...)
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
