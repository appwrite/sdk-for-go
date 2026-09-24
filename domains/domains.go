package domains

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/models"
)

// Domains service
type Domains struct {
	client client.Client
}

func New(clt client.Client) *Domains {
	return &Domains{
		client: clt,
	}
}

type ListOptions struct {
	Queries        []string
	Search         string
	enabledSetters map[string]bool
}

func (options ListOptions) New() *ListOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Search": false}
	return &options
}

type ListOption func(*ListOptions)

func (srv *Domains) WithListQueries(v []string) ListOption {
	return func(o *ListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Domains) WithListSearch(v string) ListOption {
	return func(o *ListOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}

// List list all domains registered for this project. This endpoint supports
// pagination.
func (srv *Domains) List(optionalSetters ...ListOption) (*models.DomainsList, error) {
	path := "/domains"
	options := ListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DomainsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DomainsList
	parsed, ok := resp.Result.(models.DomainsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Create create a new domain. Before creating a domain, you need to ensure
// that your DNS provider is properly configured. After creating the domain,
// you can use the verification endpoint to check if the domain is ready to be
// used.
func (srv *Domains) Create(TeamId string, Domain string) (*models.Domain, error) {
	path := "/domains"
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	params["domain"] = Domain
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.Domain{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Domain
	parsed, ok := resp.Result.(models.Domain)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type GetPriceOptions struct {
	PeriodYears      int
	RegistrationType string
	enabledSetters   map[string]bool
}

func (options GetPriceOptions) New() *GetPriceOptions {
	options.enabledSetters = map[string]bool{"PeriodYears": false, "RegistrationType": false}
	return &options
}

type GetPriceOption func(*GetPriceOptions)

func (srv *Domains) WithGetPricePeriodYears(v int) GetPriceOption {
	return func(o *GetPriceOptions) {
		o.PeriodYears = v
		o.enabledSetters["PeriodYears"] = true
	}
}
func (srv *Domains) WithGetPriceRegistrationType(v string) GetPriceOption {
	return func(o *GetPriceOptions) {
		o.RegistrationType = v
		o.enabledSetters["RegistrationType"] = true
	}
}

// GetPrice get the registration price for a domain name.
//
// Deprecated: This API has been deprecated since 2.0.0. Please use `Domains.listPrices` instead.
func (srv *Domains) GetPrice(Domain string, optionalSetters ...GetPriceOption) (*models.DomainPrice, error) {
	path := "/domains/price"
	options := GetPriceOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["domain"] = Domain
	if options.enabledSetters["PeriodYears"] {
		params["periodYears"] = options.PeriodYears
	}
	if options.enabledSetters["RegistrationType"] {
		params["registrationType"] = options.RegistrationType
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DomainPrice{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DomainPrice
	parsed, ok := resp.Result.(models.DomainPrice)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListPricesOptions struct {
	PeriodYears      int
	RegistrationType string
	enabledSetters   map[string]bool
}

func (options ListPricesOptions) New() *ListPricesOptions {
	options.enabledSetters = map[string]bool{"PeriodYears": false, "RegistrationType": false}
	return &options
}

type ListPricesOption func(*ListPricesOptions)

func (srv *Domains) WithListPricesPeriodYears(v int) ListPricesOption {
	return func(o *ListPricesOptions) {
		o.PeriodYears = v
		o.enabledSetters["PeriodYears"] = true
	}
}
func (srv *Domains) WithListPricesRegistrationType(v string) ListPricesOption {
	return func(o *ListPricesOptions) {
		o.RegistrationType = v
		o.enabledSetters["RegistrationType"] = true
	}
}

// ListPrices check availability and get the requested registration price for
// one or more domain names. Availability is resolved for all domains in a
// single registrar lookup. Unavailable domains have a null price for new
// registrations, but can still be priced for renewal, transfer, or trade.
// Every priced domain also carries its renewal price for the same period, so
// a separate renewal lookup is not needed. A domain whose price could not be
// resolved, for example because its TLD is not supported, is returned with a
// null price.
func (srv *Domains) ListPrices(Domains []string, optionalSetters ...ListPricesOption) (*models.DomainPricesList, error) {
	path := "/domains/prices"
	options := ListPricesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["domains"] = Domains
	if options.enabledSetters["PeriodYears"] {
		params["periodYears"] = options.PeriodYears
	}
	if options.enabledSetters["RegistrationType"] {
		params["registrationType"] = options.RegistrationType
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DomainPricesList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DomainPricesList
	parsed, ok := resp.Result.(models.DomainPricesList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Get get a domain by its unique ID.
func (srv *Domains) Get(DomainId string) (*models.Domain, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.Domain{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Domain
	parsed, ok := resp.Result.(models.Domain)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// Delete delete a domain by its unique ID. This endpoint can be used to
// delete a domain from your project.
// Once deleted, the domain will no longer be available for use and all
// associated resources will be removed.
func (srv *Domains) Delete(DomainId string) (*interface{}, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		var parsed interface{}

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed interface{}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateNameserversOptions struct {
	Nameservers    []string
	enabledSetters map[string]bool
}

func (options UpdateNameserversOptions) New() *UpdateNameserversOptions {
	options.enabledSetters = map[string]bool{"Nameservers": false}
	return &options
}

type UpdateNameserversOption func(*UpdateNameserversOptions)

func (srv *Domains) WithUpdateNameserversNameservers(v []string) UpdateNameserversOption {
	return func(o *UpdateNameserversOptions) {
		o.Nameservers = v
		o.enabledSetters["Nameservers"] = true
	}
}

// UpdateNameservers update the registrar nameservers for the given domain.
// When nameservers are not provided,
// the domain will be updated to use Appwrite nameservers.
func (srv *Domains) UpdateNameservers(DomainId string, optionalSetters ...UpdateNameserversOption) (*models.Domain, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/nameservers")
	options := UpdateNameserversOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Nameservers"] {
		params["nameservers"] = options.Nameservers
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.Domain{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Domain
	parsed, ok := resp.Result.(models.Domain)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// VerifyNameservers verify which NS records are used and update the domain
// accordingly. This will check the domain's
// nameservers and update the domain's status based on whether the nameservers
// match the expected
// Appwrite nameservers.
func (srv *Domains) VerifyNameservers(DomainId string) (*models.Domain, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/nameservers/verification")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.Domain{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Domain
	parsed, ok := resp.Result.(models.Domain)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetPresetGoogleWorkspace list Google Workspace DNS records.
func (srv *Domains) GetPresetGoogleWorkspace(DomainId string) (*models.DnsRecordsList, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/presets/google-workspace")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecordsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecordsList
	parsed, ok := resp.Result.(models.DnsRecordsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreatePresetGoogleWorkspace add Google Workspace DNS records to the domain.
// This will create the required MX records
// for Google Workspace email hosting.
func (srv *Domains) CreatePresetGoogleWorkspace(DomainId string) (*models.DnsRecordsList, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/presets/google-workspace")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecordsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecordsList
	parsed, ok := resp.Result.(models.DnsRecordsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetPresetICloud list iCloud DNS records.
func (srv *Domains) GetPresetICloud(DomainId string) (*models.DnsRecordsList, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/presets/icloud")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecordsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecordsList
	parsed, ok := resp.Result.(models.DnsRecordsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreatePresetICloud add iCloud DNS records to the domain. This will create
// the required MX and SPF records
// for using iCloud email services with your domain.
func (srv *Domains) CreatePresetICloud(DomainId string) (*models.DnsRecordsList, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/presets/icloud")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecordsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecordsList
	parsed, ok := resp.Result.(models.DnsRecordsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetPresetMailgun list Mailgun DNS records.
func (srv *Domains) GetPresetMailgun(DomainId string) (*models.DnsRecordsList, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/presets/mailgun")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecordsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecordsList
	parsed, ok := resp.Result.(models.DnsRecordsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreatePresetMailgun add Mailgun DNS records to the domain. This endpoint
// will create the required DNS records
// for Mailgun in the specified domain.
func (srv *Domains) CreatePresetMailgun(DomainId string) (*models.DnsRecordsList, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/presets/mailgun")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecordsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecordsList
	parsed, ok := resp.Result.(models.DnsRecordsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetPresetOutlook list Outlook DNS records.
func (srv *Domains) GetPresetOutlook(DomainId string) (*models.DnsRecordsList, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/presets/outlook")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecordsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecordsList
	parsed, ok := resp.Result.(models.DnsRecordsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreatePresetOutlook add Outlook DNS records to the domain. This will create
// the required MX records
// for setting up Outlook email hosting for your domain.
func (srv *Domains) CreatePresetOutlook(DomainId string) (*models.DnsRecordsList, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/presets/outlook")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecordsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecordsList
	parsed, ok := resp.Result.(models.DnsRecordsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetPresetProtonMail list ProtonMail DNS records.
func (srv *Domains) GetPresetProtonMail(DomainId string) (*models.DnsRecordsList, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/presets/proton-mail")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecordsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecordsList
	parsed, ok := resp.Result.(models.DnsRecordsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreatePresetProtonMail add ProtonMail DNS records to the domain. This will
// create the required MX records
// for using ProtonMail with your custom domain.
func (srv *Domains) CreatePresetProtonMail(DomainId string) (*models.DnsRecordsList, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/presets/proton-mail")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecordsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecordsList
	parsed, ok := resp.Result.(models.DnsRecordsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetPresetZoho list Zoho DNS records.
func (srv *Domains) GetPresetZoho(DomainId string) (*models.DnsRecordsList, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/presets/zoho")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecordsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecordsList
	parsed, ok := resp.Result.(models.DnsRecordsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// CreatePresetZoho add Zoho Mail DNS records to the domain. This will create
// the required MX records
// for setting up Zoho Mail on your domain.
func (srv *Domains) CreatePresetZoho(DomainId string) (*models.DnsRecordsList, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/presets/zoho")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecordsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecordsList
	parsed, ok := resp.Result.(models.DnsRecordsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListRecordsOptions struct {
	Queries        []string
	enabledSetters map[string]bool
}

func (options ListRecordsOptions) New() *ListRecordsOptions {
	options.enabledSetters = map[string]bool{"Queries": false}
	return &options
}

type ListRecordsOption func(*ListRecordsOptions)

func (srv *Domains) WithListRecordsQueries(v []string) ListRecordsOption {
	return func(o *ListRecordsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}

// ListRecords list DNS records for a given domain. You can use this endpoint
// to list all the DNS records
// associated with your domain.
func (srv *Domains) ListRecords(DomainId string, optionalSetters ...ListRecordsOption) (*models.DnsRecordsList, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/records")
	options := ListRecordsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecordsList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecordsList
	parsed, ok := resp.Result.(models.DnsRecordsList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateRecordAOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options CreateRecordAOptions) New() *CreateRecordAOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type CreateRecordAOption func(*CreateRecordAOptions)

func (srv *Domains) WithCreateRecordAComment(v string) CreateRecordAOption {
	return func(o *CreateRecordAOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// CreateRecordA create a new A record for the given domain. A records are
// used to point a domain name
// to an IPv4 address. The record value should be a valid IPv4 address.
func (srv *Domains) CreateRecordA(DomainId string, Name string, Value string, Ttl int, optionalSetters ...CreateRecordAOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/records/a")
	options := CreateRecordAOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateRecordAOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options UpdateRecordAOptions) New() *UpdateRecordAOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type UpdateRecordAOption func(*UpdateRecordAOptions)

func (srv *Domains) WithUpdateRecordAComment(v string) UpdateRecordAOption {
	return func(o *UpdateRecordAOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// UpdateRecordA update an existing A record for the given domain. This
// endpoint allows you to modify
// the properties of an A record including its name (subdomain), IPv4 address,
// TTL,
// and optional comment.
func (srv *Domains) UpdateRecordA(DomainId string, RecordId string, Name string, Value string, Ttl int, optionalSetters ...UpdateRecordAOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}
	if RecordId == "" {
		return nil, errors.New("Missing required parameter: \"recordId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId), "{recordId}", client.EncodePath(RecordId))
	path := r.Replace("/domains/{domainId}/records/a/{recordId}")
	options := UpdateRecordAOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateRecordAAAAOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options CreateRecordAAAAOptions) New() *CreateRecordAAAAOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type CreateRecordAAAAOption func(*CreateRecordAAAAOptions)

func (srv *Domains) WithCreateRecordAAAAComment(v string) CreateRecordAAAAOption {
	return func(o *CreateRecordAAAAOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// CreateRecordAAAA create a new AAAA record for the given domain. This
// endpoint allows you to add a new IPv6 DNS record
// to your domain. The record will be used to point a hostname to an IPv6
// address.
func (srv *Domains) CreateRecordAAAA(DomainId string, Name string, Value string, Ttl int, optionalSetters ...CreateRecordAAAAOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/records/aaaa")
	options := CreateRecordAAAAOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateRecordAAAAOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options UpdateRecordAAAAOptions) New() *UpdateRecordAAAAOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type UpdateRecordAAAAOption func(*UpdateRecordAAAAOptions)

func (srv *Domains) WithUpdateRecordAAAAComment(v string) UpdateRecordAAAAOption {
	return func(o *UpdateRecordAAAAOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// UpdateRecordAAAA update an existing AAAA record for the given domain. This
// endpoint allows you to modify
// the properties of an existing AAAA record, including its name (subdomain),
// IPv6 address,
// TTL, and optional comment.
func (srv *Domains) UpdateRecordAAAA(DomainId string, RecordId string, Name string, Value string, Ttl int, optionalSetters ...UpdateRecordAAAAOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}
	if RecordId == "" {
		return nil, errors.New("Missing required parameter: \"recordId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId), "{recordId}", client.EncodePath(RecordId))
	path := r.Replace("/domains/{domainId}/records/aaaa/{recordId}")
	options := UpdateRecordAAAAOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateRecordAliasOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options CreateRecordAliasOptions) New() *CreateRecordAliasOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type CreateRecordAliasOption func(*CreateRecordAliasOptions)

func (srv *Domains) WithCreateRecordAliasComment(v string) CreateRecordAliasOption {
	return func(o *CreateRecordAliasOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// CreateRecordAlias create a new ALIAS record for the given domain. This
// record type can be used to point your domain
// to another domain name that will serve as an alias. This is particularly
// useful when you want to
// map your domain to a target domain that may change its IP address.
func (srv *Domains) CreateRecordAlias(DomainId string, Name string, Value string, Ttl int, optionalSetters ...CreateRecordAliasOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/records/alias")
	options := CreateRecordAliasOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateRecordAliasOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options UpdateRecordAliasOptions) New() *UpdateRecordAliasOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type UpdateRecordAliasOption func(*UpdateRecordAliasOptions)

func (srv *Domains) WithUpdateRecordAliasComment(v string) UpdateRecordAliasOption {
	return func(o *UpdateRecordAliasOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// UpdateRecordAlias update an existing ALIAS record for the specified domain.
// This endpoint allows you to modify
// the properties of an existing ALIAS record including its name, target
// domain, TTL, and comment.
//
// The ALIAS record type is similar to a CNAME record but can be used at the
// zone apex (root domain).
// It provides a way to map one domain name to another.
func (srv *Domains) UpdateRecordAlias(DomainId string, RecordId string, Name string, Value string, Ttl int, optionalSetters ...UpdateRecordAliasOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}
	if RecordId == "" {
		return nil, errors.New("Missing required parameter: \"recordId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId), "{recordId}", client.EncodePath(RecordId))
	path := r.Replace("/domains/{domainId}/records/alias/{recordId}")
	options := UpdateRecordAliasOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateRecordCAAOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options CreateRecordCAAOptions) New() *CreateRecordCAAOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type CreateRecordCAAOption func(*CreateRecordCAAOptions)

func (srv *Domains) WithCreateRecordCAAComment(v string) CreateRecordCAAOption {
	return func(o *CreateRecordCAAOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// CreateRecordCAA create a new CAA record for the given domain. CAA records
// are used to specify which
// Certificate Authorities (CAs) are allowed to issue SSL/TLS certificates for
// your domain.
func (srv *Domains) CreateRecordCAA(DomainId string, Name string, Value string, Ttl int, optionalSetters ...CreateRecordCAAOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/records/caa")
	options := CreateRecordCAAOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateRecordCAAOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options UpdateRecordCAAOptions) New() *UpdateRecordCAAOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type UpdateRecordCAAOption func(*UpdateRecordCAAOptions)

func (srv *Domains) WithUpdateRecordCAAComment(v string) UpdateRecordCAAOption {
	return func(o *UpdateRecordCAAOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// UpdateRecordCAA update an existing CAA record for the given domain. A CAA
// (Certification Authority Authorization)
// record is used to specify which certificate authorities (CAs) are
// authorized to issue certificates
// for a domain.
func (srv *Domains) UpdateRecordCAA(DomainId string, RecordId string, Name string, Value string, Ttl int, optionalSetters ...UpdateRecordCAAOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}
	if RecordId == "" {
		return nil, errors.New("Missing required parameter: \"recordId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId), "{recordId}", client.EncodePath(RecordId))
	path := r.Replace("/domains/{domainId}/records/caa/{recordId}")
	options := UpdateRecordCAAOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateRecordCNAMEOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options CreateRecordCNAMEOptions) New() *CreateRecordCNAMEOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type CreateRecordCNAMEOption func(*CreateRecordCNAMEOptions)

func (srv *Domains) WithCreateRecordCNAMEComment(v string) CreateRecordCNAMEOption {
	return func(o *CreateRecordCNAMEOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// CreateRecordCNAME create a new CNAME record for the given domain.
//
// A CNAME record maps a subdomain to another domain name, allowing you to
// create aliases
// for your domain. For example, you can create a CNAME record to point
// 'blog.example.com'
// to 'example.wordpress.com'.
func (srv *Domains) CreateRecordCNAME(DomainId string, Name string, Value string, Ttl int, optionalSetters ...CreateRecordCNAMEOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/records/cname")
	options := CreateRecordCNAMEOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateRecordCNAMEOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options UpdateRecordCNAMEOptions) New() *UpdateRecordCNAMEOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type UpdateRecordCNAMEOption func(*UpdateRecordCNAMEOptions)

func (srv *Domains) WithUpdateRecordCNAMEComment(v string) UpdateRecordCNAMEOption {
	return func(o *UpdateRecordCNAMEOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// UpdateRecordCNAME update an existing CNAME record for the given domain.
func (srv *Domains) UpdateRecordCNAME(DomainId string, RecordId string, Name string, Value string, Ttl int, optionalSetters ...UpdateRecordCNAMEOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}
	if RecordId == "" {
		return nil, errors.New("Missing required parameter: \"recordId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId), "{recordId}", client.EncodePath(RecordId))
	path := r.Replace("/domains/{domainId}/records/cname/{recordId}")
	options := UpdateRecordCNAMEOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateRecordHTTPSOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options CreateRecordHTTPSOptions) New() *CreateRecordHTTPSOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type CreateRecordHTTPSOption func(*CreateRecordHTTPSOptions)

func (srv *Domains) WithCreateRecordHTTPSComment(v string) CreateRecordHTTPSOption {
	return func(o *CreateRecordHTTPSOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// CreateRecordHTTPS create a new HTTPS record for the given domain. This
// record is used to configure HTTPS
// settings for your domain, enabling secure communication over SSL/TLS.
func (srv *Domains) CreateRecordHTTPS(DomainId string, Name string, Value string, Ttl int, optionalSetters ...CreateRecordHTTPSOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/records/https")
	options := CreateRecordHTTPSOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateRecordHTTPSOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options UpdateRecordHTTPSOptions) New() *UpdateRecordHTTPSOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type UpdateRecordHTTPSOption func(*UpdateRecordHTTPSOptions)

func (srv *Domains) WithUpdateRecordHTTPSComment(v string) UpdateRecordHTTPSOption {
	return func(o *UpdateRecordHTTPSOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// UpdateRecordHTTPS update an existing HTTPS record for the given domain.
// This endpoint allows you to modify
// the properties of an HTTPS record associated with your domain, including
// the name (subdomain),
// target value, TTL, and optional comment.
func (srv *Domains) UpdateRecordHTTPS(DomainId string, RecordId string, Name string, Value string, Ttl int, optionalSetters ...UpdateRecordHTTPSOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}
	if RecordId == "" {
		return nil, errors.New("Missing required parameter: \"recordId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId), "{recordId}", client.EncodePath(RecordId))
	path := r.Replace("/domains/{domainId}/records/https/{recordId}")
	options := UpdateRecordHTTPSOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateRecordMXOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options CreateRecordMXOptions) New() *CreateRecordMXOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type CreateRecordMXOption func(*CreateRecordMXOptions)

func (srv *Domains) WithCreateRecordMXComment(v string) CreateRecordMXOption {
	return func(o *CreateRecordMXOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// CreateRecordMX create a new MX record for the given domain. MX records are
// used to define the mail servers responsible
// for accepting email messages for the domain. Multiple MX records can be
// created with different priorities.
// The priority parameter determines the order in which mail servers are used,
// with lower values indicating
// higher priority.
func (srv *Domains) CreateRecordMX(DomainId string, Name string, Value string, Ttl int, Priority int, optionalSetters ...CreateRecordMXOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/records/mx")
	options := CreateRecordMXOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	params["priority"] = Priority
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateRecordMXOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options UpdateRecordMXOptions) New() *UpdateRecordMXOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type UpdateRecordMXOption func(*UpdateRecordMXOptions)

func (srv *Domains) WithUpdateRecordMXComment(v string) UpdateRecordMXOption {
	return func(o *UpdateRecordMXOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// UpdateRecordMX update an existing MX record for the given domain.
func (srv *Domains) UpdateRecordMX(DomainId string, RecordId string, Name string, Value string, Ttl int, Priority int, optionalSetters ...UpdateRecordMXOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}
	if RecordId == "" {
		return nil, errors.New("Missing required parameter: \"recordId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId), "{recordId}", client.EncodePath(RecordId))
	path := r.Replace("/domains/{domainId}/records/mx/{recordId}")
	options := UpdateRecordMXOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	params["priority"] = Priority
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateRecordNSOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options CreateRecordNSOptions) New() *CreateRecordNSOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type CreateRecordNSOption func(*CreateRecordNSOptions)

func (srv *Domains) WithCreateRecordNSComment(v string) CreateRecordNSOption {
	return func(o *CreateRecordNSOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// CreateRecordNS create a new NS record for the given domain. NS records
// specify the nameservers that are used
// to resolve the domain name to IP addresses. Each domain can have multiple
// NS records.
func (srv *Domains) CreateRecordNS(DomainId string, Name string, Value string, Ttl int, optionalSetters ...CreateRecordNSOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/records/ns")
	options := CreateRecordNSOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateRecordNSOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options UpdateRecordNSOptions) New() *UpdateRecordNSOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type UpdateRecordNSOption func(*UpdateRecordNSOptions)

func (srv *Domains) WithUpdateRecordNSComment(v string) UpdateRecordNSOption {
	return func(o *UpdateRecordNSOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// UpdateRecordNS update an existing NS record for the given domain. This
// endpoint allows you to modify
// the properties of an NS (nameserver) record associated with your domain.
// You can update
// the record name (subdomain), target nameserver value, TTL, and add or
// modify comments
// for better record management.
func (srv *Domains) UpdateRecordNS(DomainId string, RecordId string, Name string, Value string, Ttl int, optionalSetters ...UpdateRecordNSOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}
	if RecordId == "" {
		return nil, errors.New("Missing required parameter: \"recordId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId), "{recordId}", client.EncodePath(RecordId))
	path := r.Replace("/domains/{domainId}/records/ns/{recordId}")
	options := UpdateRecordNSOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateRecordSRVOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options CreateRecordSRVOptions) New() *CreateRecordSRVOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type CreateRecordSRVOption func(*CreateRecordSRVOptions)

func (srv *Domains) WithCreateRecordSRVComment(v string) CreateRecordSRVOption {
	return func(o *CreateRecordSRVOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// CreateRecordSRV create a new SRV record for the given domain. SRV records
// are used to define the location
// of servers for specific services. For example, they can be used to specify
// which server
// handles a specific service like SIP or XMPP for the domain.
func (srv *Domains) CreateRecordSRV(DomainId string, Name string, Value string, Ttl int, Priority int, Weight int, Port int, optionalSetters ...CreateRecordSRVOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/records/srv")
	options := CreateRecordSRVOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	params["priority"] = Priority
	params["weight"] = Weight
	params["port"] = Port
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateRecordSRVOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options UpdateRecordSRVOptions) New() *UpdateRecordSRVOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type UpdateRecordSRVOption func(*UpdateRecordSRVOptions)

func (srv *Domains) WithUpdateRecordSRVComment(v string) UpdateRecordSRVOption {
	return func(o *UpdateRecordSRVOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// UpdateRecordSRV update an existing SRV record for the given domain.
//
// Required parameters:
// - domainId: Domain unique ID
// - recordId: DNS record unique ID
// - name: Record name (service name)
// - value: Target hostname for this SRV record
// - ttl: Time to live, in seconds
// - priority: Record priority
// - weight: Record weight
// - port: Port number for the service
//
// Optional parameters:
// - comment: A comment for this record
func (srv *Domains) UpdateRecordSRV(DomainId string, RecordId string, Name string, Value string, Ttl int, Priority int, Weight int, Port int, optionalSetters ...UpdateRecordSRVOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}
	if RecordId == "" {
		return nil, errors.New("Missing required parameter: \"recordId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId), "{recordId}", client.EncodePath(RecordId))
	path := r.Replace("/domains/{domainId}/records/srv/{recordId}")
	options := UpdateRecordSRVOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	params["priority"] = Priority
	params["weight"] = Weight
	params["port"] = Port
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateRecordTXTOptions struct {
	Value          string
	Comment        string
	enabledSetters map[string]bool
}

func (options CreateRecordTXTOptions) New() *CreateRecordTXTOptions {
	options.enabledSetters = map[string]bool{"Value": false, "Comment": false}
	return &options
}

type CreateRecordTXTOption func(*CreateRecordTXTOptions)

func (srv *Domains) WithCreateRecordTXTValue(v string) CreateRecordTXTOption {
	return func(o *CreateRecordTXTOptions) {
		o.Value = v
		o.enabledSetters["Value"] = true
	}
}
func (srv *Domains) WithCreateRecordTXTComment(v string) CreateRecordTXTOption {
	return func(o *CreateRecordTXTOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// CreateRecordTXT create a new TXT record for the given domain. TXT records
// can be used
// to provide additional information about your domain, such as domain
// verification records, SPF records, or DKIM records.
func (srv *Domains) CreateRecordTXT(DomainId string, Name string, Ttl int, optionalSetters ...CreateRecordTXTOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/records/txt")
	options := CreateRecordTXTOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	if options.enabledSetters["Value"] {
		params["value"] = options.Value
	}
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateRecordTXTOptions struct {
	Comment        string
	enabledSetters map[string]bool
}

func (options UpdateRecordTXTOptions) New() *UpdateRecordTXTOptions {
	options.enabledSetters = map[string]bool{"Comment": false}
	return &options
}

type UpdateRecordTXTOption func(*UpdateRecordTXTOptions)

func (srv *Domains) WithUpdateRecordTXTComment(v string) UpdateRecordTXTOption {
	return func(o *UpdateRecordTXTOptions) {
		o.Comment = v
		o.enabledSetters["Comment"] = true
	}
}

// UpdateRecordTXT update an existing TXT record for the given domain.
//
// Update the TXT record details for a specific domain by providing the domain
// ID,
// record ID, and the new record configuration including name, value, TTL, and
// an optional comment.
func (srv *Domains) UpdateRecordTXT(DomainId string, RecordId string, Name string, Value string, Ttl int, optionalSetters ...UpdateRecordTXTOption) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}
	if RecordId == "" {
		return nil, errors.New("Missing required parameter: \"recordId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId), "{recordId}", client.EncodePath(RecordId))
	path := r.Replace("/domains/{domainId}/records/txt/{recordId}")
	options := UpdateRecordTXTOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["value"] = Value
	params["ttl"] = Ttl
	if options.enabledSetters["Comment"] {
		params["comment"] = options.Comment
	}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetRecord get a single DNS record for a given domain by record ID.
//
// This endpoint allows you to retrieve a specific DNS record associated with
// a domain
// using its unique identifier. The record contains information about the DNS
// configuration
// such as type, value, and TTL settings.
func (srv *Domains) GetRecord(DomainId string, RecordId string) (*models.DnsRecord, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}
	if RecordId == "" {
		return nil, errors.New("Missing required parameter: \"recordId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId), "{recordId}", client.EncodePath(RecordId))
	path := r.Replace("/domains/{domainId}/records/{recordId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DnsRecord{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DnsRecord
	parsed, ok := resp.Result.(models.DnsRecord)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// DeleteRecord delete a DNS record for the given domain. This endpoint allows
// you to delete an existing DNS record
// from a specific domain.
func (srv *Domains) DeleteRecord(DomainId string, RecordId string) (*interface{}, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}
	if RecordId == "" {
		return nil, errors.New("Missing required parameter: \"recordId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId), "{recordId}", client.EncodePath(RecordId))
	path := r.Replace("/domains/{domainId}/records/{recordId}")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		var parsed interface{}

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed interface{}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateTeam update the team ID for a specific domain. The caller must
// administer the current
// team and be an owner of the destination team.
//
// Updating the team ID will transfer ownership and access control of the
// domain
// and all its DNS records to the new team.
func (srv *Domains) UpdateTeam(DomainId string, TeamId string) (*models.Domain, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/team")
	params := map[string]interface{}{}
	params["teamId"] = TeamId
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.Domain{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Domain
	parsed, ok := resp.Result.(models.Domain)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetTransferStatus retrieve the current transfer status for a domain.
// Returns the status, an optional reason, and a timestamp of the last status
// change.
func (srv *Domains) GetTransferStatus(DomainId string) (*models.DomainTransferStatus, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/transfers/status")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.DomainTransferStatus{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DomainTransferStatus
	parsed, ok := resp.Result.(models.DomainTransferStatus)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetZone retrieve the DNS zone file for the given domain. This endpoint will
// return the DNS
// zone file in a standardized format that can be used to configure DNS
// servers.
func (srv *Domains) GetZone(DomainId string) (*interface{}, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/zone")
	params := map[string]interface{}{}
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["accept"] = "text/plain"

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		var parsed interface{}

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed interface{}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// UpdateZone update the DNS zone for the given domain using the provided zone
// file content.
// All parsed records are imported and then the main domain document is
// returned.
func (srv *Domains) UpdateZone(DomainId string, Content string) (*models.Domain, error) {
	if DomainId == "" {
		return nil, errors.New("Missing required parameter: \"domainId\"")
	}

	r := strings.NewReplacer("{domainId}", client.EncodePath(DomainId))
	path := r.Replace("/domains/{domainId}/zone")
	params := map[string]interface{}{}
	params["content"] = Content
	headers := map[string]interface{}{}
	headers["X-Appwrite-Project"] = srv.client.Config["project"]
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes, err := client.ResponseBody(resp)
		if err != nil {
			return nil, err
		}

		parsed := models.Domain{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Domain
	parsed, ok := resp.Result.(models.Domain)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
