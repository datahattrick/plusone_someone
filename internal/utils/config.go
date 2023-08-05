package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		Host string `envconfig:"HOSTNAME"`
		Port string `envconfig:"PORT"`
	}
	LDAP struct {
		Host          string   `envconfig:"LDAP_ENDPOINT"`
		Bind          string   `envconfig:"LDAP_BIND_ACCOUNT"`
		Password      string   `envconfig:"LDAP_BIND_PASSWORD"`
		UserGroup     string   `envconfig:"LDAP_USER_GROUP"`
		Attribute     []string `envconfig:"LDAP_USER_ATTRIBUTES"`
		FirstNameAttr string   `envconfig:"LDAP_USER_FIRSTNAME"`
		LastNameAttr  string   `envconfig:"LDAP_USER_LASTNAME"`
		EmailAttr     string   `envconfig:"LDAP_USER_EMAIL"`
		UsernameAttr  string   `envconfig:"LDAP_USER_USERNAME"`
	}
}

func LoadDotEnvFile(cfg *Config) error {
	log.Println("Loading .env File")
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to load .env file, loading defaults", err)
		os.Setenv("HOSTNAME", "localhost")
		os.Setenv("PORT", "8000")
	}
	log.Println("Attempting to process config")
	err = envconfig.Process("", cfg)
	if err != nil {
		return err
	}
	log.Println("Loading complete")
	return nil
}
