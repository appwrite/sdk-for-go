package appwrite

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client is the client struct to access Appwrite services
type Client struct {
	client     *http.Client
	endpoint   string
	headers    map[string]string
	selfSigned bool
}

// SetEndpoint sets the default endpoint to which the Client connects to
func (clt *Client) SetEndpoint(endpoint string) {
	clt.endpoint = endpoint
}

// SetSelfSigned sets the condition that specify if the Client should allow connections to a server using a self-signed certificate
func (clt *Client) SetSelfSigned(status bool) {
	clt.selfSigned = status
}

// AddHeader add a new custom header that the Client should send on each request
func (clt *Client) AddHeader(key string, value string) {
	clt.headers[key] = value
}

// Your project ID
func (clt *Client) SetProject(value string) {
	clt.headers["X-Appwrite-Project"] = value
}

// Your secret API key
func (clt *Client) SetKey(value string) {
	clt.headers["X-Appwrite-Key"] = value
}

func (clt *Client) SetLocale(value string) {
	clt.headers["X-Appwrite-Locale"] = value
}

func (clt *Client) SetMode(value string) {
	clt.headers["X-Appwrite-Mode"] = value
}

// Call an API using Client
func (clt *Client) Call(method string, path string, headers map[string]interface{}, params map[string]interface{}) (map[string]interface{}, error) {
	if clt.client == nil {
		// Create HTTP client
		clt.client = &http.Client{}
	}

	if clt.selfSigned {
		// Allow self signed requests
	}

	isGet := strings.ToUpper(method) == "GET"

	urlPath := clt.endpoint + path

	var http_req *http.Request
	var reqBody []byte
	if isGet {
		http_req, _ = http.NewRequest(strings.ToUpper(method), urlPath, strings.NewReader(string(reqBody)))
		q := http_req.URL.Query()
		for key, val := range params {
			q.Add(key, ToString(val))
		}
		http_req.URL.RawQuery = q.Encode()
	} else {
		reqBody, _ = json.Marshal(params)
		http_req, _ = http.NewRequest(strings.ToUpper(method), urlPath, strings.NewReader(string(reqBody)))
	}

	// set client headers
	for key, val := range clt.headers {
		http_req.Header.Set(key, val)
	}

	// set custom headers
	for key, val := range headers {
		http_req.Header.Set(key, ToString(val))
	}

	// submit the request
	resp, err := clt.client.Do(http_req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response data
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var jsonResp map[string]interface{}
	json.Unmarshal(respData, &jsonResp)
	return jsonResp, nil
}
