package models


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
    Permissions []interface{} `json:"$permissions"`
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
    AllowedFileExtensions []interface{} `json:"allowedFileExtensions"`
    // Compression algorithm choosen for compression. Will be one of none,
    // [gzip](https://en.wikipedia.org/wiki/Gzip), or
    // [zstd](https://en.wikipedia.org/wiki/Zstd).
    Compression string `json:"compression"`
    // Bucket is encrypted.
    Encryption bool `json:"encryption"`
    // Virus scanning is enabled.
    Antivirus bool `json:"antivirus"`

}
