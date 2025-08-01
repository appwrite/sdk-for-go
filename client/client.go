package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"runtime"

	"github.com/appwrite/sdk-for-go/file"
)

const (
	fileNameKey      = "file"
	defaultTimeout   = 10 * time.Second
	defaultChunkSize = 5 * 1024 * 1024
)

// AppwriteError represents an error of a client request
type AppwriteError struct {
	statusCode int
	message    string
	response   string
}

// ClientResponse - represents the client response
type ClientResponse struct {
	Status     string
	StatusCode int
	Header     http.Header
	Result     interface{}
	Type	   string
}

func (ce *AppwriteError) Error() string {
	return ce.message
}

func (ce *AppwriteError) GetMessage() string {
	return ce.message
}

func (ce *AppwriteError) GetStatusCode() int {
	return ce.statusCode
}

func (ce *AppwriteError) GetResponse() string {
	return ce.response
}

// Client is the client struct to access Appwrite  services
type Client struct {
	Client     *http.Client
	Headers    map[string]string
	Endpoint   string
	Timeout    time.Duration
	SelfSigned bool
	ChunkSize  int64
}

// Initialize a new Appwrite client with a given timeout
func New(optionalSetters ...ClientOption) Client {
	headers := map[string]string{
		"X-Appwrite-Response-Format" : "1.7.0",
		"user-agent" : fmt.Sprintf("AppwriteGoSDK/0.9.0 (%s; %s)", runtime.GOOS, runtime.GOARCH),
		"x-sdk-name": "Go",
		"x-sdk-platform": "server",
		"x-sdk-language": "go",
		"x-sdk-version": "0.9.0",
	}
	httpClient, err := GetDefaultClient(defaultTimeout)
	if err != nil {
		panic(err)
	}

	client := Client{
		Endpoint:  "https://cloud.appwrite.io/v1",
		Client:    httpClient,
		Timeout:   defaultTimeout,
		Headers:   headers,
		ChunkSize: defaultChunkSize,
	}

	for _, opt := range optionalSetters {
		err = opt(&client)
		if err != nil {
			panic(err)
		}
	}

	return client
}

func (client *Client) String() string {
	return fmt.Sprintf("%s\n%s\n%v", client.Endpoint, client.Headers, client.Timeout)
}

func GetDefaultClient(timeout time.Duration) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	return &http.Client{
		Jar:     jar,
		Timeout: timeout,
	}, nil
}

type ClientOption func(*Client) error

// AddHeader add a new custom header that the Client should send on each request
func (client *Client) AddHeader(key string, value string) {
	client.Headers[key] = value
}

func isFileUpload(headers map[string]interface{}) bool {
	contentType, ok := headers["content-type"].(string)
	if ok {
		return strings.Contains(strings.ToLower(contentType), "multipart/form-data")
	}
	return false
}

func (client *Client) FileUpload(url string, headers map[string]interface{}, params map[string]interface{}, paramName string, uploadId string) (*ClientResponse, error) {
	inputFile, ok := params[paramName].(file.InputFile)
	if !ok {
		msg := fmt.Sprintf("invalid input file. params[%s] must be of type file.InputFile", paramName)
		return nil, errors.New(msg)
	}

	file, err := os.Open(inputFile.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	inputFile.Data = make([]byte, client.ChunkSize)

	var result *ClientResponse

	numChunks := fileInfo.Size() / client.ChunkSize
	if fileInfo.Size()%client.ChunkSize != 0 {
		numChunks++
	}
	var currentChunk int64 = 0
	if uploadId != "" {
		resp, err := client.Call("GET", url+"/"+uploadId, nil, nil)
		if err == nil {
			currentChunk = int64(resp.Result.(map[string]interface{})["chunksUploaded"].(float64))
		}
	}

	if fileInfo.Size() <= client.ChunkSize {
		if uploadId != "" {
			headers["x-appwrite-id"] = uploadId
		}
		inputFile.Data = make([]byte, fileInfo.Size())
		_, err := file.Read(inputFile.Data)
		if err != nil && err != io.EOF {
			return nil, err
		}
		params[paramName] = inputFile

		result, err = client.Call("POST", url, headers, params)
		if err != nil {
			return nil, err
		}

		var parsed map[string]interface{}
		if strings.HasPrefix(result.Type, "application/json") {
			err = json.Unmarshal([]byte(result.Result.(string)), &parsed)
			if err == nil {
				uploadId, _ = parsed["$id"].(string)
			}
		}

		return result, nil
	}

	for i := currentChunk; i < numChunks; i++ {
		chunkSize := client.ChunkSize
		offset := int64(i) * chunkSize
		if i == numChunks-1 {
			chunkSize = fileInfo.Size() - offset
			inputFile.Data = make([]byte, chunkSize)
		}
		_, err := file.ReadAt(inputFile.Data, offset)
		if err != nil && err != io.EOF {
			return nil, err
		}
		params[paramName] = inputFile
		if uploadId != "" {
			headers["x-appwrite-id"] = uploadId
		}
		totalSize := fileInfo.Size()
		start := offset
		end := offset + client.ChunkSize - 1
		if end >= totalSize {
			end = totalSize - 1
		}
		headers["content-range"] = fmt.Sprintf("bytes %d-%d/%d", start, end, totalSize)
		result, err = client.Call("POST", url, headers, params)
		if err != nil {
			return nil, err
		}

		var parsed map[string]interface{}
		if strings.HasPrefix(result.Type, "application/json") {
			err = json.Unmarshal([]byte(result.Result.(string)), &parsed)
			if err == nil {
				uploadId, _ = parsed["$id"].(string)
			}
		}
	}
	return result, nil
}

// Call an API using Client
func (client *Client) Call(method string, path string, headers map[string]interface{}, params map[string]interface{}) (*ClientResponse, error) {
	if client.Client == nil {
		// Create HTTP client
		httpClient, err := GetDefaultClient(client.Timeout)
		if err != nil {
			panic(err)
		}
		client.Client = httpClient
	}

	if client.SelfSigned {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	urlPath := client.Endpoint + path
	isGet := strings.ToUpper(method) == "GET"
	isPost := strings.ToUpper(method) == "POST"
	isJsonRequest := headers["content-type"] == "application/json"
	isFileUpload := isFileUpload(headers)

	var req *http.Request
	var err error
	if isFileUpload {
		if !isPost {
			return nil, errors.New("fileupload needs POST Request")
		}
		var body bytes.Buffer
		writer := multipart.NewWriter(&body)
		for key, val := range params {
			if file, ok := val.(file.InputFile); ok {
				fileName := file.Name
				fileData := file.Data
				fw, err := writer.CreateFormFile(key, fileName)
				if err != nil {
					return nil, err
				}
				_, err = io.Copy(fw, bytes.NewReader(fileData))
				if err != nil {
					return nil, err
				}
				delete(params, key)
			}
		}
		flatParams := make(map[string]string)
		flatten(params, "", &flatParams)
		for key, val := range flatParams {
			err = writer.WriteField(key, val)
			if err != nil {
				return nil, err
			}
		}
		err = writer.Close()
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, urlPath, &body)
		if err != nil {
			return nil, err
		}
		headers["content-type"] = writer.FormDataContentType()
	} else {
		if !isGet {
			var reqBody *strings.Reader
			if isJsonRequest {
				json, err := json.Marshal(params)
				if err != nil {
					return nil, err
				}
				reqBody = strings.NewReader(string(json))
			} else {
				frm := url.Values{}
				for key, val := range params {
					frm.Add(key, toString(val))
				}
				reqBody = strings.NewReader(frm.Encode())
			}
			// Create and modify HTTP request before sending
			req, err = http.NewRequest(method, urlPath, reqBody)
			if err != nil {
				return nil, err
			}
		} else {
			req, err = http.NewRequest(method, urlPath, nil)
			if err != nil {
				return nil, err
			}
		}

		if isGet {
			q := req.URL.Query()
			for key, val := range params {
				rt := reflect.TypeOf(val)
				switch rt.Kind() {
				case reflect.Array:
				case reflect.Slice:
					arr := reflect.ValueOf(val)
					for i := 0; i < arr.Len(); i++ {
						q.Add(fmt.Sprintf("%s[]", key), toString(arr.Index(i)))
					}
				default:
					if strVal := toString(val); strVal != "" {
						q.Add(key, strVal)
					}
				}
			}
			rawQuery := q.Encode()
			req.URL.RawQuery = rawQuery
		}
	}

	// Set Custom headers
	for key, val := range headers {
		req.Header.Set(key, toString(val))
	}

	// Set Client headers
	for key, val := range client.Headers {
		req.Header.Set(key, toString(val))
	}

	// Make request
	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}

	// Handle response
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	contentType := resp.Header.Get("content-type")
	var isJson = strings.HasPrefix(contentType, "application/json")
	if isJson {
		if resp.StatusCode < 200 || resp.StatusCode > 399 {
			var jsonResponse map[string]interface{}
			json.Unmarshal(responseData, &jsonResponse)
			message, ok := jsonResponse["message"].(string)
			if !ok {
				message = "N/A"
			}
			return nil, &AppwriteError{
				statusCode: resp.StatusCode,
				message:    message,
				response:   string(responseData),
			}
		}
		return &ClientResponse{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			Header:     resp.Header,
			Result:     string(responseData),
			Type:	   contentType,
		}, nil
	}

	if resp.StatusCode < 200 || resp.StatusCode > 399 {
		return nil, &AppwriteError{
			statusCode: resp.StatusCode,
			message:    string(responseData),
			response:   string(responseData),
		}
	}
	return &ClientResponse{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Result:     responseData,
		Type:	   contentType,
	}, nil
}

// toString changes arg to string
func toString(arg interface{}) string {
	var tmp = reflect.Indirect(reflect.ValueOf(arg)).Interface()
	switch v := tmp.(type) {
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case string:
		return v
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case reflect.Value:
		return toString(v.Interface())
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%s", v)
	}
}


// flatten recursively flattens params into a map[string]string and writes it to result
func flatten(params interface{}, prefix string, result *map[string]string) error {
	if result == nil {
		return errors.New("result is nil")
	}

	paramsType := reflect.TypeOf(params)
	if paramsType.Kind() == reflect.Ptr {
		paramsType = paramsType.Elem()
	}
	switch paramsType.Kind() {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		arr := reflect.Indirect(reflect.ValueOf(params))
		for i := 0; i < arr.Len(); i++ {
			currentPrefix := fmt.Sprintf("%s[%d]", prefix, i)
			err := flatten(arr.Index(i).Interface(), currentPrefix, result)
			if err != nil {
				return err
			}
		}
	case reflect.Map:
		m := reflect.Indirect(reflect.ValueOf(params))
		for _, key := range m.MapKeys() {
			var currentPrefix string
			if prefix == "" {
				currentPrefix = toString(key)
			} else {
				currentPrefix = fmt.Sprintf("%s[%s]", prefix, toString(key))
			}
			err := flatten(m.MapIndex(key).Interface(), currentPrefix, result)
			if err != nil {
				return err
			}
		}
		default:
			if prefix == "" {
				return fmt.Errorf("prefix is empty for %s", params)
			}
			(*result)[prefix] = toString(params)
	}
	return nil
}
