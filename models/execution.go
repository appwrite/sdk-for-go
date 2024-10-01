package models

import (
    "encoding/json"
    "errors"
    "github.com/appwrite/sdk-for-go/payload"
)

// Execution Model
type Execution struct {
    // Execution ID.
    Id string `json:"$id"`
    // Execution creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Execution upate date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Execution roles.
    Permissions []string `json:"$permissions"`
    // Function ID.
    FunctionId string `json:"functionId"`
    // The trigger that caused the function to execute. Possible values can be:
    // `http`, `schedule`, or `event`.
    Trigger string `json:"trigger"`
    // The status of the function execution. Possible values can be: `waiting`,
    // `processing`, `completed`, or `failed`.
    Status string `json:"status"`
    // HTTP request method type.
    RequestMethod string `json:"requestMethod"`
    // HTTP request path and query.
    RequestPath string `json:"requestPath"`
    // HTTP response headers as a key-value object. This will return only
    // whitelisted headers. All headers are returned if execution is created as
    // synchronous.
    RequestHeaders []Headers `json:"requestHeaders"`
    // HTTP response status code.
    ResponseStatusCode int `json:"responseStatusCode"`
    // HTTP response body. This will return empty unless execution is created as
    // synchronous.
    ResponseBody *payload.Payload `json:"responseBody"`
    // HTTP response headers as a key-value object. This will return only
    // whitelisted headers. All headers are returned if execution is created as
    // synchronous.
    ResponseHeaders []Headers `json:"responseHeaders"`
    // Function logs. Includes the last 4,000 characters. This will return an
    // empty string unless the response is returned using an API key or as part of
    // a webhook payload.
    Logs string `json:"logs"`
    // Function errors. Includes the last 4,000 characters. This will return an
    // empty string unless the response is returned using an API key or as part of
    // a webhook payload.
    Errors string `json:"errors"`
    // Function execution duration in seconds.
    Duration float64 `json:"duration"`
    // The scheduled time for execution. If left empty, execution will be queued
    // immediately.
    ScheduledAt string `json:"scheduledAt"`

    // Used by Decode() method
    data []byte
}

func (model Execution) New(data []byte) *Execution {
    model.data = data
    return &model
}

func (model *Execution) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
