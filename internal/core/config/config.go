package config

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var (
	// EnvTest is a const value of test environment.
	EnvTest = "test"
	// EnvLocal is a const value of local environment.
	EnvDevelopment = "development"
	// EnvProd is a const value of production environment.
	EnvProd = "prod"
)

type (
	Config struct {
		App      App      `json:"app,omitempty" yaml:"app,omitempty" `
		Casbin   Casbin   `json:"casbin,omitempty" yaml:"casbin,omitempty" `
		Database Database `json:"database,omitempty" yaml:"database,omitempty" `
		JWT      JWT      `json:"jwt,omitempty" yaml:"jwt,omitempty" `
		Server   Server   `json:"server,omitempty" yaml:"server" `
		Mail     Mail     `json:"mail,omitempty" yaml:"mail" `
	}

	App struct {
		// The application uniqueID. Once generated, don't modify
		AppID   string `json:"app_id,omitempty" yaml:"app_id,omitempty" `
		AppName string `json:"app_name,omitempty" yaml:"app_name,omitempty" `
		// Env is the environment,which maybe development, testing, production.
		Env string `json:"env,omitempty" yaml:"env,omitempty" `

		Encryption string `json:"encryption,omitempty" yaml:"encryption,omitempty" `

		MaxPageSize int `json:"max_page_size,omitempty" yaml:"max_page_size,omitempty" `
		MinPageSize int `json:"min_page_size,omitempty" yaml:"min_page_size,omitempty" `

		Resources map[string]string `json:"resources,omitempty" yaml:"resources,omitempty" `

		UploadPath string `json:"upload_path,omitempty" yaml:"upload_path,omitempty" `
	}

	Casbin struct {
		Enable           bool   `json:"enable,omitempty" yaml:"enable,omitempty" `
		ModelPath        string `json:"model_path,omitempty" yaml:"model_path,omitempty" `
		FilePath         string `json:"file_path,omitempty" yaml:"file_path,omitempty" `
		AutoLoad         bool   `json:"auto_load,omitempty" yaml:"auto_load,omitempty" `
		AutoLoadInternal uint   `json:"auto_load_internal,omitempty" yaml:"auto_load_internal,omitempty" `
	}

	Database struct {
		Host       string            `json:"host,omitempty" yaml:"host,omitempty" `
		Port       string            `json:"port,omitempty" yaml:"port,omitempty" `
		User       string            `json:"user,omitempty" yaml:"user,omitempty" `
		Pwd        string            `json:"pwd,omitempty" yaml:"pwd,omitempty" `
		Name       string            `json:"name,omitempty" yaml:"name,omitempty" `
		Driver     string            `json:"driver,omitempty" yaml:"driver,omitempty" `
		DNS        string            `json:"dns,omitempty" yaml:"dns,omitempty" `
		Params     map[string]string `json:"params,omitempty" yaml:"params,omitempty" `
		MaxIdleCon int               `json:"max_idle_con,omitempty" yaml:"max_idle_con,omitempty" `
		MaxOpenCon int               `json:"max_open_con,omitempty" yaml:"max_open_con,omitempty" `
		BatchSize  int               `json:"batch_size,omitempty" yaml:"batch_size,omitempty" `
	}

	JWT struct {
		SigningKey  string `json:"signing_key,omitempty" yaml:"signing_key,omitempty" `
		ExpiresTime int64  `json:"expires_time,omitempty" yaml:"expires_time,omitempty" `
		Issuer      string `json:"issuer,omitempty" yaml:"issuer,omitempty" `
	}

	Server struct {
		Host string `json:"host,omitempty" yaml:"host,omitempty" `
		Port string `json:"port,omitempty" yaml:"port,omitempty" `
	}

	Log struct {
		Level         string `json:"level,omitempty" yaml:"level,omitempty" `
		Prefix        string `json:"prefix,omitempty" yaml:"prefix,omitempty" `
		StacktraceKey string `json:"stack_strace_key,omitempty" yaml:"stack_strace_key,omitempty" `
	}

	Mail struct {
		Host string `json:"host,omitempty" yaml:"host,omitempty" `
		Port int    `json:"port,omitempty" yaml:"port,omitempty" `
		User string `json:"user,omitempty" yaml:"user,omitempty" `
		Pwd  string `json:"pwd,omitempty" yaml:"pwd,omitempty" `
	}
)

func (s *Config) IsProduction() bool {
	return s.App.Env == EnvProd
}

func (s *Config) IsDevelopment() bool {
	return s.App.Env == EnvDevelopment
}

func (s *Config) IsTesting() bool {
	return s.App.Env == EnvTest
}

// ReadFromJson read the Config from a JSON file.
func ReadFromJson(path string) Config {
	jsonByte, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var cfg Config

	err = json.Unmarshal(jsonByte, &cfg)

	if err != nil {
		panic(err)
	}

	return cfg
}

// ReadFromYaml read the Config from a YAML file.
func ReadFromYaml(path string) Config {
	jsonByte, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var cfg Config

	err = yaml.Unmarshal(jsonByte, &cfg)

	if err != nil {
		panic(err)
	}

	return cfg
}
