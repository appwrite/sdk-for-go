package appwrite

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// Database service
type Database struct {
	client Client
}

func NewDatabase(clt Client) Database {
	service := Database{
		client: clt,
	}

	return service
}

// ListCollections get a list of all the user collections. You can use the
// query params to filter your results. On admin mode, this endpoint will
// return a list of all of the project collections. [Learn more about
// different API modes](/docs/admin).
func (srv *Database) ListCollections(Search string, Limit int, Offset int, OrderType string) (map[string]interface{}, error) {
	path := "/database/collections"

	params := map[string]interface{}{
		"search":    Search,
		"limit":     Limit,
		"offset":    Offset,
		"orderType": OrderType,
	}

	header := map[string]interface{}{
		"content-type": "application/json",
	}

	return srv.client.Call("GET", path, header, params)
}

// CreateCollection create a new Collection.
func (srv *Database) CreateCollection(CollectionId string, Name string, Permission string, Read []string, Write []string) (map[string]interface{}, error) {
	path := "/database/collections"

	params := map[string]interface{}{
		"collectionId": CollectionId,
		"name":         Name,
		"read":         Read,
		"write":        Write,
		"permission":   Permission,
	}

	header := map[string]interface{}{
		"content-type": "application/json",
	}

	return srv.client.Call("POST", path, header, params)
}

// GetCollection get collection by its unique ID. This endpoint response
// returns a JSON object with the collection metadata.
func (srv *Database) GetCollection(CollectionId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}")

	params := map[string]interface{}{}

	return srv.client.Call("GET", path, nil, params)
}

// UpdateCollection update collection by its unique ID.
func (srv *Database) UpdateCollection(CollectionId string, Name string, Read []interface{}, Write []interface{}, Rules []interface{}) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}")

	params := map[string]interface{}{
		"name":  Name,
		"read":  Read,
		"write": Write,
		"rules": Rules,
	}

	return srv.client.Call("PUT", path, nil, params)
}

// DeleteCollection delete a collection by its unique ID. Only users with
// write permissions have access to delete this resource.
func (srv *Database) DeleteCollection(CollectionId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}")

	params := map[string]interface{}{}

	return srv.client.Call("DELETE", path, nil, params)
}

// ListDocuments get a list of all the user documents. You can use the query
// params to filter your results. On admin mode, this endpoint will return a
// list of all of the project documents. [Learn more about different API
// modes](/docs/admin).
func (srv *Database) ListDocuments(CollectionId string, Filters []interface{}, Offset int, Limit int, OrderField string, OrderType string, OrderCast string, Search string, First int, Last int) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/documents")

	params := map[string]interface{}{
		"filters":     Filters,
		"offset":      Offset,
		"limit":       Limit,
		"order-field": OrderField,
		"order-type":  OrderType,
		"order-cast":  OrderCast,
		"search":      Search,
		"first":       First,
		"last":        Last,
	}

	return srv.client.Call("GET", path, nil, params)
}

// CreateDocument create a new Document.
func (srv *Database) CreateDocument(CollectionId string, Data interface{}, Read []interface{}, Write []interface{}, ParentDocument string, ParentProperty string, ParentPropertyType string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId)
	path := r.Replace("/database/collections/{collectionId}/documents")

	params := map[string]interface{}{
		"data":               Data,
		"read":               Read,
		"write":              Write,
		"parentDocument":     ParentDocument,
		"parentProperty":     ParentProperty,
		"parentPropertyType": ParentPropertyType,
	}

	return srv.client.Call("POST", path, nil, params)
}

// GetDocument get document by its unique ID. This endpoint response returns a
// JSON object with the document data.
func (srv *Database) GetDocument(CollectionId string, DocumentId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/database/collections/{collectionId}/documents/{documentId}")

	params := map[string]interface{}{}

	return srv.client.Call("GET", path, nil, params)
}

// UpdateDocument
func (srv *Database) UpdateDocument(CollectionId string, DocumentId string, Data interface{}, Read []interface{}, Write []interface{}) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/database/collections/{collectionId}/documents/{documentId}")

	params := map[string]interface{}{
		"data":  Data,
		"read":  Read,
		"write": Write,
	}

	return srv.client.Call("PATCH", path, nil, params)
}

// DeleteDocument delete document by its unique ID. This endpoint deletes only
// the parent documents, his attributes and relations to other documents.
// Child documents **will not** be deleted.
func (srv *Database) DeleteDocument(CollectionId string, DocumentId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{collectionId}", CollectionId, "{documentId}", DocumentId)
	path := r.Replace("/database/collections/{collectionId}/documents/{documentId}")

	params := map[string]interface{}{}

	return srv.client.Call("DELETE", path, nil, params)
}

func (srv *Database) CreateStringAttribute(CollectionId string, key string, size int, required bool, xdefault Optional[string], isArray Optional[bool]) (map[string]interface{}, error) {
	Type := "string"
	path := "/database/collections/{collectionId}/attributes/{type}"
	r := strings.NewReplacer("{collectionId}", CollectionId, "{type}", Type)
	path = r.Replace(path)

	params := map[string]interface{}{
		"key":      key,
		"required": required,
		"size":     size,
	}
	// optionals
	if xdefault.Specified {
		params["default"] = xdefault.Value
	}
	if isArray.Specified {
		params["array"] = isArray.Value
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

func (srv *Database) CreateEmailAttribute(CollectionId string, key string, required bool, xdefault Optional[string], isArray Optional[bool]) (map[string]interface{}, error) {
	Type := "email"
	path := "/database/collections/{collectionId}/attributes/{type}"
	r := strings.NewReplacer("{collectionId}", CollectionId, "{type}", Type)
	path = r.Replace(path)

	params := map[string]interface{}{
		"key":      key,
		"required": required,
	}
	// optionals
	if xdefault.Specified {
		params["default"] = xdefault.Value
	}
	if isArray.Specified {
		params["arrray"] = isArray.Value
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

func (srv *Database) CreateEnumAttribute(CollectionId string, key string, elements []string, required bool, xdefault, isArray Optional[string]) (map[string]interface{}, error) {
	Type := "enum"
	path := "/database/collections/{collectionId}/attributes/{type}"
	r := strings.NewReplacer("{collectionId}", CollectionId, "{type}", Type)
	path = r.Replace(path)

	params := map[string]interface{}{
		"key":      key,
		"required": required,
		"elements": elements,
	}
	// optionals
	if xdefault.Specified {
		if contains(elements, xdefault.Value) {
			params["default"] = xdefault.Value
		} else {
			fmt.Println("❌ The default value is not contained in the array of elementes")
			return nil, nil
		}
	}
	if isArray.Specified {
		params["array"] = isArray.Value
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

func (srv *Database) CreateIpAttribute(CollectionId string, key string, required bool, xdefault, isArray Optional[bool]) (map[string]interface{}, error) {
	Type := "ip"
	path := "/database/collections/{collectionId}/attributes/{type}"
	r := strings.NewReplacer("{collectionId}", CollectionId, "{type}", Type)
	path = r.Replace(path)

	params := map[string]interface{}{
		"key":      key,
		"required": required,
	}
	if xdefault.Specified {
		params["default"] = xdefault.Value
	}
	if isArray.Specified {
		params["array"] = isArray
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// Url Layout url://something
func (srv *Database) CreateUrlAttribute(CollectionId string, key string, required bool, xdefault Optional[string], isArray Optional[bool]) (map[string]interface{}, error) {
	Type := "url"
	path := "/database/collections/{collectionId}/attributes/{type}"
	r := strings.NewReplacer("{collectionId}", CollectionId, "{type}", Type)
	path = r.Replace(path)

	params := map[string]interface{}{
		"key":      key,
		"required": required,
	}
	// optionals
	if xdefault.Specified {
		params["default"] = xdefault.Value
	}
	if isArray.Specified {
		params["array"] = isArray.Value
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

func (srv *Database) CreateIntegerAttribute(CollectionId string, key string, required bool, min, max, xdefault Optional[int], isArray Optional[bool]) (map[string]interface{}, error) {
	return createNumAttribute(srv, CollectionId, key, required, isArray, min, max, xdefault)
}

func (srv *Database) CreateFloatAttribute(CollectionId string, key string, required bool, min, max, xdefault Optional[float64], isArray Optional[bool]) (map[string]interface{}, error) {
	return createNumAttribute(srv, CollectionId, key, required, isArray, min, max, xdefault)
}

func createNumAttribute[T any](srv *Database, CollectionId string, key string, required bool, isArray Optional[bool], min, max, xdefault Optional[T]) (map[string]interface{}, error) {
	reg, _ := regexp.Compile(`\D+`)
	Type := reg.FindString(reflect.TypeOf(min.Value).String())
	if Type == "int" {
		Type = "integer"
	}
	path := "/database/collections/{collectionId}/attributes/{type}"
	r := strings.NewReplacer("{collectionId}", CollectionId, "{type}", Type)
	path = r.Replace(path)
	params := map[string]interface{}{
		"key":      key,
		"required": required,
	}
	// optionals
	if min.Specified {
		params["min"] = min.Value
	}
	if max.Specified {
		params["max"] = max.Value
	}
	if xdefault.Specified {
		params["default"] = xdefault.Value
	}
	if isArray.Specified {
		params["array"] = isArray.Value
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

func (srv *Database) CreateBooleanAttribute(CollectionId string, key string, required bool, isArray, xdefault Optional[bool]) (map[string]interface{}, error) {
	Type := "boolean"
	path := "/database/collections/{collectionId}/attributes/{type}"
	r := strings.NewReplacer("{collectionId}", CollectionId, "{type}", Type)
	path = r.Replace(path)

	params := map[string]interface{}{
		"key":      key,
		"required": required,
	}
	if xdefault.Specified {
		params["default"] = xdefault.Value
	}
	if isArray.Specified {
		params["array"] = isArray
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)

}
