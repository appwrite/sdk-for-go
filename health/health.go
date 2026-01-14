package health

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// Health service
type Health struct {
	client client.Client
}

func New(clt client.Client) *Health {
	return &Health{
		client: clt,
	}
}


// Get check the Appwrite HTTP server is up and responsive.
func (srv *Health) Get()(*models.HealthStatus, error) {
	path := "/health"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthStatus{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthStatus
	parsed, ok := resp.Result.(models.HealthStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetAntivirus check the Appwrite Antivirus server is up and connection is
// successful.
func (srv *Health) GetAntivirus()(*models.HealthAntivirus, error) {
	path := "/health/anti-virus"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthAntivirus{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthAntivirus
	parsed, ok := resp.Result.(models.HealthAntivirus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetCache check the Appwrite in-memory cache servers are up and connection
// is successful.
func (srv *Health) GetCache()(*models.HealthStatus, error) {
	path := "/health/cache"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthStatus{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthStatus
	parsed, ok := resp.Result.(models.HealthStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetCertificateOptions struct {
	Domain string
	enabledSetters map[string]bool
}
func (options GetCertificateOptions) New() *GetCertificateOptions {
	options.enabledSetters = map[string]bool{
		"Domain": false,
	}
	return &options
}
type GetCertificateOption func(*GetCertificateOptions)
func (srv *Health) WithGetCertificateDomain(v string) GetCertificateOption {
	return func(o *GetCertificateOptions) {
		o.Domain = v
		o.enabledSetters["Domain"] = true
	}
}
	
// GetCertificate get the SSL certificate for a domain
func (srv *Health) GetCertificate(optionalSetters ...GetCertificateOption)(*models.HealthCertificate, error) {
	path := "/health/certificate"
	options := GetCertificateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Domain"] {
		params["domain"] = options.Domain
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthCertificate{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthCertificate
	parsed, ok := resp.Result.(models.HealthCertificate)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetDB check the Appwrite database servers are up and connection is
// successful.
func (srv *Health) GetDB()(*models.HealthStatus, error) {
	path := "/health/db"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthStatus{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthStatus
	parsed, ok := resp.Result.(models.HealthStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetPubSub check the Appwrite pub-sub servers are up and connection is
// successful.
func (srv *Health) GetPubSub()(*models.HealthStatus, error) {
	path := "/health/pubsub"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthStatus{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthStatus
	parsed, ok := resp.Result.(models.HealthStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueBillingProjectAggregationOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueBillingProjectAggregationOptions) New() *GetQueueBillingProjectAggregationOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueBillingProjectAggregationOption func(*GetQueueBillingProjectAggregationOptions)
func (srv *Health) WithGetQueueBillingProjectAggregationThreshold(v int) GetQueueBillingProjectAggregationOption {
	return func(o *GetQueueBillingProjectAggregationOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueBillingProjectAggregation get billing project aggregation queue.
func (srv *Health) GetQueueBillingProjectAggregation(optionalSetters ...GetQueueBillingProjectAggregationOption)(*models.HealthQueue, error) {
	path := "/health/queue/billing-project-aggregation"
	options := GetQueueBillingProjectAggregationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueBillingTeamAggregationOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueBillingTeamAggregationOptions) New() *GetQueueBillingTeamAggregationOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueBillingTeamAggregationOption func(*GetQueueBillingTeamAggregationOptions)
func (srv *Health) WithGetQueueBillingTeamAggregationThreshold(v int) GetQueueBillingTeamAggregationOption {
	return func(o *GetQueueBillingTeamAggregationOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueBillingTeamAggregation get billing team aggregation queue.
func (srv *Health) GetQueueBillingTeamAggregation(optionalSetters ...GetQueueBillingTeamAggregationOption)(*models.HealthQueue, error) {
	path := "/health/queue/billing-team-aggregation"
	options := GetQueueBillingTeamAggregationOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueBuildsOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueBuildsOptions) New() *GetQueueBuildsOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueBuildsOption func(*GetQueueBuildsOptions)
func (srv *Health) WithGetQueueBuildsThreshold(v int) GetQueueBuildsOption {
	return func(o *GetQueueBuildsOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueBuilds get the number of builds that are waiting to be processed in
// the Appwrite internal queue server.
func (srv *Health) GetQueueBuilds(optionalSetters ...GetQueueBuildsOption)(*models.HealthQueue, error) {
	path := "/health/queue/builds"
	options := GetQueueBuildsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueuePriorityBuildsOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueuePriorityBuildsOptions) New() *GetQueuePriorityBuildsOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueuePriorityBuildsOption func(*GetQueuePriorityBuildsOptions)
func (srv *Health) WithGetQueuePriorityBuildsThreshold(v int) GetQueuePriorityBuildsOption {
	return func(o *GetQueuePriorityBuildsOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueuePriorityBuilds get the priority builds queue size.
func (srv *Health) GetQueuePriorityBuilds(optionalSetters ...GetQueuePriorityBuildsOption)(*models.HealthQueue, error) {
	path := "/health/queue/builds-priority"
	options := GetQueuePriorityBuildsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueCertificatesOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueCertificatesOptions) New() *GetQueueCertificatesOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueCertificatesOption func(*GetQueueCertificatesOptions)
func (srv *Health) WithGetQueueCertificatesThreshold(v int) GetQueueCertificatesOption {
	return func(o *GetQueueCertificatesOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueCertificates get the number of certificates that are waiting to be
// issued against [Letsencrypt](https://letsencrypt.org/) in the Appwrite
// internal queue server.
func (srv *Health) GetQueueCertificates(optionalSetters ...GetQueueCertificatesOption)(*models.HealthQueue, error) {
	path := "/health/queue/certificates"
	options := GetQueueCertificatesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueDatabasesOptions struct {
	Name string
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueDatabasesOptions) New() *GetQueueDatabasesOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Threshold": false,
	}
	return &options
}
type GetQueueDatabasesOption func(*GetQueueDatabasesOptions)
func (srv *Health) WithGetQueueDatabasesName(v string) GetQueueDatabasesOption {
	return func(o *GetQueueDatabasesOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Health) WithGetQueueDatabasesThreshold(v int) GetQueueDatabasesOption {
	return func(o *GetQueueDatabasesOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueDatabases get the number of database changes that are waiting to be
// processed in the Appwrite internal queue server.
func (srv *Health) GetQueueDatabases(optionalSetters ...GetQueueDatabasesOption)(*models.HealthQueue, error) {
	path := "/health/queue/databases"
	options := GetQueueDatabasesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueDeletesOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueDeletesOptions) New() *GetQueueDeletesOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueDeletesOption func(*GetQueueDeletesOptions)
func (srv *Health) WithGetQueueDeletesThreshold(v int) GetQueueDeletesOption {
	return func(o *GetQueueDeletesOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueDeletes get the number of background destructive changes that are
// waiting to be processed in the Appwrite internal queue server.
func (srv *Health) GetQueueDeletes(optionalSetters ...GetQueueDeletesOption)(*models.HealthQueue, error) {
	path := "/health/queue/deletes"
	options := GetQueueDeletesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetFailedJobsOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetFailedJobsOptions) New() *GetFailedJobsOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetFailedJobsOption func(*GetFailedJobsOptions)
func (srv *Health) WithGetFailedJobsThreshold(v int) GetFailedJobsOption {
	return func(o *GetFailedJobsOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
			
// GetFailedJobs returns the amount of failed jobs in a given queue.
func (srv *Health) GetFailedJobs(Name string, optionalSetters ...GetFailedJobsOption)(*models.HealthQueue, error) {
	r := strings.NewReplacer("{name}", Name)
	path := r.Replace("/health/queue/failed/{name}")
	options := GetFailedJobsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueFunctionsOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueFunctionsOptions) New() *GetQueueFunctionsOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueFunctionsOption func(*GetQueueFunctionsOptions)
func (srv *Health) WithGetQueueFunctionsThreshold(v int) GetQueueFunctionsOption {
	return func(o *GetQueueFunctionsOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueFunctions get the number of function executions that are waiting to
// be processed in the Appwrite internal queue server.
func (srv *Health) GetQueueFunctions(optionalSetters ...GetQueueFunctionsOption)(*models.HealthQueue, error) {
	path := "/health/queue/functions"
	options := GetQueueFunctionsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueLogsOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueLogsOptions) New() *GetQueueLogsOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueLogsOption func(*GetQueueLogsOptions)
func (srv *Health) WithGetQueueLogsThreshold(v int) GetQueueLogsOption {
	return func(o *GetQueueLogsOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueLogs get the number of logs that are waiting to be processed in the
// Appwrite internal queue server.
func (srv *Health) GetQueueLogs(optionalSetters ...GetQueueLogsOption)(*models.HealthQueue, error) {
	path := "/health/queue/logs"
	options := GetQueueLogsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueMailsOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueMailsOptions) New() *GetQueueMailsOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueMailsOption func(*GetQueueMailsOptions)
func (srv *Health) WithGetQueueMailsThreshold(v int) GetQueueMailsOption {
	return func(o *GetQueueMailsOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueMails get the number of mails that are waiting to be processed in
// the Appwrite internal queue server.
func (srv *Health) GetQueueMails(optionalSetters ...GetQueueMailsOption)(*models.HealthQueue, error) {
	path := "/health/queue/mails"
	options := GetQueueMailsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueMessagingOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueMessagingOptions) New() *GetQueueMessagingOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueMessagingOption func(*GetQueueMessagingOptions)
func (srv *Health) WithGetQueueMessagingThreshold(v int) GetQueueMessagingOption {
	return func(o *GetQueueMessagingOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueMessaging get the number of messages that are waiting to be
// processed in the Appwrite internal queue server.
func (srv *Health) GetQueueMessaging(optionalSetters ...GetQueueMessagingOption)(*models.HealthQueue, error) {
	path := "/health/queue/messaging"
	options := GetQueueMessagingOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueMigrationsOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueMigrationsOptions) New() *GetQueueMigrationsOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueMigrationsOption func(*GetQueueMigrationsOptions)
func (srv *Health) WithGetQueueMigrationsThreshold(v int) GetQueueMigrationsOption {
	return func(o *GetQueueMigrationsOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueMigrations get the number of migrations that are waiting to be
// processed in the Appwrite internal queue server.
func (srv *Health) GetQueueMigrations(optionalSetters ...GetQueueMigrationsOption)(*models.HealthQueue, error) {
	path := "/health/queue/migrations"
	options := GetQueueMigrationsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueRegionManagerOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueRegionManagerOptions) New() *GetQueueRegionManagerOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueRegionManagerOption func(*GetQueueRegionManagerOptions)
func (srv *Health) WithGetQueueRegionManagerThreshold(v int) GetQueueRegionManagerOption {
	return func(o *GetQueueRegionManagerOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueRegionManager get region manager queue.
func (srv *Health) GetQueueRegionManager(optionalSetters ...GetQueueRegionManagerOption)(*models.HealthQueue, error) {
	path := "/health/queue/region-manager"
	options := GetQueueRegionManagerOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueStatsResourcesOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueStatsResourcesOptions) New() *GetQueueStatsResourcesOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueStatsResourcesOption func(*GetQueueStatsResourcesOptions)
func (srv *Health) WithGetQueueStatsResourcesThreshold(v int) GetQueueStatsResourcesOption {
	return func(o *GetQueueStatsResourcesOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueStatsResources get the number of metrics that are waiting to be
// processed in the Appwrite stats resources queue.
func (srv *Health) GetQueueStatsResources(optionalSetters ...GetQueueStatsResourcesOption)(*models.HealthQueue, error) {
	path := "/health/queue/stats-resources"
	options := GetQueueStatsResourcesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueUsageOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueUsageOptions) New() *GetQueueUsageOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueUsageOption func(*GetQueueUsageOptions)
func (srv *Health) WithGetQueueUsageThreshold(v int) GetQueueUsageOption {
	return func(o *GetQueueUsageOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueUsage get the number of metrics that are waiting to be processed in
// the Appwrite internal queue server.
func (srv *Health) GetQueueUsage(optionalSetters ...GetQueueUsageOption)(*models.HealthQueue, error) {
	path := "/health/queue/stats-usage"
	options := GetQueueUsageOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueThreatsOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueThreatsOptions) New() *GetQueueThreatsOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueThreatsOption func(*GetQueueThreatsOptions)
func (srv *Health) WithGetQueueThreatsThreshold(v int) GetQueueThreatsOption {
	return func(o *GetQueueThreatsOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueThreats get threats queue.
func (srv *Health) GetQueueThreats(optionalSetters ...GetQueueThreatsOption)(*models.HealthQueue, error) {
	path := "/health/queue/threats"
	options := GetQueueThreatsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GetQueueWebhooksOptions struct {
	Threshold int
	enabledSetters map[string]bool
}
func (options GetQueueWebhooksOptions) New() *GetQueueWebhooksOptions {
	options.enabledSetters = map[string]bool{
		"Threshold": false,
	}
	return &options
}
type GetQueueWebhooksOption func(*GetQueueWebhooksOptions)
func (srv *Health) WithGetQueueWebhooksThreshold(v int) GetQueueWebhooksOption {
	return func(o *GetQueueWebhooksOptions) {
		o.Threshold = v
		o.enabledSetters["Threshold"] = true
	}
}
	
// GetQueueWebhooks get the number of webhooks that are waiting to be
// processed in the Appwrite internal queue server.
func (srv *Health) GetQueueWebhooks(optionalSetters ...GetQueueWebhooksOption)(*models.HealthQueue, error) {
	path := "/health/queue/webhooks"
	options := GetQueueWebhooksOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Threshold"] {
		params["threshold"] = options.Threshold
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthQueue{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthQueue
	parsed, ok := resp.Result.(models.HealthQueue)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetStorage check the Appwrite storage device is up and connection is
// successful.
func (srv *Health) GetStorage()(*models.HealthStatus, error) {
	path := "/health/storage"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthStatus{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthStatus
	parsed, ok := resp.Result.(models.HealthStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetStorageLocal check the Appwrite local storage device is up and
// connection is successful.
func (srv *Health) GetStorageLocal()(*models.HealthStatus, error) {
	path := "/health/storage/local"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthStatus{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthStatus
	parsed, ok := resp.Result.(models.HealthStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetTime check the Appwrite server time is synced with Google remote NTP
// server. We use this technology to smoothly handle leap seconds with no
// disruptive events. The [Network Time
// Protocol](https://en.wikipedia.org/wiki/Network_Time_Protocol) (NTP) is
// used by hundreds of millions of computers and devices to synchronize their
// clocks over the Internet. If your computer sets its own clock, it likely
// uses NTP.
func (srv *Health) GetTime()(*models.HealthTime, error) {
	path := "/health/time"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.HealthTime{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.HealthTime
	parsed, ok := resp.Result.(models.HealthTime)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
