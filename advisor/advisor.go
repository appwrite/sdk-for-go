package advisor

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/v4/client"
	"github.com/appwrite/sdk-for-go/v4/models"
	"strings"
)

// Advisor service
type Advisor struct {
	client client.Client
}

func New(clt client.Client) *Advisor {
	return &Advisor{
		client: clt,
	}
}

type ListReportsOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListReportsOptions) New() *ListReportsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListReportsOption func(*ListReportsOptions)
func (srv *Advisor) WithListReportsQueries(v []string) ListReportsOption {
	return func(o *ListReportsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Advisor) WithListReportsTotal(v bool) ListReportsOption {
	return func(o *ListReportsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// ListReports get a list of all the project's analyzer reports. You can use
// the query params to filter your results.
func (srv *Advisor) ListReports(optionalSetters ...ListReportsOption)(*models.ReportList, error) {
	path := "/reports"
	options := ListReportsOptions{}.New()
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

		parsed := models.ReportList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ReportList
	parsed, ok := resp.Result.(models.ReportList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// GetReport get an analyzer report by its unique ID. The response includes
// the report's metadata and the nested insights it produced.
func (srv *Advisor) GetReport(ReportId string)(*models.Report, error) {
	r := strings.NewReplacer("{reportId}", ReportId)
	path := r.Replace("/reports/{reportId}")
	params := map[string]interface{}{}
	params["reportId"] = ReportId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Report{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Report
	parsed, ok := resp.Result.(models.Report)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// DeleteReport delete an analyzer report by its unique ID. Nested insights
// and CTA metadata are removed asynchronously by the deletes worker.
func (srv *Advisor) DeleteReport(ReportId string)(*interface{}, error) {
	r := strings.NewReplacer("{reportId}", ReportId)
	path := r.Replace("/reports/{reportId}")
	params := map[string]interface{}{}
	params["reportId"] = ReportId
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
type ListInsightsOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options ListInsightsOptions) New() *ListInsightsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type ListInsightsOption func(*ListInsightsOptions)
func (srv *Advisor) WithListInsightsQueries(v []string) ListInsightsOption {
	return func(o *ListInsightsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Advisor) WithListInsightsTotal(v bool) ListInsightsOption {
	return func(o *ListInsightsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// ListInsights list the insights produced under a single analyzer report. You
// can use the query params to filter your results further.
func (srv *Advisor) ListInsights(ReportId string, optionalSetters ...ListInsightsOption)(*models.InsightList, error) {
	r := strings.NewReplacer("{reportId}", ReportId)
	path := r.Replace("/reports/{reportId}/insights")
	options := ListInsightsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["reportId"] = ReportId
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

		parsed := models.InsightList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.InsightList
	parsed, ok := resp.Result.(models.InsightList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// GetInsight get an insight by its unique ID, scoped to its parent report.
func (srv *Advisor) GetInsight(ReportId string, InsightId string)(*models.Insight, error) {
	r := strings.NewReplacer("{reportId}", ReportId, "{insightId}", InsightId)
	path := r.Replace("/reports/{reportId}/insights/{insightId}")
	params := map[string]interface{}{}
	params["reportId"] = ReportId
	params["insightId"] = InsightId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Insight{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Insight
	parsed, ok := resp.Result.(models.Insight)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
