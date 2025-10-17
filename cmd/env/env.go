package env

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type EnvKey string

func (key EnvKey) GetValue() string {
	return os.Getenv(string(key))
}

const (
	Env                 EnvKey = "ENV"
	DBURL               EnvKey = "DATABASE_URL"
	GithubWebhookSecret EnvKey = "GITHUB_WEBHOOK_SECRET"
	Port                EnvKey = "PORT"
)

func Load() error {
	_, b, _, _ := runtime.Caller(0)
	rootPath := filepath.Join(filepath.Dir(b), "../../")
	envPath := rootPath + "/.env"

	fmt.Printf("Looking for .env file at: %s\n", envPath)

	// Check if file exists
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		fmt.Printf(".env file does not exist at: %s\n", envPath)
	} else {
		fmt.Printf(".env file found at: %s\n", envPath)
	}

	return godotenv.Load(envPath)
}
