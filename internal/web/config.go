package web

import (
	"github.com/bigredeye/notmanytask/pkg/conf"
	"github.com/pkg/errors"
)

type Config struct {
	GitLab struct {
		Application struct {
			ClientID string
			Secret   string
		}
		Api struct {
			Token string
		}
	}
	Endpoints struct {
		HostName      string
		Home          string
		Login         string
		Logout        string
		Signup        string
		OauthCallback string
	}
	Server struct {
		ListenAddress string
		Cookies       struct {
			AuthenticationKey string
			EncryptionKey     string
		}
	}
	Postgres struct {
		Host     string
		Port     uint16
		Username string
		Password string
		DataBase string
	}
}

func ParseConfig() (*Config, error) {
	config := &Config{}
	if err := conf.ParseConfig(config, conf.EnvPrefix("NMT")); err != nil {
		return nil, errors.Wrap(err, "Failed to parse config")
	}
	return config, nil
}
