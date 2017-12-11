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
	ReadSize int       `yaml:"read_size"`
	Write    *DBConfig `yaml:"write"`
	Read1    *DBConfig `yaml:"read1"`
	Read2    *DBConfig `yaml:"read2"`
	Read3    *DBConfig `yaml:"read3"`
	Read4    *DBConfig `yaml:"read4"`
	Read5    *DBConfig `yaml:"read5"`
}

type DBConfig struct {
	Name         string `yaml:"name"`
	Driver       string `yaml:"driver"`
	Host         string `yaml:"host"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	ShowSql      bool   `yaml:"show_sql"`
	TableFix     string `yaml:"table_fix"`
	TableSpace   string `yaml:"table_space"`
	TableSnake   bool   `yaml:"table_snake"`
	ColumnFix    string `yaml:"column_fix"`
	ColumnSpace  string `yaml:"column_space"`
	ColumnSnake  bool   `yaml:"column_snake"`
	DisableCache bool   `yaml:"disable_cache"`
	ShowExecTime bool   `yaml:"show_exec_time"`
}
