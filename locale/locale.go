package locale

import (
	"encoding/json"
	"errors"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"strings"
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
func (srv *Locale) Get()(*models.Locale, error) {
	path := "/locale"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.Locale{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
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

// ListCodes list of all locale codes in [ISO
// 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).
func (srv *Locale) ListCodes()(*models.LocaleCodeList, error) {
	path := "/locale/codes"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.LocaleCodeList{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
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

// ListContinents list of all continents. You can use the locale header to get
// the data in a supported language.
func (srv *Locale) ListContinents()(*models.ContinentList, error) {
	path := "/locale/continents"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.ContinentList{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
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

// ListCountries list of all countries. You can use the locale header to get
// the data in a supported language.
func (srv *Locale) ListCountries()(*models.CountryList, error) {
	path := "/locale/countries"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.CountryList{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
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

// ListCountriesEU list of all countries that are currently members of the EU.
// You can use the locale header to get the data in a supported language.
func (srv *Locale) ListCountriesEU()(*models.CountryList, error) {
	path := "/locale/countries/eu"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.CountryList{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
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

// ListCountriesPhones list of all countries phone codes. You can use the
// locale header to get the data in a supported language.
func (srv *Locale) ListCountriesPhones()(*models.PhoneList, error) {
	path := "/locale/countries/phones"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.PhoneList{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
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

// ListCurrencies list of all currencies, including currency symbol, name,
// plural, and decimal digits for all major and minor currencies. You can use
// the locale header to get the data in a supported language.
func (srv *Locale) ListCurrencies()(*models.CurrencyList, error) {
	path := "/locale/currencies"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.CurrencyList{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
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

// ListLanguages list of all languages classified by ISO 639-1 including
// 2-letter code, name in English, and name in the respective language.
func (srv *Locale) ListLanguages()(*models.LanguageList, error) {
	path := "/locale/languages"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytesData := []byte(resp.Result.(string))

		parsed := models.LanguageList{}.New(bytesData)

		err = json.Unmarshal(bytesData, parsed)
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
