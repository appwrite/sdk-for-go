package locale

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/appwrite/sdk-for-go/v7/client"
	"github.com/appwrite/sdk-for-go/v7/models"
)

// Locale service
type Locale struct {
	client client.Client
}

func New(clt client.Client) *Locale {
	return &Locale{
		client: clt,
	}
}

// Get get the current user location based on IP. Returns an object with user
// country code, country name, continent name, continent code, ip address and
// suggested currency. You can use the locale header to get the data in a
// supported language.
//
// ([IP Geolocation by DB-IP](https://db-ip.com))
func (srv *Locale) Get() (*models.Locale, error) {
	path := "/locale"
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

		parsed := models.Locale{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Locale
	parsed, ok := resp.Result.(models.Locale)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListCodesOptions struct {
	Total          bool
	enabledSetters map[string]bool
}

func (options ListCodesOptions) New() *ListCodesOptions {
	options.enabledSetters = map[string]bool{"Total": false}
	return &options
}

type ListCodesOption func(*ListCodesOptions)

func (srv *Locale) WithListCodesTotal(v bool) ListCodesOption {
	return func(o *ListCodesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListCodes list of all locale codes in [ISO
// 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).
func (srv *Locale) ListCodes(optionalSetters ...ListCodesOption) (*models.LocaleCodeList, error) {
	path := "/locale/codes"
	options := ListCodesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
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

		parsed := models.LocaleCodeList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LocaleCodeList
	parsed, ok := resp.Result.(models.LocaleCodeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListContinentsOptions struct {
	Total          bool
	enabledSetters map[string]bool
}

func (options ListContinentsOptions) New() *ListContinentsOptions {
	options.enabledSetters = map[string]bool{"Total": false}
	return &options
}

type ListContinentsOption func(*ListContinentsOptions)

func (srv *Locale) WithListContinentsTotal(v bool) ListContinentsOption {
	return func(o *ListContinentsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListContinents list of all continents. You can use the locale header to get
// the data in a supported language.
func (srv *Locale) ListContinents(optionalSetters ...ListContinentsOption) (*models.ContinentList, error) {
	path := "/locale/continents"
	options := ListContinentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
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

		parsed := models.ContinentList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ContinentList
	parsed, ok := resp.Result.(models.ContinentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListCountriesOptions struct {
	Total          bool
	enabledSetters map[string]bool
}

func (options ListCountriesOptions) New() *ListCountriesOptions {
	options.enabledSetters = map[string]bool{"Total": false}
	return &options
}

type ListCountriesOption func(*ListCountriesOptions)

func (srv *Locale) WithListCountriesTotal(v bool) ListCountriesOption {
	return func(o *ListCountriesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListCountries list of all countries. You can use the locale header to get
// the data in a supported language.
func (srv *Locale) ListCountries(optionalSetters ...ListCountriesOption) (*models.CountryList, error) {
	path := "/locale/countries"
	options := ListCountriesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
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

		parsed := models.CountryList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.CountryList
	parsed, ok := resp.Result.(models.CountryList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListCountriesEUOptions struct {
	Total          bool
	enabledSetters map[string]bool
}

func (options ListCountriesEUOptions) New() *ListCountriesEUOptions {
	options.enabledSetters = map[string]bool{"Total": false}
	return &options
}

type ListCountriesEUOption func(*ListCountriesEUOptions)

func (srv *Locale) WithListCountriesEUTotal(v bool) ListCountriesEUOption {
	return func(o *ListCountriesEUOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListCountriesEU list of all countries that are currently members of the EU.
// You can use the locale header to get the data in a supported language.
func (srv *Locale) ListCountriesEU(optionalSetters ...ListCountriesEUOption) (*models.CountryList, error) {
	path := "/locale/countries/eu"
	options := ListCountriesEUOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
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

		parsed := models.CountryList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.CountryList
	parsed, ok := resp.Result.(models.CountryList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListCountriesPhonesOptions struct {
	Total          bool
	enabledSetters map[string]bool
}

func (options ListCountriesPhonesOptions) New() *ListCountriesPhonesOptions {
	options.enabledSetters = map[string]bool{"Total": false}
	return &options
}

type ListCountriesPhonesOption func(*ListCountriesPhonesOptions)

func (srv *Locale) WithListCountriesPhonesTotal(v bool) ListCountriesPhonesOption {
	return func(o *ListCountriesPhonesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListCountriesPhones list of all countries phone codes. You can use the
// locale header to get the data in a supported language.
func (srv *Locale) ListCountriesPhones(optionalSetters ...ListCountriesPhonesOption) (*models.PhoneList, error) {
	path := "/locale/countries/phones"
	options := ListCountriesPhonesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
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

		parsed := models.PhoneList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PhoneList
	parsed, ok := resp.Result.(models.PhoneList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListCurrenciesOptions struct {
	Total          bool
	enabledSetters map[string]bool
}

func (options ListCurrenciesOptions) New() *ListCurrenciesOptions {
	options.enabledSetters = map[string]bool{"Total": false}
	return &options
}

type ListCurrenciesOption func(*ListCurrenciesOptions)

func (srv *Locale) WithListCurrenciesTotal(v bool) ListCurrenciesOption {
	return func(o *ListCurrenciesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListCurrencies list of all currencies, including currency symbol, name,
// plural, and decimal digits for all major and minor currencies. You can use
// the locale header to get the data in a supported language.
func (srv *Locale) ListCurrencies(optionalSetters ...ListCurrenciesOption) (*models.CurrencyList, error) {
	path := "/locale/currencies"
	options := ListCurrenciesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
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

		parsed := models.CurrencyList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.CurrencyList
	parsed, ok := resp.Result.(models.CurrencyList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

type ListLanguagesOptions struct {
	Total          bool
	enabledSetters map[string]bool
}

func (options ListLanguagesOptions) New() *ListLanguagesOptions {
	options.enabledSetters = map[string]bool{"Total": false}
	return &options
}

type ListLanguagesOption func(*ListLanguagesOptions)

func (srv *Locale) WithListLanguagesTotal(v bool) ListLanguagesOption {
	return func(o *ListLanguagesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}

// ListLanguages list of all languages classified by ISO 639-1 including
// 2-letter code, name in English, and name in the respective language.
func (srv *Locale) ListLanguages(optionalSetters ...ListLanguagesOption) (*models.LanguageList, error) {
	path := "/locale/languages"
	options := ListLanguagesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
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

		parsed := models.LanguageList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LanguageList
	parsed, ok := resp.Result.(models.LanguageList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
