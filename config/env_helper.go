package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetDBHost() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbhost := os.Getenv("DB_HOST")

	return dbhost + " "
}

func GetDBUser() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbuser := os.Getenv("DB_USER")

	return dbuser + " "
}

func GetDBPassword() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbpassword := os.Getenv("DB_PASSWORD")

	return dbpassword + " "
}

func GetDBName() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbname := os.Getenv("DB_NAME")

	return dbname + " "
}

func GetDBPort() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbport := os.Getenv("DB_PORT")

	return dbport + " "
}

func GetDBSSLMode() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbsslmode := os.Getenv("DB_SSLMODE")

	return dbsslmode + " "
}

func GetDBTimezone() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbtimezone := os.Getenv("DB_TIMEZONE")

	return dbtimezone + " "
}

func GetJWTKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	jwtkey := os.Getenv("JWT_KEY")

	return jwtkey
}

func GetCLDURL() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	cldurl := os.Getenv("CLD_URL")

	return cldurl
}

func GetCLDFolder() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	cldfolder := os.Getenv("CLD_FOLDER")

	return cldfolder
}
