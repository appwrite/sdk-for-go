package appwrite

// NewClient initializes a new Appwrite client
func NewClient() Client {
	clt := Client{}
	clt.headers = make(map[string]string)
	return clt
}
