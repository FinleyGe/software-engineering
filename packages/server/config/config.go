package config

type ConfigStruct struct {
	Dev      bool     `mapstructure:"dev"`
	Server   server   `mapstructure:"server"`
	Database database `mapstructure:"database"`
	Jwt      jwt      `mapstructure:"jwt"`
}

type server struct {
	Port            string   `mapstructure:"port"`
	AllowOrigins    []string `mapstructure:"allow_origins"`
	LogFilePath     string   `mapstructure:"log_file_path"`
	ApiRoute        string   `mapstructure:"api_route"`
	Domain          string   `mapstructure:"domain"`
	DefaultPassword string   `mapstructure:"default_password"`
}

type database struct {
	DBName   string `mapstructure:"db_name"`
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type jwt struct {
	Issuer  string `mapstructure:"issuer"`
	Expires int    `mapstructure:"expires"`
	Secret  string `mapstructure:"secret"`
}
