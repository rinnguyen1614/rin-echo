package config

import (
	"errors"

	"github.com/spf13/viper"
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
		App struct {
			// The application uniqueID. Once generated, don't modify
			AppID   string `mapstructure:"app_id,omitempty"`
			AppName string `mapstructure:"app_name,omitempty" `
			// Env is the environment,which maybe development, testing, production.
			Env string `mapstructure:"env,omitempty"`

			Encryption string `mapstructure:"encryption,omitempty"`

			MaxPageSize int `mapstructure:"max_page_size,omitempty"`
			MinPageSize int `mapstructure:"min_page_size,omitempty"`

			Resources map[string]string `mapstructure:"resources,omitempty"`

			UploadPath string `mapstructure:"upload_path,omitempty"`
		}

		Casbin struct {
			Enable           bool   `mapstructure:"enable,omitempty"`
			ModelPath        string `mapstructure:"model_path,omitempty"`
			FilePath         string `mapstructure:"file_path,omitempty"`
			AutoLoad         bool   `mapstructure:"auto_load,omitempty"`
			AutoLoadInternal uint   `mapstructure:"auto_load_internal,omitempty"`
		}

		Server struct {
			Host string `mapstructure:"host,omitempty"`
			Port string `mapstructure:"port,omitempty"`
		}

		Mail struct {
			Host string `mapstructure:"host,omitempty"`
			Port int    `mapstructure:"port,omitempty"`
			User string `mapstructure:"user,omitempty"`
			Pwd  string `mapstructure:"pwd,omitempty"`
		}

		Log struct {
			Level         string `mapstructure:"level,omitempty"`
			Prefix        string `mapstructure:"prefix,omitempty"`
			StacktraceKey string `mapstructure:"stack_strace_key,omitempty"`
		}

		JWT struct {
			SigningKey  string `mapstructure:"signing_key,omitempty"`
			ExpiresTime int64  `mapstructure:"expires_time,omitempty"`
			Issuer      string `mapstructure:"issuer,omitempty"`
		}

		Database struct {
			Driver       string `mapstructure:"driver,omitempty"`
			URL          string `mapstructure:"url,omitempty"`
			BatchSize    int    `mapstructure:"batch_size,omitempty"`
			MigrationURL string `mapstructure:"migration_url,omitempty"`
			InitData     bool   `mapstructure:"init_data,omitempty"`
		}
	}
)

func LoadConfig(path string) (conf *Config, err error) {
	v := viper.New()

	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Config file not found; ignore error if desired")
		}
		return nil, errors.New("Config file was found but another error was produced")
	}

	err = v.Unmarshal(&conf)
	return
}

func (s *Config) IsProduction() bool {
	return s.App.Env == EnvProd
}

func (s *Config) IsDevelopment() bool {
	return s.App.Env == EnvDevelopment
}

func (s *Config) IsTesting() bool {
	return s.App.Env == EnvTest
}
