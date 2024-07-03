package models


// Document Model
type Document struct {
    // Document ID.
    Id string `json:"$id"`
    // Collection ID.
    CollectionId string `json:"$collectionId"`
    // Database ID.
    DatabaseId string `json:"$databaseId"`
    // Document creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Document update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Document permissions. [Learn more about
    // permissions](https://appwrite.io/docs/permissions).
    Permissions []interface{} `json:"$permissions"`

    // Additional properties
    data interface{}
}
