package models


// Variable Model
type Variable struct {
    // Variable ID.
    Id string `json:"$id"`
    // Variable creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Variable creation date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Variable key.
    Key string `json:"key"`
    // Variable value.
    Value string `json:"value"`
    // Service to which the variable belongs. Possible values are "project",
    // "function"
    ResourceType string `json:"resourceType"`
    // ID of resource to which the variable belongs. If resourceType is "project",
    // it is empty. If resourceType is "function", it is ID of the function.
    ResourceId string `json:"resourceId"`

}
