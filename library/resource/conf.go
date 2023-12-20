package resource

var Config struct {
	DB mysqlEnv `toml:"mysql"`
}

type mysqlEnv struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Charset  string

	MaxOpen  int
	MaxIdle  int
	MaxLife  int
	Debug    bool
	Log      bool
	LogFile  string
	LogLevel int
}
