package mkconfig

import (
	"io/ioutil"
	"time"
	"gopkg.in/yaml.v2"
)

var Config RootConfig

const globalConfigFile = "./mkgo.yaml"

func Init() {
	data, err := ioutil.ReadFile(globalConfigFile)
	if err != nil {
		panic("Config file mkgo.yaml not found")
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		panic("Config file mkgo.yaml parse failed")
	}
}

type RootConfig struct {
	MKGo `yaml:"mkgo"`
}

type MKGo struct {
	Name         string        `yaml:"name"`
	ServerPort   string        `yaml:"server_port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	Debug        bool          `yaml:"debug"`
	Log          Log           `yaml:"log"`
	Redis                      `yaml:"redis"`
	DataSource                 `yaml:"data_source"`
}

type Log struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
}

type Redis struct {
	Host      string `yaml:"host"`
	MaxIdle   int    `yaml:"max_idle"`
	MaxActive int    `yaml:"max_active"`
}

type DataSource struct {
	MaxOpenConns int      `yaml:"max_open_conns"`
	MaxIdleConns int      `yaml:"max_idle_conns"`
	Write        []string `yaml:"write"`
	Read         []string `yaml:"read"`
}
