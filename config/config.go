package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"

	"go-skeleton/pkg/utils/file"
	"go-skeleton/pkg/utils/json"
)

// Config global setting
var Config = struct {
	DB struct {
		Username     string `env:"DB_USERNAME" envDefault:""`
		Password     string `env:"DB_PASSWORD" envDefault:""`
		Name         string `env:"DB_NAME" envDefault:""`
		Host         string `env:"DB_HOST" envDefault:"db"`
		Port         string `env:"DB_PORT" envDefault:"3306"`
		Encoding     string `env:"DB_ENCODING" envDefault:"utf8mb4"`
		Maxconns     uint64 `env:"DB_MAXCONNS" envDefault:"10"`
		Maxidleconns uint64 `env:"DB_MAXIDLECONNS" envDefault:"10"`
		Timeout      uint64 `env:"DB_TIMEOUNT" envDefault:"5000"`
		Debug        bool   `env:"DB_DEBUG" envDefault:"true"`
	}

	Redis struct {
		Address    string `env:"REDIS_SERVER" envDefault:"127.0.0.1:6379"`
		Password   string `env:"REDIS_PASSWORD" envDefault:""`
		DB         int    `env:"REDIS_DB" envDefault:"0"`
		TLSEnabled bool   `env:"REDIS_TLS_ENABLED" envDefault:"false"`
	}

	System struct {
		AppServer string `env:"SYSTEM_SERVER" envDefault:"127.0.0.1"`
		AppAddr   string `env:"SYSTEM_ADDR" envDefault:":7000"`
		Mode      string `env:"SYSTEM_MODE" envDefault:"debug"`
		TimeZone  string `env:"SYSTEM_TIME_ZONE" envDefault:"Asia/Jakarta"`
	}

	APIServer struct {
		LocalePath string `env:"APISERVER_LOCALE_PATH" envDefault:"cmd/apiserver/app/locale/"`
	}

	AESCryptKey string `env:"AES_CRYPT_KEY" envDefault:""`

	GoogleRecaptcha struct {
		VerifyUrl string `env:"GOOGLE_RECAPTCHA_VERIFY_URL" envDefault:"https://www.google.com/recaptcha/api/siteverify"`
		Secret    string `env:"GOOGLE_RECAPTCHA_SECRET" envDefault:""`
	}
}{}

// Init initialize config
func Init() {
	if file.Exists("./.env") {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}

	if err := env.Parse(&Config); err != nil {
		panic(err)
	}
}

// Show for cli print setting
func Show() {
	Init()
	str, _ := json.MarshalIndent(Config, "", " ")
	fmt.Printf("Config: %s\n", str)
}
