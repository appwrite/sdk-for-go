package appwrite

import (
	"time"

	"github.com/appwrite/sdk-for-go/v6/client"
	"github.com/appwrite/sdk-for-go/v6/account"
	"github.com/appwrite/sdk-for-go/v6/activities"
	"github.com/appwrite/sdk-for-go/v6/apps"
	"github.com/appwrite/sdk-for-go/v6/avatars"
	"github.com/appwrite/sdk-for-go/v6/backups"
	"github.com/appwrite/sdk-for-go/v6/databases"
	"github.com/appwrite/sdk-for-go/v6/embeddings"
	"github.com/appwrite/sdk-for-go/v6/functions"
	"github.com/appwrite/sdk-for-go/v6/graphql"
	"github.com/appwrite/sdk-for-go/v6/locale"
	"github.com/appwrite/sdk-for-go/v6/messaging"
	"github.com/appwrite/sdk-for-go/v6/oauth2"
	"github.com/appwrite/sdk-for-go/v6/organization"
	"github.com/appwrite/sdk-for-go/v6/presences"
	"github.com/appwrite/sdk-for-go/v6/project"
	"github.com/appwrite/sdk-for-go/v6/proxy"
	"github.com/appwrite/sdk-for-go/v6/advisor"
	"github.com/appwrite/sdk-for-go/v6/sites"
	"github.com/appwrite/sdk-for-go/v6/storage"
	"github.com/appwrite/sdk-for-go/v6/tablesdb"
	"github.com/appwrite/sdk-for-go/v6/teams"
	"github.com/appwrite/sdk-for-go/v6/tokens"
	"github.com/appwrite/sdk-for-go/v6/users"
	"github.com/appwrite/sdk-for-go/v6/webhooks"
)

func NewAccount(clt client.Client) *account.Account {
	return account.New(clt)
}
func NewActivities(clt client.Client) *activities.Activities {
	return activities.New(clt)
}
func NewApps(clt client.Client) *apps.Apps {
	return apps.New(clt)
}
func NewAvatars(clt client.Client) *avatars.Avatars {
	return avatars.New(clt)
}
func NewBackups(clt client.Client) *backups.Backups {
	return backups.New(clt)
}
func NewDatabases(clt client.Client) *databases.Databases {
	return databases.New(clt)
}
func NewEmbeddings(clt client.Client) *embeddings.Embeddings {
	return embeddings.New(clt)
}
func NewFunctions(clt client.Client) *functions.Functions {
	return functions.New(clt)
}
func NewGraphql(clt client.Client) *graphql.Graphql {
	return graphql.New(clt)
}
func NewLocale(clt client.Client) *locale.Locale {
	return locale.New(clt)
}
func NewMessaging(clt client.Client) *messaging.Messaging {
	return messaging.New(clt)
}
func NewOauth2(clt client.Client) *oauth2.Oauth2 {
	return oauth2.New(clt)
}
func NewOrganization(clt client.Client) *organization.Organization {
	return organization.New(clt)
}
func NewPresences(clt client.Client) *presences.Presences {
	return presences.New(clt)
}
func NewProject(clt client.Client) *project.Project {
	return project.New(clt)
}
func NewProxy(clt client.Client) *proxy.Proxy {
	return proxy.New(clt)
}
func NewAdvisor(clt client.Client) *advisor.Advisor {
	return advisor.New(clt)
}
func NewSites(clt client.Client) *sites.Sites {
	return sites.New(clt)
}
func NewStorage(clt client.Client) *storage.Storage {
	return storage.New(clt)
}
func NewTablesDB(clt client.Client) *tablesdb.TablesDB {
	return tablesdb.New(clt)
}
func NewTeams(clt client.Client) *teams.Teams {
	return teams.New(clt)
}
func NewTokens(clt client.Client) *tokens.Tokens {
	return tokens.New(clt)
}
func NewUsers(clt client.Client) *users.Users {
	return users.New(clt)
}
func NewWebhooks(clt client.Client) *webhooks.Webhooks {
	return webhooks.New(clt)
}

// NewClient initializes a new Appwrite client with a given timeout
func NewClient(optionalSetters ...client.ClientOption) client.Client {
	return client.New(optionalSetters...)
}

// Helper method to construct NewClient()
func WithEndpoint(endpoint string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Endpoint = endpoint
		return nil
	}
}

// Helper method to construct NewClient()
func WithTimeout(timeout time.Duration) client.ClientOption {
	return func(clt *client.Client) error {
		httpClient, err := client.GetDefaultClient(timeout)
		if err != nil {
			return err
		}

		clt.Timeout = timeout
		clt.Client = httpClient

		return nil
	}
}

// Helper method to construct NewClient()
func WithSelfSigned(status bool) client.ClientOption {
	return func(clt *client.Client) error {
		clt.SelfSigned = status
		return nil
	}
}

// Helper method to construct NewClient()
func WithChunkSize(size int64) client.ClientOption {
	return func(clt *client.Client) error {
		clt.ChunkSize = size
		return nil
	}
}

// Helper method to construct NewClient()
// 
// Your project ID
func WithProject(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["project"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// Your secret API key
func WithKey(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["key"] = value
		clt.Headers["X-Appwrite-Key"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// Your organization ID
func WithOrganization(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["organization"] = value
		clt.Headers["X-Appwrite-Organization"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// Your secret JSON Web Token
func WithJWT(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["jwt"] = value
		clt.Headers["X-Appwrite-JWT"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// The OAuth access token to authenticate with
func WithBearer(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["bearer"] = value
		clt.Headers["Authorization"] = "Bearer " + value
		return nil
	}
}
// Helper method to construct NewClient()
func WithLocale(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["locale"] = value
		clt.Headers["X-Appwrite-Locale"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// The user session to authenticate with
func WithSession(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["session"] = value
		clt.Headers["X-Appwrite-Session"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// The user agent string of the client that made the request
func WithForwardedUserAgent(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["forwardeduseragent"] = value
		clt.Headers["X-Forwarded-User-Agent"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// Your secret dev API key
func WithDevKey(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["devkey"] = value
		clt.Headers["X-Appwrite-Dev-Key"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// The user cookie to authenticate with. Used by SDKs that forward an incoming Cookie header in server-side runtimes.
func WithCookie(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["cookie"] = value
		clt.Headers["Cookie"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// Impersonate a user by ID
func WithImpersonateUserId(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["impersonateuserid"] = value
		clt.Headers["X-Appwrite-Impersonate-User-Id"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// Impersonate a user by email
func WithImpersonateUserEmail(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["impersonateuseremail"] = value
		clt.Headers["X-Appwrite-Impersonate-User-Email"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// Impersonate a user by phone
func WithImpersonateUserPhone(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Config["impersonateuserphone"] = value
		clt.Headers["X-Appwrite-Impersonate-User-Phone"] = value
		return nil
	}
}
