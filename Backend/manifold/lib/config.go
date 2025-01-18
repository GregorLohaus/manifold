package lib

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
	"path"
)

type (
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
	}

	S3Bucket struct {
		Endpoint  string
		AccessKey string
		Secret    string
		SSL       bool
	}

	Server struct {
		Port              int
		AdminPasswordHash string
		FrontendProtocol  string
		FrontendHost      string
		FrontendPort      int
	}

	Config struct {
		Database Database
		S3Bucket S3Bucket
		Server   Server
	}
)

const CONFIG_HOME = "manifold"
const DATA_HOME = "manifold"
const BIN_HOME = ".local/bin"
const CONFIG_HOME_SYSTEMD_USER = "user"
const CONFIG_FILE_NAME = "config.toml"
const CORE_LOG_FILE_NAME = "manifold.log"

var CurrentConfig *Config

func ConfigFilePath() (*string, error) {
	xdgConfigHome, err := GetXDGVar(XDG_CONFIG_HOME)
	if err != nil {
		return nil, err
	}
	str := path.Join(*xdgConfigHome, CONFIG_HOME, CONFIG_FILE_NAME)
	return &str, nil
}

func CoreLogFilePath() (*string, error) {
	xdgDataHome, err := GetXDGVar(XDG_DATA_HOME)
	if err != nil {
		return nil, err
	}
	str := path.Join(*xdgDataHome, DATA_HOME, CORE_LOG_FILE_NAME)
	return &str, nil
}

func GetConfig(path *string) (*Config, error) {
	if CurrentConfig != nil {
		return CurrentConfig, nil
	}
	configPath := path
	err := (func() error { return nil })()
	if configPath == nil || (configPath != nil && *configPath == "") {
		configPath, err = ConfigFilePath()
		if err != nil {
			return nil, err
		}
	}
	str, error := os.ReadFile(*configPath)
	if error != nil {
		return nil, error
	}
	var config Config
	err = toml.Unmarshal(str, &config)
	if err != nil {
		return nil, err
	}
	CurrentConfig = &config
	return &config, nil
}

func CreateDefaultConfig() error {
	configPath, err := ConfigFilePath()
	if err != nil {
		return err
	}
	err = os.MkdirAll(path.Dir(*configPath), os.FileMode(0755))
	if err != nil {
		return err
	}
	str, err := os.ReadFile(*configPath)
	if err == nil {
		var config Config
		err = toml.Unmarshal(str, &config)
		if err == nil {
			return nil
		}
	}
	file, err := os.OpenFile(*configPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(0755))
	defer file.Close()
	if err != nil {
		return err
	}
	config := Config{
		Database: Database{},
		S3Bucket: S3Bucket{},
		Server:   Server{},
	}
	str, err = toml.Marshal(config)
	if err != nil {
		return err
	}
	status, err := file.Write(str)
	if err != nil {
		return err
	}
	os.Stderr.WriteString(fmt.Sprintf("Config creation finished with status: %d \n", status))
	return nil
}
