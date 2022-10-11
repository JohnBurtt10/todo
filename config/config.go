package config

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
	O            interface{}
	errorHandler fiber.ErrorHandler
	fiber        *fiber.Config
	jwtSecret    string
	// clientTLSConfig *tls.Config
	AppVersion    string
	AppGitCommit  string
	AppGitBranch  string
	AppGitState   string
	AppGitSummary string
	AppBuildDate  string
}

var instantiated *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		instantiated = &Config{
			Viper: viper.New(),
		}

		// Set default configurations
		instantiated.setDefaults()

		// Select the .env file
		instantiated.SetConfigName(".env")
		instantiated.SetConfigType("dotenv")
		instantiated.AddConfigPath(".")

		// Automatically refresh environment variables
		instantiated.AutomaticEnv()

	})
	return instantiated
}

func (config *Config) setDefaults() {
	config.SetDefault("MW_FIBER_SESSION_STORAGE_HOST", "localhost")
	config.SetDefault("MW_FIBER_SESSION_STORAGE_PORT", 6379)
	config.SetDefault("MW_FIBER_SESSION_STORAGE_USERNAME", "")
	config.SetDefault("MW_FIBER_SESSION_STORAGE_PASSWORD", "")
	config.SetDefault("MW_FIBER_SESSION_STORAGE_DATABASE", "mytestdb")
	config.SetDefault("MW_FIBER_SESSION_STORAGE_RESET", false)
	config.SetDefault("MW_FIBER_SESSION_EXPIRATION", "24h")

}

func (config *Config) GetSessionConfig() session.Config {
	var store fiber.Storage
	redisCfg := redis.Config{
		Host:     config.GetString("MW_FIBER_SESSION_STORAGE_HOST"),
		Port:     config.GetInt("MW_FIBER_SESSION_STORAGE_PORT"),
		Username: config.GetString("MW_FIBER_SESSION_STORAGE_USERNAME"),
		Password: config.GetString("MW_FIBER_SESSION_STORAGE_PASSWORD"),
		Database: config.GetInt("MW_FIBER_SESSION_STORAGE_DATABASE"),
		Reset:    config.GetBool("MW_FIBER_SESSION_STORAGE_RESET"),
	}
	store = redis.New(redisCfg)

	return session.Config{
		Expiration: config.GetDuration("MW_FIBER_SESSION_EXPIRATION"),
		Storage:    store,
		// KeyLookup:      fmt.Sprintf("cookie:%s", config.GetString("MW_FIBER_SESSION_COOKIENAME")),
		// CookieDomain:   config.GetString("MW_FIBER_SESSION_COOKIEDOMAIN"),
		// CookiePath:     config.GetString("MW_FIBER_SESSION_COOKIEPATH"),
		// CookieSecure:   config.GetBool("MW_FIBER_SESSION_COOKIESECURE"),
		// CookieHTTPOnly: config.GetBool("MW_FIBER_SESSION_COOKIEHTTPONLY"),
		// CookieSameSite: config.GetString("MW_FIBER_SESSION_COOKIESAMESITE"),
	}
}
