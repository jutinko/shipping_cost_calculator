package utilities

import (
	"encoding/json"
	"strconv"
)

type RedisCredentials struct {
	Redis []struct {
		Credentials struct {
			Host     string
			Password string
			Port     int
		}
	} `json:"p-redis"`
}

type EnvReader struct {
	EnvVars RedisCredentials
}

func NewEnvReader(jsonData []byte) *EnvReader {
	var mapping RedisCredentials
	json.Unmarshal(jsonData, &mapping)

	return &EnvReader{EnvVars: mapping}
}

func (e *EnvReader) GetHost() string {
	return e.EnvVars.Redis[0].Credentials.Host
}

func (e *EnvReader) GetPassword() string {
	return e.EnvVars.Redis[0].Credentials.Password
}

func (e *EnvReader) GetPort() string {
	return strconv.Itoa(e.EnvVars.Redis[0].Credentials.Port)
}
