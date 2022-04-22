package config

type HttpConfig struct {
	Port        int    `json:"port"`
	Mode        string `json:"mode"`
	OpenRedis   bool   `json:"openRedis"`   //是否开启redis
	OpenCaptcha bool   `json:"openCaptcha"` //是否开启验证码
}

type Mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	DataBase string `json:"database"`
	LogoMode bool   `json:"logoMode"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type Jwt struct {
	Key        string `json:"key"`
	ExpireTime int64  `json:"expireTime"`
}

type Log struct {
	Path       string `json:"path"`
	Maxsize    int    `json:"maxsize"`
	MaxBackups int    `json:"maxBackups"`
	MaxAge     int    `json:"maxAge"`
	Compress   bool   `json:"compress"`
}

// Captcha 验证码
type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"key-long" yaml:"key-long"`       // 验证码长度
	ImgWidth  int `mapstructure:"img-width" json:"img-width" yaml:"img-width"`    // 验证码宽度
	ImgHeight int `mapstructure:"img-height" json:"img-height" yaml:"img-height"` // 验证码高度
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"model-path" yaml:"model-path"` // 存放casbin模型的相对路径
}

type Config struct {
	HttpConfig HttpConfig `json:"httpConfig"`
	Redis      Redis      `json:"redis"`
	Mysql      Mysql      `json:"mysql"`
	Jwt        Jwt        `json:"jwt"`
	Log        Log        `json:"log"`
	Salt       string     `json:"salt"` //加密盐
	Casbin     Casbin     `json:"casbin" yaml:"casbin"`
	Captcha    Captcha    `json:"captcha" yaml:"captcha"`
}

var config Config

func SetConfig(cfg *Config) {
	config = *cfg
}

func GetConfig() *Config {
	return &config
}
