package models

import (
	"encoding/json"
	"errors"
)

// Session Model
type Session struct {
	// Session ID.
	Id string `json:"$id"`
	// Session creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Session update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// User ID.
	UserId string `json:"userId"`
	// Session expiration date in ISO 8601 format.
	Expire string `json:"expire"`
	// Session Provider.
	Provider string `json:"provider"`
	// Session Provider User ID.
	ProviderUid string `json:"providerUid"`
	// Session Provider Access Token.
	ProviderAccessToken string `json:"providerAccessToken"`
	// The date of when the access token expires in ISO 8601 format.
	ProviderAccessTokenExpiry string `json:"providerAccessTokenExpiry"`
	// Session Provider Refresh Token.
	ProviderRefreshToken string `json:"providerRefreshToken"`
	// IP in use when the session was created.
	Ip string `json:"ip"`
	// Operating system code name. View list of [available
	// options](https://github.com/appwrite/appwrite/blob/master/docs/lists/os.json).
	OsCode string `json:"osCode"`
	// Operating system name.
	OsName string `json:"osName"`
	// Operating system version.
	OsVersion string `json:"osVersion"`
	// Client type.
	ClientType string `json:"clientType"`
	// Client code name. View list of [available
	// options](https://github.com/appwrite/appwrite/blob/master/docs/lists/clients.json).
	ClientCode string `json:"clientCode"`
	// Client name.
	ClientName string `json:"clientName"`
	// Client version.
	ClientVersion string `json:"clientVersion"`
	// Client engine name.
	ClientEngine string `json:"clientEngine"`
	// Client engine name.
	ClientEngineVersion string `json:"clientEngineVersion"`
	// Device name.
	DeviceName string `json:"deviceName"`
	// Device brand name.
	DeviceBrand string `json:"deviceBrand"`
	// Device model name.
	DeviceModel string `json:"deviceModel"`
	// Country two-character ISO 3166-1 alpha code.
	CountryCode string `json:"countryCode"`
	// Country name.
	CountryName string `json:"countryName"`
	// Returns true if this the current user session.
	Current bool `json:"current"`
	// Returns a list of active session factors.
	Factors []string `json:"factors"`
	// Secret used to authenticate the user. Only included if the request was made
	// with an API key
	Secret string `json:"secret"`
	// Most recent date in ISO 8601 format when the session successfully passed
	// MFA challenge.
	MfaUpdatedAt string `json:"mfaUpdatedAt"`

	// Used by Decode() method
	data []byte
}

func (model Session) New(data []byte) *Session {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Session) UnmarshalJSON(data []byte) error {
	type alias Session
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                        json.RawMessage `json:"$id"`
		CreatedAt                 json.RawMessage `json:"$createdAt"`
		UpdatedAt                 json.RawMessage `json:"$updatedAt"`
		UserId                    json.RawMessage `json:"userId"`
		Expire                    json.RawMessage `json:"expire"`
		Provider                  json.RawMessage `json:"provider"`
		ProviderUid               json.RawMessage `json:"providerUid"`
		ProviderAccessToken       json.RawMessage `json:"providerAccessToken"`
		ProviderAccessTokenExpiry json.RawMessage `json:"providerAccessTokenExpiry"`
		ProviderRefreshToken      json.RawMessage `json:"providerRefreshToken"`
		Ip                        json.RawMessage `json:"ip"`
		OsCode                    json.RawMessage `json:"osCode"`
		OsName                    json.RawMessage `json:"osName"`
		OsVersion                 json.RawMessage `json:"osVersion"`
		ClientType                json.RawMessage `json:"clientType"`
		ClientCode                json.RawMessage `json:"clientCode"`
		ClientName                json.RawMessage `json:"clientName"`
		ClientVersion             json.RawMessage `json:"clientVersion"`
		ClientEngine              json.RawMessage `json:"clientEngine"`
		ClientEngineVersion       json.RawMessage `json:"clientEngineVersion"`
		DeviceName                json.RawMessage `json:"deviceName"`
		DeviceBrand               json.RawMessage `json:"deviceBrand"`
		DeviceModel               json.RawMessage `json:"deviceModel"`
		CountryCode               json.RawMessage `json:"countryCode"`
		CountryName               json.RawMessage `json:"countryName"`
		Current                   json.RawMessage `json:"current"`
		Factors                   json.RawMessage `json:"factors"`
		Secret                    json.RawMessage `json:"secret"`
		MfaUpdatedAt              json.RawMessage `json:"mfaUpdatedAt"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.UpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UpdatedAt); err != nil {
			return err
		}
	}
	if present.UserId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserId); err != nil {
			return err
		}
	}
	if present.Expire == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Expire); err != nil {
			return err
		}
	}
	if present.Provider == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Provider); err != nil {
			return err
		}
	}
	if present.ProviderUid == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderUid); err != nil {
			return err
		}
	}
	if present.ProviderAccessToken == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderAccessToken); err != nil {
			return err
		}
	}
	if present.ProviderAccessTokenExpiry == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderAccessTokenExpiry); err != nil {
			return err
		}
	}
	if present.ProviderRefreshToken == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderRefreshToken); err != nil {
			return err
		}
	}
	if present.Ip == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Ip); err != nil {
			return err
		}
	}
	if present.OsCode == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.OsCode); err != nil {
			return err
		}
	}
	if present.OsName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.OsName); err != nil {
			return err
		}
	}
	if present.OsVersion == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.OsVersion); err != nil {
			return err
		}
	}
	if present.ClientType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ClientType); err != nil {
			return err
		}
	}
	if present.ClientCode == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ClientCode); err != nil {
			return err
		}
	}
	if present.ClientName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ClientName); err != nil {
			return err
		}
	}
	if present.ClientVersion == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ClientVersion); err != nil {
			return err
		}
	}
	if present.ClientEngine == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ClientEngine); err != nil {
			return err
		}
	}
	if present.ClientEngineVersion == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ClientEngineVersion); err != nil {
			return err
		}
	}
	if present.DeviceName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DeviceName); err != nil {
			return err
		}
	}
	if present.DeviceBrand == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DeviceBrand); err != nil {
			return err
		}
	}
	if present.DeviceModel == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DeviceModel); err != nil {
			return err
		}
	}
	if present.CountryCode == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CountryCode); err != nil {
			return err
		}
	}
	if present.CountryName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CountryName); err != nil {
			return err
		}
	}
	if present.Current == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Current); err != nil {
			return err
		}
	}
	if present.Factors == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Factors); err != nil {
			return err
		}
	}
	if present.Secret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Secret); err != nil {
			return err
		}
	}
	if present.MfaUpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.MfaUpdatedAt); err != nil {
			return err
		}
	}

	*model = Session(decoded)
	return nil
}

func (model *Session) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
