package appwrite

import (
	"time"

	"github.com/appwrite/sdk-for-go/v3/client"
	"github.com/appwrite/sdk-for-go/v3/account"
	"github.com/appwrite/sdk-for-go/v3/activities"
	"github.com/appwrite/sdk-for-go/v3/avatars"
	"github.com/appwrite/sdk-for-go/v3/backups"
	"github.com/appwrite/sdk-for-go/v3/databases"
	"github.com/appwrite/sdk-for-go/v3/functions"
	"github.com/appwrite/sdk-for-go/v3/graphql"
	"github.com/appwrite/sdk-for-go/v3/health"
	"github.com/appwrite/sdk-for-go/v3/locale"
	"github.com/appwrite/sdk-for-go/v3/messaging"
	"github.com/appwrite/sdk-for-go/v3/project"
	"github.com/appwrite/sdk-for-go/v3/sites"
	"github.com/appwrite/sdk-for-go/v3/storage"
	"github.com/appwrite/sdk-for-go/v3/tablesdb"
	"github.com/appwrite/sdk-for-go/v3/teams"
	"github.com/appwrite/sdk-for-go/v3/tokens"
	"github.com/appwrite/sdk-for-go/v3/users"
	"github.com/appwrite/sdk-for-go/v3/webhooks"
)

func NewAccount(clt client.Client) *account.Account {
	return account.New(clt)
}
func NewActivities(clt client.Client) *activities.Activities {
	return activities.New(clt)
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
func NewFunctions(clt client.Client) *functions.Functions {
	return functions.New(clt)
}
func NewGraphql(clt client.Client) *graphql.Graphql {
	return graphql.New(clt)
}
func NewHealth(clt client.Client) *health.Health {
	return health.New(clt)
}
func NewLocale(clt client.Client) *locale.Locale {
	return locale.New(clt)
}
func NewMessaging(clt client.Client) *messaging.Messaging {
	return messaging.New(clt)
}
func NewProject(clt client.Client) *project.Project {
	return project.New(clt)
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
		clt.Headers["X-Appwrite-Project"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// Your secret API key
func WithKey(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Appwrite-Key"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// Your secret JSON Web Token
func WithJWT(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Appwrite-JWT"] = value
		return nil
	}
}
// Helper method to construct NewClient()
func WithLocale(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Appwrite-Locale"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// The user session to authenticate with
func WithSession(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Appwrite-Session"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// The user agent string of the client that made the request
func WithForwardedUserAgent(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Forwarded-User-Agent"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// Impersonate a user by ID on an already user-authenticated request. Requires the current request to be authenticated as a user with impersonator capability; X-Appwrite-Key alone is not sufficient. Impersonator users are intentionally granted users.read so they can discover a target before impersonation begins. Internal audit logs still attribute actions to the original impersonator and record the impersonated target only in internal audit payload data.
func WithImpersonateUserId(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Appwrite-Impersonate-User-Id"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// Impersonate a user by email on an already user-authenticated request. Requires the current request to be authenticated as a user with impersonator capability; X-Appwrite-Key alone is not sufficient. Impersonator users are intentionally granted users.read so they can discover a target before impersonation begins. Internal audit logs still attribute actions to the original impersonator and record the impersonated target only in internal audit payload data.
func WithImpersonateUserEmail(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Appwrite-Impersonate-User-Email"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// Impersonate a user by phone on an already user-authenticated request. Requires the current request to be authenticated as a user with impersonator capability; X-Appwrite-Key alone is not sufficient. Impersonator users are intentionally granted users.read so they can discover a target before impersonation begins. Internal audit logs still attribute actions to the original impersonator and record the impersonated target only in internal audit payload data.
func WithImpersonateUserPhone(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Appwrite-Impersonate-User-Phone"] = value
		return nil
	}
}
