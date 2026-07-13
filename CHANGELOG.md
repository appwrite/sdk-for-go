# Change Log

## v6.0.0

* Breaking: Removed `Health` and `Usage` services and their models
* Breaking: Removed `ListMessageLogs`, `ListProviderLogs`, `ListSubscriberLogs`, `ListTopicLogs` from `Messaging`
* Breaking: Changed spatial default setter types to `[][]interface{}` and `[]float64` in `Databases`/`TablesDB`
* Breaking: Removed OS, client, and device fields plus `countryCode`/`countryName` from `ActivityEvent` (`country` remains)
* Added: `Organization` `Get`, `Update`, `Delete`, and membership management methods
* Added: `WithBearer` client option for OAuth bearer-token authentication
* Added: `VectorDot`, `VectorCosine`, `VectorEuclidean` query helpers
* Added: `NewSpecification` option to `Backups.CreateRestoration`, `Specification` option to `TablesDB.Create`
* Added: `Type` filter option to `Functions.ListSpecifications` and `Sites.ListSpecifications`
* Added: `UpdateOAuth2Appwrite` and `UpdateDenyCorporateEmailPolicy` methods to `Project`
* Added: `Token` option to `Functions.GetDeploymentDownload`, new OAuth2 OIDC/server options
* Added: billing plan, `Organization`, `OAuth2Appwrite`, `Program` response models
* Added: email canonicalization and classification fields to `User` model
* Added: `Type` field to `BackupPolicy` model
* Updated: deprecated `Databases` transaction methods in favor of `TablesDB` equivalents

## v5.1.0

* Added: `createSesProvider` and `updateSesProvider` to `messaging`
* Added: `updateOAuth2Server` to `project` for OAuth2 server settings
* Added: `updatePasswordStrengthPolicy` and `PolicyPasswordStrength` to `project`
* Added: `getAuditsDB` health check to `health`
* Added: `password-strength` to `ProjectPolicyId`
* Added: `apps.read` and `apps.write` to `ProjectKeyScopes`


## v5.0.0

* Breaking: Removed `githubImagine` and `googleImagine` from `ProjectOAuthProviderId`
* Breaking: Removed `deno-1.21`, `deno-1.24`, and `deno-1.35` from `Runtime` and `BuildRuntime`
* Breaking: Dropped numeric suffixes from `StatusCode` redirect members
* Added: `Organization` service for managing projects and API keys
* Added: `PolicyDenyAliasedEmail`, `PolicyDenyDisposableEmail`, and `PolicyDenyFreeEmail` policy models
* Added: `deny-aliased-email`, `deny-disposable-email`, and `deny-free-email` to `ProjectPolicyId`
* Replaced: `BrowserTheme`, `HealthQueueName`, `OrganizationKeyScopes`, and `Region` enums
* Added: `dart-3.12` and `flutter-3.44` runtimes
* Added: `ProjectList` model and new attributes on `Function`, `Site`, and `UsageGauge`
* Updated: `functions`, `sites`, `usage`, `health`, and `avatars` services
* Updated: Renamed `updatePresence` to `update` in the `presences` service

## v4.0.0

* Breaking: Renamed `AuthMethod` enum to `ProjectAuthMethodId`
* Breaking: Renamed `EmailTemplateType` to `ProjectEmailTemplateId` and `EmailTemplateLocale` to `ProjectEmailTemplateLocale`
* Breaking: Renamed `ServiceId` to `ProjectServiceId`, `ProtocolId` to `ProjectProtocolId`, `Secure` to `ProjectSMTPSecure`, `ProjectPolicy` to `ProjectPolicyId`
* Breaking: Replaced `Scopes` enum with `ProjectKeyScopes` for project key endpoints
* Breaking: Removed `Project.UpdateDenyCanonicalEmailPolicy`; replaced with `UpdateDenyAliasedEmailPolicy`, `UpdateDenyDisposableEmailPolicy`, and `UpdateDenyFreeEmailPolicy`
* Breaking: Removed `AuthProvider` model; use new `ProjectOAuthProviderId` enum instead
* Added: `Project.Get` method to fetch current project details
* Added: `Advisor`, `Presences`, and `Usage` services
* Added: `Insight`, `Presence`, `Report`, `UsageEvent`, and `UsageGauge` models with list variants
* Added: `ProjectAuthMethod`, `ProjectProtocol`, and `ProjectService` models
* Added: `ProjectOAuthProviderId` and `ProjectOAuth2GooglePrompt` enums
* Updated: `Project`, `Database`, and `OAuth2Google` model schemas
* Updated: `X-Appwrite-Response-Format` header to `1.9.5`

## v3.1.0

* Added: Introduced `bigint` create/update APIs for legacy Databases attributes
* Added: Introduced `bigint` create/update APIs for `TablesDB` columns
* Updated: Extended key-list query filters with `key`, `resourceType`, `resourceId`, and `secret`

## v3.0.0

* [BREAKING] Renamed Webhook model fields: `security` → `tls`, `httpUser` → `authUsername`, `httpPass` → `authPassword`, `signatureKey` → `secret`
* [BREAKING] Renamed Webhook service parameters to match: `security` → `tls`, `httpUser` → `authUsername`, `httpPass` → `authPassword`
* [BREAKING] Renamed `Webhooks.UpdateSignature()` to `Webhooks.UpdateSecret()` with new optional `secret` parameter
* Added `Client.GetHeaders()` method to retrieve request headers
* Added `secret` parameter to Webhook create and update methods
* Added `x` OAuth provider to `OAuthProvider` enum
* Added `userType` field to `Log` model
* Added `purge` parameter to `UpdateCollection` and `UpdateTable` for cache invalidation
* Added Project service: platform CRUD, key CRUD, protocol/service status management
* Added new models: `Key`, `KeyList`, `Project`, `DevKey`, `MockNumber`, `AuthProvider`, `PlatformAndroid`, `PlatformApple`, `PlatformLinux`, `PlatformList`, `PlatformWeb`, `PlatformWindows`, `BillingLimits`, `Block`
* Added new enums: `PlatformType`, `ProtocolId`, `ServiceId`
* Updated `BuildRuntime`, `Runtime` enums with `dart-3.11` and `flutter-3.41`
* Updated `Scopes` enum with `KeysRead`, `KeysWrite`, `PlatformsRead`, `PlatformsWrite`
* Updated `X-Appwrite-Response-Format` header to `1.9.1`
* Updated TTL description for list caching in Databases and TablesDB

## v2.1.0

* Added `GetHeaders` method to `Client` for retrieving current request headers
* Fixed chunked upload resume handling for responses returned as strings
* Updated Go module path to `v2` for proper Go modules compatibility

## v2.0.0

* [BREAKING] Changed `$sequence` type from `int` to `string` for rows and documents
* Added `NewProject` and `NewWebhooks` client constructors
* Added impersonation helpers `WithImpersonateUserId`, `WithImpersonateUserEmail`, `WithImpersonateUserPhone`
* Added avatar URL helpers: `GetBrowserURL`, `GetCreditCardURL`, `GetFaviconURL`, `GetFlagURL`, `GetImageURL`, `GetInitialsURL`, `GetQRURL`, `GetScreenshotURL`
* Updated README badge to API version 1.9.0

## v1.0.0

* Breaking: Activate parameter was removed from CreateDeployment; use WithCreateDeploymentActivate.
* Breaking: UpdateRelationshipAttribute API path changed and old overload removed.
* Added: GetConsolePausing endpoint to monitor console pausing status.
* Added: TTL option to list operations for documents and rows.
* Updated: Document and Row sequence comments to reflect sequence IDs.

## v0.17.0

* Added new Activities service to the Go SDK with a NewActivities constructor to access Activities endpoints.
* Extended Databases attribute APIs to support encryption: introduced Encrypt option for Longtext, Mediumtext, Text, and Varchar attributes, along with corresponding New/Create option builders (WithCreateLongtextAttributeEncrypt, WithCreateMediumtextAttributeEncrypt, WithCreateTextAttributeEncrypt, WithCreateVarcharAttributeEncrypt) and wiring to send the encrypt parameter when enabled.
* Updated documentation and examples to demonstrate the new encrypt option for attribute creation (e.g., docs/examples/databases/create-longtext-attribute.md, create-mediumtext-attribute.md, create-text-attribute.md, create-varchar-attribute.md).
* Updated README to reflect compatibility with Appwrite server version 1.8.x.
* Add support for the new `Backups` service

## v0.16.1

* Fix doc examples with proper formatting

## v0.16.0

* Added ability to create columns and indexes synchronously while creating a table
* Breaking change: `Output` enum has been removed; use `ImageFormat` instead.
* Add `getQueueAudits` support to `Health` service.
* Add longtext/mediumtext/text/varchar attribute and column helpers to `Databases` and `TablesDB` services.

## v0.15.0

* Rename `VCSDeploymentType` enum to `VCSReferenceType`
* Change `CreateTemplateDeployment` method signature: replace `Version` parameter with `Type` (TemplateReferenceType) and `Reference` parameters
* Add `GetScreenshot` method to `Avatars` service
* Add `Theme`, `Timezone` and `Output` enums

## v0.14.0

* Add `total` parameter to list queries allowing skipping counting rows in a table for improved performance
* Add `Operator` class for atomic modification of rows via update, bulk update, upsert, and bulk upsert operations
* Add `createResendProvider` and `updateResendProvider` methods to `Messaging` service

## v0.13.1

* Add transaction support for Databases and TablesDB

## v0.12.0

* Deprecate `createVerification` method in `Account` service
* Add `createEmailVerification` method in `Account` service
* Add `orderRandom` query support

## 0.9.0

* Add `incrementDocumentAttribute` and `decrementDocumentAttribute` support to `Databases` service
* Add `upsertDocument` support to `Databases` service
* Update doc examples to use correct syntax

## 0.8.0

* Add `<REGION>` to doc examples due to the new multi region endpoints
* Add doc examples and methods for bulk api transactions: `createDocuments`, `deleteDocuments` etc.
* Add doc examples, class and methods for new `Sites` service
* Add doc examples, class and methods for new `Tokens` service
* Add enums for `BuildRuntime `, `Adapter`, `Framework`, `DeploymentDownloadType` and `VCSDeploymentType`
* Update enum for `runtimes` with Pythonml312, Dart219, Flutter327 and Flutter329
* Add `token` param to `getFilePreview` and `getFileView` for File tokens usage
* Add `queries` and `search` params to `listMemberships` method
* Remove `search` param from `listExecutions` method

## 0.7.0

* Version skipped

## 0.6.0

* Add bulk API methods: `createDocuments`, `deleteDocuments` etc.

## 0.5.0

* Fix requests failing by removing `Content-Type` header from `GET` and `HEAD` requests

## 0.4.0

* Fix pong response and chunked upload

## 0.3.0

* Add new push message parameters
