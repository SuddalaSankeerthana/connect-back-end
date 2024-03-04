package env

import (
	"fmt"
	"log"
	"os"

	"path/filepath"

	"github.com/joho/godotenv"
)

func GetDotEnvVariable(key string) string {

	abs, err := filepath.Abs(".env")

	fmt.Println(abs)
	err = godotenv.Load(abs)

	if err != nil {
		log.Println(err)
	}

	return os.Getenv(key)
}
