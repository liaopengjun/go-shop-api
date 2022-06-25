package config

type AppConfig struct {
	*ApplicationConfig `mapstructure:"application"`
	*MySQLConfig       `mapstructure:"mysql"`
	*RedisConfig       `mapstructure:"redis"`
	*LogConfig         `mapstructure:"log"`
	*JwtConfig         `mapstructure:"jwt"`
	*CasbinConfig      `mapstructure:"casbin"`
	*LocalConfig       `mapstructure:"local"`
}

type ApplicationConfig struct {
	Mode       string `mapstructure:"mode"`
	Port       int    `mapstructure:"port"`
	UploadType string `mapstructure:"upload_type"`
	UserRedis  bool   `mapstructure:"user_redis"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type LogConfig struct {
	Level         string `mapstructure:"level"`          // 级别
	Format        string `mapstructure:"format"`         // 输出
	Prefix        string `mapstructure:"prefix"`         // 日志前缀
	Director      string `mapstructure:"director"`       // 日志文件夹
	ShowLine      bool   `mapstructure:"show-line"`      // 显示行
	EncodeLevel   string `mapstructure:"encode-level"`   // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console"` // 输出控制台
}

type JwtConfig struct {
	SigningKey  string `mapstructure:"signing-key"`  // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer-time"`  // 缓冲时间
	Issuer      string `mapstructure:"issuer"`       // 签发者
}

type CasbinConfig struct {
	ModelPath string `mapstructure:"model-path"` // 存放casbin模型的相对路径
}

type LocalConfig struct {
	Path            string `mapstructure:"path"`
	ImageMaxSize    int    `mapstructure:"image_max_size"`
	ImageAllowExits string `mapstructure:"image_allow_exits"`
}
