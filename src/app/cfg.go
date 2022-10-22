package app

import (
	"os"

	"github.com/joho/godotenv"
)

const HTTP_ROUTES_PREFIX = "api/v2"

var (
	STORAGE_PATH    string
	WEB_SERVER_PORT string
)

func LoadEnv(dotenvPath string) {
	godotenv.Load(dotenvPath)

	STORAGE_PATH = os.Getenv("STORAGE_PATH")
	WEB_SERVER_PORT = os.Getenv("WEB_SERVER_PORT")
}
