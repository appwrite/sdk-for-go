package models

import (
    "encoding/json"
    "errors"
)

// Bucket Model
type Bucket struct {
    // Bucket ID.
    Id string `json:"$id"`
    // Bucket creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Bucket update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Bucket permissions. [Learn more about
    // permissions](https://appwrite.io/docs/permissions).
    Permissions []string `json:"$permissions"`
    // Whether file-level security is enabled. [Learn more about
    // permissions](https://appwrite.io/docs/permissions).
    FileSecurity bool `json:"fileSecurity"`
    // Bucket name.
    Name string `json:"name"`
    // Bucket enabled.
    Enabled bool `json:"enabled"`
    // Maximum file size supported.
    MaximumFileSize int `json:"maximumFileSize"`
    // Allowed file extensions.
    AllowedFileExtensions []string `json:"allowedFileExtensions"`
    // Compression algorithm choosen for compression. Will be one of none,
    // [gzip](https://en.wikipedia.org/wiki/Gzip), or
    // [zstd](https://en.wikipedia.org/wiki/Zstd).
    Compression string `json:"compression"`
    // Bucket is encrypted.
    Encryption bool `json:"encryption"`
    // Virus scanning is enabled.
    Antivirus bool `json:"antivirus"`
    // Image transformations are enabled.
    Transformations bool `json:"transformations"`

    // Used by Decode() method
    data []byte
}

func (model Bucket) New(data []byte) *Bucket {
    model.data = data
    return &model
}

func (model *Bucket) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}