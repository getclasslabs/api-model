package config

type Configuration struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Jaeger struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"jaeger"`
	Mysql Mysql `yaml:"mysql"`
}

type Mysql struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (m Mysql) GetUser() string {
	return m.User
}

func (m Mysql) GetPassword() string {
	return m.Password
}

func (m Mysql) GetHost() string {
	return m.Host
}

func (m Mysql) GetPort() string {
	return m.Port
}

func (m Mysql) GetDatabase() string {
	return m.Database
}

var Config Configuration
