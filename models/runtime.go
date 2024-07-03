package models


// Runtime Model
type Runtime struct {
    // Runtime ID.
    Id string `json:"$id"`
    // Runtime Name.
    Name string `json:"name"`
    // Runtime version.
    Version string `json:"version"`
    // Base Docker image used to build the runtime.
    Base string `json:"base"`
    // Image name of Docker Hub.
    Image string `json:"image"`
    // Name of the logo image.
    Logo string `json:"logo"`
    // List of supported architectures.
    Supports []interface{} `json:"supports"`

}
