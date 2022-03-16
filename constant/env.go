package constant

import (
	"log"
	"os"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type envConfig struct {
	Environment                 string `env:"ENV" envDefault:"development"`
	ServiceName                 string `env:"SERVICE_NAME" envDefault:"template-new-api"`
	ServiceVersion              string `env:"SERVICE_VERSION" envDefault:"1.0.0"`
	Port                        int    `env:"PORT" envDefault:"10000"`
	MySqlUri                    string `env:"SQL_URI" envDefault:"root:@tcp(127.0.0.1:3306)/transformation_ivr_db"`
	InternalSenlightApiHost     string `env:"INTERNAL_API_TRANSFORMATION_ADMIN_API_HOST" envDefault:""`
	InternalChatBusinessApiHost string `env:"INTERNAL_SERVICES_CHAT_BUSINESS_API_HOST" envDefault:""`
}

// AppEnv provide config values
type AppEnv interface {
	loadEnvFromFile()
	GetAppEnvironment() string
	GetServiceName() string
	GetServiceVersion() string
	GetPort() int
	GetMySqlUri() string
	GetInternalSenlightApiHost() string
	GetInternalChatBusinessApiHost() string
}

type appEnv struct {
	config envConfig
}

var _env AppEnv

// init initialize ONLY ONCE default AppEnv instance
// this function will be called when import the package
// https://golang.org/doc/effective_go.html#init
func init() {
	_env = &appEnv{}
	_env.loadEnvFromFile()
	log.Println("Initialized app environment successfully!")
}

// Env get app configs. Example: constant.Env().GetPort()
func Env() AppEnv {
	return _env
}

// loadEnvFromFile load and parse environment variables
func (ae *appEnv) loadEnvFromFile() {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	loadEnvErr := godotenv.Load(envFile)
	if loadEnvErr != nil {
		log.Fatal(loadEnvErr)
	}

	parseErr := env.Parse(&ae.config)
	if parseErr != nil {
		log.Fatal(parseErr)
	}
}

func (ae *appEnv) GetMySqlUri() string {
	dbPath := ae.config.MySqlUri
	dbPath = strings.Replace(dbPath, "mysql://", "", 1)
	return dbPath
}

// GetAppEnvironment get app environment
func (ae *appEnv) GetAppEnvironment() string {
	return ae.config.Environment
}

// GetServiceName get service name
func (ae *appEnv) GetServiceName() string {
	return ae.config.ServiceName
}

// GetServiceVersion get service version
func (ae *appEnv) GetServiceVersion() string {
	return ae.config.ServiceVersion
}

// GetPort get app port
func (ae *appEnv) GetPort() int {
	return ae.config.Port
}

// GetInternalSenlightApiHost
func (ae *appEnv) GetInternalSenlightApiHost() string {
	return ae.config.InternalSenlightApiHost
}

// Get internal chat Business host
func (ae *appEnv) GetInternalChatBusinessApiHost() string {
	return ae.config.InternalChatBusinessApiHost
}
