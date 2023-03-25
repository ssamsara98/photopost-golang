package lib

import (
	"log"

	"github.com/spf13/viper"
)

// Env has environment stored
type Env struct {
	ServerPort         string `mapstructure:"PORT"`
	Environment        string `mapstructure:"ENV"`
	DBUsername         string `mapstructure:"DB_USER"`
	DBPassword         string `mapstructure:"DB_PASS"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBName             string `mapstructure:"DB_NAME"`
	JWTSecret          string `mapstructure:"JWT_SECRET"`
	AWSAccessKeyId     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AWSRegion          string `mapstructure:"AWS_REGION"`
	AWSS3Bucket        string `mapstructure:"AWS_S3_BUCKET"`
}

// NewEnv creates a new environment
func NewEnv() *Env {

	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	return &env
}
