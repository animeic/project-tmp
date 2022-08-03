package model

type Config struct {
	App   `yaml:"app"`
	Mysql `yaml:"mysql"`
	Jwt   `yaml:"jwt"`
	Redis `yaml:"redis"`
	Log   `yaml:"log"`
	Asset `yaml:"asset"`
}

type App struct {
	Env  string `yaml:"env"`
	Mode string `yaml:"mode"`
	Port string `yaml:"port"`
	Name string `yaml:"app_name"`
	Host string `yaml:"app_url"`
}

type Mysql struct {
	Host                string `yaml:"host"`
	Port                int    `yaml:"port"`
	Database            string `yaml:"database"`
	UserName            string `yaml:"username"`
	Password            string `yaml:"password"`
	Charset             string `yaml:"charset"`
	MaxIdleConns        int    `mapstructure:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns        int    `mapstructure:"max_open_conns" yaml:"max_open_conns"`
	LogMode             string `mapstructure:"log_mode" yaml:"log_mode"`
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" yaml:"enable_file_log_writer"`
	LogFilename         string `mapstructure:"log_filename" yaml:"log_filename"`
}
type Jwt struct {
	Key                     string `yaml:"key"`
	Secret                  string `yaml:"secret"`
	JwtTtl                  int64  `mapstructure:"jwt_ttl" yaml:"jwt_ttl"` // token 有效期（秒）
	JwtBlacklistGracePeriod int64  `mapstructure:"jwt_blacklist_grace_period" yaml:"jwt_blacklist_grace_period"`
	RefreshGracePeriod      int64  `mapstructure:"refresh_grace_period" yaml:"refresh_grace_period"` // token 自动刷新宽限时间（秒）
}
type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
}

type Log struct {
	Level      string `yaml:"level"`
	RootDir    string `mapstructure:"root_dir" yaml:"root_dir"`
	Filename   string `yaml:"filename"`
	Format     string `yaml:"format"`
	ShowLine   bool   `mapstructure:"show_line" yaml:"show_line"`
	MaxBackups int    `mapstructure:"max_backups" yaml:"max_backups"`
	MaxSize    int    `mapstructure:"max_size" yaml:"max_size"` // MB
	MaxAge     int    `mapstructure:"max_age" yaml:"max_age"`   // day
	Compress   bool   `yaml:"compress"`
}

type Asset struct {
	ImageDir   string `mapstructure:"image_dir" yaml:"image_dir"`
	PrefixHost string `mapstructure:"prefix_host" yaml:"prefix_host"`
}
