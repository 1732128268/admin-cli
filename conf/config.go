package config

type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	DataBase string `json:"database"`
	LogoMode bool   `json:"logoMode"`
}

type Jwt struct {
	Key string `json:"key"`
}

type Log struct {
	Path       string `json:"path"`
	Maxsize    int    `json:"maxsize"`
	MaxBackups int    `json:"maxBackups"`
	MaxAge     int    `json:"maxAge"`
	Compress   bool   `json:"compress"`
}

type Config struct {
	Mysql     MysqlConf `json:"mysql"`
	Jwt       Jwt       `json:"jwt"`
	Log       Log       `json:"log"`
	OpenCache bool      `json:"openCache"`
	HostName  string    `json:"hostName"`
}

var config Config

func SetConfig(cfg *Config) {
	config = *cfg
}

func GetConfig() *Config {
	return &config
}
