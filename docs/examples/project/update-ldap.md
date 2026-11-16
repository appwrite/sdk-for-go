```go
package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/v7/appwrite"
	"github.com/appwrite/sdk-for-go/v7/project"
)

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://<REGION>.cloud.appwrite.io/v1"),
		appwrite.WithProject("<YOUR_PROJECT_ID>"),
		appwrite.WithKey("<YOUR_API_KEY>"),
	)

	service := project.New(client)

	response, err := service.UpdateLDAP(
		service.WithUpdateLDAPEnabled(false),
		service.WithUpdateLDAPHosts([]string{"ldap.example.com"}),
		service.WithUpdateLDAPPort(636),
		service.WithUpdateLDAPEncryption("none"),
		service.WithUpdateLDAPVerification(false),
		service.WithUpdateLDAPCertificate("<CERTIFICATE>"),
		service.WithUpdateLDAPTimeout(10),
		service.WithUpdateLDAPBaseDn("ou=people,dc=example,dc=com"),
		service.WithUpdateLDAPBindDn("cn=appwrite,ou=services,dc=example,dc=com"),
		service.WithUpdateLDAPBindPassword("password"),
		service.WithUpdateLDAPUserFilter("(uid={username})"),
		service.WithUpdateLDAPRequiredGroupDn("cn=staff,ou=groups,dc=example,dc=com"),
		service.WithUpdateLDAPIdAttribute("entryUUID"),
		service.WithUpdateLDAPEmailAttribute("mail"),
		service.WithUpdateLDAPNameAttribute("cn"),
		service.WithUpdateLDAPPhoneAttribute("mobile"),
	)
	fmt.Println(response, err)
}
```
