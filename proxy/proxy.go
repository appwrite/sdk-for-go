package proxy

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/v4/client"
	"github.com/appwrite/sdk-for-go/v4/models"
	"strings"
)

// Proxy service
type Proxy struct {
	client client.Client
}

func New(clt client.Client) *Proxy {
	return &Proxy{
		client: clt,
	}
}

type ListRulesOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListRulesOptions) New() *ListRulesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListRulesOption func(*ListRulesOptions)
func (srv *Proxy) WithListRulesQueries(v []string) ListRulesOption {
	return func(o *ListRulesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Proxy) WithListRulesTotal(v bool) ListRulesOption {
	return func(o *ListRulesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListRules get a list of all the proxy rules. You can use the query params
// to filter your results.
func (srv *Proxy) ListRules(optionalSetters ...ListRulesOption)(*models.ProxyRuleList, error) {
	path := "/proxy/rules"
	options := ListRulesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ProxyRuleList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProxyRuleList
	parsed, ok := resp.Result.(models.ProxyRuleList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CreateAPIRule create a new proxy rule for serving Appwrite's API on custom
// domain.
// 
// Rule ID is automatically generated as MD5 hash of a rule domain for
// performance purposes.
func (srv *Proxy) CreateAPIRule(Domain string)(*models.ProxyRule, error) {
	path := "/proxy/rules/api"
	params := map[string]interface{}{}
	params["domain"] = Domain
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ProxyRule{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProxyRule
	parsed, ok := resp.Result.(models.ProxyRule)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateFunctionRuleOptions struct {
	Branch string
	enabledSetters map[string]bool
}
func (options CreateFunctionRuleOptions) New() *CreateFunctionRuleOptions {
	options.enabledSetters = map[string]bool{
		"Branch": false,
	}
	return &options
}
type CreateFunctionRuleOption func(*CreateFunctionRuleOptions)
func (srv *Proxy) WithCreateFunctionRuleBranch(v string) CreateFunctionRuleOption {
	return func(o *CreateFunctionRuleOptions) {
		o.Branch = v
		o.enabledSetters["Branch"] = true
	}
}
					
// CreateFunctionRule create a new proxy rule for executing Appwrite Function
// on custom domain.
// 
// Rule ID is automatically generated as MD5 hash of a rule domain for
// performance purposes.
func (srv *Proxy) CreateFunctionRule(Domain string, FunctionId string, optionalSetters ...CreateFunctionRuleOption)(*models.ProxyRule, error) {
	path := "/proxy/rules/function"
	options := CreateFunctionRuleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["domain"] = Domain
	params["functionId"] = FunctionId
	if options.enabledSetters["Branch"] {
		params["branch"] = options.Branch
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ProxyRule{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProxyRule
	parsed, ok := resp.Result.(models.ProxyRule)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
									
// CreateRedirectRule create a new proxy rule for to redirect from custom
// domain to another domain.
// 
// Rule ID is automatically generated as MD5 hash of a rule domain for
// performance purposes.
func (srv *Proxy) CreateRedirectRule(Domain string, Url string, StatusCode string, ResourceId string, ResourceType string)(*models.ProxyRule, error) {
	path := "/proxy/rules/redirect"
	params := map[string]interface{}{}
	params["domain"] = Domain
	params["url"] = Url
	params["statusCode"] = StatusCode
	params["resourceId"] = ResourceId
	params["resourceType"] = ResourceType
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ProxyRule{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProxyRule
	parsed, ok := resp.Result.(models.ProxyRule)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateSiteRuleOptions struct {
	Branch string
	enabledSetters map[string]bool
}
func (options CreateSiteRuleOptions) New() *CreateSiteRuleOptions {
	options.enabledSetters = map[string]bool{
		"Branch": false,
	}
	return &options
}
type CreateSiteRuleOption func(*CreateSiteRuleOptions)
func (srv *Proxy) WithCreateSiteRuleBranch(v string) CreateSiteRuleOption {
	return func(o *CreateSiteRuleOptions) {
		o.Branch = v
		o.enabledSetters["Branch"] = true
	}
}
					
// CreateSiteRule create a new proxy rule for serving Appwrite Site on custom
// domain.
// 
// Rule ID is automatically generated as MD5 hash of a rule domain for
// performance purposes.
func (srv *Proxy) CreateSiteRule(Domain string, SiteId string, optionalSetters ...CreateSiteRuleOption)(*models.ProxyRule, error) {
	path := "/proxy/rules/site"
	options := CreateSiteRuleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["domain"] = Domain
	params["siteId"] = SiteId
	if options.enabledSetters["Branch"] {
		params["branch"] = options.Branch
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ProxyRule{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProxyRule
	parsed, ok := resp.Result.(models.ProxyRule)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetRule get a proxy rule by its unique ID.
func (srv *Proxy) GetRule(RuleId string)(*models.ProxyRule, error) {
	r := strings.NewReplacer("{ruleId}", RuleId)
	path := r.Replace("/proxy/rules/{ruleId}")
	params := map[string]interface{}{}
	params["ruleId"] = RuleId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ProxyRule{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProxyRule
	parsed, ok := resp.Result.(models.ProxyRule)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteRule delete a proxy rule by its unique ID.
func (srv *Proxy) DeleteRule(RuleId string)(*interface{}, error) {
	r := strings.NewReplacer("{ruleId}", RuleId)
	path := r.Replace("/proxy/rules/{ruleId}")
	params := map[string]interface{}{}
	params["ruleId"] = RuleId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

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
	
// UpdateRuleStatus if not succeeded yet, retry verification process of a
// proxy rule domain. This endpoint triggers domain verification by checking
// DNS records. If verification is successful, a TLS certificate will be
// automatically provisioned for the domain asynchronously in the background.
func (srv *Proxy) UpdateRuleStatus(RuleId string)(*models.ProxyRule, error) {
	r := strings.NewReplacer("{ruleId}", RuleId)
	path := r.Replace("/proxy/rules/{ruleId}/status")
	params := map[string]interface{}{}
	params["ruleId"] = RuleId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ProxyRule{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProxyRule
	parsed, ok := resp.Result.(models.ProxyRule)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
