package waf

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/models"
)

// Waf service
type Waf struct {
	client client.Client
}

func New(clt client.Client) *Waf {
	return &Waf{
		client: clt,
	}
}

type ListRulesOptions struct {
	Queries        []string
	Search         string
	Total          bool
	enabledSetters map[string]bool
}

func (options ListRulesOptions) New() *ListRulesOptions {
	options.enabledSetters = map[string]bool{"Queries": false, "Search": false, "Total": false}
	return &options
}

type ListRulesOption func(*ListRulesOptions)

func (srv *Waf) WithListRulesQueries(v []string) ListRulesOption {
	return func(o *ListRulesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Waf) WithListRulesSearch(v string) ListRulesOption {
	return func(o *ListRulesOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Waf) WithListRulesTotal(v bool) ListRulesOption {
	return func(o *ListRulesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListRules list WAF rules for the current project.
func (srv *Waf) ListRules(optionalSetters ...ListRulesOption) (*models.WafRuleList, error) {
	path := "/waf/rules"
	options := ListRulesOptions{}.New()
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
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
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

		parsed := models.WafRuleList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.WafRuleList
	parsed, ok := resp.Result.(models.WafRuleList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateBypassRuleOptions struct {
	ResourceId     string
	Description    string
	Priority       int
	Enabled        bool
	Conditions     []string
	enabledSetters map[string]bool
}

func (options CreateBypassRuleOptions) New() *CreateBypassRuleOptions {
	options.enabledSetters = map[string]bool{"ResourceId": false, "Description": false, "Priority": false, "Enabled": false, "Conditions": false}
	return &options
}

type CreateBypassRuleOption func(*CreateBypassRuleOptions)

func (srv *Waf) WithCreateBypassRuleResourceId(v string) CreateBypassRuleOption {
	return func(o *CreateBypassRuleOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
func (srv *Waf) WithCreateBypassRuleDescription(v string) CreateBypassRuleOption {
	return func(o *CreateBypassRuleOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Waf) WithCreateBypassRulePriority(v int) CreateBypassRuleOption {
	return func(o *CreateBypassRuleOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Waf) WithCreateBypassRuleEnabled(v bool) CreateBypassRuleOption {
	return func(o *CreateBypassRuleOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Waf) WithCreateBypassRuleConditions(v []string) CreateBypassRuleOption {
	return func(o *CreateBypassRuleOptions) {
		o.Conditions = v
		o.enabledSetters["Conditions"] = true
	}
}

// CreateBypassRule create a bypass WAF rule. Conditions can match request
// attributes including `ip` (plain IPs or CIDR blocks like `10.0.0.0/8`),
// `method`, `path`, `host`, `country`, `continent`, `headers.<name>`,
// `query.<key>`, `queryKeys`, `userAgent`, `os`, `osVersion`, `browser`, and
// `browserVersion`. Conditions on `city` and `state` require the premium Geo
// DB addon.
func (srv *Waf) CreateBypassRule(RuleId string, ResourceType string, Name string, optionalSetters ...CreateBypassRuleOption) (*models.WafRuleBypass, error) {
	path := "/waf/rules/bypass"
	options := CreateBypassRuleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["ruleId"] = RuleId
	params["resourceType"] = ResourceType
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
	}
	params["name"] = Name
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Conditions"] {
		params["conditions"] = options.Conditions
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

		parsed := models.WafRuleBypass{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.WafRuleBypass
	parsed, ok := resp.Result.(models.WafRuleBypass)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateBypassRuleOptions struct {
	ResourceType   string
	ResourceId     string
	Name           string
	Description    string
	Priority       int
	Enabled        bool
	Conditions     []string
	enabledSetters map[string]bool
}

func (options UpdateBypassRuleOptions) New() *UpdateBypassRuleOptions {
	options.enabledSetters = map[string]bool{"ResourceType": false, "ResourceId": false, "Name": false, "Description": false, "Priority": false, "Enabled": false, "Conditions": false}
	return &options
}

type UpdateBypassRuleOption func(*UpdateBypassRuleOptions)

func (srv *Waf) WithUpdateBypassRuleResourceType(v string) UpdateBypassRuleOption {
	return func(o *UpdateBypassRuleOptions) {
		o.ResourceType = v
		o.enabledSetters["ResourceType"] = true
	}
}
func (srv *Waf) WithUpdateBypassRuleResourceId(v string) UpdateBypassRuleOption {
	return func(o *UpdateBypassRuleOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
func (srv *Waf) WithUpdateBypassRuleName(v string) UpdateBypassRuleOption {
	return func(o *UpdateBypassRuleOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Waf) WithUpdateBypassRuleDescription(v string) UpdateBypassRuleOption {
	return func(o *UpdateBypassRuleOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Waf) WithUpdateBypassRulePriority(v int) UpdateBypassRuleOption {
	return func(o *UpdateBypassRuleOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Waf) WithUpdateBypassRuleEnabled(v bool) UpdateBypassRuleOption {
	return func(o *UpdateBypassRuleOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Waf) WithUpdateBypassRuleConditions(v []string) UpdateBypassRuleOption {
	return func(o *UpdateBypassRuleOptions) {
		o.Conditions = v
		o.enabledSetters["Conditions"] = true
	}
}

// UpdateBypassRule update a bypass WAF rule. Conditions can match request
// attributes including `ip` (plain IPs or CIDR blocks like `10.0.0.0/8`),
// `method`, `path`, `host`, `country`, `continent`, `headers.<name>`,
// `query.<key>`, `queryKeys`, `userAgent`, `os`, `osVersion`, `browser`, and
// `browserVersion`. Conditions on `city` and `state` require the premium Geo
// DB addon.
func (srv *Waf) UpdateBypassRule(RuleId string, optionalSetters ...UpdateBypassRuleOption) (*models.WafRuleBypass, error) {
	if RuleId == "" {
		return nil, errors.New("Missing required parameter: \"ruleId\"")
	}

	r := strings.NewReplacer("{ruleId}", client.EncodePath(RuleId))
	path := r.Replace("/waf/rules/bypass/{ruleId}")
	options := UpdateBypassRuleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ResourceType"] {
		params["resourceType"] = options.ResourceType
	}
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Conditions"] {
		params["conditions"] = options.Conditions
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

		parsed := models.WafRuleBypass{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.WafRuleBypass
	parsed, ok := resp.Result.(models.WafRuleBypass)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateChallengeRuleOptions struct {
	ResourceId     string
	Description    string
	ChallengeType  string
	Priority       int
	Enabled        bool
	Conditions     []string
	Difficulty     int
	Ttl            int
	enabledSetters map[string]bool
}

func (options CreateChallengeRuleOptions) New() *CreateChallengeRuleOptions {
	options.enabledSetters = map[string]bool{"ResourceId": false, "Description": false, "ChallengeType": false, "Priority": false, "Enabled": false, "Conditions": false, "Difficulty": false, "Ttl": false}
	return &options
}

type CreateChallengeRuleOption func(*CreateChallengeRuleOptions)

func (srv *Waf) WithCreateChallengeRuleResourceId(v string) CreateChallengeRuleOption {
	return func(o *CreateChallengeRuleOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
func (srv *Waf) WithCreateChallengeRuleDescription(v string) CreateChallengeRuleOption {
	return func(o *CreateChallengeRuleOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Waf) WithCreateChallengeRuleChallengeType(v string) CreateChallengeRuleOption {
	return func(o *CreateChallengeRuleOptions) {
		o.ChallengeType = v
		o.enabledSetters["ChallengeType"] = true
	}
}
func (srv *Waf) WithCreateChallengeRulePriority(v int) CreateChallengeRuleOption {
	return func(o *CreateChallengeRuleOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Waf) WithCreateChallengeRuleEnabled(v bool) CreateChallengeRuleOption {
	return func(o *CreateChallengeRuleOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Waf) WithCreateChallengeRuleConditions(v []string) CreateChallengeRuleOption {
	return func(o *CreateChallengeRuleOptions) {
		o.Conditions = v
		o.enabledSetters["Conditions"] = true
	}
}
func (srv *Waf) WithCreateChallengeRuleDifficulty(v int) CreateChallengeRuleOption {
	return func(o *CreateChallengeRuleOptions) {
		o.Difficulty = v
		o.enabledSetters["Difficulty"] = true
	}
}
func (srv *Waf) WithCreateChallengeRuleTtl(v int) CreateChallengeRuleOption {
	return func(o *CreateChallengeRuleOptions) {
		o.Ttl = v
		o.enabledSetters["Ttl"] = true
	}
}

// CreateChallengeRule create a challenge WAF rule. Use `difficulty` (1
// easiest to 5 hardest) to tune the client-side proof-of-work cost, and `ttl`
// to control how long, in seconds, a visitor stays cleared after passing the
// challenge before being challenged again. Conditions can match request
// attributes including `ip` (plain IPs or CIDR blocks like `10.0.0.0/8`),
// `method`, `path`, `host`, `country`, `continent`, `headers.<name>`,
// `query.<key>`, `queryKeys`, `userAgent`, `os`, `osVersion`, `browser`, and
// `browserVersion`. Conditions on `city` and `state` require the premium Geo
// DB addon.
func (srv *Waf) CreateChallengeRule(RuleId string, ResourceType string, Name string, optionalSetters ...CreateChallengeRuleOption) (*models.WafRuleChallenge, error) {
	path := "/waf/rules/challenge"
	options := CreateChallengeRuleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["ruleId"] = RuleId
	params["resourceType"] = ResourceType
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
	}
	params["name"] = Name
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["ChallengeType"] {
		params["challengeType"] = options.ChallengeType
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Conditions"] {
		params["conditions"] = options.Conditions
	}
	if options.enabledSetters["Difficulty"] {
		params["difficulty"] = options.Difficulty
	}
	if options.enabledSetters["Ttl"] {
		params["ttl"] = options.Ttl
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

		parsed := models.WafRuleChallenge{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.WafRuleChallenge
	parsed, ok := resp.Result.(models.WafRuleChallenge)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateChallengeRuleOptions struct {
	ResourceType   string
	ResourceId     string
	Name           string
	Description    string
	ChallengeType  string
	Priority       int
	Enabled        bool
	Conditions     []string
	Difficulty     int
	Ttl            int
	enabledSetters map[string]bool
}

func (options UpdateChallengeRuleOptions) New() *UpdateChallengeRuleOptions {
	options.enabledSetters = map[string]bool{"ResourceType": false, "ResourceId": false, "Name": false, "Description": false, "ChallengeType": false, "Priority": false, "Enabled": false, "Conditions": false, "Difficulty": false, "Ttl": false}
	return &options
}

type UpdateChallengeRuleOption func(*UpdateChallengeRuleOptions)

func (srv *Waf) WithUpdateChallengeRuleResourceType(v string) UpdateChallengeRuleOption {
	return func(o *UpdateChallengeRuleOptions) {
		o.ResourceType = v
		o.enabledSetters["ResourceType"] = true
	}
}
func (srv *Waf) WithUpdateChallengeRuleResourceId(v string) UpdateChallengeRuleOption {
	return func(o *UpdateChallengeRuleOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
func (srv *Waf) WithUpdateChallengeRuleName(v string) UpdateChallengeRuleOption {
	return func(o *UpdateChallengeRuleOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Waf) WithUpdateChallengeRuleDescription(v string) UpdateChallengeRuleOption {
	return func(o *UpdateChallengeRuleOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Waf) WithUpdateChallengeRuleChallengeType(v string) UpdateChallengeRuleOption {
	return func(o *UpdateChallengeRuleOptions) {
		o.ChallengeType = v
		o.enabledSetters["ChallengeType"] = true
	}
}
func (srv *Waf) WithUpdateChallengeRulePriority(v int) UpdateChallengeRuleOption {
	return func(o *UpdateChallengeRuleOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Waf) WithUpdateChallengeRuleEnabled(v bool) UpdateChallengeRuleOption {
	return func(o *UpdateChallengeRuleOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Waf) WithUpdateChallengeRuleConditions(v []string) UpdateChallengeRuleOption {
	return func(o *UpdateChallengeRuleOptions) {
		o.Conditions = v
		o.enabledSetters["Conditions"] = true
	}
}
func (srv *Waf) WithUpdateChallengeRuleDifficulty(v int) UpdateChallengeRuleOption {
	return func(o *UpdateChallengeRuleOptions) {
		o.Difficulty = v
		o.enabledSetters["Difficulty"] = true
	}
}
func (srv *Waf) WithUpdateChallengeRuleTtl(v int) UpdateChallengeRuleOption {
	return func(o *UpdateChallengeRuleOptions) {
		o.Ttl = v
		o.enabledSetters["Ttl"] = true
	}
}

// UpdateChallengeRule update a challenge WAF rule. Use `difficulty` (1
// easiest to 5 hardest) to tune the client-side proof-of-work cost, and `ttl`
// to control how long, in seconds, a visitor stays cleared after passing the
// challenge before being challenged again. Conditions can match request
// attributes including `ip` (plain IPs or CIDR blocks like `10.0.0.0/8`),
// `method`, `path`, `host`, `country`, `continent`, `headers.<name>`,
// `query.<key>`, `queryKeys`, `userAgent`, `os`, `osVersion`, `browser`, and
// `browserVersion`. Conditions on `city` and `state` require the premium Geo
// DB addon.
func (srv *Waf) UpdateChallengeRule(RuleId string, optionalSetters ...UpdateChallengeRuleOption) (*models.WafRuleChallenge, error) {
	if RuleId == "" {
		return nil, errors.New("Missing required parameter: \"ruleId\"")
	}

	r := strings.NewReplacer("{ruleId}", client.EncodePath(RuleId))
	path := r.Replace("/waf/rules/challenge/{ruleId}")
	options := UpdateChallengeRuleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ResourceType"] {
		params["resourceType"] = options.ResourceType
	}
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["ChallengeType"] {
		params["challengeType"] = options.ChallengeType
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Conditions"] {
		params["conditions"] = options.Conditions
	}
	if options.enabledSetters["Difficulty"] {
		params["difficulty"] = options.Difficulty
	}
	if options.enabledSetters["Ttl"] {
		params["ttl"] = options.Ttl
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

		parsed := models.WafRuleChallenge{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.WafRuleChallenge
	parsed, ok := resp.Result.(models.WafRuleChallenge)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateDenyRuleOptions struct {
	ResourceId     string
	Description    string
	Priority       int
	Enabled        bool
	Conditions     []string
	enabledSetters map[string]bool
}

func (options CreateDenyRuleOptions) New() *CreateDenyRuleOptions {
	options.enabledSetters = map[string]bool{"ResourceId": false, "Description": false, "Priority": false, "Enabled": false, "Conditions": false}
	return &options
}

type CreateDenyRuleOption func(*CreateDenyRuleOptions)

func (srv *Waf) WithCreateDenyRuleResourceId(v string) CreateDenyRuleOption {
	return func(o *CreateDenyRuleOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
func (srv *Waf) WithCreateDenyRuleDescription(v string) CreateDenyRuleOption {
	return func(o *CreateDenyRuleOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Waf) WithCreateDenyRulePriority(v int) CreateDenyRuleOption {
	return func(o *CreateDenyRuleOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Waf) WithCreateDenyRuleEnabled(v bool) CreateDenyRuleOption {
	return func(o *CreateDenyRuleOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Waf) WithCreateDenyRuleConditions(v []string) CreateDenyRuleOption {
	return func(o *CreateDenyRuleOptions) {
		o.Conditions = v
		o.enabledSetters["Conditions"] = true
	}
}

// CreateDenyRule create a deny WAF rule. Conditions can match request
// attributes including `ip` (plain IPs or CIDR blocks like `10.0.0.0/8`),
// `method`, `path`, `host`, `country`, `continent`, `headers.<name>`,
// `query.<key>`, `queryKeys`, `userAgent`, `os`, `osVersion`, `browser`, and
// `browserVersion`. Conditions on `city` and `state` require the premium Geo
// DB addon.
func (srv *Waf) CreateDenyRule(RuleId string, ResourceType string, Name string, optionalSetters ...CreateDenyRuleOption) (*models.WafRuleDeny, error) {
	path := "/waf/rules/deny"
	options := CreateDenyRuleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["ruleId"] = RuleId
	params["resourceType"] = ResourceType
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
	}
	params["name"] = Name
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Conditions"] {
		params["conditions"] = options.Conditions
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

		parsed := models.WafRuleDeny{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.WafRuleDeny
	parsed, ok := resp.Result.(models.WafRuleDeny)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateDenyRuleOptions struct {
	ResourceType   string
	ResourceId     string
	Name           string
	Description    string
	Priority       int
	Enabled        bool
	Conditions     []string
	enabledSetters map[string]bool
}

func (options UpdateDenyRuleOptions) New() *UpdateDenyRuleOptions {
	options.enabledSetters = map[string]bool{"ResourceType": false, "ResourceId": false, "Name": false, "Description": false, "Priority": false, "Enabled": false, "Conditions": false}
	return &options
}

type UpdateDenyRuleOption func(*UpdateDenyRuleOptions)

func (srv *Waf) WithUpdateDenyRuleResourceType(v string) UpdateDenyRuleOption {
	return func(o *UpdateDenyRuleOptions) {
		o.ResourceType = v
		o.enabledSetters["ResourceType"] = true
	}
}
func (srv *Waf) WithUpdateDenyRuleResourceId(v string) UpdateDenyRuleOption {
	return func(o *UpdateDenyRuleOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
func (srv *Waf) WithUpdateDenyRuleName(v string) UpdateDenyRuleOption {
	return func(o *UpdateDenyRuleOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Waf) WithUpdateDenyRuleDescription(v string) UpdateDenyRuleOption {
	return func(o *UpdateDenyRuleOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Waf) WithUpdateDenyRulePriority(v int) UpdateDenyRuleOption {
	return func(o *UpdateDenyRuleOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Waf) WithUpdateDenyRuleEnabled(v bool) UpdateDenyRuleOption {
	return func(o *UpdateDenyRuleOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Waf) WithUpdateDenyRuleConditions(v []string) UpdateDenyRuleOption {
	return func(o *UpdateDenyRuleOptions) {
		o.Conditions = v
		o.enabledSetters["Conditions"] = true
	}
}

// UpdateDenyRule update a deny WAF rule. Conditions can match request
// attributes including `ip` (plain IPs or CIDR blocks like `10.0.0.0/8`),
// `method`, `path`, `host`, `country`, `continent`, `headers.<name>`,
// `query.<key>`, `queryKeys`, `userAgent`, `os`, `osVersion`, `browser`, and
// `browserVersion`. Conditions on `city` and `state` require the premium Geo
// DB addon.
func (srv *Waf) UpdateDenyRule(RuleId string, optionalSetters ...UpdateDenyRuleOption) (*models.WafRuleDeny, error) {
	if RuleId == "" {
		return nil, errors.New("Missing required parameter: \"ruleId\"")
	}

	r := strings.NewReplacer("{ruleId}", client.EncodePath(RuleId))
	path := r.Replace("/waf/rules/deny/{ruleId}")
	options := UpdateDenyRuleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ResourceType"] {
		params["resourceType"] = options.ResourceType
	}
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Conditions"] {
		params["conditions"] = options.Conditions
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

		parsed := models.WafRuleDeny{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.WafRuleDeny
	parsed, ok := resp.Result.(models.WafRuleDeny)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateRateLimitRuleOptions struct {
	ResourceId     string
	Description    string
	Key            string
	Strategy       string
	MaxBucketSize  int
	Priority       int
	Enabled        bool
	Conditions     []string
	enabledSetters map[string]bool
}

func (options CreateRateLimitRuleOptions) New() *CreateRateLimitRuleOptions {
	options.enabledSetters = map[string]bool{"ResourceId": false, "Description": false, "Key": false, "Strategy": false, "MaxBucketSize": false, "Priority": false, "Enabled": false, "Conditions": false}
	return &options
}

type CreateRateLimitRuleOption func(*CreateRateLimitRuleOptions)

func (srv *Waf) WithCreateRateLimitRuleResourceId(v string) CreateRateLimitRuleOption {
	return func(o *CreateRateLimitRuleOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
func (srv *Waf) WithCreateRateLimitRuleDescription(v string) CreateRateLimitRuleOption {
	return func(o *CreateRateLimitRuleOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Waf) WithCreateRateLimitRuleKey(v string) CreateRateLimitRuleOption {
	return func(o *CreateRateLimitRuleOptions) {
		o.Key = v
		o.enabledSetters["Key"] = true
	}
}
func (srv *Waf) WithCreateRateLimitRuleStrategy(v string) CreateRateLimitRuleOption {
	return func(o *CreateRateLimitRuleOptions) {
		o.Strategy = v
		o.enabledSetters["Strategy"] = true
	}
}
func (srv *Waf) WithCreateRateLimitRuleMaxBucketSize(v int) CreateRateLimitRuleOption {
	return func(o *CreateRateLimitRuleOptions) {
		o.MaxBucketSize = v
		o.enabledSetters["MaxBucketSize"] = true
	}
}
func (srv *Waf) WithCreateRateLimitRulePriority(v int) CreateRateLimitRuleOption {
	return func(o *CreateRateLimitRuleOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Waf) WithCreateRateLimitRuleEnabled(v bool) CreateRateLimitRuleOption {
	return func(o *CreateRateLimitRuleOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Waf) WithCreateRateLimitRuleConditions(v []string) CreateRateLimitRuleOption {
	return func(o *CreateRateLimitRuleOptions) {
		o.Conditions = v
		o.enabledSetters["Conditions"] = true
	}
}

// CreateRateLimitRule create a rate limit WAF rule. Use `key` to choose the
// counter: `ip` limits per client IP, while `userId` limits per authenticated
// user (requests without an authenticated user skip `userId` rules).
// Conditions can match request attributes including `ip` (plain IPs or CIDR
// blocks like `10.0.0.0/8`), `method`, `path`, `host`, `country`,
// `continent`, `headers.<name>`, `query.<key>`, `queryKeys`, `userAgent`,
// `os`, `osVersion`, `browser`, and `browserVersion`. Conditions on `city`
// and `state` require the premium Geo DB addon.
func (srv *Waf) CreateRateLimitRule(RuleId string, ResourceType string, Name string, Limit int, Interval int, optionalSetters ...CreateRateLimitRuleOption) (*models.WafRuleRateLimit, error) {
	path := "/waf/rules/rate-limit"
	options := CreateRateLimitRuleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["ruleId"] = RuleId
	params["resourceType"] = ResourceType
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
	}
	params["name"] = Name
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	params["limit"] = Limit
	params["interval"] = Interval
	if options.enabledSetters["Key"] {
		params["key"] = options.Key
	}
	if options.enabledSetters["Strategy"] {
		params["strategy"] = options.Strategy
	}
	if options.enabledSetters["MaxBucketSize"] {
		params["maxBucketSize"] = options.MaxBucketSize
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Conditions"] {
		params["conditions"] = options.Conditions
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

		parsed := models.WafRuleRateLimit{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.WafRuleRateLimit
	parsed, ok := resp.Result.(models.WafRuleRateLimit)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateRateLimitRuleOptions struct {
	ResourceType   string
	ResourceId     string
	Name           string
	Description    string
	Limit          int
	Interval       int
	Key            string
	MaxBucketSize  int
	Priority       int
	Enabled        bool
	Conditions     []string
	enabledSetters map[string]bool
}

func (options UpdateRateLimitRuleOptions) New() *UpdateRateLimitRuleOptions {
	options.enabledSetters = map[string]bool{"ResourceType": false, "ResourceId": false, "Name": false, "Description": false, "Limit": false, "Interval": false, "Key": false, "MaxBucketSize": false, "Priority": false, "Enabled": false, "Conditions": false}
	return &options
}

type UpdateRateLimitRuleOption func(*UpdateRateLimitRuleOptions)

func (srv *Waf) WithUpdateRateLimitRuleResourceType(v string) UpdateRateLimitRuleOption {
	return func(o *UpdateRateLimitRuleOptions) {
		o.ResourceType = v
		o.enabledSetters["ResourceType"] = true
	}
}
func (srv *Waf) WithUpdateRateLimitRuleResourceId(v string) UpdateRateLimitRuleOption {
	return func(o *UpdateRateLimitRuleOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
func (srv *Waf) WithUpdateRateLimitRuleName(v string) UpdateRateLimitRuleOption {
	return func(o *UpdateRateLimitRuleOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Waf) WithUpdateRateLimitRuleDescription(v string) UpdateRateLimitRuleOption {
	return func(o *UpdateRateLimitRuleOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Waf) WithUpdateRateLimitRuleLimit(v int) UpdateRateLimitRuleOption {
	return func(o *UpdateRateLimitRuleOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Waf) WithUpdateRateLimitRuleInterval(v int) UpdateRateLimitRuleOption {
	return func(o *UpdateRateLimitRuleOptions) {
		o.Interval = v
		o.enabledSetters["Interval"] = true
	}
}
func (srv *Waf) WithUpdateRateLimitRuleKey(v string) UpdateRateLimitRuleOption {
	return func(o *UpdateRateLimitRuleOptions) {
		o.Key = v
		o.enabledSetters["Key"] = true
	}
}
func (srv *Waf) WithUpdateRateLimitRuleMaxBucketSize(v int) UpdateRateLimitRuleOption {
	return func(o *UpdateRateLimitRuleOptions) {
		o.MaxBucketSize = v
		o.enabledSetters["MaxBucketSize"] = true
	}
}
func (srv *Waf) WithUpdateRateLimitRulePriority(v int) UpdateRateLimitRuleOption {
	return func(o *UpdateRateLimitRuleOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Waf) WithUpdateRateLimitRuleEnabled(v bool) UpdateRateLimitRuleOption {
	return func(o *UpdateRateLimitRuleOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Waf) WithUpdateRateLimitRuleConditions(v []string) UpdateRateLimitRuleOption {
	return func(o *UpdateRateLimitRuleOptions) {
		o.Conditions = v
		o.enabledSetters["Conditions"] = true
	}
}

// UpdateRateLimitRule update a rate limit WAF rule. Use `key` to choose the
// counter: `ip` limits per client IP, while `userId` limits per authenticated
// user (requests without an authenticated user skip `userId` rules).
// Conditions can match request attributes including `ip` (plain IPs or CIDR
// blocks like `10.0.0.0/8`), `method`, `path`, `host`, `country`,
// `continent`, `headers.<name>`, `query.<key>`, `queryKeys`, `userAgent`,
// `os`, `osVersion`, `browser`, and `browserVersion`. Conditions on `city`
// and `state` require the premium Geo DB addon.
func (srv *Waf) UpdateRateLimitRule(RuleId string, optionalSetters ...UpdateRateLimitRuleOption) (*models.WafRuleRateLimit, error) {
	if RuleId == "" {
		return nil, errors.New("Missing required parameter: \"ruleId\"")
	}

	r := strings.NewReplacer("{ruleId}", client.EncodePath(RuleId))
	path := r.Replace("/waf/rules/rate-limit/{ruleId}")
	options := UpdateRateLimitRuleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ResourceType"] {
		params["resourceType"] = options.ResourceType
	}
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Interval"] {
		params["interval"] = options.Interval
	}
	if options.enabledSetters["Key"] {
		params["key"] = options.Key
	}
	if options.enabledSetters["MaxBucketSize"] {
		params["maxBucketSize"] = options.MaxBucketSize
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Conditions"] {
		params["conditions"] = options.Conditions
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

		parsed := models.WafRuleRateLimit{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.WafRuleRateLimit
	parsed, ok := resp.Result.(models.WafRuleRateLimit)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type CreateRedirectRuleOptions struct {
	ResourceId     string
	Description    string
	Priority       int
	Enabled        bool
	Conditions     []string
	enabledSetters map[string]bool
}

func (options CreateRedirectRuleOptions) New() *CreateRedirectRuleOptions {
	options.enabledSetters = map[string]bool{"ResourceId": false, "Description": false, "Priority": false, "Enabled": false, "Conditions": false}
	return &options
}

type CreateRedirectRuleOption func(*CreateRedirectRuleOptions)

func (srv *Waf) WithCreateRedirectRuleResourceId(v string) CreateRedirectRuleOption {
	return func(o *CreateRedirectRuleOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
func (srv *Waf) WithCreateRedirectRuleDescription(v string) CreateRedirectRuleOption {
	return func(o *CreateRedirectRuleOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Waf) WithCreateRedirectRulePriority(v int) CreateRedirectRuleOption {
	return func(o *CreateRedirectRuleOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Waf) WithCreateRedirectRuleEnabled(v bool) CreateRedirectRuleOption {
	return func(o *CreateRedirectRuleOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Waf) WithCreateRedirectRuleConditions(v []string) CreateRedirectRuleOption {
	return func(o *CreateRedirectRuleOptions) {
		o.Conditions = v
		o.enabledSetters["Conditions"] = true
	}
}

// CreateRedirectRule create a redirect WAF rule. Conditions can match request
// attributes including `ip` (plain IPs or CIDR blocks like `10.0.0.0/8`),
// `method`, `path`, `host`, `country`, `continent`, `headers.<name>`,
// `query.<key>`, `queryKeys`, `userAgent`, `os`, `osVersion`, `browser`, and
// `browserVersion`. Conditions on `city` and `state` require the premium Geo
// DB addon.
func (srv *Waf) CreateRedirectRule(RuleId string, ResourceType string, Name string, Location string, StatusCode int, optionalSetters ...CreateRedirectRuleOption) (*models.WafRuleRedirect, error) {
	path := "/waf/rules/redirect"
	options := CreateRedirectRuleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["ruleId"] = RuleId
	params["resourceType"] = ResourceType
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
	}
	params["name"] = Name
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	params["location"] = Location
	params["statusCode"] = StatusCode
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Conditions"] {
		params["conditions"] = options.Conditions
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

		parsed := models.WafRuleRedirect{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.WafRuleRedirect
	parsed, ok := resp.Result.(models.WafRuleRedirect)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type UpdateRedirectRuleOptions struct {
	ResourceType   string
	ResourceId     string
	Name           string
	Description    string
	Location       string
	StatusCode     int
	Priority       int
	Enabled        bool
	Conditions     []string
	enabledSetters map[string]bool
}

func (options UpdateRedirectRuleOptions) New() *UpdateRedirectRuleOptions {
	options.enabledSetters = map[string]bool{"ResourceType": false, "ResourceId": false, "Name": false, "Description": false, "Location": false, "StatusCode": false, "Priority": false, "Enabled": false, "Conditions": false}
	return &options
}

type UpdateRedirectRuleOption func(*UpdateRedirectRuleOptions)

func (srv *Waf) WithUpdateRedirectRuleResourceType(v string) UpdateRedirectRuleOption {
	return func(o *UpdateRedirectRuleOptions) {
		o.ResourceType = v
		o.enabledSetters["ResourceType"] = true
	}
}
func (srv *Waf) WithUpdateRedirectRuleResourceId(v string) UpdateRedirectRuleOption {
	return func(o *UpdateRedirectRuleOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
func (srv *Waf) WithUpdateRedirectRuleName(v string) UpdateRedirectRuleOption {
	return func(o *UpdateRedirectRuleOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Waf) WithUpdateRedirectRuleDescription(v string) UpdateRedirectRuleOption {
	return func(o *UpdateRedirectRuleOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Waf) WithUpdateRedirectRuleLocation(v string) UpdateRedirectRuleOption {
	return func(o *UpdateRedirectRuleOptions) {
		o.Location = v
		o.enabledSetters["Location"] = true
	}
}
func (srv *Waf) WithUpdateRedirectRuleStatusCode(v int) UpdateRedirectRuleOption {
	return func(o *UpdateRedirectRuleOptions) {
		o.StatusCode = v
		o.enabledSetters["StatusCode"] = true
	}
}
func (srv *Waf) WithUpdateRedirectRulePriority(v int) UpdateRedirectRuleOption {
	return func(o *UpdateRedirectRuleOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Waf) WithUpdateRedirectRuleEnabled(v bool) UpdateRedirectRuleOption {
	return func(o *UpdateRedirectRuleOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Waf) WithUpdateRedirectRuleConditions(v []string) UpdateRedirectRuleOption {
	return func(o *UpdateRedirectRuleOptions) {
		o.Conditions = v
		o.enabledSetters["Conditions"] = true
	}
}

// UpdateRedirectRule update a redirect WAF rule. Conditions can match request
// attributes including `ip` (plain IPs or CIDR blocks like `10.0.0.0/8`),
// `method`, `path`, `host`, `country`, `continent`, `headers.<name>`,
// `query.<key>`, `queryKeys`, `userAgent`, `os`, `osVersion`, `browser`, and
// `browserVersion`. Conditions on `city` and `state` require the premium Geo
// DB addon.
func (srv *Waf) UpdateRedirectRule(RuleId string, optionalSetters ...UpdateRedirectRuleOption) (*models.WafRuleRedirect, error) {
	if RuleId == "" {
		return nil, errors.New("Missing required parameter: \"ruleId\"")
	}

	r := strings.NewReplacer("{ruleId}", client.EncodePath(RuleId))
	path := r.Replace("/waf/rules/redirect/{ruleId}")
	options := UpdateRedirectRuleOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ResourceType"] {
		params["resourceType"] = options.ResourceType
	}
	if options.enabledSetters["ResourceId"] {
		params["resourceId"] = options.ResourceId
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Location"] {
		params["location"] = options.Location
	}
	if options.enabledSetters["StatusCode"] {
		params["statusCode"] = options.StatusCode
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Conditions"] {
		params["conditions"] = options.Conditions
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

		parsed := models.WafRuleRedirect{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.WafRuleRedirect
	parsed, ok := resp.Result.(models.WafRuleRedirect)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// GetRule get a WAF rule by its ID.
func (srv *Waf) GetRule(RuleId string) (models.Model, error) {
	if RuleId == "" {
		return nil, errors.New("Missing required parameter: \"ruleId\"")
	}

	r := strings.NewReplacer("{ruleId}", client.EncodePath(RuleId))
	path := r.Replace("/waf/rules/{ruleId}")
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

		var response map[string]interface{}
		if err := json.Unmarshal(bytes, &response); err != nil {
			return nil, err
		}
		if fmt.Sprint(response["action"]) == "bypass" {
			parsed := models.WafRuleBypass{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["action"]) == "deny" {
			parsed := models.WafRuleDeny{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["action"]) == "challenge" {
			parsed := models.WafRuleChallenge{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["action"]) == "rateLimit" {
			parsed := models.WafRuleRateLimit{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}
		if fmt.Sprint(response["action"]) == "redirect" {
			parsed := models.WafRuleRedirect{}.New(bytes)
			if err := json.Unmarshal(bytes, parsed); err != nil {
				return nil, err
			}

			return parsed, nil
		}

		return nil, errors.New("unable to match response to any expected response model")
	}
	parsed, ok := resp.Result.(models.Model)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return parsed, nil

}

// DeleteRule delete a WAF rule.
func (srv *Waf) DeleteRule(RuleId string) (*interface{}, error) {
	if RuleId == "" {
		return nil, errors.New("Missing required parameter: \"ruleId\"")
	}

	r := strings.NewReplacer("{ruleId}", client.EncodePath(RuleId))
	path := r.Replace("/waf/rules/{ruleId}")
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
