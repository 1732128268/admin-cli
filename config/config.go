package config

type HttpConfig struct {
	Port        int    `json:"port"`
	Mode        string `json:"mode"`
	OpenRedis   bool   `json:"openRedis"`                                        //是否开启redis
	OpenCaptcha bool   `json:"openCaptcha"`                                      //是否开启验证码
	OssType     string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"` // Oss类型
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

type AliyunOSS struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `mapstructure:"access-key-id" json:"access-key-id" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"access-key-secret" yaml:"access-key-secret"`
	BucketName      string `mapstructure:"bucket-name" json:"bucket-name" yaml:"bucket-name"`
	BucketUrl       string `mapstructure:"bucket-url" json:"bucket-url" yaml:"bucket-url"`
	BasePath        string `mapstructure:"base-path" json:"base-path" yaml:"base-path"`
}

type HuaWeiObs struct {
	Path      string `mapstructure:"path" json:"path" yaml:"path"`
	Bucket    string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Endpoint  string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKey string `mapstructure:"access-key" json:"access-key" yaml:"access-key"`
	SecretKey string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`
}
type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`                                  // 存储区域
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`                            // 空间名称
	ImgPath       string `mapstructure:"img-path" json:"img-path" yaml:"img-path"`                      // CDN加速域名
	UseHTTPS      bool   `mapstructure:"use-https" json:"use-https" yaml:"use-https"`                   // 是否使用https
	AccessKey     string `mapstructure:"access-key" json:"access-key" yaml:"access-key"`                // 秘钥AK
	SecretKey     string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`                // 秘钥SK
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"use-cdn-domains" yaml:"use-cdn-domains"` // 上传是否使用CDN上传加速
}

type TencentCOS struct {
	Bucket     string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Region     string `mapstructure:"region" json:"region" yaml:"region"`
	SecretID   string `mapstructure:"secret-id" json:"secret-id" yaml:"secret-id"`
	SecretKey  string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`
	BaseURL    string `mapstructure:"base-url" json:"base-url" yaml:"base-url"`
	PathPrefix string `mapstructure:"path-prefix" json:"path-prefix" yaml:"path-prefix"`
}

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件路径
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
	Local      Local      `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu      Qiniu      `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	AliyunOSS  AliyunOSS  `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`
	HuaWeiObs  HuaWeiObs  `mapstructure:"hua-wei-obs" json:"hua-wei-obs" yaml:"hua-wei-obs"`
	TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencent-cos" yaml:"tencent-cos"`
}

var config Config

func SetConfig(cfg *Config) {
	config = *cfg
}

func GetConfig() *Config {
	return &config
}
