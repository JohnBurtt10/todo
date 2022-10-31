package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"github.com/gofiber/template/html"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
	O            interface{}
	errorHandler fiber.ErrorHandler
	fiber        *fiber.Config
	database     *DatabaseConfig
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

var defaultErrorHandler = func(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Set error message
	// var message interface{}

	// Check if it's a fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		// if m, ok := e.Message.(string); ok {
		// 	message = m
		// }
	}

	// TODO: Check return type for the client, JSON, HTML, YAML or any other (API vs web)

	// Return HTTP response
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	c.Status(code)
	// Return statuscode with error message
	return c.Status(code).SendString(err.Error())

	// // Render default error view
	// if renderErr := c.Render("errors/"+strconv.Itoa(code), fiber.Map{"message": message}); renderErr != nil {
	// 	// Set Content-Type: text/plain; charset=utf-8
	// 	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// 	// Return statuscode with error message
	// 	return c.Status(code).SendString(err.Error())
	// }
	// return err
}

func GetInstance() *Config {
	once.Do(func() {
		instantiated = &Config{
			Viper: viper.New(),
		}

		// Set default configurations
		instantiated.setDefaults()

		// Set Fiber configurations
		instantiated.setFiberConfig()

		// Select the .env file
		instantiated.SetConfigName(".env")
		instantiated.SetConfigType("dotenv")
		instantiated.AddConfigPath(".")

		// Read configuration
		if err := instantiated.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				fmt.Println("failed to read configuration:", err.Error())
				os.Exit(1)
			}
		}

		instantiated.SetErrorHandler(defaultErrorHandler)

		// Automatically refresh environment variables
		instantiated.AutomaticEnv()
		instantiated.setDatabaseConfig()

	})
	return instantiated
}

func (config *Config) SetErrorHandler(errorHandler fiber.ErrorHandler) {
	config.errorHandler = errorHandler
}

func (config *Config) setDefaults() {
	// Set default database configuration
	config.SetDefault("DB_HOST", "localhost")
	config.SetDefault("DB_USERNAME", "postgres")
	config.SetDefault("DB_PASSWORD", "mysecretpassword")
	config.SetDefault("DB_PORT", 5432)
	config.SetDefault("DB_NAME", "todo")

	// Set default session configuration
	config.SetDefault("MW_FIBER_SESSION_STORAGE_HOST", "localhost")
	config.SetDefault("MW_FIBER_SESSION_STORAGE_PORT", 6379)
	config.SetDefault("MW_FIBER_SESSION_STORAGE_USERNAME", "")
	config.SetDefault("MW_FIBER_SESSION_STORAGE_PASSWORD", "")
	config.SetDefault("MW_FIBER_SESSION_STORAGE_DATABASE", "todo")
	config.SetDefault("MW_FIBER_SESSION_STORAGE_RESET", false)
	config.SetDefault("MW_FIBER_SESSION_EXPIRATION", "24h")

}

func (config *Config) setDatabaseConfig() {
	config.database = &DatabaseConfig{
		Default: DatabaseDriver{
			Driver:   config.GetString("DB_DRIVER"),
			Host:     config.GetString("DB_HOST"),
			Username: config.GetString("DB_USERNAME"),
			Password: config.GetString("DB_PASSWORD"),
			DBName:   config.GetString("DB_NAME"),
			Port:     config.GetInt("DB_PORT"),
		},
	}
}

func (config *Config) GetDatabaseConfig() *DatabaseConfig {
	return config.database
}

func (config *Config) setFiberConfig() {
	config.fiber = &fiber.Config{
		Views:        config.getFiberViewsEngine(),
		ErrorHandler: config.errorHandler,
	}
}

func (config *Config) GetFiberConfig() *fiber.Config {
	return config.fiber
}

func (config *Config) getFiberViewsEngine() fiber.Views {
	engine := html.New("./views", ".html")
	return engine
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
