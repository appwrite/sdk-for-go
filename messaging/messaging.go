package messaging

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
)

// Messaging service
type Messaging struct {
	client client.Client
}

func New(clt client.Client) *Messaging {
	return &Messaging{
		client: clt,
	}
}

type ListMessagesOptions struct {
	Queries []string
	Search string
	enabledSetters map[string]bool
}
func (options ListMessagesOptions) New() *ListMessagesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
	}
	return &options
}
type ListMessagesOption func(*ListMessagesOptions)
func (srv *Messaging) WithListMessagesQueries(v []string) ListMessagesOption {
	return func(o *ListMessagesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Messaging) WithListMessagesSearch(v string) ListMessagesOption {
	return func(o *ListMessagesOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
	
// ListMessages get a list of all messages from the current Appwrite project.
func (srv *Messaging) ListMessages(optionalSetters ...ListMessagesOption)(*models.MessageList, error) {
	path := "/messaging/messages"
	options := ListMessagesOptions{}.New()
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
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MessageList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MessageList
	parsed, ok := resp.Result.(models.MessageList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateEmailOptions struct {
	Topics []string
	Users []string
	Targets []string
	Cc []string
	Bcc []string
	Attachments []string
	Draft bool
	Html bool
	ScheduledAt string
	enabledSetters map[string]bool
}
func (options CreateEmailOptions) New() *CreateEmailOptions {
	options.enabledSetters = map[string]bool{
		"Topics": false,
		"Users": false,
		"Targets": false,
		"Cc": false,
		"Bcc": false,
		"Attachments": false,
		"Draft": false,
		"Html": false,
		"ScheduledAt": false,
	}
	return &options
}
type CreateEmailOption func(*CreateEmailOptions)
func (srv *Messaging) WithCreateEmailTopics(v []string) CreateEmailOption {
	return func(o *CreateEmailOptions) {
		o.Topics = v
		o.enabledSetters["Topics"] = true
	}
}
func (srv *Messaging) WithCreateEmailUsers(v []string) CreateEmailOption {
	return func(o *CreateEmailOptions) {
		o.Users = v
		o.enabledSetters["Users"] = true
	}
}
func (srv *Messaging) WithCreateEmailTargets(v []string) CreateEmailOption {
	return func(o *CreateEmailOptions) {
		o.Targets = v
		o.enabledSetters["Targets"] = true
	}
}
func (srv *Messaging) WithCreateEmailCc(v []string) CreateEmailOption {
	return func(o *CreateEmailOptions) {
		o.Cc = v
		o.enabledSetters["Cc"] = true
	}
}
func (srv *Messaging) WithCreateEmailBcc(v []string) CreateEmailOption {
	return func(o *CreateEmailOptions) {
		o.Bcc = v
		o.enabledSetters["Bcc"] = true
	}
}
func (srv *Messaging) WithCreateEmailAttachments(v []string) CreateEmailOption {
	return func(o *CreateEmailOptions) {
		o.Attachments = v
		o.enabledSetters["Attachments"] = true
	}
}
func (srv *Messaging) WithCreateEmailDraft(v bool) CreateEmailOption {
	return func(o *CreateEmailOptions) {
		o.Draft = v
		o.enabledSetters["Draft"] = true
	}
}
func (srv *Messaging) WithCreateEmailHtml(v bool) CreateEmailOption {
	return func(o *CreateEmailOptions) {
		o.Html = v
		o.enabledSetters["Html"] = true
	}
}
func (srv *Messaging) WithCreateEmailScheduledAt(v string) CreateEmailOption {
	return func(o *CreateEmailOptions) {
		o.ScheduledAt = v
		o.enabledSetters["ScheduledAt"] = true
	}
}
							
// CreateEmail create a new email message.
func (srv *Messaging) CreateEmail(MessageId string, Subject string, Content string, optionalSetters ...CreateEmailOption)(*models.Message, error) {
	path := "/messaging/messages/email"
	options := CreateEmailOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	params["subject"] = Subject
	params["content"] = Content
	if options.enabledSetters["Topics"] {
		params["topics"] = options.Topics
	}
	if options.enabledSetters["Users"] {
		params["users"] = options.Users
	}
	if options.enabledSetters["Targets"] {
		params["targets"] = options.Targets
	}
	if options.enabledSetters["Cc"] {
		params["cc"] = options.Cc
	}
	if options.enabledSetters["Bcc"] {
		params["bcc"] = options.Bcc
	}
	if options.enabledSetters["Attachments"] {
		params["attachments"] = options.Attachments
	}
	if options.enabledSetters["Draft"] {
		params["draft"] = options.Draft
	}
	if options.enabledSetters["Html"] {
		params["html"] = options.Html
	}
	if options.enabledSetters["ScheduledAt"] {
		params["scheduledAt"] = options.ScheduledAt
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

		parsed := models.Message{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Message
	parsed, ok := resp.Result.(models.Message)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateEmailOptions struct {
	Topics []string
	Users []string
	Targets []string
	Subject string
	Content string
	Draft bool
	Html bool
	Cc []string
	Bcc []string
	ScheduledAt string
	Attachments []string
	enabledSetters map[string]bool
}
func (options UpdateEmailOptions) New() *UpdateEmailOptions {
	options.enabledSetters = map[string]bool{
		"Topics": false,
		"Users": false,
		"Targets": false,
		"Subject": false,
		"Content": false,
		"Draft": false,
		"Html": false,
		"Cc": false,
		"Bcc": false,
		"ScheduledAt": false,
		"Attachments": false,
	}
	return &options
}
type UpdateEmailOption func(*UpdateEmailOptions)
func (srv *Messaging) WithUpdateEmailTopics(v []string) UpdateEmailOption {
	return func(o *UpdateEmailOptions) {
		o.Topics = v
		o.enabledSetters["Topics"] = true
	}
}
func (srv *Messaging) WithUpdateEmailUsers(v []string) UpdateEmailOption {
	return func(o *UpdateEmailOptions) {
		o.Users = v
		o.enabledSetters["Users"] = true
	}
}
func (srv *Messaging) WithUpdateEmailTargets(v []string) UpdateEmailOption {
	return func(o *UpdateEmailOptions) {
		o.Targets = v
		o.enabledSetters["Targets"] = true
	}
}
func (srv *Messaging) WithUpdateEmailSubject(v string) UpdateEmailOption {
	return func(o *UpdateEmailOptions) {
		o.Subject = v
		o.enabledSetters["Subject"] = true
	}
}
func (srv *Messaging) WithUpdateEmailContent(v string) UpdateEmailOption {
	return func(o *UpdateEmailOptions) {
		o.Content = v
		o.enabledSetters["Content"] = true
	}
}
func (srv *Messaging) WithUpdateEmailDraft(v bool) UpdateEmailOption {
	return func(o *UpdateEmailOptions) {
		o.Draft = v
		o.enabledSetters["Draft"] = true
	}
}
func (srv *Messaging) WithUpdateEmailHtml(v bool) UpdateEmailOption {
	return func(o *UpdateEmailOptions) {
		o.Html = v
		o.enabledSetters["Html"] = true
	}
}
func (srv *Messaging) WithUpdateEmailCc(v []string) UpdateEmailOption {
	return func(o *UpdateEmailOptions) {
		o.Cc = v
		o.enabledSetters["Cc"] = true
	}
}
func (srv *Messaging) WithUpdateEmailBcc(v []string) UpdateEmailOption {
	return func(o *UpdateEmailOptions) {
		o.Bcc = v
		o.enabledSetters["Bcc"] = true
	}
}
func (srv *Messaging) WithUpdateEmailScheduledAt(v string) UpdateEmailOption {
	return func(o *UpdateEmailOptions) {
		o.ScheduledAt = v
		o.enabledSetters["ScheduledAt"] = true
	}
}
func (srv *Messaging) WithUpdateEmailAttachments(v []string) UpdateEmailOption {
	return func(o *UpdateEmailOptions) {
		o.Attachments = v
		o.enabledSetters["Attachments"] = true
	}
}
			
// UpdateEmail update an email message by its unique ID. This endpoint only
// works on messages that are in draft status. Messages that are already
// processing, sent, or failed cannot be updated.
func (srv *Messaging) UpdateEmail(MessageId string, optionalSetters ...UpdateEmailOption)(*models.Message, error) {
	r := strings.NewReplacer("{messageId}", MessageId)
	path := r.Replace("/messaging/messages/email/{messageId}")
	options := UpdateEmailOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	if options.enabledSetters["Topics"] {
		params["topics"] = options.Topics
	}
	if options.enabledSetters["Users"] {
		params["users"] = options.Users
	}
	if options.enabledSetters["Targets"] {
		params["targets"] = options.Targets
	}
	if options.enabledSetters["Subject"] {
		params["subject"] = options.Subject
	}
	if options.enabledSetters["Content"] {
		params["content"] = options.Content
	}
	if options.enabledSetters["Draft"] {
		params["draft"] = options.Draft
	}
	if options.enabledSetters["Html"] {
		params["html"] = options.Html
	}
	if options.enabledSetters["Cc"] {
		params["cc"] = options.Cc
	}
	if options.enabledSetters["Bcc"] {
		params["bcc"] = options.Bcc
	}
	if options.enabledSetters["ScheduledAt"] {
		params["scheduledAt"] = options.ScheduledAt
	}
	if options.enabledSetters["Attachments"] {
		params["attachments"] = options.Attachments
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Message{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Message
	parsed, ok := resp.Result.(models.Message)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreatePushOptions struct {
	Title string
	Body string
	Topics []string
	Users []string
	Targets []string
	Data interface{}
	Action string
	Image string
	Icon string
	Sound string
	Color string
	Tag string
	Badge int
	Draft bool
	ScheduledAt string
	ContentAvailable bool
	Critical bool
	Priority string
	enabledSetters map[string]bool
}
func (options CreatePushOptions) New() *CreatePushOptions {
	options.enabledSetters = map[string]bool{
		"Title": false,
		"Body": false,
		"Topics": false,
		"Users": false,
		"Targets": false,
		"Data": false,
		"Action": false,
		"Image": false,
		"Icon": false,
		"Sound": false,
		"Color": false,
		"Tag": false,
		"Badge": false,
		"Draft": false,
		"ScheduledAt": false,
		"ContentAvailable": false,
		"Critical": false,
		"Priority": false,
	}
	return &options
}
type CreatePushOption func(*CreatePushOptions)
func (srv *Messaging) WithCreatePushTitle(v string) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *Messaging) WithCreatePushBody(v string) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Body = v
		o.enabledSetters["Body"] = true
	}
}
func (srv *Messaging) WithCreatePushTopics(v []string) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Topics = v
		o.enabledSetters["Topics"] = true
	}
}
func (srv *Messaging) WithCreatePushUsers(v []string) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Users = v
		o.enabledSetters["Users"] = true
	}
}
func (srv *Messaging) WithCreatePushTargets(v []string) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Targets = v
		o.enabledSetters["Targets"] = true
	}
}
func (srv *Messaging) WithCreatePushData(v interface{}) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func (srv *Messaging) WithCreatePushAction(v string) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Action = v
		o.enabledSetters["Action"] = true
	}
}
func (srv *Messaging) WithCreatePushImage(v string) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Image = v
		o.enabledSetters["Image"] = true
	}
}
func (srv *Messaging) WithCreatePushIcon(v string) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Icon = v
		o.enabledSetters["Icon"] = true
	}
}
func (srv *Messaging) WithCreatePushSound(v string) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Sound = v
		o.enabledSetters["Sound"] = true
	}
}
func (srv *Messaging) WithCreatePushColor(v string) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Color = v
		o.enabledSetters["Color"] = true
	}
}
func (srv *Messaging) WithCreatePushTag(v string) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Tag = v
		o.enabledSetters["Tag"] = true
	}
}
func (srv *Messaging) WithCreatePushBadge(v int) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Badge = v
		o.enabledSetters["Badge"] = true
	}
}
func (srv *Messaging) WithCreatePushDraft(v bool) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Draft = v
		o.enabledSetters["Draft"] = true
	}
}
func (srv *Messaging) WithCreatePushScheduledAt(v string) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.ScheduledAt = v
		o.enabledSetters["ScheduledAt"] = true
	}
}
func (srv *Messaging) WithCreatePushContentAvailable(v bool) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.ContentAvailable = v
		o.enabledSetters["ContentAvailable"] = true
	}
}
func (srv *Messaging) WithCreatePushCritical(v bool) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Critical = v
		o.enabledSetters["Critical"] = true
	}
}
func (srv *Messaging) WithCreatePushPriority(v string) CreatePushOption {
	return func(o *CreatePushOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
			
// CreatePush create a new push notification.
func (srv *Messaging) CreatePush(MessageId string, optionalSetters ...CreatePushOption)(*models.Message, error) {
	path := "/messaging/messages/push"
	options := CreatePushOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	if options.enabledSetters["Title"] {
		params["title"] = options.Title
	}
	if options.enabledSetters["Body"] {
		params["body"] = options.Body
	}
	if options.enabledSetters["Topics"] {
		params["topics"] = options.Topics
	}
	if options.enabledSetters["Users"] {
		params["users"] = options.Users
	}
	if options.enabledSetters["Targets"] {
		params["targets"] = options.Targets
	}
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Action"] {
		params["action"] = options.Action
	}
	if options.enabledSetters["Image"] {
		params["image"] = options.Image
	}
	if options.enabledSetters["Icon"] {
		params["icon"] = options.Icon
	}
	if options.enabledSetters["Sound"] {
		params["sound"] = options.Sound
	}
	if options.enabledSetters["Color"] {
		params["color"] = options.Color
	}
	if options.enabledSetters["Tag"] {
		params["tag"] = options.Tag
	}
	if options.enabledSetters["Badge"] {
		params["badge"] = options.Badge
	}
	if options.enabledSetters["Draft"] {
		params["draft"] = options.Draft
	}
	if options.enabledSetters["ScheduledAt"] {
		params["scheduledAt"] = options.ScheduledAt
	}
	if options.enabledSetters["ContentAvailable"] {
		params["contentAvailable"] = options.ContentAvailable
	}
	if options.enabledSetters["Critical"] {
		params["critical"] = options.Critical
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
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

		parsed := models.Message{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Message
	parsed, ok := resp.Result.(models.Message)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdatePushOptions struct {
	Topics []string
	Users []string
	Targets []string
	Title string
	Body string
	Data interface{}
	Action string
	Image string
	Icon string
	Sound string
	Color string
	Tag string
	Badge int
	Draft bool
	ScheduledAt string
	ContentAvailable bool
	Critical bool
	Priority string
	enabledSetters map[string]bool
}
func (options UpdatePushOptions) New() *UpdatePushOptions {
	options.enabledSetters = map[string]bool{
		"Topics": false,
		"Users": false,
		"Targets": false,
		"Title": false,
		"Body": false,
		"Data": false,
		"Action": false,
		"Image": false,
		"Icon": false,
		"Sound": false,
		"Color": false,
		"Tag": false,
		"Badge": false,
		"Draft": false,
		"ScheduledAt": false,
		"ContentAvailable": false,
		"Critical": false,
		"Priority": false,
	}
	return &options
}
type UpdatePushOption func(*UpdatePushOptions)
func (srv *Messaging) WithUpdatePushTopics(v []string) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Topics = v
		o.enabledSetters["Topics"] = true
	}
}
func (srv *Messaging) WithUpdatePushUsers(v []string) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Users = v
		o.enabledSetters["Users"] = true
	}
}
func (srv *Messaging) WithUpdatePushTargets(v []string) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Targets = v
		o.enabledSetters["Targets"] = true
	}
}
func (srv *Messaging) WithUpdatePushTitle(v string) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *Messaging) WithUpdatePushBody(v string) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Body = v
		o.enabledSetters["Body"] = true
	}
}
func (srv *Messaging) WithUpdatePushData(v interface{}) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func (srv *Messaging) WithUpdatePushAction(v string) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Action = v
		o.enabledSetters["Action"] = true
	}
}
func (srv *Messaging) WithUpdatePushImage(v string) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Image = v
		o.enabledSetters["Image"] = true
	}
}
func (srv *Messaging) WithUpdatePushIcon(v string) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Icon = v
		o.enabledSetters["Icon"] = true
	}
}
func (srv *Messaging) WithUpdatePushSound(v string) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Sound = v
		o.enabledSetters["Sound"] = true
	}
}
func (srv *Messaging) WithUpdatePushColor(v string) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Color = v
		o.enabledSetters["Color"] = true
	}
}
func (srv *Messaging) WithUpdatePushTag(v string) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Tag = v
		o.enabledSetters["Tag"] = true
	}
}
func (srv *Messaging) WithUpdatePushBadge(v int) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Badge = v
		o.enabledSetters["Badge"] = true
	}
}
func (srv *Messaging) WithUpdatePushDraft(v bool) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Draft = v
		o.enabledSetters["Draft"] = true
	}
}
func (srv *Messaging) WithUpdatePushScheduledAt(v string) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.ScheduledAt = v
		o.enabledSetters["ScheduledAt"] = true
	}
}
func (srv *Messaging) WithUpdatePushContentAvailable(v bool) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.ContentAvailable = v
		o.enabledSetters["ContentAvailable"] = true
	}
}
func (srv *Messaging) WithUpdatePushCritical(v bool) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Critical = v
		o.enabledSetters["Critical"] = true
	}
}
func (srv *Messaging) WithUpdatePushPriority(v string) UpdatePushOption {
	return func(o *UpdatePushOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
			
// UpdatePush update a push notification by its unique ID. This endpoint only
// works on messages that are in draft status. Messages that are already
// processing, sent, or failed cannot be updated.
func (srv *Messaging) UpdatePush(MessageId string, optionalSetters ...UpdatePushOption)(*models.Message, error) {
	r := strings.NewReplacer("{messageId}", MessageId)
	path := r.Replace("/messaging/messages/push/{messageId}")
	options := UpdatePushOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	if options.enabledSetters["Topics"] {
		params["topics"] = options.Topics
	}
	if options.enabledSetters["Users"] {
		params["users"] = options.Users
	}
	if options.enabledSetters["Targets"] {
		params["targets"] = options.Targets
	}
	if options.enabledSetters["Title"] {
		params["title"] = options.Title
	}
	if options.enabledSetters["Body"] {
		params["body"] = options.Body
	}
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Action"] {
		params["action"] = options.Action
	}
	if options.enabledSetters["Image"] {
		params["image"] = options.Image
	}
	if options.enabledSetters["Icon"] {
		params["icon"] = options.Icon
	}
	if options.enabledSetters["Sound"] {
		params["sound"] = options.Sound
	}
	if options.enabledSetters["Color"] {
		params["color"] = options.Color
	}
	if options.enabledSetters["Tag"] {
		params["tag"] = options.Tag
	}
	if options.enabledSetters["Badge"] {
		params["badge"] = options.Badge
	}
	if options.enabledSetters["Draft"] {
		params["draft"] = options.Draft
	}
	if options.enabledSetters["ScheduledAt"] {
		params["scheduledAt"] = options.ScheduledAt
	}
	if options.enabledSetters["ContentAvailable"] {
		params["contentAvailable"] = options.ContentAvailable
	}
	if options.enabledSetters["Critical"] {
		params["critical"] = options.Critical
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Message{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Message
	parsed, ok := resp.Result.(models.Message)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateSmsOptions struct {
	Topics []string
	Users []string
	Targets []string
	Draft bool
	ScheduledAt string
	enabledSetters map[string]bool
}
func (options CreateSmsOptions) New() *CreateSmsOptions {
	options.enabledSetters = map[string]bool{
		"Topics": false,
		"Users": false,
		"Targets": false,
		"Draft": false,
		"ScheduledAt": false,
	}
	return &options
}
type CreateSmsOption func(*CreateSmsOptions)
func (srv *Messaging) WithCreateSmsTopics(v []string) CreateSmsOption {
	return func(o *CreateSmsOptions) {
		o.Topics = v
		o.enabledSetters["Topics"] = true
	}
}
func (srv *Messaging) WithCreateSmsUsers(v []string) CreateSmsOption {
	return func(o *CreateSmsOptions) {
		o.Users = v
		o.enabledSetters["Users"] = true
	}
}
func (srv *Messaging) WithCreateSmsTargets(v []string) CreateSmsOption {
	return func(o *CreateSmsOptions) {
		o.Targets = v
		o.enabledSetters["Targets"] = true
	}
}
func (srv *Messaging) WithCreateSmsDraft(v bool) CreateSmsOption {
	return func(o *CreateSmsOptions) {
		o.Draft = v
		o.enabledSetters["Draft"] = true
	}
}
func (srv *Messaging) WithCreateSmsScheduledAt(v string) CreateSmsOption {
	return func(o *CreateSmsOptions) {
		o.ScheduledAt = v
		o.enabledSetters["ScheduledAt"] = true
	}
}
					
// CreateSms create a new SMS message.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `Messaging.createSMS` instead.
func (srv *Messaging) CreateSms(MessageId string, Content string, optionalSetters ...CreateSmsOption)(*models.Message, error) {
	path := "/messaging/messages/sms"
	options := CreateSmsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	params["content"] = Content
	if options.enabledSetters["Topics"] {
		params["topics"] = options.Topics
	}
	if options.enabledSetters["Users"] {
		params["users"] = options.Users
	}
	if options.enabledSetters["Targets"] {
		params["targets"] = options.Targets
	}
	if options.enabledSetters["Draft"] {
		params["draft"] = options.Draft
	}
	if options.enabledSetters["ScheduledAt"] {
		params["scheduledAt"] = options.ScheduledAt
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

		parsed := models.Message{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Message
	parsed, ok := resp.Result.(models.Message)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateSMSOptions struct {
	Topics []string
	Users []string
	Targets []string
	Draft bool
	ScheduledAt string
	enabledSetters map[string]bool
}
func (options CreateSMSOptions) New() *CreateSMSOptions {
	options.enabledSetters = map[string]bool{
		"Topics": false,
		"Users": false,
		"Targets": false,
		"Draft": false,
		"ScheduledAt": false,
	}
	return &options
}
type CreateSMSOption func(*CreateSMSOptions)
func (srv *Messaging) WithCreateSMSTopics(v []string) CreateSMSOption {
	return func(o *CreateSMSOptions) {
		o.Topics = v
		o.enabledSetters["Topics"] = true
	}
}
func (srv *Messaging) WithCreateSMSUsers(v []string) CreateSMSOption {
	return func(o *CreateSMSOptions) {
		o.Users = v
		o.enabledSetters["Users"] = true
	}
}
func (srv *Messaging) WithCreateSMSTargets(v []string) CreateSMSOption {
	return func(o *CreateSMSOptions) {
		o.Targets = v
		o.enabledSetters["Targets"] = true
	}
}
func (srv *Messaging) WithCreateSMSDraft(v bool) CreateSMSOption {
	return func(o *CreateSMSOptions) {
		o.Draft = v
		o.enabledSetters["Draft"] = true
	}
}
func (srv *Messaging) WithCreateSMSScheduledAt(v string) CreateSMSOption {
	return func(o *CreateSMSOptions) {
		o.ScheduledAt = v
		o.enabledSetters["ScheduledAt"] = true
	}
}
					
// CreateSMS create a new SMS message.
func (srv *Messaging) CreateSMS(MessageId string, Content string, optionalSetters ...CreateSMSOption)(*models.Message, error) {
	path := "/messaging/messages/sms"
	options := CreateSMSOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	params["content"] = Content
	if options.enabledSetters["Topics"] {
		params["topics"] = options.Topics
	}
	if options.enabledSetters["Users"] {
		params["users"] = options.Users
	}
	if options.enabledSetters["Targets"] {
		params["targets"] = options.Targets
	}
	if options.enabledSetters["Draft"] {
		params["draft"] = options.Draft
	}
	if options.enabledSetters["ScheduledAt"] {
		params["scheduledAt"] = options.ScheduledAt
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

		parsed := models.Message{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Message
	parsed, ok := resp.Result.(models.Message)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateSmsOptions struct {
	Topics []string
	Users []string
	Targets []string
	Content string
	Draft bool
	ScheduledAt string
	enabledSetters map[string]bool
}
func (options UpdateSmsOptions) New() *UpdateSmsOptions {
	options.enabledSetters = map[string]bool{
		"Topics": false,
		"Users": false,
		"Targets": false,
		"Content": false,
		"Draft": false,
		"ScheduledAt": false,
	}
	return &options
}
type UpdateSmsOption func(*UpdateSmsOptions)
func (srv *Messaging) WithUpdateSmsTopics(v []string) UpdateSmsOption {
	return func(o *UpdateSmsOptions) {
		o.Topics = v
		o.enabledSetters["Topics"] = true
	}
}
func (srv *Messaging) WithUpdateSmsUsers(v []string) UpdateSmsOption {
	return func(o *UpdateSmsOptions) {
		o.Users = v
		o.enabledSetters["Users"] = true
	}
}
func (srv *Messaging) WithUpdateSmsTargets(v []string) UpdateSmsOption {
	return func(o *UpdateSmsOptions) {
		o.Targets = v
		o.enabledSetters["Targets"] = true
	}
}
func (srv *Messaging) WithUpdateSmsContent(v string) UpdateSmsOption {
	return func(o *UpdateSmsOptions) {
		o.Content = v
		o.enabledSetters["Content"] = true
	}
}
func (srv *Messaging) WithUpdateSmsDraft(v bool) UpdateSmsOption {
	return func(o *UpdateSmsOptions) {
		o.Draft = v
		o.enabledSetters["Draft"] = true
	}
}
func (srv *Messaging) WithUpdateSmsScheduledAt(v string) UpdateSmsOption {
	return func(o *UpdateSmsOptions) {
		o.ScheduledAt = v
		o.enabledSetters["ScheduledAt"] = true
	}
}
			
// UpdateSms update an SMS message by its unique ID. This endpoint only works
// on messages that are in draft status. Messages that are already processing,
// sent, or failed cannot be updated.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `Messaging.updateSMS` instead.
func (srv *Messaging) UpdateSms(MessageId string, optionalSetters ...UpdateSmsOption)(*models.Message, error) {
	r := strings.NewReplacer("{messageId}", MessageId)
	path := r.Replace("/messaging/messages/sms/{messageId}")
	options := UpdateSmsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	if options.enabledSetters["Topics"] {
		params["topics"] = options.Topics
	}
	if options.enabledSetters["Users"] {
		params["users"] = options.Users
	}
	if options.enabledSetters["Targets"] {
		params["targets"] = options.Targets
	}
	if options.enabledSetters["Content"] {
		params["content"] = options.Content
	}
	if options.enabledSetters["Draft"] {
		params["draft"] = options.Draft
	}
	if options.enabledSetters["ScheduledAt"] {
		params["scheduledAt"] = options.ScheduledAt
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Message{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Message
	parsed, ok := resp.Result.(models.Message)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateSMSOptions struct {
	Topics []string
	Users []string
	Targets []string
	Content string
	Draft bool
	ScheduledAt string
	enabledSetters map[string]bool
}
func (options UpdateSMSOptions) New() *UpdateSMSOptions {
	options.enabledSetters = map[string]bool{
		"Topics": false,
		"Users": false,
		"Targets": false,
		"Content": false,
		"Draft": false,
		"ScheduledAt": false,
	}
	return &options
}
type UpdateSMSOption func(*UpdateSMSOptions)
func (srv *Messaging) WithUpdateSMSTopics(v []string) UpdateSMSOption {
	return func(o *UpdateSMSOptions) {
		o.Topics = v
		o.enabledSetters["Topics"] = true
	}
}
func (srv *Messaging) WithUpdateSMSUsers(v []string) UpdateSMSOption {
	return func(o *UpdateSMSOptions) {
		o.Users = v
		o.enabledSetters["Users"] = true
	}
}
func (srv *Messaging) WithUpdateSMSTargets(v []string) UpdateSMSOption {
	return func(o *UpdateSMSOptions) {
		o.Targets = v
		o.enabledSetters["Targets"] = true
	}
}
func (srv *Messaging) WithUpdateSMSContent(v string) UpdateSMSOption {
	return func(o *UpdateSMSOptions) {
		o.Content = v
		o.enabledSetters["Content"] = true
	}
}
func (srv *Messaging) WithUpdateSMSDraft(v bool) UpdateSMSOption {
	return func(o *UpdateSMSOptions) {
		o.Draft = v
		o.enabledSetters["Draft"] = true
	}
}
func (srv *Messaging) WithUpdateSMSScheduledAt(v string) UpdateSMSOption {
	return func(o *UpdateSMSOptions) {
		o.ScheduledAt = v
		o.enabledSetters["ScheduledAt"] = true
	}
}
			
// UpdateSMS update an SMS message by its unique ID. This endpoint only works
// on messages that are in draft status. Messages that are already processing,
// sent, or failed cannot be updated.
func (srv *Messaging) UpdateSMS(MessageId string, optionalSetters ...UpdateSMSOption)(*models.Message, error) {
	r := strings.NewReplacer("{messageId}", MessageId)
	path := r.Replace("/messaging/messages/sms/{messageId}")
	options := UpdateSMSOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	if options.enabledSetters["Topics"] {
		params["topics"] = options.Topics
	}
	if options.enabledSetters["Users"] {
		params["users"] = options.Users
	}
	if options.enabledSetters["Targets"] {
		params["targets"] = options.Targets
	}
	if options.enabledSetters["Content"] {
		params["content"] = options.Content
	}
	if options.enabledSetters["Draft"] {
		params["draft"] = options.Draft
	}
	if options.enabledSetters["ScheduledAt"] {
		params["scheduledAt"] = options.ScheduledAt
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Message{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Message
	parsed, ok := resp.Result.(models.Message)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetMessage get a message by its unique ID.
func (srv *Messaging) GetMessage(MessageId string)(*models.Message, error) {
	r := strings.NewReplacer("{messageId}", MessageId)
	path := r.Replace("/messaging/messages/{messageId}")
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Message{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Message
	parsed, ok := resp.Result.(models.Message)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// Delete delete a message. If the message is not a draft or scheduled, but
// has been sent, this will not recall the message.
func (srv *Messaging) Delete(MessageId string)(*interface{}, error) {
	r := strings.NewReplacer("{messageId}", MessageId)
	path := r.Replace("/messaging/messages/{messageId}")
	params := map[string]interface{}{}
	params["messageId"] = MessageId
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
type ListMessageLogsOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListMessageLogsOptions) New() *ListMessageLogsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListMessageLogsOption func(*ListMessageLogsOptions)
func (srv *Messaging) WithListMessageLogsQueries(v []string) ListMessageLogsOption {
	return func(o *ListMessageLogsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
			
// ListMessageLogs get the message activity logs listed by its unique ID.
func (srv *Messaging) ListMessageLogs(MessageId string, optionalSetters ...ListMessageLogsOption)(*models.LogList, error) {
	r := strings.NewReplacer("{messageId}", MessageId)
	path := r.Replace("/messaging/messages/{messageId}/logs")
	options := ListMessageLogsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.LogList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LogList
	parsed, ok := resp.Result.(models.LogList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListTargetsOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListTargetsOptions) New() *ListTargetsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListTargetsOption func(*ListTargetsOptions)
func (srv *Messaging) WithListTargetsQueries(v []string) ListTargetsOption {
	return func(o *ListTargetsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
			
// ListTargets get a list of the targets associated with a message.
func (srv *Messaging) ListTargets(MessageId string, optionalSetters ...ListTargetsOption)(*models.TargetList, error) {
	r := strings.NewReplacer("{messageId}", MessageId)
	path := r.Replace("/messaging/messages/{messageId}/targets")
	options := ListTargetsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.TargetList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.TargetList
	parsed, ok := resp.Result.(models.TargetList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListProvidersOptions struct {
	Queries []string
	Search string
	enabledSetters map[string]bool
}
func (options ListProvidersOptions) New() *ListProvidersOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
	}
	return &options
}
type ListProvidersOption func(*ListProvidersOptions)
func (srv *Messaging) WithListProvidersQueries(v []string) ListProvidersOption {
	return func(o *ListProvidersOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Messaging) WithListProvidersSearch(v string) ListProvidersOption {
	return func(o *ListProvidersOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
	
// ListProviders get a list of all providers from the current Appwrite
// project.
func (srv *Messaging) ListProviders(optionalSetters ...ListProvidersOption)(*models.ProviderList, error) {
	path := "/messaging/providers"
	options := ListProvidersOptions{}.New()
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
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ProviderList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProviderList
	parsed, ok := resp.Result.(models.ProviderList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateApnsProviderOptions struct {
	AuthKey string
	AuthKeyId string
	TeamId string
	BundleId string
	Sandbox bool
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateApnsProviderOptions) New() *CreateApnsProviderOptions {
	options.enabledSetters = map[string]bool{
		"AuthKey": false,
		"AuthKeyId": false,
		"TeamId": false,
		"BundleId": false,
		"Sandbox": false,
		"Enabled": false,
	}
	return &options
}
type CreateApnsProviderOption func(*CreateApnsProviderOptions)
func (srv *Messaging) WithCreateApnsProviderAuthKey(v string) CreateApnsProviderOption {
	return func(o *CreateApnsProviderOptions) {
		o.AuthKey = v
		o.enabledSetters["AuthKey"] = true
	}
}
func (srv *Messaging) WithCreateApnsProviderAuthKeyId(v string) CreateApnsProviderOption {
	return func(o *CreateApnsProviderOptions) {
		o.AuthKeyId = v
		o.enabledSetters["AuthKeyId"] = true
	}
}
func (srv *Messaging) WithCreateApnsProviderTeamId(v string) CreateApnsProviderOption {
	return func(o *CreateApnsProviderOptions) {
		o.TeamId = v
		o.enabledSetters["TeamId"] = true
	}
}
func (srv *Messaging) WithCreateApnsProviderBundleId(v string) CreateApnsProviderOption {
	return func(o *CreateApnsProviderOptions) {
		o.BundleId = v
		o.enabledSetters["BundleId"] = true
	}
}
func (srv *Messaging) WithCreateApnsProviderSandbox(v bool) CreateApnsProviderOption {
	return func(o *CreateApnsProviderOptions) {
		o.Sandbox = v
		o.enabledSetters["Sandbox"] = true
	}
}
func (srv *Messaging) WithCreateApnsProviderEnabled(v bool) CreateApnsProviderOption {
	return func(o *CreateApnsProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
					
// CreateApnsProvider create a new Apple Push Notification service provider.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `Messaging.createAPNSProvider` instead.
func (srv *Messaging) CreateApnsProvider(ProviderId string, Name string, optionalSetters ...CreateApnsProviderOption)(*models.Provider, error) {
	path := "/messaging/providers/apns"
	options := CreateApnsProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	params["name"] = Name
	if options.enabledSetters["AuthKey"] {
		params["authKey"] = options.AuthKey
	}
	if options.enabledSetters["AuthKeyId"] {
		params["authKeyId"] = options.AuthKeyId
	}
	if options.enabledSetters["TeamId"] {
		params["teamId"] = options.TeamId
	}
	if options.enabledSetters["BundleId"] {
		params["bundleId"] = options.BundleId
	}
	if options.enabledSetters["Sandbox"] {
		params["sandbox"] = options.Sandbox
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateAPNSProviderOptions struct {
	AuthKey string
	AuthKeyId string
	TeamId string
	BundleId string
	Sandbox bool
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateAPNSProviderOptions) New() *CreateAPNSProviderOptions {
	options.enabledSetters = map[string]bool{
		"AuthKey": false,
		"AuthKeyId": false,
		"TeamId": false,
		"BundleId": false,
		"Sandbox": false,
		"Enabled": false,
	}
	return &options
}
type CreateAPNSProviderOption func(*CreateAPNSProviderOptions)
func (srv *Messaging) WithCreateAPNSProviderAuthKey(v string) CreateAPNSProviderOption {
	return func(o *CreateAPNSProviderOptions) {
		o.AuthKey = v
		o.enabledSetters["AuthKey"] = true
	}
}
func (srv *Messaging) WithCreateAPNSProviderAuthKeyId(v string) CreateAPNSProviderOption {
	return func(o *CreateAPNSProviderOptions) {
		o.AuthKeyId = v
		o.enabledSetters["AuthKeyId"] = true
	}
}
func (srv *Messaging) WithCreateAPNSProviderTeamId(v string) CreateAPNSProviderOption {
	return func(o *CreateAPNSProviderOptions) {
		o.TeamId = v
		o.enabledSetters["TeamId"] = true
	}
}
func (srv *Messaging) WithCreateAPNSProviderBundleId(v string) CreateAPNSProviderOption {
	return func(o *CreateAPNSProviderOptions) {
		o.BundleId = v
		o.enabledSetters["BundleId"] = true
	}
}
func (srv *Messaging) WithCreateAPNSProviderSandbox(v bool) CreateAPNSProviderOption {
	return func(o *CreateAPNSProviderOptions) {
		o.Sandbox = v
		o.enabledSetters["Sandbox"] = true
	}
}
func (srv *Messaging) WithCreateAPNSProviderEnabled(v bool) CreateAPNSProviderOption {
	return func(o *CreateAPNSProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
					
// CreateAPNSProvider create a new Apple Push Notification service provider.
func (srv *Messaging) CreateAPNSProvider(ProviderId string, Name string, optionalSetters ...CreateAPNSProviderOption)(*models.Provider, error) {
	path := "/messaging/providers/apns"
	options := CreateAPNSProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	params["name"] = Name
	if options.enabledSetters["AuthKey"] {
		params["authKey"] = options.AuthKey
	}
	if options.enabledSetters["AuthKeyId"] {
		params["authKeyId"] = options.AuthKeyId
	}
	if options.enabledSetters["TeamId"] {
		params["teamId"] = options.TeamId
	}
	if options.enabledSetters["BundleId"] {
		params["bundleId"] = options.BundleId
	}
	if options.enabledSetters["Sandbox"] {
		params["sandbox"] = options.Sandbox
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateApnsProviderOptions struct {
	Name string
	Enabled bool
	AuthKey string
	AuthKeyId string
	TeamId string
	BundleId string
	Sandbox bool
	enabledSetters map[string]bool
}
func (options UpdateApnsProviderOptions) New() *UpdateApnsProviderOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Enabled": false,
		"AuthKey": false,
		"AuthKeyId": false,
		"TeamId": false,
		"BundleId": false,
		"Sandbox": false,
	}
	return &options
}
type UpdateApnsProviderOption func(*UpdateApnsProviderOptions)
func (srv *Messaging) WithUpdateApnsProviderName(v string) UpdateApnsProviderOption {
	return func(o *UpdateApnsProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithUpdateApnsProviderEnabled(v bool) UpdateApnsProviderOption {
	return func(o *UpdateApnsProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithUpdateApnsProviderAuthKey(v string) UpdateApnsProviderOption {
	return func(o *UpdateApnsProviderOptions) {
		o.AuthKey = v
		o.enabledSetters["AuthKey"] = true
	}
}
func (srv *Messaging) WithUpdateApnsProviderAuthKeyId(v string) UpdateApnsProviderOption {
	return func(o *UpdateApnsProviderOptions) {
		o.AuthKeyId = v
		o.enabledSetters["AuthKeyId"] = true
	}
}
func (srv *Messaging) WithUpdateApnsProviderTeamId(v string) UpdateApnsProviderOption {
	return func(o *UpdateApnsProviderOptions) {
		o.TeamId = v
		o.enabledSetters["TeamId"] = true
	}
}
func (srv *Messaging) WithUpdateApnsProviderBundleId(v string) UpdateApnsProviderOption {
	return func(o *UpdateApnsProviderOptions) {
		o.BundleId = v
		o.enabledSetters["BundleId"] = true
	}
}
func (srv *Messaging) WithUpdateApnsProviderSandbox(v bool) UpdateApnsProviderOption {
	return func(o *UpdateApnsProviderOptions) {
		o.Sandbox = v
		o.enabledSetters["Sandbox"] = true
	}
}
			
// UpdateApnsProvider update a Apple Push Notification service provider by its
// unique ID.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `Messaging.updateAPNSProvider` instead.
func (srv *Messaging) UpdateApnsProvider(ProviderId string, optionalSetters ...UpdateApnsProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/apns/{providerId}")
	options := UpdateApnsProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["AuthKey"] {
		params["authKey"] = options.AuthKey
	}
	if options.enabledSetters["AuthKeyId"] {
		params["authKeyId"] = options.AuthKeyId
	}
	if options.enabledSetters["TeamId"] {
		params["teamId"] = options.TeamId
	}
	if options.enabledSetters["BundleId"] {
		params["bundleId"] = options.BundleId
	}
	if options.enabledSetters["Sandbox"] {
		params["sandbox"] = options.Sandbox
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateAPNSProviderOptions struct {
	Name string
	Enabled bool
	AuthKey string
	AuthKeyId string
	TeamId string
	BundleId string
	Sandbox bool
	enabledSetters map[string]bool
}
func (options UpdateAPNSProviderOptions) New() *UpdateAPNSProviderOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Enabled": false,
		"AuthKey": false,
		"AuthKeyId": false,
		"TeamId": false,
		"BundleId": false,
		"Sandbox": false,
	}
	return &options
}
type UpdateAPNSProviderOption func(*UpdateAPNSProviderOptions)
func (srv *Messaging) WithUpdateAPNSProviderName(v string) UpdateAPNSProviderOption {
	return func(o *UpdateAPNSProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithUpdateAPNSProviderEnabled(v bool) UpdateAPNSProviderOption {
	return func(o *UpdateAPNSProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithUpdateAPNSProviderAuthKey(v string) UpdateAPNSProviderOption {
	return func(o *UpdateAPNSProviderOptions) {
		o.AuthKey = v
		o.enabledSetters["AuthKey"] = true
	}
}
func (srv *Messaging) WithUpdateAPNSProviderAuthKeyId(v string) UpdateAPNSProviderOption {
	return func(o *UpdateAPNSProviderOptions) {
		o.AuthKeyId = v
		o.enabledSetters["AuthKeyId"] = true
	}
}
func (srv *Messaging) WithUpdateAPNSProviderTeamId(v string) UpdateAPNSProviderOption {
	return func(o *UpdateAPNSProviderOptions) {
		o.TeamId = v
		o.enabledSetters["TeamId"] = true
	}
}
func (srv *Messaging) WithUpdateAPNSProviderBundleId(v string) UpdateAPNSProviderOption {
	return func(o *UpdateAPNSProviderOptions) {
		o.BundleId = v
		o.enabledSetters["BundleId"] = true
	}
}
func (srv *Messaging) WithUpdateAPNSProviderSandbox(v bool) UpdateAPNSProviderOption {
	return func(o *UpdateAPNSProviderOptions) {
		o.Sandbox = v
		o.enabledSetters["Sandbox"] = true
	}
}
			
// UpdateAPNSProvider update a Apple Push Notification service provider by its
// unique ID.
func (srv *Messaging) UpdateAPNSProvider(ProviderId string, optionalSetters ...UpdateAPNSProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/apns/{providerId}")
	options := UpdateAPNSProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["AuthKey"] {
		params["authKey"] = options.AuthKey
	}
	if options.enabledSetters["AuthKeyId"] {
		params["authKeyId"] = options.AuthKeyId
	}
	if options.enabledSetters["TeamId"] {
		params["teamId"] = options.TeamId
	}
	if options.enabledSetters["BundleId"] {
		params["bundleId"] = options.BundleId
	}
	if options.enabledSetters["Sandbox"] {
		params["sandbox"] = options.Sandbox
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateFcmProviderOptions struct {
	ServiceAccountJSON interface{}
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateFcmProviderOptions) New() *CreateFcmProviderOptions {
	options.enabledSetters = map[string]bool{
		"ServiceAccountJSON": false,
		"Enabled": false,
	}
	return &options
}
type CreateFcmProviderOption func(*CreateFcmProviderOptions)
func (srv *Messaging) WithCreateFcmProviderServiceAccountJSON(v interface{}) CreateFcmProviderOption {
	return func(o *CreateFcmProviderOptions) {
		o.ServiceAccountJSON = v
		o.enabledSetters["ServiceAccountJSON"] = true
	}
}
func (srv *Messaging) WithCreateFcmProviderEnabled(v bool) CreateFcmProviderOption {
	return func(o *CreateFcmProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
					
// CreateFcmProvider create a new Firebase Cloud Messaging provider.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `Messaging.createFCMProvider` instead.
func (srv *Messaging) CreateFcmProvider(ProviderId string, Name string, optionalSetters ...CreateFcmProviderOption)(*models.Provider, error) {
	path := "/messaging/providers/fcm"
	options := CreateFcmProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	params["name"] = Name
	if options.enabledSetters["ServiceAccountJSON"] {
		params["serviceAccountJSON"] = options.ServiceAccountJSON
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateFCMProviderOptions struct {
	ServiceAccountJSON interface{}
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateFCMProviderOptions) New() *CreateFCMProviderOptions {
	options.enabledSetters = map[string]bool{
		"ServiceAccountJSON": false,
		"Enabled": false,
	}
	return &options
}
type CreateFCMProviderOption func(*CreateFCMProviderOptions)
func (srv *Messaging) WithCreateFCMProviderServiceAccountJSON(v interface{}) CreateFCMProviderOption {
	return func(o *CreateFCMProviderOptions) {
		o.ServiceAccountJSON = v
		o.enabledSetters["ServiceAccountJSON"] = true
	}
}
func (srv *Messaging) WithCreateFCMProviderEnabled(v bool) CreateFCMProviderOption {
	return func(o *CreateFCMProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
					
// CreateFCMProvider create a new Firebase Cloud Messaging provider.
func (srv *Messaging) CreateFCMProvider(ProviderId string, Name string, optionalSetters ...CreateFCMProviderOption)(*models.Provider, error) {
	path := "/messaging/providers/fcm"
	options := CreateFCMProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	params["name"] = Name
	if options.enabledSetters["ServiceAccountJSON"] {
		params["serviceAccountJSON"] = options.ServiceAccountJSON
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateFcmProviderOptions struct {
	Name string
	Enabled bool
	ServiceAccountJSON interface{}
	enabledSetters map[string]bool
}
func (options UpdateFcmProviderOptions) New() *UpdateFcmProviderOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Enabled": false,
		"ServiceAccountJSON": false,
	}
	return &options
}
type UpdateFcmProviderOption func(*UpdateFcmProviderOptions)
func (srv *Messaging) WithUpdateFcmProviderName(v string) UpdateFcmProviderOption {
	return func(o *UpdateFcmProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithUpdateFcmProviderEnabled(v bool) UpdateFcmProviderOption {
	return func(o *UpdateFcmProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithUpdateFcmProviderServiceAccountJSON(v interface{}) UpdateFcmProviderOption {
	return func(o *UpdateFcmProviderOptions) {
		o.ServiceAccountJSON = v
		o.enabledSetters["ServiceAccountJSON"] = true
	}
}
			
// UpdateFcmProvider update a Firebase Cloud Messaging provider by its unique
// ID.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `Messaging.updateFCMProvider` instead.
func (srv *Messaging) UpdateFcmProvider(ProviderId string, optionalSetters ...UpdateFcmProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/fcm/{providerId}")
	options := UpdateFcmProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["ServiceAccountJSON"] {
		params["serviceAccountJSON"] = options.ServiceAccountJSON
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateFCMProviderOptions struct {
	Name string
	Enabled bool
	ServiceAccountJSON interface{}
	enabledSetters map[string]bool
}
func (options UpdateFCMProviderOptions) New() *UpdateFCMProviderOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Enabled": false,
		"ServiceAccountJSON": false,
	}
	return &options
}
type UpdateFCMProviderOption func(*UpdateFCMProviderOptions)
func (srv *Messaging) WithUpdateFCMProviderName(v string) UpdateFCMProviderOption {
	return func(o *UpdateFCMProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithUpdateFCMProviderEnabled(v bool) UpdateFCMProviderOption {
	return func(o *UpdateFCMProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithUpdateFCMProviderServiceAccountJSON(v interface{}) UpdateFCMProviderOption {
	return func(o *UpdateFCMProviderOptions) {
		o.ServiceAccountJSON = v
		o.enabledSetters["ServiceAccountJSON"] = true
	}
}
			
// UpdateFCMProvider update a Firebase Cloud Messaging provider by its unique
// ID.
func (srv *Messaging) UpdateFCMProvider(ProviderId string, optionalSetters ...UpdateFCMProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/fcm/{providerId}")
	options := UpdateFCMProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["ServiceAccountJSON"] {
		params["serviceAccountJSON"] = options.ServiceAccountJSON
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateMailgunProviderOptions struct {
	ApiKey string
	Domain string
	IsEuRegion bool
	FromName string
	FromEmail string
	ReplyToName string
	ReplyToEmail string
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateMailgunProviderOptions) New() *CreateMailgunProviderOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"Domain": false,
		"IsEuRegion": false,
		"FromName": false,
		"FromEmail": false,
		"ReplyToName": false,
		"ReplyToEmail": false,
		"Enabled": false,
	}
	return &options
}
type CreateMailgunProviderOption func(*CreateMailgunProviderOptions)
func (srv *Messaging) WithCreateMailgunProviderApiKey(v string) CreateMailgunProviderOption {
	return func(o *CreateMailgunProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithCreateMailgunProviderDomain(v string) CreateMailgunProviderOption {
	return func(o *CreateMailgunProviderOptions) {
		o.Domain = v
		o.enabledSetters["Domain"] = true
	}
}
func (srv *Messaging) WithCreateMailgunProviderIsEuRegion(v bool) CreateMailgunProviderOption {
	return func(o *CreateMailgunProviderOptions) {
		o.IsEuRegion = v
		o.enabledSetters["IsEuRegion"] = true
	}
}
func (srv *Messaging) WithCreateMailgunProviderFromName(v string) CreateMailgunProviderOption {
	return func(o *CreateMailgunProviderOptions) {
		o.FromName = v
		o.enabledSetters["FromName"] = true
	}
}
func (srv *Messaging) WithCreateMailgunProviderFromEmail(v string) CreateMailgunProviderOption {
	return func(o *CreateMailgunProviderOptions) {
		o.FromEmail = v
		o.enabledSetters["FromEmail"] = true
	}
}
func (srv *Messaging) WithCreateMailgunProviderReplyToName(v string) CreateMailgunProviderOption {
	return func(o *CreateMailgunProviderOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
func (srv *Messaging) WithCreateMailgunProviderReplyToEmail(v string) CreateMailgunProviderOption {
	return func(o *CreateMailgunProviderOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
func (srv *Messaging) WithCreateMailgunProviderEnabled(v bool) CreateMailgunProviderOption {
	return func(o *CreateMailgunProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
					
// CreateMailgunProvider create a new Mailgun provider.
func (srv *Messaging) CreateMailgunProvider(ProviderId string, Name string, optionalSetters ...CreateMailgunProviderOption)(*models.Provider, error) {
	path := "/messaging/providers/mailgun"
	options := CreateMailgunProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	params["name"] = Name
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["Domain"] {
		params["domain"] = options.Domain
	}
	if options.enabledSetters["IsEuRegion"] {
		params["isEuRegion"] = options.IsEuRegion
	}
	if options.enabledSetters["FromName"] {
		params["fromName"] = options.FromName
	}
	if options.enabledSetters["FromEmail"] {
		params["fromEmail"] = options.FromEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateMailgunProviderOptions struct {
	Name string
	ApiKey string
	Domain string
	IsEuRegion bool
	Enabled bool
	FromName string
	FromEmail string
	ReplyToName string
	ReplyToEmail string
	enabledSetters map[string]bool
}
func (options UpdateMailgunProviderOptions) New() *UpdateMailgunProviderOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"ApiKey": false,
		"Domain": false,
		"IsEuRegion": false,
		"Enabled": false,
		"FromName": false,
		"FromEmail": false,
		"ReplyToName": false,
		"ReplyToEmail": false,
	}
	return &options
}
type UpdateMailgunProviderOption func(*UpdateMailgunProviderOptions)
func (srv *Messaging) WithUpdateMailgunProviderName(v string) UpdateMailgunProviderOption {
	return func(o *UpdateMailgunProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithUpdateMailgunProviderApiKey(v string) UpdateMailgunProviderOption {
	return func(o *UpdateMailgunProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithUpdateMailgunProviderDomain(v string) UpdateMailgunProviderOption {
	return func(o *UpdateMailgunProviderOptions) {
		o.Domain = v
		o.enabledSetters["Domain"] = true
	}
}
func (srv *Messaging) WithUpdateMailgunProviderIsEuRegion(v bool) UpdateMailgunProviderOption {
	return func(o *UpdateMailgunProviderOptions) {
		o.IsEuRegion = v
		o.enabledSetters["IsEuRegion"] = true
	}
}
func (srv *Messaging) WithUpdateMailgunProviderEnabled(v bool) UpdateMailgunProviderOption {
	return func(o *UpdateMailgunProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithUpdateMailgunProviderFromName(v string) UpdateMailgunProviderOption {
	return func(o *UpdateMailgunProviderOptions) {
		o.FromName = v
		o.enabledSetters["FromName"] = true
	}
}
func (srv *Messaging) WithUpdateMailgunProviderFromEmail(v string) UpdateMailgunProviderOption {
	return func(o *UpdateMailgunProviderOptions) {
		o.FromEmail = v
		o.enabledSetters["FromEmail"] = true
	}
}
func (srv *Messaging) WithUpdateMailgunProviderReplyToName(v string) UpdateMailgunProviderOption {
	return func(o *UpdateMailgunProviderOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
func (srv *Messaging) WithUpdateMailgunProviderReplyToEmail(v string) UpdateMailgunProviderOption {
	return func(o *UpdateMailgunProviderOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
			
// UpdateMailgunProvider update a Mailgun provider by its unique ID.
func (srv *Messaging) UpdateMailgunProvider(ProviderId string, optionalSetters ...UpdateMailgunProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/mailgun/{providerId}")
	options := UpdateMailgunProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["Domain"] {
		params["domain"] = options.Domain
	}
	if options.enabledSetters["IsEuRegion"] {
		params["isEuRegion"] = options.IsEuRegion
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FromName"] {
		params["fromName"] = options.FromName
	}
	if options.enabledSetters["FromEmail"] {
		params["fromEmail"] = options.FromEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateMsg91ProviderOptions struct {
	TemplateId string
	SenderId string
	AuthKey string
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateMsg91ProviderOptions) New() *CreateMsg91ProviderOptions {
	options.enabledSetters = map[string]bool{
		"TemplateId": false,
		"SenderId": false,
		"AuthKey": false,
		"Enabled": false,
	}
	return &options
}
type CreateMsg91ProviderOption func(*CreateMsg91ProviderOptions)
func (srv *Messaging) WithCreateMsg91ProviderTemplateId(v string) CreateMsg91ProviderOption {
	return func(o *CreateMsg91ProviderOptions) {
		o.TemplateId = v
		o.enabledSetters["TemplateId"] = true
	}
}
func (srv *Messaging) WithCreateMsg91ProviderSenderId(v string) CreateMsg91ProviderOption {
	return func(o *CreateMsg91ProviderOptions) {
		o.SenderId = v
		o.enabledSetters["SenderId"] = true
	}
}
func (srv *Messaging) WithCreateMsg91ProviderAuthKey(v string) CreateMsg91ProviderOption {
	return func(o *CreateMsg91ProviderOptions) {
		o.AuthKey = v
		o.enabledSetters["AuthKey"] = true
	}
}
func (srv *Messaging) WithCreateMsg91ProviderEnabled(v bool) CreateMsg91ProviderOption {
	return func(o *CreateMsg91ProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
					
// CreateMsg91Provider create a new MSG91 provider.
func (srv *Messaging) CreateMsg91Provider(ProviderId string, Name string, optionalSetters ...CreateMsg91ProviderOption)(*models.Provider, error) {
	path := "/messaging/providers/msg91"
	options := CreateMsg91ProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	params["name"] = Name
	if options.enabledSetters["TemplateId"] {
		params["templateId"] = options.TemplateId
	}
	if options.enabledSetters["SenderId"] {
		params["senderId"] = options.SenderId
	}
	if options.enabledSetters["AuthKey"] {
		params["authKey"] = options.AuthKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateMsg91ProviderOptions struct {
	Name string
	Enabled bool
	TemplateId string
	SenderId string
	AuthKey string
	enabledSetters map[string]bool
}
func (options UpdateMsg91ProviderOptions) New() *UpdateMsg91ProviderOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Enabled": false,
		"TemplateId": false,
		"SenderId": false,
		"AuthKey": false,
	}
	return &options
}
type UpdateMsg91ProviderOption func(*UpdateMsg91ProviderOptions)
func (srv *Messaging) WithUpdateMsg91ProviderName(v string) UpdateMsg91ProviderOption {
	return func(o *UpdateMsg91ProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithUpdateMsg91ProviderEnabled(v bool) UpdateMsg91ProviderOption {
	return func(o *UpdateMsg91ProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithUpdateMsg91ProviderTemplateId(v string) UpdateMsg91ProviderOption {
	return func(o *UpdateMsg91ProviderOptions) {
		o.TemplateId = v
		o.enabledSetters["TemplateId"] = true
	}
}
func (srv *Messaging) WithUpdateMsg91ProviderSenderId(v string) UpdateMsg91ProviderOption {
	return func(o *UpdateMsg91ProviderOptions) {
		o.SenderId = v
		o.enabledSetters["SenderId"] = true
	}
}
func (srv *Messaging) WithUpdateMsg91ProviderAuthKey(v string) UpdateMsg91ProviderOption {
	return func(o *UpdateMsg91ProviderOptions) {
		o.AuthKey = v
		o.enabledSetters["AuthKey"] = true
	}
}
			
// UpdateMsg91Provider update a MSG91 provider by its unique ID.
func (srv *Messaging) UpdateMsg91Provider(ProviderId string, optionalSetters ...UpdateMsg91ProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/msg91/{providerId}")
	options := UpdateMsg91ProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["TemplateId"] {
		params["templateId"] = options.TemplateId
	}
	if options.enabledSetters["SenderId"] {
		params["senderId"] = options.SenderId
	}
	if options.enabledSetters["AuthKey"] {
		params["authKey"] = options.AuthKey
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateSendgridProviderOptions struct {
	ApiKey string
	FromName string
	FromEmail string
	ReplyToName string
	ReplyToEmail string
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateSendgridProviderOptions) New() *CreateSendgridProviderOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"FromName": false,
		"FromEmail": false,
		"ReplyToName": false,
		"ReplyToEmail": false,
		"Enabled": false,
	}
	return &options
}
type CreateSendgridProviderOption func(*CreateSendgridProviderOptions)
func (srv *Messaging) WithCreateSendgridProviderApiKey(v string) CreateSendgridProviderOption {
	return func(o *CreateSendgridProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithCreateSendgridProviderFromName(v string) CreateSendgridProviderOption {
	return func(o *CreateSendgridProviderOptions) {
		o.FromName = v
		o.enabledSetters["FromName"] = true
	}
}
func (srv *Messaging) WithCreateSendgridProviderFromEmail(v string) CreateSendgridProviderOption {
	return func(o *CreateSendgridProviderOptions) {
		o.FromEmail = v
		o.enabledSetters["FromEmail"] = true
	}
}
func (srv *Messaging) WithCreateSendgridProviderReplyToName(v string) CreateSendgridProviderOption {
	return func(o *CreateSendgridProviderOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
func (srv *Messaging) WithCreateSendgridProviderReplyToEmail(v string) CreateSendgridProviderOption {
	return func(o *CreateSendgridProviderOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
func (srv *Messaging) WithCreateSendgridProviderEnabled(v bool) CreateSendgridProviderOption {
	return func(o *CreateSendgridProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
					
// CreateSendgridProvider create a new Sendgrid provider.
func (srv *Messaging) CreateSendgridProvider(ProviderId string, Name string, optionalSetters ...CreateSendgridProviderOption)(*models.Provider, error) {
	path := "/messaging/providers/sendgrid"
	options := CreateSendgridProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	params["name"] = Name
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["FromName"] {
		params["fromName"] = options.FromName
	}
	if options.enabledSetters["FromEmail"] {
		params["fromEmail"] = options.FromEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateSendgridProviderOptions struct {
	Name string
	Enabled bool
	ApiKey string
	FromName string
	FromEmail string
	ReplyToName string
	ReplyToEmail string
	enabledSetters map[string]bool
}
func (options UpdateSendgridProviderOptions) New() *UpdateSendgridProviderOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Enabled": false,
		"ApiKey": false,
		"FromName": false,
		"FromEmail": false,
		"ReplyToName": false,
		"ReplyToEmail": false,
	}
	return &options
}
type UpdateSendgridProviderOption func(*UpdateSendgridProviderOptions)
func (srv *Messaging) WithUpdateSendgridProviderName(v string) UpdateSendgridProviderOption {
	return func(o *UpdateSendgridProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithUpdateSendgridProviderEnabled(v bool) UpdateSendgridProviderOption {
	return func(o *UpdateSendgridProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithUpdateSendgridProviderApiKey(v string) UpdateSendgridProviderOption {
	return func(o *UpdateSendgridProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithUpdateSendgridProviderFromName(v string) UpdateSendgridProviderOption {
	return func(o *UpdateSendgridProviderOptions) {
		o.FromName = v
		o.enabledSetters["FromName"] = true
	}
}
func (srv *Messaging) WithUpdateSendgridProviderFromEmail(v string) UpdateSendgridProviderOption {
	return func(o *UpdateSendgridProviderOptions) {
		o.FromEmail = v
		o.enabledSetters["FromEmail"] = true
	}
}
func (srv *Messaging) WithUpdateSendgridProviderReplyToName(v string) UpdateSendgridProviderOption {
	return func(o *UpdateSendgridProviderOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
func (srv *Messaging) WithUpdateSendgridProviderReplyToEmail(v string) UpdateSendgridProviderOption {
	return func(o *UpdateSendgridProviderOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
			
// UpdateSendgridProvider update a Sendgrid provider by its unique ID.
func (srv *Messaging) UpdateSendgridProvider(ProviderId string, optionalSetters ...UpdateSendgridProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/sendgrid/{providerId}")
	options := UpdateSendgridProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["FromName"] {
		params["fromName"] = options.FromName
	}
	if options.enabledSetters["FromEmail"] {
		params["fromEmail"] = options.FromEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateSmtpProviderOptions struct {
	Port int
	Username string
	Password string
	Encryption string
	AutoTLS bool
	Mailer string
	FromName string
	FromEmail string
	ReplyToName string
	ReplyToEmail string
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateSmtpProviderOptions) New() *CreateSmtpProviderOptions {
	options.enabledSetters = map[string]bool{
		"Port": false,
		"Username": false,
		"Password": false,
		"Encryption": false,
		"AutoTLS": false,
		"Mailer": false,
		"FromName": false,
		"FromEmail": false,
		"ReplyToName": false,
		"ReplyToEmail": false,
		"Enabled": false,
	}
	return &options
}
type CreateSmtpProviderOption func(*CreateSmtpProviderOptions)
func (srv *Messaging) WithCreateSmtpProviderPort(v int) CreateSmtpProviderOption {
	return func(o *CreateSmtpProviderOptions) {
		o.Port = v
		o.enabledSetters["Port"] = true
	}
}
func (srv *Messaging) WithCreateSmtpProviderUsername(v string) CreateSmtpProviderOption {
	return func(o *CreateSmtpProviderOptions) {
		o.Username = v
		o.enabledSetters["Username"] = true
	}
}
func (srv *Messaging) WithCreateSmtpProviderPassword(v string) CreateSmtpProviderOption {
	return func(o *CreateSmtpProviderOptions) {
		o.Password = v
		o.enabledSetters["Password"] = true
	}
}
func (srv *Messaging) WithCreateSmtpProviderEncryption(v string) CreateSmtpProviderOption {
	return func(o *CreateSmtpProviderOptions) {
		o.Encryption = v
		o.enabledSetters["Encryption"] = true
	}
}
func (srv *Messaging) WithCreateSmtpProviderAutoTLS(v bool) CreateSmtpProviderOption {
	return func(o *CreateSmtpProviderOptions) {
		o.AutoTLS = v
		o.enabledSetters["AutoTLS"] = true
	}
}
func (srv *Messaging) WithCreateSmtpProviderMailer(v string) CreateSmtpProviderOption {
	return func(o *CreateSmtpProviderOptions) {
		o.Mailer = v
		o.enabledSetters["Mailer"] = true
	}
}
func (srv *Messaging) WithCreateSmtpProviderFromName(v string) CreateSmtpProviderOption {
	return func(o *CreateSmtpProviderOptions) {
		o.FromName = v
		o.enabledSetters["FromName"] = true
	}
}
func (srv *Messaging) WithCreateSmtpProviderFromEmail(v string) CreateSmtpProviderOption {
	return func(o *CreateSmtpProviderOptions) {
		o.FromEmail = v
		o.enabledSetters["FromEmail"] = true
	}
}
func (srv *Messaging) WithCreateSmtpProviderReplyToName(v string) CreateSmtpProviderOption {
	return func(o *CreateSmtpProviderOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
func (srv *Messaging) WithCreateSmtpProviderReplyToEmail(v string) CreateSmtpProviderOption {
	return func(o *CreateSmtpProviderOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
func (srv *Messaging) WithCreateSmtpProviderEnabled(v bool) CreateSmtpProviderOption {
	return func(o *CreateSmtpProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
							
// CreateSmtpProvider create a new SMTP provider.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `Messaging.createSMTPProvider` instead.
func (srv *Messaging) CreateSmtpProvider(ProviderId string, Name string, Host string, optionalSetters ...CreateSmtpProviderOption)(*models.Provider, error) {
	path := "/messaging/providers/smtp"
	options := CreateSmtpProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	params["name"] = Name
	params["host"] = Host
	if options.enabledSetters["Port"] {
		params["port"] = options.Port
	}
	if options.enabledSetters["Username"] {
		params["username"] = options.Username
	}
	if options.enabledSetters["Password"] {
		params["password"] = options.Password
	}
	if options.enabledSetters["Encryption"] {
		params["encryption"] = options.Encryption
	}
	if options.enabledSetters["AutoTLS"] {
		params["autoTLS"] = options.AutoTLS
	}
	if options.enabledSetters["Mailer"] {
		params["mailer"] = options.Mailer
	}
	if options.enabledSetters["FromName"] {
		params["fromName"] = options.FromName
	}
	if options.enabledSetters["FromEmail"] {
		params["fromEmail"] = options.FromEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateSMTPProviderOptions struct {
	Port int
	Username string
	Password string
	Encryption string
	AutoTLS bool
	Mailer string
	FromName string
	FromEmail string
	ReplyToName string
	ReplyToEmail string
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateSMTPProviderOptions) New() *CreateSMTPProviderOptions {
	options.enabledSetters = map[string]bool{
		"Port": false,
		"Username": false,
		"Password": false,
		"Encryption": false,
		"AutoTLS": false,
		"Mailer": false,
		"FromName": false,
		"FromEmail": false,
		"ReplyToName": false,
		"ReplyToEmail": false,
		"Enabled": false,
	}
	return &options
}
type CreateSMTPProviderOption func(*CreateSMTPProviderOptions)
func (srv *Messaging) WithCreateSMTPProviderPort(v int) CreateSMTPProviderOption {
	return func(o *CreateSMTPProviderOptions) {
		o.Port = v
		o.enabledSetters["Port"] = true
	}
}
func (srv *Messaging) WithCreateSMTPProviderUsername(v string) CreateSMTPProviderOption {
	return func(o *CreateSMTPProviderOptions) {
		o.Username = v
		o.enabledSetters["Username"] = true
	}
}
func (srv *Messaging) WithCreateSMTPProviderPassword(v string) CreateSMTPProviderOption {
	return func(o *CreateSMTPProviderOptions) {
		o.Password = v
		o.enabledSetters["Password"] = true
	}
}
func (srv *Messaging) WithCreateSMTPProviderEncryption(v string) CreateSMTPProviderOption {
	return func(o *CreateSMTPProviderOptions) {
		o.Encryption = v
		o.enabledSetters["Encryption"] = true
	}
}
func (srv *Messaging) WithCreateSMTPProviderAutoTLS(v bool) CreateSMTPProviderOption {
	return func(o *CreateSMTPProviderOptions) {
		o.AutoTLS = v
		o.enabledSetters["AutoTLS"] = true
	}
}
func (srv *Messaging) WithCreateSMTPProviderMailer(v string) CreateSMTPProviderOption {
	return func(o *CreateSMTPProviderOptions) {
		o.Mailer = v
		o.enabledSetters["Mailer"] = true
	}
}
func (srv *Messaging) WithCreateSMTPProviderFromName(v string) CreateSMTPProviderOption {
	return func(o *CreateSMTPProviderOptions) {
		o.FromName = v
		o.enabledSetters["FromName"] = true
	}
}
func (srv *Messaging) WithCreateSMTPProviderFromEmail(v string) CreateSMTPProviderOption {
	return func(o *CreateSMTPProviderOptions) {
		o.FromEmail = v
		o.enabledSetters["FromEmail"] = true
	}
}
func (srv *Messaging) WithCreateSMTPProviderReplyToName(v string) CreateSMTPProviderOption {
	return func(o *CreateSMTPProviderOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
func (srv *Messaging) WithCreateSMTPProviderReplyToEmail(v string) CreateSMTPProviderOption {
	return func(o *CreateSMTPProviderOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
func (srv *Messaging) WithCreateSMTPProviderEnabled(v bool) CreateSMTPProviderOption {
	return func(o *CreateSMTPProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
							
// CreateSMTPProvider create a new SMTP provider.
func (srv *Messaging) CreateSMTPProvider(ProviderId string, Name string, Host string, optionalSetters ...CreateSMTPProviderOption)(*models.Provider, error) {
	path := "/messaging/providers/smtp"
	options := CreateSMTPProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	params["name"] = Name
	params["host"] = Host
	if options.enabledSetters["Port"] {
		params["port"] = options.Port
	}
	if options.enabledSetters["Username"] {
		params["username"] = options.Username
	}
	if options.enabledSetters["Password"] {
		params["password"] = options.Password
	}
	if options.enabledSetters["Encryption"] {
		params["encryption"] = options.Encryption
	}
	if options.enabledSetters["AutoTLS"] {
		params["autoTLS"] = options.AutoTLS
	}
	if options.enabledSetters["Mailer"] {
		params["mailer"] = options.Mailer
	}
	if options.enabledSetters["FromName"] {
		params["fromName"] = options.FromName
	}
	if options.enabledSetters["FromEmail"] {
		params["fromEmail"] = options.FromEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateSmtpProviderOptions struct {
	Name string
	Host string
	Port int
	Username string
	Password string
	Encryption string
	AutoTLS bool
	Mailer string
	FromName string
	FromEmail string
	ReplyToName string
	ReplyToEmail string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateSmtpProviderOptions) New() *UpdateSmtpProviderOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Host": false,
		"Port": false,
		"Username": false,
		"Password": false,
		"Encryption": false,
		"AutoTLS": false,
		"Mailer": false,
		"FromName": false,
		"FromEmail": false,
		"ReplyToName": false,
		"ReplyToEmail": false,
		"Enabled": false,
	}
	return &options
}
type UpdateSmtpProviderOption func(*UpdateSmtpProviderOptions)
func (srv *Messaging) WithUpdateSmtpProviderName(v string) UpdateSmtpProviderOption {
	return func(o *UpdateSmtpProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithUpdateSmtpProviderHost(v string) UpdateSmtpProviderOption {
	return func(o *UpdateSmtpProviderOptions) {
		o.Host = v
		o.enabledSetters["Host"] = true
	}
}
func (srv *Messaging) WithUpdateSmtpProviderPort(v int) UpdateSmtpProviderOption {
	return func(o *UpdateSmtpProviderOptions) {
		o.Port = v
		o.enabledSetters["Port"] = true
	}
}
func (srv *Messaging) WithUpdateSmtpProviderUsername(v string) UpdateSmtpProviderOption {
	return func(o *UpdateSmtpProviderOptions) {
		o.Username = v
		o.enabledSetters["Username"] = true
	}
}
func (srv *Messaging) WithUpdateSmtpProviderPassword(v string) UpdateSmtpProviderOption {
	return func(o *UpdateSmtpProviderOptions) {
		o.Password = v
		o.enabledSetters["Password"] = true
	}
}
func (srv *Messaging) WithUpdateSmtpProviderEncryption(v string) UpdateSmtpProviderOption {
	return func(o *UpdateSmtpProviderOptions) {
		o.Encryption = v
		o.enabledSetters["Encryption"] = true
	}
}
func (srv *Messaging) WithUpdateSmtpProviderAutoTLS(v bool) UpdateSmtpProviderOption {
	return func(o *UpdateSmtpProviderOptions) {
		o.AutoTLS = v
		o.enabledSetters["AutoTLS"] = true
	}
}
func (srv *Messaging) WithUpdateSmtpProviderMailer(v string) UpdateSmtpProviderOption {
	return func(o *UpdateSmtpProviderOptions) {
		o.Mailer = v
		o.enabledSetters["Mailer"] = true
	}
}
func (srv *Messaging) WithUpdateSmtpProviderFromName(v string) UpdateSmtpProviderOption {
	return func(o *UpdateSmtpProviderOptions) {
		o.FromName = v
		o.enabledSetters["FromName"] = true
	}
}
func (srv *Messaging) WithUpdateSmtpProviderFromEmail(v string) UpdateSmtpProviderOption {
	return func(o *UpdateSmtpProviderOptions) {
		o.FromEmail = v
		o.enabledSetters["FromEmail"] = true
	}
}
func (srv *Messaging) WithUpdateSmtpProviderReplyToName(v string) UpdateSmtpProviderOption {
	return func(o *UpdateSmtpProviderOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
func (srv *Messaging) WithUpdateSmtpProviderReplyToEmail(v string) UpdateSmtpProviderOption {
	return func(o *UpdateSmtpProviderOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
func (srv *Messaging) WithUpdateSmtpProviderEnabled(v bool) UpdateSmtpProviderOption {
	return func(o *UpdateSmtpProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
			
// UpdateSmtpProvider update a SMTP provider by its unique ID.
//
// Deprecated: This API has been deprecated since 1.8.0. Please use `Messaging.updateSMTPProvider` instead.
func (srv *Messaging) UpdateSmtpProvider(ProviderId string, optionalSetters ...UpdateSmtpProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/smtp/{providerId}")
	options := UpdateSmtpProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Host"] {
		params["host"] = options.Host
	}
	if options.enabledSetters["Port"] {
		params["port"] = options.Port
	}
	if options.enabledSetters["Username"] {
		params["username"] = options.Username
	}
	if options.enabledSetters["Password"] {
		params["password"] = options.Password
	}
	if options.enabledSetters["Encryption"] {
		params["encryption"] = options.Encryption
	}
	if options.enabledSetters["AutoTLS"] {
		params["autoTLS"] = options.AutoTLS
	}
	if options.enabledSetters["Mailer"] {
		params["mailer"] = options.Mailer
	}
	if options.enabledSetters["FromName"] {
		params["fromName"] = options.FromName
	}
	if options.enabledSetters["FromEmail"] {
		params["fromEmail"] = options.FromEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateSMTPProviderOptions struct {
	Name string
	Host string
	Port int
	Username string
	Password string
	Encryption string
	AutoTLS bool
	Mailer string
	FromName string
	FromEmail string
	ReplyToName string
	ReplyToEmail string
	Enabled bool
	enabledSetters map[string]bool
}
func (options UpdateSMTPProviderOptions) New() *UpdateSMTPProviderOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Host": false,
		"Port": false,
		"Username": false,
		"Password": false,
		"Encryption": false,
		"AutoTLS": false,
		"Mailer": false,
		"FromName": false,
		"FromEmail": false,
		"ReplyToName": false,
		"ReplyToEmail": false,
		"Enabled": false,
	}
	return &options
}
type UpdateSMTPProviderOption func(*UpdateSMTPProviderOptions)
func (srv *Messaging) WithUpdateSMTPProviderName(v string) UpdateSMTPProviderOption {
	return func(o *UpdateSMTPProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithUpdateSMTPProviderHost(v string) UpdateSMTPProviderOption {
	return func(o *UpdateSMTPProviderOptions) {
		o.Host = v
		o.enabledSetters["Host"] = true
	}
}
func (srv *Messaging) WithUpdateSMTPProviderPort(v int) UpdateSMTPProviderOption {
	return func(o *UpdateSMTPProviderOptions) {
		o.Port = v
		o.enabledSetters["Port"] = true
	}
}
func (srv *Messaging) WithUpdateSMTPProviderUsername(v string) UpdateSMTPProviderOption {
	return func(o *UpdateSMTPProviderOptions) {
		o.Username = v
		o.enabledSetters["Username"] = true
	}
}
func (srv *Messaging) WithUpdateSMTPProviderPassword(v string) UpdateSMTPProviderOption {
	return func(o *UpdateSMTPProviderOptions) {
		o.Password = v
		o.enabledSetters["Password"] = true
	}
}
func (srv *Messaging) WithUpdateSMTPProviderEncryption(v string) UpdateSMTPProviderOption {
	return func(o *UpdateSMTPProviderOptions) {
		o.Encryption = v
		o.enabledSetters["Encryption"] = true
	}
}
func (srv *Messaging) WithUpdateSMTPProviderAutoTLS(v bool) UpdateSMTPProviderOption {
	return func(o *UpdateSMTPProviderOptions) {
		o.AutoTLS = v
		o.enabledSetters["AutoTLS"] = true
	}
}
func (srv *Messaging) WithUpdateSMTPProviderMailer(v string) UpdateSMTPProviderOption {
	return func(o *UpdateSMTPProviderOptions) {
		o.Mailer = v
		o.enabledSetters["Mailer"] = true
	}
}
func (srv *Messaging) WithUpdateSMTPProviderFromName(v string) UpdateSMTPProviderOption {
	return func(o *UpdateSMTPProviderOptions) {
		o.FromName = v
		o.enabledSetters["FromName"] = true
	}
}
func (srv *Messaging) WithUpdateSMTPProviderFromEmail(v string) UpdateSMTPProviderOption {
	return func(o *UpdateSMTPProviderOptions) {
		o.FromEmail = v
		o.enabledSetters["FromEmail"] = true
	}
}
func (srv *Messaging) WithUpdateSMTPProviderReplyToName(v string) UpdateSMTPProviderOption {
	return func(o *UpdateSMTPProviderOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
func (srv *Messaging) WithUpdateSMTPProviderReplyToEmail(v string) UpdateSMTPProviderOption {
	return func(o *UpdateSMTPProviderOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
func (srv *Messaging) WithUpdateSMTPProviderEnabled(v bool) UpdateSMTPProviderOption {
	return func(o *UpdateSMTPProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
			
// UpdateSMTPProvider update a SMTP provider by its unique ID.
func (srv *Messaging) UpdateSMTPProvider(ProviderId string, optionalSetters ...UpdateSMTPProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/smtp/{providerId}")
	options := UpdateSMTPProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Host"] {
		params["host"] = options.Host
	}
	if options.enabledSetters["Port"] {
		params["port"] = options.Port
	}
	if options.enabledSetters["Username"] {
		params["username"] = options.Username
	}
	if options.enabledSetters["Password"] {
		params["password"] = options.Password
	}
	if options.enabledSetters["Encryption"] {
		params["encryption"] = options.Encryption
	}
	if options.enabledSetters["AutoTLS"] {
		params["autoTLS"] = options.AutoTLS
	}
	if options.enabledSetters["Mailer"] {
		params["mailer"] = options.Mailer
	}
	if options.enabledSetters["FromName"] {
		params["fromName"] = options.FromName
	}
	if options.enabledSetters["FromEmail"] {
		params["fromEmail"] = options.FromEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateTelesignProviderOptions struct {
	From string
	CustomerId string
	ApiKey string
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateTelesignProviderOptions) New() *CreateTelesignProviderOptions {
	options.enabledSetters = map[string]bool{
		"From": false,
		"CustomerId": false,
		"ApiKey": false,
		"Enabled": false,
	}
	return &options
}
type CreateTelesignProviderOption func(*CreateTelesignProviderOptions)
func (srv *Messaging) WithCreateTelesignProviderFrom(v string) CreateTelesignProviderOption {
	return func(o *CreateTelesignProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
func (srv *Messaging) WithCreateTelesignProviderCustomerId(v string) CreateTelesignProviderOption {
	return func(o *CreateTelesignProviderOptions) {
		o.CustomerId = v
		o.enabledSetters["CustomerId"] = true
	}
}
func (srv *Messaging) WithCreateTelesignProviderApiKey(v string) CreateTelesignProviderOption {
	return func(o *CreateTelesignProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithCreateTelesignProviderEnabled(v bool) CreateTelesignProviderOption {
	return func(o *CreateTelesignProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
					
// CreateTelesignProvider create a new Telesign provider.
func (srv *Messaging) CreateTelesignProvider(ProviderId string, Name string, optionalSetters ...CreateTelesignProviderOption)(*models.Provider, error) {
	path := "/messaging/providers/telesign"
	options := CreateTelesignProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	params["name"] = Name
	if options.enabledSetters["From"] {
		params["from"] = options.From
	}
	if options.enabledSetters["CustomerId"] {
		params["customerId"] = options.CustomerId
	}
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateTelesignProviderOptions struct {
	Name string
	Enabled bool
	CustomerId string
	ApiKey string
	From string
	enabledSetters map[string]bool
}
func (options UpdateTelesignProviderOptions) New() *UpdateTelesignProviderOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Enabled": false,
		"CustomerId": false,
		"ApiKey": false,
		"From": false,
	}
	return &options
}
type UpdateTelesignProviderOption func(*UpdateTelesignProviderOptions)
func (srv *Messaging) WithUpdateTelesignProviderName(v string) UpdateTelesignProviderOption {
	return func(o *UpdateTelesignProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithUpdateTelesignProviderEnabled(v bool) UpdateTelesignProviderOption {
	return func(o *UpdateTelesignProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithUpdateTelesignProviderCustomerId(v string) UpdateTelesignProviderOption {
	return func(o *UpdateTelesignProviderOptions) {
		o.CustomerId = v
		o.enabledSetters["CustomerId"] = true
	}
}
func (srv *Messaging) WithUpdateTelesignProviderApiKey(v string) UpdateTelesignProviderOption {
	return func(o *UpdateTelesignProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithUpdateTelesignProviderFrom(v string) UpdateTelesignProviderOption {
	return func(o *UpdateTelesignProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
			
// UpdateTelesignProvider update a Telesign provider by its unique ID.
func (srv *Messaging) UpdateTelesignProvider(ProviderId string, optionalSetters ...UpdateTelesignProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/telesign/{providerId}")
	options := UpdateTelesignProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["CustomerId"] {
		params["customerId"] = options.CustomerId
	}
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["From"] {
		params["from"] = options.From
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateTextmagicProviderOptions struct {
	From string
	Username string
	ApiKey string
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateTextmagicProviderOptions) New() *CreateTextmagicProviderOptions {
	options.enabledSetters = map[string]bool{
		"From": false,
		"Username": false,
		"ApiKey": false,
		"Enabled": false,
	}
	return &options
}
type CreateTextmagicProviderOption func(*CreateTextmagicProviderOptions)
func (srv *Messaging) WithCreateTextmagicProviderFrom(v string) CreateTextmagicProviderOption {
	return func(o *CreateTextmagicProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
func (srv *Messaging) WithCreateTextmagicProviderUsername(v string) CreateTextmagicProviderOption {
	return func(o *CreateTextmagicProviderOptions) {
		o.Username = v
		o.enabledSetters["Username"] = true
	}
}
func (srv *Messaging) WithCreateTextmagicProviderApiKey(v string) CreateTextmagicProviderOption {
	return func(o *CreateTextmagicProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithCreateTextmagicProviderEnabled(v bool) CreateTextmagicProviderOption {
	return func(o *CreateTextmagicProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
					
// CreateTextmagicProvider create a new Textmagic provider.
func (srv *Messaging) CreateTextmagicProvider(ProviderId string, Name string, optionalSetters ...CreateTextmagicProviderOption)(*models.Provider, error) {
	path := "/messaging/providers/textmagic"
	options := CreateTextmagicProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	params["name"] = Name
	if options.enabledSetters["From"] {
		params["from"] = options.From
	}
	if options.enabledSetters["Username"] {
		params["username"] = options.Username
	}
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateTextmagicProviderOptions struct {
	Name string
	Enabled bool
	Username string
	ApiKey string
	From string
	enabledSetters map[string]bool
}
func (options UpdateTextmagicProviderOptions) New() *UpdateTextmagicProviderOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Enabled": false,
		"Username": false,
		"ApiKey": false,
		"From": false,
	}
	return &options
}
type UpdateTextmagicProviderOption func(*UpdateTextmagicProviderOptions)
func (srv *Messaging) WithUpdateTextmagicProviderName(v string) UpdateTextmagicProviderOption {
	return func(o *UpdateTextmagicProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithUpdateTextmagicProviderEnabled(v bool) UpdateTextmagicProviderOption {
	return func(o *UpdateTextmagicProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithUpdateTextmagicProviderUsername(v string) UpdateTextmagicProviderOption {
	return func(o *UpdateTextmagicProviderOptions) {
		o.Username = v
		o.enabledSetters["Username"] = true
	}
}
func (srv *Messaging) WithUpdateTextmagicProviderApiKey(v string) UpdateTextmagicProviderOption {
	return func(o *UpdateTextmagicProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithUpdateTextmagicProviderFrom(v string) UpdateTextmagicProviderOption {
	return func(o *UpdateTextmagicProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
			
// UpdateTextmagicProvider update a Textmagic provider by its unique ID.
func (srv *Messaging) UpdateTextmagicProvider(ProviderId string, optionalSetters ...UpdateTextmagicProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/textmagic/{providerId}")
	options := UpdateTextmagicProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Username"] {
		params["username"] = options.Username
	}
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["From"] {
		params["from"] = options.From
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateTwilioProviderOptions struct {
	From string
	AccountSid string
	AuthToken string
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateTwilioProviderOptions) New() *CreateTwilioProviderOptions {
	options.enabledSetters = map[string]bool{
		"From": false,
		"AccountSid": false,
		"AuthToken": false,
		"Enabled": false,
	}
	return &options
}
type CreateTwilioProviderOption func(*CreateTwilioProviderOptions)
func (srv *Messaging) WithCreateTwilioProviderFrom(v string) CreateTwilioProviderOption {
	return func(o *CreateTwilioProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
func (srv *Messaging) WithCreateTwilioProviderAccountSid(v string) CreateTwilioProviderOption {
	return func(o *CreateTwilioProviderOptions) {
		o.AccountSid = v
		o.enabledSetters["AccountSid"] = true
	}
}
func (srv *Messaging) WithCreateTwilioProviderAuthToken(v string) CreateTwilioProviderOption {
	return func(o *CreateTwilioProviderOptions) {
		o.AuthToken = v
		o.enabledSetters["AuthToken"] = true
	}
}
func (srv *Messaging) WithCreateTwilioProviderEnabled(v bool) CreateTwilioProviderOption {
	return func(o *CreateTwilioProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
					
// CreateTwilioProvider create a new Twilio provider.
func (srv *Messaging) CreateTwilioProvider(ProviderId string, Name string, optionalSetters ...CreateTwilioProviderOption)(*models.Provider, error) {
	path := "/messaging/providers/twilio"
	options := CreateTwilioProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	params["name"] = Name
	if options.enabledSetters["From"] {
		params["from"] = options.From
	}
	if options.enabledSetters["AccountSid"] {
		params["accountSid"] = options.AccountSid
	}
	if options.enabledSetters["AuthToken"] {
		params["authToken"] = options.AuthToken
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateTwilioProviderOptions struct {
	Name string
	Enabled bool
	AccountSid string
	AuthToken string
	From string
	enabledSetters map[string]bool
}
func (options UpdateTwilioProviderOptions) New() *UpdateTwilioProviderOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Enabled": false,
		"AccountSid": false,
		"AuthToken": false,
		"From": false,
	}
	return &options
}
type UpdateTwilioProviderOption func(*UpdateTwilioProviderOptions)
func (srv *Messaging) WithUpdateTwilioProviderName(v string) UpdateTwilioProviderOption {
	return func(o *UpdateTwilioProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithUpdateTwilioProviderEnabled(v bool) UpdateTwilioProviderOption {
	return func(o *UpdateTwilioProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithUpdateTwilioProviderAccountSid(v string) UpdateTwilioProviderOption {
	return func(o *UpdateTwilioProviderOptions) {
		o.AccountSid = v
		o.enabledSetters["AccountSid"] = true
	}
}
func (srv *Messaging) WithUpdateTwilioProviderAuthToken(v string) UpdateTwilioProviderOption {
	return func(o *UpdateTwilioProviderOptions) {
		o.AuthToken = v
		o.enabledSetters["AuthToken"] = true
	}
}
func (srv *Messaging) WithUpdateTwilioProviderFrom(v string) UpdateTwilioProviderOption {
	return func(o *UpdateTwilioProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
			
// UpdateTwilioProvider update a Twilio provider by its unique ID.
func (srv *Messaging) UpdateTwilioProvider(ProviderId string, optionalSetters ...UpdateTwilioProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/twilio/{providerId}")
	options := UpdateTwilioProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["AccountSid"] {
		params["accountSid"] = options.AccountSid
	}
	if options.enabledSetters["AuthToken"] {
		params["authToken"] = options.AuthToken
	}
	if options.enabledSetters["From"] {
		params["from"] = options.From
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateVonageProviderOptions struct {
	From string
	ApiKey string
	ApiSecret string
	Enabled bool
	enabledSetters map[string]bool
}
func (options CreateVonageProviderOptions) New() *CreateVonageProviderOptions {
	options.enabledSetters = map[string]bool{
		"From": false,
		"ApiKey": false,
		"ApiSecret": false,
		"Enabled": false,
	}
	return &options
}
type CreateVonageProviderOption func(*CreateVonageProviderOptions)
func (srv *Messaging) WithCreateVonageProviderFrom(v string) CreateVonageProviderOption {
	return func(o *CreateVonageProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
func (srv *Messaging) WithCreateVonageProviderApiKey(v string) CreateVonageProviderOption {
	return func(o *CreateVonageProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithCreateVonageProviderApiSecret(v string) CreateVonageProviderOption {
	return func(o *CreateVonageProviderOptions) {
		o.ApiSecret = v
		o.enabledSetters["ApiSecret"] = true
	}
}
func (srv *Messaging) WithCreateVonageProviderEnabled(v bool) CreateVonageProviderOption {
	return func(o *CreateVonageProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
					
// CreateVonageProvider create a new Vonage provider.
func (srv *Messaging) CreateVonageProvider(ProviderId string, Name string, optionalSetters ...CreateVonageProviderOption)(*models.Provider, error) {
	path := "/messaging/providers/vonage"
	options := CreateVonageProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	params["name"] = Name
	if options.enabledSetters["From"] {
		params["from"] = options.From
	}
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["ApiSecret"] {
		params["apiSecret"] = options.ApiSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateVonageProviderOptions struct {
	Name string
	Enabled bool
	ApiKey string
	ApiSecret string
	From string
	enabledSetters map[string]bool
}
func (options UpdateVonageProviderOptions) New() *UpdateVonageProviderOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Enabled": false,
		"ApiKey": false,
		"ApiSecret": false,
		"From": false,
	}
	return &options
}
type UpdateVonageProviderOption func(*UpdateVonageProviderOptions)
func (srv *Messaging) WithUpdateVonageProviderName(v string) UpdateVonageProviderOption {
	return func(o *UpdateVonageProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithUpdateVonageProviderEnabled(v bool) UpdateVonageProviderOption {
	return func(o *UpdateVonageProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithUpdateVonageProviderApiKey(v string) UpdateVonageProviderOption {
	return func(o *UpdateVonageProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithUpdateVonageProviderApiSecret(v string) UpdateVonageProviderOption {
	return func(o *UpdateVonageProviderOptions) {
		o.ApiSecret = v
		o.enabledSetters["ApiSecret"] = true
	}
}
func (srv *Messaging) WithUpdateVonageProviderFrom(v string) UpdateVonageProviderOption {
	return func(o *UpdateVonageProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
			
// UpdateVonageProvider update a Vonage provider by its unique ID.
func (srv *Messaging) UpdateVonageProvider(ProviderId string, optionalSetters ...UpdateVonageProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/vonage/{providerId}")
	options := UpdateVonageProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["ApiSecret"] {
		params["apiSecret"] = options.ApiSecret
	}
	if options.enabledSetters["From"] {
		params["from"] = options.From
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetProvider get a provider by its unique ID.
func (srv *Messaging) GetProvider(ProviderId string)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/{providerId}")
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteProvider delete a provider by its unique ID.
func (srv *Messaging) DeleteProvider(ProviderId string)(*interface{}, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/{providerId}")
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
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
type ListProviderLogsOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListProviderLogsOptions) New() *ListProviderLogsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListProviderLogsOption func(*ListProviderLogsOptions)
func (srv *Messaging) WithListProviderLogsQueries(v []string) ListProviderLogsOption {
	return func(o *ListProviderLogsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
			
// ListProviderLogs get the provider activity logs listed by its unique ID.
func (srv *Messaging) ListProviderLogs(ProviderId string, optionalSetters ...ListProviderLogsOption)(*models.LogList, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/messaging/providers/{providerId}/logs")
	options := ListProviderLogsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.LogList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LogList
	parsed, ok := resp.Result.(models.LogList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListSubscriberLogsOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListSubscriberLogsOptions) New() *ListSubscriberLogsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListSubscriberLogsOption func(*ListSubscriberLogsOptions)
func (srv *Messaging) WithListSubscriberLogsQueries(v []string) ListSubscriberLogsOption {
	return func(o *ListSubscriberLogsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
			
// ListSubscriberLogs get the subscriber activity logs listed by its unique
// ID.
func (srv *Messaging) ListSubscriberLogs(SubscriberId string, optionalSetters ...ListSubscriberLogsOption)(*models.LogList, error) {
	r := strings.NewReplacer("{subscriberId}", SubscriberId)
	path := r.Replace("/messaging/subscribers/{subscriberId}/logs")
	options := ListSubscriberLogsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["subscriberId"] = SubscriberId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.LogList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LogList
	parsed, ok := resp.Result.(models.LogList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListTopicsOptions struct {
	Queries []string
	Search string
	enabledSetters map[string]bool
}
func (options ListTopicsOptions) New() *ListTopicsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
	}
	return &options
}
type ListTopicsOption func(*ListTopicsOptions)
func (srv *Messaging) WithListTopicsQueries(v []string) ListTopicsOption {
	return func(o *ListTopicsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Messaging) WithListTopicsSearch(v string) ListTopicsOption {
	return func(o *ListTopicsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
	
// ListTopics get a list of all topics from the current Appwrite project.
func (srv *Messaging) ListTopics(optionalSetters ...ListTopicsOption)(*models.TopicList, error) {
	path := "/messaging/topics"
	options := ListTopicsOptions{}.New()
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
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.TopicList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.TopicList
	parsed, ok := resp.Result.(models.TopicList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CreateTopicOptions struct {
	Subscribe []string
	enabledSetters map[string]bool
}
func (options CreateTopicOptions) New() *CreateTopicOptions {
	options.enabledSetters = map[string]bool{
		"Subscribe": false,
	}
	return &options
}
type CreateTopicOption func(*CreateTopicOptions)
func (srv *Messaging) WithCreateTopicSubscribe(v []string) CreateTopicOption {
	return func(o *CreateTopicOptions) {
		o.Subscribe = v
		o.enabledSetters["Subscribe"] = true
	}
}
					
// CreateTopic create a new topic.
func (srv *Messaging) CreateTopic(TopicId string, Name string, optionalSetters ...CreateTopicOption)(*models.Topic, error) {
	path := "/messaging/topics"
	options := CreateTopicOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	params["name"] = Name
	if options.enabledSetters["Subscribe"] {
		params["subscribe"] = options.Subscribe
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

		parsed := models.Topic{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Topic
	parsed, ok := resp.Result.(models.Topic)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetTopic get a topic by its unique ID.
func (srv *Messaging) GetTopic(TopicId string)(*models.Topic, error) {
	r := strings.NewReplacer("{topicId}", TopicId)
	path := r.Replace("/messaging/topics/{topicId}")
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Topic{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Topic
	parsed, ok := resp.Result.(models.Topic)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type UpdateTopicOptions struct {
	Name string
	Subscribe []string
	enabledSetters map[string]bool
}
func (options UpdateTopicOptions) New() *UpdateTopicOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Subscribe": false,
	}
	return &options
}
type UpdateTopicOption func(*UpdateTopicOptions)
func (srv *Messaging) WithUpdateTopicName(v string) UpdateTopicOption {
	return func(o *UpdateTopicOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithUpdateTopicSubscribe(v []string) UpdateTopicOption {
	return func(o *UpdateTopicOptions) {
		o.Subscribe = v
		o.enabledSetters["Subscribe"] = true
	}
}
			
// UpdateTopic update a topic by its unique ID.
func (srv *Messaging) UpdateTopic(TopicId string, optionalSetters ...UpdateTopicOption)(*models.Topic, error) {
	r := strings.NewReplacer("{topicId}", TopicId)
	path := r.Replace("/messaging/topics/{topicId}")
	options := UpdateTopicOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Subscribe"] {
		params["subscribe"] = options.Subscribe
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Topic{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Topic
	parsed, ok := resp.Result.(models.Topic)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteTopic delete a topic by its unique ID.
func (srv *Messaging) DeleteTopic(TopicId string)(*interface{}, error) {
	r := strings.NewReplacer("{topicId}", TopicId)
	path := r.Replace("/messaging/topics/{topicId}")
	params := map[string]interface{}{}
	params["topicId"] = TopicId
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
type ListTopicLogsOptions struct {
	Queries []string
	enabledSetters map[string]bool
}
func (options ListTopicLogsOptions) New() *ListTopicLogsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
	}
	return &options
}
type ListTopicLogsOption func(*ListTopicLogsOptions)
func (srv *Messaging) WithListTopicLogsQueries(v []string) ListTopicLogsOption {
	return func(o *ListTopicLogsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
			
// ListTopicLogs get the topic activity logs listed by its unique ID.
func (srv *Messaging) ListTopicLogs(TopicId string, optionalSetters ...ListTopicLogsOption)(*models.LogList, error) {
	r := strings.NewReplacer("{topicId}", TopicId)
	path := r.Replace("/messaging/topics/{topicId}/logs")
	options := ListTopicLogsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.LogList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LogList
	parsed, ok := resp.Result.(models.LogList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type ListSubscribersOptions struct {
	Queries []string
	Search string
	enabledSetters map[string]bool
}
func (options ListSubscribersOptions) New() *ListSubscribersOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
	}
	return &options
}
type ListSubscribersOption func(*ListSubscribersOptions)
func (srv *Messaging) WithListSubscribersQueries(v []string) ListSubscribersOption {
	return func(o *ListSubscribersOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Messaging) WithListSubscribersSearch(v string) ListSubscribersOption {
	return func(o *ListSubscribersOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
			
// ListSubscribers get a list of all subscribers from the current Appwrite
// project.
func (srv *Messaging) ListSubscribers(TopicId string, optionalSetters ...ListSubscribersOption)(*models.SubscriberList, error) {
	r := strings.NewReplacer("{topicId}", TopicId)
	path := r.Replace("/messaging/topics/{topicId}/subscribers")
	options := ListSubscribersOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.SubscriberList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.SubscriberList
	parsed, ok := resp.Result.(models.SubscriberList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// CreateSubscriber create a new subscriber.
func (srv *Messaging) CreateSubscriber(TopicId string, SubscriberId string, TargetId string)(*models.Subscriber, error) {
	r := strings.NewReplacer("{topicId}", TopicId)
	path := r.Replace("/messaging/topics/{topicId}/subscribers")
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	params["subscriberId"] = SubscriberId
	params["targetId"] = TargetId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Subscriber{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Subscriber
	parsed, ok := resp.Result.(models.Subscriber)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetSubscriber get a subscriber by its unique ID.
func (srv *Messaging) GetSubscriber(TopicId string, SubscriberId string)(*models.Subscriber, error) {
	r := strings.NewReplacer("{topicId}", TopicId, "{subscriberId}", SubscriberId)
	path := r.Replace("/messaging/topics/{topicId}/subscribers/{subscriberId}")
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	params["subscriberId"] = SubscriberId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Subscriber{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Subscriber
	parsed, ok := resp.Result.(models.Subscriber)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// DeleteSubscriber delete a subscriber by its unique ID.
func (srv *Messaging) DeleteSubscriber(TopicId string, SubscriberId string)(*interface{}, error) {
	r := strings.NewReplacer("{topicId}", TopicId, "{subscriberId}", SubscriberId)
	path := r.Replace("/messaging/topics/{topicId}/subscribers/{subscriberId}")
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	params["subscriberId"] = SubscriberId
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
