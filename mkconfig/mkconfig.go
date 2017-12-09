package mkconfig

import (
	"io/ioutil"
	"log"
	"mkgo/utils/yaml"
	"os"
)

var Config RootConfig

const GlobalConfigFile = "./mkgo.yaml"

func Init() {
	bytes, err := readAll(GlobalConfigFile)
	if err != nil {
		log.Fatal("Failed read config file mkgo.yaml.", err)
	}
	err = yaml.Unmarshal(bytes, &Config)
	if err != nil {
		log.Fatal("Failed read config file mkgo.yaml.", err)
	}
}

type RootConfig struct {
	MLGO MLGOConfig `json:"mkgo"`
}

type MLGOConfig struct {
	Name         string     `json:"name"`
	ServerPort   string     `json:"server_port"`
	ReadTimeout  int        `json:"read_timeout"`
	WriteTimeout int        `json:"write_timeout"`
	Debug        bool       `json:"debug"`
	Log          Log        `json:"log"`
	Datasource   DataSource `json:"datasource"`
}

type Log struct {
	Path  string `json:"path"`
	Level string `json:"level"`
}

type DataSource struct {
	ReadSize int       `json:"read_size"`
	Write    *DBConfig `json:"write"`
	Read1    *DBConfig `json:"read1"`
	Read2    *DBConfig `json:"read2"`
	Read3    *DBConfig `json:"read3"`
	Read4    *DBConfig `json:"read4"`
	Read5    *DBConfig `json:"read5"`
}

type DBConfig struct {
	Name         string `json:"name"`
	Driver       string `json:"driver"`
	Host         string `json:"host"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
	ShowSql      bool   `json:"show_sql"`
	TableFix     string `json:"table_fix"`
	TableSpace   string `json:"table_space"`
	TableSnake   bool   `json:"table_snake"`
	ColumnFix    string `json:"column_fix"`
	ColumnSpace  string `json:"column_space"`
	ColumnSnake  bool   `json:"column_snake"`
	DisableCache bool   `json:"disable_cache"`
	ShowExecTime bool   `json:"show_exec_time"`
}

func readAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
