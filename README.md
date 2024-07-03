# Appwrite Go SDK

![License](https://img.shields.io/github/license/appwrite/sdk-for-go.svg?style=flat-square)
![Version](https://img.shields.io/badge/api%20version-1.5.7-blue.svg?style=flat-square)
[![Build Status](https://img.shields.io/travis/com/appwrite/sdk-generator?style=flat-square)](https://travis-ci.com/appwrite/sdk-generator)
[![Twitter Account](https://img.shields.io/twitter/follow/appwrite?color=00acee&label=twitter&style=flat-square)](https://twitter.com/appwrite)
[![Discord](https://img.shields.io/discord/564160730845151244?label=discord&style=flat-square)](https://appwrite.io/discord)

**This SDK is compatible with Appwrite server version 1.6.x. For older versions, please check [previous releases](https://github.com/appwrite/sdk-for-go/releases).**

Appwrite is an open-source backend as a service server that abstract and simplify complex and repetitive development tasks behind a very simple to use REST API. Appwrite aims to help you develop your apps faster and in a more secure way. Use the Go SDK to integrate your app with the Appwrite server to easily start interacting with all of Appwrite backend APIs and tools. For full API documentation and tutorials go to [https://appwrite.io/docs](https://appwrite.io/docs)

![Appwrite](https://github.com/appwrite/appwrite/raw/main/public/images/github.png)

## Installation

To install using `go get`:

```bash
go get github.com/appwrite/sdk-for-go
```

## Testing the SDK

* clone this repo.
* create a project and within this project a collection.
* configure the documents in the collection to have a _key_ = __hello__.
* Then inject these environment variables:

  ```bash
  export YOUR_ENDPOINT=https://appwrite.io/v1  
  export YOUR_PROJECT_ID=6…8  
  export YOUR_KEY="7055781…cd95"  
  export COLLECTION_ID=616a095b20180  
  ```

Create `main.go` file with:

```go
package main

import (
	"log"
	"os"
	"time"

	"github.com/appwrite/sdk-for-go/sdk-for-go"
)

func main() {
	client := sdk-for-go.NewClient(10 * time.Second)
	client.SetEndpoint(os.Getenv("YOUR_ENDPOINT"))
	client.SetProject(os.Getenv("YOUR_PROJECT_ID"))
	client.SetKey(os.Getenv("YOUR_KEY"))

	db := sdk-for-go.NewDatabase(client)
	data := map[string]string{
		"hello": "world",
	}
	var EmptyArray = []interface{}{}
	doc, err := db.CreateDocument(
		os.Getenv("COLLECTION_ID"),
		data,
		EmptyArray,
		EmptyArray,
		"",
		"",
		"",
	)
	if err != nil {
		log.Printf("Error creating document: %v", err)
	}
	log.Printf("Created document: %v", doc)
}
```

* After that, run the following

  > % go run main.go  
  > 2021/10/16 03:41:17 Created document: map[$collection:616a095b20180 $id:616a2dbd4df16 $permissions:map[read:[] write:[]] hello:world]  


## Contribution

This library is auto-generated by Appwrite custom [SDK Generator](https://github.com/appwrite/sdk-generator). To learn more about how you can help us improve this SDK, please check the [contribution guide](https://github.com/appwrite/sdk-generator/blob/master/CONTRIBUTING.md) before sending a pull-request.

## License

Please see the [BSD-3-Clause license](https://raw.githubusercontent.com/appwrite/appwrite/master/LICENSE) file for more information.
