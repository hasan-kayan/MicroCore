package config

import (
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Name            string        `env:"APP_NAME" envDefault:"authservice"`
		Env             string        `env:"APP_ENV" envDefault:"dev"`
		HTTPAddr        string        `env:"HTTP_ADDR" envDefault:":8080"`
		Version         string        `env:"VERSION" envDefault:"0.1.0"`
		ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT_SEC" envDefault:"15s"`
	}
	Logging struct {
		Level string `env:"LOG_LEVEL" envDefault:"debug"`
	}
	Mongo struct {
		URI              string `env:"MONGO_URI" envDefault:"mongodb://localhost:27017"`
		DB               string `env:"MONGO_DB" envDefault:"authservice"`
		HealthCollection string `env:"MONGO_HEALTH_COLLECTION" envDefault:"health_events"`
	}
}

// Load loads env (optionally from .env) and returns Config.
// YAML dosyası örnek amaçlıdır; prod’da değerler ENV’den gelir.
func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	// configs/app.example.yaml içinden okumayı zorunlu kılmıyoruz;
	// istersen burada YAML parse ekleyebiliriz. ENV her zaman kaynak.
	_ = os.Setenv("APP_NAME", cfg.App.Name) // (örnek: başka paketlerin ihtiyaç duyması için)

	return cfg, nil
}
