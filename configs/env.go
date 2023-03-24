package configs

import (
	"os"
)

func EnvMongoURI() string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error Loading .env file")
	// }

	return os.Getenv("MONGO_URL")
}
